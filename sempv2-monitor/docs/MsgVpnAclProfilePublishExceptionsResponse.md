# MsgVpnAclProfilePublishExceptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MsgVpnAclProfilePublishException**](MsgVpnAclProfilePublishException.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnAclProfilePublishExceptionLinks**](MsgVpnAclProfilePublishExceptionLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnAclProfilePublishExceptionsResponse

`func NewMsgVpnAclProfilePublishExceptionsResponse(meta SempMeta, ) *MsgVpnAclProfilePublishExceptionsResponse`

NewMsgVpnAclProfilePublishExceptionsResponse instantiates a new MsgVpnAclProfilePublishExceptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAclProfilePublishExceptionsResponseWithDefaults

`func NewMsgVpnAclProfilePublishExceptionsResponseWithDefaults() *MsgVpnAclProfilePublishExceptionsResponse`

NewMsgVpnAclProfilePublishExceptionsResponseWithDefaults instantiates a new MsgVpnAclProfilePublishExceptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnAclProfilePublishExceptionsResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnAclProfilePublishExceptionsResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnAclProfilePublishExceptionsResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnAclProfilePublishExceptionsResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnAclProfilePublishExceptionsResponse) GetData() []MsgVpnAclProfilePublishException`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnAclProfilePublishExceptionsResponse) GetDataOk() (*[]MsgVpnAclProfilePublishException, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnAclProfilePublishExceptionsResponse) SetData(v []MsgVpnAclProfilePublishException)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnAclProfilePublishExceptionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnAclProfilePublishExceptionsResponse) GetLinks() []MsgVpnAclProfilePublishExceptionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnAclProfilePublishExceptionsResponse) GetLinksOk() (*[]MsgVpnAclProfilePublishExceptionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnAclProfilePublishExceptionsResponse) SetLinks(v []MsgVpnAclProfilePublishExceptionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnAclProfilePublishExceptionsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnAclProfilePublishExceptionsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnAclProfilePublishExceptionsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnAclProfilePublishExceptionsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


