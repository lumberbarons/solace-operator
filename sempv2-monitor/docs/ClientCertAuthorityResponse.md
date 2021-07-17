# ClientCertAuthorityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**ClientCertAuthorityCollections**](ClientCertAuthorityCollections.md) |  | [optional] 
**Data** | Pointer to [**ClientCertAuthority**](ClientCertAuthority.md) |  | [optional] 
**Links** | Pointer to [**ClientCertAuthorityLinks**](ClientCertAuthorityLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewClientCertAuthorityResponse

`func NewClientCertAuthorityResponse(meta SempMeta, ) *ClientCertAuthorityResponse`

NewClientCertAuthorityResponse instantiates a new ClientCertAuthorityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCertAuthorityResponseWithDefaults

`func NewClientCertAuthorityResponseWithDefaults() *ClientCertAuthorityResponse`

NewClientCertAuthorityResponseWithDefaults instantiates a new ClientCertAuthorityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *ClientCertAuthorityResponse) GetCollections() ClientCertAuthorityCollections`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *ClientCertAuthorityResponse) GetCollectionsOk() (*ClientCertAuthorityCollections, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *ClientCertAuthorityResponse) SetCollections(v ClientCertAuthorityCollections)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *ClientCertAuthorityResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *ClientCertAuthorityResponse) GetData() ClientCertAuthority`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ClientCertAuthorityResponse) GetDataOk() (*ClientCertAuthority, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ClientCertAuthorityResponse) SetData(v ClientCertAuthority)`

SetData sets Data field to given value.

### HasData

`func (o *ClientCertAuthorityResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *ClientCertAuthorityResponse) GetLinks() ClientCertAuthorityLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ClientCertAuthorityResponse) GetLinksOk() (*ClientCertAuthorityLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ClientCertAuthorityResponse) SetLinks(v ClientCertAuthorityLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ClientCertAuthorityResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *ClientCertAuthorityResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ClientCertAuthorityResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ClientCertAuthorityResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


