# MsgVpnCollections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfiles** | Pointer to [**MsgVpnCollectionsAclprofiles**](MsgVpnCollectionsAclprofiles.md) |  | [optional] 
**AuthenticationOauthProviders** | Pointer to [**MsgVpnCollectionsAuthenticationoauthproviders**](MsgVpnCollectionsAuthenticationoauthproviders.md) |  | [optional] 
**AuthorizationGroups** | Pointer to [**MsgVpnCollectionsAuthorizationgroups**](MsgVpnCollectionsAuthorizationgroups.md) |  | [optional] 
**Bridges** | Pointer to [**MsgVpnCollectionsBridges**](MsgVpnCollectionsBridges.md) |  | [optional] 
**ClientProfiles** | Pointer to [**MsgVpnCollectionsClientprofiles**](MsgVpnCollectionsClientprofiles.md) |  | [optional] 
**ClientUsernames** | Pointer to [**MsgVpnCollectionsClientusernames**](MsgVpnCollectionsClientusernames.md) |  | [optional] 
**Clients** | Pointer to [**MsgVpnCollectionsClients**](MsgVpnCollectionsClients.md) |  | [optional] 
**ConfigSyncRemoteNodes** | Pointer to [**MsgVpnCollectionsConfigsyncremotenodes**](MsgVpnCollectionsConfigsyncremotenodes.md) |  | [optional] 
**DistributedCaches** | Pointer to [**MsgVpnCollectionsDistributedcaches**](MsgVpnCollectionsDistributedcaches.md) |  | [optional] 
**DmrBridges** | Pointer to [**MsgVpnCollectionsDmrbridges**](MsgVpnCollectionsDmrbridges.md) |  | [optional] 
**JndiConnectionFactories** | Pointer to [**MsgVpnCollectionsJndiconnectionfactories**](MsgVpnCollectionsJndiconnectionfactories.md) |  | [optional] 
**JndiQueues** | Pointer to [**MsgVpnCollectionsJndiqueues**](MsgVpnCollectionsJndiqueues.md) |  | [optional] 
**JndiTopics** | Pointer to [**MsgVpnCollectionsJnditopics**](MsgVpnCollectionsJnditopics.md) |  | [optional] 
**MqttRetainCaches** | Pointer to [**MsgVpnCollectionsMqttretaincaches**](MsgVpnCollectionsMqttretaincaches.md) |  | [optional] 
**MqttSessions** | Pointer to [**MsgVpnCollectionsMqttsessions**](MsgVpnCollectionsMqttsessions.md) |  | [optional] 
**QueueTemplates** | Pointer to [**MsgVpnCollectionsQueuetemplates**](MsgVpnCollectionsQueuetemplates.md) |  | [optional] 
**Queues** | Pointer to [**MsgVpnCollectionsQueues**](MsgVpnCollectionsQueues.md) |  | [optional] 
**ReplayLogs** | Pointer to [**MsgVpnCollectionsReplaylogs**](MsgVpnCollectionsReplaylogs.md) |  | [optional] 
**ReplicatedTopics** | Pointer to [**MsgVpnCollectionsReplicatedtopics**](MsgVpnCollectionsReplicatedtopics.md) |  | [optional] 
**RestDeliveryPoints** | Pointer to [**MsgVpnCollectionsRestdeliverypoints**](MsgVpnCollectionsRestdeliverypoints.md) |  | [optional] 
**TopicEndpointTemplates** | Pointer to [**MsgVpnCollectionsTopicendpointtemplates**](MsgVpnCollectionsTopicendpointtemplates.md) |  | [optional] 
**TopicEndpoints** | Pointer to [**MsgVpnCollectionsTopicendpoints**](MsgVpnCollectionsTopicendpoints.md) |  | [optional] 
**Transactions** | Pointer to [**MsgVpnCollectionsTransactions**](MsgVpnCollectionsTransactions.md) |  | [optional] 

## Methods

### NewMsgVpnCollections

`func NewMsgVpnCollections() *MsgVpnCollections`

NewMsgVpnCollections instantiates a new MsgVpnCollections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnCollectionsWithDefaults

`func NewMsgVpnCollectionsWithDefaults() *MsgVpnCollections`

NewMsgVpnCollectionsWithDefaults instantiates a new MsgVpnCollections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclProfiles

`func (o *MsgVpnCollections) GetAclProfiles() MsgVpnCollectionsAclprofiles`

GetAclProfiles returns the AclProfiles field if non-nil, zero value otherwise.

### GetAclProfilesOk

`func (o *MsgVpnCollections) GetAclProfilesOk() (*MsgVpnCollectionsAclprofiles, bool)`

GetAclProfilesOk returns a tuple with the AclProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclProfiles

`func (o *MsgVpnCollections) SetAclProfiles(v MsgVpnCollectionsAclprofiles)`

SetAclProfiles sets AclProfiles field to given value.

### HasAclProfiles

`func (o *MsgVpnCollections) HasAclProfiles() bool`

HasAclProfiles returns a boolean if a field has been set.

### GetAuthenticationOauthProviders

`func (o *MsgVpnCollections) GetAuthenticationOauthProviders() MsgVpnCollectionsAuthenticationoauthproviders`

GetAuthenticationOauthProviders returns the AuthenticationOauthProviders field if non-nil, zero value otherwise.

### GetAuthenticationOauthProvidersOk

`func (o *MsgVpnCollections) GetAuthenticationOauthProvidersOk() (*MsgVpnCollectionsAuthenticationoauthproviders, bool)`

GetAuthenticationOauthProvidersOk returns a tuple with the AuthenticationOauthProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthProviders

`func (o *MsgVpnCollections) SetAuthenticationOauthProviders(v MsgVpnCollectionsAuthenticationoauthproviders)`

SetAuthenticationOauthProviders sets AuthenticationOauthProviders field to given value.

### HasAuthenticationOauthProviders

`func (o *MsgVpnCollections) HasAuthenticationOauthProviders() bool`

HasAuthenticationOauthProviders returns a boolean if a field has been set.

### GetAuthorizationGroups

`func (o *MsgVpnCollections) GetAuthorizationGroups() MsgVpnCollectionsAuthorizationgroups`

GetAuthorizationGroups returns the AuthorizationGroups field if non-nil, zero value otherwise.

### GetAuthorizationGroupsOk

`func (o *MsgVpnCollections) GetAuthorizationGroupsOk() (*MsgVpnCollectionsAuthorizationgroups, bool)`

GetAuthorizationGroupsOk returns a tuple with the AuthorizationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationGroups

`func (o *MsgVpnCollections) SetAuthorizationGroups(v MsgVpnCollectionsAuthorizationgroups)`

SetAuthorizationGroups sets AuthorizationGroups field to given value.

### HasAuthorizationGroups

`func (o *MsgVpnCollections) HasAuthorizationGroups() bool`

HasAuthorizationGroups returns a boolean if a field has been set.

### GetBridges

`func (o *MsgVpnCollections) GetBridges() MsgVpnCollectionsBridges`

GetBridges returns the Bridges field if non-nil, zero value otherwise.

### GetBridgesOk

`func (o *MsgVpnCollections) GetBridgesOk() (*MsgVpnCollectionsBridges, bool)`

GetBridgesOk returns a tuple with the Bridges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridges

`func (o *MsgVpnCollections) SetBridges(v MsgVpnCollectionsBridges)`

SetBridges sets Bridges field to given value.

### HasBridges

`func (o *MsgVpnCollections) HasBridges() bool`

HasBridges returns a boolean if a field has been set.

### GetClientProfiles

`func (o *MsgVpnCollections) GetClientProfiles() MsgVpnCollectionsClientprofiles`

GetClientProfiles returns the ClientProfiles field if non-nil, zero value otherwise.

### GetClientProfilesOk

`func (o *MsgVpnCollections) GetClientProfilesOk() (*MsgVpnCollectionsClientprofiles, bool)`

GetClientProfilesOk returns a tuple with the ClientProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfiles

`func (o *MsgVpnCollections) SetClientProfiles(v MsgVpnCollectionsClientprofiles)`

SetClientProfiles sets ClientProfiles field to given value.

### HasClientProfiles

`func (o *MsgVpnCollections) HasClientProfiles() bool`

HasClientProfiles returns a boolean if a field has been set.

### GetClientUsernames

`func (o *MsgVpnCollections) GetClientUsernames() MsgVpnCollectionsClientusernames`

GetClientUsernames returns the ClientUsernames field if non-nil, zero value otherwise.

### GetClientUsernamesOk

`func (o *MsgVpnCollections) GetClientUsernamesOk() (*MsgVpnCollectionsClientusernames, bool)`

GetClientUsernamesOk returns a tuple with the ClientUsernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUsernames

`func (o *MsgVpnCollections) SetClientUsernames(v MsgVpnCollectionsClientusernames)`

SetClientUsernames sets ClientUsernames field to given value.

### HasClientUsernames

`func (o *MsgVpnCollections) HasClientUsernames() bool`

HasClientUsernames returns a boolean if a field has been set.

### GetClients

`func (o *MsgVpnCollections) GetClients() MsgVpnCollectionsClients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *MsgVpnCollections) GetClientsOk() (*MsgVpnCollectionsClients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *MsgVpnCollections) SetClients(v MsgVpnCollectionsClients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *MsgVpnCollections) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetConfigSyncRemoteNodes

`func (o *MsgVpnCollections) GetConfigSyncRemoteNodes() MsgVpnCollectionsConfigsyncremotenodes`

GetConfigSyncRemoteNodes returns the ConfigSyncRemoteNodes field if non-nil, zero value otherwise.

### GetConfigSyncRemoteNodesOk

`func (o *MsgVpnCollections) GetConfigSyncRemoteNodesOk() (*MsgVpnCollectionsConfigsyncremotenodes, bool)`

GetConfigSyncRemoteNodesOk returns a tuple with the ConfigSyncRemoteNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSyncRemoteNodes

`func (o *MsgVpnCollections) SetConfigSyncRemoteNodes(v MsgVpnCollectionsConfigsyncremotenodes)`

SetConfigSyncRemoteNodes sets ConfigSyncRemoteNodes field to given value.

### HasConfigSyncRemoteNodes

`func (o *MsgVpnCollections) HasConfigSyncRemoteNodes() bool`

HasConfigSyncRemoteNodes returns a boolean if a field has been set.

### GetDistributedCaches

`func (o *MsgVpnCollections) GetDistributedCaches() MsgVpnCollectionsDistributedcaches`

GetDistributedCaches returns the DistributedCaches field if non-nil, zero value otherwise.

### GetDistributedCachesOk

`func (o *MsgVpnCollections) GetDistributedCachesOk() (*MsgVpnCollectionsDistributedcaches, bool)`

GetDistributedCachesOk returns a tuple with the DistributedCaches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedCaches

`func (o *MsgVpnCollections) SetDistributedCaches(v MsgVpnCollectionsDistributedcaches)`

SetDistributedCaches sets DistributedCaches field to given value.

### HasDistributedCaches

`func (o *MsgVpnCollections) HasDistributedCaches() bool`

HasDistributedCaches returns a boolean if a field has been set.

### GetDmrBridges

`func (o *MsgVpnCollections) GetDmrBridges() MsgVpnCollectionsDmrbridges`

GetDmrBridges returns the DmrBridges field if non-nil, zero value otherwise.

### GetDmrBridgesOk

`func (o *MsgVpnCollections) GetDmrBridgesOk() (*MsgVpnCollectionsDmrbridges, bool)`

GetDmrBridgesOk returns a tuple with the DmrBridges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmrBridges

`func (o *MsgVpnCollections) SetDmrBridges(v MsgVpnCollectionsDmrbridges)`

SetDmrBridges sets DmrBridges field to given value.

### HasDmrBridges

`func (o *MsgVpnCollections) HasDmrBridges() bool`

HasDmrBridges returns a boolean if a field has been set.

### GetJndiConnectionFactories

`func (o *MsgVpnCollections) GetJndiConnectionFactories() MsgVpnCollectionsJndiconnectionfactories`

GetJndiConnectionFactories returns the JndiConnectionFactories field if non-nil, zero value otherwise.

### GetJndiConnectionFactoriesOk

`func (o *MsgVpnCollections) GetJndiConnectionFactoriesOk() (*MsgVpnCollectionsJndiconnectionfactories, bool)`

GetJndiConnectionFactoriesOk returns a tuple with the JndiConnectionFactories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJndiConnectionFactories

`func (o *MsgVpnCollections) SetJndiConnectionFactories(v MsgVpnCollectionsJndiconnectionfactories)`

SetJndiConnectionFactories sets JndiConnectionFactories field to given value.

### HasJndiConnectionFactories

`func (o *MsgVpnCollections) HasJndiConnectionFactories() bool`

HasJndiConnectionFactories returns a boolean if a field has been set.

### GetJndiQueues

`func (o *MsgVpnCollections) GetJndiQueues() MsgVpnCollectionsJndiqueues`

GetJndiQueues returns the JndiQueues field if non-nil, zero value otherwise.

### GetJndiQueuesOk

`func (o *MsgVpnCollections) GetJndiQueuesOk() (*MsgVpnCollectionsJndiqueues, bool)`

GetJndiQueuesOk returns a tuple with the JndiQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJndiQueues

`func (o *MsgVpnCollections) SetJndiQueues(v MsgVpnCollectionsJndiqueues)`

SetJndiQueues sets JndiQueues field to given value.

### HasJndiQueues

`func (o *MsgVpnCollections) HasJndiQueues() bool`

HasJndiQueues returns a boolean if a field has been set.

### GetJndiTopics

`func (o *MsgVpnCollections) GetJndiTopics() MsgVpnCollectionsJnditopics`

GetJndiTopics returns the JndiTopics field if non-nil, zero value otherwise.

### GetJndiTopicsOk

`func (o *MsgVpnCollections) GetJndiTopicsOk() (*MsgVpnCollectionsJnditopics, bool)`

GetJndiTopicsOk returns a tuple with the JndiTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJndiTopics

`func (o *MsgVpnCollections) SetJndiTopics(v MsgVpnCollectionsJnditopics)`

SetJndiTopics sets JndiTopics field to given value.

### HasJndiTopics

`func (o *MsgVpnCollections) HasJndiTopics() bool`

HasJndiTopics returns a boolean if a field has been set.

### GetMqttRetainCaches

`func (o *MsgVpnCollections) GetMqttRetainCaches() MsgVpnCollectionsMqttretaincaches`

GetMqttRetainCaches returns the MqttRetainCaches field if non-nil, zero value otherwise.

### GetMqttRetainCachesOk

`func (o *MsgVpnCollections) GetMqttRetainCachesOk() (*MsgVpnCollectionsMqttretaincaches, bool)`

GetMqttRetainCachesOk returns a tuple with the MqttRetainCaches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttRetainCaches

`func (o *MsgVpnCollections) SetMqttRetainCaches(v MsgVpnCollectionsMqttretaincaches)`

SetMqttRetainCaches sets MqttRetainCaches field to given value.

### HasMqttRetainCaches

`func (o *MsgVpnCollections) HasMqttRetainCaches() bool`

HasMqttRetainCaches returns a boolean if a field has been set.

### GetMqttSessions

`func (o *MsgVpnCollections) GetMqttSessions() MsgVpnCollectionsMqttsessions`

GetMqttSessions returns the MqttSessions field if non-nil, zero value otherwise.

### GetMqttSessionsOk

`func (o *MsgVpnCollections) GetMqttSessionsOk() (*MsgVpnCollectionsMqttsessions, bool)`

GetMqttSessionsOk returns a tuple with the MqttSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttSessions

`func (o *MsgVpnCollections) SetMqttSessions(v MsgVpnCollectionsMqttsessions)`

SetMqttSessions sets MqttSessions field to given value.

### HasMqttSessions

`func (o *MsgVpnCollections) HasMqttSessions() bool`

HasMqttSessions returns a boolean if a field has been set.

### GetQueueTemplates

`func (o *MsgVpnCollections) GetQueueTemplates() MsgVpnCollectionsQueuetemplates`

GetQueueTemplates returns the QueueTemplates field if non-nil, zero value otherwise.

### GetQueueTemplatesOk

`func (o *MsgVpnCollections) GetQueueTemplatesOk() (*MsgVpnCollectionsQueuetemplates, bool)`

GetQueueTemplatesOk returns a tuple with the QueueTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueTemplates

`func (o *MsgVpnCollections) SetQueueTemplates(v MsgVpnCollectionsQueuetemplates)`

SetQueueTemplates sets QueueTemplates field to given value.

### HasQueueTemplates

`func (o *MsgVpnCollections) HasQueueTemplates() bool`

HasQueueTemplates returns a boolean if a field has been set.

### GetQueues

`func (o *MsgVpnCollections) GetQueues() MsgVpnCollectionsQueues`

GetQueues returns the Queues field if non-nil, zero value otherwise.

### GetQueuesOk

`func (o *MsgVpnCollections) GetQueuesOk() (*MsgVpnCollectionsQueues, bool)`

GetQueuesOk returns a tuple with the Queues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueues

`func (o *MsgVpnCollections) SetQueues(v MsgVpnCollectionsQueues)`

SetQueues sets Queues field to given value.

### HasQueues

`func (o *MsgVpnCollections) HasQueues() bool`

HasQueues returns a boolean if a field has been set.

### GetReplayLogs

`func (o *MsgVpnCollections) GetReplayLogs() MsgVpnCollectionsReplaylogs`

GetReplayLogs returns the ReplayLogs field if non-nil, zero value otherwise.

### GetReplayLogsOk

`func (o *MsgVpnCollections) GetReplayLogsOk() (*MsgVpnCollectionsReplaylogs, bool)`

GetReplayLogsOk returns a tuple with the ReplayLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayLogs

`func (o *MsgVpnCollections) SetReplayLogs(v MsgVpnCollectionsReplaylogs)`

SetReplayLogs sets ReplayLogs field to given value.

### HasReplayLogs

`func (o *MsgVpnCollections) HasReplayLogs() bool`

HasReplayLogs returns a boolean if a field has been set.

### GetReplicatedTopics

`func (o *MsgVpnCollections) GetReplicatedTopics() MsgVpnCollectionsReplicatedtopics`

GetReplicatedTopics returns the ReplicatedTopics field if non-nil, zero value otherwise.

### GetReplicatedTopicsOk

`func (o *MsgVpnCollections) GetReplicatedTopicsOk() (*MsgVpnCollectionsReplicatedtopics, bool)`

GetReplicatedTopicsOk returns a tuple with the ReplicatedTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicatedTopics

`func (o *MsgVpnCollections) SetReplicatedTopics(v MsgVpnCollectionsReplicatedtopics)`

SetReplicatedTopics sets ReplicatedTopics field to given value.

### HasReplicatedTopics

`func (o *MsgVpnCollections) HasReplicatedTopics() bool`

HasReplicatedTopics returns a boolean if a field has been set.

### GetRestDeliveryPoints

`func (o *MsgVpnCollections) GetRestDeliveryPoints() MsgVpnCollectionsRestdeliverypoints`

GetRestDeliveryPoints returns the RestDeliveryPoints field if non-nil, zero value otherwise.

### GetRestDeliveryPointsOk

`func (o *MsgVpnCollections) GetRestDeliveryPointsOk() (*MsgVpnCollectionsRestdeliverypoints, bool)`

GetRestDeliveryPointsOk returns a tuple with the RestDeliveryPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestDeliveryPoints

`func (o *MsgVpnCollections) SetRestDeliveryPoints(v MsgVpnCollectionsRestdeliverypoints)`

SetRestDeliveryPoints sets RestDeliveryPoints field to given value.

### HasRestDeliveryPoints

`func (o *MsgVpnCollections) HasRestDeliveryPoints() bool`

HasRestDeliveryPoints returns a boolean if a field has been set.

### GetTopicEndpointTemplates

`func (o *MsgVpnCollections) GetTopicEndpointTemplates() MsgVpnCollectionsTopicendpointtemplates`

GetTopicEndpointTemplates returns the TopicEndpointTemplates field if non-nil, zero value otherwise.

### GetTopicEndpointTemplatesOk

`func (o *MsgVpnCollections) GetTopicEndpointTemplatesOk() (*MsgVpnCollectionsTopicendpointtemplates, bool)`

GetTopicEndpointTemplatesOk returns a tuple with the TopicEndpointTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicEndpointTemplates

`func (o *MsgVpnCollections) SetTopicEndpointTemplates(v MsgVpnCollectionsTopicendpointtemplates)`

SetTopicEndpointTemplates sets TopicEndpointTemplates field to given value.

### HasTopicEndpointTemplates

`func (o *MsgVpnCollections) HasTopicEndpointTemplates() bool`

HasTopicEndpointTemplates returns a boolean if a field has been set.

### GetTopicEndpoints

`func (o *MsgVpnCollections) GetTopicEndpoints() MsgVpnCollectionsTopicendpoints`

GetTopicEndpoints returns the TopicEndpoints field if non-nil, zero value otherwise.

### GetTopicEndpointsOk

`func (o *MsgVpnCollections) GetTopicEndpointsOk() (*MsgVpnCollectionsTopicendpoints, bool)`

GetTopicEndpointsOk returns a tuple with the TopicEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicEndpoints

`func (o *MsgVpnCollections) SetTopicEndpoints(v MsgVpnCollectionsTopicendpoints)`

SetTopicEndpoints sets TopicEndpoints field to given value.

### HasTopicEndpoints

`func (o *MsgVpnCollections) HasTopicEndpoints() bool`

HasTopicEndpoints returns a boolean if a field has been set.

### GetTransactions

`func (o *MsgVpnCollections) GetTransactions() MsgVpnCollectionsTransactions`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *MsgVpnCollections) GetTransactionsOk() (*MsgVpnCollectionsTransactions, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *MsgVpnCollections) SetTransactions(v MsgVpnCollectionsTransactions)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *MsgVpnCollections) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


