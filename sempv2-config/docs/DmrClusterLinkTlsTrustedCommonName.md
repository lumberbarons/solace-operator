# DmrClusterLinkTlsTrustedCommonName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DmrClusterName** | Pointer to **string** | The name of the Cluster. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation. | [optional] 
**RemoteNodeName** | Pointer to **string** | The name of the node at the remote end of the Link. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation. | [optional] 
**TlsTrustedCommonName** | Pointer to **string** | The expected trusted common name of the remote certificate. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation. | [optional] 

## Methods

### NewDmrClusterLinkTlsTrustedCommonName

`func NewDmrClusterLinkTlsTrustedCommonName() *DmrClusterLinkTlsTrustedCommonName`

NewDmrClusterLinkTlsTrustedCommonName instantiates a new DmrClusterLinkTlsTrustedCommonName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDmrClusterLinkTlsTrustedCommonNameWithDefaults

`func NewDmrClusterLinkTlsTrustedCommonNameWithDefaults() *DmrClusterLinkTlsTrustedCommonName`

NewDmrClusterLinkTlsTrustedCommonNameWithDefaults instantiates a new DmrClusterLinkTlsTrustedCommonName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDmrClusterName

`func (o *DmrClusterLinkTlsTrustedCommonName) GetDmrClusterName() string`

GetDmrClusterName returns the DmrClusterName field if non-nil, zero value otherwise.

### GetDmrClusterNameOk

`func (o *DmrClusterLinkTlsTrustedCommonName) GetDmrClusterNameOk() (*string, bool)`

GetDmrClusterNameOk returns a tuple with the DmrClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmrClusterName

`func (o *DmrClusterLinkTlsTrustedCommonName) SetDmrClusterName(v string)`

SetDmrClusterName sets DmrClusterName field to given value.

### HasDmrClusterName

`func (o *DmrClusterLinkTlsTrustedCommonName) HasDmrClusterName() bool`

HasDmrClusterName returns a boolean if a field has been set.

### GetRemoteNodeName

`func (o *DmrClusterLinkTlsTrustedCommonName) GetRemoteNodeName() string`

GetRemoteNodeName returns the RemoteNodeName field if non-nil, zero value otherwise.

### GetRemoteNodeNameOk

`func (o *DmrClusterLinkTlsTrustedCommonName) GetRemoteNodeNameOk() (*string, bool)`

GetRemoteNodeNameOk returns a tuple with the RemoteNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNodeName

`func (o *DmrClusterLinkTlsTrustedCommonName) SetRemoteNodeName(v string)`

SetRemoteNodeName sets RemoteNodeName field to given value.

### HasRemoteNodeName

`func (o *DmrClusterLinkTlsTrustedCommonName) HasRemoteNodeName() bool`

HasRemoteNodeName returns a boolean if a field has been set.

### GetTlsTrustedCommonName

`func (o *DmrClusterLinkTlsTrustedCommonName) GetTlsTrustedCommonName() string`

GetTlsTrustedCommonName returns the TlsTrustedCommonName field if non-nil, zero value otherwise.

### GetTlsTrustedCommonNameOk

`func (o *DmrClusterLinkTlsTrustedCommonName) GetTlsTrustedCommonNameOk() (*string, bool)`

GetTlsTrustedCommonNameOk returns a tuple with the TlsTrustedCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTrustedCommonName

`func (o *DmrClusterLinkTlsTrustedCommonName) SetTlsTrustedCommonName(v string)`

SetTlsTrustedCommonName sets TlsTrustedCommonName field to given value.

### HasTlsTrustedCommonName

`func (o *DmrClusterLinkTlsTrustedCommonName) HasTlsTrustedCommonName() bool`

HasTlsTrustedCommonName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


