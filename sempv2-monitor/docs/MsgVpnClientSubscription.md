# MsgVpnClientSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | Pointer to **string** | The name of the Client. | [optional] 
**DtoPriority** | Pointer to **string** | The priority of the Subscription topic for receiving deliver-to-one (DTO) messages. The allowed values and their meaning are:  &lt;pre&gt; \&quot;p1\&quot; - The 1st or highest priority. \&quot;p2\&quot; - The 2nd highest priority. \&quot;p3\&quot; - The 3rd highest priority. \&quot;p4\&quot; - The 4th highest priority. \&quot;da\&quot; - Ignore priority and deliver always. &lt;/pre&gt;  | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**SubscriptionTopic** | Pointer to **string** | The topic of the Subscription. | [optional] 

## Methods

### NewMsgVpnClientSubscription

`func NewMsgVpnClientSubscription() *MsgVpnClientSubscription`

NewMsgVpnClientSubscription instantiates a new MsgVpnClientSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnClientSubscriptionWithDefaults

`func NewMsgVpnClientSubscriptionWithDefaults() *MsgVpnClientSubscription`

NewMsgVpnClientSubscriptionWithDefaults instantiates a new MsgVpnClientSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *MsgVpnClientSubscription) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *MsgVpnClientSubscription) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *MsgVpnClientSubscription) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *MsgVpnClientSubscription) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetDtoPriority

`func (o *MsgVpnClientSubscription) GetDtoPriority() string`

GetDtoPriority returns the DtoPriority field if non-nil, zero value otherwise.

### GetDtoPriorityOk

`func (o *MsgVpnClientSubscription) GetDtoPriorityOk() (*string, bool)`

GetDtoPriorityOk returns a tuple with the DtoPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtoPriority

`func (o *MsgVpnClientSubscription) SetDtoPriority(v string)`

SetDtoPriority sets DtoPriority field to given value.

### HasDtoPriority

`func (o *MsgVpnClientSubscription) HasDtoPriority() bool`

HasDtoPriority returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnClientSubscription) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnClientSubscription) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnClientSubscription) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnClientSubscription) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetSubscriptionTopic

`func (o *MsgVpnClientSubscription) GetSubscriptionTopic() string`

GetSubscriptionTopic returns the SubscriptionTopic field if non-nil, zero value otherwise.

### GetSubscriptionTopicOk

`func (o *MsgVpnClientSubscription) GetSubscriptionTopicOk() (*string, bool)`

GetSubscriptionTopicOk returns a tuple with the SubscriptionTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTopic

`func (o *MsgVpnClientSubscription) SetSubscriptionTopic(v string)`

SetSubscriptionTopic sets SubscriptionTopic field to given value.

### HasSubscriptionTopic

`func (o *MsgVpnClientSubscription) HasSubscriptionTopic() bool`

HasSubscriptionTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


