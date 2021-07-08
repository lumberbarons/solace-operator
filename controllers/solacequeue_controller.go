/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"net/url"
	"os"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"

	solacev1alpha1 "github.com/lumberbarons/solace-operator/api/v1alpha1"
	semp "github.com/lumberbarons/solace-operator/sempv2-config"
	"golang.org/x/net/context"
)

// SolaceQueueReconciler reconciles a SolaceQueue object
type SolaceQueueReconciler struct {
	client.Client
	Scheme     *runtime.Scheme
	sempClient *semp.APIClient
	sempAuth   context.Context
	msgVpnName string
}

//+kubebuilder:rbac:groups=solace.lmbrn.ca,resources=solacequeues,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=solace.lmbrn.ca,resources=solacequeues/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=solace.lmbrn.ca,resources=solacequeues/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SolaceQueue object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *SolaceQueueReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := ctrllog.FromContext(ctx)

	log.Info("Reconciling SolaceQueue")

	solacequeue := &solacev1alpha1.SolaceQueue{}
	err := r.Get(ctx, req.NamespacedName, solacequeue)
	if err != nil {
		if errors.IsNotFound(err) {
			return r.deleteQueue(ctx, req)
		}
		log.Error(err, "Failed to get SolaceQueue")
		return ctrl.Result{}, err
	}

	if ok := r.isInitialized(ctx, solacequeue); !ok {
		log.Info("Initialized SolaceQueue")

		err := r.Update(ctx, solacequeue)
		if err != nil {
			log.Error(err, "Failed to update SolaceQueue")
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	err = r.createOrUpdateQueue(ctx, solacequeue)
	if err != nil {
		return ctrl.Result{}, err
	}

	solacequeue.Status.Ready = true

	err = r.Status().Update(ctx, solacequeue)
	if err != nil {
		log.Error(err, "Failed to update SolaceQueue status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *SolaceQueueReconciler) isInitialized(ctx context.Context, solacequeue *solacev1alpha1.SolaceQueue) bool {
	initialized := true

	if solacequeue.Spec.AccessType == "" {
		solacequeue.Spec.AccessType = "exclusive"
		initialized = false
	}

	if solacequeue.Spec.NonOwnerPermission == "" {
		solacequeue.Spec.NonOwnerPermission = "consume"
		initialized = false
	}

	return initialized
}

func (r *SolaceQueueReconciler) deleteQueue(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := ctrllog.FromContext(ctx)

	queueName := req.NamespacedName.Name
	log.Info("Deleting queue", "MsgVpnMame", r.msgVpnName, "QueueName", queueName)
	_, _, err := r.sempClient.MsgVpnApi.DeleteMsgVpnQueue(r.sempAuth, r.msgVpnName, queueName)
	if err != nil {
		log.Error(err, "Failed to delete queue", "MsgVpnName", r.msgVpnName, "QueueName", queueName)
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *SolaceQueueReconciler) createOrUpdateQueue(ctx context.Context, solacequeue *solacev1alpha1.SolaceQueue) error {
	log := ctrllog.FromContext(ctx)

	queueName := solacequeue.ObjectMeta.Name
	topicSubs := solacequeue.Spec.TopicSubscriptions

	_, httpResponse, _ := r.sempClient.MsgVpnApi.GetMsgVpnQueue(r.sempAuth, r.msgVpnName, queueName, nil)

	if httpResponse.StatusCode != 200 {
		queue := semp.MsgVpnQueue{QueueName: queueName, IngressEnabled: true, EgressEnabled: true,
			Permission: solacequeue.Spec.NonOwnerPermission, Owner: solacequeue.Spec.Owner,
			AccessType: solacequeue.Spec.AccessType}
		log.Info("Creating queue", "MsgVpnName", r.msgVpnName, "QueueName", queueName)
		_, _, err := r.sempClient.MsgVpnApi.CreateMsgVpnQueue(r.sempAuth, queue, r.msgVpnName, nil)
		if err != nil {
			log.Error(err, "Failed to create queue", "MsgVpnName", r.msgVpnName, "QueueName", queueName)
			return err
		}
	} else {
		/* log.Info("Updating queue", "MsgVpnName", r.msgVpnName, "QueueName", queueName)
		_, _, err := r.sempClient.MsgVpnApi.UpdateMsgVpnQueue(r.sempAuth, queue, r.msgVpnName, queueName, nil)
		if err != nil {
			log.Error(err, "Failed to update queue", "MsgVpnName", r.msgVpnName, "QueueName", queueName)
			return err
		} */
	}

	sempResponse, _, err := r.sempClient.MsgVpnApi.GetMsgVpnQueueSubscriptions(r.sempAuth,
		r.msgVpnName, queueName, nil)

	if err != nil {
		log.Error(err, "Failed to get queue topic subscriptions", "MsgVpnName", r.msgVpnName, "QueueName", queueName)
		return err
	}

	var exTopicSubs []string
	for idx := range sempResponse.Data {
		exTopicSub := sempResponse.Data[idx]
		exTopicSubs = append(exTopicSubs, exTopicSub.SubscriptionTopic)
	}

	for _, topicSub := range topicSubs {
		if !contains(exTopicSubs, topicSub) {
			log.Info("Creating topic subcription on queue", "MsgVpnName", r.msgVpnName,
				"TopicSubscription", topicSub, "QueueName", queueName)
			_, _, err := r.sempClient.MsgVpnApi.CreateMsgVpnQueueSubscription(r.sempAuth,
				semp.MsgVpnQueueSubscription{SubscriptionTopic: topicSub}, r.msgVpnName, queueName, nil)
			if err != nil {
				log.Error(err, "Failed to create queue topic subscription", "MsgVpnName",
					r.msgVpnName, "QueueName", queueName, "TopicSubscription", topicSub)
				return err
			}
		} else {
			/* log.Info("Skip topic subcription on queue, already exists", "MsgVpnName", r.msgVpnName,
			"TopicSubscription", topicSub, "QueueName", queueName) */
		}
	}

	for _, exTopicSub := range exTopicSubs {
		if !contains(topicSubs, exTopicSub) {
			log.Info("Deleting topic subcription from queue", "MsgVpnName", r.msgVpnName,
				"TopicSubscription", exTopicSub, "QueueName", queueName)
			_, _, err := r.sempClient.MsgVpnApi.DeleteMsgVpnQueueSubscription(r.sempAuth, r.msgVpnName, queueName, url.PathEscape(exTopicSub))
			if err != nil {
				log.Error(err, "Failed to delete queue topic subscription", "MsgVpnName",
					r.msgVpnName, "QueueName", queueName, "TopicSubscription", exTopicSub)
				return err
			}
		}
	}

	return nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// SetupWithManager sets up the controller with the Manager.
func (r *SolaceQueueReconciler) SetupWithManager(mgr ctrl.Manager) error {
	cfg := semp.NewConfiguration()
	cfg.BasePath = os.Getenv("SEMP_URL") + "/SEMP/v2/config"
	r.sempClient = semp.NewAPIClient(cfg)

	r.sempAuth = context.WithValue(context.Background(), semp.ContextBasicAuth, semp.BasicAuth{
		UserName: os.Getenv("SEMP_USERNAME"),
		Password: os.Getenv("SEMP_PASSWORD"),
	})

	r.msgVpnName = os.Getenv("SEMP_MSGVPN")

	return ctrl.NewControllerManagedBy(mgr).
		For(&solacev1alpha1.SolaceQueue{}).
		Complete(r)
}
