# EventThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClearPercent** | Pointer to **int64** | The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET. | [optional] 
**ClearValue** | Pointer to **int64** | The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET. | [optional] 
**SetPercent** | Pointer to **int64** | The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET. | [optional] 
**SetValue** | Pointer to **int64** | The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET. | [optional] 

## Methods

### NewEventThreshold

`func NewEventThreshold() *EventThreshold`

NewEventThreshold instantiates a new EventThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventThresholdWithDefaults

`func NewEventThresholdWithDefaults() *EventThreshold`

NewEventThresholdWithDefaults instantiates a new EventThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClearPercent

`func (o *EventThreshold) GetClearPercent() int64`

GetClearPercent returns the ClearPercent field if non-nil, zero value otherwise.

### GetClearPercentOk

`func (o *EventThreshold) GetClearPercentOk() (*int64, bool)`

GetClearPercentOk returns a tuple with the ClearPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearPercent

`func (o *EventThreshold) SetClearPercent(v int64)`

SetClearPercent sets ClearPercent field to given value.

### HasClearPercent

`func (o *EventThreshold) HasClearPercent() bool`

HasClearPercent returns a boolean if a field has been set.

### GetClearValue

`func (o *EventThreshold) GetClearValue() int64`

GetClearValue returns the ClearValue field if non-nil, zero value otherwise.

### GetClearValueOk

`func (o *EventThreshold) GetClearValueOk() (*int64, bool)`

GetClearValueOk returns a tuple with the ClearValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearValue

`func (o *EventThreshold) SetClearValue(v int64)`

SetClearValue sets ClearValue field to given value.

### HasClearValue

`func (o *EventThreshold) HasClearValue() bool`

HasClearValue returns a boolean if a field has been set.

### GetSetPercent

`func (o *EventThreshold) GetSetPercent() int64`

GetSetPercent returns the SetPercent field if non-nil, zero value otherwise.

### GetSetPercentOk

`func (o *EventThreshold) GetSetPercentOk() (*int64, bool)`

GetSetPercentOk returns a tuple with the SetPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetPercent

`func (o *EventThreshold) SetSetPercent(v int64)`

SetSetPercent sets SetPercent field to given value.

### HasSetPercent

`func (o *EventThreshold) HasSetPercent() bool`

HasSetPercent returns a boolean if a field has been set.

### GetSetValue

`func (o *EventThreshold) GetSetValue() int64`

GetSetValue returns the SetValue field if non-nil, zero value otherwise.

### GetSetValueOk

`func (o *EventThreshold) GetSetValueOk() (*int64, bool)`

GetSetValueOk returns a tuple with the SetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetValue

`func (o *EventThreshold) SetSetValue(v int64)`

SetSetValue sets SetValue field to given value.

### HasSetValue

`func (o *EventThreshold) HasSetValue() bool`

HasSetValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


