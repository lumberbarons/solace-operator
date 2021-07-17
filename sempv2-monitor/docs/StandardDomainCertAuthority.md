# StandardDomainCertAuthority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertAuthorityName** | Pointer to **string** | The name of the Certificate Authority. | [optional] 
**CertContent** | Pointer to **string** | The PEM formatted content for the trusted root certificate of a standard domain Certificate Authority. | [optional] 

## Methods

### NewStandardDomainCertAuthority

`func NewStandardDomainCertAuthority() *StandardDomainCertAuthority`

NewStandardDomainCertAuthority instantiates a new StandardDomainCertAuthority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardDomainCertAuthorityWithDefaults

`func NewStandardDomainCertAuthorityWithDefaults() *StandardDomainCertAuthority`

NewStandardDomainCertAuthorityWithDefaults instantiates a new StandardDomainCertAuthority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertAuthorityName

`func (o *StandardDomainCertAuthority) GetCertAuthorityName() string`

GetCertAuthorityName returns the CertAuthorityName field if non-nil, zero value otherwise.

### GetCertAuthorityNameOk

`func (o *StandardDomainCertAuthority) GetCertAuthorityNameOk() (*string, bool)`

GetCertAuthorityNameOk returns a tuple with the CertAuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAuthorityName

`func (o *StandardDomainCertAuthority) SetCertAuthorityName(v string)`

SetCertAuthorityName sets CertAuthorityName field to given value.

### HasCertAuthorityName

`func (o *StandardDomainCertAuthority) HasCertAuthorityName() bool`

HasCertAuthorityName returns a boolean if a field has been set.

### GetCertContent

`func (o *StandardDomainCertAuthority) GetCertContent() string`

GetCertContent returns the CertContent field if non-nil, zero value otherwise.

### GetCertContentOk

`func (o *StandardDomainCertAuthority) GetCertContentOk() (*string, bool)`

GetCertContentOk returns a tuple with the CertContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertContent

`func (o *StandardDomainCertAuthority) SetCertContent(v string)`

SetCertContent sets CertContent field to given value.

### HasCertContent

`func (o *StandardDomainCertAuthority) HasCertContent() bool`

HasCertContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


