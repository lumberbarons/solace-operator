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
	"strconv"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"

	solacev1alpha1 "github.com/lumberbarons/solace-operator/api/v1alpha1"

	semp "github.com/lumberbarons/solace-operator/sempv2-config"
	sempmon "github.com/lumberbarons/solace-operator/sempv2-monitor"

	"golang.org/x/net/context"
)

// SolaceRdpReconciler reconciles a SolaceRdp object
type SolaceRdpReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	sempClient    *semp.APIClient
	sempMonClient *sempmon.APIClient
	sempAuth      context.Context
	sempMonAuth   context.Context
	msgVpnName    string
}

//+kubebuilder:rbac:groups=solace.lmbrn.ca,resources=solacerdps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=solace.lmbrn.ca,resources=solacerdps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=solace.lmbrn.ca,resources=solacerdps/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SolaceRdp object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *SolaceRdpReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := ctrllog.FromContext(ctx)

	log.Info("Reconciling SolaceRdp")

	solacerdp := &solacev1alpha1.SolaceRdp{}
	err := r.Get(ctx, req.NamespacedName, solacerdp)
	if err != nil {
		if errors.IsNotFound(err) {
			return r.deleteRdp(ctx, req)
		}
		log.Error(err, "Failed to get SolaceRdp")
		return ctrl.Result{}, err
	}

	if ok := r.isInitialized(ctx, solacerdp); !ok {
		log.Info("Initialized SolaceRdp")

		err := r.Update(ctx, solacerdp)
		if err != nil {
			log.Error(err, "Failed to update SolaceRdp")
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	return r.reconcileCreateOrUpdate(ctx, solacerdp)
}

func (r *SolaceRdpReconciler) isInitialized(ctx context.Context, solacerdp *solacev1alpha1.SolaceRdp) bool {
	initialized := true

	if solacerdp.Spec.ClientProfile == "" {
		solacerdp.Spec.ClientProfile = "default"
		initialized = false
	}

	for idx := range solacerdp.Spec.Consumers {
		consumer := &solacerdp.Spec.Consumers[idx]
		if consumer.HttpMethod == "" {
			consumer.HttpMethod = "post"
			initialized = false
		}

		if consumer.Port == 0 {
			if consumer.TlsEnabled {
				consumer.Port = 443
			} else {
				consumer.Port = 80
			}
			initialized = false
		}

		if consumer.ConnectionCount == 0 {
			consumer.ConnectionCount = 3
			initialized = false
		}
	}

	return initialized
}

func (r *SolaceRdpReconciler) deleteRdp(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := ctrllog.FromContext(ctx)

	rdpName := req.NamespacedName.Name
	log.Info("Deleting rdp", "MsgVpnMame", r.msgVpnName, "RdpName", rdpName)
	_, _, err := r.sempClient.MsgVpnApi.DeleteMsgVpnRestDeliveryPoint(r.sempAuth, r.msgVpnName, rdpName).Execute()
	if err != nil {
		log.Error(err, "Failed to delete rdp", "MsgVpnName", r.msgVpnName, "RdpName", rdpName)
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *SolaceRdpReconciler) reconcileCreateOrUpdate(ctx context.Context, solacerdp *solacev1alpha1.SolaceRdp) (ctrl.Result, error) {
	log := ctrllog.FromContext(ctx)

	err := r.createOrUpdateRdp(ctx, solacerdp)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.reconcileConsumers(ctx, solacerdp)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.reconcileQueueBindings(ctx, solacerdp)
	if err != nil {
		return ctrl.Result{}, err
	}

	rdpName := solacerdp.ObjectMeta.Name
	sempResponse, _, err := r.sempMonClient.MsgVpnApi.GetMsgVpnRestDeliveryPoint(r.sempMonAuth, r.msgVpnName, rdpName).Execute()
	if err != nil {
		log.Error(err, "Failed to get rdp status")
		return ctrl.Result{}, err
	}

	solacerdp.Status.OperationalState = "down"
	if *sempResponse.Data.Up {
		solacerdp.Status.OperationalState = "up"
	}

	solacerdp.Status.LastFailureReason = *sempResponse.Data.LastFailureReason
	solacerdp.Status.LastFailureTime = time.Unix(int64(*sempResponse.Data.LastFailureTime), 0).UTC().Format(time.RFC3339)

	err = r.Status().Update(ctx, solacerdp)
	if err != nil {
		log.Error(err, "Failed to update SolaceRdp status")
		return ctrl.Result{}, err
	}

	if solacerdp.Status.OperationalState == "down" {
		// if down then check again in 10 seconds
		return ctrl.Result{RequeueAfter: time.Second * 10}, nil
	}

	return ctrl.Result{}, nil
}

func (r *SolaceRdpReconciler) reconcileQueueBindings(ctx context.Context, solacerdp *solacev1alpha1.SolaceRdp) error {
	log := ctrllog.FromContext(ctx)

	rdpName := solacerdp.ObjectMeta.Name

	sempResponse, _, err := r.sempClient.MsgVpnApi.GetMsgVpnRestDeliveryPointQueueBindings(r.sempAuth,
		r.msgVpnName, rdpName).Execute()

	if err != nil {
		return solaceError(ctx, err, "Failed to get rdp queue bindings",
			"MsgVpnName", r.msgVpnName, "RdpName", rdpName)
	}

	var exQueueBindings []string
	for idx := range *sempResponse.Data {
		rdpQueueBinding := (*sempResponse.Data)[idx]
		exQueueBindings = append(exQueueBindings, *rdpQueueBinding.QueueBindingName)
	}

	var queueBindings []string
	for _, queueBinding := range solacerdp.Spec.QueueBindings {
		err = r.createOrUpdateQueueBinding(ctx, solacerdp, queueBinding)
		if err != nil {
			return err
		}
		queueBindings = append(queueBindings, queueBinding.QueueName)
	}

	for _, exQueueBinding := range exQueueBindings {
		if !contains(queueBindings, exQueueBinding) {
			log.Info("Deleting queue binding", "MsgVpnMame", r.msgVpnName, "QueueName", exQueueBinding)
			_, _, err := r.sempClient.MsgVpnApi.DeleteMsgVpnRestDeliveryPointQueueBinding(r.sempAuth,
				r.msgVpnName, rdpName, url.PathEscape(exQueueBinding)).Execute()
			if err != nil {
				return solaceError(ctx, err, "Failed to delete queue binding",
					"MsgVpnName", r.msgVpnName, "QueueName", exQueueBinding)
			}
		}
	}

	return nil
}

func (r *SolaceRdpReconciler) reconcileConsumers(ctx context.Context, solacerdp *solacev1alpha1.SolaceRdp) error {
	log := ctrllog.FromContext(ctx)

	rdpName := solacerdp.ObjectMeta.Name

	sempResponse, _, err := r.sempClient.MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumers(r.sempAuth,
		r.msgVpnName, rdpName).Execute()

	if err != nil {
		return solaceError(ctx, err, "Failed to get rdp consumers",
			"MsgVpnName", r.msgVpnName, "RdpName", rdpName)
	}

	var exConsumers []string
	for _, rdpConsumer := range *sempResponse.Data {
		rdpConsumerName := strings.ToLower(*rdpConsumer.HttpMethod) + "-" +
			*rdpConsumer.RemoteHost + "-" + strconv.Itoa(int(*rdpConsumer.RemotePort))
		exConsumers = append(exConsumers, rdpConsumerName)
	}

	var consumers []string
	for _, consumer := range solacerdp.Spec.Consumers {
		rdpConsumerName := strings.ToLower(consumer.HttpMethod) + "-" + consumer.Host + "-" + strconv.Itoa(consumer.Port)
		err = r.createOrUpdateRdpConsumer(ctx, solacerdp, consumer)
		if err != nil {
			return err
		}
		consumers = append(consumers, rdpConsumerName)
	}

	for _, exConsumer := range exConsumers {
		if !contains(consumers, exConsumer) {
			log.Info("Deleting rdp consumer", "MsgVpnMame", r.msgVpnName, "RdpConsumerName", exConsumer)
			_, _, err := r.sempClient.MsgVpnApi.DeleteMsgVpnRestDeliveryPointRestConsumer(r.sempAuth,
				r.msgVpnName, rdpName, url.PathEscape(exConsumer)).Execute()
			if err != nil {
				return solaceError(ctx, err, "Failed to delete rdp consumer",
					"MsgVpnName", r.msgVpnName, "RdpConsumerName", exConsumer)
			}
		}
	}

	return nil
}

func (r *SolaceRdpReconciler) createOrUpdateRdpConsumer(ctx context.Context, solacerdp *solacev1alpha1.SolaceRdp,
	consumer solacev1alpha1.SolaceRdpConsumer) error {
	log := ctrllog.FromContext(ctx)

	rdpName := solacerdp.ObjectMeta.Name
	rdpHttpMethod := strings.ToLower(consumer.HttpMethod)
	rdpConsumerPort := int64(consumer.Port)
	rdpConnCount := int32(consumer.ConnectionCount)
	rdpEnabled := true

	rdpConsumerName := rdpHttpMethod + "-" + consumer.Host + "-" + strconv.Itoa(consumer.Port)

	sempResponse, httpResponse, _ := r.sempClient.MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumer(r.sempAuth,
		r.msgVpnName, rdpName, rdpConsumerName).Execute()

	if httpResponse.StatusCode != 200 {
		rdpConsumer := semp.MsgVpnRestDeliveryPointRestConsumer{RestConsumerName: &rdpConsumerName,
			Enabled: &rdpEnabled, RemoteHost: &consumer.Host, RemotePort: &rdpConsumerPort,
			TlsEnabled: &consumer.TlsEnabled, HttpMethod: &rdpHttpMethod, OutgoingConnectionCount: &rdpConnCount}

		log.Info("Creating rdp consumer", "RdpName", rdpName, "RdpConsumer", rdpConsumer)
		_, _, err := r.sempClient.RestDeliveryPointApi.CreateMsgVpnRestDeliveryPointRestConsumer(
			r.sempAuth, r.msgVpnName, rdpName).Body(rdpConsumer).Execute()
		if err != nil {
			return solaceError(ctx, err, "Failed to create rdp consumer", "MsgVpnName", r.msgVpnName,
				"RdpName", rdpName, "RdpConsumerName", rdpConsumerName)
		}
	} else {
		rdpConsumerResp := sempResponse.GetData()

		if *rdpConsumerResp.OutgoingConnectionCount != rdpConnCount || *rdpConsumerResp.TlsEnabled != consumer.TlsEnabled {
			rdpEnabled = false

			rdpConsumer := semp.MsgVpnRestDeliveryPointRestConsumer{RestConsumerName: &rdpConsumerName, Enabled: &rdpEnabled}

			log.Info("Disabling rdp consumer", "RdpName", rdpName, "RdpConsumer", rdpConsumer)
			_, _, err := r.sempClient.RestDeliveryPointApi.UpdateMsgVpnRestDeliveryPointRestConsumer(
				r.sempAuth, r.msgVpnName, rdpName, rdpConsumerName).Body(rdpConsumer).Execute()
			if err != nil {
				return err
			}

			rdpEnabled = true

			rdpConsumer = semp.MsgVpnRestDeliveryPointRestConsumer{RestConsumerName: &rdpConsumerName,
				Enabled: &rdpEnabled, OutgoingConnectionCount: &rdpConnCount, TlsEnabled: &consumer.TlsEnabled}

			log.Info("Updating rdp consumer", "RdpName", rdpName, "RdpConsumer", rdpConsumer)
			_, _, err = r.sempClient.RestDeliveryPointApi.UpdateMsgVpnRestDeliveryPointRestConsumer(
				r.sempAuth, r.msgVpnName, rdpName, rdpConsumerName).Body(rdpConsumer).Execute()
			if err != nil {
				return solaceError(ctx, err, "Failed to update rdp consumer", "MsgVpnName", r.msgVpnName,
					"RdpName", rdpName, "RdpConsumerName", rdpConsumerName)
			}
		}
	}

	return nil
}

func (r *SolaceRdpReconciler) createOrUpdateRdp(ctx context.Context, solacerdp *solacev1alpha1.SolaceRdp) error {
	log := ctrllog.FromContext(ctx)

	rdpName := solacerdp.ObjectMeta.Name
	rdpEnabled := true

	sempResponse, httpResponse, _ := r.sempClient.MsgVpnApi.GetMsgVpnRestDeliveryPoint(r.sempAuth, r.msgVpnName, rdpName).Execute()

	if httpResponse.StatusCode != 200 {
		rdp := semp.MsgVpnRestDeliveryPoint{RestDeliveryPointName: &rdpName,
			ClientProfileName: &solacerdp.Spec.ClientProfile, Enabled: &rdpEnabled}

		log.Info("Creating rdp", "MsgVpnName", r.msgVpnName, "Rdp", rdp)
		_, _, err := r.sempClient.RestDeliveryPointApi.CreateMsgVpnRestDeliveryPoint(r.sempAuth, r.msgVpnName).Body(rdp).Execute()
		if err != nil {
			return solaceError(ctx, err, "Failed to create rdp", "MsgVpnName", r.msgVpnName, "RdpName", rdpName)
		}
	} else {
		rdpResp := sempResponse.GetData()

		if *rdpResp.ClientProfileName != solacerdp.Spec.ClientProfile {
			rdpEnabled = false

			rdp := semp.MsgVpnRestDeliveryPoint{RestDeliveryPointName: &rdpName, Enabled: &rdpEnabled}

			log.Info("Disabling rdp", "MsgVpnName", r.msgVpnName, "Rdp", rdp)

			_, _, err := r.sempClient.RestDeliveryPointApi.UpdateMsgVpnRestDeliveryPoint(r.sempAuth, r.msgVpnName, rdpName).Body(rdp).Execute()
			if err != nil {
				return solaceError(ctx, err, "Failed to disable rdp", "MsgVpnName", r.msgVpnName, "RdpName", rdpName)
			}

			rdpEnabled = true

			rdp = semp.MsgVpnRestDeliveryPoint{RestDeliveryPointName: &rdpName,
				ClientProfileName: &solacerdp.Spec.ClientProfile, Enabled: &rdpEnabled}

			log.Info("Updating rdp", "MsgVpnName", r.msgVpnName, "Rdp", rdp)

			_, _, err = r.sempClient.RestDeliveryPointApi.UpdateMsgVpnRestDeliveryPoint(r.sempAuth, r.msgVpnName, rdpName).Body(rdp).Execute()
			if err != nil {
				return solaceError(ctx, err, "Failed to update rdp", "MsgVpnName", r.msgVpnName, "RdpName", rdpName)
			}
		}
	}

	return nil
}

func (r *SolaceRdpReconciler) createOrUpdateQueueBinding(ctx context.Context,
	solacerdp *solacev1alpha1.SolaceRdp, queueBinding solacev1alpha1.SolaceRdpQueueBinding) error {
	log := ctrllog.FromContext(ctx)

	rdpName := solacerdp.ObjectMeta.Name
	queueName := queueBinding.QueueName
	postTarget := queueBinding.PostRequestTarget

	sempResponse, httpResponse, _ := r.sempClient.MsgVpnApi.GetMsgVpnRestDeliveryPointQueueBinding(
		r.sempAuth, r.msgVpnName, rdpName, url.PathEscape(queueName)).Execute()

	rdpQueueBinding := semp.MsgVpnRestDeliveryPointQueueBinding{QueueBindingName: &queueName,
		PostRequestTarget: &postTarget}

	if httpResponse.StatusCode != 200 {
		log.Info("Creating rdp queue binding", "RdpName", rdpName,
			"QueueName", queueName, "PostTarget", postTarget)

		_, _, err := r.sempClient.RestDeliveryPointApi.CreateMsgVpnRestDeliveryPointQueueBinding(
			r.sempAuth, r.msgVpnName, rdpName).Body(rdpQueueBinding).Execute()
		if err != nil {
			return solaceError(ctx, err, "Failed to create queue binding", "MsgVpnName",
				r.msgVpnName, "RdpName", rdpName, "QueueName", queueName)
		}
	} else {
		if *sempResponse.Data.PostRequestTarget != postTarget {
			log.Info("Updating rdp queue binding", "RdpName", rdpName,
				"QueueName", queueName, "PostTarget", postTarget)

			_, _, err := r.sempClient.RestDeliveryPointApi.UpdateMsgVpnRestDeliveryPointQueueBinding(
				r.sempAuth, r.msgVpnName, rdpName, queueName).Body(rdpQueueBinding).Execute()
			if err != nil {
				return solaceError(ctx, err, "Failed to update queue binding", "MsgVpnName",
					r.msgVpnName, "RdpName", rdpName, "QueueName", queueName)
			}
		}
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SolaceRdpReconciler) SetupWithManager(mgr ctrl.Manager) error {
	cfg := semp.NewConfiguration()
	cfg.Scheme = "https"
	cfg.Host = os.Getenv("SEMP_HOST")
	r.sempClient = semp.NewAPIClient(cfg)

	r.sempAuth = context.WithValue(context.Background(), semp.ContextBasicAuth, semp.BasicAuth{
		UserName: os.Getenv("SEMP_USERNAME"),
		Password: os.Getenv("SEMP_PASSWORD"),
	})

	cfgMon := sempmon.NewConfiguration()
	cfgMon.Scheme = "https"
	cfgMon.Host = os.Getenv("SEMP_HOST")
	r.sempMonClient = sempmon.NewAPIClient(cfgMon)

	r.sempMonAuth = context.WithValue(context.Background(), sempmon.ContextBasicAuth, sempmon.BasicAuth{
		UserName: os.Getenv("SEMP_USERNAME"),
		Password: os.Getenv("SEMP_PASSWORD"),
	})

	r.msgVpnName = os.Getenv("SEMP_MSGVPN")

	return ctrl.NewControllerManagedBy(mgr).
		For(&solacev1alpha1.SolaceRdp{}).
		Complete(r)
}
