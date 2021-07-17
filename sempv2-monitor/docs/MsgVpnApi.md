# \MsgVpnApi

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsgVpn**](MsgVpnApi.md#GetMsgVpn) | **Get** /msgVpns/{msgVpnName} | Get a Message VPN object.
[**GetMsgVpnAclProfile**](MsgVpnApi.md#GetMsgVpnAclProfile) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Get an ACL Profile object.
[**GetMsgVpnAclProfileClientConnectException**](MsgVpnApi.md#GetMsgVpnAclProfileClientConnectException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress} | Get a Client Connect Exception object.
[**GetMsgVpnAclProfileClientConnectExceptions**](MsgVpnApi.md#GetMsgVpnAclProfileClientConnectExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions | Get a list of Client Connect Exception objects.
[**GetMsgVpnAclProfilePublishException**](MsgVpnApi.md#GetMsgVpnAclProfilePublishException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic} | Get a Publish Topic Exception object.
[**GetMsgVpnAclProfilePublishExceptions**](MsgVpnApi.md#GetMsgVpnAclProfilePublishExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions | Get a list of Publish Topic Exception objects.
[**GetMsgVpnAclProfilePublishTopicException**](MsgVpnApi.md#GetMsgVpnAclProfilePublishTopicException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions/{publishTopicExceptionSyntax},{publishTopicException} | Get a Publish Topic Exception object.
[**GetMsgVpnAclProfilePublishTopicExceptions**](MsgVpnApi.md#GetMsgVpnAclProfilePublishTopicExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions | Get a list of Publish Topic Exception objects.
[**GetMsgVpnAclProfileSubscribeException**](MsgVpnApi.md#GetMsgVpnAclProfileSubscribeException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic} | Get a Subscribe Topic Exception object.
[**GetMsgVpnAclProfileSubscribeExceptions**](MsgVpnApi.md#GetMsgVpnAclProfileSubscribeExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions | Get a list of Subscribe Topic Exception objects.
[**GetMsgVpnAclProfileSubscribeShareNameException**](MsgVpnApi.md#GetMsgVpnAclProfileSubscribeShareNameException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions/{subscribeShareNameExceptionSyntax},{subscribeShareNameException} | Get a Subscribe Share Name Exception object.
[**GetMsgVpnAclProfileSubscribeShareNameExceptions**](MsgVpnApi.md#GetMsgVpnAclProfileSubscribeShareNameExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions | Get a list of Subscribe Share Name Exception objects.
[**GetMsgVpnAclProfileSubscribeTopicException**](MsgVpnApi.md#GetMsgVpnAclProfileSubscribeTopicException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions/{subscribeTopicExceptionSyntax},{subscribeTopicException} | Get a Subscribe Topic Exception object.
[**GetMsgVpnAclProfileSubscribeTopicExceptions**](MsgVpnApi.md#GetMsgVpnAclProfileSubscribeTopicExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions | Get a list of Subscribe Topic Exception objects.
[**GetMsgVpnAclProfiles**](MsgVpnApi.md#GetMsgVpnAclProfiles) | **Get** /msgVpns/{msgVpnName}/aclProfiles | Get a list of ACL Profile objects.
[**GetMsgVpnAuthenticationOauthProvider**](MsgVpnApi.md#GetMsgVpnAuthenticationOauthProvider) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Get an OAuth Provider object.
[**GetMsgVpnAuthenticationOauthProviders**](MsgVpnApi.md#GetMsgVpnAuthenticationOauthProviders) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders | Get a list of OAuth Provider objects.
[**GetMsgVpnAuthorizationGroup**](MsgVpnApi.md#GetMsgVpnAuthorizationGroup) | **Get** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Get an LDAP Authorization Group object.
[**GetMsgVpnAuthorizationGroups**](MsgVpnApi.md#GetMsgVpnAuthorizationGroups) | **Get** /msgVpns/{msgVpnName}/authorizationGroups | Get a list of LDAP Authorization Group objects.
[**GetMsgVpnBridge**](MsgVpnApi.md#GetMsgVpnBridge) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Get a Bridge object.
[**GetMsgVpnBridgeLocalSubscription**](MsgVpnApi.md#GetMsgVpnBridgeLocalSubscription) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/localSubscriptions/{localSubscriptionTopic} | Get a Bridge Local Subscriptions object.
[**GetMsgVpnBridgeLocalSubscriptions**](MsgVpnApi.md#GetMsgVpnBridgeLocalSubscriptions) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/localSubscriptions | Get a list of Bridge Local Subscriptions objects.
[**GetMsgVpnBridgeRemoteMsgVpn**](MsgVpnApi.md#GetMsgVpnBridgeRemoteMsgVpn) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Get a Remote Message VPN object.
[**GetMsgVpnBridgeRemoteMsgVpns**](MsgVpnApi.md#GetMsgVpnBridgeRemoteMsgVpns) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns | Get a list of Remote Message VPN objects.
[**GetMsgVpnBridgeRemoteSubscription**](MsgVpnApi.md#GetMsgVpnBridgeRemoteSubscription) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic} | Get a Remote Subscription object.
[**GetMsgVpnBridgeRemoteSubscriptions**](MsgVpnApi.md#GetMsgVpnBridgeRemoteSubscriptions) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions | Get a list of Remote Subscription objects.
[**GetMsgVpnBridgeTlsTrustedCommonName**](MsgVpnApi.md#GetMsgVpnBridgeTlsTrustedCommonName) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
[**GetMsgVpnBridgeTlsTrustedCommonNames**](MsgVpnApi.md#GetMsgVpnBridgeTlsTrustedCommonNames) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
[**GetMsgVpnBridges**](MsgVpnApi.md#GetMsgVpnBridges) | **Get** /msgVpns/{msgVpnName}/bridges | Get a list of Bridge objects.
[**GetMsgVpnClient**](MsgVpnApi.md#GetMsgVpnClient) | **Get** /msgVpns/{msgVpnName}/clients/{clientName} | Get a Client object.
[**GetMsgVpnClientConnection**](MsgVpnApi.md#GetMsgVpnClientConnection) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/connections/{clientAddress} | Get a Client Connection object.
[**GetMsgVpnClientConnections**](MsgVpnApi.md#GetMsgVpnClientConnections) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/connections | Get a list of Client Connection objects.
[**GetMsgVpnClientProfile**](MsgVpnApi.md#GetMsgVpnClientProfile) | **Get** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Get a Client Profile object.
[**GetMsgVpnClientProfiles**](MsgVpnApi.md#GetMsgVpnClientProfiles) | **Get** /msgVpns/{msgVpnName}/clientProfiles | Get a list of Client Profile objects.
[**GetMsgVpnClientRxFlow**](MsgVpnApi.md#GetMsgVpnClientRxFlow) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/rxFlows/{flowId} | Get a Client Receive Flow object.
[**GetMsgVpnClientRxFlows**](MsgVpnApi.md#GetMsgVpnClientRxFlows) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/rxFlows | Get a list of Client Receive Flow objects.
[**GetMsgVpnClientSubscription**](MsgVpnApi.md#GetMsgVpnClientSubscription) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/subscriptions/{subscriptionTopic} | Get a Client Subscription object.
[**GetMsgVpnClientSubscriptions**](MsgVpnApi.md#GetMsgVpnClientSubscriptions) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/subscriptions | Get a list of Client Subscription objects.
[**GetMsgVpnClientTransactedSession**](MsgVpnApi.md#GetMsgVpnClientTransactedSession) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions/{sessionName} | Get a Client Transacted Session object.
[**GetMsgVpnClientTransactedSessions**](MsgVpnApi.md#GetMsgVpnClientTransactedSessions) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions | Get a list of Client Transacted Session objects.
[**GetMsgVpnClientTxFlow**](MsgVpnApi.md#GetMsgVpnClientTxFlow) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/txFlows/{flowId} | Get a Client Transmit Flow object.
[**GetMsgVpnClientTxFlows**](MsgVpnApi.md#GetMsgVpnClientTxFlows) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/txFlows | Get a list of Client Transmit Flow objects.
[**GetMsgVpnClientUsername**](MsgVpnApi.md#GetMsgVpnClientUsername) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Get a Client Username object.
[**GetMsgVpnClientUsernames**](MsgVpnApi.md#GetMsgVpnClientUsernames) | **Get** /msgVpns/{msgVpnName}/clientUsernames | Get a list of Client Username objects.
[**GetMsgVpnClients**](MsgVpnApi.md#GetMsgVpnClients) | **Get** /msgVpns/{msgVpnName}/clients | Get a list of Client objects.
[**GetMsgVpnConfigSyncRemoteNode**](MsgVpnApi.md#GetMsgVpnConfigSyncRemoteNode) | **Get** /msgVpns/{msgVpnName}/configSyncRemoteNodes/{remoteNodeName} | Get a Config Sync Remote Node object.
[**GetMsgVpnConfigSyncRemoteNodes**](MsgVpnApi.md#GetMsgVpnConfigSyncRemoteNodes) | **Get** /msgVpns/{msgVpnName}/configSyncRemoteNodes | Get a list of Config Sync Remote Node objects.
[**GetMsgVpnDistributedCache**](MsgVpnApi.md#GetMsgVpnDistributedCache) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Get a Distributed Cache object.
[**GetMsgVpnDistributedCacheCluster**](MsgVpnApi.md#GetMsgVpnDistributedCacheCluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Get a Cache Cluster object.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName} | Get a Home Cache Cluster object.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix} | Get a Topic Prefix object.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes | Get a list of Topic Prefix objects.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters | Get a list of Home Cache Cluster objects.
[**GetMsgVpnDistributedCacheClusterInstance**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterInstance) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Get a Cache Instance object.
[**GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteGlobalCachingHomeClusters/{homeClusterName} | Get a Remote Home Cache Cluster object.
[**GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteGlobalCachingHomeClusters | Get a list of Remote Home Cache Cluster objects.
[**GetMsgVpnDistributedCacheClusterInstanceRemoteTopic**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterInstanceRemoteTopic) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteTopics/{topic} | Get a Remote Topic object.
[**GetMsgVpnDistributedCacheClusterInstanceRemoteTopics**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterInstanceRemoteTopics) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteTopics | Get a list of Remote Topic objects.
[**GetMsgVpnDistributedCacheClusterInstances**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterInstances) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances | Get a list of Cache Instance objects.
[**GetMsgVpnDistributedCacheClusterTopic**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterTopic) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics/{topic} | Get a Topic object.
[**GetMsgVpnDistributedCacheClusterTopics**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterTopics) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics | Get a list of Topic objects.
[**GetMsgVpnDistributedCacheClusters**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters | Get a list of Cache Cluster objects.
[**GetMsgVpnDistributedCaches**](MsgVpnApi.md#GetMsgVpnDistributedCaches) | **Get** /msgVpns/{msgVpnName}/distributedCaches | Get a list of Distributed Cache objects.
[**GetMsgVpnDmrBridge**](MsgVpnApi.md#GetMsgVpnDmrBridge) | **Get** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Get a DMR Bridge object.
[**GetMsgVpnDmrBridges**](MsgVpnApi.md#GetMsgVpnDmrBridges) | **Get** /msgVpns/{msgVpnName}/dmrBridges | Get a list of DMR Bridge objects.
[**GetMsgVpnJndiConnectionFactories**](MsgVpnApi.md#GetMsgVpnJndiConnectionFactories) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories | Get a list of JNDI Connection Factory objects.
[**GetMsgVpnJndiConnectionFactory**](MsgVpnApi.md#GetMsgVpnJndiConnectionFactory) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Get a JNDI Connection Factory object.
[**GetMsgVpnJndiQueue**](MsgVpnApi.md#GetMsgVpnJndiQueue) | **Get** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Get a JNDI Queue object.
[**GetMsgVpnJndiQueues**](MsgVpnApi.md#GetMsgVpnJndiQueues) | **Get** /msgVpns/{msgVpnName}/jndiQueues | Get a list of JNDI Queue objects.
[**GetMsgVpnJndiTopic**](MsgVpnApi.md#GetMsgVpnJndiTopic) | **Get** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Get a JNDI Topic object.
[**GetMsgVpnJndiTopics**](MsgVpnApi.md#GetMsgVpnJndiTopics) | **Get** /msgVpns/{msgVpnName}/jndiTopics | Get a list of JNDI Topic objects.
[**GetMsgVpnMqttRetainCache**](MsgVpnApi.md#GetMsgVpnMqttRetainCache) | **Get** /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName} | Get an MQTT Retain Cache object.
[**GetMsgVpnMqttRetainCaches**](MsgVpnApi.md#GetMsgVpnMqttRetainCaches) | **Get** /msgVpns/{msgVpnName}/mqttRetainCaches | Get a list of MQTT Retain Cache objects.
[**GetMsgVpnMqttSession**](MsgVpnApi.md#GetMsgVpnMqttSession) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Get an MQTT Session object.
[**GetMsgVpnMqttSessionSubscription**](MsgVpnApi.md#GetMsgVpnMqttSessionSubscription) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Get a Subscription object.
[**GetMsgVpnMqttSessionSubscriptions**](MsgVpnApi.md#GetMsgVpnMqttSessionSubscriptions) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions | Get a list of Subscription objects.
[**GetMsgVpnMqttSessions**](MsgVpnApi.md#GetMsgVpnMqttSessions) | **Get** /msgVpns/{msgVpnName}/mqttSessions | Get a list of MQTT Session objects.
[**GetMsgVpnQueue**](MsgVpnApi.md#GetMsgVpnQueue) | **Get** /msgVpns/{msgVpnName}/queues/{queueName} | Get a Queue object.
[**GetMsgVpnQueueMsg**](MsgVpnApi.md#GetMsgVpnQueueMsg) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs/{msgId} | Get a Queue Message object.
[**GetMsgVpnQueueMsgs**](MsgVpnApi.md#GetMsgVpnQueueMsgs) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs | Get a list of Queue Message objects.
[**GetMsgVpnQueuePriorities**](MsgVpnApi.md#GetMsgVpnQueuePriorities) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/priorities | Get a list of Queue Priority objects.
[**GetMsgVpnQueuePriority**](MsgVpnApi.md#GetMsgVpnQueuePriority) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/priorities/{priority} | Get a Queue Priority object.
[**GetMsgVpnQueueSubscription**](MsgVpnApi.md#GetMsgVpnQueueSubscription) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic} | Get a Queue Subscription object.
[**GetMsgVpnQueueSubscriptions**](MsgVpnApi.md#GetMsgVpnQueueSubscriptions) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions | Get a list of Queue Subscription objects.
[**GetMsgVpnQueueTemplate**](MsgVpnApi.md#GetMsgVpnQueueTemplate) | **Get** /msgVpns/{msgVpnName}/queueTemplates/{queueTemplateName} | Get a Queue Template object.
[**GetMsgVpnQueueTemplates**](MsgVpnApi.md#GetMsgVpnQueueTemplates) | **Get** /msgVpns/{msgVpnName}/queueTemplates | Get a list of Queue Template objects.
[**GetMsgVpnQueueTxFlow**](MsgVpnApi.md#GetMsgVpnQueueTxFlow) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/txFlows/{flowId} | Get a Queue Transmit Flow object.
[**GetMsgVpnQueueTxFlows**](MsgVpnApi.md#GetMsgVpnQueueTxFlows) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/txFlows | Get a list of Queue Transmit Flow objects.
[**GetMsgVpnQueues**](MsgVpnApi.md#GetMsgVpnQueues) | **Get** /msgVpns/{msgVpnName}/queues | Get a list of Queue objects.
[**GetMsgVpnReplayLog**](MsgVpnApi.md#GetMsgVpnReplayLog) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Get a Replay Log object.
[**GetMsgVpnReplayLogMsg**](MsgVpnApi.md#GetMsgVpnReplayLogMsg) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/msgs/{msgId} | Get a Message object.
[**GetMsgVpnReplayLogMsgs**](MsgVpnApi.md#GetMsgVpnReplayLogMsgs) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/msgs | Get a list of Message objects.
[**GetMsgVpnReplayLogs**](MsgVpnApi.md#GetMsgVpnReplayLogs) | **Get** /msgVpns/{msgVpnName}/replayLogs | Get a list of Replay Log objects.
[**GetMsgVpnReplicatedTopic**](MsgVpnApi.md#GetMsgVpnReplicatedTopic) | **Get** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Get a Replicated Topic object.
[**GetMsgVpnReplicatedTopics**](MsgVpnApi.md#GetMsgVpnReplicatedTopics) | **Get** /msgVpns/{msgVpnName}/replicatedTopics | Get a list of Replicated Topic objects.
[**GetMsgVpnRestDeliveryPoint**](MsgVpnApi.md#GetMsgVpnRestDeliveryPoint) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Get a REST Delivery Point object.
[**GetMsgVpnRestDeliveryPointQueueBinding**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointQueueBinding) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Get a Queue Binding object.
[**GetMsgVpnRestDeliveryPointQueueBindings**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointQueueBindings) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings | Get a list of Queue Binding objects.
[**GetMsgVpnRestDeliveryPointRestConsumer**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointRestConsumer) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Get a REST Consumer object.
[**GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims/{oauthJwtClaimName} | Get a Claim object.
[**GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims | Get a list of Claim objects.
[**GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
[**GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
[**GetMsgVpnRestDeliveryPointRestConsumers**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointRestConsumers) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Get a list of REST Consumer objects.
[**GetMsgVpnRestDeliveryPoints**](MsgVpnApi.md#GetMsgVpnRestDeliveryPoints) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints | Get a list of REST Delivery Point objects.
[**GetMsgVpnTopicEndpoint**](MsgVpnApi.md#GetMsgVpnTopicEndpoint) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Get a Topic Endpoint object.
[**GetMsgVpnTopicEndpointMsg**](MsgVpnApi.md#GetMsgVpnTopicEndpointMsg) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId} | Get a Topic Endpoint Message object.
[**GetMsgVpnTopicEndpointMsgs**](MsgVpnApi.md#GetMsgVpnTopicEndpointMsgs) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs | Get a list of Topic Endpoint Message objects.
[**GetMsgVpnTopicEndpointPriorities**](MsgVpnApi.md#GetMsgVpnTopicEndpointPriorities) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/priorities | Get a list of Topic Endpoint Priority objects.
[**GetMsgVpnTopicEndpointPriority**](MsgVpnApi.md#GetMsgVpnTopicEndpointPriority) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/priorities/{priority} | Get a Topic Endpoint Priority object.
[**GetMsgVpnTopicEndpointTemplate**](MsgVpnApi.md#GetMsgVpnTopicEndpointTemplate) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Get a Topic Endpoint Template object.
[**GetMsgVpnTopicEndpointTemplates**](MsgVpnApi.md#GetMsgVpnTopicEndpointTemplates) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates | Get a list of Topic Endpoint Template objects.
[**GetMsgVpnTopicEndpointTxFlow**](MsgVpnApi.md#GetMsgVpnTopicEndpointTxFlow) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/txFlows/{flowId} | Get a Topic Endpoint Transmit Flow object.
[**GetMsgVpnTopicEndpointTxFlows**](MsgVpnApi.md#GetMsgVpnTopicEndpointTxFlows) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/txFlows | Get a list of Topic Endpoint Transmit Flow objects.
[**GetMsgVpnTopicEndpoints**](MsgVpnApi.md#GetMsgVpnTopicEndpoints) | **Get** /msgVpns/{msgVpnName}/topicEndpoints | Get a list of Topic Endpoint objects.
[**GetMsgVpnTransaction**](MsgVpnApi.md#GetMsgVpnTransaction) | **Get** /msgVpns/{msgVpnName}/transactions/{xid} | Get a Replicated Local Transaction or XA Transaction object.
[**GetMsgVpnTransactionConsumerMsg**](MsgVpnApi.md#GetMsgVpnTransactionConsumerMsg) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/consumerMsgs/{msgId} | Get a Transaction Consumer Message object.
[**GetMsgVpnTransactionConsumerMsgs**](MsgVpnApi.md#GetMsgVpnTransactionConsumerMsgs) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/consumerMsgs | Get a list of Transaction Consumer Message objects.
[**GetMsgVpnTransactionPublisherMsg**](MsgVpnApi.md#GetMsgVpnTransactionPublisherMsg) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/publisherMsgs/{msgId} | Get a Transaction Publisher Message object.
[**GetMsgVpnTransactionPublisherMsgs**](MsgVpnApi.md#GetMsgVpnTransactionPublisherMsgs) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/publisherMsgs | Get a list of Transaction Publisher Message objects.
[**GetMsgVpnTransactions**](MsgVpnApi.md#GetMsgVpnTransactions) | **Get** /msgVpns/{msgVpnName}/transactions | Get a list of Replicated Local Transaction or XA Transaction objects.
[**GetMsgVpns**](MsgVpnApi.md#GetMsgVpns) | **Get** /msgVpns | Get a list of Message VPN objects.



## GetMsgVpn

> MsgVpnResponse GetMsgVpn(ctx, msgVpnName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpn(context.Background(), msgVpnName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpn`: MsgVpnResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpn`: %v\n", resp)
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

> MsgVpnAclProfileResponse GetMsgVpnAclProfile(ctx, msgVpnName, aclProfileName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAclProfile(context.Background(), msgVpnName, aclProfileName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAclProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfile`: MsgVpnAclProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAclProfile`: %v\n", resp)
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

> MsgVpnAclProfileClientConnectExceptionResponse GetMsgVpnAclProfileClientConnectException(ctx, msgVpnName, aclProfileName, clientConnectExceptionAddress).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAclProfileClientConnectException(context.Background(), msgVpnName, aclProfileName, clientConnectExceptionAddress).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAclProfileClientConnectException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileClientConnectException`: MsgVpnAclProfileClientConnectExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAclProfileClientConnectException`: %v\n", resp)
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

> MsgVpnAclProfileClientConnectExceptionsResponse GetMsgVpnAclProfileClientConnectExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAclProfileClientConnectExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAclProfileClientConnectExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileClientConnectExceptions`: MsgVpnAclProfileClientConnectExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAclProfileClientConnectExceptions`: %v\n", resp)
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

> MsgVpnAclProfilePublishExceptionResponse GetMsgVpnAclProfilePublishException(ctx, msgVpnName, aclProfileName, topicSyntax, publishExceptionTopic).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAclProfilePublishException(context.Background(), msgVpnName, aclProfileName, topicSyntax, publishExceptionTopic).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAclProfilePublishException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfilePublishException`: MsgVpnAclProfilePublishExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAclProfilePublishException`: %v\n", resp)
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

> MsgVpnAclProfilePublishExceptionsResponse GetMsgVpnAclProfilePublishExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAclProfilePublishExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAclProfilePublishExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfilePublishExceptions`: MsgVpnAclProfilePublishExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAclProfilePublishExceptions`: %v\n", resp)
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

> MsgVpnAclProfilePublishTopicExceptionResponse GetMsgVpnAclProfilePublishTopicException(ctx, msgVpnName, aclProfileName, publishTopicExceptionSyntax, publishTopicException).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAclProfilePublishTopicException(context.Background(), msgVpnName, aclProfileName, publishTopicExceptionSyntax, publishTopicException).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAclProfilePublishTopicException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfilePublishTopicException`: MsgVpnAclProfilePublishTopicExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAclProfilePublishTopicException`: %v\n", resp)
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

> MsgVpnAclProfilePublishTopicExceptionsResponse GetMsgVpnAclProfilePublishTopicExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAclProfilePublishTopicExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAclProfilePublishTopicExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfilePublishTopicExceptions`: MsgVpnAclProfilePublishTopicExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAclProfilePublishTopicExceptions`: %v\n", resp)
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

> MsgVpnAclProfileSubscribeExceptionResponse GetMsgVpnAclProfileSubscribeException(ctx, msgVpnName, aclProfileName, topicSyntax, subscribeExceptionTopic).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAclProfileSubscribeException(context.Background(), msgVpnName, aclProfileName, topicSyntax, subscribeExceptionTopic).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAclProfileSubscribeException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeException`: MsgVpnAclProfileSubscribeExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAclProfileSubscribeException`: %v\n", resp)
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

> MsgVpnAclProfileSubscribeExceptionsResponse GetMsgVpnAclProfileSubscribeExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAclProfileSubscribeExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAclProfileSubscribeExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeExceptions`: MsgVpnAclProfileSubscribeExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAclProfileSubscribeExceptions`: %v\n", resp)
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

> MsgVpnAclProfileSubscribeShareNameExceptionResponse GetMsgVpnAclProfileSubscribeShareNameException(ctx, msgVpnName, aclProfileName, subscribeShareNameExceptionSyntax, subscribeShareNameException).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAclProfileSubscribeShareNameException(context.Background(), msgVpnName, aclProfileName, subscribeShareNameExceptionSyntax, subscribeShareNameException).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAclProfileSubscribeShareNameException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeShareNameException`: MsgVpnAclProfileSubscribeShareNameExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAclProfileSubscribeShareNameException`: %v\n", resp)
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

> MsgVpnAclProfileSubscribeShareNameExceptionsResponse GetMsgVpnAclProfileSubscribeShareNameExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAclProfileSubscribeShareNameExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAclProfileSubscribeShareNameExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeShareNameExceptions`: MsgVpnAclProfileSubscribeShareNameExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAclProfileSubscribeShareNameExceptions`: %v\n", resp)
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

> MsgVpnAclProfileSubscribeTopicExceptionResponse GetMsgVpnAclProfileSubscribeTopicException(ctx, msgVpnName, aclProfileName, subscribeTopicExceptionSyntax, subscribeTopicException).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAclProfileSubscribeTopicException(context.Background(), msgVpnName, aclProfileName, subscribeTopicExceptionSyntax, subscribeTopicException).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAclProfileSubscribeTopicException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeTopicException`: MsgVpnAclProfileSubscribeTopicExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAclProfileSubscribeTopicException`: %v\n", resp)
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

> MsgVpnAclProfileSubscribeTopicExceptionsResponse GetMsgVpnAclProfileSubscribeTopicExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAclProfileSubscribeTopicExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAclProfileSubscribeTopicExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeTopicExceptions`: MsgVpnAclProfileSubscribeTopicExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAclProfileSubscribeTopicExceptions`: %v\n", resp)
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

> MsgVpnAclProfilesResponse GetMsgVpnAclProfiles(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAclProfiles(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAclProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfiles`: MsgVpnAclProfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAclProfiles`: %v\n", resp)
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

> MsgVpnAuthenticationOauthProviderResponse GetMsgVpnAuthenticationOauthProvider(ctx, msgVpnName, oauthProviderName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAuthenticationOauthProvider(context.Background(), msgVpnName, oauthProviderName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAuthenticationOauthProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAuthenticationOauthProvider`: MsgVpnAuthenticationOauthProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAuthenticationOauthProvider`: %v\n", resp)
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

> MsgVpnAuthenticationOauthProvidersResponse GetMsgVpnAuthenticationOauthProviders(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAuthenticationOauthProviders(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAuthenticationOauthProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAuthenticationOauthProviders`: MsgVpnAuthenticationOauthProvidersResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAuthenticationOauthProviders`: %v\n", resp)
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

> MsgVpnAuthorizationGroupResponse GetMsgVpnAuthorizationGroup(ctx, msgVpnName, authorizationGroupName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAuthorizationGroup(context.Background(), msgVpnName, authorizationGroupName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAuthorizationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAuthorizationGroup`: MsgVpnAuthorizationGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAuthorizationGroup`: %v\n", resp)
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

> MsgVpnAuthorizationGroupsResponse GetMsgVpnAuthorizationGroups(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnAuthorizationGroups(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnAuthorizationGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAuthorizationGroups`: MsgVpnAuthorizationGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnAuthorizationGroups`: %v\n", resp)
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

> MsgVpnBridgeResponse GetMsgVpnBridge(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnBridge(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridge`: MsgVpnBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnBridge`: %v\n", resp)
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


## GetMsgVpnBridgeLocalSubscription

> MsgVpnBridgeLocalSubscriptionResponse GetMsgVpnBridgeLocalSubscription(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, localSubscriptionTopic).Select_(select_).Execute()

Get a Bridge Local Subscriptions object.



### Example

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
    localSubscriptionTopic := "localSubscriptionTopic_example" // string | The topic of the Bridge local subscription.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnBridgeLocalSubscription(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, localSubscriptionTopic).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnBridgeLocalSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeLocalSubscription`: MsgVpnBridgeLocalSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnBridgeLocalSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**localSubscriptionTopic** | **string** | The topic of the Bridge local subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgeLocalSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeLocalSubscriptionResponse**](MsgVpnBridgeLocalSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridgeLocalSubscriptions

> MsgVpnBridgeLocalSubscriptionsResponse GetMsgVpnBridgeLocalSubscriptions(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Bridge Local Subscriptions objects.



### Example

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnBridgeLocalSubscriptions(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnBridgeLocalSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeLocalSubscriptions`: MsgVpnBridgeLocalSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnBridgeLocalSubscriptions`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetMsgVpnBridgeLocalSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeLocalSubscriptionsResponse**](MsgVpnBridgeLocalSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridgeRemoteMsgVpn

> MsgVpnBridgeRemoteMsgVpnResponse GetMsgVpnBridgeRemoteMsgVpn(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnBridgeRemoteMsgVpn(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnBridgeRemoteMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeRemoteMsgVpn`: MsgVpnBridgeRemoteMsgVpnResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnBridgeRemoteMsgVpn`: %v\n", resp)
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

> MsgVpnBridgeRemoteMsgVpnsResponse GetMsgVpnBridgeRemoteMsgVpns(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnBridgeRemoteMsgVpns(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnBridgeRemoteMsgVpns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeRemoteMsgVpns`: MsgVpnBridgeRemoteMsgVpnsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnBridgeRemoteMsgVpns`: %v\n", resp)
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

> MsgVpnBridgeRemoteSubscriptionResponse GetMsgVpnBridgeRemoteSubscription(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteSubscriptionTopic).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnBridgeRemoteSubscription(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, remoteSubscriptionTopic).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnBridgeRemoteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeRemoteSubscription`: MsgVpnBridgeRemoteSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnBridgeRemoteSubscription`: %v\n", resp)
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

> MsgVpnBridgeRemoteSubscriptionsResponse GetMsgVpnBridgeRemoteSubscriptions(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnBridgeRemoteSubscriptions(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnBridgeRemoteSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeRemoteSubscriptions`: MsgVpnBridgeRemoteSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnBridgeRemoteSubscriptions`: %v\n", resp)
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

> MsgVpnBridgeTlsTrustedCommonNameResponse GetMsgVpnBridgeTlsTrustedCommonName(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, tlsTrustedCommonName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnBridgeTlsTrustedCommonName(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, tlsTrustedCommonName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnBridgeTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeTlsTrustedCommonName`: MsgVpnBridgeTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnBridgeTlsTrustedCommonName`: %v\n", resp)
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

> MsgVpnBridgeTlsTrustedCommonNamesResponse GetMsgVpnBridgeTlsTrustedCommonNames(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnBridgeTlsTrustedCommonNames(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnBridgeTlsTrustedCommonNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeTlsTrustedCommonNames`: MsgVpnBridgeTlsTrustedCommonNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnBridgeTlsTrustedCommonNames`: %v\n", resp)
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

> MsgVpnBridgesResponse GetMsgVpnBridges(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnBridges(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnBridges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridges`: MsgVpnBridgesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnBridges`: %v\n", resp)
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


## GetMsgVpnClient

> MsgVpnClientResponse GetMsgVpnClient(ctx, msgVpnName, clientName).Select_(select_).Execute()

Get a Client object.



### Example

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
    clientName := "clientName_example" // string | The name of the Client.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnClient(context.Background(), msgVpnName, clientName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClient`: MsgVpnClientResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientResponse**](MsgVpnClientResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientConnection

> MsgVpnClientConnectionResponse GetMsgVpnClientConnection(ctx, msgVpnName, clientName, clientAddress).Select_(select_).Execute()

Get a Client Connection object.



### Example

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
    clientName := "clientName_example" // string | The name of the Client.
    clientAddress := "clientAddress_example" // string | The IP address and TCP port on the Client side of the Client Connection.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnClientConnection(context.Background(), msgVpnName, clientName, clientAddress).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnClientConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientConnection`: MsgVpnClientConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnClientConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 
**clientAddress** | **string** | The IP address and TCP port on the Client side of the Client Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientConnectionResponse**](MsgVpnClientConnectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientConnections

> MsgVpnClientConnectionsResponse GetMsgVpnClientConnections(ctx, msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Client Connection objects.



### Example

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
    clientName := "clientName_example" // string | The name of the Client.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnClientConnections(context.Background(), msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnClientConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientConnections`: MsgVpnClientConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnClientConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientConnectionsResponse**](MsgVpnClientConnectionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientProfile

> MsgVpnClientProfileResponse GetMsgVpnClientProfile(ctx, msgVpnName, clientProfileName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnClientProfile(context.Background(), msgVpnName, clientProfileName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnClientProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientProfile`: MsgVpnClientProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnClientProfile`: %v\n", resp)
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

> MsgVpnClientProfilesResponse GetMsgVpnClientProfiles(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnClientProfiles(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnClientProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientProfiles`: MsgVpnClientProfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnClientProfiles`: %v\n", resp)
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


## GetMsgVpnClientRxFlow

> MsgVpnClientRxFlowResponse GetMsgVpnClientRxFlow(ctx, msgVpnName, clientName, flowId).Select_(select_).Execute()

Get a Client Receive Flow object.



### Example

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
    clientName := "clientName_example" // string | The name of the Client.
    flowId := "flowId_example" // string | The identifier (ID) of the flow.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnClientRxFlow(context.Background(), msgVpnName, clientName, flowId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnClientRxFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientRxFlow`: MsgVpnClientRxFlowResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnClientRxFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 
**flowId** | **string** | The identifier (ID) of the flow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientRxFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientRxFlowResponse**](MsgVpnClientRxFlowResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientRxFlows

> MsgVpnClientRxFlowsResponse GetMsgVpnClientRxFlows(ctx, msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Client Receive Flow objects.



### Example

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
    clientName := "clientName_example" // string | The name of the Client.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnClientRxFlows(context.Background(), msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnClientRxFlows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientRxFlows`: MsgVpnClientRxFlowsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnClientRxFlows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientRxFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientRxFlowsResponse**](MsgVpnClientRxFlowsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientSubscription

> MsgVpnClientSubscriptionResponse GetMsgVpnClientSubscription(ctx, msgVpnName, clientName, subscriptionTopic).Select_(select_).Execute()

Get a Client Subscription object.



### Example

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
    clientName := "clientName_example" // string | The name of the Client.
    subscriptionTopic := "subscriptionTopic_example" // string | The topic of the Subscription.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnClientSubscription(context.Background(), msgVpnName, clientName, subscriptionTopic).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnClientSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientSubscription`: MsgVpnClientSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnClientSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 
**subscriptionTopic** | **string** | The topic of the Subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientSubscriptionResponse**](MsgVpnClientSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientSubscriptions

> MsgVpnClientSubscriptionsResponse GetMsgVpnClientSubscriptions(ctx, msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Client Subscription objects.



### Example

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
    clientName := "clientName_example" // string | The name of the Client.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnClientSubscriptions(context.Background(), msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnClientSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientSubscriptions`: MsgVpnClientSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnClientSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientSubscriptionsResponse**](MsgVpnClientSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientTransactedSession

> MsgVpnClientTransactedSessionResponse GetMsgVpnClientTransactedSession(ctx, msgVpnName, clientName, sessionName).Select_(select_).Execute()

Get a Client Transacted Session object.



### Example

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
    clientName := "clientName_example" // string | The name of the Client.
    sessionName := "sessionName_example" // string | The name of the Transacted Session.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnClientTransactedSession(context.Background(), msgVpnName, clientName, sessionName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnClientTransactedSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientTransactedSession`: MsgVpnClientTransactedSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnClientTransactedSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 
**sessionName** | **string** | The name of the Transacted Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientTransactedSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientTransactedSessionResponse**](MsgVpnClientTransactedSessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientTransactedSessions

> MsgVpnClientTransactedSessionsResponse GetMsgVpnClientTransactedSessions(ctx, msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Client Transacted Session objects.



### Example

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
    clientName := "clientName_example" // string | The name of the Client.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnClientTransactedSessions(context.Background(), msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnClientTransactedSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientTransactedSessions`: MsgVpnClientTransactedSessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnClientTransactedSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientTransactedSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientTransactedSessionsResponse**](MsgVpnClientTransactedSessionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientTxFlow

> MsgVpnClientTxFlowResponse GetMsgVpnClientTxFlow(ctx, msgVpnName, clientName, flowId).Select_(select_).Execute()

Get a Client Transmit Flow object.



### Example

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
    clientName := "clientName_example" // string | The name of the Client.
    flowId := "flowId_example" // string | The identifier (ID) of the flow.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnClientTxFlow(context.Background(), msgVpnName, clientName, flowId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnClientTxFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientTxFlow`: MsgVpnClientTxFlowResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnClientTxFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 
**flowId** | **string** | The identifier (ID) of the flow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientTxFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientTxFlowResponse**](MsgVpnClientTxFlowResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientTxFlows

> MsgVpnClientTxFlowsResponse GetMsgVpnClientTxFlows(ctx, msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Client Transmit Flow objects.



### Example

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
    clientName := "clientName_example" // string | The name of the Client.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnClientTxFlows(context.Background(), msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnClientTxFlows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientTxFlows`: MsgVpnClientTxFlowsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnClientTxFlows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientTxFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientTxFlowsResponse**](MsgVpnClientTxFlowsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientUsername

> MsgVpnClientUsernameResponse GetMsgVpnClientUsername(ctx, msgVpnName, clientUsername).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnClientUsername(context.Background(), msgVpnName, clientUsername).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnClientUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientUsername`: MsgVpnClientUsernameResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnClientUsername`: %v\n", resp)
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

> MsgVpnClientUsernamesResponse GetMsgVpnClientUsernames(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnClientUsernames(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnClientUsernames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientUsernames`: MsgVpnClientUsernamesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnClientUsernames`: %v\n", resp)
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


## GetMsgVpnClients

> MsgVpnClientsResponse GetMsgVpnClients(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Client objects.



### Example

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnClients(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClients`: MsgVpnClientsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientsResponse**](MsgVpnClientsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnConfigSyncRemoteNode

> MsgVpnConfigSyncRemoteNodeResponse GetMsgVpnConfigSyncRemoteNode(ctx, msgVpnName, remoteNodeName).Select_(select_).Execute()

Get a Config Sync Remote Node object.



### Example

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
    remoteNodeName := "remoteNodeName_example" // string | The name of the Config Sync Remote Node.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnConfigSyncRemoteNode(context.Background(), msgVpnName, remoteNodeName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnConfigSyncRemoteNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnConfigSyncRemoteNode`: MsgVpnConfigSyncRemoteNodeResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnConfigSyncRemoteNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**remoteNodeName** | **string** | The name of the Config Sync Remote Node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnConfigSyncRemoteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnConfigSyncRemoteNodeResponse**](MsgVpnConfigSyncRemoteNodeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnConfigSyncRemoteNodes

> MsgVpnConfigSyncRemoteNodesResponse GetMsgVpnConfigSyncRemoteNodes(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Config Sync Remote Node objects.



### Example

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnConfigSyncRemoteNodes(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnConfigSyncRemoteNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnConfigSyncRemoteNodes`: MsgVpnConfigSyncRemoteNodesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnConfigSyncRemoteNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnConfigSyncRemoteNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnConfigSyncRemoteNodesResponse**](MsgVpnConfigSyncRemoteNodesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCache

> MsgVpnDistributedCacheResponse GetMsgVpnDistributedCache(ctx, msgVpnName, cacheName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDistributedCache(context.Background(), msgVpnName, cacheName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDistributedCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCache`: MsgVpnDistributedCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDistributedCache`: %v\n", resp)
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

> MsgVpnDistributedCacheClusterResponse GetMsgVpnDistributedCacheCluster(ctx, msgVpnName, cacheName, clusterName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDistributedCacheCluster(context.Background(), msgVpnName, cacheName, clusterName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDistributedCacheCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheCluster`: MsgVpnDistributedCacheClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDistributedCacheCluster`: %v\n", resp)
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

> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(ctx, msgVpnName, cacheName, clusterName, homeClusterName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(context.Background(), msgVpnName, cacheName, clusterName, homeClusterName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster`: MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster`: %v\n", resp)
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

> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(ctx, msgVpnName, cacheName, clusterName, homeClusterName, topicPrefix).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(context.Background(), msgVpnName, cacheName, clusterName, homeClusterName, topicPrefix).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix`: MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix`: %v\n", resp)
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

> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes(ctx, msgVpnName, cacheName, clusterName, homeClusterName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes(context.Background(), msgVpnName, cacheName, clusterName, homeClusterName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes`: MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes`: %v\n", resp)
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

> MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters(ctx, msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters(context.Background(), msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters`: MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters`: %v\n", resp)
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

> MsgVpnDistributedCacheClusterInstanceResponse GetMsgVpnDistributedCacheClusterInstance(ctx, msgVpnName, cacheName, clusterName, instanceName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDistributedCacheClusterInstance(context.Background(), msgVpnName, cacheName, clusterName, instanceName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDistributedCacheClusterInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterInstance`: MsgVpnDistributedCacheClusterInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDistributedCacheClusterInstance`: %v\n", resp)
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


## GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster

> MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster(ctx, msgVpnName, cacheName, clusterName, instanceName, homeClusterName).Select_(select_).Execute()

Get a Remote Home Cache Cluster object.



### Example

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
    homeClusterName := "homeClusterName_example" // string | The name of the remote Home Cache Cluster.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster(context.Background(), msgVpnName, cacheName, clusterName, instanceName, homeClusterName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster`: MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster`: %v\n", resp)
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
**homeClusterName** | **string** | The name of the remote Home Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse**](MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters

> MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters(ctx, msgVpnName, cacheName, clusterName, instanceName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Remote Home Cache Cluster objects.



### Example

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
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters(context.Background(), msgVpnName, cacheName, clusterName, instanceName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters`: MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse**](MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterInstanceRemoteTopic

> MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse GetMsgVpnDistributedCacheClusterInstanceRemoteTopic(ctx, msgVpnName, cacheName, clusterName, instanceName, topic).Select_(select_).Execute()

Get a Remote Topic object.



### Example

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
    topic := "topic_example" // string | The value of the remote Topic.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDistributedCacheClusterInstanceRemoteTopic(context.Background(), msgVpnName, cacheName, clusterName, instanceName, topic).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDistributedCacheClusterInstanceRemoteTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterInstanceRemoteTopic`: MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDistributedCacheClusterInstanceRemoteTopic`: %v\n", resp)
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
**topic** | **string** | The value of the remote Topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse**](MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterInstanceRemoteTopics

> MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse GetMsgVpnDistributedCacheClusterInstanceRemoteTopics(ctx, msgVpnName, cacheName, clusterName, instanceName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Remote Topic objects.



### Example

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
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDistributedCacheClusterInstanceRemoteTopics(context.Background(), msgVpnName, cacheName, clusterName, instanceName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDistributedCacheClusterInstanceRemoteTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterInstanceRemoteTopics`: MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDistributedCacheClusterInstanceRemoteTopics`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse**](MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterInstances

> MsgVpnDistributedCacheClusterInstancesResponse GetMsgVpnDistributedCacheClusterInstances(ctx, msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDistributedCacheClusterInstances(context.Background(), msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDistributedCacheClusterInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterInstances`: MsgVpnDistributedCacheClusterInstancesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDistributedCacheClusterInstances`: %v\n", resp)
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

> MsgVpnDistributedCacheClusterTopicResponse GetMsgVpnDistributedCacheClusterTopic(ctx, msgVpnName, cacheName, clusterName, topic).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDistributedCacheClusterTopic(context.Background(), msgVpnName, cacheName, clusterName, topic).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDistributedCacheClusterTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterTopic`: MsgVpnDistributedCacheClusterTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDistributedCacheClusterTopic`: %v\n", resp)
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

> MsgVpnDistributedCacheClusterTopicsResponse GetMsgVpnDistributedCacheClusterTopics(ctx, msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDistributedCacheClusterTopics(context.Background(), msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDistributedCacheClusterTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterTopics`: MsgVpnDistributedCacheClusterTopicsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDistributedCacheClusterTopics`: %v\n", resp)
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

> MsgVpnDistributedCacheClustersResponse GetMsgVpnDistributedCacheClusters(ctx, msgVpnName, cacheName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDistributedCacheClusters(context.Background(), msgVpnName, cacheName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDistributedCacheClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusters`: MsgVpnDistributedCacheClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDistributedCacheClusters`: %v\n", resp)
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

> MsgVpnDistributedCachesResponse GetMsgVpnDistributedCaches(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDistributedCaches(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDistributedCaches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCaches`: MsgVpnDistributedCachesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDistributedCaches`: %v\n", resp)
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

> MsgVpnDmrBridgeResponse GetMsgVpnDmrBridge(ctx, msgVpnName, remoteNodeName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDmrBridge(context.Background(), msgVpnName, remoteNodeName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDmrBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDmrBridge`: MsgVpnDmrBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDmrBridge`: %v\n", resp)
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

> MsgVpnDmrBridgesResponse GetMsgVpnDmrBridges(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnDmrBridges(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnDmrBridges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDmrBridges`: MsgVpnDmrBridgesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnDmrBridges`: %v\n", resp)
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

> MsgVpnJndiConnectionFactoriesResponse GetMsgVpnJndiConnectionFactories(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnJndiConnectionFactories(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnJndiConnectionFactories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnJndiConnectionFactories`: MsgVpnJndiConnectionFactoriesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnJndiConnectionFactories`: %v\n", resp)
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

> MsgVpnJndiConnectionFactoryResponse GetMsgVpnJndiConnectionFactory(ctx, msgVpnName, connectionFactoryName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnJndiConnectionFactory(context.Background(), msgVpnName, connectionFactoryName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnJndiConnectionFactory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnJndiConnectionFactory`: MsgVpnJndiConnectionFactoryResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnJndiConnectionFactory`: %v\n", resp)
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

> MsgVpnJndiQueueResponse GetMsgVpnJndiQueue(ctx, msgVpnName, queueName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnJndiQueue(context.Background(), msgVpnName, queueName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnJndiQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnJndiQueue`: MsgVpnJndiQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnJndiQueue`: %v\n", resp)
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

> MsgVpnJndiQueuesResponse GetMsgVpnJndiQueues(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnJndiQueues(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnJndiQueues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnJndiQueues`: MsgVpnJndiQueuesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnJndiQueues`: %v\n", resp)
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

> MsgVpnJndiTopicResponse GetMsgVpnJndiTopic(ctx, msgVpnName, topicName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnJndiTopic(context.Background(), msgVpnName, topicName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnJndiTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnJndiTopic`: MsgVpnJndiTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnJndiTopic`: %v\n", resp)
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

> MsgVpnJndiTopicsResponse GetMsgVpnJndiTopics(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnJndiTopics(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnJndiTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnJndiTopics`: MsgVpnJndiTopicsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnJndiTopics`: %v\n", resp)
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

> MsgVpnMqttRetainCacheResponse GetMsgVpnMqttRetainCache(ctx, msgVpnName, cacheName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnMqttRetainCache(context.Background(), msgVpnName, cacheName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnMqttRetainCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnMqttRetainCache`: MsgVpnMqttRetainCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnMqttRetainCache`: %v\n", resp)
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

> MsgVpnMqttRetainCachesResponse GetMsgVpnMqttRetainCaches(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnMqttRetainCaches(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnMqttRetainCaches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnMqttRetainCaches`: MsgVpnMqttRetainCachesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnMqttRetainCaches`: %v\n", resp)
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

> MsgVpnMqttSessionResponse GetMsgVpnMqttSession(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnMqttSession(context.Background(), msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnMqttSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnMqttSession`: MsgVpnMqttSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnMqttSession`: %v\n", resp)
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

> MsgVpnMqttSessionSubscriptionResponse GetMsgVpnMqttSessionSubscription(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, subscriptionTopic).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnMqttSessionSubscription(context.Background(), msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, subscriptionTopic).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnMqttSessionSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnMqttSessionSubscription`: MsgVpnMqttSessionSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnMqttSessionSubscription`: %v\n", resp)
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

> MsgVpnMqttSessionSubscriptionsResponse GetMsgVpnMqttSessionSubscriptions(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnMqttSessionSubscriptions(context.Background(), msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnMqttSessionSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnMqttSessionSubscriptions`: MsgVpnMqttSessionSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnMqttSessionSubscriptions`: %v\n", resp)
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

> MsgVpnMqttSessionsResponse GetMsgVpnMqttSessions(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnMqttSessions(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnMqttSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnMqttSessions`: MsgVpnMqttSessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnMqttSessions`: %v\n", resp)
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

> MsgVpnQueueResponse GetMsgVpnQueue(ctx, msgVpnName, queueName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnQueue(context.Background(), msgVpnName, queueName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueue`: MsgVpnQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnQueue`: %v\n", resp)
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


## GetMsgVpnQueueMsg

> MsgVpnQueueMsgResponse GetMsgVpnQueueMsg(ctx, msgVpnName, queueName, msgId).Select_(select_).Execute()

Get a Queue Message object.



### Example

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
    msgId := "msgId_example" // string | The identifier (ID) of the Message.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnQueueMsg(context.Background(), msgVpnName, queueName, msgId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnQueueMsg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueueMsg`: MsgVpnQueueMsgResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnQueueMsg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The name of the Queue. | 
**msgId** | **string** | The identifier (ID) of the Message. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnQueueMsgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueMsgResponse**](MsgVpnQueueMsgResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnQueueMsgs

> MsgVpnQueueMsgsResponse GetMsgVpnQueueMsgs(ctx, msgVpnName, queueName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Queue Message objects.



### Example

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnQueueMsgs(context.Background(), msgVpnName, queueName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnQueueMsgs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueueMsgs`: MsgVpnQueueMsgsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnQueueMsgs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The name of the Queue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnQueueMsgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueMsgsResponse**](MsgVpnQueueMsgsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnQueuePriorities

> MsgVpnQueuePrioritiesResponse GetMsgVpnQueuePriorities(ctx, msgVpnName, queueName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Queue Priority objects.



### Example

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnQueuePriorities(context.Background(), msgVpnName, queueName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnQueuePriorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueuePriorities`: MsgVpnQueuePrioritiesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnQueuePriorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The name of the Queue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnQueuePrioritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueuePrioritiesResponse**](MsgVpnQueuePrioritiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnQueuePriority

> MsgVpnQueuePriorityResponse GetMsgVpnQueuePriority(ctx, msgVpnName, queueName, priority).Select_(select_).Execute()

Get a Queue Priority object.



### Example

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
    priority := "priority_example" // string | The level of the Priority, from 9 (highest) to 0 (lowest).
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnQueuePriority(context.Background(), msgVpnName, queueName, priority).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnQueuePriority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueuePriority`: MsgVpnQueuePriorityResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnQueuePriority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The name of the Queue. | 
**priority** | **string** | The level of the Priority, from 9 (highest) to 0 (lowest). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnQueuePriorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueuePriorityResponse**](MsgVpnQueuePriorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnQueueSubscription

> MsgVpnQueueSubscriptionResponse GetMsgVpnQueueSubscription(ctx, msgVpnName, queueName, subscriptionTopic).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnQueueSubscription(context.Background(), msgVpnName, queueName, subscriptionTopic).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnQueueSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueueSubscription`: MsgVpnQueueSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnQueueSubscription`: %v\n", resp)
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

> MsgVpnQueueSubscriptionsResponse GetMsgVpnQueueSubscriptions(ctx, msgVpnName, queueName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnQueueSubscriptions(context.Background(), msgVpnName, queueName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnQueueSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueueSubscriptions`: MsgVpnQueueSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnQueueSubscriptions`: %v\n", resp)
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

> MsgVpnQueueTemplateResponse GetMsgVpnQueueTemplate(ctx, msgVpnName, queueTemplateName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnQueueTemplate(context.Background(), msgVpnName, queueTemplateName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnQueueTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueueTemplate`: MsgVpnQueueTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnQueueTemplate`: %v\n", resp)
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

> MsgVpnQueueTemplatesResponse GetMsgVpnQueueTemplates(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnQueueTemplates(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnQueueTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueueTemplates`: MsgVpnQueueTemplatesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnQueueTemplates`: %v\n", resp)
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


## GetMsgVpnQueueTxFlow

> MsgVpnQueueTxFlowResponse GetMsgVpnQueueTxFlow(ctx, msgVpnName, queueName, flowId).Select_(select_).Execute()

Get a Queue Transmit Flow object.



### Example

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
    flowId := "flowId_example" // string | The identifier (ID) of the Flow.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnQueueTxFlow(context.Background(), msgVpnName, queueName, flowId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnQueueTxFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueueTxFlow`: MsgVpnQueueTxFlowResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnQueueTxFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The name of the Queue. | 
**flowId** | **string** | The identifier (ID) of the Flow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnQueueTxFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueTxFlowResponse**](MsgVpnQueueTxFlowResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnQueueTxFlows

> MsgVpnQueueTxFlowsResponse GetMsgVpnQueueTxFlows(ctx, msgVpnName, queueName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Queue Transmit Flow objects.



### Example

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnQueueTxFlows(context.Background(), msgVpnName, queueName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnQueueTxFlows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueueTxFlows`: MsgVpnQueueTxFlowsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnQueueTxFlows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**queueName** | **string** | The name of the Queue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnQueueTxFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueTxFlowsResponse**](MsgVpnQueueTxFlowsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnQueues

> MsgVpnQueuesResponse GetMsgVpnQueues(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnQueues(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnQueues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnQueues`: MsgVpnQueuesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnQueues`: %v\n", resp)
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

> MsgVpnReplayLogResponse GetMsgVpnReplayLog(ctx, msgVpnName, replayLogName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnReplayLog(context.Background(), msgVpnName, replayLogName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnReplayLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnReplayLog`: MsgVpnReplayLogResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnReplayLog`: %v\n", resp)
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


## GetMsgVpnReplayLogMsg

> MsgVpnReplayLogMsgResponse GetMsgVpnReplayLogMsg(ctx, msgVpnName, replayLogName, msgId).Select_(select_).Execute()

Get a Message object.



### Example

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
    msgId := "msgId_example" // string | The identifier (ID) of the message.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnReplayLogMsg(context.Background(), msgVpnName, replayLogName, msgId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnReplayLogMsg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnReplayLogMsg`: MsgVpnReplayLogMsgResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnReplayLogMsg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**replayLogName** | **string** | The name of the Replay Log. | 
**msgId** | **string** | The identifier (ID) of the message. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnReplayLogMsgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogMsgResponse**](MsgVpnReplayLogMsgResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnReplayLogMsgs

> MsgVpnReplayLogMsgsResponse GetMsgVpnReplayLogMsgs(ctx, msgVpnName, replayLogName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Message objects.



### Example

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
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnReplayLogMsgs(context.Background(), msgVpnName, replayLogName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnReplayLogMsgs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnReplayLogMsgs`: MsgVpnReplayLogMsgsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnReplayLogMsgs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**replayLogName** | **string** | The name of the Replay Log. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnReplayLogMsgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogMsgsResponse**](MsgVpnReplayLogMsgsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnReplayLogs

> MsgVpnReplayLogsResponse GetMsgVpnReplayLogs(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnReplayLogs(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnReplayLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnReplayLogs`: MsgVpnReplayLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnReplayLogs`: %v\n", resp)
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

> MsgVpnReplicatedTopicResponse GetMsgVpnReplicatedTopic(ctx, msgVpnName, replicatedTopic).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnReplicatedTopic(context.Background(), msgVpnName, replicatedTopic).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnReplicatedTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnReplicatedTopic`: MsgVpnReplicatedTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnReplicatedTopic`: %v\n", resp)
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

> MsgVpnReplicatedTopicsResponse GetMsgVpnReplicatedTopics(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnReplicatedTopics(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnReplicatedTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnReplicatedTopics`: MsgVpnReplicatedTopicsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnReplicatedTopics`: %v\n", resp)
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

> MsgVpnRestDeliveryPointResponse GetMsgVpnRestDeliveryPoint(ctx, msgVpnName, restDeliveryPointName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnRestDeliveryPoint(context.Background(), msgVpnName, restDeliveryPointName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnRestDeliveryPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPoint`: MsgVpnRestDeliveryPointResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnRestDeliveryPoint`: %v\n", resp)
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

> MsgVpnRestDeliveryPointQueueBindingResponse GetMsgVpnRestDeliveryPointQueueBinding(ctx, msgVpnName, restDeliveryPointName, queueBindingName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnRestDeliveryPointQueueBinding(context.Background(), msgVpnName, restDeliveryPointName, queueBindingName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnRestDeliveryPointQueueBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPointQueueBinding`: MsgVpnRestDeliveryPointQueueBindingResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnRestDeliveryPointQueueBinding`: %v\n", resp)
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

> MsgVpnRestDeliveryPointQueueBindingsResponse GetMsgVpnRestDeliveryPointQueueBindings(ctx, msgVpnName, restDeliveryPointName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnRestDeliveryPointQueueBindings(context.Background(), msgVpnName, restDeliveryPointName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnRestDeliveryPointQueueBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPointQueueBindings`: MsgVpnRestDeliveryPointQueueBindingsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnRestDeliveryPointQueueBindings`: %v\n", resp)
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

> MsgVpnRestDeliveryPointRestConsumerResponse GetMsgVpnRestDeliveryPointRestConsumer(ctx, msgVpnName, restDeliveryPointName, restConsumerName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumer(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPointRestConsumer`: MsgVpnRestDeliveryPointRestConsumerResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumer`: %v\n", resp)
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

> MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(ctx, msgVpnName, restDeliveryPointName, restConsumerName, oauthJwtClaimName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName, oauthJwtClaimName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim`: MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim`: %v\n", resp)
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

> MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims(ctx, msgVpnName, restDeliveryPointName, restConsumerName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims`: MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims`: %v\n", resp)
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

> MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(ctx, msgVpnName, restDeliveryPointName, restConsumerName, tlsTrustedCommonName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName, tlsTrustedCommonName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName`: MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName`: %v\n", resp)
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

> MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames(ctx, msgVpnName, restDeliveryPointName, restConsumerName).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames(context.Background(), msgVpnName, restDeliveryPointName, restConsumerName).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames`: MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames`: %v\n", resp)
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

> MsgVpnRestDeliveryPointRestConsumersResponse GetMsgVpnRestDeliveryPointRestConsumers(ctx, msgVpnName, restDeliveryPointName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumers(context.Background(), msgVpnName, restDeliveryPointName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPointRestConsumers`: MsgVpnRestDeliveryPointRestConsumersResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnRestDeliveryPointRestConsumers`: %v\n", resp)
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

> MsgVpnRestDeliveryPointsResponse GetMsgVpnRestDeliveryPoints(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnRestDeliveryPoints(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnRestDeliveryPoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnRestDeliveryPoints`: MsgVpnRestDeliveryPointsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnRestDeliveryPoints`: %v\n", resp)
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


## GetMsgVpnTopicEndpoint

> MsgVpnTopicEndpointResponse GetMsgVpnTopicEndpoint(ctx, msgVpnName, topicEndpointName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnTopicEndpoint(context.Background(), msgVpnName, topicEndpointName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnTopicEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTopicEndpoint`: MsgVpnTopicEndpointResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnTopicEndpoint`: %v\n", resp)
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


## GetMsgVpnTopicEndpointMsg

> MsgVpnTopicEndpointMsgResponse GetMsgVpnTopicEndpointMsg(ctx, msgVpnName, topicEndpointName, msgId).Select_(select_).Execute()

Get a Topic Endpoint Message object.



### Example

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
    msgId := "msgId_example" // string | The identifier (ID) of the Message.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnTopicEndpointMsg(context.Background(), msgVpnName, topicEndpointName, msgId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnTopicEndpointMsg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTopicEndpointMsg`: MsgVpnTopicEndpointMsgResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnTopicEndpointMsg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicEndpointName** | **string** | The name of the Topic Endpoint. | 
**msgId** | **string** | The identifier (ID) of the Message. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnTopicEndpointMsgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointMsgResponse**](MsgVpnTopicEndpointMsgResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnTopicEndpointMsgs

> MsgVpnTopicEndpointMsgsResponse GetMsgVpnTopicEndpointMsgs(ctx, msgVpnName, topicEndpointName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Topic Endpoint Message objects.



### Example

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
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnTopicEndpointMsgs(context.Background(), msgVpnName, topicEndpointName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnTopicEndpointMsgs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTopicEndpointMsgs`: MsgVpnTopicEndpointMsgsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnTopicEndpointMsgs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicEndpointName** | **string** | The name of the Topic Endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnTopicEndpointMsgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointMsgsResponse**](MsgVpnTopicEndpointMsgsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnTopicEndpointPriorities

> MsgVpnTopicEndpointPrioritiesResponse GetMsgVpnTopicEndpointPriorities(ctx, msgVpnName, topicEndpointName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Topic Endpoint Priority objects.



### Example

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
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnTopicEndpointPriorities(context.Background(), msgVpnName, topicEndpointName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnTopicEndpointPriorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTopicEndpointPriorities`: MsgVpnTopicEndpointPrioritiesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnTopicEndpointPriorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicEndpointName** | **string** | The name of the Topic Endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnTopicEndpointPrioritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointPrioritiesResponse**](MsgVpnTopicEndpointPrioritiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnTopicEndpointPriority

> MsgVpnTopicEndpointPriorityResponse GetMsgVpnTopicEndpointPriority(ctx, msgVpnName, topicEndpointName, priority).Select_(select_).Execute()

Get a Topic Endpoint Priority object.



### Example

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
    priority := "priority_example" // string | The level of the Priority, from 9 (highest) to 0 (lowest).
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnTopicEndpointPriority(context.Background(), msgVpnName, topicEndpointName, priority).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnTopicEndpointPriority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTopicEndpointPriority`: MsgVpnTopicEndpointPriorityResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnTopicEndpointPriority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicEndpointName** | **string** | The name of the Topic Endpoint. | 
**priority** | **string** | The level of the Priority, from 9 (highest) to 0 (lowest). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnTopicEndpointPriorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointPriorityResponse**](MsgVpnTopicEndpointPriorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnTopicEndpointTemplate

> MsgVpnTopicEndpointTemplateResponse GetMsgVpnTopicEndpointTemplate(ctx, msgVpnName, topicEndpointTemplateName).Select_(select_).Execute()

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
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnTopicEndpointTemplate(context.Background(), msgVpnName, topicEndpointTemplateName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnTopicEndpointTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTopicEndpointTemplate`: MsgVpnTopicEndpointTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnTopicEndpointTemplate`: %v\n", resp)
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

> MsgVpnTopicEndpointTemplatesResponse GetMsgVpnTopicEndpointTemplates(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnTopicEndpointTemplates(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnTopicEndpointTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTopicEndpointTemplates`: MsgVpnTopicEndpointTemplatesResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnTopicEndpointTemplates`: %v\n", resp)
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


## GetMsgVpnTopicEndpointTxFlow

> MsgVpnTopicEndpointTxFlowResponse GetMsgVpnTopicEndpointTxFlow(ctx, msgVpnName, topicEndpointName, flowId).Select_(select_).Execute()

Get a Topic Endpoint Transmit Flow object.



### Example

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
    flowId := "flowId_example" // string | The identifier (ID) of the Flow.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnTopicEndpointTxFlow(context.Background(), msgVpnName, topicEndpointName, flowId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnTopicEndpointTxFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTopicEndpointTxFlow`: MsgVpnTopicEndpointTxFlowResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnTopicEndpointTxFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicEndpointName** | **string** | The name of the Topic Endpoint. | 
**flowId** | **string** | The identifier (ID) of the Flow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnTopicEndpointTxFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTxFlowResponse**](MsgVpnTopicEndpointTxFlowResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnTopicEndpointTxFlows

> MsgVpnTopicEndpointTxFlowsResponse GetMsgVpnTopicEndpointTxFlows(ctx, msgVpnName, topicEndpointName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Topic Endpoint Transmit Flow objects.



### Example

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
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnTopicEndpointTxFlows(context.Background(), msgVpnName, topicEndpointName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnTopicEndpointTxFlows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTopicEndpointTxFlows`: MsgVpnTopicEndpointTxFlowsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnTopicEndpointTxFlows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**topicEndpointName** | **string** | The name of the Topic Endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnTopicEndpointTxFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTxFlowsResponse**](MsgVpnTopicEndpointTxFlowsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnTopicEndpoints

> MsgVpnTopicEndpointsResponse GetMsgVpnTopicEndpoints(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnTopicEndpoints(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnTopicEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTopicEndpoints`: MsgVpnTopicEndpointsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnTopicEndpoints`: %v\n", resp)
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


## GetMsgVpnTransaction

> MsgVpnTransactionResponse GetMsgVpnTransaction(ctx, msgVpnName, xid).Select_(select_).Execute()

Get a Replicated Local Transaction or XA Transaction object.



### Example

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
    xid := "xid_example" // string | The identifier (ID) of the Transaction.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnTransaction(context.Background(), msgVpnName, xid).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTransaction`: MsgVpnTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**xid** | **string** | The identifier (ID) of the Transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTransactionResponse**](MsgVpnTransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnTransactionConsumerMsg

> MsgVpnTransactionConsumerMsgResponse GetMsgVpnTransactionConsumerMsg(ctx, msgVpnName, xid, msgId).Select_(select_).Execute()

Get a Transaction Consumer Message object.



### Example

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
    xid := "xid_example" // string | The identifier (ID) of the Transaction.
    msgId := "msgId_example" // string | The identifier (ID) of the Message.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnTransactionConsumerMsg(context.Background(), msgVpnName, xid, msgId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnTransactionConsumerMsg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTransactionConsumerMsg`: MsgVpnTransactionConsumerMsgResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnTransactionConsumerMsg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**xid** | **string** | The identifier (ID) of the Transaction. | 
**msgId** | **string** | The identifier (ID) of the Message. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnTransactionConsumerMsgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTransactionConsumerMsgResponse**](MsgVpnTransactionConsumerMsgResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnTransactionConsumerMsgs

> MsgVpnTransactionConsumerMsgsResponse GetMsgVpnTransactionConsumerMsgs(ctx, msgVpnName, xid).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Transaction Consumer Message objects.



### Example

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
    xid := "xid_example" // string | The identifier (ID) of the Transaction.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnTransactionConsumerMsgs(context.Background(), msgVpnName, xid).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnTransactionConsumerMsgs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTransactionConsumerMsgs`: MsgVpnTransactionConsumerMsgsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnTransactionConsumerMsgs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**xid** | **string** | The identifier (ID) of the Transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnTransactionConsumerMsgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTransactionConsumerMsgsResponse**](MsgVpnTransactionConsumerMsgsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnTransactionPublisherMsg

> MsgVpnTransactionPublisherMsgResponse GetMsgVpnTransactionPublisherMsg(ctx, msgVpnName, xid, msgId).Select_(select_).Execute()

Get a Transaction Publisher Message object.



### Example

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
    xid := "xid_example" // string | The identifier (ID) of the Transaction.
    msgId := "msgId_example" // string | The identifier (ID) of the Message.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnTransactionPublisherMsg(context.Background(), msgVpnName, xid, msgId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnTransactionPublisherMsg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTransactionPublisherMsg`: MsgVpnTransactionPublisherMsgResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnTransactionPublisherMsg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**xid** | **string** | The identifier (ID) of the Transaction. | 
**msgId** | **string** | The identifier (ID) of the Message. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnTransactionPublisherMsgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTransactionPublisherMsgResponse**](MsgVpnTransactionPublisherMsgResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnTransactionPublisherMsgs

> MsgVpnTransactionPublisherMsgsResponse GetMsgVpnTransactionPublisherMsgs(ctx, msgVpnName, xid).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Transaction Publisher Message objects.



### Example

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
    xid := "xid_example" // string | The identifier (ID) of the Transaction.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnTransactionPublisherMsgs(context.Background(), msgVpnName, xid).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnTransactionPublisherMsgs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTransactionPublisherMsgs`: MsgVpnTransactionPublisherMsgsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnTransactionPublisherMsgs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**xid** | **string** | The identifier (ID) of the Transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnTransactionPublisherMsgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTransactionPublisherMsgsResponse**](MsgVpnTransactionPublisherMsgsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnTransactions

> MsgVpnTransactionsResponse GetMsgVpnTransactions(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Replicated Local Transaction or XA Transaction objects.



### Example

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpnTransactions(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpnTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnTransactions`: MsgVpnTransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpnTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTransactionsResponse**](MsgVpnTransactionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpns

> MsgVpnsResponse GetMsgVpns(ctx).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MsgVpnApi.GetMsgVpns(context.Background()).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MsgVpnApi.GetMsgVpns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpns`: MsgVpnsResponse
    fmt.Fprintf(os.Stdout, "Response from `MsgVpnApi.GetMsgVpns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
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

