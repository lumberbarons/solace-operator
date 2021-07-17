# CertAuthorityOcspTlsTrustedCommonName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertAuthorityName** | Pointer to **string** | The name of the Certificate Authority. Deprecated since 2.19. Replaced by clientCertAuthorities. | [optional] 
**OcspTlsTrustedCommonName** | Pointer to **string** | The expected Trusted Common Name of the OCSP responder remote certificate. Deprecated since 2.19. Replaced by clientCertAuthorities. | [optional] 

## Methods

### NewCertAuthorityOcspTlsTrustedCommonName

`func NewCertAuthorityOcspTlsTrustedCommonName() *CertAuthorityOcspTlsTrustedCommonName`

NewCertAuthorityOcspTlsTrustedCommonName instantiates a new CertAuthorityOcspTlsTrustedCommonName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertAuthorityOcspTlsTrustedCommonNameWithDefaults

`func NewCertAuthorityOcspTlsTrustedCommonNameWithDefaults() *CertAuthorityOcspTlsTrustedCommonName`

NewCertAuthorityOcspTlsTrustedCommonNameWithDefaults instantiates a new CertAuthorityOcspTlsTrustedCommonName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertAuthorityName

`func (o *CertAuthorityOcspTlsTrustedCommonName) GetCertAuthorityName() string`

GetCertAuthorityName returns the CertAuthorityName field if non-nil, zero value otherwise.

### GetCertAuthorityNameOk

`func (o *CertAuthorityOcspTlsTrustedCommonName) GetCertAuthorityNameOk() (*string, bool)`

GetCertAuthorityNameOk returns a tuple with the CertAuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAuthorityName

`func (o *CertAuthorityOcspTlsTrustedCommonName) SetCertAuthorityName(v string)`

SetCertAuthorityName sets CertAuthorityName field to given value.

### HasCertAuthorityName

`func (o *CertAuthorityOcspTlsTrustedCommonName) HasCertAuthorityName() bool`

HasCertAuthorityName returns a boolean if a field has been set.

### GetOcspTlsTrustedCommonName

`func (o *CertAuthorityOcspTlsTrustedCommonName) GetOcspTlsTrustedCommonName() string`

GetOcspTlsTrustedCommonName returns the OcspTlsTrustedCommonName field if non-nil, zero value otherwise.

### GetOcspTlsTrustedCommonNameOk

`func (o *CertAuthorityOcspTlsTrustedCommonName) GetOcspTlsTrustedCommonNameOk() (*string, bool)`

GetOcspTlsTrustedCommonNameOk returns a tuple with the OcspTlsTrustedCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspTlsTrustedCommonName

`func (o *CertAuthorityOcspTlsTrustedCommonName) SetOcspTlsTrustedCommonName(v string)`

SetOcspTlsTrustedCommonName sets OcspTlsTrustedCommonName field to given value.

### HasOcspTlsTrustedCommonName

`func (o *CertAuthorityOcspTlsTrustedCommonName) HasOcspTlsTrustedCommonName() bool`

HasOcspTlsTrustedCommonName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


