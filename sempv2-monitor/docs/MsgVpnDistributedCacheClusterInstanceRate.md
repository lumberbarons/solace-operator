# MsgVpnDistributedCacheClusterInstanceRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageDataRxBytePeakRate** | Pointer to **int64** | The peak of the one minute average of the data message rate received by the Cache Instance, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**AverageDataRxByteRate** | Pointer to **int64** | The one minute average of the data message rate received by the Cache Instance, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**AverageDataRxMsgPeakRate** | Pointer to **int64** | The peak of the one minute average of the data message rate received by the Cache Instance, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**AverageDataRxMsgRate** | Pointer to **int64** | The one minute average of the data message rate received by the Cache Instance, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**AverageDataTxMsgPeakRate** | Pointer to **int64** | The peak of the one minute average of the data message rate transmitted by the Cache Instance, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**AverageDataTxMsgRate** | Pointer to **int64** | The one minute average of the data message rate transmitted by the Cache Instance, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**AverageRequestRxPeakRate** | Pointer to **int64** | The peak of the one minute average of the request rate received by the Cache Instance, in requests per second (req/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**AverageRequestRxRate** | Pointer to **int64** | The one minute average of the request rate received by the Cache Instance, in requests per second (req/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**DataRxBytePeakRate** | Pointer to **int64** | The data message peak rate received by the Cache Instance, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**DataRxByteRate** | Pointer to **int64** | The data message rate received by the Cache Instance, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**DataRxMsgPeakRate** | Pointer to **int64** | The data message peak rate received by the Cache Instance, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**DataRxMsgRate** | Pointer to **int64** | The data message rate received by the Cache Instance, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**DataTxMsgPeakRate** | Pointer to **int64** | The data message peak rate transmitted by the Cache Instance, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**DataTxMsgRate** | Pointer to **int64** | The data message rate transmitted by the Cache Instance, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**RequestRxPeakRate** | Pointer to **int64** | The request peak rate received by the Cache Instance, in requests per second (req/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**RequestRxRate** | Pointer to **int64** | The request rate received by the Cache Instance, in requests per second (req/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 

## Methods

### NewMsgVpnDistributedCacheClusterInstanceRate

`func NewMsgVpnDistributedCacheClusterInstanceRate() *MsgVpnDistributedCacheClusterInstanceRate`

NewMsgVpnDistributedCacheClusterInstanceRate instantiates a new MsgVpnDistributedCacheClusterInstanceRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnDistributedCacheClusterInstanceRateWithDefaults

`func NewMsgVpnDistributedCacheClusterInstanceRateWithDefaults() *MsgVpnDistributedCacheClusterInstanceRate`

NewMsgVpnDistributedCacheClusterInstanceRateWithDefaults instantiates a new MsgVpnDistributedCacheClusterInstanceRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageDataRxBytePeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataRxBytePeakRate() int64`

GetAverageDataRxBytePeakRate returns the AverageDataRxBytePeakRate field if non-nil, zero value otherwise.

### GetAverageDataRxBytePeakRateOk

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataRxBytePeakRateOk() (*int64, bool)`

GetAverageDataRxBytePeakRateOk returns a tuple with the AverageDataRxBytePeakRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDataRxBytePeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) SetAverageDataRxBytePeakRate(v int64)`

SetAverageDataRxBytePeakRate sets AverageDataRxBytePeakRate field to given value.

### HasAverageDataRxBytePeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) HasAverageDataRxBytePeakRate() bool`

HasAverageDataRxBytePeakRate returns a boolean if a field has been set.

### GetAverageDataRxByteRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataRxByteRate() int64`

GetAverageDataRxByteRate returns the AverageDataRxByteRate field if non-nil, zero value otherwise.

### GetAverageDataRxByteRateOk

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataRxByteRateOk() (*int64, bool)`

GetAverageDataRxByteRateOk returns a tuple with the AverageDataRxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDataRxByteRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) SetAverageDataRxByteRate(v int64)`

SetAverageDataRxByteRate sets AverageDataRxByteRate field to given value.

### HasAverageDataRxByteRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) HasAverageDataRxByteRate() bool`

HasAverageDataRxByteRate returns a boolean if a field has been set.

### GetAverageDataRxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataRxMsgPeakRate() int64`

GetAverageDataRxMsgPeakRate returns the AverageDataRxMsgPeakRate field if non-nil, zero value otherwise.

### GetAverageDataRxMsgPeakRateOk

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataRxMsgPeakRateOk() (*int64, bool)`

GetAverageDataRxMsgPeakRateOk returns a tuple with the AverageDataRxMsgPeakRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDataRxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) SetAverageDataRxMsgPeakRate(v int64)`

SetAverageDataRxMsgPeakRate sets AverageDataRxMsgPeakRate field to given value.

### HasAverageDataRxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) HasAverageDataRxMsgPeakRate() bool`

HasAverageDataRxMsgPeakRate returns a boolean if a field has been set.

### GetAverageDataRxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataRxMsgRate() int64`

GetAverageDataRxMsgRate returns the AverageDataRxMsgRate field if non-nil, zero value otherwise.

### GetAverageDataRxMsgRateOk

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataRxMsgRateOk() (*int64, bool)`

GetAverageDataRxMsgRateOk returns a tuple with the AverageDataRxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDataRxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) SetAverageDataRxMsgRate(v int64)`

SetAverageDataRxMsgRate sets AverageDataRxMsgRate field to given value.

### HasAverageDataRxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) HasAverageDataRxMsgRate() bool`

HasAverageDataRxMsgRate returns a boolean if a field has been set.

### GetAverageDataTxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataTxMsgPeakRate() int64`

GetAverageDataTxMsgPeakRate returns the AverageDataTxMsgPeakRate field if non-nil, zero value otherwise.

### GetAverageDataTxMsgPeakRateOk

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataTxMsgPeakRateOk() (*int64, bool)`

GetAverageDataTxMsgPeakRateOk returns a tuple with the AverageDataTxMsgPeakRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDataTxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) SetAverageDataTxMsgPeakRate(v int64)`

SetAverageDataTxMsgPeakRate sets AverageDataTxMsgPeakRate field to given value.

### HasAverageDataTxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) HasAverageDataTxMsgPeakRate() bool`

HasAverageDataTxMsgPeakRate returns a boolean if a field has been set.

### GetAverageDataTxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataTxMsgRate() int64`

GetAverageDataTxMsgRate returns the AverageDataTxMsgRate field if non-nil, zero value otherwise.

### GetAverageDataTxMsgRateOk

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataTxMsgRateOk() (*int64, bool)`

GetAverageDataTxMsgRateOk returns a tuple with the AverageDataTxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDataTxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) SetAverageDataTxMsgRate(v int64)`

SetAverageDataTxMsgRate sets AverageDataTxMsgRate field to given value.

### HasAverageDataTxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) HasAverageDataTxMsgRate() bool`

HasAverageDataTxMsgRate returns a boolean if a field has been set.

### GetAverageRequestRxPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageRequestRxPeakRate() int64`

GetAverageRequestRxPeakRate returns the AverageRequestRxPeakRate field if non-nil, zero value otherwise.

### GetAverageRequestRxPeakRateOk

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageRequestRxPeakRateOk() (*int64, bool)`

GetAverageRequestRxPeakRateOk returns a tuple with the AverageRequestRxPeakRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRequestRxPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) SetAverageRequestRxPeakRate(v int64)`

SetAverageRequestRxPeakRate sets AverageRequestRxPeakRate field to given value.

### HasAverageRequestRxPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) HasAverageRequestRxPeakRate() bool`

HasAverageRequestRxPeakRate returns a boolean if a field has been set.

### GetAverageRequestRxRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageRequestRxRate() int64`

GetAverageRequestRxRate returns the AverageRequestRxRate field if non-nil, zero value otherwise.

### GetAverageRequestRxRateOk

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageRequestRxRateOk() (*int64, bool)`

GetAverageRequestRxRateOk returns a tuple with the AverageRequestRxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRequestRxRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) SetAverageRequestRxRate(v int64)`

SetAverageRequestRxRate sets AverageRequestRxRate field to given value.

### HasAverageRequestRxRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) HasAverageRequestRxRate() bool`

HasAverageRequestRxRate returns a boolean if a field has been set.

### GetDataRxBytePeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataRxBytePeakRate() int64`

GetDataRxBytePeakRate returns the DataRxBytePeakRate field if non-nil, zero value otherwise.

### GetDataRxBytePeakRateOk

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataRxBytePeakRateOk() (*int64, bool)`

GetDataRxBytePeakRateOk returns a tuple with the DataRxBytePeakRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxBytePeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) SetDataRxBytePeakRate(v int64)`

SetDataRxBytePeakRate sets DataRxBytePeakRate field to given value.

### HasDataRxBytePeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) HasDataRxBytePeakRate() bool`

HasDataRxBytePeakRate returns a boolean if a field has been set.

### GetDataRxByteRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataRxByteRate() int64`

GetDataRxByteRate returns the DataRxByteRate field if non-nil, zero value otherwise.

### GetDataRxByteRateOk

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataRxByteRateOk() (*int64, bool)`

GetDataRxByteRateOk returns a tuple with the DataRxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxByteRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) SetDataRxByteRate(v int64)`

SetDataRxByteRate sets DataRxByteRate field to given value.

### HasDataRxByteRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) HasDataRxByteRate() bool`

HasDataRxByteRate returns a boolean if a field has been set.

### GetDataRxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataRxMsgPeakRate() int64`

GetDataRxMsgPeakRate returns the DataRxMsgPeakRate field if non-nil, zero value otherwise.

### GetDataRxMsgPeakRateOk

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataRxMsgPeakRateOk() (*int64, bool)`

GetDataRxMsgPeakRateOk returns a tuple with the DataRxMsgPeakRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) SetDataRxMsgPeakRate(v int64)`

SetDataRxMsgPeakRate sets DataRxMsgPeakRate field to given value.

### HasDataRxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) HasDataRxMsgPeakRate() bool`

HasDataRxMsgPeakRate returns a boolean if a field has been set.

### GetDataRxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataRxMsgRate() int64`

GetDataRxMsgRate returns the DataRxMsgRate field if non-nil, zero value otherwise.

### GetDataRxMsgRateOk

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataRxMsgRateOk() (*int64, bool)`

GetDataRxMsgRateOk returns a tuple with the DataRxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) SetDataRxMsgRate(v int64)`

SetDataRxMsgRate sets DataRxMsgRate field to given value.

### HasDataRxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) HasDataRxMsgRate() bool`

HasDataRxMsgRate returns a boolean if a field has been set.

### GetDataTxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataTxMsgPeakRate() int64`

GetDataTxMsgPeakRate returns the DataTxMsgPeakRate field if non-nil, zero value otherwise.

### GetDataTxMsgPeakRateOk

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataTxMsgPeakRateOk() (*int64, bool)`

GetDataTxMsgPeakRateOk returns a tuple with the DataTxMsgPeakRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) SetDataTxMsgPeakRate(v int64)`

SetDataTxMsgPeakRate sets DataTxMsgPeakRate field to given value.

### HasDataTxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) HasDataTxMsgPeakRate() bool`

HasDataTxMsgPeakRate returns a boolean if a field has been set.

### GetDataTxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataTxMsgRate() int64`

GetDataTxMsgRate returns the DataTxMsgRate field if non-nil, zero value otherwise.

### GetDataTxMsgRateOk

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataTxMsgRateOk() (*int64, bool)`

GetDataTxMsgRateOk returns a tuple with the DataTxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) SetDataTxMsgRate(v int64)`

SetDataTxMsgRate sets DataTxMsgRate field to given value.

### HasDataTxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) HasDataTxMsgRate() bool`

HasDataTxMsgRate returns a boolean if a field has been set.

### GetRequestRxPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetRequestRxPeakRate() int64`

GetRequestRxPeakRate returns the RequestRxPeakRate field if non-nil, zero value otherwise.

### GetRequestRxPeakRateOk

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetRequestRxPeakRateOk() (*int64, bool)`

GetRequestRxPeakRateOk returns a tuple with the RequestRxPeakRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRxPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) SetRequestRxPeakRate(v int64)`

SetRequestRxPeakRate sets RequestRxPeakRate field to given value.

### HasRequestRxPeakRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) HasRequestRxPeakRate() bool`

HasRequestRxPeakRate returns a boolean if a field has been set.

### GetRequestRxRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetRequestRxRate() int64`

GetRequestRxRate returns the RequestRxRate field if non-nil, zero value otherwise.

### GetRequestRxRateOk

`func (o *MsgVpnDistributedCacheClusterInstanceRate) GetRequestRxRateOk() (*int64, bool)`

GetRequestRxRateOk returns a tuple with the RequestRxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRxRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) SetRequestRxRate(v int64)`

SetRequestRxRate sets RequestRxRate field to given value.

### HasRequestRxRate

`func (o *MsgVpnDistributedCacheClusterInstanceRate) HasRequestRxRate() bool`

HasRequestRxRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


