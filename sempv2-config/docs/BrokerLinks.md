# BrokerLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AboutUri** | Pointer to **string** | The URI of this Broker&#39;s About object. | [optional] 
**CertAuthoritiesUri** | Pointer to **string** | The URI of this Broker&#39;s collection of Certificate Authority objects. Deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities. | [optional] 
**ClientCertAuthoritiesUri** | Pointer to **string** | The URI of this Broker&#39;s collection of Client Certificate Authority objects. Available since 2.19. | [optional] 
**DmrClustersUri** | Pointer to **string** | The URI of this Broker&#39;s collection of Cluster objects. Available since 2.11. | [optional] 
**DomainCertAuthoritiesUri** | Pointer to **string** | The URI of this Broker&#39;s collection of Domain Certificate Authority objects. Available since 2.19. | [optional] 
**MsgVpnsUri** | Pointer to **string** | The URI of this Broker&#39;s collection of Message VPN objects. Available since 2.0. | [optional] 
**SystemInformationUri** | Pointer to **string** | The URI of this Broker&#39;s System Information object. Deprecated since 2.2. /systemInformation was replaced by /about/api. | [optional] 
**Uri** | Pointer to **string** | The URI of this Broker object. | [optional] 
**VirtualHostnamesUri** | Pointer to **string** | The URI of this Broker&#39;s collection of Virtual Hostname objects. Available since 2.17. | [optional] 

## Methods

### NewBrokerLinks

`func NewBrokerLinks() *BrokerLinks`

NewBrokerLinks instantiates a new BrokerLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerLinksWithDefaults

`func NewBrokerLinksWithDefaults() *BrokerLinks`

NewBrokerLinksWithDefaults instantiates a new BrokerLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAboutUri

`func (o *BrokerLinks) GetAboutUri() string`

GetAboutUri returns the AboutUri field if non-nil, zero value otherwise.

### GetAboutUriOk

`func (o *BrokerLinks) GetAboutUriOk() (*string, bool)`

GetAboutUriOk returns a tuple with the AboutUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAboutUri

`func (o *BrokerLinks) SetAboutUri(v string)`

SetAboutUri sets AboutUri field to given value.

### HasAboutUri

`func (o *BrokerLinks) HasAboutUri() bool`

HasAboutUri returns a boolean if a field has been set.

### GetCertAuthoritiesUri

`func (o *BrokerLinks) GetCertAuthoritiesUri() string`

GetCertAuthoritiesUri returns the CertAuthoritiesUri field if non-nil, zero value otherwise.

### GetCertAuthoritiesUriOk

`func (o *BrokerLinks) GetCertAuthoritiesUriOk() (*string, bool)`

GetCertAuthoritiesUriOk returns a tuple with the CertAuthoritiesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAuthoritiesUri

`func (o *BrokerLinks) SetCertAuthoritiesUri(v string)`

SetCertAuthoritiesUri sets CertAuthoritiesUri field to given value.

### HasCertAuthoritiesUri

`func (o *BrokerLinks) HasCertAuthoritiesUri() bool`

HasCertAuthoritiesUri returns a boolean if a field has been set.

### GetClientCertAuthoritiesUri

`func (o *BrokerLinks) GetClientCertAuthoritiesUri() string`

GetClientCertAuthoritiesUri returns the ClientCertAuthoritiesUri field if non-nil, zero value otherwise.

### GetClientCertAuthoritiesUriOk

`func (o *BrokerLinks) GetClientCertAuthoritiesUriOk() (*string, bool)`

GetClientCertAuthoritiesUriOk returns a tuple with the ClientCertAuthoritiesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertAuthoritiesUri

`func (o *BrokerLinks) SetClientCertAuthoritiesUri(v string)`

SetClientCertAuthoritiesUri sets ClientCertAuthoritiesUri field to given value.

### HasClientCertAuthoritiesUri

`func (o *BrokerLinks) HasClientCertAuthoritiesUri() bool`

HasClientCertAuthoritiesUri returns a boolean if a field has been set.

### GetDmrClustersUri

`func (o *BrokerLinks) GetDmrClustersUri() string`

GetDmrClustersUri returns the DmrClustersUri field if non-nil, zero value otherwise.

### GetDmrClustersUriOk

`func (o *BrokerLinks) GetDmrClustersUriOk() (*string, bool)`

GetDmrClustersUriOk returns a tuple with the DmrClustersUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmrClustersUri

`func (o *BrokerLinks) SetDmrClustersUri(v string)`

SetDmrClustersUri sets DmrClustersUri field to given value.

### HasDmrClustersUri

`func (o *BrokerLinks) HasDmrClustersUri() bool`

HasDmrClustersUri returns a boolean if a field has been set.

### GetDomainCertAuthoritiesUri

`func (o *BrokerLinks) GetDomainCertAuthoritiesUri() string`

GetDomainCertAuthoritiesUri returns the DomainCertAuthoritiesUri field if non-nil, zero value otherwise.

### GetDomainCertAuthoritiesUriOk

`func (o *BrokerLinks) GetDomainCertAuthoritiesUriOk() (*string, bool)`

GetDomainCertAuthoritiesUriOk returns a tuple with the DomainCertAuthoritiesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainCertAuthoritiesUri

`func (o *BrokerLinks) SetDomainCertAuthoritiesUri(v string)`

SetDomainCertAuthoritiesUri sets DomainCertAuthoritiesUri field to given value.

### HasDomainCertAuthoritiesUri

`func (o *BrokerLinks) HasDomainCertAuthoritiesUri() bool`

HasDomainCertAuthoritiesUri returns a boolean if a field has been set.

### GetMsgVpnsUri

`func (o *BrokerLinks) GetMsgVpnsUri() string`

GetMsgVpnsUri returns the MsgVpnsUri field if non-nil, zero value otherwise.

### GetMsgVpnsUriOk

`func (o *BrokerLinks) GetMsgVpnsUriOk() (*string, bool)`

GetMsgVpnsUriOk returns a tuple with the MsgVpnsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnsUri

`func (o *BrokerLinks) SetMsgVpnsUri(v string)`

SetMsgVpnsUri sets MsgVpnsUri field to given value.

### HasMsgVpnsUri

`func (o *BrokerLinks) HasMsgVpnsUri() bool`

HasMsgVpnsUri returns a boolean if a field has been set.

### GetSystemInformationUri

`func (o *BrokerLinks) GetSystemInformationUri() string`

GetSystemInformationUri returns the SystemInformationUri field if non-nil, zero value otherwise.

### GetSystemInformationUriOk

`func (o *BrokerLinks) GetSystemInformationUriOk() (*string, bool)`

GetSystemInformationUriOk returns a tuple with the SystemInformationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInformationUri

`func (o *BrokerLinks) SetSystemInformationUri(v string)`

SetSystemInformationUri sets SystemInformationUri field to given value.

### HasSystemInformationUri

`func (o *BrokerLinks) HasSystemInformationUri() bool`

HasSystemInformationUri returns a boolean if a field has been set.

### GetUri

`func (o *BrokerLinks) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BrokerLinks) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BrokerLinks) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *BrokerLinks) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetVirtualHostnamesUri

`func (o *BrokerLinks) GetVirtualHostnamesUri() string`

GetVirtualHostnamesUri returns the VirtualHostnamesUri field if non-nil, zero value otherwise.

### GetVirtualHostnamesUriOk

`func (o *BrokerLinks) GetVirtualHostnamesUriOk() (*string, bool)`

GetVirtualHostnamesUriOk returns a tuple with the VirtualHostnamesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHostnamesUri

`func (o *BrokerLinks) SetVirtualHostnamesUri(v string)`

SetVirtualHostnamesUri sets VirtualHostnamesUri field to given value.

### HasVirtualHostnamesUri

`func (o *BrokerLinks) HasVirtualHostnamesUri() bool`

HasVirtualHostnamesUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


