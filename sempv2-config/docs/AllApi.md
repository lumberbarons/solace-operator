# \AllApi

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertAuthority**](AllApi.md#CreateCertAuthority) | **Post** /certAuthorities | Create a Certificate Authority object.
[**CreateCertAuthorityOcspTlsTrustedCommonName**](AllApi.md#CreateCertAuthorityOcspTlsTrustedCommonName) | **Post** /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames | Create an OCSP Responder Trusted Common Name object.
[**CreateClientCertAuthority**](AllApi.md#CreateClientCertAuthority) | **Post** /clientCertAuthorities | Create a Client Certificate Authority object.
[**CreateClientCertAuthorityOcspTlsTrustedCommonName**](AllApi.md#CreateClientCertAuthorityOcspTlsTrustedCommonName) | **Post** /clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames | Create an OCSP Responder Trusted Common Name object.
[**CreateDmrCluster**](AllApi.md#CreateDmrCluster) | **Post** /dmrClusters | Create a Cluster object.
[**CreateDmrClusterLink**](AllApi.md#CreateDmrClusterLink) | **Post** /dmrClusters/{dmrClusterName}/links | Create a Link object.
[**CreateDmrClusterLinkRemoteAddress**](AllApi.md#CreateDmrClusterLinkRemoteAddress) | **Post** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses | Create a Remote Address object.
[**CreateDmrClusterLinkTlsTrustedCommonName**](AllApi.md#CreateDmrClusterLinkTlsTrustedCommonName) | **Post** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames | Create a Trusted Common Name object.
[**CreateDomainCertAuthority**](AllApi.md#CreateDomainCertAuthority) | **Post** /domainCertAuthorities | Create a Domain Certificate Authority object.
[**CreateMsgVpn**](AllApi.md#CreateMsgVpn) | **Post** /msgVpns | Create a Message VPN object.
[**CreateMsgVpnAclProfile**](AllApi.md#CreateMsgVpnAclProfile) | **Post** /msgVpns/{msgVpnName}/aclProfiles | Create an ACL Profile object.
[**CreateMsgVpnAclProfileClientConnectException**](AllApi.md#CreateMsgVpnAclProfileClientConnectException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions | Create a Client Connect Exception object.
[**CreateMsgVpnAclProfilePublishException**](AllApi.md#CreateMsgVpnAclProfilePublishException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions | Create a Publish Topic Exception object.
[**CreateMsgVpnAclProfilePublishTopicException**](AllApi.md#CreateMsgVpnAclProfilePublishTopicException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions | Create a Publish Topic Exception object.
[**CreateMsgVpnAclProfileSubscribeException**](AllApi.md#CreateMsgVpnAclProfileSubscribeException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions | Create a Subscribe Topic Exception object.
[**CreateMsgVpnAclProfileSubscribeShareNameException**](AllApi.md#CreateMsgVpnAclProfileSubscribeShareNameException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions | Create a Subscribe Share Name Exception object.
[**CreateMsgVpnAclProfileSubscribeTopicException**](AllApi.md#CreateMsgVpnAclProfileSubscribeTopicException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions | Create a Subscribe Topic Exception object.
[**CreateMsgVpnAuthenticationOauthProvider**](AllApi.md#CreateMsgVpnAuthenticationOauthProvider) | **Post** /msgVpns/{msgVpnName}/authenticationOauthProviders | Create an OAuth Provider object.
[**CreateMsgVpnAuthorizationGroup**](AllApi.md#CreateMsgVpnAuthorizationGroup) | **Post** /msgVpns/{msgVpnName}/authorizationGroups | Create an LDAP Authorization Group object.
[**CreateMsgVpnBridge**](AllApi.md#CreateMsgVpnBridge) | **Post** /msgVpns/{msgVpnName}/bridges | Create a Bridge object.
[**CreateMsgVpnBridgeRemoteMsgVpn**](AllApi.md#CreateMsgVpnBridgeRemoteMsgVpn) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns | Create a Remote Message VPN object.
[**CreateMsgVpnBridgeRemoteSubscription**](AllApi.md#CreateMsgVpnBridgeRemoteSubscription) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions | Create a Remote Subscription object.
[**CreateMsgVpnBridgeTlsTrustedCommonName**](AllApi.md#CreateMsgVpnBridgeTlsTrustedCommonName) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames | Create a Trusted Common Name object.
[**CreateMsgVpnClientProfile**](AllApi.md#CreateMsgVpnClientProfile) | **Post** /msgVpns/{msgVpnName}/clientProfiles | Create a Client Profile object.
[**CreateMsgVpnClientUsername**](AllApi.md#CreateMsgVpnClientUsername) | **Post** /msgVpns/{msgVpnName}/clientUsernames | Create a Client Username object.
[**CreateMsgVpnDistributedCache**](AllApi.md#CreateMsgVpnDistributedCache) | **Post** /msgVpns/{msgVpnName}/distributedCaches | Create a Distributed Cache object.
[**CreateMsgVpnDistributedCacheCluster**](AllApi.md#CreateMsgVpnDistributedCacheCluster) | **Post** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters | Create a Cache Cluster object.
[**CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](AllApi.md#CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster) | **Post** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters | Create a Home Cache Cluster object.
[**CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](AllApi.md#CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix) | **Post** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes | Create a Topic Prefix object.
[**CreateMsgVpnDistributedCacheClusterInstance**](AllApi.md#CreateMsgVpnDistributedCacheClusterInstance) | **Post** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances | Create a Cache Instance object.
[**CreateMsgVpnDistributedCacheClusterTopic**](AllApi.md#CreateMsgVpnDistributedCacheClusterTopic) | **Post** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics | Create a Topic object.
[**CreateMsgVpnDmrBridge**](AllApi.md#CreateMsgVpnDmrBridge) | **Post** /msgVpns/{msgVpnName}/dmrBridges | Create a DMR Bridge object.
[**CreateMsgVpnJndiConnectionFactory**](AllApi.md#CreateMsgVpnJndiConnectionFactory) | **Post** /msgVpns/{msgVpnName}/jndiConnectionFactories | Create a JNDI Connection Factory object.
[**CreateMsgVpnJndiQueue**](AllApi.md#CreateMsgVpnJndiQueue) | **Post** /msgVpns/{msgVpnName}/jndiQueues | Create a JNDI Queue object.
[**CreateMsgVpnJndiTopic**](AllApi.md#CreateMsgVpnJndiTopic) | **Post** /msgVpns/{msgVpnName}/jndiTopics | Create a JNDI Topic object.
[**CreateMsgVpnMqttRetainCache**](AllApi.md#CreateMsgVpnMqttRetainCache) | **Post** /msgVpns/{msgVpnName}/mqttRetainCaches | Create an MQTT Retain Cache object.
[**CreateMsgVpnMqttSession**](AllApi.md#CreateMsgVpnMqttSession) | **Post** /msgVpns/{msgVpnName}/mqttSessions | Create an MQTT Session object.
[**CreateMsgVpnMqttSessionSubscription**](AllApi.md#CreateMsgVpnMqttSessionSubscription) | **Post** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions | Create a Subscription object.
[**CreateMsgVpnQueue**](AllApi.md#CreateMsgVpnQueue) | **Post** /msgVpns/{msgVpnName}/queues | Create a Queue object.
[**CreateMsgVpnQueueSubscription**](AllApi.md#CreateMsgVpnQueueSubscription) | **Post** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions | Create a Queue Subscription object.
[**CreateMsgVpnQueueTemplate**](AllApi.md#CreateMsgVpnQueueTemplate) | **Post** /msgVpns/{msgVpnName}/queueTemplates | Create a Queue Template object.
[**CreateMsgVpnReplayLog**](AllApi.md#CreateMsgVpnReplayLog) | **Post** /msgVpns/{msgVpnName}/replayLogs | Create a Replay Log object.
[**CreateMsgVpnReplicatedTopic**](AllApi.md#CreateMsgVpnReplicatedTopic) | **Post** /msgVpns/{msgVpnName}/replicatedTopics | Create a Replicated Topic object.
[**CreateMsgVpnRestDeliveryPoint**](AllApi.md#CreateMsgVpnRestDeliveryPoint) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints | Create a REST Delivery Point object.
[**CreateMsgVpnRestDeliveryPointQueueBinding**](AllApi.md#CreateMsgVpnRestDeliveryPointQueueBinding) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings | Create a Queue Binding object.
[**CreateMsgVpnRestDeliveryPointRestConsumer**](AllApi.md#CreateMsgVpnRestDeliveryPointRestConsumer) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Create a REST Consumer object.
[**CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**](AllApi.md#CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims | Create a Claim object.
[**CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**](AllApi.md#CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames | Create a Trusted Common Name object.
[**CreateMsgVpnSequencedTopic**](AllApi.md#CreateMsgVpnSequencedTopic) | **Post** /msgVpns/{msgVpnName}/sequencedTopics | Create a Sequenced Topic object.
[**CreateMsgVpnTopicEndpoint**](AllApi.md#CreateMsgVpnTopicEndpoint) | **Post** /msgVpns/{msgVpnName}/topicEndpoints | Create a Topic Endpoint object.
[**CreateMsgVpnTopicEndpointTemplate**](AllApi.md#CreateMsgVpnTopicEndpointTemplate) | **Post** /msgVpns/{msgVpnName}/topicEndpointTemplates | Create a Topic Endpoint Template object.
[**CreateVirtualHostname**](AllApi.md#CreateVirtualHostname) | **Post** /virtualHostnames | Create a Virtual Hostname object.
[**DeleteCertAuthority**](AllApi.md#DeleteCertAuthority) | **Delete** /certAuthorities/{certAuthorityName} | Delete a Certificate Authority object.
[**DeleteCertAuthorityOcspTlsTrustedCommonName**](AllApi.md#DeleteCertAuthorityOcspTlsTrustedCommonName) | **Delete** /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName} | Delete an OCSP Responder Trusted Common Name object.
[**DeleteClientCertAuthority**](AllApi.md#DeleteClientCertAuthority) | **Delete** /clientCertAuthorities/{certAuthorityName} | Delete a Client Certificate Authority object.
[**DeleteClientCertAuthorityOcspTlsTrustedCommonName**](AllApi.md#DeleteClientCertAuthorityOcspTlsTrustedCommonName) | **Delete** /clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName} | Delete an OCSP Responder Trusted Common Name object.
[**DeleteDmrCluster**](AllApi.md#DeleteDmrCluster) | **Delete** /dmrClusters/{dmrClusterName} | Delete a Cluster object.
[**DeleteDmrClusterLink**](AllApi.md#DeleteDmrClusterLink) | **Delete** /dmrClusters/{dmrClusterName}/links/{remoteNodeName} | Delete a Link object.
[**DeleteDmrClusterLinkRemoteAddress**](AllApi.md#DeleteDmrClusterLinkRemoteAddress) | **Delete** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses/{remoteAddress} | Delete a Remote Address object.
[**DeleteDmrClusterLinkTlsTrustedCommonName**](AllApi.md#DeleteDmrClusterLinkTlsTrustedCommonName) | **Delete** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Delete a Trusted Common Name object.
[**DeleteDomainCertAuthority**](AllApi.md#DeleteDomainCertAuthority) | **Delete** /domainCertAuthorities/{certAuthorityName} | Delete a Domain Certificate Authority object.
[**DeleteMsgVpn**](AllApi.md#DeleteMsgVpn) | **Delete** /msgVpns/{msgVpnName} | Delete a Message VPN object.
[**DeleteMsgVpnAclProfile**](AllApi.md#DeleteMsgVpnAclProfile) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Delete an ACL Profile object.
[**DeleteMsgVpnAclProfileClientConnectException**](AllApi.md#DeleteMsgVpnAclProfileClientConnectException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress} | Delete a Client Connect Exception object.
[**DeleteMsgVpnAclProfilePublishException**](AllApi.md#DeleteMsgVpnAclProfilePublishException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic} | Delete a Publish Topic Exception object.
[**DeleteMsgVpnAclProfilePublishTopicException**](AllApi.md#DeleteMsgVpnAclProfilePublishTopicException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions/{publishTopicExceptionSyntax},{publishTopicException} | Delete a Publish Topic Exception object.
[**DeleteMsgVpnAclProfileSubscribeException**](AllApi.md#DeleteMsgVpnAclProfileSubscribeException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic} | Delete a Subscribe Topic Exception object.
[**DeleteMsgVpnAclProfileSubscribeShareNameException**](AllApi.md#DeleteMsgVpnAclProfileSubscribeShareNameException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions/{subscribeShareNameExceptionSyntax},{subscribeShareNameException} | Delete a Subscribe Share Name Exception object.
[**DeleteMsgVpnAclProfileSubscribeTopicException**](AllApi.md#DeleteMsgVpnAclProfileSubscribeTopicException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions/{subscribeTopicExceptionSyntax},{subscribeTopicException} | Delete a Subscribe Topic Exception object.
[**DeleteMsgVpnAuthenticationOauthProvider**](AllApi.md#DeleteMsgVpnAuthenticationOauthProvider) | **Delete** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Delete an OAuth Provider object.
[**DeleteMsgVpnAuthorizationGroup**](AllApi.md#DeleteMsgVpnAuthorizationGroup) | **Delete** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Delete an LDAP Authorization Group object.
[**DeleteMsgVpnBridge**](AllApi.md#DeleteMsgVpnBridge) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Delete a Bridge object.
[**DeleteMsgVpnBridgeRemoteMsgVpn**](AllApi.md#DeleteMsgVpnBridgeRemoteMsgVpn) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Delete a Remote Message VPN object.
[**DeleteMsgVpnBridgeRemoteSubscription**](AllApi.md#DeleteMsgVpnBridgeRemoteSubscription) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic} | Delete a Remote Subscription object.
[**DeleteMsgVpnBridgeTlsTrustedCommonName**](AllApi.md#DeleteMsgVpnBridgeTlsTrustedCommonName) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Delete a Trusted Common Name object.
[**DeleteMsgVpnClientProfile**](AllApi.md#DeleteMsgVpnClientProfile) | **Delete** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Delete a Client Profile object.
[**DeleteMsgVpnClientUsername**](AllApi.md#DeleteMsgVpnClientUsername) | **Delete** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Delete a Client Username object.
[**DeleteMsgVpnDistributedCache**](AllApi.md#DeleteMsgVpnDistributedCache) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Delete a Distributed Cache object.
[**DeleteMsgVpnDistributedCacheCluster**](AllApi.md#DeleteMsgVpnDistributedCacheCluster) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Delete a Cache Cluster object.
[**DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](AllApi.md#DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName} | Delete a Home Cache Cluster object.
[**DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](AllApi.md#DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix} | Delete a Topic Prefix object.
[**DeleteMsgVpnDistributedCacheClusterInstance**](AllApi.md#DeleteMsgVpnDistributedCacheClusterInstance) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Delete a Cache Instance object.
[**DeleteMsgVpnDistributedCacheClusterTopic**](AllApi.md#DeleteMsgVpnDistributedCacheClusterTopic) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics/{topic} | Delete a Topic object.
[**DeleteMsgVpnDmrBridge**](AllApi.md#DeleteMsgVpnDmrBridge) | **Delete** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Delete a DMR Bridge object.
[**DeleteMsgVpnJndiConnectionFactory**](AllApi.md#DeleteMsgVpnJndiConnectionFactory) | **Delete** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Delete a JNDI Connection Factory object.
[**DeleteMsgVpnJndiQueue**](AllApi.md#DeleteMsgVpnJndiQueue) | **Delete** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Delete a JNDI Queue object.
[**DeleteMsgVpnJndiTopic**](AllApi.md#DeleteMsgVpnJndiTopic) | **Delete** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Delete a JNDI Topic object.
[**DeleteMsgVpnMqttRetainCache**](AllApi.md#DeleteMsgVpnMqttRetainCache) | **Delete** /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName} | Delete an MQTT Retain Cache object.
[**DeleteMsgVpnMqttSession**](AllApi.md#DeleteMsgVpnMqttSession) | **Delete** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Delete an MQTT Session object.
[**DeleteMsgVpnMqttSessionSubscription**](AllApi.md#DeleteMsgVpnMqttSessionSubscription) | **Delete** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Delete a Subscription object.
[**DeleteMsgVpnQueue**](AllApi.md#DeleteMsgVpnQueue) | **Delete** /msgVpns/{msgVpnName}/queues/{queueName} | Delete a Queue object.
[**DeleteMsgVpnQueueSubscription**](AllApi.md#DeleteMsgVpnQueueSubscription) | **Delete** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic} | Delete a Queue Subscription object.
[**DeleteMsgVpnQueueTemplate**](AllApi.md#DeleteMsgVpnQueueTemplate) | **Delete** /msgVpns/{msgVpnName}/queueTemplates/{queueTemplateName} | Delete a Queue Template object.
[**DeleteMsgVpnReplayLog**](AllApi.md#DeleteMsgVpnReplayLog) | **Delete** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Delete a Replay Log object.
[**DeleteMsgVpnReplicatedTopic**](AllApi.md#DeleteMsgVpnReplicatedTopic) | **Delete** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Delete a Replicated Topic object.
[**DeleteMsgVpnRestDeliveryPoint**](AllApi.md#DeleteMsgVpnRestDeliveryPoint) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Delete a REST Delivery Point object.
[**DeleteMsgVpnRestDeliveryPointQueueBinding**](AllApi.md#DeleteMsgVpnRestDeliveryPointQueueBinding) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Delete a Queue Binding object.
[**DeleteMsgVpnRestDeliveryPointRestConsumer**](AllApi.md#DeleteMsgVpnRestDeliveryPointRestConsumer) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Delete a REST Consumer object.
[**DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**](AllApi.md#DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims/{oauthJwtClaimName} | Delete a Claim object.
[**DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**](AllApi.md#DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Delete a Trusted Common Name object.
[**DeleteMsgVpnSequencedTopic**](AllApi.md#DeleteMsgVpnSequencedTopic) | **Delete** /msgVpns/{msgVpnName}/sequencedTopics/{sequencedTopic} | Delete a Sequenced Topic object.
[**DeleteMsgVpnTopicEndpoint**](AllApi.md#DeleteMsgVpnTopicEndpoint) | **Delete** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Delete a Topic Endpoint object.
[**DeleteMsgVpnTopicEndpointTemplate**](AllApi.md#DeleteMsgVpnTopicEndpointTemplate) | **Delete** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Delete a Topic Endpoint Template object.
[**DeleteVirtualHostname**](AllApi.md#DeleteVirtualHostname) | **Delete** /virtualHostnames/{virtualHostname} | Delete a Virtual Hostname object.
[**GetAbout**](AllApi.md#GetAbout) | **Get** /about | Get an About object.
[**GetAboutApi**](AllApi.md#GetAboutApi) | **Get** /about/api | Get an API Description object.
[**GetAboutUser**](AllApi.md#GetAboutUser) | **Get** /about/user | Get a User object.
[**GetAboutUserMsgVpn**](AllApi.md#GetAboutUserMsgVpn) | **Get** /about/user/msgVpns/{msgVpnName} | Get a User Message VPN object.
[**GetAboutUserMsgVpns**](AllApi.md#GetAboutUserMsgVpns) | **Get** /about/user/msgVpns | Get a list of User Message VPN objects.
[**GetBroker**](AllApi.md#GetBroker) | **Get** / | Get a Broker object.
[**GetCertAuthorities**](AllApi.md#GetCertAuthorities) | **Get** /certAuthorities | Get a list of Certificate Authority objects.
[**GetCertAuthority**](AllApi.md#GetCertAuthority) | **Get** /certAuthorities/{certAuthorityName} | Get a Certificate Authority object.
[**GetCertAuthorityOcspTlsTrustedCommonName**](AllApi.md#GetCertAuthorityOcspTlsTrustedCommonName) | **Get** /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName} | Get an OCSP Responder Trusted Common Name object.
[**GetCertAuthorityOcspTlsTrustedCommonNames**](AllApi.md#GetCertAuthorityOcspTlsTrustedCommonNames) | **Get** /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames | Get a list of OCSP Responder Trusted Common Name objects.
[**GetClientCertAuthorities**](AllApi.md#GetClientCertAuthorities) | **Get** /clientCertAuthorities | Get a list of Client Certificate Authority objects.
[**GetClientCertAuthority**](AllApi.md#GetClientCertAuthority) | **Get** /clientCertAuthorities/{certAuthorityName} | Get a Client Certificate Authority object.
[**GetClientCertAuthorityOcspTlsTrustedCommonName**](AllApi.md#GetClientCertAuthorityOcspTlsTrustedCommonName) | **Get** /clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName} | Get an OCSP Responder Trusted Common Name object.
[**GetClientCertAuthorityOcspTlsTrustedCommonNames**](AllApi.md#GetClientCertAuthorityOcspTlsTrustedCommonNames) | **Get** /clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames | Get a list of OCSP Responder Trusted Common Name objects.
[**GetDmrCluster**](AllApi.md#GetDmrCluster) | **Get** /dmrClusters/{dmrClusterName} | Get a Cluster object.
[**GetDmrClusterLink**](AllApi.md#GetDmrClusterLink) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName} | Get a Link object.
[**GetDmrClusterLinkRemoteAddress**](AllApi.md#GetDmrClusterLinkRemoteAddress) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses/{remoteAddress} | Get a Remote Address object.
[**GetDmrClusterLinkRemoteAddresses**](AllApi.md#GetDmrClusterLinkRemoteAddresses) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses | Get a list of Remote Address objects.
[**GetDmrClusterLinkTlsTrustedCommonName**](AllApi.md#GetDmrClusterLinkTlsTrustedCommonName) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
[**GetDmrClusterLinkTlsTrustedCommonNames**](AllApi.md#GetDmrClusterLinkTlsTrustedCommonNames) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
[**GetDmrClusterLinks**](AllApi.md#GetDmrClusterLinks) | **Get** /dmrClusters/{dmrClusterName}/links | Get a list of Link objects.
[**GetDmrClusters**](AllApi.md#GetDmrClusters) | **Get** /dmrClusters | Get a list of Cluster objects.
[**GetDomainCertAuthorities**](AllApi.md#GetDomainCertAuthorities) | **Get** /domainCertAuthorities | Get a list of Domain Certificate Authority objects.
[**GetDomainCertAuthority**](AllApi.md#GetDomainCertAuthority) | **Get** /domainCertAuthorities/{certAuthorityName} | Get a Domain Certificate Authority object.
[**GetMsgVpn**](AllApi.md#GetMsgVpn) | **Get** /msgVpns/{msgVpnName} | Get a Message VPN object.
[**GetMsgVpnAclProfile**](AllApi.md#GetMsgVpnAclProfile) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Get an ACL Profile object.
[**GetMsgVpnAclProfileClientConnectException**](AllApi.md#GetMsgVpnAclProfileClientConnectException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress} | Get a Client Connect Exception object.
[**GetMsgVpnAclProfileClientConnectExceptions**](AllApi.md#GetMsgVpnAclProfileClientConnectExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions | Get a list of Client Connect Exception objects.
[**GetMsgVpnAclProfilePublishException**](AllApi.md#GetMsgVpnAclProfilePublishException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic} | Get a Publish Topic Exception object.
[**GetMsgVpnAclProfilePublishExceptions**](AllApi.md#GetMsgVpnAclProfilePublishExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions | Get a list of Publish Topic Exception objects.
[**GetMsgVpnAclProfilePublishTopicException**](AllApi.md#GetMsgVpnAclProfilePublishTopicException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions/{publishTopicExceptionSyntax},{publishTopicException} | Get a Publish Topic Exception object.
[**GetMsgVpnAclProfilePublishTopicExceptions**](AllApi.md#GetMsgVpnAclProfilePublishTopicExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions | Get a list of Publish Topic Exception objects.
[**GetMsgVpnAclProfileSubscribeException**](AllApi.md#GetMsgVpnAclProfileSubscribeException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic} | Get a Subscribe Topic Exception object.
[**GetMsgVpnAclProfileSubscribeExceptions**](AllApi.md#GetMsgVpnAclProfileSubscribeExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions | Get a list of Subscribe Topic Exception objects.
[**GetMsgVpnAclProfileSubscribeShareNameException**](AllApi.md#GetMsgVpnAclProfileSubscribeShareNameException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions/{subscribeShareNameExceptionSyntax},{subscribeShareNameException} | Get a Subscribe Share Name Exception object.
[**GetMsgVpnAclProfileSubscribeShareNameExceptions**](AllApi.md#GetMsgVpnAclProfileSubscribeShareNameExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions | Get a list of Subscribe Share Name Exception objects.
[**GetMsgVpnAclProfileSubscribeTopicException**](AllApi.md#GetMsgVpnAclProfileSubscribeTopicException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions/{subscribeTopicExceptionSyntax},{subscribeTopicException} | Get a Subscribe Topic Exception object.
[**GetMsgVpnAclProfileSubscribeTopicExceptions**](AllApi.md#GetMsgVpnAclProfileSubscribeTopicExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions | Get a list of Subscribe Topic Exception objects.
[**GetMsgVpnAclProfiles**](AllApi.md#GetMsgVpnAclProfiles) | **Get** /msgVpns/{msgVpnName}/aclProfiles | Get a list of ACL Profile objects.
[**GetMsgVpnAuthenticationOauthProvider**](AllApi.md#GetMsgVpnAuthenticationOauthProvider) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Get an OAuth Provider object.
[**GetMsgVpnAuthenticationOauthProviders**](AllApi.md#GetMsgVpnAuthenticationOauthProviders) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders | Get a list of OAuth Provider objects.
[**GetMsgVpnAuthorizationGroup**](AllApi.md#GetMsgVpnAuthorizationGroup) | **Get** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Get an LDAP Authorization Group object.
[**GetMsgVpnAuthorizationGroups**](AllApi.md#GetMsgVpnAuthorizationGroups) | **Get** /msgVpns/{msgVpnName}/authorizationGroups | Get a list of LDAP Authorization Group objects.
[**GetMsgVpnBridge**](AllApi.md#GetMsgVpnBridge) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Get a Bridge object.
[**GetMsgVpnBridgeRemoteMsgVpn**](AllApi.md#GetMsgVpnBridgeRemoteMsgVpn) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Get a Remote Message VPN object.
[**GetMsgVpnBridgeRemoteMsgVpns**](AllApi.md#GetMsgVpnBridgeRemoteMsgVpns) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns | Get a list of Remote Message VPN objects.
[**GetMsgVpnBridgeRemoteSubscription**](AllApi.md#GetMsgVpnBridgeRemoteSubscription) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic} | Get a Remote Subscription object.
[**GetMsgVpnBridgeRemoteSubscriptions**](AllApi.md#GetMsgVpnBridgeRemoteSubscriptions) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions | Get a list of Remote Subscription objects.
[**GetMsgVpnBridgeTlsTrustedCommonName**](AllApi.md#GetMsgVpnBridgeTlsTrustedCommonName) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
[**GetMsgVpnBridgeTlsTrustedCommonNames**](AllApi.md#GetMsgVpnBridgeTlsTrustedCommonNames) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
[**GetMsgVpnBridges**](AllApi.md#GetMsgVpnBridges) | **Get** /msgVpns/{msgVpnName}/bridges | Get a list of Bridge objects.
[**GetMsgVpnClientProfile**](AllApi.md#GetMsgVpnClientProfile) | **Get** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Get a Client Profile object.
[**GetMsgVpnClientProfiles**](AllApi.md#GetMsgVpnClientProfiles) | **Get** /msgVpns/{msgVpnName}/clientProfiles | Get a list of Client Profile objects.
[**GetMsgVpnClientUsername**](AllApi.md#GetMsgVpnClientUsername) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Get a Client Username object.
[**GetMsgVpnClientUsernames**](AllApi.md#GetMsgVpnClientUsernames) | **Get** /msgVpns/{msgVpnName}/clientUsernames | Get a list of Client Username objects.
[**GetMsgVpnDistributedCache**](AllApi.md#GetMsgVpnDistributedCache) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Get a Distributed Cache object.
[**GetMsgVpnDistributedCacheCluster**](AllApi.md#GetMsgVpnDistributedCacheCluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Get a Cache Cluster object.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](AllApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName} | Get a Home Cache Cluster object.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](AllApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix} | Get a Topic Prefix object.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes**](AllApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes | Get a list of Topic Prefix objects.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters**](AllApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters | Get a list of Home Cache Cluster objects.
[**GetMsgVpnDistributedCacheClusterInstance**](AllApi.md#GetMsgVpnDistributedCacheClusterInstance) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Get a Cache Instance object.
[**GetMsgVpnDistributedCacheClusterInstances**](AllApi.md#GetMsgVpnDistributedCacheClusterInstances) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances | Get a list of Cache Instance objects.
[**GetMsgVpnDistributedCacheClusterTopic**](AllApi.md#GetMsgVpnDistributedCacheClusterTopic) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics/{topic} | Get a Topic object.
[**GetMsgVpnDistributedCacheClusterTopics**](AllApi.md#GetMsgVpnDistributedCacheClusterTopics) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics | Get a list of Topic objects.
[**GetMsgVpnDistributedCacheClusters**](AllApi.md#GetMsgVpnDistributedCacheClusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters | Get a list of Cache Cluster objects.
[**GetMsgVpnDistributedCaches**](AllApi.md#GetMsgVpnDistributedCaches) | **Get** /msgVpns/{msgVpnName}/distributedCaches | Get a list of Distributed Cache objects.
[**GetMsgVpnDmrBridge**](AllApi.md#GetMsgVpnDmrBridge) | **Get** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Get a DMR Bridge object.
[**GetMsgVpnDmrBridges**](AllApi.md#GetMsgVpnDmrBridges) | **Get** /msgVpns/{msgVpnName}/dmrBridges | Get a list of DMR Bridge objects.
[**GetMsgVpnJndiConnectionFactories**](AllApi.md#GetMsgVpnJndiConnectionFactories) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories | Get a list of JNDI Connection Factory objects.
[**GetMsgVpnJndiConnectionFactory**](AllApi.md#GetMsgVpnJndiConnectionFactory) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Get a JNDI Connection Factory object.
[**GetMsgVpnJndiQueue**](AllApi.md#GetMsgVpnJndiQueue) | **Get** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Get a JNDI Queue object.
[**GetMsgVpnJndiQueues**](AllApi.md#GetMsgVpnJndiQueues) | **Get** /msgVpns/{msgVpnName}/jndiQueues | Get a list of JNDI Queue objects.
[**GetMsgVpnJndiTopic**](AllApi.md#GetMsgVpnJndiTopic) | **Get** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Get a JNDI Topic object.
[**GetMsgVpnJndiTopics**](AllApi.md#GetMsgVpnJndiTopics) | **Get** /msgVpns/{msgVpnName}/jndiTopics | Get a list of JNDI Topic objects.
[**GetMsgVpnMqttRetainCache**](AllApi.md#GetMsgVpnMqttRetainCache) | **Get** /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName} | Get an MQTT Retain Cache object.
[**GetMsgVpnMqttRetainCaches**](AllApi.md#GetMsgVpnMqttRetainCaches) | **Get** /msgVpns/{msgVpnName}/mqttRetainCaches | Get a list of MQTT Retain Cache objects.
[**GetMsgVpnMqttSession**](AllApi.md#GetMsgVpnMqttSession) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Get an MQTT Session object.
[**GetMsgVpnMqttSessionSubscription**](AllApi.md#GetMsgVpnMqttSessionSubscription) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Get a Subscription object.
[**GetMsgVpnMqttSessionSubscriptions**](AllApi.md#GetMsgVpnMqttSessionSubscriptions) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions | Get a list of Subscription objects.
[**GetMsgVpnMqttSessions**](AllApi.md#GetMsgVpnMqttSessions) | **Get** /msgVpns/{msgVpnName}/mqttSessions | Get a list of MQTT Session objects.
[**GetMsgVpnQueue**](AllApi.md#GetMsgVpnQueue) | **Get** /msgVpns/{msgVpnName}/queues/{queueName} | Get a Queue object.
[**GetMsgVpnQueueSubscription**](AllApi.md#GetMsgVpnQueueSubscription) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic} | Get a Queue Subscription object.
[**GetMsgVpnQueueSubscriptions**](AllApi.md#GetMsgVpnQueueSubscriptions) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions | Get a list of Queue Subscription objects.
[**GetMsgVpnQueueTemplate**](AllApi.md#GetMsgVpnQueueTemplate) | **Get** /msgVpns/{msgVpnName}/queueTemplates/{queueTemplateName} | Get a Queue Template object.
[**GetMsgVpnQueueTemplates**](AllApi.md#GetMsgVpnQueueTemplates) | **Get** /msgVpns/{msgVpnName}/queueTemplates | Get a list of Queue Template objects.
[**GetMsgVpnQueues**](AllApi.md#GetMsgVpnQueues) | **Get** /msgVpns/{msgVpnName}/queues | Get a list of Queue objects.
[**GetMsgVpnReplayLog**](AllApi.md#GetMsgVpnReplayLog) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Get a Replay Log object.
[**GetMsgVpnReplayLogs**](AllApi.md#GetMsgVpnReplayLogs) | **Get** /msgVpns/{msgVpnName}/replayLogs | Get a list of Replay Log objects.
[**GetMsgVpnReplicatedTopic**](AllApi.md#GetMsgVpnReplicatedTopic) | **Get** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Get a Replicated Topic object.
[**GetMsgVpnReplicatedTopics**](AllApi.md#GetMsgVpnReplicatedTopics) | **Get** /msgVpns/{msgVpnName}/replicatedTopics | Get a list of Replicated Topic objects.
[**GetMsgVpnRestDeliveryPoint**](AllApi.md#GetMsgVpnRestDeliveryPoint) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Get a REST Delivery Point object.
[**GetMsgVpnRestDeliveryPointQueueBinding**](AllApi.md#GetMsgVpnRestDeliveryPointQueueBinding) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Get a Queue Binding object.
[**GetMsgVpnRestDeliveryPointQueueBindings**](AllApi.md#GetMsgVpnRestDeliveryPointQueueBindings) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings | Get a list of Queue Binding objects.
[**GetMsgVpnRestDeliveryPointRestConsumer**](AllApi.md#GetMsgVpnRestDeliveryPointRestConsumer) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Get a REST Consumer object.
[**GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**](AllApi.md#GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims/{oauthJwtClaimName} | Get a Claim object.
[**GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims**](AllApi.md#GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims | Get a list of Claim objects.
[**GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**](AllApi.md#GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
[**GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames**](AllApi.md#GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
[**GetMsgVpnRestDeliveryPointRestConsumers**](AllApi.md#GetMsgVpnRestDeliveryPointRestConsumers) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Get a list of REST Consumer objects.
[**GetMsgVpnRestDeliveryPoints**](AllApi.md#GetMsgVpnRestDeliveryPoints) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints | Get a list of REST Delivery Point objects.
[**GetMsgVpnSequencedTopic**](AllApi.md#GetMsgVpnSequencedTopic) | **Get** /msgVpns/{msgVpnName}/sequencedTopics/{sequencedTopic} | Get a Sequenced Topic object.
[**GetMsgVpnSequencedTopics**](AllApi.md#GetMsgVpnSequencedTopics) | **Get** /msgVpns/{msgVpnName}/sequencedTopics | Get a list of Sequenced Topic objects.
[**GetMsgVpnTopicEndpoint**](AllApi.md#GetMsgVpnTopicEndpoint) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Get a Topic Endpoint object.
[**GetMsgVpnTopicEndpointTemplate**](AllApi.md#GetMsgVpnTopicEndpointTemplate) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Get a Topic Endpoint Template object.
[**GetMsgVpnTopicEndpointTemplates**](AllApi.md#GetMsgVpnTopicEndpointTemplates) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates | Get a list of Topic Endpoint Template objects.
[**GetMsgVpnTopicEndpoints**](AllApi.md#GetMsgVpnTopicEndpoints) | **Get** /msgVpns/{msgVpnName}/topicEndpoints | Get a list of Topic Endpoint objects.
[**GetMsgVpns**](AllApi.md#GetMsgVpns) | **Get** /msgVpns | Get a list of Message VPN objects.
[**GetSystemInformation**](AllApi.md#GetSystemInformation) | **Get** /systemInformation | Get a System Information object.
[**GetVirtualHostname**](AllApi.md#GetVirtualHostname) | **Get** /virtualHostnames/{virtualHostname} | Get a Virtual Hostname object.
[**GetVirtualHostnames**](AllApi.md#GetVirtualHostnames) | **Get** /virtualHostnames | Get a list of Virtual Hostname objects.
[**ReplaceCertAuthority**](AllApi.md#ReplaceCertAuthority) | **Put** /certAuthorities/{certAuthorityName} | Replace a Certificate Authority object.
[**ReplaceClientCertAuthority**](AllApi.md#ReplaceClientCertAuthority) | **Put** /clientCertAuthorities/{certAuthorityName} | Replace a Client Certificate Authority object.
[**ReplaceDmrCluster**](AllApi.md#ReplaceDmrCluster) | **Put** /dmrClusters/{dmrClusterName} | Replace a Cluster object.
[**ReplaceDmrClusterLink**](AllApi.md#ReplaceDmrClusterLink) | **Put** /dmrClusters/{dmrClusterName}/links/{remoteNodeName} | Replace a Link object.
[**ReplaceDomainCertAuthority**](AllApi.md#ReplaceDomainCertAuthority) | **Put** /domainCertAuthorities/{certAuthorityName} | Replace a Domain Certificate Authority object.
[**ReplaceMsgVpn**](AllApi.md#ReplaceMsgVpn) | **Put** /msgVpns/{msgVpnName} | Replace a Message VPN object.
[**ReplaceMsgVpnAclProfile**](AllApi.md#ReplaceMsgVpnAclProfile) | **Put** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Replace an ACL Profile object.
[**ReplaceMsgVpnAuthenticationOauthProvider**](AllApi.md#ReplaceMsgVpnAuthenticationOauthProvider) | **Put** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Replace an OAuth Provider object.
[**ReplaceMsgVpnAuthorizationGroup**](AllApi.md#ReplaceMsgVpnAuthorizationGroup) | **Put** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Replace an LDAP Authorization Group object.
[**ReplaceMsgVpnBridge**](AllApi.md#ReplaceMsgVpnBridge) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Replace a Bridge object.
[**ReplaceMsgVpnBridgeRemoteMsgVpn**](AllApi.md#ReplaceMsgVpnBridgeRemoteMsgVpn) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Replace a Remote Message VPN object.
[**ReplaceMsgVpnClientProfile**](AllApi.md#ReplaceMsgVpnClientProfile) | **Put** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Replace a Client Profile object.
[**ReplaceMsgVpnClientUsername**](AllApi.md#ReplaceMsgVpnClientUsername) | **Put** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Replace a Client Username object.
[**ReplaceMsgVpnDistributedCache**](AllApi.md#ReplaceMsgVpnDistributedCache) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Replace a Distributed Cache object.
[**ReplaceMsgVpnDistributedCacheCluster**](AllApi.md#ReplaceMsgVpnDistributedCacheCluster) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Replace a Cache Cluster object.
[**ReplaceMsgVpnDistributedCacheClusterInstance**](AllApi.md#ReplaceMsgVpnDistributedCacheClusterInstance) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Replace a Cache Instance object.
[**ReplaceMsgVpnDmrBridge**](AllApi.md#ReplaceMsgVpnDmrBridge) | **Put** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Replace a DMR Bridge object.
[**ReplaceMsgVpnJndiConnectionFactory**](AllApi.md#ReplaceMsgVpnJndiConnectionFactory) | **Put** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Replace a JNDI Connection Factory object.
[**ReplaceMsgVpnJndiQueue**](AllApi.md#ReplaceMsgVpnJndiQueue) | **Put** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Replace a JNDI Queue object.
[**ReplaceMsgVpnJndiTopic**](AllApi.md#ReplaceMsgVpnJndiTopic) | **Put** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Replace a JNDI Topic object.
[**ReplaceMsgVpnMqttRetainCache**](AllApi.md#ReplaceMsgVpnMqttRetainCache) | **Put** /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName} | Replace an MQTT Retain Cache object.
[**ReplaceMsgVpnMqttSession**](AllApi.md#ReplaceMsgVpnMqttSession) | **Put** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Replace an MQTT Session object.
[**ReplaceMsgVpnMqttSessionSubscription**](AllApi.md#ReplaceMsgVpnMqttSessionSubscription) | **Put** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Replace a Subscription object.
[**ReplaceMsgVpnQueue**](AllApi.md#ReplaceMsgVpnQueue) | **Put** /msgVpns/{msgVpnName}/queues/{queueName} | Replace a Queue object.
[**ReplaceMsgVpnQueueTemplate**](AllApi.md#ReplaceMsgVpnQueueTemplate) | **Put** /msgVpns/{msgVpnName}/queueTemplates/{queueTemplateName} | Replace a Queue Template object.
[**ReplaceMsgVpnReplayLog**](AllApi.md#ReplaceMsgVpnReplayLog) | **Put** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Replace a Replay Log object.
[**ReplaceMsgVpnReplicatedTopic**](AllApi.md#ReplaceMsgVpnReplicatedTopic) | **Put** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Replace a Replicated Topic object.
[**ReplaceMsgVpnRestDeliveryPoint**](AllApi.md#ReplaceMsgVpnRestDeliveryPoint) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Replace a REST Delivery Point object.
[**ReplaceMsgVpnRestDeliveryPointQueueBinding**](AllApi.md#ReplaceMsgVpnRestDeliveryPointQueueBinding) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Replace a Queue Binding object.
[**ReplaceMsgVpnRestDeliveryPointRestConsumer**](AllApi.md#ReplaceMsgVpnRestDeliveryPointRestConsumer) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Replace a REST Consumer object.
[**ReplaceMsgVpnTopicEndpoint**](AllApi.md#ReplaceMsgVpnTopicEndpoint) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Replace a Topic Endpoint object.
[**ReplaceMsgVpnTopicEndpointTemplate**](AllApi.md#ReplaceMsgVpnTopicEndpointTemplate) | **Put** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Replace a Topic Endpoint Template object.
[**ReplaceVirtualHostname**](AllApi.md#ReplaceVirtualHostname) | **Put** /virtualHostnames/{virtualHostname} | Replace a Virtual Hostname object.
[**UpdateBroker**](AllApi.md#UpdateBroker) | **Patch** / | Update a Broker object.
[**UpdateCertAuthority**](AllApi.md#UpdateCertAuthority) | **Patch** /certAuthorities/{certAuthorityName} | Update a Certificate Authority object.
[**UpdateClientCertAuthority**](AllApi.md#UpdateClientCertAuthority) | **Patch** /clientCertAuthorities/{certAuthorityName} | Update a Client Certificate Authority object.
[**UpdateDmrCluster**](AllApi.md#UpdateDmrCluster) | **Patch** /dmrClusters/{dmrClusterName} | Update a Cluster object.
[**UpdateDmrClusterLink**](AllApi.md#UpdateDmrClusterLink) | **Patch** /dmrClusters/{dmrClusterName}/links/{remoteNodeName} | Update a Link object.
[**UpdateDomainCertAuthority**](AllApi.md#UpdateDomainCertAuthority) | **Patch** /domainCertAuthorities/{certAuthorityName} | Update a Domain Certificate Authority object.
[**UpdateMsgVpn**](AllApi.md#UpdateMsgVpn) | **Patch** /msgVpns/{msgVpnName} | Update a Message VPN object.
[**UpdateMsgVpnAclProfile**](AllApi.md#UpdateMsgVpnAclProfile) | **Patch** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Update an ACL Profile object.
[**UpdateMsgVpnAuthenticationOauthProvider**](AllApi.md#UpdateMsgVpnAuthenticationOauthProvider) | **Patch** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Update an OAuth Provider object.
[**UpdateMsgVpnAuthorizationGroup**](AllApi.md#UpdateMsgVpnAuthorizationGroup) | **Patch** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Update an LDAP Authorization Group object.
[**UpdateMsgVpnBridge**](AllApi.md#UpdateMsgVpnBridge) | **Patch** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Update a Bridge object.
[**UpdateMsgVpnBridgeRemoteMsgVpn**](AllApi.md#UpdateMsgVpnBridgeRemoteMsgVpn) | **Patch** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Update a Remote Message VPN object.
[**UpdateMsgVpnClientProfile**](AllApi.md#UpdateMsgVpnClientProfile) | **Patch** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Update a Client Profile object.
[**UpdateMsgVpnClientUsername**](AllApi.md#UpdateMsgVpnClientUsername) | **Patch** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Update a Client Username object.
[**UpdateMsgVpnDistributedCache**](AllApi.md#UpdateMsgVpnDistributedCache) | **Patch** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Update a Distributed Cache object.
[**UpdateMsgVpnDistributedCacheCluster**](AllApi.md#UpdateMsgVpnDistributedCacheCluster) | **Patch** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Update a Cache Cluster object.
[**UpdateMsgVpnDistributedCacheClusterInstance**](AllApi.md#UpdateMsgVpnDistributedCacheClusterInstance) | **Patch** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Update a Cache Instance object.
[**UpdateMsgVpnDmrBridge**](AllApi.md#UpdateMsgVpnDmrBridge) | **Patch** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Update a DMR Bridge object.
[**UpdateMsgVpnJndiConnectionFactory**](AllApi.md#UpdateMsgVpnJndiConnectionFactory) | **Patch** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Update a JNDI Connection Factory object.
[**UpdateMsgVpnJndiQueue**](AllApi.md#UpdateMsgVpnJndiQueue) | **Patch** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Update a JNDI Queue object.
[**UpdateMsgVpnJndiTopic**](AllApi.md#UpdateMsgVpnJndiTopic) | **Patch** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Update a JNDI Topic object.
[**UpdateMsgVpnMqttRetainCache**](AllApi.md#UpdateMsgVpnMqttRetainCache) | **Patch** /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName} | Update an MQTT Retain Cache object.
[**UpdateMsgVpnMqttSession**](AllApi.md#UpdateMsgVpnMqttSession) | **Patch** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Update an MQTT Session object.
[**UpdateMsgVpnMqttSessionSubscription**](AllApi.md#UpdateMsgVpnMqttSessionSubscription) | **Patch** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Update a Subscription object.
[**UpdateMsgVpnQueue**](AllApi.md#UpdateMsgVpnQueue) | **Patch** /msgVpns/{msgVpnName}/queues/{queueName} | Update a Queue object.
[**UpdateMsgVpnQueueTemplate**](AllApi.md#UpdateMsgVpnQueueTemplate) | **Patch** /msgVpns/{msgVpnName}/queueTemplates/{queueTemplateName} | Update a Queue Template object.
[**UpdateMsgVpnReplayLog**](AllApi.md#UpdateMsgVpnReplayLog) | **Patch** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Update a Replay Log object.
[**UpdateMsgVpnReplicatedTopic**](AllApi.md#UpdateMsgVpnReplicatedTopic) | **Patch** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Update a Replicated Topic object.
[**UpdateMsgVpnRestDeliveryPoint**](AllApi.md#UpdateMsgVpnRestDeliveryPoint) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Update a REST Delivery Point object.
[**UpdateMsgVpnRestDeliveryPointQueueBinding**](AllApi.md#UpdateMsgVpnRestDeliveryPointQueueBinding) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Update a Queue Binding object.
[**UpdateMsgVpnRestDeliveryPointRestConsumer**](AllApi.md#UpdateMsgVpnRestDeliveryPointRestConsumer) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Update a REST Consumer object.
[**UpdateMsgVpnTopicEndpoint**](AllApi.md#UpdateMsgVpnTopicEndpoint) | **Patch** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Update a Topic Endpoint object.
[**UpdateMsgVpnTopicEndpointTemplate**](AllApi.md#UpdateMsgVpnTopicEndpointTemplate) | **Patch** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Update a Topic Endpoint Template object.
[**UpdateVirtualHostname**](AllApi.md#UpdateVirtualHostname) | **Patch** /virtualHostnames/{virtualHostname} | Update a Virtual Hostname object.



## CreateCertAuthority

> CertAuthorityResponse CreateCertAuthority(ctx).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Certificate Authority object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewCertAuthority() // CertAuthority | The Certificate Authority object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateCertAuthority(context.Background()).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertAuthority`: CertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateCertAuthority`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CertAuthority**](CertAuthority.md) | The Certificate Authority object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityResponse**](CertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertAuthorityOcspTlsTrustedCommonName

> CertAuthorityOcspTlsTrustedCommonNameResponse CreateCertAuthorityOcspTlsTrustedCommonName(ctx, certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create an OCSP Responder Trusted Common Name object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    body := *openapiclient.NewCertAuthorityOcspTlsTrustedCommonName() // CertAuthorityOcspTlsTrustedCommonName | The OCSP Responder Trusted Common Name object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateCertAuthorityOcspTlsTrustedCommonName(context.Background(), certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateCertAuthorityOcspTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertAuthorityOcspTlsTrustedCommonName`: CertAuthorityOcspTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateCertAuthorityOcspTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertAuthorityOcspTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CertAuthorityOcspTlsTrustedCommonName**](CertAuthorityOcspTlsTrustedCommonName.md) | The OCSP Responder Trusted Common Name object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityOcspTlsTrustedCommonNameResponse**](CertAuthorityOcspTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClientCertAuthority

> ClientCertAuthorityResponse CreateClientCertAuthority(ctx).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Client Certificate Authority object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewClientCertAuthority() // ClientCertAuthority | The Client Certificate Authority object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateClientCertAuthority(context.Background()).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateClientCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateClientCertAuthority`: ClientCertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateClientCertAuthority`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClientCertAuthority**](ClientCertAuthority.md) | The Client Certificate Authority object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthorityResponse**](ClientCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClientCertAuthorityOcspTlsTrustedCommonName

> ClientCertAuthorityOcspTlsTrustedCommonNameResponse CreateClientCertAuthorityOcspTlsTrustedCommonName(ctx, certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create an OCSP Responder Trusted Common Name object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    body := *openapiclient.NewClientCertAuthorityOcspTlsTrustedCommonName() // ClientCertAuthorityOcspTlsTrustedCommonName | The OCSP Responder Trusted Common Name object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateClientCertAuthorityOcspTlsTrustedCommonName(context.Background(), certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateClientCertAuthorityOcspTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateClientCertAuthorityOcspTlsTrustedCommonName`: ClientCertAuthorityOcspTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateClientCertAuthorityOcspTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientCertAuthorityOcspTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ClientCertAuthorityOcspTlsTrustedCommonName**](ClientCertAuthorityOcspTlsTrustedCommonName.md) | The OCSP Responder Trusted Common Name object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthorityOcspTlsTrustedCommonNameResponse**](ClientCertAuthorityOcspTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDmrCluster

> DmrClusterResponse CreateDmrCluster(ctx).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Cluster object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewDmrCluster() // DmrCluster | The Cluster object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateDmrCluster(context.Background()).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateDmrCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDmrCluster`: DmrClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateDmrCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDmrClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DmrCluster**](DmrCluster.md) | The Cluster object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterResponse**](DmrClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDmrClusterLink

> DmrClusterLinkResponse CreateDmrClusterLink(ctx, dmrClusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Link object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    body := *openapiclient.NewDmrClusterLink() // DmrClusterLink | The Link object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateDmrClusterLink(context.Background(), dmrClusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateDmrClusterLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDmrClusterLink`: DmrClusterLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateDmrClusterLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDmrClusterLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DmrClusterLink**](DmrClusterLink.md) | The Link object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkResponse**](DmrClusterLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDmrClusterLinkRemoteAddress

> DmrClusterLinkRemoteAddressResponse CreateDmrClusterLinkRemoteAddress(ctx, dmrClusterName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Remote Address object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    body := *openapiclient.NewDmrClusterLinkRemoteAddress() // DmrClusterLinkRemoteAddress | The Remote Address object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateDmrClusterLinkRemoteAddress(context.Background(), dmrClusterName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateDmrClusterLinkRemoteAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDmrClusterLinkRemoteAddress`: DmrClusterLinkRemoteAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateDmrClusterLinkRemoteAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDmrClusterLinkRemoteAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DmrClusterLinkRemoteAddress**](DmrClusterLinkRemoteAddress.md) | The Remote Address object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkRemoteAddressResponse**](DmrClusterLinkRemoteAddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDmrClusterLinkTlsTrustedCommonName

> DmrClusterLinkTlsTrustedCommonNameResponse CreateDmrClusterLinkTlsTrustedCommonName(ctx, dmrClusterName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Trusted Common Name object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    body := *openapiclient.NewDmrClusterLinkTlsTrustedCommonName() // DmrClusterLinkTlsTrustedCommonName | The Trusted Common Name object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateDmrClusterLinkTlsTrustedCommonName(context.Background(), dmrClusterName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateDmrClusterLinkTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDmrClusterLinkTlsTrustedCommonName`: DmrClusterLinkTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateDmrClusterLinkTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDmrClusterLinkTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DmrClusterLinkTlsTrustedCommonName**](DmrClusterLinkTlsTrustedCommonName.md) | The Trusted Common Name object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkTlsTrustedCommonNameResponse**](DmrClusterLinkTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDomainCertAuthority

> DomainCertAuthorityResponse CreateDomainCertAuthority(ctx).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Domain Certificate Authority object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewDomainCertAuthority() // DomainCertAuthority | The Domain Certificate Authority object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateDomainCertAuthority(context.Background()).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateDomainCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDomainCertAuthority`: DomainCertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateDomainCertAuthority`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DomainCertAuthority**](DomainCertAuthority.md) | The Domain Certificate Authority object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DomainCertAuthorityResponse**](DomainCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpn

> MsgVpnResponse CreateMsgVpn(ctx).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Message VPN object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewMsgVpn() // MsgVpn | The Message VPN object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpn(context.Background()).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpn`: MsgVpnResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MsgVpn**](MsgVpn.md) | The Message VPN object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnResponse**](MsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnAclProfile

> MsgVpnAclProfileResponse CreateMsgVpnAclProfile(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create an ACL Profile object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnAclProfile() // MsgVpnAclProfile | The ACL Profile object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnAclProfile(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnAclProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnAclProfile`: MsgVpnAclProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnAclProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnAclProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnAclProfile**](MsgVpnAclProfile.md) | The ACL Profile object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileResponse**](MsgVpnAclProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnAclProfileClientConnectException

> MsgVpnAclProfileClientConnectExceptionResponse CreateMsgVpnAclProfileClientConnectException(ctx, msgVpnName, aclProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Client Connect Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    body := *openapiclient.NewMsgVpnAclProfileClientConnectException() // MsgVpnAclProfileClientConnectException | The Client Connect Exception object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnAclProfileClientConnectException(context.Background(), msgVpnName, aclProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnAclProfileClientConnectException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnAclProfileClientConnectException`: MsgVpnAclProfileClientConnectExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnAclProfileClientConnectException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnAclProfileClientConnectExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnAclProfileClientConnectException**](MsgVpnAclProfileClientConnectException.md) | The Client Connect Exception object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileClientConnectExceptionResponse**](MsgVpnAclProfileClientConnectExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnAclProfilePublishException

> MsgVpnAclProfilePublishExceptionResponse CreateMsgVpnAclProfilePublishException(ctx, msgVpnName, aclProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Publish Topic Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    body := *openapiclient.NewMsgVpnAclProfilePublishException() // MsgVpnAclProfilePublishException | The Publish Topic Exception object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnAclProfilePublishException(context.Background(), msgVpnName, aclProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnAclProfilePublishException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnAclProfilePublishException`: MsgVpnAclProfilePublishExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnAclProfilePublishException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnAclProfilePublishExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnAclProfilePublishException**](MsgVpnAclProfilePublishException.md) | The Publish Topic Exception object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishExceptionResponse**](MsgVpnAclProfilePublishExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnAclProfilePublishTopicException

> MsgVpnAclProfilePublishTopicExceptionResponse CreateMsgVpnAclProfilePublishTopicException(ctx, msgVpnName, aclProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Publish Topic Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    body := *openapiclient.NewMsgVpnAclProfilePublishTopicException() // MsgVpnAclProfilePublishTopicException | The Publish Topic Exception object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnAclProfilePublishTopicException(context.Background(), msgVpnName, aclProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnAclProfilePublishTopicException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnAclProfilePublishTopicException`: MsgVpnAclProfilePublishTopicExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnAclProfilePublishTopicException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnAclProfilePublishTopicExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnAclProfilePublishTopicException**](MsgVpnAclProfilePublishTopicException.md) | The Publish Topic Exception object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishTopicExceptionResponse**](MsgVpnAclProfilePublishTopicExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnAclProfileSubscribeException

> MsgVpnAclProfileSubscribeExceptionResponse CreateMsgVpnAclProfileSubscribeException(ctx, msgVpnName, aclProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Subscribe Topic Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    body := *openapiclient.NewMsgVpnAclProfileSubscribeException() // MsgVpnAclProfileSubscribeException | The Subscribe Topic Exception object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnAclProfileSubscribeException(context.Background(), msgVpnName, aclProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnAclProfileSubscribeException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnAclProfileSubscribeException`: MsgVpnAclProfileSubscribeExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnAclProfileSubscribeException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnAclProfileSubscribeExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnAclProfileSubscribeException**](MsgVpnAclProfileSubscribeException.md) | The Subscribe Topic Exception object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeExceptionResponse**](MsgVpnAclProfileSubscribeExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnAclProfileSubscribeShareNameException

> MsgVpnAclProfileSubscribeShareNameExceptionResponse CreateMsgVpnAclProfileSubscribeShareNameException(ctx, msgVpnName, aclProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Subscribe Share Name Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    body := *openapiclient.NewMsgVpnAclProfileSubscribeShareNameException() // MsgVpnAclProfileSubscribeShareNameException | The Subscribe Share Name Exception object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnAclProfileSubscribeShareNameException(context.Background(), msgVpnName, aclProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnAclProfileSubscribeShareNameException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnAclProfileSubscribeShareNameException`: MsgVpnAclProfileSubscribeShareNameExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnAclProfileSubscribeShareNameException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnAclProfileSubscribeShareNameExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnAclProfileSubscribeShareNameException**](MsgVpnAclProfileSubscribeShareNameException.md) | The Subscribe Share Name Exception object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeShareNameExceptionResponse**](MsgVpnAclProfileSubscribeShareNameExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnAclProfileSubscribeTopicException

> MsgVpnAclProfileSubscribeTopicExceptionResponse CreateMsgVpnAclProfileSubscribeTopicException(ctx, msgVpnName, aclProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Subscribe Topic Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    body := *openapiclient.NewMsgVpnAclProfileSubscribeTopicException() // MsgVpnAclProfileSubscribeTopicException | The Subscribe Topic Exception object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnAclProfileSubscribeTopicException(context.Background(), msgVpnName, aclProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnAclProfileSubscribeTopicException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnAclProfileSubscribeTopicException`: MsgVpnAclProfileSubscribeTopicExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnAclProfileSubscribeTopicException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnAclProfileSubscribeTopicExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnAclProfileSubscribeTopicException**](MsgVpnAclProfileSubscribeTopicException.md) | The Subscribe Topic Exception object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeTopicExceptionResponse**](MsgVpnAclProfileSubscribeTopicExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnAuthenticationOauthProvider

> MsgVpnAuthenticationOauthProviderResponse CreateMsgVpnAuthenticationOauthProvider(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create an OAuth Provider object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnAuthenticationOauthProvider() // MsgVpnAuthenticationOauthProvider | The OAuth Provider object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnAuthenticationOauthProvider(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnAuthenticationOauthProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnAuthenticationOauthProvider`: MsgVpnAuthenticationOauthProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnAuthenticationOauthProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnAuthenticationOauthProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnAuthenticationOauthProvider**](MsgVpnAuthenticationOauthProvider.md) | The OAuth Provider object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProviderResponse**](MsgVpnAuthenticationOauthProviderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnAuthorizationGroup

> MsgVpnAuthorizationGroupResponse CreateMsgVpnAuthorizationGroup(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create an LDAP Authorization Group object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnAuthorizationGroup() // MsgVpnAuthorizationGroup | The LDAP Authorization Group object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnAuthorizationGroup(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnAuthorizationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnAuthorizationGroup`: MsgVpnAuthorizationGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnAuthorizationGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnAuthorizationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnAuthorizationGroup**](MsgVpnAuthorizationGroup.md) | The LDAP Authorization Group object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupResponse**](MsgVpnAuthorizationGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnBridge

> MsgVpnBridgeResponse CreateMsgVpnBridge(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Bridge object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnBridge() // MsgVpnBridge | The Bridge object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnBridge(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnBridge`: MsgVpnBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnBridge**](MsgVpnBridge.md) | The Bridge object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnBridgeRemoteMsgVpn

> MsgVpnBridgeRemoteMsgVpnResponse CreateMsgVpnBridgeRemoteMsgVpn(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Remote Message VPN object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    body := *openapiclient.NewMsgVpnBridgeRemoteMsgVpn() // MsgVpnBridgeRemoteMsgVpn | The Remote Message VPN object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnBridgeRemoteMsgVpn(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnBridgeRemoteMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnBridgeRemoteMsgVpn`: MsgVpnBridgeRemoteMsgVpnResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnBridgeRemoteMsgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnBridgeRemoteMsgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnBridgeRemoteMsgVpn**](MsgVpnBridgeRemoteMsgVpn.md) | The Remote Message VPN object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteMsgVpnResponse**](MsgVpnBridgeRemoteMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnBridgeRemoteSubscription

> MsgVpnBridgeRemoteSubscriptionResponse CreateMsgVpnBridgeRemoteSubscription(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Remote Subscription object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    body := *openapiclient.NewMsgVpnBridgeRemoteSubscription() // MsgVpnBridgeRemoteSubscription | The Remote Subscription object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnBridgeRemoteSubscription(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnBridgeRemoteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnBridgeRemoteSubscription`: MsgVpnBridgeRemoteSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnBridgeRemoteSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnBridgeRemoteSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnBridgeRemoteSubscription**](MsgVpnBridgeRemoteSubscription.md) | The Remote Subscription object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteSubscriptionResponse**](MsgVpnBridgeRemoteSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnBridgeTlsTrustedCommonName

> MsgVpnBridgeTlsTrustedCommonNameResponse CreateMsgVpnBridgeTlsTrustedCommonName(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Trusted Common Name object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    body := *openapiclient.NewMsgVpnBridgeTlsTrustedCommonName() // MsgVpnBridgeTlsTrustedCommonName | The Trusted Common Name object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnBridgeTlsTrustedCommonName(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnBridgeTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnBridgeTlsTrustedCommonName`: MsgVpnBridgeTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnBridgeTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnBridgeTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnBridgeTlsTrustedCommonName**](MsgVpnBridgeTlsTrustedCommonName.md) | The Trusted Common Name object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeTlsTrustedCommonNameResponse**](MsgVpnBridgeTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnClientProfile

> MsgVpnClientProfileResponse CreateMsgVpnClientProfile(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Client Profile object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnClientProfile() // MsgVpnClientProfile | The Client Profile object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnClientProfile(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnClientProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnClientProfile`: MsgVpnClientProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnClientProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnClientProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnClientProfile**](MsgVpnClientProfile.md) | The Client Profile object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientProfileResponse**](MsgVpnClientProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnClientUsername

> MsgVpnClientUsernameResponse CreateMsgVpnClientUsername(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Client Username object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnClientUsername() // MsgVpnClientUsername | The Client Username object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnClientUsername(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnClientUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnClientUsername`: MsgVpnClientUsernameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnClientUsername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnClientUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnClientUsername**](MsgVpnClientUsername.md) | The Client Username object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernameResponse**](MsgVpnClientUsernameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnDistributedCache

> MsgVpnDistributedCacheResponse CreateMsgVpnDistributedCache(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Distributed Cache object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnDistributedCache() // MsgVpnDistributedCache | The Distributed Cache object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnDistributedCache(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnDistributedCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnDistributedCache`: MsgVpnDistributedCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnDistributedCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnDistributedCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnDistributedCache**](MsgVpnDistributedCache.md) | The Distributed Cache object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheResponse**](MsgVpnDistributedCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnDistributedCacheCluster

> MsgVpnDistributedCacheClusterResponse CreateMsgVpnDistributedCacheCluster(ctx, msgVpnName, cacheName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Cache Cluster object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    body := *openapiclient.NewMsgVpnDistributedCacheCluster() // MsgVpnDistributedCacheCluster | The Cache Cluster object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnDistributedCacheCluster(context.Background(), msgVpnName, cacheName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnDistributedCacheCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnDistributedCacheCluster`: MsgVpnDistributedCacheClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnDistributedCacheCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnDistributedCacheClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnDistributedCacheCluster**](MsgVpnDistributedCacheCluster.md) | The Cache Cluster object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterResponse**](MsgVpnDistributedCacheClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster

> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(ctx, msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Home Cache Cluster object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    body := *openapiclient.NewMsgVpnDistributedCacheClusterGlobalCachingHomeCluster() // MsgVpnDistributedCacheClusterGlobalCachingHomeCluster | The Home Cache Cluster object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(context.Background(), msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster`: MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](MsgVpnDistributedCacheClusterGlobalCachingHomeCluster.md) | The Home Cache Cluster object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix

> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(ctx, msgVpnName, cacheName, clusterName, homeClusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Topic Prefix object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    homeClusterName := "homeClusterName_example" // string | The name of the remote Home Cache Cluster.
    body := *openapiclient.NewMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix() // MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix | The Topic Prefix object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(context.Background(), msgVpnName, cacheName, clusterName, homeClusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix`: MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**homeClusterName** | **string** | The name of the remote Home Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix.md) | The Topic Prefix object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnDistributedCacheClusterInstance

> MsgVpnDistributedCacheClusterInstanceResponse CreateMsgVpnDistributedCacheClusterInstance(ctx, msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Cache Instance object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    body := *openapiclient.NewMsgVpnDistributedCacheClusterInstance() // MsgVpnDistributedCacheClusterInstance | The Cache Instance object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnDistributedCacheClusterInstance(context.Background(), msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnDistributedCacheClusterInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnDistributedCacheClusterInstance`: MsgVpnDistributedCacheClusterInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnDistributedCacheClusterInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnDistributedCacheClusterInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnDistributedCacheClusterInstance**](MsgVpnDistributedCacheClusterInstance.md) | The Cache Instance object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceResponse**](MsgVpnDistributedCacheClusterInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnDistributedCacheClusterTopic

> MsgVpnDistributedCacheClusterTopicResponse CreateMsgVpnDistributedCacheClusterTopic(ctx, msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Topic object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    body := *openapiclient.NewMsgVpnDistributedCacheClusterTopic() // MsgVpnDistributedCacheClusterTopic | The Topic object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnDistributedCacheClusterTopic(context.Background(), msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnDistributedCacheClusterTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnDistributedCacheClusterTopic`: MsgVpnDistributedCacheClusterTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnDistributedCacheClusterTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnDistributedCacheClusterTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnDistributedCacheClusterTopic**](MsgVpnDistributedCacheClusterTopic.md) | The Topic object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterTopicResponse**](MsgVpnDistributedCacheClusterTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnDmrBridge

> MsgVpnDmrBridgeResponse CreateMsgVpnDmrBridge(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a DMR Bridge object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnDmrBridge() // MsgVpnDmrBridge | The DMR Bridge object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnDmrBridge(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnDmrBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnDmrBridge`: MsgVpnDmrBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnDmrBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnDmrBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnDmrBridge**](MsgVpnDmrBridge.md) | The DMR Bridge object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDmrBridgeResponse**](MsgVpnDmrBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnJndiConnectionFactory

> MsgVpnJndiConnectionFactoryResponse CreateMsgVpnJndiConnectionFactory(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a JNDI Connection Factory object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnJndiConnectionFactory() // MsgVpnJndiConnectionFactory | The JNDI Connection Factory object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnJndiConnectionFactory(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnJndiConnectionFactory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnJndiConnectionFactory`: MsgVpnJndiConnectionFactoryResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnJndiConnectionFactory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnJndiConnectionFactoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnJndiConnectionFactory**](MsgVpnJndiConnectionFactory.md) | The JNDI Connection Factory object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiConnectionFactoryResponse**](MsgVpnJndiConnectionFactoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnJndiQueue

> MsgVpnJndiQueueResponse CreateMsgVpnJndiQueue(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a JNDI Queue object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnJndiQueue() // MsgVpnJndiQueue | The JNDI Queue object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnJndiQueue(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnJndiQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnJndiQueue`: MsgVpnJndiQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnJndiQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnJndiQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnJndiQueue**](MsgVpnJndiQueue.md) | The JNDI Queue object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiQueueResponse**](MsgVpnJndiQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnJndiTopic

> MsgVpnJndiTopicResponse CreateMsgVpnJndiTopic(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a JNDI Topic object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnJndiTopic() // MsgVpnJndiTopic | The JNDI Topic object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnJndiTopic(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnJndiTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnJndiTopic`: MsgVpnJndiTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnJndiTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnJndiTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnJndiTopic**](MsgVpnJndiTopic.md) | The JNDI Topic object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiTopicResponse**](MsgVpnJndiTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnMqttRetainCache

> MsgVpnMqttRetainCacheResponse CreateMsgVpnMqttRetainCache(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create an MQTT Retain Cache object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnMqttRetainCache() // MsgVpnMqttRetainCache | The MQTT Retain Cache object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnMqttRetainCache(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnMqttRetainCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnMqttRetainCache`: MsgVpnMqttRetainCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnMqttRetainCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnMqttRetainCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnMqttRetainCache**](MsgVpnMqttRetainCache.md) | The MQTT Retain Cache object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttRetainCacheResponse**](MsgVpnMqttRetainCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnMqttSession

> MsgVpnMqttSessionResponse CreateMsgVpnMqttSession(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create an MQTT Session object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnMqttSession() // MsgVpnMqttSession | The MQTT Session object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnMqttSession(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnMqttSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnMqttSession`: MsgVpnMqttSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnMqttSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnMqttSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnMqttSession**](MsgVpnMqttSession.md) | The MQTT Session object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionResponse**](MsgVpnMqttSessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnMqttSessionSubscription

> MsgVpnMqttSessionSubscriptionResponse CreateMsgVpnMqttSessionSubscription(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Subscription object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    mqttSessionClientId := "mqttSessionClientId_example" // string | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
    mqttSessionVirtualRouter := "mqttSessionVirtualRouter_example" // string | The virtual router of the MQTT Session.
    body := *openapiclient.NewMsgVpnMqttSessionSubscription() // MsgVpnMqttSessionSubscription | The Subscription object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnMqttSessionSubscription(context.Background(), msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnMqttSessionSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnMqttSessionSubscription`: MsgVpnMqttSessionSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnMqttSessionSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**mqttSessionClientId** | **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
**mqttSessionVirtualRouter** | **string** | The virtual router of the MQTT Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnMqttSessionSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnMqttSessionSubscription**](MsgVpnMqttSessionSubscription.md) | The Subscription object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionSubscriptionResponse**](MsgVpnMqttSessionSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnQueue

> MsgVpnQueueResponse CreateMsgVpnQueue(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Queue object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnQueue() // MsgVpnQueue | The Queue object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnQueue(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnQueue`: MsgVpnQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnQueue**](MsgVpnQueue.md) | The Queue object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueResponse**](MsgVpnQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnQueueSubscription

> MsgVpnQueueSubscriptionResponse CreateMsgVpnQueueSubscription(ctx, msgVpnName, queueName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Queue Subscription object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    queueName := "queueName_example" // string | The name of the Queue.
    body := *openapiclient.NewMsgVpnQueueSubscription() // MsgVpnQueueSubscription | The Queue Subscription object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnQueueSubscription(context.Background(), msgVpnName, queueName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnQueueSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnQueueSubscription`: MsgVpnQueueSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnQueueSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The name of the Queue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnQueueSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnQueueSubscription**](MsgVpnQueueSubscription.md) | The Queue Subscription object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueSubscriptionResponse**](MsgVpnQueueSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnQueueTemplate

> MsgVpnQueueTemplateResponse CreateMsgVpnQueueTemplate(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Queue Template object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnQueueTemplate() // MsgVpnQueueTemplate | The Queue Template object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnQueueTemplate(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnQueueTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnQueueTemplate`: MsgVpnQueueTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnQueueTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnQueueTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnQueueTemplate**](MsgVpnQueueTemplate.md) | The Queue Template object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueTemplateResponse**](MsgVpnQueueTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnReplayLog

> MsgVpnReplayLogResponse CreateMsgVpnReplayLog(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Replay Log object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnReplayLog() // MsgVpnReplayLog | The Replay Log object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnReplayLog(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnReplayLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnReplayLog`: MsgVpnReplayLogResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnReplayLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnReplayLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnReplayLog**](MsgVpnReplayLog.md) | The Replay Log object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogResponse**](MsgVpnReplayLogResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnReplicatedTopic

> MsgVpnReplicatedTopicResponse CreateMsgVpnReplicatedTopic(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Replicated Topic object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnReplicatedTopic() // MsgVpnReplicatedTopic | The Replicated Topic object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnReplicatedTopic(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnReplicatedTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnReplicatedTopic`: MsgVpnReplicatedTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnReplicatedTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnReplicatedTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnReplicatedTopic**](MsgVpnReplicatedTopic.md) | The Replicated Topic object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplicatedTopicResponse**](MsgVpnReplicatedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnRestDeliveryPoint

> MsgVpnRestDeliveryPointResponse CreateMsgVpnRestDeliveryPoint(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a REST Delivery Point object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnRestDeliveryPoint() // MsgVpnRestDeliveryPoint | The REST Delivery Point object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnRestDeliveryPoint(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnRestDeliveryPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnRestDeliveryPoint`: MsgVpnRestDeliveryPointResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnRestDeliveryPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnRestDeliveryPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnRestDeliveryPoint**](MsgVpnRestDeliveryPoint.md) | The REST Delivery Point object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointResponse**](MsgVpnRestDeliveryPointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnRestDeliveryPointQueueBinding

> MsgVpnRestDeliveryPointQueueBindingResponse CreateMsgVpnRestDeliveryPointQueueBinding(ctx, msgVpnName, restDeliveryPointName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Queue Binding object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    body := *openapiclient.NewMsgVpnRestDeliveryPointQueueBinding() // MsgVpnRestDeliveryPointQueueBinding | The Queue Binding object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnRestDeliveryPointQueueBinding(context.Background(), msgVpnName, restDeliveryPointName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnRestDeliveryPointQueueBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnRestDeliveryPointQueueBinding`: MsgVpnRestDeliveryPointQueueBindingResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnRestDeliveryPointQueueBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnRestDeliveryPointQueueBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnRestDeliveryPointQueueBinding**](MsgVpnRestDeliveryPointQueueBinding.md) | The Queue Binding object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingResponse**](MsgVpnRestDeliveryPointQueueBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnRestDeliveryPointRestConsumer

> MsgVpnRestDeliveryPointRestConsumerResponse CreateMsgVpnRestDeliveryPointRestConsumer(ctx, msgVpnName, restDeliveryPointName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a REST Consumer object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    body := *openapiclient.NewMsgVpnRestDeliveryPointRestConsumer() // MsgVpnRestDeliveryPointRestConsumer | The REST Consumer object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnRestDeliveryPointRestConsumer(context.Background(), msgVpnName, restDeliveryPointName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnRestDeliveryPointRestConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnRestDeliveryPointRestConsumer`: MsgVpnRestDeliveryPointRestConsumerResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnRestDeliveryPointRestConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnRestDeliveryPointRestConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnRestDeliveryPointRestConsumer**](MsgVpnRestDeliveryPointRestConsumer.md) | The REST Consumer object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerResponse**](MsgVpnRestDeliveryPointRestConsumerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim

> MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(ctx, msgVpnName, restDeliveryPointName, restConsumerName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Claim object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    restConsumerName := "restConsumerName_example" // string | The name of the REST Consumer.
    body := *openapiclient.NewMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim() // MsgVpnRestDeliveryPointRestConsumerOauthJwtClaim | The Claim object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim`: MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 
**restConsumerName** | **string** | The name of the REST Consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**](MsgVpnRestDeliveryPointRestConsumerOauthJwtClaim.md) | The Claim object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse**](MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName

> MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(ctx, msgVpnName, restDeliveryPointName, restConsumerName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Trusted Common Name object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    restConsumerName := "restConsumerName_example" // string | The name of the REST Consumer.
    body := *openapiclient.NewMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName() // MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName | The Trusted Common Name object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName`: MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 
**restConsumerName** | **string** | The name of the REST Consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**](MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName.md) | The Trusted Common Name object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse**](MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnSequencedTopic

> MsgVpnSequencedTopicResponse CreateMsgVpnSequencedTopic(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Sequenced Topic object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnSequencedTopic() // MsgVpnSequencedTopic | The Sequenced Topic object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnSequencedTopic(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnSequencedTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnSequencedTopic`: MsgVpnSequencedTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnSequencedTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnSequencedTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnSequencedTopic**](MsgVpnSequencedTopic.md) | The Sequenced Topic object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnSequencedTopicResponse**](MsgVpnSequencedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnTopicEndpoint

> MsgVpnTopicEndpointResponse CreateMsgVpnTopicEndpoint(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Topic Endpoint object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnTopicEndpoint() // MsgVpnTopicEndpoint | The Topic Endpoint object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnTopicEndpoint(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnTopicEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnTopicEndpoint`: MsgVpnTopicEndpointResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnTopicEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnTopicEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnTopicEndpoint**](MsgVpnTopicEndpoint.md) | The Topic Endpoint object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointResponse**](MsgVpnTopicEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnTopicEndpointTemplate

> MsgVpnTopicEndpointTemplateResponse CreateMsgVpnTopicEndpointTemplate(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Topic Endpoint Template object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnTopicEndpointTemplate() // MsgVpnTopicEndpointTemplate | The Topic Endpoint Template object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateMsgVpnTopicEndpointTemplate(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateMsgVpnTopicEndpointTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnTopicEndpointTemplate`: MsgVpnTopicEndpointTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateMsgVpnTopicEndpointTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnTopicEndpointTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnTopicEndpointTemplate**](MsgVpnTopicEndpointTemplate.md) | The Topic Endpoint Template object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplateResponse**](MsgVpnTopicEndpointTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVirtualHostname

> VirtualHostnameResponse CreateVirtualHostname(ctx).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Virtual Hostname object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewVirtualHostname() // VirtualHostname | The Virtual Hostname object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.CreateVirtualHostname(context.Background()).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.CreateVirtualHostname``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVirtualHostname`: VirtualHostnameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.CreateVirtualHostname`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualHostnameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VirtualHostname**](VirtualHostname.md) | The Virtual Hostname object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**VirtualHostnameResponse**](VirtualHostnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertAuthority

> SempMetaOnlyResponse DeleteCertAuthority(ctx, certAuthorityName).Execute()

Delete a Certificate Authority object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteCertAuthority(context.Background(), certAuthorityName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCertAuthority`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertAuthorityOcspTlsTrustedCommonName

> SempMetaOnlyResponse DeleteCertAuthorityOcspTlsTrustedCommonName(ctx, certAuthorityName, ocspTlsTrustedCommonName).Execute()

Delete an OCSP Responder Trusted Common Name object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    ocspTlsTrustedCommonName := "ocspTlsTrustedCommonName_example" // string | The expected Trusted Common Name of the OCSP responder remote certificate.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteCertAuthorityOcspTlsTrustedCommonName(context.Background(), certAuthorityName, ocspTlsTrustedCommonName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteCertAuthorityOcspTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCertAuthorityOcspTlsTrustedCommonName`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteCertAuthorityOcspTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 
**ocspTlsTrustedCommonName** | **string** | The expected Trusted Common Name of the OCSP responder remote certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertAuthorityOcspTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClientCertAuthority

> SempMetaOnlyResponse DeleteClientCertAuthority(ctx, certAuthorityName).Execute()

Delete a Client Certificate Authority object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteClientCertAuthority(context.Background(), certAuthorityName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteClientCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClientCertAuthority`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteClientCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClientCertAuthorityOcspTlsTrustedCommonName

> SempMetaOnlyResponse DeleteClientCertAuthorityOcspTlsTrustedCommonName(ctx, certAuthorityName, ocspTlsTrustedCommonName).Execute()

Delete an OCSP Responder Trusted Common Name object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    ocspTlsTrustedCommonName := "ocspTlsTrustedCommonName_example" // string | The expected Trusted Common Name of the OCSP responder remote certificate.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteClientCertAuthorityOcspTlsTrustedCommonName(context.Background(), certAuthorityName, ocspTlsTrustedCommonName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteClientCertAuthorityOcspTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClientCertAuthorityOcspTlsTrustedCommonName`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteClientCertAuthorityOcspTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 
**ocspTlsTrustedCommonName** | **string** | The expected Trusted Common Name of the OCSP responder remote certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientCertAuthorityOcspTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDmrCluster

> SempMetaOnlyResponse DeleteDmrCluster(ctx, dmrClusterName).Execute()

Delete a Cluster object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteDmrCluster(context.Background(), dmrClusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteDmrCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDmrCluster`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteDmrCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDmrClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDmrClusterLink

> SempMetaOnlyResponse DeleteDmrClusterLink(ctx, dmrClusterName, remoteNodeName).Execute()

Delete a Link object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteDmrClusterLink(context.Background(), dmrClusterName, remoteNodeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteDmrClusterLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDmrClusterLink`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteDmrClusterLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDmrClusterLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDmrClusterLinkRemoteAddress

> SempMetaOnlyResponse DeleteDmrClusterLinkRemoteAddress(ctx, dmrClusterName, remoteNodeName, remoteAddress).Execute()

Delete a Remote Address object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    remoteAddress := "remoteAddress_example" // string | The FQDN or IP address (and optional port) of the remote node. If a port is not provided, it will vary based on the transport encoding: 55555 (plain-text), 55443 (encrypted), or 55003 (compressed).

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteDmrClusterLinkRemoteAddress(context.Background(), dmrClusterName, remoteNodeName, remoteAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteDmrClusterLinkRemoteAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDmrClusterLinkRemoteAddress`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteDmrClusterLinkRemoteAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 
**remoteAddress** | **string** | The FQDN or IP address (and optional port) of the remote node. If a port is not provided, it will vary based on the transport encoding: 55555 (plain-text), 55443 (encrypted), or 55003 (compressed). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDmrClusterLinkRemoteAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDmrClusterLinkTlsTrustedCommonName

> SempMetaOnlyResponse DeleteDmrClusterLinkTlsTrustedCommonName(ctx, dmrClusterName, remoteNodeName, tlsTrustedCommonName).Execute()

Delete a Trusted Common Name object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    tlsTrustedCommonName := "tlsTrustedCommonName_example" // string | The expected trusted common name of the remote certificate.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteDmrClusterLinkTlsTrustedCommonName(context.Background(), dmrClusterName, remoteNodeName, tlsTrustedCommonName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteDmrClusterLinkTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDmrClusterLinkTlsTrustedCommonName`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteDmrClusterLinkTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 
**tlsTrustedCommonName** | **string** | The expected trusted common name of the remote certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDmrClusterLinkTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDomainCertAuthority

> SempMetaOnlyResponse DeleteDomainCertAuthority(ctx, certAuthorityName).Execute()

Delete a Domain Certificate Authority object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteDomainCertAuthority(context.Background(), certAuthorityName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteDomainCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDomainCertAuthority`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteDomainCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpn

> SempMetaOnlyResponse DeleteMsgVpn(ctx, msgVpnName).Execute()

Delete a Message VPN object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpn(context.Background(), msgVpnName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpn`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnAclProfile

> SempMetaOnlyResponse DeleteMsgVpnAclProfile(ctx, msgVpnName, aclProfileName).Execute()

Delete an ACL Profile object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnAclProfile(context.Background(), msgVpnName, aclProfileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnAclProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnAclProfile`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnAclProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnAclProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnAclProfileClientConnectException

> SempMetaOnlyResponse DeleteMsgVpnAclProfileClientConnectException(ctx, msgVpnName, aclProfileName, clientConnectExceptionAddress).Execute()

Delete a Client Connect Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    clientConnectExceptionAddress := "clientConnectExceptionAddress_example" // string | The IP address/netmask of the client connect exception in CIDR form.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnAclProfileClientConnectException(context.Background(), msgVpnName, aclProfileName, clientConnectExceptionAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnAclProfileClientConnectException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnAclProfileClientConnectException`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnAclProfileClientConnectException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**clientConnectExceptionAddress** | **string** | The IP address/netmask of the client connect exception in CIDR form. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnAclProfileClientConnectExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnAclProfilePublishException

> SempMetaOnlyResponse DeleteMsgVpnAclProfilePublishException(ctx, msgVpnName, aclProfileName, topicSyntax, publishExceptionTopic).Execute()

Delete a Publish Topic Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    topicSyntax := "topicSyntax_example" // string | The syntax of the topic for the exception to the default action taken.
    publishExceptionTopic := "publishExceptionTopic_example" // string | The topic for the exception to the default action taken. May include wildcard characters.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnAclProfilePublishException(context.Background(), msgVpnName, aclProfileName, topicSyntax, publishExceptionTopic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnAclProfilePublishException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnAclProfilePublishException`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnAclProfilePublishException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**topicSyntax** | **string** | The syntax of the topic for the exception to the default action taken. | 
**publishExceptionTopic** | **string** | The topic for the exception to the default action taken. May include wildcard characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnAclProfilePublishExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnAclProfilePublishTopicException

> SempMetaOnlyResponse DeleteMsgVpnAclProfilePublishTopicException(ctx, msgVpnName, aclProfileName, publishTopicExceptionSyntax, publishTopicException).Execute()

Delete a Publish Topic Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    publishTopicExceptionSyntax := "publishTopicExceptionSyntax_example" // string | The syntax of the topic for the exception to the default action taken.
    publishTopicException := "publishTopicException_example" // string | The topic for the exception to the default action taken. May include wildcard characters.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnAclProfilePublishTopicException(context.Background(), msgVpnName, aclProfileName, publishTopicExceptionSyntax, publishTopicException).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnAclProfilePublishTopicException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnAclProfilePublishTopicException`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnAclProfilePublishTopicException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**publishTopicExceptionSyntax** | **string** | The syntax of the topic for the exception to the default action taken. | 
**publishTopicException** | **string** | The topic for the exception to the default action taken. May include wildcard characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnAclProfilePublishTopicExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnAclProfileSubscribeException

> SempMetaOnlyResponse DeleteMsgVpnAclProfileSubscribeException(ctx, msgVpnName, aclProfileName, topicSyntax, subscribeExceptionTopic).Execute()

Delete a Subscribe Topic Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    topicSyntax := "topicSyntax_example" // string | The syntax of the topic for the exception to the default action taken.
    subscribeExceptionTopic := "subscribeExceptionTopic_example" // string | The topic for the exception to the default action taken. May include wildcard characters.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnAclProfileSubscribeException(context.Background(), msgVpnName, aclProfileName, topicSyntax, subscribeExceptionTopic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnAclProfileSubscribeException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnAclProfileSubscribeException`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnAclProfileSubscribeException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**topicSyntax** | **string** | The syntax of the topic for the exception to the default action taken. | 
**subscribeExceptionTopic** | **string** | The topic for the exception to the default action taken. May include wildcard characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnAclProfileSubscribeExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnAclProfileSubscribeShareNameException

> SempMetaOnlyResponse DeleteMsgVpnAclProfileSubscribeShareNameException(ctx, msgVpnName, aclProfileName, subscribeShareNameExceptionSyntax, subscribeShareNameException).Execute()

Delete a Subscribe Share Name Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    subscribeShareNameExceptionSyntax := "subscribeShareNameExceptionSyntax_example" // string | The syntax of the subscribe share name for the exception to the default action taken.
    subscribeShareNameException := "subscribeShareNameException_example" // string | The subscribe share name exception to the default action taken. May include wildcard characters.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnAclProfileSubscribeShareNameException(context.Background(), msgVpnName, aclProfileName, subscribeShareNameExceptionSyntax, subscribeShareNameException).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnAclProfileSubscribeShareNameException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnAclProfileSubscribeShareNameException`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnAclProfileSubscribeShareNameException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**subscribeShareNameExceptionSyntax** | **string** | The syntax of the subscribe share name for the exception to the default action taken. | 
**subscribeShareNameException** | **string** | The subscribe share name exception to the default action taken. May include wildcard characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnAclProfileSubscribeShareNameExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnAclProfileSubscribeTopicException

> SempMetaOnlyResponse DeleteMsgVpnAclProfileSubscribeTopicException(ctx, msgVpnName, aclProfileName, subscribeTopicExceptionSyntax, subscribeTopicException).Execute()

Delete a Subscribe Topic Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    subscribeTopicExceptionSyntax := "subscribeTopicExceptionSyntax_example" // string | The syntax of the topic for the exception to the default action taken.
    subscribeTopicException := "subscribeTopicException_example" // string | The topic for the exception to the default action taken. May include wildcard characters.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnAclProfileSubscribeTopicException(context.Background(), msgVpnName, aclProfileName, subscribeTopicExceptionSyntax, subscribeTopicException).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnAclProfileSubscribeTopicException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnAclProfileSubscribeTopicException`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnAclProfileSubscribeTopicException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**subscribeTopicExceptionSyntax** | **string** | The syntax of the topic for the exception to the default action taken. | 
**subscribeTopicException** | **string** | The topic for the exception to the default action taken. May include wildcard characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnAclProfileSubscribeTopicExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnAuthenticationOauthProvider

> SempMetaOnlyResponse DeleteMsgVpnAuthenticationOauthProvider(ctx, msgVpnName, oauthProviderName).Execute()

Delete an OAuth Provider object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    oauthProviderName := "oauthProviderName_example" // string | The name of the OAuth Provider.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnAuthenticationOauthProvider(context.Background(), msgVpnName, oauthProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnAuthenticationOauthProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnAuthenticationOauthProvider`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnAuthenticationOauthProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**oauthProviderName** | **string** | The name of the OAuth Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnAuthenticationOauthProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnAuthorizationGroup

> SempMetaOnlyResponse DeleteMsgVpnAuthorizationGroup(ctx, msgVpnName, authorizationGroupName).Execute()

Delete an LDAP Authorization Group object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    authorizationGroupName := "authorizationGroupName_example" // string | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as '#', '+', ';', '=' as the value of the group name returned from the LDAP server might prepend those characters with '\\'. For example a group name called 'test#,lab,com' will be returned from the LDAP server as 'test\\#,lab,com'.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnAuthorizationGroup(context.Background(), msgVpnName, authorizationGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnAuthorizationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnAuthorizationGroup`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnAuthorizationGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**authorizationGroupName** | **string** | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as &#39;#&#39;, &#39;+&#39;, &#39;;&#39;, &#39;&#x3D;&#39; as the value of the group name returned from the LDAP server might prepend those characters with &#39;\\&#39;. For example a group name called &#39;test#,lab,com&#39; will be returned from the LDAP server as &#39;test\\#,lab,com&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnAuthorizationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnBridge

> SempMetaOnlyResponse DeleteMsgVpnBridge(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Execute()

Delete a Bridge object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnBridge(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnBridge`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnBridgeRemoteMsgVpn

> SempMetaOnlyResponse DeleteMsgVpnBridgeRemoteMsgVpn(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).Execute()

Delete a Remote Message VPN object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    remoteMsgVpnName := "remoteMsgVpnName_example" // string | The name of the remote Message VPN.
    remoteMsgVpnLocation := "remoteMsgVpnLocation_example" // string | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \"v:\").
    remoteMsgVpnInterface := "remoteMsgVpnInterface_example" // string | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, `remoteMsgVpnLocation` must not be a virtual router name.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnBridgeRemoteMsgVpn(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnBridgeRemoteMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnBridgeRemoteMsgVpn`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnBridgeRemoteMsgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**remoteMsgVpnName** | **string** | The name of the remote Message VPN. | 
**remoteMsgVpnLocation** | **string** | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \&quot;v:\&quot;). | 
**remoteMsgVpnInterface** | **string** | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, &#x60;remoteMsgVpnLocation&#x60; must not be a virtual router name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnBridgeRemoteMsgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnBridgeRemoteSubscription

> SempMetaOnlyResponse DeleteMsgVpnBridgeRemoteSubscription(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteSubscriptionTopic).Execute()

Delete a Remote Subscription object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    remoteSubscriptionTopic := "remoteSubscriptionTopic_example" // string | The topic of the Bridge remote subscription.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnBridgeRemoteSubscription(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, remoteSubscriptionTopic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnBridgeRemoteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnBridgeRemoteSubscription`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnBridgeRemoteSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**remoteSubscriptionTopic** | **string** | The topic of the Bridge remote subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnBridgeRemoteSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnBridgeTlsTrustedCommonName

> SempMetaOnlyResponse DeleteMsgVpnBridgeTlsTrustedCommonName(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, tlsTrustedCommonName).Execute()

Delete a Trusted Common Name object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    tlsTrustedCommonName := "tlsTrustedCommonName_example" // string | The expected trusted common name of the remote certificate.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnBridgeTlsTrustedCommonName(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, tlsTrustedCommonName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnBridgeTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnBridgeTlsTrustedCommonName`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnBridgeTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**tlsTrustedCommonName** | **string** | The expected trusted common name of the remote certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnBridgeTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnClientProfile

> SempMetaOnlyResponse DeleteMsgVpnClientProfile(ctx, msgVpnName, clientProfileName).Execute()

Delete a Client Profile object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    clientProfileName := "clientProfileName_example" // string | The name of the Client Profile.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnClientProfile(context.Background(), msgVpnName, clientProfileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnClientProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnClientProfile`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnClientProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientProfileName** | **string** | The name of the Client Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnClientProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnClientUsername

> SempMetaOnlyResponse DeleteMsgVpnClientUsername(ctx, msgVpnName, clientUsername).Execute()

Delete a Client Username object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    clientUsername := "clientUsername_example" // string | The name of the Client Username.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnClientUsername(context.Background(), msgVpnName, clientUsername).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnClientUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnClientUsername`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnClientUsername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientUsername** | **string** | The name of the Client Username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnClientUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnDistributedCache

> SempMetaOnlyResponse DeleteMsgVpnDistributedCache(ctx, msgVpnName, cacheName).Execute()

Delete a Distributed Cache object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnDistributedCache(context.Background(), msgVpnName, cacheName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnDistributedCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnDistributedCache`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnDistributedCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnDistributedCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnDistributedCacheCluster

> SempMetaOnlyResponse DeleteMsgVpnDistributedCacheCluster(ctx, msgVpnName, cacheName, clusterName).Execute()

Delete a Cache Cluster object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnDistributedCacheCluster(context.Background(), msgVpnName, cacheName, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnDistributedCacheCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnDistributedCacheCluster`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnDistributedCacheCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnDistributedCacheClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster

> SempMetaOnlyResponse DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(ctx, msgVpnName, cacheName, clusterName, homeClusterName).Execute()

Delete a Home Cache Cluster object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    homeClusterName := "homeClusterName_example" // string | The name of the remote Home Cache Cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(context.Background(), msgVpnName, cacheName, clusterName, homeClusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**homeClusterName** | **string** | The name of the remote Home Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix

> SempMetaOnlyResponse DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(ctx, msgVpnName, cacheName, clusterName, homeClusterName, topicPrefix).Execute()

Delete a Topic Prefix object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    homeClusterName := "homeClusterName_example" // string | The name of the remote Home Cache Cluster.
    topicPrefix := "topicPrefix_example" // string | A topic prefix for global topics available from the remote Home Cache Cluster. A wildcard (/>) is implied at the end of the prefix.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(context.Background(), msgVpnName, cacheName, clusterName, homeClusterName, topicPrefix).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**homeClusterName** | **string** | The name of the remote Home Cache Cluster. | 
**topicPrefix** | **string** | A topic prefix for global topics available from the remote Home Cache Cluster. A wildcard (/&gt;) is implied at the end of the prefix. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnDistributedCacheClusterInstance

> SempMetaOnlyResponse DeleteMsgVpnDistributedCacheClusterInstance(ctx, msgVpnName, cacheName, clusterName, instanceName).Execute()

Delete a Cache Instance object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    instanceName := "instanceName_example" // string | The name of the Cache Instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnDistributedCacheClusterInstance(context.Background(), msgVpnName, cacheName, clusterName, instanceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnDistributedCacheClusterInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnDistributedCacheClusterInstance`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnDistributedCacheClusterInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**instanceName** | **string** | The name of the Cache Instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnDistributedCacheClusterInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnDistributedCacheClusterTopic

> SempMetaOnlyResponse DeleteMsgVpnDistributedCacheClusterTopic(ctx, msgVpnName, cacheName, clusterName, topic).Execute()

Delete a Topic object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    topic := "topic_example" // string | The value of the Topic in the form a/b/c.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnDistributedCacheClusterTopic(context.Background(), msgVpnName, cacheName, clusterName, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnDistributedCacheClusterTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnDistributedCacheClusterTopic`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnDistributedCacheClusterTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**topic** | **string** | The value of the Topic in the form a/b/c. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnDistributedCacheClusterTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnDmrBridge

> SempMetaOnlyResponse DeleteMsgVpnDmrBridge(ctx, msgVpnName, remoteNodeName).Execute()

Delete a DMR Bridge object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the DMR Bridge.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnDmrBridge(context.Background(), msgVpnName, remoteNodeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnDmrBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnDmrBridge`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnDmrBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the DMR Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnDmrBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnJndiConnectionFactory

> SempMetaOnlyResponse DeleteMsgVpnJndiConnectionFactory(ctx, msgVpnName, connectionFactoryName).Execute()

Delete a JNDI Connection Factory object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    connectionFactoryName := "connectionFactoryName_example" // string | The name of the JMS Connection Factory.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnJndiConnectionFactory(context.Background(), msgVpnName, connectionFactoryName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnJndiConnectionFactory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnJndiConnectionFactory`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnJndiConnectionFactory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**connectionFactoryName** | **string** | The name of the JMS Connection Factory. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnJndiConnectionFactoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnJndiQueue

> SempMetaOnlyResponse DeleteMsgVpnJndiQueue(ctx, msgVpnName, queueName).Execute()

Delete a JNDI Queue object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    queueName := "queueName_example" // string | The JNDI name of the JMS Queue.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnJndiQueue(context.Background(), msgVpnName, queueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnJndiQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnJndiQueue`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnJndiQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The JNDI name of the JMS Queue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnJndiQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnJndiTopic

> SempMetaOnlyResponse DeleteMsgVpnJndiTopic(ctx, msgVpnName, topicName).Execute()

Delete a JNDI Topic object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    topicName := "topicName_example" // string | The JNDI name of the JMS Topic.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnJndiTopic(context.Background(), msgVpnName, topicName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnJndiTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnJndiTopic`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnJndiTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicName** | **string** | The JNDI name of the JMS Topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnJndiTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnMqttRetainCache

> SempMetaOnlyResponse DeleteMsgVpnMqttRetainCache(ctx, msgVpnName, cacheName).Execute()

Delete an MQTT Retain Cache object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the MQTT Retain Cache.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnMqttRetainCache(context.Background(), msgVpnName, cacheName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnMqttRetainCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnMqttRetainCache`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnMqttRetainCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the MQTT Retain Cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnMqttRetainCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnMqttSession

> SempMetaOnlyResponse DeleteMsgVpnMqttSession(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Execute()

Delete an MQTT Session object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    mqttSessionClientId := "mqttSessionClientId_example" // string | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
    mqttSessionVirtualRouter := "mqttSessionVirtualRouter_example" // string | The virtual router of the MQTT Session.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnMqttSession(context.Background(), msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnMqttSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnMqttSession`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnMqttSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**mqttSessionClientId** | **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
**mqttSessionVirtualRouter** | **string** | The virtual router of the MQTT Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnMqttSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnMqttSessionSubscription

> SempMetaOnlyResponse DeleteMsgVpnMqttSessionSubscription(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, subscriptionTopic).Execute()

Delete a Subscription object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    mqttSessionClientId := "mqttSessionClientId_example" // string | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
    mqttSessionVirtualRouter := "mqttSessionVirtualRouter_example" // string | The virtual router of the MQTT Session.
    subscriptionTopic := "subscriptionTopic_example" // string | The MQTT subscription topic.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnMqttSessionSubscription(context.Background(), msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, subscriptionTopic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnMqttSessionSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnMqttSessionSubscription`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnMqttSessionSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**mqttSessionClientId** | **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
**mqttSessionVirtualRouter** | **string** | The virtual router of the MQTT Session. | 
**subscriptionTopic** | **string** | The MQTT subscription topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnMqttSessionSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnQueue

> SempMetaOnlyResponse DeleteMsgVpnQueue(ctx, msgVpnName, queueName).Execute()

Delete a Queue object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    queueName := "queueName_example" // string | The name of the Queue.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnQueue(context.Background(), msgVpnName, queueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnQueue`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The name of the Queue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnQueueSubscription

> SempMetaOnlyResponse DeleteMsgVpnQueueSubscription(ctx, msgVpnName, queueName, subscriptionTopic).Execute()

Delete a Queue Subscription object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    queueName := "queueName_example" // string | The name of the Queue.
    subscriptionTopic := "subscriptionTopic_example" // string | The topic of the Subscription.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnQueueSubscription(context.Background(), msgVpnName, queueName, subscriptionTopic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnQueueSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnQueueSubscription`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnQueueSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The name of the Queue. | 
**subscriptionTopic** | **string** | The topic of the Subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnQueueSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnQueueTemplate

> SempMetaOnlyResponse DeleteMsgVpnQueueTemplate(ctx, msgVpnName, queueTemplateName).Execute()

Delete a Queue Template object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    queueTemplateName := "queueTemplateName_example" // string | The name of the Queue Template.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnQueueTemplate(context.Background(), msgVpnName, queueTemplateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnQueueTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnQueueTemplate`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnQueueTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueTemplateName** | **string** | The name of the Queue Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnQueueTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnReplayLog

> SempMetaOnlyResponse DeleteMsgVpnReplayLog(ctx, msgVpnName, replayLogName).Execute()

Delete a Replay Log object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    replayLogName := "replayLogName_example" // string | The name of the Replay Log.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnReplayLog(context.Background(), msgVpnName, replayLogName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnReplayLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnReplayLog`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnReplayLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**replayLogName** | **string** | The name of the Replay Log. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnReplayLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnReplicatedTopic

> SempMetaOnlyResponse DeleteMsgVpnReplicatedTopic(ctx, msgVpnName, replicatedTopic).Execute()

Delete a Replicated Topic object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    replicatedTopic := "replicatedTopic_example" // string | The topic for applying replication. Published messages matching this topic will be replicated to the standby site.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnReplicatedTopic(context.Background(), msgVpnName, replicatedTopic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnReplicatedTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnReplicatedTopic`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnReplicatedTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**replicatedTopic** | **string** | The topic for applying replication. Published messages matching this topic will be replicated to the standby site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnReplicatedTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnRestDeliveryPoint

> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPoint(ctx, msgVpnName, restDeliveryPointName).Execute()

Delete a REST Delivery Point object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnRestDeliveryPoint(context.Background(), msgVpnName, restDeliveryPointName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnRestDeliveryPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnRestDeliveryPoint`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnRestDeliveryPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnRestDeliveryPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnRestDeliveryPointQueueBinding

> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPointQueueBinding(ctx, msgVpnName, restDeliveryPointName, queueBindingName).Execute()

Delete a Queue Binding object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    queueBindingName := "queueBindingName_example" // string | The name of a queue in the Message VPN.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnRestDeliveryPointQueueBinding(context.Background(), msgVpnName, restDeliveryPointName, queueBindingName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnRestDeliveryPointQueueBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnRestDeliveryPointQueueBinding`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnRestDeliveryPointQueueBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 
**queueBindingName** | **string** | The name of a queue in the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnRestDeliveryPointQueueBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnRestDeliveryPointRestConsumer

> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPointRestConsumer(ctx, msgVpnName, restDeliveryPointName, restConsumerName).Execute()

Delete a REST Consumer object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    restConsumerName := "restConsumerName_example" // string | The name of the REST Consumer.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnRestDeliveryPointRestConsumer(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnRestDeliveryPointRestConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnRestDeliveryPointRestConsumer`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnRestDeliveryPointRestConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 
**restConsumerName** | **string** | The name of the REST Consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnRestDeliveryPointRestConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim

> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(ctx, msgVpnName, restDeliveryPointName, restConsumerName, oauthJwtClaimName).Execute()

Delete a Claim object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    restConsumerName := "restConsumerName_example" // string | The name of the REST Consumer.
    oauthJwtClaimName := "oauthJwtClaimName_example" // string | The name of the additional claim. Cannot be \"exp\", \"iat\", or \"jti\".

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName, oauthJwtClaimName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 
**restConsumerName** | **string** | The name of the REST Consumer. | 
**oauthJwtClaimName** | **string** | The name of the additional claim. Cannot be \&quot;exp\&quot;, \&quot;iat\&quot;, or \&quot;jti\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName

> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(ctx, msgVpnName, restDeliveryPointName, restConsumerName, tlsTrustedCommonName).Execute()

Delete a Trusted Common Name object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    restConsumerName := "restConsumerName_example" // string | The name of the REST Consumer.
    tlsTrustedCommonName := "tlsTrustedCommonName_example" // string | The expected trusted common name of the remote certificate.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName, tlsTrustedCommonName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 
**restConsumerName** | **string** | The name of the REST Consumer. | 
**tlsTrustedCommonName** | **string** | The expected trusted common name of the remote certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnSequencedTopic

> SempMetaOnlyResponse DeleteMsgVpnSequencedTopic(ctx, msgVpnName, sequencedTopic).Execute()

Delete a Sequenced Topic object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    sequencedTopic := "sequencedTopic_example" // string | Topic for applying sequence numbers.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnSequencedTopic(context.Background(), msgVpnName, sequencedTopic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnSequencedTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnSequencedTopic`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnSequencedTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**sequencedTopic** | **string** | Topic for applying sequence numbers. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnSequencedTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnTopicEndpoint

> SempMetaOnlyResponse DeleteMsgVpnTopicEndpoint(ctx, msgVpnName, topicEndpointName).Execute()

Delete a Topic Endpoint object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    topicEndpointName := "topicEndpointName_example" // string | The name of the Topic Endpoint.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnTopicEndpoint(context.Background(), msgVpnName, topicEndpointName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnTopicEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnTopicEndpoint`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnTopicEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicEndpointName** | **string** | The name of the Topic Endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnTopicEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnTopicEndpointTemplate

> SempMetaOnlyResponse DeleteMsgVpnTopicEndpointTemplate(ctx, msgVpnName, topicEndpointTemplateName).Execute()

Delete a Topic Endpoint Template object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    topicEndpointTemplateName := "topicEndpointTemplateName_example" // string | The name of the Topic Endpoint Template.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteMsgVpnTopicEndpointTemplate(context.Background(), msgVpnName, topicEndpointTemplateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteMsgVpnTopicEndpointTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnTopicEndpointTemplate`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteMsgVpnTopicEndpointTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicEndpointTemplateName** | **string** | The name of the Topic Endpoint Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnTopicEndpointTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualHostname

> SempMetaOnlyResponse DeleteVirtualHostname(ctx, virtualHostname).Execute()

Delete a Virtual Hostname object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    virtualHostname := "virtualHostname_example" // string | The virtual hostname.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.DeleteVirtualHostname(context.Background(), virtualHostname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.DeleteVirtualHostname``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVirtualHostname`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.DeleteVirtualHostname`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualHostname** | **string** | The virtual hostname. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualHostnameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAbout

> AboutResponse GetAbout(ctx).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get an About object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetAbout(context.Background()).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetAbout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAbout`: AboutResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetAbout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAboutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**AboutResponse**](AboutResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAboutApi

> AboutApiResponse GetAboutApi(ctx).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get an API Description object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetAboutApi(context.Background()).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetAboutApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAboutApi`: AboutApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetAboutApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAboutApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**AboutApiResponse**](AboutApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAboutUser

> AboutUserResponse GetAboutUser(ctx).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a User object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetAboutUser(context.Background()).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetAboutUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAboutUser`: AboutUserResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetAboutUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAboutUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**AboutUserResponse**](AboutUserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAboutUserMsgVpn

> AboutUserMsgVpnResponse GetAboutUserMsgVpn(ctx, msgVpnName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a User Message VPN object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetAboutUserMsgVpn(context.Background(), msgVpnName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetAboutUserMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAboutUserMsgVpn`: AboutUserMsgVpnResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetAboutUserMsgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAboutUserMsgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**AboutUserMsgVpnResponse**](AboutUserMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAboutUserMsgVpns

> AboutUserMsgVpnsResponse GetAboutUserMsgVpns(ctx).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of User Message VPN objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetAboutUserMsgVpns(context.Background()).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetAboutUserMsgVpns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAboutUserMsgVpns`: AboutUserMsgVpnsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetAboutUserMsgVpns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAboutUserMsgVpnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**AboutUserMsgVpnsResponse**](AboutUserMsgVpnsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBroker

> BrokerResponse GetBroker(ctx).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Broker object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetBroker(context.Background()).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetBroker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBroker`: BrokerResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetBroker`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**BrokerResponse**](BrokerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertAuthorities

> CertAuthoritiesResponse GetCertAuthorities(ctx).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Certificate Authority objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetCertAuthorities(context.Background()).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetCertAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertAuthorities`: CertAuthoritiesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetCertAuthorities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertAuthoritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthoritiesResponse**](CertAuthoritiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertAuthority

> CertAuthorityResponse GetCertAuthority(ctx, certAuthorityName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Certificate Authority object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetCertAuthority(context.Background(), certAuthorityName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertAuthority`: CertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityResponse**](CertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertAuthorityOcspTlsTrustedCommonName

> CertAuthorityOcspTlsTrustedCommonNameResponse GetCertAuthorityOcspTlsTrustedCommonName(ctx, certAuthorityName, ocspTlsTrustedCommonName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get an OCSP Responder Trusted Common Name object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    ocspTlsTrustedCommonName := "ocspTlsTrustedCommonName_example" // string | The expected Trusted Common Name of the OCSP responder remote certificate.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetCertAuthorityOcspTlsTrustedCommonName(context.Background(), certAuthorityName, ocspTlsTrustedCommonName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetCertAuthorityOcspTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertAuthorityOcspTlsTrustedCommonName`: CertAuthorityOcspTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetCertAuthorityOcspTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 
**ocspTlsTrustedCommonName** | **string** | The expected Trusted Common Name of the OCSP responder remote certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertAuthorityOcspTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityOcspTlsTrustedCommonNameResponse**](CertAuthorityOcspTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertAuthorityOcspTlsTrustedCommonNames

> CertAuthorityOcspTlsTrustedCommonNamesResponse GetCertAuthorityOcspTlsTrustedCommonNames(ctx, certAuthorityName).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of OCSP Responder Trusted Common Name objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetCertAuthorityOcspTlsTrustedCommonNames(context.Background(), certAuthorityName).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetCertAuthorityOcspTlsTrustedCommonNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertAuthorityOcspTlsTrustedCommonNames`: CertAuthorityOcspTlsTrustedCommonNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetCertAuthorityOcspTlsTrustedCommonNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertAuthorityOcspTlsTrustedCommonNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityOcspTlsTrustedCommonNamesResponse**](CertAuthorityOcspTlsTrustedCommonNamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientCertAuthorities

> ClientCertAuthoritiesResponse GetClientCertAuthorities(ctx).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Client Certificate Authority objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetClientCertAuthorities(context.Background()).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetClientCertAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientCertAuthorities`: ClientCertAuthoritiesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetClientCertAuthorities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientCertAuthoritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthoritiesResponse**](ClientCertAuthoritiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientCertAuthority

> ClientCertAuthorityResponse GetClientCertAuthority(ctx, certAuthorityName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Client Certificate Authority object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetClientCertAuthority(context.Background(), certAuthorityName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetClientCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientCertAuthority`: ClientCertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetClientCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthorityResponse**](ClientCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientCertAuthorityOcspTlsTrustedCommonName

> ClientCertAuthorityOcspTlsTrustedCommonNameResponse GetClientCertAuthorityOcspTlsTrustedCommonName(ctx, certAuthorityName, ocspTlsTrustedCommonName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get an OCSP Responder Trusted Common Name object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    ocspTlsTrustedCommonName := "ocspTlsTrustedCommonName_example" // string | The expected Trusted Common Name of the OCSP responder remote certificate.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetClientCertAuthorityOcspTlsTrustedCommonName(context.Background(), certAuthorityName, ocspTlsTrustedCommonName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetClientCertAuthorityOcspTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientCertAuthorityOcspTlsTrustedCommonName`: ClientCertAuthorityOcspTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetClientCertAuthorityOcspTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 
**ocspTlsTrustedCommonName** | **string** | The expected Trusted Common Name of the OCSP responder remote certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthorityOcspTlsTrustedCommonNameResponse**](ClientCertAuthorityOcspTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientCertAuthorityOcspTlsTrustedCommonNames

> ClientCertAuthorityOcspTlsTrustedCommonNamesResponse GetClientCertAuthorityOcspTlsTrustedCommonNames(ctx, certAuthorityName).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of OCSP Responder Trusted Common Name objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetClientCertAuthorityOcspTlsTrustedCommonNames(context.Background(), certAuthorityName).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetClientCertAuthorityOcspTlsTrustedCommonNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientCertAuthorityOcspTlsTrustedCommonNames`: ClientCertAuthorityOcspTlsTrustedCommonNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetClientCertAuthorityOcspTlsTrustedCommonNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthorityOcspTlsTrustedCommonNamesResponse**](ClientCertAuthorityOcspTlsTrustedCommonNamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrCluster

> DmrClusterResponse GetDmrCluster(ctx, dmrClusterName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Cluster object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetDmrCluster(context.Background(), dmrClusterName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetDmrCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrCluster`: DmrClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetDmrCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterResponse**](DmrClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusterLink

> DmrClusterLinkResponse GetDmrClusterLink(ctx, dmrClusterName, remoteNodeName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Link object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetDmrClusterLink(context.Background(), dmrClusterName, remoteNodeName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetDmrClusterLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusterLink`: DmrClusterLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetDmrClusterLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkResponse**](DmrClusterLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusterLinkRemoteAddress

> DmrClusterLinkRemoteAddressResponse GetDmrClusterLinkRemoteAddress(ctx, dmrClusterName, remoteNodeName, remoteAddress).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Remote Address object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    remoteAddress := "remoteAddress_example" // string | The FQDN or IP address (and optional port) of the remote node. If a port is not provided, it will vary based on the transport encoding: 55555 (plain-text), 55443 (encrypted), or 55003 (compressed).
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetDmrClusterLinkRemoteAddress(context.Background(), dmrClusterName, remoteNodeName, remoteAddress).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetDmrClusterLinkRemoteAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusterLinkRemoteAddress`: DmrClusterLinkRemoteAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetDmrClusterLinkRemoteAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 
**remoteAddress** | **string** | The FQDN or IP address (and optional port) of the remote node. If a port is not provided, it will vary based on the transport encoding: 55555 (plain-text), 55443 (encrypted), or 55003 (compressed). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterLinkRemoteAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkRemoteAddressResponse**](DmrClusterLinkRemoteAddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusterLinkRemoteAddresses

> DmrClusterLinkRemoteAddressesResponse GetDmrClusterLinkRemoteAddresses(ctx, dmrClusterName, remoteNodeName).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Remote Address objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetDmrClusterLinkRemoteAddresses(context.Background(), dmrClusterName, remoteNodeName).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetDmrClusterLinkRemoteAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusterLinkRemoteAddresses`: DmrClusterLinkRemoteAddressesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetDmrClusterLinkRemoteAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterLinkRemoteAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkRemoteAddressesResponse**](DmrClusterLinkRemoteAddressesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusterLinkTlsTrustedCommonName

> DmrClusterLinkTlsTrustedCommonNameResponse GetDmrClusterLinkTlsTrustedCommonName(ctx, dmrClusterName, remoteNodeName, tlsTrustedCommonName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Trusted Common Name object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    tlsTrustedCommonName := "tlsTrustedCommonName_example" // string | The expected trusted common name of the remote certificate.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetDmrClusterLinkTlsTrustedCommonName(context.Background(), dmrClusterName, remoteNodeName, tlsTrustedCommonName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetDmrClusterLinkTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusterLinkTlsTrustedCommonName`: DmrClusterLinkTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetDmrClusterLinkTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 
**tlsTrustedCommonName** | **string** | The expected trusted common name of the remote certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterLinkTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkTlsTrustedCommonNameResponse**](DmrClusterLinkTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusterLinkTlsTrustedCommonNames

> DmrClusterLinkTlsTrustedCommonNamesResponse GetDmrClusterLinkTlsTrustedCommonNames(ctx, dmrClusterName, remoteNodeName).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Trusted Common Name objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetDmrClusterLinkTlsTrustedCommonNames(context.Background(), dmrClusterName, remoteNodeName).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetDmrClusterLinkTlsTrustedCommonNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusterLinkTlsTrustedCommonNames`: DmrClusterLinkTlsTrustedCommonNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetDmrClusterLinkTlsTrustedCommonNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterLinkTlsTrustedCommonNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkTlsTrustedCommonNamesResponse**](DmrClusterLinkTlsTrustedCommonNamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusterLinks

> DmrClusterLinksResponse GetDmrClusterLinks(ctx, dmrClusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Link objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetDmrClusterLinks(context.Background(), dmrClusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetDmrClusterLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusterLinks`: DmrClusterLinksResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetDmrClusterLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinksResponse**](DmrClusterLinksResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusters

> DmrClustersResponse GetDmrClusters(ctx).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Cluster objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetDmrClusters(context.Background()).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetDmrClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusters`: DmrClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetDmrClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClustersResponse**](DmrClustersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainCertAuthorities

> DomainCertAuthoritiesResponse GetDomainCertAuthorities(ctx).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Domain Certificate Authority objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetDomainCertAuthorities(context.Background()).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetDomainCertAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainCertAuthorities`: DomainCertAuthoritiesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetDomainCertAuthorities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainCertAuthoritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DomainCertAuthoritiesResponse**](DomainCertAuthoritiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainCertAuthority

> DomainCertAuthorityResponse GetDomainCertAuthority(ctx, certAuthorityName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Domain Certificate Authority object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetDomainCertAuthority(context.Background(), certAuthorityName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetDomainCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainCertAuthority`: DomainCertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetDomainCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DomainCertAuthorityResponse**](DomainCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpn

> MsgVpnResponse GetMsgVpn(ctx, msgVpnName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Message VPN object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpn(context.Background(), msgVpnName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpn`: MsgVpnResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnResponse**](MsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfile

> MsgVpnAclProfileResponse GetMsgVpnAclProfile(ctx, msgVpnName, aclProfileName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get an ACL Profile object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAclProfile(context.Background(), msgVpnName, aclProfileName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAclProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfile`: MsgVpnAclProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAclProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileResponse**](MsgVpnAclProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfileClientConnectException

> MsgVpnAclProfileClientConnectExceptionResponse GetMsgVpnAclProfileClientConnectException(ctx, msgVpnName, aclProfileName, clientConnectExceptionAddress).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Client Connect Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    clientConnectExceptionAddress := "clientConnectExceptionAddress_example" // string | The IP address/netmask of the client connect exception in CIDR form.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAclProfileClientConnectException(context.Background(), msgVpnName, aclProfileName, clientConnectExceptionAddress).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAclProfileClientConnectException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileClientConnectException`: MsgVpnAclProfileClientConnectExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAclProfileClientConnectException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**clientConnectExceptionAddress** | **string** | The IP address/netmask of the client connect exception in CIDR form. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileClientConnectExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileClientConnectExceptionResponse**](MsgVpnAclProfileClientConnectExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfileClientConnectExceptions

> MsgVpnAclProfileClientConnectExceptionsResponse GetMsgVpnAclProfileClientConnectExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Client Connect Exception objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAclProfileClientConnectExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAclProfileClientConnectExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileClientConnectExceptions`: MsgVpnAclProfileClientConnectExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAclProfileClientConnectExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileClientConnectExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileClientConnectExceptionsResponse**](MsgVpnAclProfileClientConnectExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfilePublishException

> MsgVpnAclProfilePublishExceptionResponse GetMsgVpnAclProfilePublishException(ctx, msgVpnName, aclProfileName, topicSyntax, publishExceptionTopic).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Publish Topic Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    topicSyntax := "topicSyntax_example" // string | The syntax of the topic for the exception to the default action taken.
    publishExceptionTopic := "publishExceptionTopic_example" // string | The topic for the exception to the default action taken. May include wildcard characters.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAclProfilePublishException(context.Background(), msgVpnName, aclProfileName, topicSyntax, publishExceptionTopic).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAclProfilePublishException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfilePublishException`: MsgVpnAclProfilePublishExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAclProfilePublishException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**topicSyntax** | **string** | The syntax of the topic for the exception to the default action taken. | 
**publishExceptionTopic** | **string** | The topic for the exception to the default action taken. May include wildcard characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfilePublishExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishExceptionResponse**](MsgVpnAclProfilePublishExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfilePublishExceptions

> MsgVpnAclProfilePublishExceptionsResponse GetMsgVpnAclProfilePublishExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Publish Topic Exception objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAclProfilePublishExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAclProfilePublishExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfilePublishExceptions`: MsgVpnAclProfilePublishExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAclProfilePublishExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfilePublishExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishExceptionsResponse**](MsgVpnAclProfilePublishExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfilePublishTopicException

> MsgVpnAclProfilePublishTopicExceptionResponse GetMsgVpnAclProfilePublishTopicException(ctx, msgVpnName, aclProfileName, publishTopicExceptionSyntax, publishTopicException).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Publish Topic Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    publishTopicExceptionSyntax := "publishTopicExceptionSyntax_example" // string | The syntax of the topic for the exception to the default action taken.
    publishTopicException := "publishTopicException_example" // string | The topic for the exception to the default action taken. May include wildcard characters.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAclProfilePublishTopicException(context.Background(), msgVpnName, aclProfileName, publishTopicExceptionSyntax, publishTopicException).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAclProfilePublishTopicException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfilePublishTopicException`: MsgVpnAclProfilePublishTopicExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAclProfilePublishTopicException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**publishTopicExceptionSyntax** | **string** | The syntax of the topic for the exception to the default action taken. | 
**publishTopicException** | **string** | The topic for the exception to the default action taken. May include wildcard characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfilePublishTopicExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishTopicExceptionResponse**](MsgVpnAclProfilePublishTopicExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfilePublishTopicExceptions

> MsgVpnAclProfilePublishTopicExceptionsResponse GetMsgVpnAclProfilePublishTopicExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Publish Topic Exception objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAclProfilePublishTopicExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAclProfilePublishTopicExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfilePublishTopicExceptions`: MsgVpnAclProfilePublishTopicExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAclProfilePublishTopicExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfilePublishTopicExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishTopicExceptionsResponse**](MsgVpnAclProfilePublishTopicExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfileSubscribeException

> MsgVpnAclProfileSubscribeExceptionResponse GetMsgVpnAclProfileSubscribeException(ctx, msgVpnName, aclProfileName, topicSyntax, subscribeExceptionTopic).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Subscribe Topic Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    topicSyntax := "topicSyntax_example" // string | The syntax of the topic for the exception to the default action taken.
    subscribeExceptionTopic := "subscribeExceptionTopic_example" // string | The topic for the exception to the default action taken. May include wildcard characters.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAclProfileSubscribeException(context.Background(), msgVpnName, aclProfileName, topicSyntax, subscribeExceptionTopic).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAclProfileSubscribeException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeException`: MsgVpnAclProfileSubscribeExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAclProfileSubscribeException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**topicSyntax** | **string** | The syntax of the topic for the exception to the default action taken. | 
**subscribeExceptionTopic** | **string** | The topic for the exception to the default action taken. May include wildcard characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileSubscribeExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeExceptionResponse**](MsgVpnAclProfileSubscribeExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfileSubscribeExceptions

> MsgVpnAclProfileSubscribeExceptionsResponse GetMsgVpnAclProfileSubscribeExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Subscribe Topic Exception objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAclProfileSubscribeExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAclProfileSubscribeExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeExceptions`: MsgVpnAclProfileSubscribeExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAclProfileSubscribeExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileSubscribeExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeExceptionsResponse**](MsgVpnAclProfileSubscribeExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfileSubscribeShareNameException

> MsgVpnAclProfileSubscribeShareNameExceptionResponse GetMsgVpnAclProfileSubscribeShareNameException(ctx, msgVpnName, aclProfileName, subscribeShareNameExceptionSyntax, subscribeShareNameException).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Subscribe Share Name Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    subscribeShareNameExceptionSyntax := "subscribeShareNameExceptionSyntax_example" // string | The syntax of the subscribe share name for the exception to the default action taken.
    subscribeShareNameException := "subscribeShareNameException_example" // string | The subscribe share name exception to the default action taken. May include wildcard characters.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAclProfileSubscribeShareNameException(context.Background(), msgVpnName, aclProfileName, subscribeShareNameExceptionSyntax, subscribeShareNameException).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAclProfileSubscribeShareNameException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeShareNameException`: MsgVpnAclProfileSubscribeShareNameExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAclProfileSubscribeShareNameException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**subscribeShareNameExceptionSyntax** | **string** | The syntax of the subscribe share name for the exception to the default action taken. | 
**subscribeShareNameException** | **string** | The subscribe share name exception to the default action taken. May include wildcard characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeShareNameExceptionResponse**](MsgVpnAclProfileSubscribeShareNameExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfileSubscribeShareNameExceptions

> MsgVpnAclProfileSubscribeShareNameExceptionsResponse GetMsgVpnAclProfileSubscribeShareNameExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Subscribe Share Name Exception objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAclProfileSubscribeShareNameExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAclProfileSubscribeShareNameExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeShareNameExceptions`: MsgVpnAclProfileSubscribeShareNameExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAclProfileSubscribeShareNameExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeShareNameExceptionsResponse**](MsgVpnAclProfileSubscribeShareNameExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfileSubscribeTopicException

> MsgVpnAclProfileSubscribeTopicExceptionResponse GetMsgVpnAclProfileSubscribeTopicException(ctx, msgVpnName, aclProfileName, subscribeTopicExceptionSyntax, subscribeTopicException).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Subscribe Topic Exception object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    subscribeTopicExceptionSyntax := "subscribeTopicExceptionSyntax_example" // string | The syntax of the topic for the exception to the default action taken.
    subscribeTopicException := "subscribeTopicException_example" // string | The topic for the exception to the default action taken. May include wildcard characters.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAclProfileSubscribeTopicException(context.Background(), msgVpnName, aclProfileName, subscribeTopicExceptionSyntax, subscribeTopicException).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAclProfileSubscribeTopicException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeTopicException`: MsgVpnAclProfileSubscribeTopicExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAclProfileSubscribeTopicException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**subscribeTopicExceptionSyntax** | **string** | The syntax of the topic for the exception to the default action taken. | 
**subscribeTopicException** | **string** | The topic for the exception to the default action taken. May include wildcard characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileSubscribeTopicExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeTopicExceptionResponse**](MsgVpnAclProfileSubscribeTopicExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfileSubscribeTopicExceptions

> MsgVpnAclProfileSubscribeTopicExceptionsResponse GetMsgVpnAclProfileSubscribeTopicExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Subscribe Topic Exception objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAclProfileSubscribeTopicExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAclProfileSubscribeTopicExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeTopicExceptions`: MsgVpnAclProfileSubscribeTopicExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAclProfileSubscribeTopicExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeTopicExceptionsResponse**](MsgVpnAclProfileSubscribeTopicExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfiles

> MsgVpnAclProfilesResponse GetMsgVpnAclProfiles(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of ACL Profile objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAclProfiles(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAclProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfiles`: MsgVpnAclProfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAclProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilesResponse**](MsgVpnAclProfilesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAuthenticationOauthProvider

> MsgVpnAuthenticationOauthProviderResponse GetMsgVpnAuthenticationOauthProvider(ctx, msgVpnName, oauthProviderName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get an OAuth Provider object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    oauthProviderName := "oauthProviderName_example" // string | The name of the OAuth Provider.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAuthenticationOauthProvider(context.Background(), msgVpnName, oauthProviderName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAuthenticationOauthProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAuthenticationOauthProvider`: MsgVpnAuthenticationOauthProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAuthenticationOauthProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**oauthProviderName** | **string** | The name of the OAuth Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAuthenticationOauthProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProviderResponse**](MsgVpnAuthenticationOauthProviderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAuthenticationOauthProviders

> MsgVpnAuthenticationOauthProvidersResponse GetMsgVpnAuthenticationOauthProviders(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of OAuth Provider objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAuthenticationOauthProviders(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAuthenticationOauthProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAuthenticationOauthProviders`: MsgVpnAuthenticationOauthProvidersResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAuthenticationOauthProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAuthenticationOauthProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProvidersResponse**](MsgVpnAuthenticationOauthProvidersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAuthorizationGroup

> MsgVpnAuthorizationGroupResponse GetMsgVpnAuthorizationGroup(ctx, msgVpnName, authorizationGroupName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get an LDAP Authorization Group object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    authorizationGroupName := "authorizationGroupName_example" // string | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as '#', '+', ';', '=' as the value of the group name returned from the LDAP server might prepend those characters with '\\'. For example a group name called 'test#,lab,com' will be returned from the LDAP server as 'test\\#,lab,com'.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAuthorizationGroup(context.Background(), msgVpnName, authorizationGroupName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAuthorizationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAuthorizationGroup`: MsgVpnAuthorizationGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAuthorizationGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**authorizationGroupName** | **string** | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as &#39;#&#39;, &#39;+&#39;, &#39;;&#39;, &#39;&#x3D;&#39; as the value of the group name returned from the LDAP server might prepend those characters with &#39;\\&#39;. For example a group name called &#39;test#,lab,com&#39; will be returned from the LDAP server as &#39;test\\#,lab,com&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAuthorizationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupResponse**](MsgVpnAuthorizationGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAuthorizationGroups

> MsgVpnAuthorizationGroupsResponse GetMsgVpnAuthorizationGroups(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of LDAP Authorization Group objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnAuthorizationGroups(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnAuthorizationGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAuthorizationGroups`: MsgVpnAuthorizationGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnAuthorizationGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAuthorizationGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupsResponse**](MsgVpnAuthorizationGroupsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridge

> MsgVpnBridgeResponse GetMsgVpnBridge(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Bridge object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnBridge(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridge`: MsgVpnBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridgeRemoteMsgVpn

> MsgVpnBridgeRemoteMsgVpnResponse GetMsgVpnBridgeRemoteMsgVpn(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Remote Message VPN object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    remoteMsgVpnName := "remoteMsgVpnName_example" // string | The name of the remote Message VPN.
    remoteMsgVpnLocation := "remoteMsgVpnLocation_example" // string | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \"v:\").
    remoteMsgVpnInterface := "remoteMsgVpnInterface_example" // string | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, `remoteMsgVpnLocation` must not be a virtual router name.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnBridgeRemoteMsgVpn(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnBridgeRemoteMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeRemoteMsgVpn`: MsgVpnBridgeRemoteMsgVpnResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnBridgeRemoteMsgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**remoteMsgVpnName** | **string** | The name of the remote Message VPN. | 
**remoteMsgVpnLocation** | **string** | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \&quot;v:\&quot;). | 
**remoteMsgVpnInterface** | **string** | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, &#x60;remoteMsgVpnLocation&#x60; must not be a virtual router name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgeRemoteMsgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteMsgVpnResponse**](MsgVpnBridgeRemoteMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridgeRemoteMsgVpns

> MsgVpnBridgeRemoteMsgVpnsResponse GetMsgVpnBridgeRemoteMsgVpns(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Remote Message VPN objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnBridgeRemoteMsgVpns(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnBridgeRemoteMsgVpns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeRemoteMsgVpns`: MsgVpnBridgeRemoteMsgVpnsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnBridgeRemoteMsgVpns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgeRemoteMsgVpnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteMsgVpnsResponse**](MsgVpnBridgeRemoteMsgVpnsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridgeRemoteSubscription

> MsgVpnBridgeRemoteSubscriptionResponse GetMsgVpnBridgeRemoteSubscription(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteSubscriptionTopic).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Remote Subscription object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    remoteSubscriptionTopic := "remoteSubscriptionTopic_example" // string | The topic of the Bridge remote subscription.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnBridgeRemoteSubscription(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, remoteSubscriptionTopic).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnBridgeRemoteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeRemoteSubscription`: MsgVpnBridgeRemoteSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnBridgeRemoteSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**remoteSubscriptionTopic** | **string** | The topic of the Bridge remote subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgeRemoteSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteSubscriptionResponse**](MsgVpnBridgeRemoteSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridgeRemoteSubscriptions

> MsgVpnBridgeRemoteSubscriptionsResponse GetMsgVpnBridgeRemoteSubscriptions(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Remote Subscription objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnBridgeRemoteSubscriptions(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnBridgeRemoteSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeRemoteSubscriptions`: MsgVpnBridgeRemoteSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnBridgeRemoteSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgeRemoteSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteSubscriptionsResponse**](MsgVpnBridgeRemoteSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridgeTlsTrustedCommonName

> MsgVpnBridgeTlsTrustedCommonNameResponse GetMsgVpnBridgeTlsTrustedCommonName(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, tlsTrustedCommonName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Trusted Common Name object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    tlsTrustedCommonName := "tlsTrustedCommonName_example" // string | The expected trusted common name of the remote certificate.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnBridgeTlsTrustedCommonName(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, tlsTrustedCommonName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnBridgeTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeTlsTrustedCommonName`: MsgVpnBridgeTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnBridgeTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**tlsTrustedCommonName** | **string** | The expected trusted common name of the remote certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgeTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeTlsTrustedCommonNameResponse**](MsgVpnBridgeTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridgeTlsTrustedCommonNames

> MsgVpnBridgeTlsTrustedCommonNamesResponse GetMsgVpnBridgeTlsTrustedCommonNames(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Trusted Common Name objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnBridgeTlsTrustedCommonNames(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnBridgeTlsTrustedCommonNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeTlsTrustedCommonNames`: MsgVpnBridgeTlsTrustedCommonNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnBridgeTlsTrustedCommonNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgeTlsTrustedCommonNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeTlsTrustedCommonNamesResponse**](MsgVpnBridgeTlsTrustedCommonNamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridges

> MsgVpnBridgesResponse GetMsgVpnBridges(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Bridge objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnBridges(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnBridges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridges`: MsgVpnBridgesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnBridges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgesResponse**](MsgVpnBridgesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientProfile

> MsgVpnClientProfileResponse GetMsgVpnClientProfile(ctx, msgVpnName, clientProfileName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Client Profile object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    clientProfileName := "clientProfileName_example" // string | The name of the Client Profile.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnClientProfile(context.Background(), msgVpnName, clientProfileName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnClientProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientProfile`: MsgVpnClientProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnClientProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientProfileName** | **string** | The name of the Client Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientProfileResponse**](MsgVpnClientProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientProfiles

> MsgVpnClientProfilesResponse GetMsgVpnClientProfiles(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Client Profile objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnClientProfiles(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnClientProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientProfiles`: MsgVpnClientProfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnClientProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientProfilesResponse**](MsgVpnClientProfilesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientUsername

> MsgVpnClientUsernameResponse GetMsgVpnClientUsername(ctx, msgVpnName, clientUsername).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Client Username object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    clientUsername := "clientUsername_example" // string | The name of the Client Username.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnClientUsername(context.Background(), msgVpnName, clientUsername).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnClientUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientUsername`: MsgVpnClientUsernameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnClientUsername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientUsername** | **string** | The name of the Client Username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernameResponse**](MsgVpnClientUsernameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientUsernames

> MsgVpnClientUsernamesResponse GetMsgVpnClientUsernames(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Client Username objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnClientUsernames(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnClientUsernames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientUsernames`: MsgVpnClientUsernamesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnClientUsernames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientUsernamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernamesResponse**](MsgVpnClientUsernamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCache

> MsgVpnDistributedCacheResponse GetMsgVpnDistributedCache(ctx, msgVpnName, cacheName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Distributed Cache object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnDistributedCache(context.Background(), msgVpnName, cacheName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnDistributedCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCache`: MsgVpnDistributedCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnDistributedCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheResponse**](MsgVpnDistributedCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheCluster

> MsgVpnDistributedCacheClusterResponse GetMsgVpnDistributedCacheCluster(ctx, msgVpnName, cacheName, clusterName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Cache Cluster object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnDistributedCacheCluster(context.Background(), msgVpnName, cacheName, clusterName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnDistributedCacheCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheCluster`: MsgVpnDistributedCacheClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnDistributedCacheCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterResponse**](MsgVpnDistributedCacheClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster

> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(ctx, msgVpnName, cacheName, clusterName, homeClusterName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Home Cache Cluster object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    homeClusterName := "homeClusterName_example" // string | The name of the remote Home Cache Cluster.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(context.Background(), msgVpnName, cacheName, clusterName, homeClusterName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster`: MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**homeClusterName** | **string** | The name of the remote Home Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix

> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(ctx, msgVpnName, cacheName, clusterName, homeClusterName, topicPrefix).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Topic Prefix object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    homeClusterName := "homeClusterName_example" // string | The name of the remote Home Cache Cluster.
    topicPrefix := "topicPrefix_example" // string | A topic prefix for global topics available from the remote Home Cache Cluster. A wildcard (/>) is implied at the end of the prefix.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(context.Background(), msgVpnName, cacheName, clusterName, homeClusterName, topicPrefix).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix`: MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**homeClusterName** | **string** | The name of the remote Home Cache Cluster. | 
**topicPrefix** | **string** | A topic prefix for global topics available from the remote Home Cache Cluster. A wildcard (/&gt;) is implied at the end of the prefix. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes

> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes(ctx, msgVpnName, cacheName, clusterName, homeClusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Topic Prefix objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    homeClusterName := "homeClusterName_example" // string | The name of the remote Home Cache Cluster.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes(context.Background(), msgVpnName, cacheName, clusterName, homeClusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes`: MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**homeClusterName** | **string** | The name of the remote Home Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters

> MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters(ctx, msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Home Cache Cluster objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters(context.Background(), msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters`: MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterInstance

> MsgVpnDistributedCacheClusterInstanceResponse GetMsgVpnDistributedCacheClusterInstance(ctx, msgVpnName, cacheName, clusterName, instanceName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Cache Instance object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    instanceName := "instanceName_example" // string | The name of the Cache Instance.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnDistributedCacheClusterInstance(context.Background(), msgVpnName, cacheName, clusterName, instanceName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnDistributedCacheClusterInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterInstance`: MsgVpnDistributedCacheClusterInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnDistributedCacheClusterInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**instanceName** | **string** | The name of the Cache Instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceResponse**](MsgVpnDistributedCacheClusterInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterInstances

> MsgVpnDistributedCacheClusterInstancesResponse GetMsgVpnDistributedCacheClusterInstances(ctx, msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Cache Instance objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnDistributedCacheClusterInstances(context.Background(), msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnDistributedCacheClusterInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterInstances`: MsgVpnDistributedCacheClusterInstancesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnDistributedCacheClusterInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstancesResponse**](MsgVpnDistributedCacheClusterInstancesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterTopic

> MsgVpnDistributedCacheClusterTopicResponse GetMsgVpnDistributedCacheClusterTopic(ctx, msgVpnName, cacheName, clusterName, topic).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Topic object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    topic := "topic_example" // string | The value of the Topic in the form a/b/c.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnDistributedCacheClusterTopic(context.Background(), msgVpnName, cacheName, clusterName, topic).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnDistributedCacheClusterTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterTopic`: MsgVpnDistributedCacheClusterTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnDistributedCacheClusterTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**topic** | **string** | The value of the Topic in the form a/b/c. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterTopicResponse**](MsgVpnDistributedCacheClusterTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterTopics

> MsgVpnDistributedCacheClusterTopicsResponse GetMsgVpnDistributedCacheClusterTopics(ctx, msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Topic objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnDistributedCacheClusterTopics(context.Background(), msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnDistributedCacheClusterTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterTopics`: MsgVpnDistributedCacheClusterTopicsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnDistributedCacheClusterTopics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterTopicsResponse**](MsgVpnDistributedCacheClusterTopicsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusters

> MsgVpnDistributedCacheClustersResponse GetMsgVpnDistributedCacheClusters(ctx, msgVpnName, cacheName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Cache Cluster objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnDistributedCacheClusters(context.Background(), msgVpnName, cacheName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnDistributedCacheClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusters`: MsgVpnDistributedCacheClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnDistributedCacheClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClustersResponse**](MsgVpnDistributedCacheClustersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCaches

> MsgVpnDistributedCachesResponse GetMsgVpnDistributedCaches(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Distributed Cache objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnDistributedCaches(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnDistributedCaches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCaches`: MsgVpnDistributedCachesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnDistributedCaches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCachesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCachesResponse**](MsgVpnDistributedCachesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDmrBridge

> MsgVpnDmrBridgeResponse GetMsgVpnDmrBridge(ctx, msgVpnName, remoteNodeName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a DMR Bridge object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the DMR Bridge.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnDmrBridge(context.Background(), msgVpnName, remoteNodeName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnDmrBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDmrBridge`: MsgVpnDmrBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnDmrBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the DMR Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDmrBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDmrBridgeResponse**](MsgVpnDmrBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDmrBridges

> MsgVpnDmrBridgesResponse GetMsgVpnDmrBridges(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of DMR Bridge objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnDmrBridges(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnDmrBridges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDmrBridges`: MsgVpnDmrBridgesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnDmrBridges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDmrBridgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDmrBridgesResponse**](MsgVpnDmrBridgesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnJndiConnectionFactories

> MsgVpnJndiConnectionFactoriesResponse GetMsgVpnJndiConnectionFactories(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of JNDI Connection Factory objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnJndiConnectionFactories(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnJndiConnectionFactories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnJndiConnectionFactories`: MsgVpnJndiConnectionFactoriesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnJndiConnectionFactories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnJndiConnectionFactoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiConnectionFactoriesResponse**](MsgVpnJndiConnectionFactoriesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnJndiConnectionFactory

> MsgVpnJndiConnectionFactoryResponse GetMsgVpnJndiConnectionFactory(ctx, msgVpnName, connectionFactoryName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a JNDI Connection Factory object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    connectionFactoryName := "connectionFactoryName_example" // string | The name of the JMS Connection Factory.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnJndiConnectionFactory(context.Background(), msgVpnName, connectionFactoryName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnJndiConnectionFactory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnJndiConnectionFactory`: MsgVpnJndiConnectionFactoryResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnJndiConnectionFactory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**connectionFactoryName** | **string** | The name of the JMS Connection Factory. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnJndiConnectionFactoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiConnectionFactoryResponse**](MsgVpnJndiConnectionFactoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnJndiQueue

> MsgVpnJndiQueueResponse GetMsgVpnJndiQueue(ctx, msgVpnName, queueName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a JNDI Queue object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    queueName := "queueName_example" // string | The JNDI name of the JMS Queue.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnJndiQueue(context.Background(), msgVpnName, queueName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnJndiQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnJndiQueue`: MsgVpnJndiQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnJndiQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The JNDI name of the JMS Queue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnJndiQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiQueueResponse**](MsgVpnJndiQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnJndiQueues

> MsgVpnJndiQueuesResponse GetMsgVpnJndiQueues(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of JNDI Queue objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnJndiQueues(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnJndiQueues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnJndiQueues`: MsgVpnJndiQueuesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnJndiQueues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnJndiQueuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiQueuesResponse**](MsgVpnJndiQueuesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnJndiTopic

> MsgVpnJndiTopicResponse GetMsgVpnJndiTopic(ctx, msgVpnName, topicName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a JNDI Topic object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    topicName := "topicName_example" // string | The JNDI name of the JMS Topic.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnJndiTopic(context.Background(), msgVpnName, topicName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnJndiTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnJndiTopic`: MsgVpnJndiTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnJndiTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicName** | **string** | The JNDI name of the JMS Topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnJndiTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiTopicResponse**](MsgVpnJndiTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnJndiTopics

> MsgVpnJndiTopicsResponse GetMsgVpnJndiTopics(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of JNDI Topic objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnJndiTopics(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnJndiTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnJndiTopics`: MsgVpnJndiTopicsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnJndiTopics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnJndiTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiTopicsResponse**](MsgVpnJndiTopicsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnMqttRetainCache

> MsgVpnMqttRetainCacheResponse GetMsgVpnMqttRetainCache(ctx, msgVpnName, cacheName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get an MQTT Retain Cache object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the MQTT Retain Cache.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnMqttRetainCache(context.Background(), msgVpnName, cacheName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnMqttRetainCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnMqttRetainCache`: MsgVpnMqttRetainCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnMqttRetainCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the MQTT Retain Cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnMqttRetainCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttRetainCacheResponse**](MsgVpnMqttRetainCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnMqttRetainCaches

> MsgVpnMqttRetainCachesResponse GetMsgVpnMqttRetainCaches(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of MQTT Retain Cache objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnMqttRetainCaches(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnMqttRetainCaches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnMqttRetainCaches`: MsgVpnMqttRetainCachesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnMqttRetainCaches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnMqttRetainCachesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttRetainCachesResponse**](MsgVpnMqttRetainCachesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnMqttSession

> MsgVpnMqttSessionResponse GetMsgVpnMqttSession(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get an MQTT Session object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    mqttSessionClientId := "mqttSessionClientId_example" // string | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
    mqttSessionVirtualRouter := "mqttSessionVirtualRouter_example" // string | The virtual router of the MQTT Session.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnMqttSession(context.Background(), msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnMqttSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnMqttSession`: MsgVpnMqttSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnMqttSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**mqttSessionClientId** | **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
**mqttSessionVirtualRouter** | **string** | The virtual router of the MQTT Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnMqttSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionResponse**](MsgVpnMqttSessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnMqttSessionSubscription

> MsgVpnMqttSessionSubscriptionResponse GetMsgVpnMqttSessionSubscription(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, subscriptionTopic).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Subscription object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    mqttSessionClientId := "mqttSessionClientId_example" // string | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
    mqttSessionVirtualRouter := "mqttSessionVirtualRouter_example" // string | The virtual router of the MQTT Session.
    subscriptionTopic := "subscriptionTopic_example" // string | The MQTT subscription topic.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnMqttSessionSubscription(context.Background(), msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, subscriptionTopic).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnMqttSessionSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnMqttSessionSubscription`: MsgVpnMqttSessionSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnMqttSessionSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**mqttSessionClientId** | **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
**mqttSessionVirtualRouter** | **string** | The virtual router of the MQTT Session. | 
**subscriptionTopic** | **string** | The MQTT subscription topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnMqttSessionSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionSubscriptionResponse**](MsgVpnMqttSessionSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnMqttSessionSubscriptions

> MsgVpnMqttSessionSubscriptionsResponse GetMsgVpnMqttSessionSubscriptions(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Subscription objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    mqttSessionClientId := "mqttSessionClientId_example" // string | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
    mqttSessionVirtualRouter := "mqttSessionVirtualRouter_example" // string | The virtual router of the MQTT Session.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnMqttSessionSubscriptions(context.Background(), msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnMqttSessionSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnMqttSessionSubscriptions`: MsgVpnMqttSessionSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnMqttSessionSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**mqttSessionClientId** | **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
**mqttSessionVirtualRouter** | **string** | The virtual router of the MQTT Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnMqttSessionSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionSubscriptionsResponse**](MsgVpnMqttSessionSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnMqttSessions

> MsgVpnMqttSessionsResponse GetMsgVpnMqttSessions(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of MQTT Session objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnMqttSessions(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnMqttSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnMqttSessions`: MsgVpnMqttSessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnMqttSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnMqttSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionsResponse**](MsgVpnMqttSessionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnQueue

> MsgVpnQueueResponse GetMsgVpnQueue(ctx, msgVpnName, queueName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Queue object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    queueName := "queueName_example" // string | The name of the Queue.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnQueue(context.Background(), msgVpnName, queueName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueue`: MsgVpnQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The name of the Queue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueResponse**](MsgVpnQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnQueueSubscription

> MsgVpnQueueSubscriptionResponse GetMsgVpnQueueSubscription(ctx, msgVpnName, queueName, subscriptionTopic).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Queue Subscription object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    queueName := "queueName_example" // string | The name of the Queue.
    subscriptionTopic := "subscriptionTopic_example" // string | The topic of the Subscription.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnQueueSubscription(context.Background(), msgVpnName, queueName, subscriptionTopic).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnQueueSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueueSubscription`: MsgVpnQueueSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnQueueSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The name of the Queue. | 
**subscriptionTopic** | **string** | The topic of the Subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnQueueSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueSubscriptionResponse**](MsgVpnQueueSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnQueueSubscriptions

> MsgVpnQueueSubscriptionsResponse GetMsgVpnQueueSubscriptions(ctx, msgVpnName, queueName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Queue Subscription objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    queueName := "queueName_example" // string | The name of the Queue.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnQueueSubscriptions(context.Background(), msgVpnName, queueName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnQueueSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueueSubscriptions`: MsgVpnQueueSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnQueueSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The name of the Queue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnQueueSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueSubscriptionsResponse**](MsgVpnQueueSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnQueueTemplate

> MsgVpnQueueTemplateResponse GetMsgVpnQueueTemplate(ctx, msgVpnName, queueTemplateName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Queue Template object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    queueTemplateName := "queueTemplateName_example" // string | The name of the Queue Template.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnQueueTemplate(context.Background(), msgVpnName, queueTemplateName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnQueueTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueueTemplate`: MsgVpnQueueTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnQueueTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueTemplateName** | **string** | The name of the Queue Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnQueueTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueTemplateResponse**](MsgVpnQueueTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnQueueTemplates

> MsgVpnQueueTemplatesResponse GetMsgVpnQueueTemplates(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Queue Template objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnQueueTemplates(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnQueueTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueueTemplates`: MsgVpnQueueTemplatesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnQueueTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnQueueTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueTemplatesResponse**](MsgVpnQueueTemplatesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnQueues

> MsgVpnQueuesResponse GetMsgVpnQueues(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Queue objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnQueues(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnQueues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueues`: MsgVpnQueuesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnQueues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnQueuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueuesResponse**](MsgVpnQueuesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnReplayLog

> MsgVpnReplayLogResponse GetMsgVpnReplayLog(ctx, msgVpnName, replayLogName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Replay Log object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    replayLogName := "replayLogName_example" // string | The name of the Replay Log.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnReplayLog(context.Background(), msgVpnName, replayLogName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnReplayLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnReplayLog`: MsgVpnReplayLogResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnReplayLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**replayLogName** | **string** | The name of the Replay Log. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnReplayLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogResponse**](MsgVpnReplayLogResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnReplayLogs

> MsgVpnReplayLogsResponse GetMsgVpnReplayLogs(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Replay Log objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnReplayLogs(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnReplayLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnReplayLogs`: MsgVpnReplayLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnReplayLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnReplayLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogsResponse**](MsgVpnReplayLogsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnReplicatedTopic

> MsgVpnReplicatedTopicResponse GetMsgVpnReplicatedTopic(ctx, msgVpnName, replicatedTopic).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Replicated Topic object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    replicatedTopic := "replicatedTopic_example" // string | The topic for applying replication. Published messages matching this topic will be replicated to the standby site.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnReplicatedTopic(context.Background(), msgVpnName, replicatedTopic).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnReplicatedTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnReplicatedTopic`: MsgVpnReplicatedTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnReplicatedTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**replicatedTopic** | **string** | The topic for applying replication. Published messages matching this topic will be replicated to the standby site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnReplicatedTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplicatedTopicResponse**](MsgVpnReplicatedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnReplicatedTopics

> MsgVpnReplicatedTopicsResponse GetMsgVpnReplicatedTopics(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Replicated Topic objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnReplicatedTopics(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnReplicatedTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnReplicatedTopics`: MsgVpnReplicatedTopicsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnReplicatedTopics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnReplicatedTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplicatedTopicsResponse**](MsgVpnReplicatedTopicsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnRestDeliveryPoint

> MsgVpnRestDeliveryPointResponse GetMsgVpnRestDeliveryPoint(ctx, msgVpnName, restDeliveryPointName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a REST Delivery Point object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnRestDeliveryPoint(context.Background(), msgVpnName, restDeliveryPointName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnRestDeliveryPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPoint`: MsgVpnRestDeliveryPointResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnRestDeliveryPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnRestDeliveryPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointResponse**](MsgVpnRestDeliveryPointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnRestDeliveryPointQueueBinding

> MsgVpnRestDeliveryPointQueueBindingResponse GetMsgVpnRestDeliveryPointQueueBinding(ctx, msgVpnName, restDeliveryPointName, queueBindingName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Queue Binding object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    queueBindingName := "queueBindingName_example" // string | The name of a queue in the Message VPN.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnRestDeliveryPointQueueBinding(context.Background(), msgVpnName, restDeliveryPointName, queueBindingName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnRestDeliveryPointQueueBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPointQueueBinding`: MsgVpnRestDeliveryPointQueueBindingResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnRestDeliveryPointQueueBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 
**queueBindingName** | **string** | The name of a queue in the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnRestDeliveryPointQueueBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingResponse**](MsgVpnRestDeliveryPointQueueBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnRestDeliveryPointQueueBindings

> MsgVpnRestDeliveryPointQueueBindingsResponse GetMsgVpnRestDeliveryPointQueueBindings(ctx, msgVpnName, restDeliveryPointName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Queue Binding objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnRestDeliveryPointQueueBindings(context.Background(), msgVpnName, restDeliveryPointName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnRestDeliveryPointQueueBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPointQueueBindings`: MsgVpnRestDeliveryPointQueueBindingsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnRestDeliveryPointQueueBindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnRestDeliveryPointQueueBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingsResponse**](MsgVpnRestDeliveryPointQueueBindingsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnRestDeliveryPointRestConsumer

> MsgVpnRestDeliveryPointRestConsumerResponse GetMsgVpnRestDeliveryPointRestConsumer(ctx, msgVpnName, restDeliveryPointName, restConsumerName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a REST Consumer object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    restConsumerName := "restConsumerName_example" // string | The name of the REST Consumer.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnRestDeliveryPointRestConsumer(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnRestDeliveryPointRestConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPointRestConsumer`: MsgVpnRestDeliveryPointRestConsumerResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnRestDeliveryPointRestConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 
**restConsumerName** | **string** | The name of the REST Consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnRestDeliveryPointRestConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerResponse**](MsgVpnRestDeliveryPointRestConsumerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim

> MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(ctx, msgVpnName, restDeliveryPointName, restConsumerName, oauthJwtClaimName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Claim object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    restConsumerName := "restConsumerName_example" // string | The name of the REST Consumer.
    oauthJwtClaimName := "oauthJwtClaimName_example" // string | The name of the additional claim. Cannot be \"exp\", \"iat\", or \"jti\".
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName, oauthJwtClaimName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim`: MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 
**restConsumerName** | **string** | The name of the REST Consumer. | 
**oauthJwtClaimName** | **string** | The name of the additional claim. Cannot be \&quot;exp\&quot;, \&quot;iat\&quot;, or \&quot;jti\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse**](MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims

> MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims(ctx, msgVpnName, restDeliveryPointName, restConsumerName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Claim objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    restConsumerName := "restConsumerName_example" // string | The name of the REST Consumer.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims`: MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 
**restConsumerName** | **string** | The name of the REST Consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse**](MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName

> MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(ctx, msgVpnName, restDeliveryPointName, restConsumerName, tlsTrustedCommonName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Trusted Common Name object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    restConsumerName := "restConsumerName_example" // string | The name of the REST Consumer.
    tlsTrustedCommonName := "tlsTrustedCommonName_example" // string | The expected trusted common name of the remote certificate.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName, tlsTrustedCommonName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName`: MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 
**restConsumerName** | **string** | The name of the REST Consumer. | 
**tlsTrustedCommonName** | **string** | The expected trusted common name of the remote certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse**](MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames

> MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames(ctx, msgVpnName, restDeliveryPointName, restConsumerName).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Trusted Common Name objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    restConsumerName := "restConsumerName_example" // string | The name of the REST Consumer.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames`: MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 
**restConsumerName** | **string** | The name of the REST Consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse**](MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnRestDeliveryPointRestConsumers

> MsgVpnRestDeliveryPointRestConsumersResponse GetMsgVpnRestDeliveryPointRestConsumers(ctx, msgVpnName, restDeliveryPointName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of REST Consumer objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnRestDeliveryPointRestConsumers(context.Background(), msgVpnName, restDeliveryPointName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnRestDeliveryPointRestConsumers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPointRestConsumers`: MsgVpnRestDeliveryPointRestConsumersResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnRestDeliveryPointRestConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnRestDeliveryPointRestConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumersResponse**](MsgVpnRestDeliveryPointRestConsumersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnRestDeliveryPoints

> MsgVpnRestDeliveryPointsResponse GetMsgVpnRestDeliveryPoints(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of REST Delivery Point objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnRestDeliveryPoints(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnRestDeliveryPoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPoints`: MsgVpnRestDeliveryPointsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnRestDeliveryPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnRestDeliveryPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointsResponse**](MsgVpnRestDeliveryPointsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnSequencedTopic

> MsgVpnSequencedTopicResponse GetMsgVpnSequencedTopic(ctx, msgVpnName, sequencedTopic).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Sequenced Topic object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    sequencedTopic := "sequencedTopic_example" // string | Topic for applying sequence numbers.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnSequencedTopic(context.Background(), msgVpnName, sequencedTopic).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnSequencedTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnSequencedTopic`: MsgVpnSequencedTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnSequencedTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**sequencedTopic** | **string** | Topic for applying sequence numbers. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnSequencedTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnSequencedTopicResponse**](MsgVpnSequencedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnSequencedTopics

> MsgVpnSequencedTopicsResponse GetMsgVpnSequencedTopics(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Sequenced Topic objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnSequencedTopics(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnSequencedTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnSequencedTopics`: MsgVpnSequencedTopicsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnSequencedTopics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnSequencedTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnSequencedTopicsResponse**](MsgVpnSequencedTopicsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnTopicEndpoint

> MsgVpnTopicEndpointResponse GetMsgVpnTopicEndpoint(ctx, msgVpnName, topicEndpointName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Topic Endpoint object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    topicEndpointName := "topicEndpointName_example" // string | The name of the Topic Endpoint.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnTopicEndpoint(context.Background(), msgVpnName, topicEndpointName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnTopicEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTopicEndpoint`: MsgVpnTopicEndpointResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnTopicEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicEndpointName** | **string** | The name of the Topic Endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnTopicEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointResponse**](MsgVpnTopicEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnTopicEndpointTemplate

> MsgVpnTopicEndpointTemplateResponse GetMsgVpnTopicEndpointTemplate(ctx, msgVpnName, topicEndpointTemplateName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Topic Endpoint Template object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    topicEndpointTemplateName := "topicEndpointTemplateName_example" // string | The name of the Topic Endpoint Template.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnTopicEndpointTemplate(context.Background(), msgVpnName, topicEndpointTemplateName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnTopicEndpointTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTopicEndpointTemplate`: MsgVpnTopicEndpointTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnTopicEndpointTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicEndpointTemplateName** | **string** | The name of the Topic Endpoint Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnTopicEndpointTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplateResponse**](MsgVpnTopicEndpointTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnTopicEndpointTemplates

> MsgVpnTopicEndpointTemplatesResponse GetMsgVpnTopicEndpointTemplates(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Topic Endpoint Template objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnTopicEndpointTemplates(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnTopicEndpointTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTopicEndpointTemplates`: MsgVpnTopicEndpointTemplatesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnTopicEndpointTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnTopicEndpointTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplatesResponse**](MsgVpnTopicEndpointTemplatesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnTopicEndpoints

> MsgVpnTopicEndpointsResponse GetMsgVpnTopicEndpoints(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Topic Endpoint objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpnTopicEndpoints(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpnTopicEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTopicEndpoints`: MsgVpnTopicEndpointsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpnTopicEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnTopicEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointsResponse**](MsgVpnTopicEndpointsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpns

> MsgVpnsResponse GetMsgVpns(ctx).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Message VPN objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetMsgVpns(context.Background()).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetMsgVpns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpns`: MsgVpnsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetMsgVpns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnsResponse**](MsgVpnsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemInformation

> SystemInformationResponse GetSystemInformation(ctx).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a System Information object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetSystemInformation(context.Background()).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetSystemInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemInformation`: SystemInformationResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetSystemInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**SystemInformationResponse**](SystemInformationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualHostname

> VirtualHostnameResponse GetVirtualHostname(ctx, virtualHostname).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Virtual Hostname object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    virtualHostname := "virtualHostname_example" // string | The virtual hostname.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetVirtualHostname(context.Background(), virtualHostname).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetVirtualHostname``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualHostname`: VirtualHostnameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetVirtualHostname`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualHostname** | **string** | The virtual hostname. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualHostnameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**VirtualHostnameResponse**](VirtualHostnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualHostnames

> VirtualHostnamesResponse GetVirtualHostnames(ctx).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Virtual Hostname objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.GetVirtualHostnames(context.Background()).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.GetVirtualHostnames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualHostnames`: VirtualHostnamesResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.GetVirtualHostnames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualHostnamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**VirtualHostnamesResponse**](VirtualHostnamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCertAuthority

> CertAuthorityResponse ReplaceCertAuthority(ctx, certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Certificate Authority object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    body := *openapiclient.NewCertAuthority() // CertAuthority | The Certificate Authority object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceCertAuthority(context.Background(), certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceCertAuthority`: CertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CertAuthority**](CertAuthority.md) | The Certificate Authority object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityResponse**](CertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceClientCertAuthority

> ClientCertAuthorityResponse ReplaceClientCertAuthority(ctx, certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Client Certificate Authority object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    body := *openapiclient.NewClientCertAuthority() // ClientCertAuthority | The Client Certificate Authority object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceClientCertAuthority(context.Background(), certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceClientCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceClientCertAuthority`: ClientCertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceClientCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceClientCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ClientCertAuthority**](ClientCertAuthority.md) | The Client Certificate Authority object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthorityResponse**](ClientCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceDmrCluster

> DmrClusterResponse ReplaceDmrCluster(ctx, dmrClusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Cluster object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    body := *openapiclient.NewDmrCluster() // DmrCluster | The Cluster object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceDmrCluster(context.Background(), dmrClusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceDmrCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceDmrCluster`: DmrClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceDmrCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceDmrClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DmrCluster**](DmrCluster.md) | The Cluster object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterResponse**](DmrClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceDmrClusterLink

> DmrClusterLinkResponse ReplaceDmrClusterLink(ctx, dmrClusterName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Link object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    body := *openapiclient.NewDmrClusterLink() // DmrClusterLink | The Link object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceDmrClusterLink(context.Background(), dmrClusterName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceDmrClusterLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceDmrClusterLink`: DmrClusterLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceDmrClusterLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceDmrClusterLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DmrClusterLink**](DmrClusterLink.md) | The Link object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkResponse**](DmrClusterLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceDomainCertAuthority

> DomainCertAuthorityResponse ReplaceDomainCertAuthority(ctx, certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Domain Certificate Authority object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    body := *openapiclient.NewDomainCertAuthority() // DomainCertAuthority | The Domain Certificate Authority object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceDomainCertAuthority(context.Background(), certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceDomainCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceDomainCertAuthority`: DomainCertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceDomainCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceDomainCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DomainCertAuthority**](DomainCertAuthority.md) | The Domain Certificate Authority object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DomainCertAuthorityResponse**](DomainCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpn

> MsgVpnResponse ReplaceMsgVpn(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Message VPN object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpn() // MsgVpn | The Message VPN object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpn(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpn`: MsgVpnResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpn**](MsgVpn.md) | The Message VPN object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnResponse**](MsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnAclProfile

> MsgVpnAclProfileResponse ReplaceMsgVpnAclProfile(ctx, msgVpnName, aclProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace an ACL Profile object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    body := *openapiclient.NewMsgVpnAclProfile() // MsgVpnAclProfile | The ACL Profile object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnAclProfile(context.Background(), msgVpnName, aclProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnAclProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnAclProfile`: MsgVpnAclProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnAclProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnAclProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnAclProfile**](MsgVpnAclProfile.md) | The ACL Profile object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileResponse**](MsgVpnAclProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnAuthenticationOauthProvider

> MsgVpnAuthenticationOauthProviderResponse ReplaceMsgVpnAuthenticationOauthProvider(ctx, msgVpnName, oauthProviderName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace an OAuth Provider object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    oauthProviderName := "oauthProviderName_example" // string | The name of the OAuth Provider.
    body := *openapiclient.NewMsgVpnAuthenticationOauthProvider() // MsgVpnAuthenticationOauthProvider | The OAuth Provider object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnAuthenticationOauthProvider(context.Background(), msgVpnName, oauthProviderName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnAuthenticationOauthProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnAuthenticationOauthProvider`: MsgVpnAuthenticationOauthProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnAuthenticationOauthProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**oauthProviderName** | **string** | The name of the OAuth Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnAuthenticationOauthProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnAuthenticationOauthProvider**](MsgVpnAuthenticationOauthProvider.md) | The OAuth Provider object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProviderResponse**](MsgVpnAuthenticationOauthProviderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnAuthorizationGroup

> MsgVpnAuthorizationGroupResponse ReplaceMsgVpnAuthorizationGroup(ctx, msgVpnName, authorizationGroupName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace an LDAP Authorization Group object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    authorizationGroupName := "authorizationGroupName_example" // string | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as '#', '+', ';', '=' as the value of the group name returned from the LDAP server might prepend those characters with '\\'. For example a group name called 'test#,lab,com' will be returned from the LDAP server as 'test\\#,lab,com'.
    body := *openapiclient.NewMsgVpnAuthorizationGroup() // MsgVpnAuthorizationGroup | The LDAP Authorization Group object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnAuthorizationGroup(context.Background(), msgVpnName, authorizationGroupName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnAuthorizationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnAuthorizationGroup`: MsgVpnAuthorizationGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnAuthorizationGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**authorizationGroupName** | **string** | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as &#39;#&#39;, &#39;+&#39;, &#39;;&#39;, &#39;&#x3D;&#39; as the value of the group name returned from the LDAP server might prepend those characters with &#39;\\&#39;. For example a group name called &#39;test#,lab,com&#39; will be returned from the LDAP server as &#39;test\\#,lab,com&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnAuthorizationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnAuthorizationGroup**](MsgVpnAuthorizationGroup.md) | The LDAP Authorization Group object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupResponse**](MsgVpnAuthorizationGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnBridge

> MsgVpnBridgeResponse ReplaceMsgVpnBridge(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Bridge object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    body := *openapiclient.NewMsgVpnBridge() // MsgVpnBridge | The Bridge object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnBridge(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnBridge`: MsgVpnBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnBridge**](MsgVpnBridge.md) | The Bridge object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnBridgeRemoteMsgVpn

> MsgVpnBridgeRemoteMsgVpnResponse ReplaceMsgVpnBridgeRemoteMsgVpn(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Remote Message VPN object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    remoteMsgVpnName := "remoteMsgVpnName_example" // string | The name of the remote Message VPN.
    remoteMsgVpnLocation := "remoteMsgVpnLocation_example" // string | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \"v:\").
    remoteMsgVpnInterface := "remoteMsgVpnInterface_example" // string | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, `remoteMsgVpnLocation` must not be a virtual router name.
    body := *openapiclient.NewMsgVpnBridgeRemoteMsgVpn() // MsgVpnBridgeRemoteMsgVpn | The Remote Message VPN object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnBridgeRemoteMsgVpn(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnBridgeRemoteMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnBridgeRemoteMsgVpn`: MsgVpnBridgeRemoteMsgVpnResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnBridgeRemoteMsgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**remoteMsgVpnName** | **string** | The name of the remote Message VPN. | 
**remoteMsgVpnLocation** | **string** | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \&quot;v:\&quot;). | 
**remoteMsgVpnInterface** | **string** | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, &#x60;remoteMsgVpnLocation&#x60; must not be a virtual router name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnBridgeRemoteMsgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **body** | [**MsgVpnBridgeRemoteMsgVpn**](MsgVpnBridgeRemoteMsgVpn.md) | The Remote Message VPN object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteMsgVpnResponse**](MsgVpnBridgeRemoteMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnClientProfile

> MsgVpnClientProfileResponse ReplaceMsgVpnClientProfile(ctx, msgVpnName, clientProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Client Profile object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    clientProfileName := "clientProfileName_example" // string | The name of the Client Profile.
    body := *openapiclient.NewMsgVpnClientProfile() // MsgVpnClientProfile | The Client Profile object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnClientProfile(context.Background(), msgVpnName, clientProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnClientProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnClientProfile`: MsgVpnClientProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnClientProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientProfileName** | **string** | The name of the Client Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnClientProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnClientProfile**](MsgVpnClientProfile.md) | The Client Profile object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientProfileResponse**](MsgVpnClientProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnClientUsername

> MsgVpnClientUsernameResponse ReplaceMsgVpnClientUsername(ctx, msgVpnName, clientUsername).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Client Username object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    clientUsername := "clientUsername_example" // string | The name of the Client Username.
    body := *openapiclient.NewMsgVpnClientUsername() // MsgVpnClientUsername | The Client Username object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnClientUsername(context.Background(), msgVpnName, clientUsername).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnClientUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnClientUsername`: MsgVpnClientUsernameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnClientUsername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientUsername** | **string** | The name of the Client Username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnClientUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnClientUsername**](MsgVpnClientUsername.md) | The Client Username object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernameResponse**](MsgVpnClientUsernameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnDistributedCache

> MsgVpnDistributedCacheResponse ReplaceMsgVpnDistributedCache(ctx, msgVpnName, cacheName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Distributed Cache object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    body := *openapiclient.NewMsgVpnDistributedCache() // MsgVpnDistributedCache | The Distributed Cache object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnDistributedCache(context.Background(), msgVpnName, cacheName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnDistributedCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnDistributedCache`: MsgVpnDistributedCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnDistributedCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnDistributedCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnDistributedCache**](MsgVpnDistributedCache.md) | The Distributed Cache object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheResponse**](MsgVpnDistributedCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnDistributedCacheCluster

> MsgVpnDistributedCacheClusterResponse ReplaceMsgVpnDistributedCacheCluster(ctx, msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Cache Cluster object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    body := *openapiclient.NewMsgVpnDistributedCacheCluster() // MsgVpnDistributedCacheCluster | The Cache Cluster object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnDistributedCacheCluster(context.Background(), msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnDistributedCacheCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnDistributedCacheCluster`: MsgVpnDistributedCacheClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnDistributedCacheCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnDistributedCacheClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnDistributedCacheCluster**](MsgVpnDistributedCacheCluster.md) | The Cache Cluster object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterResponse**](MsgVpnDistributedCacheClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnDistributedCacheClusterInstance

> MsgVpnDistributedCacheClusterInstanceResponse ReplaceMsgVpnDistributedCacheClusterInstance(ctx, msgVpnName, cacheName, clusterName, instanceName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Cache Instance object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    instanceName := "instanceName_example" // string | The name of the Cache Instance.
    body := *openapiclient.NewMsgVpnDistributedCacheClusterInstance() // MsgVpnDistributedCacheClusterInstance | The Cache Instance object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnDistributedCacheClusterInstance(context.Background(), msgVpnName, cacheName, clusterName, instanceName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnDistributedCacheClusterInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnDistributedCacheClusterInstance`: MsgVpnDistributedCacheClusterInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnDistributedCacheClusterInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**instanceName** | **string** | The name of the Cache Instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnDistributedCacheClusterInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**MsgVpnDistributedCacheClusterInstance**](MsgVpnDistributedCacheClusterInstance.md) | The Cache Instance object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceResponse**](MsgVpnDistributedCacheClusterInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnDmrBridge

> MsgVpnDmrBridgeResponse ReplaceMsgVpnDmrBridge(ctx, msgVpnName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a DMR Bridge object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the DMR Bridge.
    body := *openapiclient.NewMsgVpnDmrBridge() // MsgVpnDmrBridge | The DMR Bridge object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnDmrBridge(context.Background(), msgVpnName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnDmrBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnDmrBridge`: MsgVpnDmrBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnDmrBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the DMR Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnDmrBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnDmrBridge**](MsgVpnDmrBridge.md) | The DMR Bridge object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDmrBridgeResponse**](MsgVpnDmrBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnJndiConnectionFactory

> MsgVpnJndiConnectionFactoryResponse ReplaceMsgVpnJndiConnectionFactory(ctx, msgVpnName, connectionFactoryName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a JNDI Connection Factory object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    connectionFactoryName := "connectionFactoryName_example" // string | The name of the JMS Connection Factory.
    body := *openapiclient.NewMsgVpnJndiConnectionFactory() // MsgVpnJndiConnectionFactory | The JNDI Connection Factory object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnJndiConnectionFactory(context.Background(), msgVpnName, connectionFactoryName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnJndiConnectionFactory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnJndiConnectionFactory`: MsgVpnJndiConnectionFactoryResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnJndiConnectionFactory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**connectionFactoryName** | **string** | The name of the JMS Connection Factory. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnJndiConnectionFactoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnJndiConnectionFactory**](MsgVpnJndiConnectionFactory.md) | The JNDI Connection Factory object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiConnectionFactoryResponse**](MsgVpnJndiConnectionFactoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnJndiQueue

> MsgVpnJndiQueueResponse ReplaceMsgVpnJndiQueue(ctx, msgVpnName, queueName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a JNDI Queue object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    queueName := "queueName_example" // string | The JNDI name of the JMS Queue.
    body := *openapiclient.NewMsgVpnJndiQueue() // MsgVpnJndiQueue | The JNDI Queue object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnJndiQueue(context.Background(), msgVpnName, queueName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnJndiQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnJndiQueue`: MsgVpnJndiQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnJndiQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The JNDI name of the JMS Queue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnJndiQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnJndiQueue**](MsgVpnJndiQueue.md) | The JNDI Queue object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiQueueResponse**](MsgVpnJndiQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnJndiTopic

> MsgVpnJndiTopicResponse ReplaceMsgVpnJndiTopic(ctx, msgVpnName, topicName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a JNDI Topic object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    topicName := "topicName_example" // string | The JNDI name of the JMS Topic.
    body := *openapiclient.NewMsgVpnJndiTopic() // MsgVpnJndiTopic | The JNDI Topic object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnJndiTopic(context.Background(), msgVpnName, topicName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnJndiTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnJndiTopic`: MsgVpnJndiTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnJndiTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicName** | **string** | The JNDI name of the JMS Topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnJndiTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnJndiTopic**](MsgVpnJndiTopic.md) | The JNDI Topic object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiTopicResponse**](MsgVpnJndiTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnMqttRetainCache

> MsgVpnMqttRetainCacheResponse ReplaceMsgVpnMqttRetainCache(ctx, msgVpnName, cacheName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace an MQTT Retain Cache object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the MQTT Retain Cache.
    body := *openapiclient.NewMsgVpnMqttRetainCache() // MsgVpnMqttRetainCache | The MQTT Retain Cache object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnMqttRetainCache(context.Background(), msgVpnName, cacheName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnMqttRetainCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnMqttRetainCache`: MsgVpnMqttRetainCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnMqttRetainCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the MQTT Retain Cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnMqttRetainCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnMqttRetainCache**](MsgVpnMqttRetainCache.md) | The MQTT Retain Cache object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttRetainCacheResponse**](MsgVpnMqttRetainCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnMqttSession

> MsgVpnMqttSessionResponse ReplaceMsgVpnMqttSession(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace an MQTT Session object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    mqttSessionClientId := "mqttSessionClientId_example" // string | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
    mqttSessionVirtualRouter := "mqttSessionVirtualRouter_example" // string | The virtual router of the MQTT Session.
    body := *openapiclient.NewMsgVpnMqttSession() // MsgVpnMqttSession | The MQTT Session object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnMqttSession(context.Background(), msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnMqttSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnMqttSession`: MsgVpnMqttSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnMqttSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**mqttSessionClientId** | **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
**mqttSessionVirtualRouter** | **string** | The virtual router of the MQTT Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnMqttSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnMqttSession**](MsgVpnMqttSession.md) | The MQTT Session object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionResponse**](MsgVpnMqttSessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnMqttSessionSubscription

> MsgVpnMqttSessionSubscriptionResponse ReplaceMsgVpnMqttSessionSubscription(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, subscriptionTopic).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Subscription object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    mqttSessionClientId := "mqttSessionClientId_example" // string | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
    mqttSessionVirtualRouter := "mqttSessionVirtualRouter_example" // string | The virtual router of the MQTT Session.
    subscriptionTopic := "subscriptionTopic_example" // string | The MQTT subscription topic.
    body := *openapiclient.NewMsgVpnMqttSessionSubscription() // MsgVpnMqttSessionSubscription | The Subscription object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnMqttSessionSubscription(context.Background(), msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, subscriptionTopic).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnMqttSessionSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnMqttSessionSubscription`: MsgVpnMqttSessionSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnMqttSessionSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**mqttSessionClientId** | **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
**mqttSessionVirtualRouter** | **string** | The virtual router of the MQTT Session. | 
**subscriptionTopic** | **string** | The MQTT subscription topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnMqttSessionSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**MsgVpnMqttSessionSubscription**](MsgVpnMqttSessionSubscription.md) | The Subscription object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionSubscriptionResponse**](MsgVpnMqttSessionSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnQueue

> MsgVpnQueueResponse ReplaceMsgVpnQueue(ctx, msgVpnName, queueName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Queue object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    queueName := "queueName_example" // string | The name of the Queue.
    body := *openapiclient.NewMsgVpnQueue() // MsgVpnQueue | The Queue object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnQueue(context.Background(), msgVpnName, queueName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnQueue`: MsgVpnQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The name of the Queue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnQueue**](MsgVpnQueue.md) | The Queue object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueResponse**](MsgVpnQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnQueueTemplate

> MsgVpnQueueTemplateResponse ReplaceMsgVpnQueueTemplate(ctx, msgVpnName, queueTemplateName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Queue Template object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    queueTemplateName := "queueTemplateName_example" // string | The name of the Queue Template.
    body := *openapiclient.NewMsgVpnQueueTemplate() // MsgVpnQueueTemplate | The Queue Template object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnQueueTemplate(context.Background(), msgVpnName, queueTemplateName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnQueueTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnQueueTemplate`: MsgVpnQueueTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnQueueTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueTemplateName** | **string** | The name of the Queue Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnQueueTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnQueueTemplate**](MsgVpnQueueTemplate.md) | The Queue Template object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueTemplateResponse**](MsgVpnQueueTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnReplayLog

> MsgVpnReplayLogResponse ReplaceMsgVpnReplayLog(ctx, msgVpnName, replayLogName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Replay Log object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    replayLogName := "replayLogName_example" // string | The name of the Replay Log.
    body := *openapiclient.NewMsgVpnReplayLog() // MsgVpnReplayLog | The Replay Log object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnReplayLog(context.Background(), msgVpnName, replayLogName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnReplayLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnReplayLog`: MsgVpnReplayLogResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnReplayLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**replayLogName** | **string** | The name of the Replay Log. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnReplayLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnReplayLog**](MsgVpnReplayLog.md) | The Replay Log object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogResponse**](MsgVpnReplayLogResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnReplicatedTopic

> MsgVpnReplicatedTopicResponse ReplaceMsgVpnReplicatedTopic(ctx, msgVpnName, replicatedTopic).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Replicated Topic object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    replicatedTopic := "replicatedTopic_example" // string | The topic for applying replication. Published messages matching this topic will be replicated to the standby site.
    body := *openapiclient.NewMsgVpnReplicatedTopic() // MsgVpnReplicatedTopic | The Replicated Topic object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnReplicatedTopic(context.Background(), msgVpnName, replicatedTopic).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnReplicatedTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnReplicatedTopic`: MsgVpnReplicatedTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnReplicatedTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**replicatedTopic** | **string** | The topic for applying replication. Published messages matching this topic will be replicated to the standby site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnReplicatedTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnReplicatedTopic**](MsgVpnReplicatedTopic.md) | The Replicated Topic object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplicatedTopicResponse**](MsgVpnReplicatedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnRestDeliveryPoint

> MsgVpnRestDeliveryPointResponse ReplaceMsgVpnRestDeliveryPoint(ctx, msgVpnName, restDeliveryPointName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a REST Delivery Point object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    body := *openapiclient.NewMsgVpnRestDeliveryPoint() // MsgVpnRestDeliveryPoint | The REST Delivery Point object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnRestDeliveryPoint(context.Background(), msgVpnName, restDeliveryPointName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnRestDeliveryPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnRestDeliveryPoint`: MsgVpnRestDeliveryPointResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnRestDeliveryPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnRestDeliveryPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnRestDeliveryPoint**](MsgVpnRestDeliveryPoint.md) | The REST Delivery Point object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointResponse**](MsgVpnRestDeliveryPointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnRestDeliveryPointQueueBinding

> MsgVpnRestDeliveryPointQueueBindingResponse ReplaceMsgVpnRestDeliveryPointQueueBinding(ctx, msgVpnName, restDeliveryPointName, queueBindingName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Queue Binding object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    queueBindingName := "queueBindingName_example" // string | The name of a queue in the Message VPN.
    body := *openapiclient.NewMsgVpnRestDeliveryPointQueueBinding() // MsgVpnRestDeliveryPointQueueBinding | The Queue Binding object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnRestDeliveryPointQueueBinding(context.Background(), msgVpnName, restDeliveryPointName, queueBindingName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnRestDeliveryPointQueueBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnRestDeliveryPointQueueBinding`: MsgVpnRestDeliveryPointQueueBindingResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnRestDeliveryPointQueueBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 
**queueBindingName** | **string** | The name of a queue in the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnRestDeliveryPointQueueBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnRestDeliveryPointQueueBinding**](MsgVpnRestDeliveryPointQueueBinding.md) | The Queue Binding object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingResponse**](MsgVpnRestDeliveryPointQueueBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnRestDeliveryPointRestConsumer

> MsgVpnRestDeliveryPointRestConsumerResponse ReplaceMsgVpnRestDeliveryPointRestConsumer(ctx, msgVpnName, restDeliveryPointName, restConsumerName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a REST Consumer object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    restConsumerName := "restConsumerName_example" // string | The name of the REST Consumer.
    body := *openapiclient.NewMsgVpnRestDeliveryPointRestConsumer() // MsgVpnRestDeliveryPointRestConsumer | The REST Consumer object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnRestDeliveryPointRestConsumer(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnRestDeliveryPointRestConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnRestDeliveryPointRestConsumer`: MsgVpnRestDeliveryPointRestConsumerResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnRestDeliveryPointRestConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 
**restConsumerName** | **string** | The name of the REST Consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnRestDeliveryPointRestConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnRestDeliveryPointRestConsumer**](MsgVpnRestDeliveryPointRestConsumer.md) | The REST Consumer object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerResponse**](MsgVpnRestDeliveryPointRestConsumerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnTopicEndpoint

> MsgVpnTopicEndpointResponse ReplaceMsgVpnTopicEndpoint(ctx, msgVpnName, topicEndpointName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Topic Endpoint object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    topicEndpointName := "topicEndpointName_example" // string | The name of the Topic Endpoint.
    body := *openapiclient.NewMsgVpnTopicEndpoint() // MsgVpnTopicEndpoint | The Topic Endpoint object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnTopicEndpoint(context.Background(), msgVpnName, topicEndpointName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnTopicEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnTopicEndpoint`: MsgVpnTopicEndpointResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnTopicEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicEndpointName** | **string** | The name of the Topic Endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnTopicEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnTopicEndpoint**](MsgVpnTopicEndpoint.md) | The Topic Endpoint object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointResponse**](MsgVpnTopicEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnTopicEndpointTemplate

> MsgVpnTopicEndpointTemplateResponse ReplaceMsgVpnTopicEndpointTemplate(ctx, msgVpnName, topicEndpointTemplateName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Topic Endpoint Template object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    topicEndpointTemplateName := "topicEndpointTemplateName_example" // string | The name of the Topic Endpoint Template.
    body := *openapiclient.NewMsgVpnTopicEndpointTemplate() // MsgVpnTopicEndpointTemplate | The Topic Endpoint Template object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceMsgVpnTopicEndpointTemplate(context.Background(), msgVpnName, topicEndpointTemplateName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceMsgVpnTopicEndpointTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnTopicEndpointTemplate`: MsgVpnTopicEndpointTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceMsgVpnTopicEndpointTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicEndpointTemplateName** | **string** | The name of the Topic Endpoint Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnTopicEndpointTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnTopicEndpointTemplate**](MsgVpnTopicEndpointTemplate.md) | The Topic Endpoint Template object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplateResponse**](MsgVpnTopicEndpointTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceVirtualHostname

> VirtualHostnameResponse ReplaceVirtualHostname(ctx, virtualHostname).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Virtual Hostname object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    virtualHostname := "virtualHostname_example" // string | The virtual hostname.
    body := *openapiclient.NewVirtualHostname() // VirtualHostname | The Virtual Hostname object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.ReplaceVirtualHostname(context.Background(), virtualHostname).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.ReplaceVirtualHostname``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceVirtualHostname`: VirtualHostnameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.ReplaceVirtualHostname`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualHostname** | **string** | The virtual hostname. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceVirtualHostnameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VirtualHostname**](VirtualHostname.md) | The Virtual Hostname object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**VirtualHostnameResponse**](VirtualHostnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBroker

> BrokerResponse UpdateBroker(ctx).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Broker object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewBroker() // Broker | The Broker object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateBroker(context.Background()).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateBroker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBroker`: BrokerResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateBroker`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Broker**](Broker.md) | The Broker object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**BrokerResponse**](BrokerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertAuthority

> CertAuthorityResponse UpdateCertAuthority(ctx, certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Certificate Authority object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    body := *openapiclient.NewCertAuthority() // CertAuthority | The Certificate Authority object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateCertAuthority(context.Background(), certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertAuthority`: CertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CertAuthority**](CertAuthority.md) | The Certificate Authority object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityResponse**](CertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClientCertAuthority

> ClientCertAuthorityResponse UpdateClientCertAuthority(ctx, certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Client Certificate Authority object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    body := *openapiclient.NewClientCertAuthority() // ClientCertAuthority | The Client Certificate Authority object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateClientCertAuthority(context.Background(), certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateClientCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClientCertAuthority`: ClientCertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateClientCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ClientCertAuthority**](ClientCertAuthority.md) | The Client Certificate Authority object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthorityResponse**](ClientCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDmrCluster

> DmrClusterResponse UpdateDmrCluster(ctx, dmrClusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Cluster object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    body := *openapiclient.NewDmrCluster() // DmrCluster | The Cluster object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateDmrCluster(context.Background(), dmrClusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateDmrCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDmrCluster`: DmrClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateDmrCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDmrClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DmrCluster**](DmrCluster.md) | The Cluster object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterResponse**](DmrClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDmrClusterLink

> DmrClusterLinkResponse UpdateDmrClusterLink(ctx, dmrClusterName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Link object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    body := *openapiclient.NewDmrClusterLink() // DmrClusterLink | The Link object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateDmrClusterLink(context.Background(), dmrClusterName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateDmrClusterLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDmrClusterLink`: DmrClusterLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateDmrClusterLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDmrClusterLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DmrClusterLink**](DmrClusterLink.md) | The Link object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkResponse**](DmrClusterLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomainCertAuthority

> DomainCertAuthorityResponse UpdateDomainCertAuthority(ctx, certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Domain Certificate Authority object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    body := *openapiclient.NewDomainCertAuthority() // DomainCertAuthority | The Domain Certificate Authority object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateDomainCertAuthority(context.Background(), certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateDomainCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDomainCertAuthority`: DomainCertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateDomainCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DomainCertAuthority**](DomainCertAuthority.md) | The Domain Certificate Authority object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DomainCertAuthorityResponse**](DomainCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpn

> MsgVpnResponse UpdateMsgVpn(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Message VPN object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpn() // MsgVpn | The Message VPN object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpn(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpn`: MsgVpnResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpn**](MsgVpn.md) | The Message VPN object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnResponse**](MsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnAclProfile

> MsgVpnAclProfileResponse UpdateMsgVpnAclProfile(ctx, msgVpnName, aclProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update an ACL Profile object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    body := *openapiclient.NewMsgVpnAclProfile() // MsgVpnAclProfile | The ACL Profile object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnAclProfile(context.Background(), msgVpnName, aclProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnAclProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnAclProfile`: MsgVpnAclProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnAclProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnAclProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnAclProfile**](MsgVpnAclProfile.md) | The ACL Profile object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileResponse**](MsgVpnAclProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnAuthenticationOauthProvider

> MsgVpnAuthenticationOauthProviderResponse UpdateMsgVpnAuthenticationOauthProvider(ctx, msgVpnName, oauthProviderName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update an OAuth Provider object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    oauthProviderName := "oauthProviderName_example" // string | The name of the OAuth Provider.
    body := *openapiclient.NewMsgVpnAuthenticationOauthProvider() // MsgVpnAuthenticationOauthProvider | The OAuth Provider object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnAuthenticationOauthProvider(context.Background(), msgVpnName, oauthProviderName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnAuthenticationOauthProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnAuthenticationOauthProvider`: MsgVpnAuthenticationOauthProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnAuthenticationOauthProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**oauthProviderName** | **string** | The name of the OAuth Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnAuthenticationOauthProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnAuthenticationOauthProvider**](MsgVpnAuthenticationOauthProvider.md) | The OAuth Provider object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProviderResponse**](MsgVpnAuthenticationOauthProviderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnAuthorizationGroup

> MsgVpnAuthorizationGroupResponse UpdateMsgVpnAuthorizationGroup(ctx, msgVpnName, authorizationGroupName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update an LDAP Authorization Group object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    authorizationGroupName := "authorizationGroupName_example" // string | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as '#', '+', ';', '=' as the value of the group name returned from the LDAP server might prepend those characters with '\\'. For example a group name called 'test#,lab,com' will be returned from the LDAP server as 'test\\#,lab,com'.
    body := *openapiclient.NewMsgVpnAuthorizationGroup() // MsgVpnAuthorizationGroup | The LDAP Authorization Group object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnAuthorizationGroup(context.Background(), msgVpnName, authorizationGroupName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnAuthorizationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnAuthorizationGroup`: MsgVpnAuthorizationGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnAuthorizationGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**authorizationGroupName** | **string** | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as &#39;#&#39;, &#39;+&#39;, &#39;;&#39;, &#39;&#x3D;&#39; as the value of the group name returned from the LDAP server might prepend those characters with &#39;\\&#39;. For example a group name called &#39;test#,lab,com&#39; will be returned from the LDAP server as &#39;test\\#,lab,com&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnAuthorizationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnAuthorizationGroup**](MsgVpnAuthorizationGroup.md) | The LDAP Authorization Group object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupResponse**](MsgVpnAuthorizationGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnBridge

> MsgVpnBridgeResponse UpdateMsgVpnBridge(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Bridge object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    body := *openapiclient.NewMsgVpnBridge() // MsgVpnBridge | The Bridge object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnBridge(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnBridge`: MsgVpnBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnBridge**](MsgVpnBridge.md) | The Bridge object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnBridgeRemoteMsgVpn

> MsgVpnBridgeRemoteMsgVpnResponse UpdateMsgVpnBridgeRemoteMsgVpn(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Remote Message VPN object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    remoteMsgVpnName := "remoteMsgVpnName_example" // string | The name of the remote Message VPN.
    remoteMsgVpnLocation := "remoteMsgVpnLocation_example" // string | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \"v:\").
    remoteMsgVpnInterface := "remoteMsgVpnInterface_example" // string | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, `remoteMsgVpnLocation` must not be a virtual router name.
    body := *openapiclient.NewMsgVpnBridgeRemoteMsgVpn() // MsgVpnBridgeRemoteMsgVpn | The Remote Message VPN object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnBridgeRemoteMsgVpn(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnBridgeRemoteMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnBridgeRemoteMsgVpn`: MsgVpnBridgeRemoteMsgVpnResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnBridgeRemoteMsgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**remoteMsgVpnName** | **string** | The name of the remote Message VPN. | 
**remoteMsgVpnLocation** | **string** | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \&quot;v:\&quot;). | 
**remoteMsgVpnInterface** | **string** | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, &#x60;remoteMsgVpnLocation&#x60; must not be a virtual router name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnBridgeRemoteMsgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **body** | [**MsgVpnBridgeRemoteMsgVpn**](MsgVpnBridgeRemoteMsgVpn.md) | The Remote Message VPN object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteMsgVpnResponse**](MsgVpnBridgeRemoteMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnClientProfile

> MsgVpnClientProfileResponse UpdateMsgVpnClientProfile(ctx, msgVpnName, clientProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Client Profile object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    clientProfileName := "clientProfileName_example" // string | The name of the Client Profile.
    body := *openapiclient.NewMsgVpnClientProfile() // MsgVpnClientProfile | The Client Profile object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnClientProfile(context.Background(), msgVpnName, clientProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnClientProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnClientProfile`: MsgVpnClientProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnClientProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientProfileName** | **string** | The name of the Client Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnClientProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnClientProfile**](MsgVpnClientProfile.md) | The Client Profile object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientProfileResponse**](MsgVpnClientProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnClientUsername

> MsgVpnClientUsernameResponse UpdateMsgVpnClientUsername(ctx, msgVpnName, clientUsername).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Client Username object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    clientUsername := "clientUsername_example" // string | The name of the Client Username.
    body := *openapiclient.NewMsgVpnClientUsername() // MsgVpnClientUsername | The Client Username object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnClientUsername(context.Background(), msgVpnName, clientUsername).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnClientUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnClientUsername`: MsgVpnClientUsernameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnClientUsername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientUsername** | **string** | The name of the Client Username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnClientUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnClientUsername**](MsgVpnClientUsername.md) | The Client Username object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernameResponse**](MsgVpnClientUsernameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnDistributedCache

> MsgVpnDistributedCacheResponse UpdateMsgVpnDistributedCache(ctx, msgVpnName, cacheName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Distributed Cache object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    body := *openapiclient.NewMsgVpnDistributedCache() // MsgVpnDistributedCache | The Distributed Cache object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnDistributedCache(context.Background(), msgVpnName, cacheName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnDistributedCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnDistributedCache`: MsgVpnDistributedCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnDistributedCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnDistributedCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnDistributedCache**](MsgVpnDistributedCache.md) | The Distributed Cache object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheResponse**](MsgVpnDistributedCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnDistributedCacheCluster

> MsgVpnDistributedCacheClusterResponse UpdateMsgVpnDistributedCacheCluster(ctx, msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Cache Cluster object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    body := *openapiclient.NewMsgVpnDistributedCacheCluster() // MsgVpnDistributedCacheCluster | The Cache Cluster object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnDistributedCacheCluster(context.Background(), msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnDistributedCacheCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnDistributedCacheCluster`: MsgVpnDistributedCacheClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnDistributedCacheCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnDistributedCacheClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnDistributedCacheCluster**](MsgVpnDistributedCacheCluster.md) | The Cache Cluster object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterResponse**](MsgVpnDistributedCacheClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnDistributedCacheClusterInstance

> MsgVpnDistributedCacheClusterInstanceResponse UpdateMsgVpnDistributedCacheClusterInstance(ctx, msgVpnName, cacheName, clusterName, instanceName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Cache Instance object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    instanceName := "instanceName_example" // string | The name of the Cache Instance.
    body := *openapiclient.NewMsgVpnDistributedCacheClusterInstance() // MsgVpnDistributedCacheClusterInstance | The Cache Instance object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnDistributedCacheClusterInstance(context.Background(), msgVpnName, cacheName, clusterName, instanceName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnDistributedCacheClusterInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnDistributedCacheClusterInstance`: MsgVpnDistributedCacheClusterInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnDistributedCacheClusterInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**instanceName** | **string** | The name of the Cache Instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnDistributedCacheClusterInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**MsgVpnDistributedCacheClusterInstance**](MsgVpnDistributedCacheClusterInstance.md) | The Cache Instance object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceResponse**](MsgVpnDistributedCacheClusterInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnDmrBridge

> MsgVpnDmrBridgeResponse UpdateMsgVpnDmrBridge(ctx, msgVpnName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a DMR Bridge object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the DMR Bridge.
    body := *openapiclient.NewMsgVpnDmrBridge() // MsgVpnDmrBridge | The DMR Bridge object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnDmrBridge(context.Background(), msgVpnName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnDmrBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnDmrBridge`: MsgVpnDmrBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnDmrBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the DMR Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnDmrBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnDmrBridge**](MsgVpnDmrBridge.md) | The DMR Bridge object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDmrBridgeResponse**](MsgVpnDmrBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnJndiConnectionFactory

> MsgVpnJndiConnectionFactoryResponse UpdateMsgVpnJndiConnectionFactory(ctx, msgVpnName, connectionFactoryName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a JNDI Connection Factory object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    connectionFactoryName := "connectionFactoryName_example" // string | The name of the JMS Connection Factory.
    body := *openapiclient.NewMsgVpnJndiConnectionFactory() // MsgVpnJndiConnectionFactory | The JNDI Connection Factory object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnJndiConnectionFactory(context.Background(), msgVpnName, connectionFactoryName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnJndiConnectionFactory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnJndiConnectionFactory`: MsgVpnJndiConnectionFactoryResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnJndiConnectionFactory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**connectionFactoryName** | **string** | The name of the JMS Connection Factory. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnJndiConnectionFactoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnJndiConnectionFactory**](MsgVpnJndiConnectionFactory.md) | The JNDI Connection Factory object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiConnectionFactoryResponse**](MsgVpnJndiConnectionFactoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnJndiQueue

> MsgVpnJndiQueueResponse UpdateMsgVpnJndiQueue(ctx, msgVpnName, queueName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a JNDI Queue object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    queueName := "queueName_example" // string | The JNDI name of the JMS Queue.
    body := *openapiclient.NewMsgVpnJndiQueue() // MsgVpnJndiQueue | The JNDI Queue object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnJndiQueue(context.Background(), msgVpnName, queueName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnJndiQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnJndiQueue`: MsgVpnJndiQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnJndiQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The JNDI name of the JMS Queue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnJndiQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnJndiQueue**](MsgVpnJndiQueue.md) | The JNDI Queue object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiQueueResponse**](MsgVpnJndiQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnJndiTopic

> MsgVpnJndiTopicResponse UpdateMsgVpnJndiTopic(ctx, msgVpnName, topicName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a JNDI Topic object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    topicName := "topicName_example" // string | The JNDI name of the JMS Topic.
    body := *openapiclient.NewMsgVpnJndiTopic() // MsgVpnJndiTopic | The JNDI Topic object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnJndiTopic(context.Background(), msgVpnName, topicName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnJndiTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnJndiTopic`: MsgVpnJndiTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnJndiTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicName** | **string** | The JNDI name of the JMS Topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnJndiTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnJndiTopic**](MsgVpnJndiTopic.md) | The JNDI Topic object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiTopicResponse**](MsgVpnJndiTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnMqttRetainCache

> MsgVpnMqttRetainCacheResponse UpdateMsgVpnMqttRetainCache(ctx, msgVpnName, cacheName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update an MQTT Retain Cache object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    cacheName := "cacheName_example" // string | The name of the MQTT Retain Cache.
    body := *openapiclient.NewMsgVpnMqttRetainCache() // MsgVpnMqttRetainCache | The MQTT Retain Cache object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnMqttRetainCache(context.Background(), msgVpnName, cacheName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnMqttRetainCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnMqttRetainCache`: MsgVpnMqttRetainCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnMqttRetainCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the MQTT Retain Cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnMqttRetainCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnMqttRetainCache**](MsgVpnMqttRetainCache.md) | The MQTT Retain Cache object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttRetainCacheResponse**](MsgVpnMqttRetainCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnMqttSession

> MsgVpnMqttSessionResponse UpdateMsgVpnMqttSession(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update an MQTT Session object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    mqttSessionClientId := "mqttSessionClientId_example" // string | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
    mqttSessionVirtualRouter := "mqttSessionVirtualRouter_example" // string | The virtual router of the MQTT Session.
    body := *openapiclient.NewMsgVpnMqttSession() // MsgVpnMqttSession | The MQTT Session object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnMqttSession(context.Background(), msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnMqttSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnMqttSession`: MsgVpnMqttSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnMqttSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**mqttSessionClientId** | **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
**mqttSessionVirtualRouter** | **string** | The virtual router of the MQTT Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnMqttSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnMqttSession**](MsgVpnMqttSession.md) | The MQTT Session object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionResponse**](MsgVpnMqttSessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnMqttSessionSubscription

> MsgVpnMqttSessionSubscriptionResponse UpdateMsgVpnMqttSessionSubscription(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, subscriptionTopic).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Subscription object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    mqttSessionClientId := "mqttSessionClientId_example" // string | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
    mqttSessionVirtualRouter := "mqttSessionVirtualRouter_example" // string | The virtual router of the MQTT Session.
    subscriptionTopic := "subscriptionTopic_example" // string | The MQTT subscription topic.
    body := *openapiclient.NewMsgVpnMqttSessionSubscription() // MsgVpnMqttSessionSubscription | The Subscription object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnMqttSessionSubscription(context.Background(), msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, subscriptionTopic).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnMqttSessionSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnMqttSessionSubscription`: MsgVpnMqttSessionSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnMqttSessionSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**mqttSessionClientId** | **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
**mqttSessionVirtualRouter** | **string** | The virtual router of the MQTT Session. | 
**subscriptionTopic** | **string** | The MQTT subscription topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnMqttSessionSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**MsgVpnMqttSessionSubscription**](MsgVpnMqttSessionSubscription.md) | The Subscription object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionSubscriptionResponse**](MsgVpnMqttSessionSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnQueue

> MsgVpnQueueResponse UpdateMsgVpnQueue(ctx, msgVpnName, queueName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Queue object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    queueName := "queueName_example" // string | The name of the Queue.
    body := *openapiclient.NewMsgVpnQueue() // MsgVpnQueue | The Queue object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnQueue(context.Background(), msgVpnName, queueName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnQueue`: MsgVpnQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The name of the Queue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnQueue**](MsgVpnQueue.md) | The Queue object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueResponse**](MsgVpnQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnQueueTemplate

> MsgVpnQueueTemplateResponse UpdateMsgVpnQueueTemplate(ctx, msgVpnName, queueTemplateName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Queue Template object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    queueTemplateName := "queueTemplateName_example" // string | The name of the Queue Template.
    body := *openapiclient.NewMsgVpnQueueTemplate() // MsgVpnQueueTemplate | The Queue Template object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnQueueTemplate(context.Background(), msgVpnName, queueTemplateName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnQueueTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnQueueTemplate`: MsgVpnQueueTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnQueueTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueTemplateName** | **string** | The name of the Queue Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnQueueTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnQueueTemplate**](MsgVpnQueueTemplate.md) | The Queue Template object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueTemplateResponse**](MsgVpnQueueTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnReplayLog

> MsgVpnReplayLogResponse UpdateMsgVpnReplayLog(ctx, msgVpnName, replayLogName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Replay Log object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    replayLogName := "replayLogName_example" // string | The name of the Replay Log.
    body := *openapiclient.NewMsgVpnReplayLog() // MsgVpnReplayLog | The Replay Log object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnReplayLog(context.Background(), msgVpnName, replayLogName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnReplayLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnReplayLog`: MsgVpnReplayLogResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnReplayLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**replayLogName** | **string** | The name of the Replay Log. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnReplayLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnReplayLog**](MsgVpnReplayLog.md) | The Replay Log object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogResponse**](MsgVpnReplayLogResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnReplicatedTopic

> MsgVpnReplicatedTopicResponse UpdateMsgVpnReplicatedTopic(ctx, msgVpnName, replicatedTopic).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Replicated Topic object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    replicatedTopic := "replicatedTopic_example" // string | The topic for applying replication. Published messages matching this topic will be replicated to the standby site.
    body := *openapiclient.NewMsgVpnReplicatedTopic() // MsgVpnReplicatedTopic | The Replicated Topic object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnReplicatedTopic(context.Background(), msgVpnName, replicatedTopic).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnReplicatedTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnReplicatedTopic`: MsgVpnReplicatedTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnReplicatedTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**replicatedTopic** | **string** | The topic for applying replication. Published messages matching this topic will be replicated to the standby site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnReplicatedTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnReplicatedTopic**](MsgVpnReplicatedTopic.md) | The Replicated Topic object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplicatedTopicResponse**](MsgVpnReplicatedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnRestDeliveryPoint

> MsgVpnRestDeliveryPointResponse UpdateMsgVpnRestDeliveryPoint(ctx, msgVpnName, restDeliveryPointName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a REST Delivery Point object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    body := *openapiclient.NewMsgVpnRestDeliveryPoint() // MsgVpnRestDeliveryPoint | The REST Delivery Point object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnRestDeliveryPoint(context.Background(), msgVpnName, restDeliveryPointName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnRestDeliveryPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnRestDeliveryPoint`: MsgVpnRestDeliveryPointResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnRestDeliveryPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnRestDeliveryPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnRestDeliveryPoint**](MsgVpnRestDeliveryPoint.md) | The REST Delivery Point object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointResponse**](MsgVpnRestDeliveryPointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnRestDeliveryPointQueueBinding

> MsgVpnRestDeliveryPointQueueBindingResponse UpdateMsgVpnRestDeliveryPointQueueBinding(ctx, msgVpnName, restDeliveryPointName, queueBindingName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Queue Binding object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    queueBindingName := "queueBindingName_example" // string | The name of a queue in the Message VPN.
    body := *openapiclient.NewMsgVpnRestDeliveryPointQueueBinding() // MsgVpnRestDeliveryPointQueueBinding | The Queue Binding object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnRestDeliveryPointQueueBinding(context.Background(), msgVpnName, restDeliveryPointName, queueBindingName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnRestDeliveryPointQueueBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnRestDeliveryPointQueueBinding`: MsgVpnRestDeliveryPointQueueBindingResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnRestDeliveryPointQueueBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 
**queueBindingName** | **string** | The name of a queue in the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnRestDeliveryPointQueueBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnRestDeliveryPointQueueBinding**](MsgVpnRestDeliveryPointQueueBinding.md) | The Queue Binding object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingResponse**](MsgVpnRestDeliveryPointQueueBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnRestDeliveryPointRestConsumer

> MsgVpnRestDeliveryPointRestConsumerResponse UpdateMsgVpnRestDeliveryPointRestConsumer(ctx, msgVpnName, restDeliveryPointName, restConsumerName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a REST Consumer object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    restDeliveryPointName := "restDeliveryPointName_example" // string | The name of the REST Delivery Point.
    restConsumerName := "restConsumerName_example" // string | The name of the REST Consumer.
    body := *openapiclient.NewMsgVpnRestDeliveryPointRestConsumer() // MsgVpnRestDeliveryPointRestConsumer | The REST Consumer object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnRestDeliveryPointRestConsumer(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnRestDeliveryPointRestConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnRestDeliveryPointRestConsumer`: MsgVpnRestDeliveryPointRestConsumerResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnRestDeliveryPointRestConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**restDeliveryPointName** | **string** | The name of the REST Delivery Point. | 
**restConsumerName** | **string** | The name of the REST Consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnRestDeliveryPointRestConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnRestDeliveryPointRestConsumer**](MsgVpnRestDeliveryPointRestConsumer.md) | The REST Consumer object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerResponse**](MsgVpnRestDeliveryPointRestConsumerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnTopicEndpoint

> MsgVpnTopicEndpointResponse UpdateMsgVpnTopicEndpoint(ctx, msgVpnName, topicEndpointName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Topic Endpoint object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    topicEndpointName := "topicEndpointName_example" // string | The name of the Topic Endpoint.
    body := *openapiclient.NewMsgVpnTopicEndpoint() // MsgVpnTopicEndpoint | The Topic Endpoint object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnTopicEndpoint(context.Background(), msgVpnName, topicEndpointName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnTopicEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnTopicEndpoint`: MsgVpnTopicEndpointResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnTopicEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicEndpointName** | **string** | The name of the Topic Endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnTopicEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnTopicEndpoint**](MsgVpnTopicEndpoint.md) | The Topic Endpoint object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointResponse**](MsgVpnTopicEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnTopicEndpointTemplate

> MsgVpnTopicEndpointTemplateResponse UpdateMsgVpnTopicEndpointTemplate(ctx, msgVpnName, topicEndpointTemplateName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Topic Endpoint Template object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    topicEndpointTemplateName := "topicEndpointTemplateName_example" // string | The name of the Topic Endpoint Template.
    body := *openapiclient.NewMsgVpnTopicEndpointTemplate() // MsgVpnTopicEndpointTemplate | The Topic Endpoint Template object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateMsgVpnTopicEndpointTemplate(context.Background(), msgVpnName, topicEndpointTemplateName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateMsgVpnTopicEndpointTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnTopicEndpointTemplate`: MsgVpnTopicEndpointTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateMsgVpnTopicEndpointTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicEndpointTemplateName** | **string** | The name of the Topic Endpoint Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnTopicEndpointTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnTopicEndpointTemplate**](MsgVpnTopicEndpointTemplate.md) | The Topic Endpoint Template object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplateResponse**](MsgVpnTopicEndpointTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualHostname

> VirtualHostnameResponse UpdateVirtualHostname(ctx, virtualHostname).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Virtual Hostname object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    virtualHostname := "virtualHostname_example" // string | The virtual hostname.
    body := *openapiclient.NewVirtualHostname() // VirtualHostname | The Virtual Hostname object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllApi.UpdateVirtualHostname(context.Background(), virtualHostname).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllApi.UpdateVirtualHostname``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualHostname`: VirtualHostnameResponse
    fmt.Fprintf(os.Stdout, "Response from `AllApi.UpdateVirtualHostname`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualHostname** | **string** | The virtual hostname. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualHostnameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VirtualHostname**](VirtualHostname.md) | The Virtual Hostname object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**VirtualHostnameResponse**](VirtualHostnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

