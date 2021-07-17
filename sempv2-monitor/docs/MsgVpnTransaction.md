# MsgVpnTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **int32** | The identifier (ID) of the Client. | [optional] 
**ClientName** | Pointer to **string** | The name of the Client. | [optional] 
**ClientUsername** | Pointer to **string** | The username of the Client. | [optional] 
**IdleTimeout** | Pointer to **int32** | The number of seconds before an idle Transaction may be automatically rolled back and freed. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**Replicated** | Pointer to **bool** | Indicates whether the Transaction is replicated. | [optional] 
**SessionName** | Pointer to **string** | The name of the Transacted Session for the Transaction. | [optional] 
**State** | Pointer to **string** | The state of the Transaction. The allowed values and their meaning are:  &lt;pre&gt; \&quot;active\&quot; - The Transaction was started. \&quot;suspended\&quot; - The Transaction was suspended. \&quot;idle\&quot; - The Transaction was ended. \&quot;prepared\&quot; - The Transaction was prepared. \&quot;complete\&quot; - The Transaction was committed or rolled back. &lt;/pre&gt;  | [optional] 
**TimeInState** | Pointer to **int32** | The number of seconds the Transaction has remained in the current state. | [optional] 
**Type** | Pointer to **string** | The type of Transaction. The allowed values and their meaning are:  &lt;pre&gt; \&quot;xa\&quot; - The Transaction is an XA Transaction. \&quot;local\&quot; - The Transaction is a local Transaction. &lt;/pre&gt;  | [optional] 
**Xid** | Pointer to **string** | The identifier (ID) of the Transaction. | [optional] 

## Methods

### NewMsgVpnTransaction

`func NewMsgVpnTransaction() *MsgVpnTransaction`

NewMsgVpnTransaction instantiates a new MsgVpnTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnTransactionWithDefaults

`func NewMsgVpnTransactionWithDefaults() *MsgVpnTransaction`

NewMsgVpnTransactionWithDefaults instantiates a new MsgVpnTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *MsgVpnTransaction) GetClientId() int32`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MsgVpnTransaction) GetClientIdOk() (*int32, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MsgVpnTransaction) SetClientId(v int32)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *MsgVpnTransaction) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientName

`func (o *MsgVpnTransaction) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *MsgVpnTransaction) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *MsgVpnTransaction) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *MsgVpnTransaction) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientUsername

`func (o *MsgVpnTransaction) GetClientUsername() string`

GetClientUsername returns the ClientUsername field if non-nil, zero value otherwise.

### GetClientUsernameOk

`func (o *MsgVpnTransaction) GetClientUsernameOk() (*string, bool)`

GetClientUsernameOk returns a tuple with the ClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUsername

`func (o *MsgVpnTransaction) SetClientUsername(v string)`

SetClientUsername sets ClientUsername field to given value.

### HasClientUsername

`func (o *MsgVpnTransaction) HasClientUsername() bool`

HasClientUsername returns a boolean if a field has been set.

### GetIdleTimeout

`func (o *MsgVpnTransaction) GetIdleTimeout() int32`

GetIdleTimeout returns the IdleTimeout field if non-nil, zero value otherwise.

### GetIdleTimeoutOk

`func (o *MsgVpnTransaction) GetIdleTimeoutOk() (*int32, bool)`

GetIdleTimeoutOk returns a tuple with the IdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeout

`func (o *MsgVpnTransaction) SetIdleTimeout(v int32)`

SetIdleTimeout sets IdleTimeout field to given value.

### HasIdleTimeout

`func (o *MsgVpnTransaction) HasIdleTimeout() bool`

HasIdleTimeout returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnTransaction) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnTransaction) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnTransaction) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnTransaction) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetReplicated

`func (o *MsgVpnTransaction) GetReplicated() bool`

GetReplicated returns the Replicated field if non-nil, zero value otherwise.

### GetReplicatedOk

`func (o *MsgVpnTransaction) GetReplicatedOk() (*bool, bool)`

GetReplicatedOk returns a tuple with the Replicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicated

`func (o *MsgVpnTransaction) SetReplicated(v bool)`

SetReplicated sets Replicated field to given value.

### HasReplicated

`func (o *MsgVpnTransaction) HasReplicated() bool`

HasReplicated returns a boolean if a field has been set.

### GetSessionName

`func (o *MsgVpnTransaction) GetSessionName() string`

GetSessionName returns the SessionName field if non-nil, zero value otherwise.

### GetSessionNameOk

`func (o *MsgVpnTransaction) GetSessionNameOk() (*string, bool)`

GetSessionNameOk returns a tuple with the SessionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionName

`func (o *MsgVpnTransaction) SetSessionName(v string)`

SetSessionName sets SessionName field to given value.

### HasSessionName

`func (o *MsgVpnTransaction) HasSessionName() bool`

HasSessionName returns a boolean if a field has been set.

### GetState

`func (o *MsgVpnTransaction) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MsgVpnTransaction) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MsgVpnTransaction) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MsgVpnTransaction) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTimeInState

`func (o *MsgVpnTransaction) GetTimeInState() int32`

GetTimeInState returns the TimeInState field if non-nil, zero value otherwise.

### GetTimeInStateOk

`func (o *MsgVpnTransaction) GetTimeInStateOk() (*int32, bool)`

GetTimeInStateOk returns a tuple with the TimeInState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInState

`func (o *MsgVpnTransaction) SetTimeInState(v int32)`

SetTimeInState sets TimeInState field to given value.

### HasTimeInState

`func (o *MsgVpnTransaction) HasTimeInState() bool`

HasTimeInState returns a boolean if a field has been set.

### GetType

`func (o *MsgVpnTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MsgVpnTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MsgVpnTransaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MsgVpnTransaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetXid

`func (o *MsgVpnTransaction) GetXid() string`

GetXid returns the Xid field if non-nil, zero value otherwise.

### GetXidOk

`func (o *MsgVpnTransaction) GetXidOk() (*string, bool)`

GetXidOk returns a tuple with the Xid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXid

`func (o *MsgVpnTransaction) SetXid(v string)`

SetXid sets Xid field to given value.

### HasXid

`func (o *MsgVpnTransaction) HasXid() bool`

HasXid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


