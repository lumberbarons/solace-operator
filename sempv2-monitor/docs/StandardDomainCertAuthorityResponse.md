# StandardDomainCertAuthorityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**StandardDomainCertAuthority**](StandardDomainCertAuthority.md) |  | [optional] 
**Links** | Pointer to [**StandardDomainCertAuthorityLinks**](StandardDomainCertAuthorityLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewStandardDomainCertAuthorityResponse

`func NewStandardDomainCertAuthorityResponse(meta SempMeta, ) *StandardDomainCertAuthorityResponse`

NewStandardDomainCertAuthorityResponse instantiates a new StandardDomainCertAuthorityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardDomainCertAuthorityResponseWithDefaults

`func NewStandardDomainCertAuthorityResponseWithDefaults() *StandardDomainCertAuthorityResponse`

NewStandardDomainCertAuthorityResponseWithDefaults instantiates a new StandardDomainCertAuthorityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *StandardDomainCertAuthorityResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *StandardDomainCertAuthorityResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *StandardDomainCertAuthorityResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *StandardDomainCertAuthorityResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *StandardDomainCertAuthorityResponse) GetData() StandardDomainCertAuthority`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StandardDomainCertAuthorityResponse) GetDataOk() (*StandardDomainCertAuthority, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StandardDomainCertAuthorityResponse) SetData(v StandardDomainCertAuthority)`

SetData sets Data field to given value.

### HasData

`func (o *StandardDomainCertAuthorityResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *StandardDomainCertAuthorityResponse) GetLinks() StandardDomainCertAuthorityLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StandardDomainCertAuthorityResponse) GetLinksOk() (*StandardDomainCertAuthorityLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StandardDomainCertAuthorityResponse) SetLinks(v StandardDomainCertAuthorityLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StandardDomainCertAuthorityResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *StandardDomainCertAuthorityResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StandardDomainCertAuthorityResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StandardDomainCertAuthorityResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


