# ClientCertAuthority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertAuthorityName** | Pointer to **string** | The name of the Certificate Authority. | [optional] 
**CertContent** | Pointer to **string** | The PEM formatted content for the trusted root certificate of a client Certificate Authority. | [optional] 
**CrlDayList** | Pointer to **string** | The scheduled CRL refresh day(s), specified as \&quot;daily\&quot; or a comma-separated list of days. Days must be specified as \&quot;Sun\&quot;, \&quot;Mon\&quot;, \&quot;Tue\&quot;, \&quot;Wed\&quot;, \&quot;Thu\&quot;, \&quot;Fri\&quot;, or \&quot;Sat\&quot;, with no spaces, and in sorted order from Sunday to Saturday. | [optional] 
**CrlLastDownloadTime** | Pointer to **int32** | The timestamp of the last successful CRL download. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**CrlLastFailureReason** | Pointer to **string** | The reason for the last CRL failure. | [optional] 
**CrlLastFailureTime** | Pointer to **int32** | The timestamp of the last CRL failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**CrlNextDownloadTime** | Pointer to **int32** | The scheduled time of the next CRL download. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**CrlTimeList** | Pointer to **string** | The scheduled CRL refresh time(s), specified as \&quot;hourly\&quot; or a comma-separated list of 24-hour times in the form hh:mm, or h:mm. There must be no spaces, and times must be in sorted order from 0:00 to 23:59. | [optional] 
**CrlUp** | Pointer to **bool** | Indicates whether CRL revocation checking is operationally up. | [optional] 
**CrlUrl** | Pointer to **string** | The URL for the CRL source. This is a required attribute for CRL to be operational and the URL must be complete with http:// included. | [optional] 
**OcspLastFailureReason** | Pointer to **string** | The reason for the last OCSP failure. | [optional] 
**OcspLastFailureTime** | Pointer to **int32** | The timestamp of the last OCSP failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**OcspLastFailureUrl** | Pointer to **string** | The URL involved in the last OCSP failure. | [optional] 
**OcspNonResponderCertEnabled** | Pointer to **bool** | Indicates whether a non-responder certificate is allowed to sign an OCSP response. Typically used with an OCSP override URL in cases where a single certificate is used to sign client certificates and OCSP responses. | [optional] 
**OcspOverrideUrl** | Pointer to **string** | The OCSP responder URL to use for overriding the one supplied in the client certificate. The URL must be complete with http:// included. | [optional] 
**OcspTimeout** | Pointer to **int64** | The timeout in seconds to receive a response from the OCSP responder after sending a request or making the initial connection attempt. | [optional] 
**RevocationCheckEnabled** | Pointer to **bool** | Indicates whether Certificate Authority revocation checking is enabled. | [optional] 

## Methods

### NewClientCertAuthority

`func NewClientCertAuthority() *ClientCertAuthority`

NewClientCertAuthority instantiates a new ClientCertAuthority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCertAuthorityWithDefaults

`func NewClientCertAuthorityWithDefaults() *ClientCertAuthority`

NewClientCertAuthorityWithDefaults instantiates a new ClientCertAuthority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertAuthorityName

`func (o *ClientCertAuthority) GetCertAuthorityName() string`

GetCertAuthorityName returns the CertAuthorityName field if non-nil, zero value otherwise.

### GetCertAuthorityNameOk

`func (o *ClientCertAuthority) GetCertAuthorityNameOk() (*string, bool)`

GetCertAuthorityNameOk returns a tuple with the CertAuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAuthorityName

`func (o *ClientCertAuthority) SetCertAuthorityName(v string)`

SetCertAuthorityName sets CertAuthorityName field to given value.

### HasCertAuthorityName

`func (o *ClientCertAuthority) HasCertAuthorityName() bool`

HasCertAuthorityName returns a boolean if a field has been set.

### GetCertContent

`func (o *ClientCertAuthority) GetCertContent() string`

GetCertContent returns the CertContent field if non-nil, zero value otherwise.

### GetCertContentOk

`func (o *ClientCertAuthority) GetCertContentOk() (*string, bool)`

GetCertContentOk returns a tuple with the CertContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertContent

`func (o *ClientCertAuthority) SetCertContent(v string)`

SetCertContent sets CertContent field to given value.

### HasCertContent

`func (o *ClientCertAuthority) HasCertContent() bool`

HasCertContent returns a boolean if a field has been set.

### GetCrlDayList

`func (o *ClientCertAuthority) GetCrlDayList() string`

GetCrlDayList returns the CrlDayList field if non-nil, zero value otherwise.

### GetCrlDayListOk

`func (o *ClientCertAuthority) GetCrlDayListOk() (*string, bool)`

GetCrlDayListOk returns a tuple with the CrlDayList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDayList

`func (o *ClientCertAuthority) SetCrlDayList(v string)`

SetCrlDayList sets CrlDayList field to given value.

### HasCrlDayList

`func (o *ClientCertAuthority) HasCrlDayList() bool`

HasCrlDayList returns a boolean if a field has been set.

### GetCrlLastDownloadTime

`func (o *ClientCertAuthority) GetCrlLastDownloadTime() int32`

GetCrlLastDownloadTime returns the CrlLastDownloadTime field if non-nil, zero value otherwise.

### GetCrlLastDownloadTimeOk

`func (o *ClientCertAuthority) GetCrlLastDownloadTimeOk() (*int32, bool)`

GetCrlLastDownloadTimeOk returns a tuple with the CrlLastDownloadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlLastDownloadTime

`func (o *ClientCertAuthority) SetCrlLastDownloadTime(v int32)`

SetCrlLastDownloadTime sets CrlLastDownloadTime field to given value.

### HasCrlLastDownloadTime

`func (o *ClientCertAuthority) HasCrlLastDownloadTime() bool`

HasCrlLastDownloadTime returns a boolean if a field has been set.

### GetCrlLastFailureReason

`func (o *ClientCertAuthority) GetCrlLastFailureReason() string`

GetCrlLastFailureReason returns the CrlLastFailureReason field if non-nil, zero value otherwise.

### GetCrlLastFailureReasonOk

`func (o *ClientCertAuthority) GetCrlLastFailureReasonOk() (*string, bool)`

GetCrlLastFailureReasonOk returns a tuple with the CrlLastFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlLastFailureReason

`func (o *ClientCertAuthority) SetCrlLastFailureReason(v string)`

SetCrlLastFailureReason sets CrlLastFailureReason field to given value.

### HasCrlLastFailureReason

`func (o *ClientCertAuthority) HasCrlLastFailureReason() bool`

HasCrlLastFailureReason returns a boolean if a field has been set.

### GetCrlLastFailureTime

`func (o *ClientCertAuthority) GetCrlLastFailureTime() int32`

GetCrlLastFailureTime returns the CrlLastFailureTime field if non-nil, zero value otherwise.

### GetCrlLastFailureTimeOk

`func (o *ClientCertAuthority) GetCrlLastFailureTimeOk() (*int32, bool)`

GetCrlLastFailureTimeOk returns a tuple with the CrlLastFailureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlLastFailureTime

`func (o *ClientCertAuthority) SetCrlLastFailureTime(v int32)`

SetCrlLastFailureTime sets CrlLastFailureTime field to given value.

### HasCrlLastFailureTime

`func (o *ClientCertAuthority) HasCrlLastFailureTime() bool`

HasCrlLastFailureTime returns a boolean if a field has been set.

### GetCrlNextDownloadTime

`func (o *ClientCertAuthority) GetCrlNextDownloadTime() int32`

GetCrlNextDownloadTime returns the CrlNextDownloadTime field if non-nil, zero value otherwise.

### GetCrlNextDownloadTimeOk

`func (o *ClientCertAuthority) GetCrlNextDownloadTimeOk() (*int32, bool)`

GetCrlNextDownloadTimeOk returns a tuple with the CrlNextDownloadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlNextDownloadTime

`func (o *ClientCertAuthority) SetCrlNextDownloadTime(v int32)`

SetCrlNextDownloadTime sets CrlNextDownloadTime field to given value.

### HasCrlNextDownloadTime

`func (o *ClientCertAuthority) HasCrlNextDownloadTime() bool`

HasCrlNextDownloadTime returns a boolean if a field has been set.

### GetCrlTimeList

`func (o *ClientCertAuthority) GetCrlTimeList() string`

GetCrlTimeList returns the CrlTimeList field if non-nil, zero value otherwise.

### GetCrlTimeListOk

`func (o *ClientCertAuthority) GetCrlTimeListOk() (*string, bool)`

GetCrlTimeListOk returns a tuple with the CrlTimeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlTimeList

`func (o *ClientCertAuthority) SetCrlTimeList(v string)`

SetCrlTimeList sets CrlTimeList field to given value.

### HasCrlTimeList

`func (o *ClientCertAuthority) HasCrlTimeList() bool`

HasCrlTimeList returns a boolean if a field has been set.

### GetCrlUp

`func (o *ClientCertAuthority) GetCrlUp() bool`

GetCrlUp returns the CrlUp field if non-nil, zero value otherwise.

### GetCrlUpOk

`func (o *ClientCertAuthority) GetCrlUpOk() (*bool, bool)`

GetCrlUpOk returns a tuple with the CrlUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlUp

`func (o *ClientCertAuthority) SetCrlUp(v bool)`

SetCrlUp sets CrlUp field to given value.

### HasCrlUp

`func (o *ClientCertAuthority) HasCrlUp() bool`

HasCrlUp returns a boolean if a field has been set.

### GetCrlUrl

`func (o *ClientCertAuthority) GetCrlUrl() string`

GetCrlUrl returns the CrlUrl field if non-nil, zero value otherwise.

### GetCrlUrlOk

`func (o *ClientCertAuthority) GetCrlUrlOk() (*string, bool)`

GetCrlUrlOk returns a tuple with the CrlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlUrl

`func (o *ClientCertAuthority) SetCrlUrl(v string)`

SetCrlUrl sets CrlUrl field to given value.

### HasCrlUrl

`func (o *ClientCertAuthority) HasCrlUrl() bool`

HasCrlUrl returns a boolean if a field has been set.

### GetOcspLastFailureReason

`func (o *ClientCertAuthority) GetOcspLastFailureReason() string`

GetOcspLastFailureReason returns the OcspLastFailureReason field if non-nil, zero value otherwise.

### GetOcspLastFailureReasonOk

`func (o *ClientCertAuthority) GetOcspLastFailureReasonOk() (*string, bool)`

GetOcspLastFailureReasonOk returns a tuple with the OcspLastFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspLastFailureReason

`func (o *ClientCertAuthority) SetOcspLastFailureReason(v string)`

SetOcspLastFailureReason sets OcspLastFailureReason field to given value.

### HasOcspLastFailureReason

`func (o *ClientCertAuthority) HasOcspLastFailureReason() bool`

HasOcspLastFailureReason returns a boolean if a field has been set.

### GetOcspLastFailureTime

`func (o *ClientCertAuthority) GetOcspLastFailureTime() int32`

GetOcspLastFailureTime returns the OcspLastFailureTime field if non-nil, zero value otherwise.

### GetOcspLastFailureTimeOk

`func (o *ClientCertAuthority) GetOcspLastFailureTimeOk() (*int32, bool)`

GetOcspLastFailureTimeOk returns a tuple with the OcspLastFailureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspLastFailureTime

`func (o *ClientCertAuthority) SetOcspLastFailureTime(v int32)`

SetOcspLastFailureTime sets OcspLastFailureTime field to given value.

### HasOcspLastFailureTime

`func (o *ClientCertAuthority) HasOcspLastFailureTime() bool`

HasOcspLastFailureTime returns a boolean if a field has been set.

### GetOcspLastFailureUrl

`func (o *ClientCertAuthority) GetOcspLastFailureUrl() string`

GetOcspLastFailureUrl returns the OcspLastFailureUrl field if non-nil, zero value otherwise.

### GetOcspLastFailureUrlOk

`func (o *ClientCertAuthority) GetOcspLastFailureUrlOk() (*string, bool)`

GetOcspLastFailureUrlOk returns a tuple with the OcspLastFailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspLastFailureUrl

`func (o *ClientCertAuthority) SetOcspLastFailureUrl(v string)`

SetOcspLastFailureUrl sets OcspLastFailureUrl field to given value.

### HasOcspLastFailureUrl

`func (o *ClientCertAuthority) HasOcspLastFailureUrl() bool`

HasOcspLastFailureUrl returns a boolean if a field has been set.

### GetOcspNonResponderCertEnabled

`func (o *ClientCertAuthority) GetOcspNonResponderCertEnabled() bool`

GetOcspNonResponderCertEnabled returns the OcspNonResponderCertEnabled field if non-nil, zero value otherwise.

### GetOcspNonResponderCertEnabledOk

`func (o *ClientCertAuthority) GetOcspNonResponderCertEnabledOk() (*bool, bool)`

GetOcspNonResponderCertEnabledOk returns a tuple with the OcspNonResponderCertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspNonResponderCertEnabled

`func (o *ClientCertAuthority) SetOcspNonResponderCertEnabled(v bool)`

SetOcspNonResponderCertEnabled sets OcspNonResponderCertEnabled field to given value.

### HasOcspNonResponderCertEnabled

`func (o *ClientCertAuthority) HasOcspNonResponderCertEnabled() bool`

HasOcspNonResponderCertEnabled returns a boolean if a field has been set.

### GetOcspOverrideUrl

`func (o *ClientCertAuthority) GetOcspOverrideUrl() string`

GetOcspOverrideUrl returns the OcspOverrideUrl field if non-nil, zero value otherwise.

### GetOcspOverrideUrlOk

`func (o *ClientCertAuthority) GetOcspOverrideUrlOk() (*string, bool)`

GetOcspOverrideUrlOk returns a tuple with the OcspOverrideUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspOverrideUrl

`func (o *ClientCertAuthority) SetOcspOverrideUrl(v string)`

SetOcspOverrideUrl sets OcspOverrideUrl field to given value.

### HasOcspOverrideUrl

`func (o *ClientCertAuthority) HasOcspOverrideUrl() bool`

HasOcspOverrideUrl returns a boolean if a field has been set.

### GetOcspTimeout

`func (o *ClientCertAuthority) GetOcspTimeout() int64`

GetOcspTimeout returns the OcspTimeout field if non-nil, zero value otherwise.

### GetOcspTimeoutOk

`func (o *ClientCertAuthority) GetOcspTimeoutOk() (*int64, bool)`

GetOcspTimeoutOk returns a tuple with the OcspTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspTimeout

`func (o *ClientCertAuthority) SetOcspTimeout(v int64)`

SetOcspTimeout sets OcspTimeout field to given value.

### HasOcspTimeout

`func (o *ClientCertAuthority) HasOcspTimeout() bool`

HasOcspTimeout returns a boolean if a field has been set.

### GetRevocationCheckEnabled

`func (o *ClientCertAuthority) GetRevocationCheckEnabled() bool`

GetRevocationCheckEnabled returns the RevocationCheckEnabled field if non-nil, zero value otherwise.

### GetRevocationCheckEnabledOk

`func (o *ClientCertAuthority) GetRevocationCheckEnabledOk() (*bool, bool)`

GetRevocationCheckEnabledOk returns a tuple with the RevocationCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationCheckEnabled

`func (o *ClientCertAuthority) SetRevocationCheckEnabled(v bool)`

SetRevocationCheckEnabled sets RevocationCheckEnabled field to given value.

### HasRevocationCheckEnabled

`func (o *ClientCertAuthority) HasRevocationCheckEnabled() bool`

HasRevocationCheckEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


