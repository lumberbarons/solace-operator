# MsgVpnRestDeliveryPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | Pointer to **string** | The name of the Client for the REST Delivery Point. | [optional] 
**ClientProfileName** | Pointer to **string** | The name of the Client Profile for the REST Delivery Point. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the REST Delivery Point is enabled. | [optional] 
**LastFailureReason** | Pointer to **string** | The reason for the last REST Delivery Point failure. | [optional] 
**LastFailureTime** | Pointer to **int32** | The timestamp of the last REST Delivery Point failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**RestDeliveryPointName** | Pointer to **string** | The name of the REST Delivery Point. | [optional] 
**Service** | Pointer to **string** | The name of the service that this REST Delivery Point connects to. Internally the broker does not use this value; it is informational only. Available since 2.19. | [optional] 
**TimeConnectionsBlocked** | Pointer to **int64** | The percentage of time the REST Delivery Point connections are blocked from transmitting data. | [optional] 
**Up** | Pointer to **bool** | Indicates whether the operational state of the REST Delivery Point is up. | [optional] 
**Vendor** | Pointer to **string** | The name of the vendor that this REST Delivery Point connects to. Internally the broker does not use this value; it is informational only. Available since 2.19. | [optional] 

## Methods

### NewMsgVpnRestDeliveryPoint

`func NewMsgVpnRestDeliveryPoint() *MsgVpnRestDeliveryPoint`

NewMsgVpnRestDeliveryPoint instantiates a new MsgVpnRestDeliveryPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnRestDeliveryPointWithDefaults

`func NewMsgVpnRestDeliveryPointWithDefaults() *MsgVpnRestDeliveryPoint`

NewMsgVpnRestDeliveryPointWithDefaults instantiates a new MsgVpnRestDeliveryPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *MsgVpnRestDeliveryPoint) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *MsgVpnRestDeliveryPoint) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *MsgVpnRestDeliveryPoint) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *MsgVpnRestDeliveryPoint) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientProfileName

`func (o *MsgVpnRestDeliveryPoint) GetClientProfileName() string`

GetClientProfileName returns the ClientProfileName field if non-nil, zero value otherwise.

### GetClientProfileNameOk

`func (o *MsgVpnRestDeliveryPoint) GetClientProfileNameOk() (*string, bool)`

GetClientProfileNameOk returns a tuple with the ClientProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileName

`func (o *MsgVpnRestDeliveryPoint) SetClientProfileName(v string)`

SetClientProfileName sets ClientProfileName field to given value.

### HasClientProfileName

`func (o *MsgVpnRestDeliveryPoint) HasClientProfileName() bool`

HasClientProfileName returns a boolean if a field has been set.

### GetEnabled

`func (o *MsgVpnRestDeliveryPoint) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MsgVpnRestDeliveryPoint) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MsgVpnRestDeliveryPoint) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MsgVpnRestDeliveryPoint) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLastFailureReason

`func (o *MsgVpnRestDeliveryPoint) GetLastFailureReason() string`

GetLastFailureReason returns the LastFailureReason field if non-nil, zero value otherwise.

### GetLastFailureReasonOk

`func (o *MsgVpnRestDeliveryPoint) GetLastFailureReasonOk() (*string, bool)`

GetLastFailureReasonOk returns a tuple with the LastFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureReason

`func (o *MsgVpnRestDeliveryPoint) SetLastFailureReason(v string)`

SetLastFailureReason sets LastFailureReason field to given value.

### HasLastFailureReason

`func (o *MsgVpnRestDeliveryPoint) HasLastFailureReason() bool`

HasLastFailureReason returns a boolean if a field has been set.

### GetLastFailureTime

`func (o *MsgVpnRestDeliveryPoint) GetLastFailureTime() int32`

GetLastFailureTime returns the LastFailureTime field if non-nil, zero value otherwise.

### GetLastFailureTimeOk

`func (o *MsgVpnRestDeliveryPoint) GetLastFailureTimeOk() (*int32, bool)`

GetLastFailureTimeOk returns a tuple with the LastFailureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureTime

`func (o *MsgVpnRestDeliveryPoint) SetLastFailureTime(v int32)`

SetLastFailureTime sets LastFailureTime field to given value.

### HasLastFailureTime

`func (o *MsgVpnRestDeliveryPoint) HasLastFailureTime() bool`

HasLastFailureTime returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnRestDeliveryPoint) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnRestDeliveryPoint) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnRestDeliveryPoint) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnRestDeliveryPoint) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetRestDeliveryPointName

`func (o *MsgVpnRestDeliveryPoint) GetRestDeliveryPointName() string`

GetRestDeliveryPointName returns the RestDeliveryPointName field if non-nil, zero value otherwise.

### GetRestDeliveryPointNameOk

`func (o *MsgVpnRestDeliveryPoint) GetRestDeliveryPointNameOk() (*string, bool)`

GetRestDeliveryPointNameOk returns a tuple with the RestDeliveryPointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestDeliveryPointName

`func (o *MsgVpnRestDeliveryPoint) SetRestDeliveryPointName(v string)`

SetRestDeliveryPointName sets RestDeliveryPointName field to given value.

### HasRestDeliveryPointName

`func (o *MsgVpnRestDeliveryPoint) HasRestDeliveryPointName() bool`

HasRestDeliveryPointName returns a boolean if a field has been set.

### GetService

`func (o *MsgVpnRestDeliveryPoint) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MsgVpnRestDeliveryPoint) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *MsgVpnRestDeliveryPoint) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *MsgVpnRestDeliveryPoint) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTimeConnectionsBlocked

`func (o *MsgVpnRestDeliveryPoint) GetTimeConnectionsBlocked() int64`

GetTimeConnectionsBlocked returns the TimeConnectionsBlocked field if non-nil, zero value otherwise.

### GetTimeConnectionsBlockedOk

`func (o *MsgVpnRestDeliveryPoint) GetTimeConnectionsBlockedOk() (*int64, bool)`

GetTimeConnectionsBlockedOk returns a tuple with the TimeConnectionsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeConnectionsBlocked

`func (o *MsgVpnRestDeliveryPoint) SetTimeConnectionsBlocked(v int64)`

SetTimeConnectionsBlocked sets TimeConnectionsBlocked field to given value.

### HasTimeConnectionsBlocked

`func (o *MsgVpnRestDeliveryPoint) HasTimeConnectionsBlocked() bool`

HasTimeConnectionsBlocked returns a boolean if a field has been set.

### GetUp

`func (o *MsgVpnRestDeliveryPoint) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *MsgVpnRestDeliveryPoint) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *MsgVpnRestDeliveryPoint) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *MsgVpnRestDeliveryPoint) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetVendor

`func (o *MsgVpnRestDeliveryPoint) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *MsgVpnRestDeliveryPoint) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *MsgVpnRestDeliveryPoint) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *MsgVpnRestDeliveryPoint) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


