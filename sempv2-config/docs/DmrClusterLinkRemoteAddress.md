# DmrClusterLinkRemoteAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DmrClusterName** | Pointer to **string** | The name of the Cluster. | [optional] 
**RemoteAddress** | Pointer to **string** | The FQDN or IP address (and optional port) of the remote node. If a port is not provided, it will vary based on the transport encoding: 55555 (plain-text), 55443 (encrypted), or 55003 (compressed). | [optional] 
**RemoteNodeName** | Pointer to **string** | The name of the node at the remote end of the Link. | [optional] 

## Methods

### NewDmrClusterLinkRemoteAddress

`func NewDmrClusterLinkRemoteAddress() *DmrClusterLinkRemoteAddress`

NewDmrClusterLinkRemoteAddress instantiates a new DmrClusterLinkRemoteAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDmrClusterLinkRemoteAddressWithDefaults

`func NewDmrClusterLinkRemoteAddressWithDefaults() *DmrClusterLinkRemoteAddress`

NewDmrClusterLinkRemoteAddressWithDefaults instantiates a new DmrClusterLinkRemoteAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDmrClusterName

`func (o *DmrClusterLinkRemoteAddress) GetDmrClusterName() string`

GetDmrClusterName returns the DmrClusterName field if non-nil, zero value otherwise.

### GetDmrClusterNameOk

`func (o *DmrClusterLinkRemoteAddress) GetDmrClusterNameOk() (*string, bool)`

GetDmrClusterNameOk returns a tuple with the DmrClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmrClusterName

`func (o *DmrClusterLinkRemoteAddress) SetDmrClusterName(v string)`

SetDmrClusterName sets DmrClusterName field to given value.

### HasDmrClusterName

`func (o *DmrClusterLinkRemoteAddress) HasDmrClusterName() bool`

HasDmrClusterName returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *DmrClusterLinkRemoteAddress) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *DmrClusterLinkRemoteAddress) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *DmrClusterLinkRemoteAddress) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *DmrClusterLinkRemoteAddress) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetRemoteNodeName

`func (o *DmrClusterLinkRemoteAddress) GetRemoteNodeName() string`

GetRemoteNodeName returns the RemoteNodeName field if non-nil, zero value otherwise.

### GetRemoteNodeNameOk

`func (o *DmrClusterLinkRemoteAddress) GetRemoteNodeNameOk() (*string, bool)`

GetRemoteNodeNameOk returns a tuple with the RemoteNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNodeName

`func (o *DmrClusterLinkRemoteAddress) SetRemoteNodeName(v string)`

SetRemoteNodeName sets RemoteNodeName field to given value.

### HasRemoteNodeName

`func (o *DmrClusterLinkRemoteAddress) HasRemoteNodeName() bool`

HasRemoteNodeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


