# MsgVpnAclProfileLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientConnectExceptionsUri** | Pointer to **string** | The URI of this ACL Profile&#39;s collection of Client Connect Exception objects. | [optional] 
**PublishExceptionsUri** | Pointer to **string** | The URI of this ACL Profile&#39;s collection of Publish Topic Exception objects. Deprecated since 2.14. Replaced by publishTopicExceptions. | [optional] 
**PublishTopicExceptionsUri** | Pointer to **string** | The URI of this ACL Profile&#39;s collection of Publish Topic Exception objects. Available since 2.14. | [optional] 
**SubscribeExceptionsUri** | Pointer to **string** | The URI of this ACL Profile&#39;s collection of Subscribe Topic Exception objects. Deprecated since 2.14. Replaced by subscribeTopicExceptions. | [optional] 
**SubscribeShareNameExceptionsUri** | Pointer to **string** | The URI of this ACL Profile&#39;s collection of Subscribe Share Name Exception objects. Available since 2.14. | [optional] 
**SubscribeTopicExceptionsUri** | Pointer to **string** | The URI of this ACL Profile&#39;s collection of Subscribe Topic Exception objects. Available since 2.14. | [optional] 
**Uri** | Pointer to **string** | The URI of this ACL Profile object. | [optional] 

## Methods

### NewMsgVpnAclProfileLinks

`func NewMsgVpnAclProfileLinks() *MsgVpnAclProfileLinks`

NewMsgVpnAclProfileLinks instantiates a new MsgVpnAclProfileLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAclProfileLinksWithDefaults

`func NewMsgVpnAclProfileLinksWithDefaults() *MsgVpnAclProfileLinks`

NewMsgVpnAclProfileLinksWithDefaults instantiates a new MsgVpnAclProfileLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientConnectExceptionsUri

`func (o *MsgVpnAclProfileLinks) GetClientConnectExceptionsUri() string`

GetClientConnectExceptionsUri returns the ClientConnectExceptionsUri field if non-nil, zero value otherwise.

### GetClientConnectExceptionsUriOk

`func (o *MsgVpnAclProfileLinks) GetClientConnectExceptionsUriOk() (*string, bool)`

GetClientConnectExceptionsUriOk returns a tuple with the ClientConnectExceptionsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectExceptionsUri

`func (o *MsgVpnAclProfileLinks) SetClientConnectExceptionsUri(v string)`

SetClientConnectExceptionsUri sets ClientConnectExceptionsUri field to given value.

### HasClientConnectExceptionsUri

`func (o *MsgVpnAclProfileLinks) HasClientConnectExceptionsUri() bool`

HasClientConnectExceptionsUri returns a boolean if a field has been set.

### GetPublishExceptionsUri

`func (o *MsgVpnAclProfileLinks) GetPublishExceptionsUri() string`

GetPublishExceptionsUri returns the PublishExceptionsUri field if non-nil, zero value otherwise.

### GetPublishExceptionsUriOk

`func (o *MsgVpnAclProfileLinks) GetPublishExceptionsUriOk() (*string, bool)`

GetPublishExceptionsUriOk returns a tuple with the PublishExceptionsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishExceptionsUri

`func (o *MsgVpnAclProfileLinks) SetPublishExceptionsUri(v string)`

SetPublishExceptionsUri sets PublishExceptionsUri field to given value.

### HasPublishExceptionsUri

`func (o *MsgVpnAclProfileLinks) HasPublishExceptionsUri() bool`

HasPublishExceptionsUri returns a boolean if a field has been set.

### GetPublishTopicExceptionsUri

`func (o *MsgVpnAclProfileLinks) GetPublishTopicExceptionsUri() string`

GetPublishTopicExceptionsUri returns the PublishTopicExceptionsUri field if non-nil, zero value otherwise.

### GetPublishTopicExceptionsUriOk

`func (o *MsgVpnAclProfileLinks) GetPublishTopicExceptionsUriOk() (*string, bool)`

GetPublishTopicExceptionsUriOk returns a tuple with the PublishTopicExceptionsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishTopicExceptionsUri

`func (o *MsgVpnAclProfileLinks) SetPublishTopicExceptionsUri(v string)`

SetPublishTopicExceptionsUri sets PublishTopicExceptionsUri field to given value.

### HasPublishTopicExceptionsUri

`func (o *MsgVpnAclProfileLinks) HasPublishTopicExceptionsUri() bool`

HasPublishTopicExceptionsUri returns a boolean if a field has been set.

### GetSubscribeExceptionsUri

`func (o *MsgVpnAclProfileLinks) GetSubscribeExceptionsUri() string`

GetSubscribeExceptionsUri returns the SubscribeExceptionsUri field if non-nil, zero value otherwise.

### GetSubscribeExceptionsUriOk

`func (o *MsgVpnAclProfileLinks) GetSubscribeExceptionsUriOk() (*string, bool)`

GetSubscribeExceptionsUriOk returns a tuple with the SubscribeExceptionsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeExceptionsUri

`func (o *MsgVpnAclProfileLinks) SetSubscribeExceptionsUri(v string)`

SetSubscribeExceptionsUri sets SubscribeExceptionsUri field to given value.

### HasSubscribeExceptionsUri

`func (o *MsgVpnAclProfileLinks) HasSubscribeExceptionsUri() bool`

HasSubscribeExceptionsUri returns a boolean if a field has been set.

### GetSubscribeShareNameExceptionsUri

`func (o *MsgVpnAclProfileLinks) GetSubscribeShareNameExceptionsUri() string`

GetSubscribeShareNameExceptionsUri returns the SubscribeShareNameExceptionsUri field if non-nil, zero value otherwise.

### GetSubscribeShareNameExceptionsUriOk

`func (o *MsgVpnAclProfileLinks) GetSubscribeShareNameExceptionsUriOk() (*string, bool)`

GetSubscribeShareNameExceptionsUriOk returns a tuple with the SubscribeShareNameExceptionsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeShareNameExceptionsUri

`func (o *MsgVpnAclProfileLinks) SetSubscribeShareNameExceptionsUri(v string)`

SetSubscribeShareNameExceptionsUri sets SubscribeShareNameExceptionsUri field to given value.

### HasSubscribeShareNameExceptionsUri

`func (o *MsgVpnAclProfileLinks) HasSubscribeShareNameExceptionsUri() bool`

HasSubscribeShareNameExceptionsUri returns a boolean if a field has been set.

### GetSubscribeTopicExceptionsUri

`func (o *MsgVpnAclProfileLinks) GetSubscribeTopicExceptionsUri() string`

GetSubscribeTopicExceptionsUri returns the SubscribeTopicExceptionsUri field if non-nil, zero value otherwise.

### GetSubscribeTopicExceptionsUriOk

`func (o *MsgVpnAclProfileLinks) GetSubscribeTopicExceptionsUriOk() (*string, bool)`

GetSubscribeTopicExceptionsUriOk returns a tuple with the SubscribeTopicExceptionsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeTopicExceptionsUri

`func (o *MsgVpnAclProfileLinks) SetSubscribeTopicExceptionsUri(v string)`

SetSubscribeTopicExceptionsUri sets SubscribeTopicExceptionsUri field to given value.

### HasSubscribeTopicExceptionsUri

`func (o *MsgVpnAclProfileLinks) HasSubscribeTopicExceptionsUri() bool`

HasSubscribeTopicExceptionsUri returns a boolean if a field has been set.

### GetUri

`func (o *MsgVpnAclProfileLinks) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *MsgVpnAclProfileLinks) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *MsgVpnAclProfileLinks) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *MsgVpnAclProfileLinks) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


