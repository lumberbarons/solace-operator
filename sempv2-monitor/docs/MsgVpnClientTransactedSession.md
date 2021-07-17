# MsgVpnClientTransactedSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | Pointer to **string** | The name of the Client. | [optional] 
**CommitCount** | Pointer to **int64** | The number of transactions committed within the Transacted Session. | [optional] 
**CommitFailureCount** | Pointer to **int64** | The number of transaction commit operations that failed. | [optional] 
**CommitSuccessCount** | Pointer to **int64** | The number of transaction commit operations that succeeded. | [optional] 
**ConsumedMsgCount** | Pointer to **int64** | The number of messages consumed within the Transacted Session. | [optional] 
**EndFailFailureCount** | Pointer to **int64** | The number of transaction end fail operations that failed. | [optional] 
**EndFailSuccessCount** | Pointer to **int64** | The number of transaction end fail operations that succeeded. | [optional] 
**EndFailureCount** | Pointer to **int64** | The number of transaction end operations that failed. | [optional] 
**EndRollbackFailureCount** | Pointer to **int64** | The number of transaction end rollback operations that failed. | [optional] 
**EndRollbackSuccessCount** | Pointer to **int64** | The number of transaction end rollback operations that succeeded. | [optional] 
**EndSuccessCount** | Pointer to **int64** | The number of transaction end operations that succeeded. | [optional] 
**FailureCount** | Pointer to **int64** | The number of transactions that failed within the Transacted Session. | [optional] 
**ForgetFailureCount** | Pointer to **int64** | The number of transaction forget operations that failed. | [optional] 
**ForgetSuccessCount** | Pointer to **int64** | The number of transaction forget operations that succeeded. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**OnePhaseCommitFailureCount** | Pointer to **int64** | The number of transaction one-phase commit operations that failed. | [optional] 
**OnePhaseCommitSuccessCount** | Pointer to **int64** | The number of transaction one-phase commit operations that succeeded. | [optional] 
**PendingConsumedMsgCount** | Pointer to **int32** | The number of messages to be consumed when the transaction is committed. | [optional] 
**PendingPublishedMsgCount** | Pointer to **int32** | The number of messages to be published when the transaction is committed. | [optional] 
**PrepareFailureCount** | Pointer to **int64** | The number of transaction prepare operations that failed. | [optional] 
**PrepareSuccessCount** | Pointer to **int64** | The number of transaction prepare operations that succeeded. | [optional] 
**PreviousTransactionState** | Pointer to **string** | The state of the previous transaction. The allowed values and their meaning are:  &lt;pre&gt; \&quot;none\&quot; - The previous transaction had no state. \&quot;committed\&quot; - The previous transaction was committed. \&quot;rolled-back\&quot; - The previous transaction was rolled back. \&quot;failed\&quot; - The previous transaction failed. &lt;/pre&gt;  | [optional] 
**PublishedMsgCount** | Pointer to **int64** | The number of messages published within the Transacted Session. | [optional] 
**ResumeFailureCount** | Pointer to **int64** | The number of transaction resume operations that failed. | [optional] 
**ResumeSuccessCount** | Pointer to **int64** | The number of transaction resume operations that succeeded. | [optional] 
**RetrievedMsgCount** | Pointer to **int64** | The number of messages retrieved within the Transacted Session. | [optional] 
**RollbackCount** | Pointer to **int64** | The number of transactions rolled back within the Transacted Session. | [optional] 
**RollbackFailureCount** | Pointer to **int64** | The number of transaction rollback operations that failed. | [optional] 
**RollbackSuccessCount** | Pointer to **int64** | The number of transaction rollback operations that succeeded. | [optional] 
**SessionName** | Pointer to **string** | The name of the Transacted Session. | [optional] 
**SpooledMsgCount** | Pointer to **int64** | The number of messages spooled within the Transacted Session. | [optional] 
**StartFailureCount** | Pointer to **int64** | The number of transaction start operations that failed. | [optional] 
**StartSuccessCount** | Pointer to **int64** | The number of transaction start operations that succeeded. | [optional] 
**SuccessCount** | Pointer to **int64** | The number of transactions that succeeded within the Transacted Session. | [optional] 
**SuspendFailureCount** | Pointer to **int64** | The number of transaction suspend operations that failed. | [optional] 
**SuspendSuccessCount** | Pointer to **int64** | The number of transaction suspend operations that succeeded. | [optional] 
**TransactionId** | Pointer to **int32** | The identifier (ID) of the transaction in the Transacted Session. | [optional] 
**TransactionState** | Pointer to **string** | The state of the current transaction. The allowed values and their meaning are:  &lt;pre&gt; \&quot;in-progress\&quot; - The current transaction is in progress. \&quot;committing\&quot; - The current transaction is committing. \&quot;rolling-back\&quot; - The current transaction is rolling back. \&quot;failing\&quot; - The current transaction is failing. &lt;/pre&gt;  | [optional] 
**TwoPhaseCommitFailureCount** | Pointer to **int64** | The number of transaction two-phase commit operations that failed. | [optional] 
**TwoPhaseCommitSuccessCount** | Pointer to **int64** | The number of transaction two-phase commit operations that succeeded. | [optional] 

## Methods

### NewMsgVpnClientTransactedSession

`func NewMsgVpnClientTransactedSession() *MsgVpnClientTransactedSession`

NewMsgVpnClientTransactedSession instantiates a new MsgVpnClientTransactedSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnClientTransactedSessionWithDefaults

`func NewMsgVpnClientTransactedSessionWithDefaults() *MsgVpnClientTransactedSession`

NewMsgVpnClientTransactedSessionWithDefaults instantiates a new MsgVpnClientTransactedSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *MsgVpnClientTransactedSession) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *MsgVpnClientTransactedSession) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *MsgVpnClientTransactedSession) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *MsgVpnClientTransactedSession) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetCommitCount

`func (o *MsgVpnClientTransactedSession) GetCommitCount() int64`

GetCommitCount returns the CommitCount field if non-nil, zero value otherwise.

### GetCommitCountOk

`func (o *MsgVpnClientTransactedSession) GetCommitCountOk() (*int64, bool)`

GetCommitCountOk returns a tuple with the CommitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitCount

`func (o *MsgVpnClientTransactedSession) SetCommitCount(v int64)`

SetCommitCount sets CommitCount field to given value.

### HasCommitCount

`func (o *MsgVpnClientTransactedSession) HasCommitCount() bool`

HasCommitCount returns a boolean if a field has been set.

### GetCommitFailureCount

`func (o *MsgVpnClientTransactedSession) GetCommitFailureCount() int64`

GetCommitFailureCount returns the CommitFailureCount field if non-nil, zero value otherwise.

### GetCommitFailureCountOk

`func (o *MsgVpnClientTransactedSession) GetCommitFailureCountOk() (*int64, bool)`

GetCommitFailureCountOk returns a tuple with the CommitFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitFailureCount

`func (o *MsgVpnClientTransactedSession) SetCommitFailureCount(v int64)`

SetCommitFailureCount sets CommitFailureCount field to given value.

### HasCommitFailureCount

`func (o *MsgVpnClientTransactedSession) HasCommitFailureCount() bool`

HasCommitFailureCount returns a boolean if a field has been set.

### GetCommitSuccessCount

`func (o *MsgVpnClientTransactedSession) GetCommitSuccessCount() int64`

GetCommitSuccessCount returns the CommitSuccessCount field if non-nil, zero value otherwise.

### GetCommitSuccessCountOk

`func (o *MsgVpnClientTransactedSession) GetCommitSuccessCountOk() (*int64, bool)`

GetCommitSuccessCountOk returns a tuple with the CommitSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSuccessCount

`func (o *MsgVpnClientTransactedSession) SetCommitSuccessCount(v int64)`

SetCommitSuccessCount sets CommitSuccessCount field to given value.

### HasCommitSuccessCount

`func (o *MsgVpnClientTransactedSession) HasCommitSuccessCount() bool`

HasCommitSuccessCount returns a boolean if a field has been set.

### GetConsumedMsgCount

`func (o *MsgVpnClientTransactedSession) GetConsumedMsgCount() int64`

GetConsumedMsgCount returns the ConsumedMsgCount field if non-nil, zero value otherwise.

### GetConsumedMsgCountOk

`func (o *MsgVpnClientTransactedSession) GetConsumedMsgCountOk() (*int64, bool)`

GetConsumedMsgCountOk returns a tuple with the ConsumedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedMsgCount

`func (o *MsgVpnClientTransactedSession) SetConsumedMsgCount(v int64)`

SetConsumedMsgCount sets ConsumedMsgCount field to given value.

### HasConsumedMsgCount

`func (o *MsgVpnClientTransactedSession) HasConsumedMsgCount() bool`

HasConsumedMsgCount returns a boolean if a field has been set.

### GetEndFailFailureCount

`func (o *MsgVpnClientTransactedSession) GetEndFailFailureCount() int64`

GetEndFailFailureCount returns the EndFailFailureCount field if non-nil, zero value otherwise.

### GetEndFailFailureCountOk

`func (o *MsgVpnClientTransactedSession) GetEndFailFailureCountOk() (*int64, bool)`

GetEndFailFailureCountOk returns a tuple with the EndFailFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndFailFailureCount

`func (o *MsgVpnClientTransactedSession) SetEndFailFailureCount(v int64)`

SetEndFailFailureCount sets EndFailFailureCount field to given value.

### HasEndFailFailureCount

`func (o *MsgVpnClientTransactedSession) HasEndFailFailureCount() bool`

HasEndFailFailureCount returns a boolean if a field has been set.

### GetEndFailSuccessCount

`func (o *MsgVpnClientTransactedSession) GetEndFailSuccessCount() int64`

GetEndFailSuccessCount returns the EndFailSuccessCount field if non-nil, zero value otherwise.

### GetEndFailSuccessCountOk

`func (o *MsgVpnClientTransactedSession) GetEndFailSuccessCountOk() (*int64, bool)`

GetEndFailSuccessCountOk returns a tuple with the EndFailSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndFailSuccessCount

`func (o *MsgVpnClientTransactedSession) SetEndFailSuccessCount(v int64)`

SetEndFailSuccessCount sets EndFailSuccessCount field to given value.

### HasEndFailSuccessCount

`func (o *MsgVpnClientTransactedSession) HasEndFailSuccessCount() bool`

HasEndFailSuccessCount returns a boolean if a field has been set.

### GetEndFailureCount

`func (o *MsgVpnClientTransactedSession) GetEndFailureCount() int64`

GetEndFailureCount returns the EndFailureCount field if non-nil, zero value otherwise.

### GetEndFailureCountOk

`func (o *MsgVpnClientTransactedSession) GetEndFailureCountOk() (*int64, bool)`

GetEndFailureCountOk returns a tuple with the EndFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndFailureCount

`func (o *MsgVpnClientTransactedSession) SetEndFailureCount(v int64)`

SetEndFailureCount sets EndFailureCount field to given value.

### HasEndFailureCount

`func (o *MsgVpnClientTransactedSession) HasEndFailureCount() bool`

HasEndFailureCount returns a boolean if a field has been set.

### GetEndRollbackFailureCount

`func (o *MsgVpnClientTransactedSession) GetEndRollbackFailureCount() int64`

GetEndRollbackFailureCount returns the EndRollbackFailureCount field if non-nil, zero value otherwise.

### GetEndRollbackFailureCountOk

`func (o *MsgVpnClientTransactedSession) GetEndRollbackFailureCountOk() (*int64, bool)`

GetEndRollbackFailureCountOk returns a tuple with the EndRollbackFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndRollbackFailureCount

`func (o *MsgVpnClientTransactedSession) SetEndRollbackFailureCount(v int64)`

SetEndRollbackFailureCount sets EndRollbackFailureCount field to given value.

### HasEndRollbackFailureCount

`func (o *MsgVpnClientTransactedSession) HasEndRollbackFailureCount() bool`

HasEndRollbackFailureCount returns a boolean if a field has been set.

### GetEndRollbackSuccessCount

`func (o *MsgVpnClientTransactedSession) GetEndRollbackSuccessCount() int64`

GetEndRollbackSuccessCount returns the EndRollbackSuccessCount field if non-nil, zero value otherwise.

### GetEndRollbackSuccessCountOk

`func (o *MsgVpnClientTransactedSession) GetEndRollbackSuccessCountOk() (*int64, bool)`

GetEndRollbackSuccessCountOk returns a tuple with the EndRollbackSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndRollbackSuccessCount

`func (o *MsgVpnClientTransactedSession) SetEndRollbackSuccessCount(v int64)`

SetEndRollbackSuccessCount sets EndRollbackSuccessCount field to given value.

### HasEndRollbackSuccessCount

`func (o *MsgVpnClientTransactedSession) HasEndRollbackSuccessCount() bool`

HasEndRollbackSuccessCount returns a boolean if a field has been set.

### GetEndSuccessCount

`func (o *MsgVpnClientTransactedSession) GetEndSuccessCount() int64`

GetEndSuccessCount returns the EndSuccessCount field if non-nil, zero value otherwise.

### GetEndSuccessCountOk

`func (o *MsgVpnClientTransactedSession) GetEndSuccessCountOk() (*int64, bool)`

GetEndSuccessCountOk returns a tuple with the EndSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSuccessCount

`func (o *MsgVpnClientTransactedSession) SetEndSuccessCount(v int64)`

SetEndSuccessCount sets EndSuccessCount field to given value.

### HasEndSuccessCount

`func (o *MsgVpnClientTransactedSession) HasEndSuccessCount() bool`

HasEndSuccessCount returns a boolean if a field has been set.

### GetFailureCount

`func (o *MsgVpnClientTransactedSession) GetFailureCount() int64`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *MsgVpnClientTransactedSession) GetFailureCountOk() (*int64, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *MsgVpnClientTransactedSession) SetFailureCount(v int64)`

SetFailureCount sets FailureCount field to given value.

### HasFailureCount

`func (o *MsgVpnClientTransactedSession) HasFailureCount() bool`

HasFailureCount returns a boolean if a field has been set.

### GetForgetFailureCount

`func (o *MsgVpnClientTransactedSession) GetForgetFailureCount() int64`

GetForgetFailureCount returns the ForgetFailureCount field if non-nil, zero value otherwise.

### GetForgetFailureCountOk

`func (o *MsgVpnClientTransactedSession) GetForgetFailureCountOk() (*int64, bool)`

GetForgetFailureCountOk returns a tuple with the ForgetFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgetFailureCount

`func (o *MsgVpnClientTransactedSession) SetForgetFailureCount(v int64)`

SetForgetFailureCount sets ForgetFailureCount field to given value.

### HasForgetFailureCount

`func (o *MsgVpnClientTransactedSession) HasForgetFailureCount() bool`

HasForgetFailureCount returns a boolean if a field has been set.

### GetForgetSuccessCount

`func (o *MsgVpnClientTransactedSession) GetForgetSuccessCount() int64`

GetForgetSuccessCount returns the ForgetSuccessCount field if non-nil, zero value otherwise.

### GetForgetSuccessCountOk

`func (o *MsgVpnClientTransactedSession) GetForgetSuccessCountOk() (*int64, bool)`

GetForgetSuccessCountOk returns a tuple with the ForgetSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgetSuccessCount

`func (o *MsgVpnClientTransactedSession) SetForgetSuccessCount(v int64)`

SetForgetSuccessCount sets ForgetSuccessCount field to given value.

### HasForgetSuccessCount

`func (o *MsgVpnClientTransactedSession) HasForgetSuccessCount() bool`

HasForgetSuccessCount returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnClientTransactedSession) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnClientTransactedSession) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnClientTransactedSession) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnClientTransactedSession) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetOnePhaseCommitFailureCount

`func (o *MsgVpnClientTransactedSession) GetOnePhaseCommitFailureCount() int64`

GetOnePhaseCommitFailureCount returns the OnePhaseCommitFailureCount field if non-nil, zero value otherwise.

### GetOnePhaseCommitFailureCountOk

`func (o *MsgVpnClientTransactedSession) GetOnePhaseCommitFailureCountOk() (*int64, bool)`

GetOnePhaseCommitFailureCountOk returns a tuple with the OnePhaseCommitFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnePhaseCommitFailureCount

`func (o *MsgVpnClientTransactedSession) SetOnePhaseCommitFailureCount(v int64)`

SetOnePhaseCommitFailureCount sets OnePhaseCommitFailureCount field to given value.

### HasOnePhaseCommitFailureCount

`func (o *MsgVpnClientTransactedSession) HasOnePhaseCommitFailureCount() bool`

HasOnePhaseCommitFailureCount returns a boolean if a field has been set.

### GetOnePhaseCommitSuccessCount

`func (o *MsgVpnClientTransactedSession) GetOnePhaseCommitSuccessCount() int64`

GetOnePhaseCommitSuccessCount returns the OnePhaseCommitSuccessCount field if non-nil, zero value otherwise.

### GetOnePhaseCommitSuccessCountOk

`func (o *MsgVpnClientTransactedSession) GetOnePhaseCommitSuccessCountOk() (*int64, bool)`

GetOnePhaseCommitSuccessCountOk returns a tuple with the OnePhaseCommitSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnePhaseCommitSuccessCount

`func (o *MsgVpnClientTransactedSession) SetOnePhaseCommitSuccessCount(v int64)`

SetOnePhaseCommitSuccessCount sets OnePhaseCommitSuccessCount field to given value.

### HasOnePhaseCommitSuccessCount

`func (o *MsgVpnClientTransactedSession) HasOnePhaseCommitSuccessCount() bool`

HasOnePhaseCommitSuccessCount returns a boolean if a field has been set.

### GetPendingConsumedMsgCount

`func (o *MsgVpnClientTransactedSession) GetPendingConsumedMsgCount() int32`

GetPendingConsumedMsgCount returns the PendingConsumedMsgCount field if non-nil, zero value otherwise.

### GetPendingConsumedMsgCountOk

`func (o *MsgVpnClientTransactedSession) GetPendingConsumedMsgCountOk() (*int32, bool)`

GetPendingConsumedMsgCountOk returns a tuple with the PendingConsumedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingConsumedMsgCount

`func (o *MsgVpnClientTransactedSession) SetPendingConsumedMsgCount(v int32)`

SetPendingConsumedMsgCount sets PendingConsumedMsgCount field to given value.

### HasPendingConsumedMsgCount

`func (o *MsgVpnClientTransactedSession) HasPendingConsumedMsgCount() bool`

HasPendingConsumedMsgCount returns a boolean if a field has been set.

### GetPendingPublishedMsgCount

`func (o *MsgVpnClientTransactedSession) GetPendingPublishedMsgCount() int32`

GetPendingPublishedMsgCount returns the PendingPublishedMsgCount field if non-nil, zero value otherwise.

### GetPendingPublishedMsgCountOk

`func (o *MsgVpnClientTransactedSession) GetPendingPublishedMsgCountOk() (*int32, bool)`

GetPendingPublishedMsgCountOk returns a tuple with the PendingPublishedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingPublishedMsgCount

`func (o *MsgVpnClientTransactedSession) SetPendingPublishedMsgCount(v int32)`

SetPendingPublishedMsgCount sets PendingPublishedMsgCount field to given value.

### HasPendingPublishedMsgCount

`func (o *MsgVpnClientTransactedSession) HasPendingPublishedMsgCount() bool`

HasPendingPublishedMsgCount returns a boolean if a field has been set.

### GetPrepareFailureCount

`func (o *MsgVpnClientTransactedSession) GetPrepareFailureCount() int64`

GetPrepareFailureCount returns the PrepareFailureCount field if non-nil, zero value otherwise.

### GetPrepareFailureCountOk

`func (o *MsgVpnClientTransactedSession) GetPrepareFailureCountOk() (*int64, bool)`

GetPrepareFailureCountOk returns a tuple with the PrepareFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepareFailureCount

`func (o *MsgVpnClientTransactedSession) SetPrepareFailureCount(v int64)`

SetPrepareFailureCount sets PrepareFailureCount field to given value.

### HasPrepareFailureCount

`func (o *MsgVpnClientTransactedSession) HasPrepareFailureCount() bool`

HasPrepareFailureCount returns a boolean if a field has been set.

### GetPrepareSuccessCount

`func (o *MsgVpnClientTransactedSession) GetPrepareSuccessCount() int64`

GetPrepareSuccessCount returns the PrepareSuccessCount field if non-nil, zero value otherwise.

### GetPrepareSuccessCountOk

`func (o *MsgVpnClientTransactedSession) GetPrepareSuccessCountOk() (*int64, bool)`

GetPrepareSuccessCountOk returns a tuple with the PrepareSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepareSuccessCount

`func (o *MsgVpnClientTransactedSession) SetPrepareSuccessCount(v int64)`

SetPrepareSuccessCount sets PrepareSuccessCount field to given value.

### HasPrepareSuccessCount

`func (o *MsgVpnClientTransactedSession) HasPrepareSuccessCount() bool`

HasPrepareSuccessCount returns a boolean if a field has been set.

### GetPreviousTransactionState

`func (o *MsgVpnClientTransactedSession) GetPreviousTransactionState() string`

GetPreviousTransactionState returns the PreviousTransactionState field if non-nil, zero value otherwise.

### GetPreviousTransactionStateOk

`func (o *MsgVpnClientTransactedSession) GetPreviousTransactionStateOk() (*string, bool)`

GetPreviousTransactionStateOk returns a tuple with the PreviousTransactionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTransactionState

`func (o *MsgVpnClientTransactedSession) SetPreviousTransactionState(v string)`

SetPreviousTransactionState sets PreviousTransactionState field to given value.

### HasPreviousTransactionState

`func (o *MsgVpnClientTransactedSession) HasPreviousTransactionState() bool`

HasPreviousTransactionState returns a boolean if a field has been set.

### GetPublishedMsgCount

`func (o *MsgVpnClientTransactedSession) GetPublishedMsgCount() int64`

GetPublishedMsgCount returns the PublishedMsgCount field if non-nil, zero value otherwise.

### GetPublishedMsgCountOk

`func (o *MsgVpnClientTransactedSession) GetPublishedMsgCountOk() (*int64, bool)`

GetPublishedMsgCountOk returns a tuple with the PublishedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedMsgCount

`func (o *MsgVpnClientTransactedSession) SetPublishedMsgCount(v int64)`

SetPublishedMsgCount sets PublishedMsgCount field to given value.

### HasPublishedMsgCount

`func (o *MsgVpnClientTransactedSession) HasPublishedMsgCount() bool`

HasPublishedMsgCount returns a boolean if a field has been set.

### GetResumeFailureCount

`func (o *MsgVpnClientTransactedSession) GetResumeFailureCount() int64`

GetResumeFailureCount returns the ResumeFailureCount field if non-nil, zero value otherwise.

### GetResumeFailureCountOk

`func (o *MsgVpnClientTransactedSession) GetResumeFailureCountOk() (*int64, bool)`

GetResumeFailureCountOk returns a tuple with the ResumeFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeFailureCount

`func (o *MsgVpnClientTransactedSession) SetResumeFailureCount(v int64)`

SetResumeFailureCount sets ResumeFailureCount field to given value.

### HasResumeFailureCount

`func (o *MsgVpnClientTransactedSession) HasResumeFailureCount() bool`

HasResumeFailureCount returns a boolean if a field has been set.

### GetResumeSuccessCount

`func (o *MsgVpnClientTransactedSession) GetResumeSuccessCount() int64`

GetResumeSuccessCount returns the ResumeSuccessCount field if non-nil, zero value otherwise.

### GetResumeSuccessCountOk

`func (o *MsgVpnClientTransactedSession) GetResumeSuccessCountOk() (*int64, bool)`

GetResumeSuccessCountOk returns a tuple with the ResumeSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeSuccessCount

`func (o *MsgVpnClientTransactedSession) SetResumeSuccessCount(v int64)`

SetResumeSuccessCount sets ResumeSuccessCount field to given value.

### HasResumeSuccessCount

`func (o *MsgVpnClientTransactedSession) HasResumeSuccessCount() bool`

HasResumeSuccessCount returns a boolean if a field has been set.

### GetRetrievedMsgCount

`func (o *MsgVpnClientTransactedSession) GetRetrievedMsgCount() int64`

GetRetrievedMsgCount returns the RetrievedMsgCount field if non-nil, zero value otherwise.

### GetRetrievedMsgCountOk

`func (o *MsgVpnClientTransactedSession) GetRetrievedMsgCountOk() (*int64, bool)`

GetRetrievedMsgCountOk returns a tuple with the RetrievedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievedMsgCount

`func (o *MsgVpnClientTransactedSession) SetRetrievedMsgCount(v int64)`

SetRetrievedMsgCount sets RetrievedMsgCount field to given value.

### HasRetrievedMsgCount

`func (o *MsgVpnClientTransactedSession) HasRetrievedMsgCount() bool`

HasRetrievedMsgCount returns a boolean if a field has been set.

### GetRollbackCount

`func (o *MsgVpnClientTransactedSession) GetRollbackCount() int64`

GetRollbackCount returns the RollbackCount field if non-nil, zero value otherwise.

### GetRollbackCountOk

`func (o *MsgVpnClientTransactedSession) GetRollbackCountOk() (*int64, bool)`

GetRollbackCountOk returns a tuple with the RollbackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackCount

`func (o *MsgVpnClientTransactedSession) SetRollbackCount(v int64)`

SetRollbackCount sets RollbackCount field to given value.

### HasRollbackCount

`func (o *MsgVpnClientTransactedSession) HasRollbackCount() bool`

HasRollbackCount returns a boolean if a field has been set.

### GetRollbackFailureCount

`func (o *MsgVpnClientTransactedSession) GetRollbackFailureCount() int64`

GetRollbackFailureCount returns the RollbackFailureCount field if non-nil, zero value otherwise.

### GetRollbackFailureCountOk

`func (o *MsgVpnClientTransactedSession) GetRollbackFailureCountOk() (*int64, bool)`

GetRollbackFailureCountOk returns a tuple with the RollbackFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackFailureCount

`func (o *MsgVpnClientTransactedSession) SetRollbackFailureCount(v int64)`

SetRollbackFailureCount sets RollbackFailureCount field to given value.

### HasRollbackFailureCount

`func (o *MsgVpnClientTransactedSession) HasRollbackFailureCount() bool`

HasRollbackFailureCount returns a boolean if a field has been set.

### GetRollbackSuccessCount

`func (o *MsgVpnClientTransactedSession) GetRollbackSuccessCount() int64`

GetRollbackSuccessCount returns the RollbackSuccessCount field if non-nil, zero value otherwise.

### GetRollbackSuccessCountOk

`func (o *MsgVpnClientTransactedSession) GetRollbackSuccessCountOk() (*int64, bool)`

GetRollbackSuccessCountOk returns a tuple with the RollbackSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackSuccessCount

`func (o *MsgVpnClientTransactedSession) SetRollbackSuccessCount(v int64)`

SetRollbackSuccessCount sets RollbackSuccessCount field to given value.

### HasRollbackSuccessCount

`func (o *MsgVpnClientTransactedSession) HasRollbackSuccessCount() bool`

HasRollbackSuccessCount returns a boolean if a field has been set.

### GetSessionName

`func (o *MsgVpnClientTransactedSession) GetSessionName() string`

GetSessionName returns the SessionName field if non-nil, zero value otherwise.

### GetSessionNameOk

`func (o *MsgVpnClientTransactedSession) GetSessionNameOk() (*string, bool)`

GetSessionNameOk returns a tuple with the SessionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionName

`func (o *MsgVpnClientTransactedSession) SetSessionName(v string)`

SetSessionName sets SessionName field to given value.

### HasSessionName

`func (o *MsgVpnClientTransactedSession) HasSessionName() bool`

HasSessionName returns a boolean if a field has been set.

### GetSpooledMsgCount

`func (o *MsgVpnClientTransactedSession) GetSpooledMsgCount() int64`

GetSpooledMsgCount returns the SpooledMsgCount field if non-nil, zero value otherwise.

### GetSpooledMsgCountOk

`func (o *MsgVpnClientTransactedSession) GetSpooledMsgCountOk() (*int64, bool)`

GetSpooledMsgCountOk returns a tuple with the SpooledMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpooledMsgCount

`func (o *MsgVpnClientTransactedSession) SetSpooledMsgCount(v int64)`

SetSpooledMsgCount sets SpooledMsgCount field to given value.

### HasSpooledMsgCount

`func (o *MsgVpnClientTransactedSession) HasSpooledMsgCount() bool`

HasSpooledMsgCount returns a boolean if a field has been set.

### GetStartFailureCount

`func (o *MsgVpnClientTransactedSession) GetStartFailureCount() int64`

GetStartFailureCount returns the StartFailureCount field if non-nil, zero value otherwise.

### GetStartFailureCountOk

`func (o *MsgVpnClientTransactedSession) GetStartFailureCountOk() (*int64, bool)`

GetStartFailureCountOk returns a tuple with the StartFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFailureCount

`func (o *MsgVpnClientTransactedSession) SetStartFailureCount(v int64)`

SetStartFailureCount sets StartFailureCount field to given value.

### HasStartFailureCount

`func (o *MsgVpnClientTransactedSession) HasStartFailureCount() bool`

HasStartFailureCount returns a boolean if a field has been set.

### GetStartSuccessCount

`func (o *MsgVpnClientTransactedSession) GetStartSuccessCount() int64`

GetStartSuccessCount returns the StartSuccessCount field if non-nil, zero value otherwise.

### GetStartSuccessCountOk

`func (o *MsgVpnClientTransactedSession) GetStartSuccessCountOk() (*int64, bool)`

GetStartSuccessCountOk returns a tuple with the StartSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSuccessCount

`func (o *MsgVpnClientTransactedSession) SetStartSuccessCount(v int64)`

SetStartSuccessCount sets StartSuccessCount field to given value.

### HasStartSuccessCount

`func (o *MsgVpnClientTransactedSession) HasStartSuccessCount() bool`

HasStartSuccessCount returns a boolean if a field has been set.

### GetSuccessCount

`func (o *MsgVpnClientTransactedSession) GetSuccessCount() int64`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *MsgVpnClientTransactedSession) GetSuccessCountOk() (*int64, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *MsgVpnClientTransactedSession) SetSuccessCount(v int64)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *MsgVpnClientTransactedSession) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetSuspendFailureCount

`func (o *MsgVpnClientTransactedSession) GetSuspendFailureCount() int64`

GetSuspendFailureCount returns the SuspendFailureCount field if non-nil, zero value otherwise.

### GetSuspendFailureCountOk

`func (o *MsgVpnClientTransactedSession) GetSuspendFailureCountOk() (*int64, bool)`

GetSuspendFailureCountOk returns a tuple with the SuspendFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendFailureCount

`func (o *MsgVpnClientTransactedSession) SetSuspendFailureCount(v int64)`

SetSuspendFailureCount sets SuspendFailureCount field to given value.

### HasSuspendFailureCount

`func (o *MsgVpnClientTransactedSession) HasSuspendFailureCount() bool`

HasSuspendFailureCount returns a boolean if a field has been set.

### GetSuspendSuccessCount

`func (o *MsgVpnClientTransactedSession) GetSuspendSuccessCount() int64`

GetSuspendSuccessCount returns the SuspendSuccessCount field if non-nil, zero value otherwise.

### GetSuspendSuccessCountOk

`func (o *MsgVpnClientTransactedSession) GetSuspendSuccessCountOk() (*int64, bool)`

GetSuspendSuccessCountOk returns a tuple with the SuspendSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendSuccessCount

`func (o *MsgVpnClientTransactedSession) SetSuspendSuccessCount(v int64)`

SetSuspendSuccessCount sets SuspendSuccessCount field to given value.

### HasSuspendSuccessCount

`func (o *MsgVpnClientTransactedSession) HasSuspendSuccessCount() bool`

HasSuspendSuccessCount returns a boolean if a field has been set.

### GetTransactionId

`func (o *MsgVpnClientTransactedSession) GetTransactionId() int32`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *MsgVpnClientTransactedSession) GetTransactionIdOk() (*int32, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *MsgVpnClientTransactedSession) SetTransactionId(v int32)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *MsgVpnClientTransactedSession) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionState

`func (o *MsgVpnClientTransactedSession) GetTransactionState() string`

GetTransactionState returns the TransactionState field if non-nil, zero value otherwise.

### GetTransactionStateOk

`func (o *MsgVpnClientTransactedSession) GetTransactionStateOk() (*string, bool)`

GetTransactionStateOk returns a tuple with the TransactionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionState

`func (o *MsgVpnClientTransactedSession) SetTransactionState(v string)`

SetTransactionState sets TransactionState field to given value.

### HasTransactionState

`func (o *MsgVpnClientTransactedSession) HasTransactionState() bool`

HasTransactionState returns a boolean if a field has been set.

### GetTwoPhaseCommitFailureCount

`func (o *MsgVpnClientTransactedSession) GetTwoPhaseCommitFailureCount() int64`

GetTwoPhaseCommitFailureCount returns the TwoPhaseCommitFailureCount field if non-nil, zero value otherwise.

### GetTwoPhaseCommitFailureCountOk

`func (o *MsgVpnClientTransactedSession) GetTwoPhaseCommitFailureCountOk() (*int64, bool)`

GetTwoPhaseCommitFailureCountOk returns a tuple with the TwoPhaseCommitFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoPhaseCommitFailureCount

`func (o *MsgVpnClientTransactedSession) SetTwoPhaseCommitFailureCount(v int64)`

SetTwoPhaseCommitFailureCount sets TwoPhaseCommitFailureCount field to given value.

### HasTwoPhaseCommitFailureCount

`func (o *MsgVpnClientTransactedSession) HasTwoPhaseCommitFailureCount() bool`

HasTwoPhaseCommitFailureCount returns a boolean if a field has been set.

### GetTwoPhaseCommitSuccessCount

`func (o *MsgVpnClientTransactedSession) GetTwoPhaseCommitSuccessCount() int64`

GetTwoPhaseCommitSuccessCount returns the TwoPhaseCommitSuccessCount field if non-nil, zero value otherwise.

### GetTwoPhaseCommitSuccessCountOk

`func (o *MsgVpnClientTransactedSession) GetTwoPhaseCommitSuccessCountOk() (*int64, bool)`

GetTwoPhaseCommitSuccessCountOk returns a tuple with the TwoPhaseCommitSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoPhaseCommitSuccessCount

`func (o *MsgVpnClientTransactedSession) SetTwoPhaseCommitSuccessCount(v int64)`

SetTwoPhaseCommitSuccessCount sets TwoPhaseCommitSuccessCount field to given value.

### HasTwoPhaseCommitSuccessCount

`func (o *MsgVpnClientTransactedSession) HasTwoPhaseCommitSuccessCount() bool`

HasTwoPhaseCommitSuccessCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


