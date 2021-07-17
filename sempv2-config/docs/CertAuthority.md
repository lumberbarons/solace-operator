# CertAuthority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertAuthorityName** | Pointer to **string** | The name of the Certificate Authority. Deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities. | [optional] 
**CertContent** | Pointer to **string** | The PEM formatted content for the trusted root certificate of a Certificate Authority. The default value is &#x60;\&quot;\&quot;&#x60;. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] 
**CrlDayList** | Pointer to **string** | The scheduled CRL refresh day(s), specified as \&quot;daily\&quot; or a comma-separated list of days. Days must be specified as \&quot;Sun\&quot;, \&quot;Mon\&quot;, \&quot;Tue\&quot;, \&quot;Wed\&quot;, \&quot;Thu\&quot;, \&quot;Fri\&quot;, or \&quot;Sat\&quot;, with no spaces, and in sorted order from Sunday to Saturday. The default value is &#x60;\&quot;daily\&quot;&#x60;. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] 
**CrlTimeList** | Pointer to **string** | The scheduled CRL refresh time(s), specified as \&quot;hourly\&quot; or a comma-separated list of 24-hour times in the form hh:mm, or h:mm. There must be no spaces, and times must be in sorted order from 0:00 to 23:59. The default value is &#x60;\&quot;3:00\&quot;&#x60;. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] 
**CrlUrl** | Pointer to **string** | The URL for the CRL source. This is a required attribute for CRL to be operational and the URL must be complete with http:// included. The default value is &#x60;\&quot;\&quot;&#x60;. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] 
**OcspNonResponderCertEnabled** | Pointer to **bool** | Enable or disable allowing a non-responder certificate to sign an OCSP response. Typically used with an OCSP override URL in cases where a single certificate is used to sign client certificates and OCSP responses. The default value is &#x60;false&#x60;. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] 
**OcspOverrideUrl** | Pointer to **string** | The OCSP responder URL to use for overriding the one supplied in the client certificate. The URL must be complete with http:// included. The default value is &#x60;\&quot;\&quot;&#x60;. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] 
**OcspTimeout** | Pointer to **int64** | The timeout in seconds to receive a response from the OCSP responder after sending a request or making the initial connection attempt. The default value is &#x60;5&#x60;. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] 
**RevocationCheckEnabled** | Pointer to **bool** | Enable or disable Certificate Authority revocation checking. The default value is &#x60;false&#x60;. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] 

## Methods

### NewCertAuthority

`func NewCertAuthority() *CertAuthority`

NewCertAuthority instantiates a new CertAuthority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertAuthorityWithDefaults

`func NewCertAuthorityWithDefaults() *CertAuthority`

NewCertAuthorityWithDefaults instantiates a new CertAuthority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertAuthorityName

`func (o *CertAuthority) GetCertAuthorityName() string`

GetCertAuthorityName returns the CertAuthorityName field if non-nil, zero value otherwise.

### GetCertAuthorityNameOk

`func (o *CertAuthority) GetCertAuthorityNameOk() (*string, bool)`

GetCertAuthorityNameOk returns a tuple with the CertAuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAuthorityName

`func (o *CertAuthority) SetCertAuthorityName(v string)`

SetCertAuthorityName sets CertAuthorityName field to given value.

### HasCertAuthorityName

`func (o *CertAuthority) HasCertAuthorityName() bool`

HasCertAuthorityName returns a boolean if a field has been set.

### GetCertContent

`func (o *CertAuthority) GetCertContent() string`

GetCertContent returns the CertContent field if non-nil, zero value otherwise.

### GetCertContentOk

`func (o *CertAuthority) GetCertContentOk() (*string, bool)`

GetCertContentOk returns a tuple with the CertContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertContent

`func (o *CertAuthority) SetCertContent(v string)`

SetCertContent sets CertContent field to given value.

### HasCertContent

`func (o *CertAuthority) HasCertContent() bool`

HasCertContent returns a boolean if a field has been set.

### GetCrlDayList

`func (o *CertAuthority) GetCrlDayList() string`

GetCrlDayList returns the CrlDayList field if non-nil, zero value otherwise.

### GetCrlDayListOk

`func (o *CertAuthority) GetCrlDayListOk() (*string, bool)`

GetCrlDayListOk returns a tuple with the CrlDayList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDayList

`func (o *CertAuthority) SetCrlDayList(v string)`

SetCrlDayList sets CrlDayList field to given value.

### HasCrlDayList

`func (o *CertAuthority) HasCrlDayList() bool`

HasCrlDayList returns a boolean if a field has been set.

### GetCrlTimeList

`func (o *CertAuthority) GetCrlTimeList() string`

GetCrlTimeList returns the CrlTimeList field if non-nil, zero value otherwise.

### GetCrlTimeListOk

`func (o *CertAuthority) GetCrlTimeListOk() (*string, bool)`

GetCrlTimeListOk returns a tuple with the CrlTimeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlTimeList

`func (o *CertAuthority) SetCrlTimeList(v string)`

SetCrlTimeList sets CrlTimeList field to given value.

### HasCrlTimeList

`func (o *CertAuthority) HasCrlTimeList() bool`

HasCrlTimeList returns a boolean if a field has been set.

### GetCrlUrl

`func (o *CertAuthority) GetCrlUrl() string`

GetCrlUrl returns the CrlUrl field if non-nil, zero value otherwise.

### GetCrlUrlOk

`func (o *CertAuthority) GetCrlUrlOk() (*string, bool)`

GetCrlUrlOk returns a tuple with the CrlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlUrl

`func (o *CertAuthority) SetCrlUrl(v string)`

SetCrlUrl sets CrlUrl field to given value.

### HasCrlUrl

`func (o *CertAuthority) HasCrlUrl() bool`

HasCrlUrl returns a boolean if a field has been set.

### GetOcspNonResponderCertEnabled

`func (o *CertAuthority) GetOcspNonResponderCertEnabled() bool`

GetOcspNonResponderCertEnabled returns the OcspNonResponderCertEnabled field if non-nil, zero value otherwise.

### GetOcspNonResponderCertEnabledOk

`func (o *CertAuthority) GetOcspNonResponderCertEnabledOk() (*bool, bool)`

GetOcspNonResponderCertEnabledOk returns a tuple with the OcspNonResponderCertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspNonResponderCertEnabled

`func (o *CertAuthority) SetOcspNonResponderCertEnabled(v bool)`

SetOcspNonResponderCertEnabled sets OcspNonResponderCertEnabled field to given value.

### HasOcspNonResponderCertEnabled

`func (o *CertAuthority) HasOcspNonResponderCertEnabled() bool`

HasOcspNonResponderCertEnabled returns a boolean if a field has been set.

### GetOcspOverrideUrl

`func (o *CertAuthority) GetOcspOverrideUrl() string`

GetOcspOverrideUrl returns the OcspOverrideUrl field if non-nil, zero value otherwise.

### GetOcspOverrideUrlOk

`func (o *CertAuthority) GetOcspOverrideUrlOk() (*string, bool)`

GetOcspOverrideUrlOk returns a tuple with the OcspOverrideUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspOverrideUrl

`func (o *CertAuthority) SetOcspOverrideUrl(v string)`

SetOcspOverrideUrl sets OcspOverrideUrl field to given value.

### HasOcspOverrideUrl

`func (o *CertAuthority) HasOcspOverrideUrl() bool`

HasOcspOverrideUrl returns a boolean if a field has been set.

### GetOcspTimeout

`func (o *CertAuthority) GetOcspTimeout() int64`

GetOcspTimeout returns the OcspTimeout field if non-nil, zero value otherwise.

### GetOcspTimeoutOk

`func (o *CertAuthority) GetOcspTimeoutOk() (*int64, bool)`

GetOcspTimeoutOk returns a tuple with the OcspTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspTimeout

`func (o *CertAuthority) SetOcspTimeout(v int64)`

SetOcspTimeout sets OcspTimeout field to given value.

### HasOcspTimeout

`func (o *CertAuthority) HasOcspTimeout() bool`

HasOcspTimeout returns a boolean if a field has been set.

### GetRevocationCheckEnabled

`func (o *CertAuthority) GetRevocationCheckEnabled() bool`

GetRevocationCheckEnabled returns the RevocationCheckEnabled field if non-nil, zero value otherwise.

### GetRevocationCheckEnabledOk

`func (o *CertAuthority) GetRevocationCheckEnabledOk() (*bool, bool)`

GetRevocationCheckEnabledOk returns a tuple with the RevocationCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationCheckEnabled

`func (o *CertAuthority) SetRevocationCheckEnabled(v bool)`

SetRevocationCheckEnabled sets RevocationCheckEnabled field to given value.

### HasRevocationCheckEnabled

`func (o *CertAuthority) HasRevocationCheckEnabled() bool`

HasRevocationCheckEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


