# BrokerCollections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertAuthorities** | Pointer to [**BrokerCollectionsCertauthorities**](BrokerCollectionsCertauthorities.md) |  | [optional] 
**ClientCertAuthorities** | Pointer to [**BrokerCollectionsClientcertauthorities**](BrokerCollectionsClientcertauthorities.md) |  | [optional] 
**DmrClusters** | Pointer to [**BrokerCollectionsDmrclusters**](BrokerCollectionsDmrclusters.md) |  | [optional] 
**DomainCertAuthorities** | Pointer to [**BrokerCollectionsDomaincertauthorities**](BrokerCollectionsDomaincertauthorities.md) |  | [optional] 
**MsgVpns** | Pointer to [**BrokerCollectionsMsgvpns**](BrokerCollectionsMsgvpns.md) |  | [optional] 
**Sessions** | Pointer to [**BrokerCollectionsSessions**](BrokerCollectionsSessions.md) |  | [optional] 
**StandardDomainCertAuthorities** | Pointer to [**BrokerCollectionsStandarddomaincertauthorities**](BrokerCollectionsStandarddomaincertauthorities.md) |  | [optional] 
**VirtualHostnames** | Pointer to [**BrokerCollectionsVirtualhostnames**](BrokerCollectionsVirtualhostnames.md) |  | [optional] 

## Methods

### NewBrokerCollections

`func NewBrokerCollections() *BrokerCollections`

NewBrokerCollections instantiates a new BrokerCollections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerCollectionsWithDefaults

`func NewBrokerCollectionsWithDefaults() *BrokerCollections`

NewBrokerCollectionsWithDefaults instantiates a new BrokerCollections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertAuthorities

`func (o *BrokerCollections) GetCertAuthorities() BrokerCollectionsCertauthorities`

GetCertAuthorities returns the CertAuthorities field if non-nil, zero value otherwise.

### GetCertAuthoritiesOk

`func (o *BrokerCollections) GetCertAuthoritiesOk() (*BrokerCollectionsCertauthorities, bool)`

GetCertAuthoritiesOk returns a tuple with the CertAuthorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAuthorities

`func (o *BrokerCollections) SetCertAuthorities(v BrokerCollectionsCertauthorities)`

SetCertAuthorities sets CertAuthorities field to given value.

### HasCertAuthorities

`func (o *BrokerCollections) HasCertAuthorities() bool`

HasCertAuthorities returns a boolean if a field has been set.

### GetClientCertAuthorities

`func (o *BrokerCollections) GetClientCertAuthorities() BrokerCollectionsClientcertauthorities`

GetClientCertAuthorities returns the ClientCertAuthorities field if non-nil, zero value otherwise.

### GetClientCertAuthoritiesOk

`func (o *BrokerCollections) GetClientCertAuthoritiesOk() (*BrokerCollectionsClientcertauthorities, bool)`

GetClientCertAuthoritiesOk returns a tuple with the ClientCertAuthorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertAuthorities

`func (o *BrokerCollections) SetClientCertAuthorities(v BrokerCollectionsClientcertauthorities)`

SetClientCertAuthorities sets ClientCertAuthorities field to given value.

### HasClientCertAuthorities

`func (o *BrokerCollections) HasClientCertAuthorities() bool`

HasClientCertAuthorities returns a boolean if a field has been set.

### GetDmrClusters

`func (o *BrokerCollections) GetDmrClusters() BrokerCollectionsDmrclusters`

GetDmrClusters returns the DmrClusters field if non-nil, zero value otherwise.

### GetDmrClustersOk

`func (o *BrokerCollections) GetDmrClustersOk() (*BrokerCollectionsDmrclusters, bool)`

GetDmrClustersOk returns a tuple with the DmrClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmrClusters

`func (o *BrokerCollections) SetDmrClusters(v BrokerCollectionsDmrclusters)`

SetDmrClusters sets DmrClusters field to given value.

### HasDmrClusters

`func (o *BrokerCollections) HasDmrClusters() bool`

HasDmrClusters returns a boolean if a field has been set.

### GetDomainCertAuthorities

`func (o *BrokerCollections) GetDomainCertAuthorities() BrokerCollectionsDomaincertauthorities`

GetDomainCertAuthorities returns the DomainCertAuthorities field if non-nil, zero value otherwise.

### GetDomainCertAuthoritiesOk

`func (o *BrokerCollections) GetDomainCertAuthoritiesOk() (*BrokerCollectionsDomaincertauthorities, bool)`

GetDomainCertAuthoritiesOk returns a tuple with the DomainCertAuthorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainCertAuthorities

`func (o *BrokerCollections) SetDomainCertAuthorities(v BrokerCollectionsDomaincertauthorities)`

SetDomainCertAuthorities sets DomainCertAuthorities field to given value.

### HasDomainCertAuthorities

`func (o *BrokerCollections) HasDomainCertAuthorities() bool`

HasDomainCertAuthorities returns a boolean if a field has been set.

### GetMsgVpns

`func (o *BrokerCollections) GetMsgVpns() BrokerCollectionsMsgvpns`

GetMsgVpns returns the MsgVpns field if non-nil, zero value otherwise.

### GetMsgVpnsOk

`func (o *BrokerCollections) GetMsgVpnsOk() (*BrokerCollectionsMsgvpns, bool)`

GetMsgVpnsOk returns a tuple with the MsgVpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpns

`func (o *BrokerCollections) SetMsgVpns(v BrokerCollectionsMsgvpns)`

SetMsgVpns sets MsgVpns field to given value.

### HasMsgVpns

`func (o *BrokerCollections) HasMsgVpns() bool`

HasMsgVpns returns a boolean if a field has been set.

### GetSessions

`func (o *BrokerCollections) GetSessions() BrokerCollectionsSessions`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *BrokerCollections) GetSessionsOk() (*BrokerCollectionsSessions, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *BrokerCollections) SetSessions(v BrokerCollectionsSessions)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *BrokerCollections) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetStandardDomainCertAuthorities

`func (o *BrokerCollections) GetStandardDomainCertAuthorities() BrokerCollectionsStandarddomaincertauthorities`

GetStandardDomainCertAuthorities returns the StandardDomainCertAuthorities field if non-nil, zero value otherwise.

### GetStandardDomainCertAuthoritiesOk

`func (o *BrokerCollections) GetStandardDomainCertAuthoritiesOk() (*BrokerCollectionsStandarddomaincertauthorities, bool)`

GetStandardDomainCertAuthoritiesOk returns a tuple with the StandardDomainCertAuthorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardDomainCertAuthorities

`func (o *BrokerCollections) SetStandardDomainCertAuthorities(v BrokerCollectionsStandarddomaincertauthorities)`

SetStandardDomainCertAuthorities sets StandardDomainCertAuthorities field to given value.

### HasStandardDomainCertAuthorities

`func (o *BrokerCollections) HasStandardDomainCertAuthorities() bool`

HasStandardDomainCertAuthorities returns a boolean if a field has been set.

### GetVirtualHostnames

`func (o *BrokerCollections) GetVirtualHostnames() BrokerCollectionsVirtualhostnames`

GetVirtualHostnames returns the VirtualHostnames field if non-nil, zero value otherwise.

### GetVirtualHostnamesOk

`func (o *BrokerCollections) GetVirtualHostnamesOk() (*BrokerCollectionsVirtualhostnames, bool)`

GetVirtualHostnamesOk returns a tuple with the VirtualHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHostnames

`func (o *BrokerCollections) SetVirtualHostnames(v BrokerCollectionsVirtualhostnames)`

SetVirtualHostnames sets VirtualHostnames field to given value.

### HasVirtualHostnames

`func (o *BrokerCollections) HasVirtualHostnames() bool`

HasVirtualHostnames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


