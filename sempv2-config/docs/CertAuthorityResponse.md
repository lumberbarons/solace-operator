# CertAuthorityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CertAuthority**](CertAuthority.md) |  | [optional] 
**Links** | Pointer to [**CertAuthorityLinks**](CertAuthorityLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewCertAuthorityResponse

`func NewCertAuthorityResponse(meta SempMeta, ) *CertAuthorityResponse`

NewCertAuthorityResponse instantiates a new CertAuthorityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertAuthorityResponseWithDefaults

`func NewCertAuthorityResponseWithDefaults() *CertAuthorityResponse`

NewCertAuthorityResponseWithDefaults instantiates a new CertAuthorityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CertAuthorityResponse) GetData() CertAuthority`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CertAuthorityResponse) GetDataOk() (*CertAuthority, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CertAuthorityResponse) SetData(v CertAuthority)`

SetData sets Data field to given value.

### HasData

`func (o *CertAuthorityResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *CertAuthorityResponse) GetLinks() CertAuthorityLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CertAuthorityResponse) GetLinksOk() (*CertAuthorityLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CertAuthorityResponse) SetLinks(v CertAuthorityLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CertAuthorityResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *CertAuthorityResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CertAuthorityResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CertAuthorityResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


