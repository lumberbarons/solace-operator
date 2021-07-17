# EventThresholdByPercent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClearPercent** | Pointer to **int64** | The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. | [optional] 
**SetPercent** | Pointer to **int64** | The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. | [optional] 

## Methods

### NewEventThresholdByPercent

`func NewEventThresholdByPercent() *EventThresholdByPercent`

NewEventThresholdByPercent instantiates a new EventThresholdByPercent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventThresholdByPercentWithDefaults

`func NewEventThresholdByPercentWithDefaults() *EventThresholdByPercent`

NewEventThresholdByPercentWithDefaults instantiates a new EventThresholdByPercent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClearPercent

`func (o *EventThresholdByPercent) GetClearPercent() int64`

GetClearPercent returns the ClearPercent field if non-nil, zero value otherwise.

### GetClearPercentOk

`func (o *EventThresholdByPercent) GetClearPercentOk() (*int64, bool)`

GetClearPercentOk returns a tuple with the ClearPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearPercent

`func (o *EventThresholdByPercent) SetClearPercent(v int64)`

SetClearPercent sets ClearPercent field to given value.

### HasClearPercent

`func (o *EventThresholdByPercent) HasClearPercent() bool`

HasClearPercent returns a boolean if a field has been set.

### GetSetPercent

`func (o *EventThresholdByPercent) GetSetPercent() int64`

GetSetPercent returns the SetPercent field if non-nil, zero value otherwise.

### GetSetPercentOk

`func (o *EventThresholdByPercent) GetSetPercentOk() (*int64, bool)`

GetSetPercentOk returns a tuple with the SetPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetPercent

`func (o *EventThresholdByPercent) SetSetPercent(v int64)`

SetSetPercent sets SetPercent field to given value.

### HasSetPercent

`func (o *EventThresholdByPercent) HasSetPercent() bool`

HasSetPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


