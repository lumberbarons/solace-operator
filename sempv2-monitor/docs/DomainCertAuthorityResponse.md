# DomainCertAuthorityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**DomainCertAuthority**](DomainCertAuthority.md) |  | [optional] 
**Links** | Pointer to [**DomainCertAuthorityLinks**](DomainCertAuthorityLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewDomainCertAuthorityResponse

`func NewDomainCertAuthorityResponse(meta SempMeta, ) *DomainCertAuthorityResponse`

NewDomainCertAuthorityResponse instantiates a new DomainCertAuthorityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainCertAuthorityResponseWithDefaults

`func NewDomainCertAuthorityResponseWithDefaults() *DomainCertAuthorityResponse`

NewDomainCertAuthorityResponseWithDefaults instantiates a new DomainCertAuthorityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *DomainCertAuthorityResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *DomainCertAuthorityResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *DomainCertAuthorityResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *DomainCertAuthorityResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *DomainCertAuthorityResponse) GetData() DomainCertAuthority`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainCertAuthorityResponse) GetDataOk() (*DomainCertAuthority, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainCertAuthorityResponse) SetData(v DomainCertAuthority)`

SetData sets Data field to given value.

### HasData

`func (o *DomainCertAuthorityResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *DomainCertAuthorityResponse) GetLinks() DomainCertAuthorityLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DomainCertAuthorityResponse) GetLinksOk() (*DomainCertAuthorityLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DomainCertAuthorityResponse) SetLinks(v DomainCertAuthorityLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DomainCertAuthorityResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *DomainCertAuthorityResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DomainCertAuthorityResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DomainCertAuthorityResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


