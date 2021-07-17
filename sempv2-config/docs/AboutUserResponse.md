# AboutUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AboutUser**](AboutUser.md) |  | [optional] 
**Links** | Pointer to [**AboutUserLinks**](AboutUserLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewAboutUserResponse

`func NewAboutUserResponse(meta SempMeta, ) *AboutUserResponse`

NewAboutUserResponse instantiates a new AboutUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAboutUserResponseWithDefaults

`func NewAboutUserResponseWithDefaults() *AboutUserResponse`

NewAboutUserResponseWithDefaults instantiates a new AboutUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AboutUserResponse) GetData() AboutUser`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AboutUserResponse) GetDataOk() (*AboutUser, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AboutUserResponse) SetData(v AboutUser)`

SetData sets Data field to given value.

### HasData

`func (o *AboutUserResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *AboutUserResponse) GetLinks() AboutUserLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AboutUserResponse) GetLinksOk() (*AboutUserLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AboutUserResponse) SetLinks(v AboutUserLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AboutUserResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *AboutUserResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AboutUserResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AboutUserResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


