# EventThresholdByValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClearValue** | Pointer to **int64** | The clear threshold for the absolute value of this counter or rate. Falling below this value will trigger a corresponding event. | [optional] 
**SetValue** | Pointer to **int64** | The set threshold for the absolute value of this counter or rate. Exceeding this value will trigger a corresponding event. | [optional] 

## Methods

### NewEventThresholdByValue

`func NewEventThresholdByValue() *EventThresholdByValue`

NewEventThresholdByValue instantiates a new EventThresholdByValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventThresholdByValueWithDefaults

`func NewEventThresholdByValueWithDefaults() *EventThresholdByValue`

NewEventThresholdByValueWithDefaults instantiates a new EventThresholdByValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClearValue

`func (o *EventThresholdByValue) GetClearValue() int64`

GetClearValue returns the ClearValue field if non-nil, zero value otherwise.

### GetClearValueOk

`func (o *EventThresholdByValue) GetClearValueOk() (*int64, bool)`

GetClearValueOk returns a tuple with the ClearValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearValue

`func (o *EventThresholdByValue) SetClearValue(v int64)`

SetClearValue sets ClearValue field to given value.

### HasClearValue

`func (o *EventThresholdByValue) HasClearValue() bool`

HasClearValue returns a boolean if a field has been set.

### GetSetValue

`func (o *EventThresholdByValue) GetSetValue() int64`

GetSetValue returns the SetValue field if non-nil, zero value otherwise.

### GetSetValueOk

`func (o *EventThresholdByValue) GetSetValueOk() (*int64, bool)`

GetSetValueOk returns a tuple with the SetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetValue

`func (o *EventThresholdByValue) SetSetValue(v int64)`

SetSetValue sets SetValue field to given value.

### HasSetValue

`func (o *EventThresholdByValue) HasSetValue() bool`

HasSetValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


