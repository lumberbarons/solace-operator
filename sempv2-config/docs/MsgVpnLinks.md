# MsgVpnLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfilesUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of ACL Profile objects. | [optional] 
**AuthenticationOauthProvidersUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of OAuth Provider objects. Available since 2.13. | [optional] 
**AuthorizationGroupsUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of LDAP Authorization Group objects. | [optional] 
**BridgesUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of Bridge objects. | [optional] 
**ClientProfilesUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of Client Profile objects. | [optional] 
**ClientUsernamesUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of Client Username objects. | [optional] 
**DistributedCachesUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of Distributed Cache objects. Available since 2.11. | [optional] 
**DmrBridgesUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of DMR Bridge objects. Available since 2.11. | [optional] 
**JndiConnectionFactoriesUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of JNDI Connection Factory objects. Available since 2.2. | [optional] 
**JndiQueuesUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of JNDI Queue objects. Available since 2.2. | [optional] 
**JndiTopicsUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of JNDI Topic objects. Available since 2.2. | [optional] 
**MqttRetainCachesUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of MQTT Retain Cache objects. Available since 2.11. | [optional] 
**MqttSessionsUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of MQTT Session objects. Available since 2.1. | [optional] 
**QueueTemplatesUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of Queue Template objects. Available since 2.14. | [optional] 
**QueuesUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of Queue objects. | [optional] 
**ReplayLogsUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of Replay Log objects. Available since 2.10. | [optional] 
**ReplicatedTopicsUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of Replicated Topic objects. Available since 2.1. | [optional] 
**RestDeliveryPointsUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of REST Delivery Point objects. | [optional] 
**SequencedTopicsUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of Sequenced Topic objects. | [optional] 
**TopicEndpointTemplatesUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of Topic Endpoint Template objects. Available since 2.14. | [optional] 
**TopicEndpointsUri** | Pointer to **string** | The URI of this Message VPN&#39;s collection of Topic Endpoint objects. Available since 2.1. | [optional] 
**Uri** | Pointer to **string** | The URI of this Message VPN object. | [optional] 

## Methods

### NewMsgVpnLinks

`func NewMsgVpnLinks() *MsgVpnLinks`

NewMsgVpnLinks instantiates a new MsgVpnLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnLinksWithDefaults

`func NewMsgVpnLinksWithDefaults() *MsgVpnLinks`

NewMsgVpnLinksWithDefaults instantiates a new MsgVpnLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclProfilesUri

`func (o *MsgVpnLinks) GetAclProfilesUri() string`

GetAclProfilesUri returns the AclProfilesUri field if non-nil, zero value otherwise.

### GetAclProfilesUriOk

`func (o *MsgVpnLinks) GetAclProfilesUriOk() (*string, bool)`

GetAclProfilesUriOk returns a tuple with the AclProfilesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclProfilesUri

`func (o *MsgVpnLinks) SetAclProfilesUri(v string)`

SetAclProfilesUri sets AclProfilesUri field to given value.

### HasAclProfilesUri

`func (o *MsgVpnLinks) HasAclProfilesUri() bool`

HasAclProfilesUri returns a boolean if a field has been set.

### GetAuthenticationOauthProvidersUri

`func (o *MsgVpnLinks) GetAuthenticationOauthProvidersUri() string`

GetAuthenticationOauthProvidersUri returns the AuthenticationOauthProvidersUri field if non-nil, zero value otherwise.

### GetAuthenticationOauthProvidersUriOk

`func (o *MsgVpnLinks) GetAuthenticationOauthProvidersUriOk() (*string, bool)`

GetAuthenticationOauthProvidersUriOk returns a tuple with the AuthenticationOauthProvidersUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthProvidersUri

`func (o *MsgVpnLinks) SetAuthenticationOauthProvidersUri(v string)`

SetAuthenticationOauthProvidersUri sets AuthenticationOauthProvidersUri field to given value.

### HasAuthenticationOauthProvidersUri

`func (o *MsgVpnLinks) HasAuthenticationOauthProvidersUri() bool`

HasAuthenticationOauthProvidersUri returns a boolean if a field has been set.

### GetAuthorizationGroupsUri

`func (o *MsgVpnLinks) GetAuthorizationGroupsUri() string`

GetAuthorizationGroupsUri returns the AuthorizationGroupsUri field if non-nil, zero value otherwise.

### GetAuthorizationGroupsUriOk

`func (o *MsgVpnLinks) GetAuthorizationGroupsUriOk() (*string, bool)`

GetAuthorizationGroupsUriOk returns a tuple with the AuthorizationGroupsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationGroupsUri

`func (o *MsgVpnLinks) SetAuthorizationGroupsUri(v string)`

SetAuthorizationGroupsUri sets AuthorizationGroupsUri field to given value.

### HasAuthorizationGroupsUri

`func (o *MsgVpnLinks) HasAuthorizationGroupsUri() bool`

HasAuthorizationGroupsUri returns a boolean if a field has been set.

### GetBridgesUri

`func (o *MsgVpnLinks) GetBridgesUri() string`

GetBridgesUri returns the BridgesUri field if non-nil, zero value otherwise.

### GetBridgesUriOk

`func (o *MsgVpnLinks) GetBridgesUriOk() (*string, bool)`

GetBridgesUriOk returns a tuple with the BridgesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgesUri

`func (o *MsgVpnLinks) SetBridgesUri(v string)`

SetBridgesUri sets BridgesUri field to given value.

### HasBridgesUri

`func (o *MsgVpnLinks) HasBridgesUri() bool`

HasBridgesUri returns a boolean if a field has been set.

### GetClientProfilesUri

`func (o *MsgVpnLinks) GetClientProfilesUri() string`

GetClientProfilesUri returns the ClientProfilesUri field if non-nil, zero value otherwise.

### GetClientProfilesUriOk

`func (o *MsgVpnLinks) GetClientProfilesUriOk() (*string, bool)`

GetClientProfilesUriOk returns a tuple with the ClientProfilesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfilesUri

`func (o *MsgVpnLinks) SetClientProfilesUri(v string)`

SetClientProfilesUri sets ClientProfilesUri field to given value.

### HasClientProfilesUri

`func (o *MsgVpnLinks) HasClientProfilesUri() bool`

HasClientProfilesUri returns a boolean if a field has been set.

### GetClientUsernamesUri

`func (o *MsgVpnLinks) GetClientUsernamesUri() string`

GetClientUsernamesUri returns the ClientUsernamesUri field if non-nil, zero value otherwise.

### GetClientUsernamesUriOk

`func (o *MsgVpnLinks) GetClientUsernamesUriOk() (*string, bool)`

GetClientUsernamesUriOk returns a tuple with the ClientUsernamesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUsernamesUri

`func (o *MsgVpnLinks) SetClientUsernamesUri(v string)`

SetClientUsernamesUri sets ClientUsernamesUri field to given value.

### HasClientUsernamesUri

`func (o *MsgVpnLinks) HasClientUsernamesUri() bool`

HasClientUsernamesUri returns a boolean if a field has been set.

### GetDistributedCachesUri

`func (o *MsgVpnLinks) GetDistributedCachesUri() string`

GetDistributedCachesUri returns the DistributedCachesUri field if non-nil, zero value otherwise.

### GetDistributedCachesUriOk

`func (o *MsgVpnLinks) GetDistributedCachesUriOk() (*string, bool)`

GetDistributedCachesUriOk returns a tuple with the DistributedCachesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedCachesUri

`func (o *MsgVpnLinks) SetDistributedCachesUri(v string)`

SetDistributedCachesUri sets DistributedCachesUri field to given value.

### HasDistributedCachesUri

`func (o *MsgVpnLinks) HasDistributedCachesUri() bool`

HasDistributedCachesUri returns a boolean if a field has been set.

### GetDmrBridgesUri

`func (o *MsgVpnLinks) GetDmrBridgesUri() string`

GetDmrBridgesUri returns the DmrBridgesUri field if non-nil, zero value otherwise.

### GetDmrBridgesUriOk

`func (o *MsgVpnLinks) GetDmrBridgesUriOk() (*string, bool)`

GetDmrBridgesUriOk returns a tuple with the DmrBridgesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmrBridgesUri

`func (o *MsgVpnLinks) SetDmrBridgesUri(v string)`

SetDmrBridgesUri sets DmrBridgesUri field to given value.

### HasDmrBridgesUri

`func (o *MsgVpnLinks) HasDmrBridgesUri() bool`

HasDmrBridgesUri returns a boolean if a field has been set.

### GetJndiConnectionFactoriesUri

`func (o *MsgVpnLinks) GetJndiConnectionFactoriesUri() string`

GetJndiConnectionFactoriesUri returns the JndiConnectionFactoriesUri field if non-nil, zero value otherwise.

### GetJndiConnectionFactoriesUriOk

`func (o *MsgVpnLinks) GetJndiConnectionFactoriesUriOk() (*string, bool)`

GetJndiConnectionFactoriesUriOk returns a tuple with the JndiConnectionFactoriesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJndiConnectionFactoriesUri

`func (o *MsgVpnLinks) SetJndiConnectionFactoriesUri(v string)`

SetJndiConnectionFactoriesUri sets JndiConnectionFactoriesUri field to given value.

### HasJndiConnectionFactoriesUri

`func (o *MsgVpnLinks) HasJndiConnectionFactoriesUri() bool`

HasJndiConnectionFactoriesUri returns a boolean if a field has been set.

### GetJndiQueuesUri

`func (o *MsgVpnLinks) GetJndiQueuesUri() string`

GetJndiQueuesUri returns the JndiQueuesUri field if non-nil, zero value otherwise.

### GetJndiQueuesUriOk

`func (o *MsgVpnLinks) GetJndiQueuesUriOk() (*string, bool)`

GetJndiQueuesUriOk returns a tuple with the JndiQueuesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJndiQueuesUri

`func (o *MsgVpnLinks) SetJndiQueuesUri(v string)`

SetJndiQueuesUri sets JndiQueuesUri field to given value.

### HasJndiQueuesUri

`func (o *MsgVpnLinks) HasJndiQueuesUri() bool`

HasJndiQueuesUri returns a boolean if a field has been set.

### GetJndiTopicsUri

`func (o *MsgVpnLinks) GetJndiTopicsUri() string`

GetJndiTopicsUri returns the JndiTopicsUri field if non-nil, zero value otherwise.

### GetJndiTopicsUriOk

`func (o *MsgVpnLinks) GetJndiTopicsUriOk() (*string, bool)`

GetJndiTopicsUriOk returns a tuple with the JndiTopicsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJndiTopicsUri

`func (o *MsgVpnLinks) SetJndiTopicsUri(v string)`

SetJndiTopicsUri sets JndiTopicsUri field to given value.

### HasJndiTopicsUri

`func (o *MsgVpnLinks) HasJndiTopicsUri() bool`

HasJndiTopicsUri returns a boolean if a field has been set.

### GetMqttRetainCachesUri

`func (o *MsgVpnLinks) GetMqttRetainCachesUri() string`

GetMqttRetainCachesUri returns the MqttRetainCachesUri field if non-nil, zero value otherwise.

### GetMqttRetainCachesUriOk

`func (o *MsgVpnLinks) GetMqttRetainCachesUriOk() (*string, bool)`

GetMqttRetainCachesUriOk returns a tuple with the MqttRetainCachesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttRetainCachesUri

`func (o *MsgVpnLinks) SetMqttRetainCachesUri(v string)`

SetMqttRetainCachesUri sets MqttRetainCachesUri field to given value.

### HasMqttRetainCachesUri

`func (o *MsgVpnLinks) HasMqttRetainCachesUri() bool`

HasMqttRetainCachesUri returns a boolean if a field has been set.

### GetMqttSessionsUri

`func (o *MsgVpnLinks) GetMqttSessionsUri() string`

GetMqttSessionsUri returns the MqttSessionsUri field if non-nil, zero value otherwise.

### GetMqttSessionsUriOk

`func (o *MsgVpnLinks) GetMqttSessionsUriOk() (*string, bool)`

GetMqttSessionsUriOk returns a tuple with the MqttSessionsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttSessionsUri

`func (o *MsgVpnLinks) SetMqttSessionsUri(v string)`

SetMqttSessionsUri sets MqttSessionsUri field to given value.

### HasMqttSessionsUri

`func (o *MsgVpnLinks) HasMqttSessionsUri() bool`

HasMqttSessionsUri returns a boolean if a field has been set.

### GetQueueTemplatesUri

`func (o *MsgVpnLinks) GetQueueTemplatesUri() string`

GetQueueTemplatesUri returns the QueueTemplatesUri field if non-nil, zero value otherwise.

### GetQueueTemplatesUriOk

`func (o *MsgVpnLinks) GetQueueTemplatesUriOk() (*string, bool)`

GetQueueTemplatesUriOk returns a tuple with the QueueTemplatesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueTemplatesUri

`func (o *MsgVpnLinks) SetQueueTemplatesUri(v string)`

SetQueueTemplatesUri sets QueueTemplatesUri field to given value.

### HasQueueTemplatesUri

`func (o *MsgVpnLinks) HasQueueTemplatesUri() bool`

HasQueueTemplatesUri returns a boolean if a field has been set.

### GetQueuesUri

`func (o *MsgVpnLinks) GetQueuesUri() string`

GetQueuesUri returns the QueuesUri field if non-nil, zero value otherwise.

### GetQueuesUriOk

`func (o *MsgVpnLinks) GetQueuesUriOk() (*string, bool)`

GetQueuesUriOk returns a tuple with the QueuesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuesUri

`func (o *MsgVpnLinks) SetQueuesUri(v string)`

SetQueuesUri sets QueuesUri field to given value.

### HasQueuesUri

`func (o *MsgVpnLinks) HasQueuesUri() bool`

HasQueuesUri returns a boolean if a field has been set.

### GetReplayLogsUri

`func (o *MsgVpnLinks) GetReplayLogsUri() string`

GetReplayLogsUri returns the ReplayLogsUri field if non-nil, zero value otherwise.

### GetReplayLogsUriOk

`func (o *MsgVpnLinks) GetReplayLogsUriOk() (*string, bool)`

GetReplayLogsUriOk returns a tuple with the ReplayLogsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayLogsUri

`func (o *MsgVpnLinks) SetReplayLogsUri(v string)`

SetReplayLogsUri sets ReplayLogsUri field to given value.

### HasReplayLogsUri

`func (o *MsgVpnLinks) HasReplayLogsUri() bool`

HasReplayLogsUri returns a boolean if a field has been set.

### GetReplicatedTopicsUri

`func (o *MsgVpnLinks) GetReplicatedTopicsUri() string`

GetReplicatedTopicsUri returns the ReplicatedTopicsUri field if non-nil, zero value otherwise.

### GetReplicatedTopicsUriOk

`func (o *MsgVpnLinks) GetReplicatedTopicsUriOk() (*string, bool)`

GetReplicatedTopicsUriOk returns a tuple with the ReplicatedTopicsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicatedTopicsUri

`func (o *MsgVpnLinks) SetReplicatedTopicsUri(v string)`

SetReplicatedTopicsUri sets ReplicatedTopicsUri field to given value.

### HasReplicatedTopicsUri

`func (o *MsgVpnLinks) HasReplicatedTopicsUri() bool`

HasReplicatedTopicsUri returns a boolean if a field has been set.

### GetRestDeliveryPointsUri

`func (o *MsgVpnLinks) GetRestDeliveryPointsUri() string`

GetRestDeliveryPointsUri returns the RestDeliveryPointsUri field if non-nil, zero value otherwise.

### GetRestDeliveryPointsUriOk

`func (o *MsgVpnLinks) GetRestDeliveryPointsUriOk() (*string, bool)`

GetRestDeliveryPointsUriOk returns a tuple with the RestDeliveryPointsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestDeliveryPointsUri

`func (o *MsgVpnLinks) SetRestDeliveryPointsUri(v string)`

SetRestDeliveryPointsUri sets RestDeliveryPointsUri field to given value.

### HasRestDeliveryPointsUri

`func (o *MsgVpnLinks) HasRestDeliveryPointsUri() bool`

HasRestDeliveryPointsUri returns a boolean if a field has been set.

### GetSequencedTopicsUri

`func (o *MsgVpnLinks) GetSequencedTopicsUri() string`

GetSequencedTopicsUri returns the SequencedTopicsUri field if non-nil, zero value otherwise.

### GetSequencedTopicsUriOk

`func (o *MsgVpnLinks) GetSequencedTopicsUriOk() (*string, bool)`

GetSequencedTopicsUriOk returns a tuple with the SequencedTopicsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequencedTopicsUri

`func (o *MsgVpnLinks) SetSequencedTopicsUri(v string)`

SetSequencedTopicsUri sets SequencedTopicsUri field to given value.

### HasSequencedTopicsUri

`func (o *MsgVpnLinks) HasSequencedTopicsUri() bool`

HasSequencedTopicsUri returns a boolean if a field has been set.

### GetTopicEndpointTemplatesUri

`func (o *MsgVpnLinks) GetTopicEndpointTemplatesUri() string`

GetTopicEndpointTemplatesUri returns the TopicEndpointTemplatesUri field if non-nil, zero value otherwise.

### GetTopicEndpointTemplatesUriOk

`func (o *MsgVpnLinks) GetTopicEndpointTemplatesUriOk() (*string, bool)`

GetTopicEndpointTemplatesUriOk returns a tuple with the TopicEndpointTemplatesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicEndpointTemplatesUri

`func (o *MsgVpnLinks) SetTopicEndpointTemplatesUri(v string)`

SetTopicEndpointTemplatesUri sets TopicEndpointTemplatesUri field to given value.

### HasTopicEndpointTemplatesUri

`func (o *MsgVpnLinks) HasTopicEndpointTemplatesUri() bool`

HasTopicEndpointTemplatesUri returns a boolean if a field has been set.

### GetTopicEndpointsUri

`func (o *MsgVpnLinks) GetTopicEndpointsUri() string`

GetTopicEndpointsUri returns the TopicEndpointsUri field if non-nil, zero value otherwise.

### GetTopicEndpointsUriOk

`func (o *MsgVpnLinks) GetTopicEndpointsUriOk() (*string, bool)`

GetTopicEndpointsUriOk returns a tuple with the TopicEndpointsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicEndpointsUri

`func (o *MsgVpnLinks) SetTopicEndpointsUri(v string)`

SetTopicEndpointsUri sets TopicEndpointsUri field to given value.

### HasTopicEndpointsUri

`func (o *MsgVpnLinks) HasTopicEndpointsUri() bool`

HasTopicEndpointsUri returns a boolean if a field has been set.

### GetUri

`func (o *MsgVpnLinks) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *MsgVpnLinks) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *MsgVpnLinks) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *MsgVpnLinks) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


