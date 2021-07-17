# SystemInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** | The platform running the SEMP API. Deprecated since 2.2. /systemInformation was replaced by /about/api. | [optional] 
**SempVersion** | Pointer to **string** | The version of the SEMP API. Deprecated since 2.2. /systemInformation was replaced by /about/api. | [optional] 

## Methods

### NewSystemInformation

`func NewSystemInformation() *SystemInformation`

NewSystemInformation instantiates a new SystemInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInformationWithDefaults

`func NewSystemInformationWithDefaults() *SystemInformation`

NewSystemInformationWithDefaults instantiates a new SystemInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *SystemInformation) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *SystemInformation) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *SystemInformation) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *SystemInformation) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSempVersion

`func (o *SystemInformation) GetSempVersion() string`

GetSempVersion returns the SempVersion field if non-nil, zero value otherwise.

### GetSempVersionOk

`func (o *SystemInformation) GetSempVersionOk() (*string, bool)`

GetSempVersionOk returns a tuple with the SempVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSempVersion

`func (o *SystemInformation) SetSempVersion(v string)`

SetSempVersion sets SempVersion field to given value.

### HasSempVersion

`func (o *SystemInformation) HasSempVersion() bool`

HasSempVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


