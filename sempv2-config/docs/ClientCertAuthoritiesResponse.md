# ClientCertAuthoritiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ClientCertAuthority**](ClientCertAuthority.md) |  | [optional] 
**Links** | Pointer to [**[]ClientCertAuthorityLinks**](ClientCertAuthorityLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewClientCertAuthoritiesResponse

`func NewClientCertAuthoritiesResponse(meta SempMeta, ) *ClientCertAuthoritiesResponse`

NewClientCertAuthoritiesResponse instantiates a new ClientCertAuthoritiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCertAuthoritiesResponseWithDefaults

`func NewClientCertAuthoritiesResponseWithDefaults() *ClientCertAuthoritiesResponse`

NewClientCertAuthoritiesResponseWithDefaults instantiates a new ClientCertAuthoritiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ClientCertAuthoritiesResponse) GetData() []ClientCertAuthority`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ClientCertAuthoritiesResponse) GetDataOk() (*[]ClientCertAuthority, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ClientCertAuthoritiesResponse) SetData(v []ClientCertAuthority)`

SetData sets Data field to given value.

### HasData

`func (o *ClientCertAuthoritiesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *ClientCertAuthoritiesResponse) GetLinks() []ClientCertAuthorityLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ClientCertAuthoritiesResponse) GetLinksOk() (*[]ClientCertAuthorityLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ClientCertAuthoritiesResponse) SetLinks(v []ClientCertAuthorityLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ClientCertAuthoritiesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *ClientCertAuthoritiesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ClientCertAuthoritiesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ClientCertAuthoritiesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


