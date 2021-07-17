# DmrClusterTopologyIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DmrClusterName** | Pointer to **string** | The name of the Cluster. | [optional] 
**TopologyIssue** | Pointer to **string** | The topology issue discovered in the Cluster. A topology issue indicates incorrect or inconsistent configuration within the DMR network. Such issues will cause messages to be misdelivered or lost. | [optional] 

## Methods

### NewDmrClusterTopologyIssue

`func NewDmrClusterTopologyIssue() *DmrClusterTopologyIssue`

NewDmrClusterTopologyIssue instantiates a new DmrClusterTopologyIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDmrClusterTopologyIssueWithDefaults

`func NewDmrClusterTopologyIssueWithDefaults() *DmrClusterTopologyIssue`

NewDmrClusterTopologyIssueWithDefaults instantiates a new DmrClusterTopologyIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDmrClusterName

`func (o *DmrClusterTopologyIssue) GetDmrClusterName() string`

GetDmrClusterName returns the DmrClusterName field if non-nil, zero value otherwise.

### GetDmrClusterNameOk

`func (o *DmrClusterTopologyIssue) GetDmrClusterNameOk() (*string, bool)`

GetDmrClusterNameOk returns a tuple with the DmrClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmrClusterName

`func (o *DmrClusterTopologyIssue) SetDmrClusterName(v string)`

SetDmrClusterName sets DmrClusterName field to given value.

### HasDmrClusterName

`func (o *DmrClusterTopologyIssue) HasDmrClusterName() bool`

HasDmrClusterName returns a boolean if a field has been set.

### GetTopologyIssue

`func (o *DmrClusterTopologyIssue) GetTopologyIssue() string`

GetTopologyIssue returns the TopologyIssue field if non-nil, zero value otherwise.

### GetTopologyIssueOk

`func (o *DmrClusterTopologyIssue) GetTopologyIssueOk() (*string, bool)`

GetTopologyIssueOk returns a tuple with the TopologyIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyIssue

`func (o *DmrClusterTopologyIssue) SetTopologyIssue(v string)`

SetTopologyIssue sets TopologyIssue field to given value.

### HasTopologyIssue

`func (o *DmrClusterTopologyIssue) HasTopologyIssue() bool`

HasTopologyIssue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


