# AboutApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AboutApi**](AboutApi.md) |  | [optional] 
**Links** | Pointer to [**AboutApiLinks**](AboutApiLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewAboutApiResponse

`func NewAboutApiResponse(meta SempMeta, ) *AboutApiResponse`

NewAboutApiResponse instantiates a new AboutApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAboutApiResponseWithDefaults

`func NewAboutApiResponseWithDefaults() *AboutApiResponse`

NewAboutApiResponseWithDefaults instantiates a new AboutApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AboutApiResponse) GetData() AboutApi`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AboutApiResponse) GetDataOk() (*AboutApi, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AboutApiResponse) SetData(v AboutApi)`

SetData sets Data field to given value.

### HasData

`func (o *AboutApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *AboutApiResponse) GetLinks() AboutApiLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AboutApiResponse) GetLinksOk() (*AboutApiLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AboutApiResponse) SetLinks(v AboutApiLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AboutApiResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *AboutApiResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AboutApiResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AboutApiResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


