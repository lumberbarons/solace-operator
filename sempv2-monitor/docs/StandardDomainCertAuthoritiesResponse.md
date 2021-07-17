# StandardDomainCertAuthoritiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]StandardDomainCertAuthority**](StandardDomainCertAuthority.md) |  | [optional] 
**Links** | Pointer to [**[]StandardDomainCertAuthorityLinks**](StandardDomainCertAuthorityLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewStandardDomainCertAuthoritiesResponse

`func NewStandardDomainCertAuthoritiesResponse(meta SempMeta, ) *StandardDomainCertAuthoritiesResponse`

NewStandardDomainCertAuthoritiesResponse instantiates a new StandardDomainCertAuthoritiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardDomainCertAuthoritiesResponseWithDefaults

`func NewStandardDomainCertAuthoritiesResponseWithDefaults() *StandardDomainCertAuthoritiesResponse`

NewStandardDomainCertAuthoritiesResponseWithDefaults instantiates a new StandardDomainCertAuthoritiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *StandardDomainCertAuthoritiesResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *StandardDomainCertAuthoritiesResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *StandardDomainCertAuthoritiesResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *StandardDomainCertAuthoritiesResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *StandardDomainCertAuthoritiesResponse) GetData() []StandardDomainCertAuthority`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StandardDomainCertAuthoritiesResponse) GetDataOk() (*[]StandardDomainCertAuthority, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StandardDomainCertAuthoritiesResponse) SetData(v []StandardDomainCertAuthority)`

SetData sets Data field to given value.

### HasData

`func (o *StandardDomainCertAuthoritiesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *StandardDomainCertAuthoritiesResponse) GetLinks() []StandardDomainCertAuthorityLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StandardDomainCertAuthoritiesResponse) GetLinksOk() (*[]StandardDomainCertAuthorityLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StandardDomainCertAuthoritiesResponse) SetLinks(v []StandardDomainCertAuthorityLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StandardDomainCertAuthoritiesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *StandardDomainCertAuthoritiesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StandardDomainCertAuthoritiesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StandardDomainCertAuthoritiesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


