# solace-operator

## docs

https://www.openshift.com/blog/kubernetes-operators-best-practices

https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/

## semp api

swagger-codegen generate -i https://broker.lmbrn.co:31851/SEMP/v2/config/spec --lang go -o sempv2-config
swagger-codegen generate -i https://broker.lmbrn.co:31851/SEMP/v2/monitor/spec --lang go -o sempv2-monitor