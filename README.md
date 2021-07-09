# solace-operator

This is really just an experiment with operators. It creates (and deletes) queues and REST delivery points on a Solace PubSub+ event broker using operator-sdk.

## how to use

First, you'll need a running Solace PubSub+ event broker. There are a few ways to run one, since we're in k8s land the [Solace PubSub+ Kubernetes Quickstart](https://github.com/SolaceProducts/pubsubplus-kubernetes-quickstart) is probably the easiest option if you want to run one yourself.

Next, you'll need your own docker registry - you have to build the operator yourself if you want to try it. So, build and push the image:

```bash
make docker-build docker-push IMG="registry/solace-operator:v0.1.0"
```

Or, if you want to build it for a different arch, you can do this:

```bash
make docker-buildx ARCH=arm64 IMG="registry/solace-operator:v0.1.0"

Next, you can install the operator (with the CRDs) with the Helm chart. First, you need some values:

```yaml
image:
  repository: registry/solace-operator # the repostiry used for docker iamge
  tag: v0.1.0 #  the tag used for docker image

imagePullSecrets: [] # any image pull secrets required

solaceBroker:
  sempUrl: https://solace-broker:1943
  username:
  password:
  messageVpn: default
```

Then, install the thing:

```bash
helm install test ./charts/solace-operator -f values.yaml
```

Finally, you can use the example yaml to create some queues and a REST delivery point:

```bash
kubectl apply -f config/samples/queue-and-rdp.yaml
```

Check on the broker and, if the operator actually worked, you should see two queues and an RDP. Magic!

## docs

https://www.openshift.com/blog/kubernetes-operators-best-practices

https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/

## semp api

https://docs.solace.com/API-Developer-Online-Ref-Documentation/swagger-ui/config/index.html

https://docs.solace.com/API-Developer-Online-Ref-Documentation/swagger-ui/monitor/index.html

swagger-codegen generate -i https://broker.lmbrn.co:31851/SEMP/v2/config/spec --lang go -o sempv2-config
swagger-codegen generate -i https://broker.lmbrn.co:31851/SEMP/v2/monitor/spec --lang go -o sempv2-monitor