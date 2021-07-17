# DmrClusterTopologyIssueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**DmrClusterTopologyIssue**](DmrClusterTopologyIssue.md) |  | [optional] 
**Links** | Pointer to [**DmrClusterTopologyIssueLinks**](DmrClusterTopologyIssueLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewDmrClusterTopologyIssueResponse

`func NewDmrClusterTopologyIssueResponse(meta SempMeta, ) *DmrClusterTopologyIssueResponse`

NewDmrClusterTopologyIssueResponse instantiates a new DmrClusterTopologyIssueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDmrClusterTopologyIssueResponseWithDefaults

`func NewDmrClusterTopologyIssueResponseWithDefaults() *DmrClusterTopologyIssueResponse`

NewDmrClusterTopologyIssueResponseWithDefaults instantiates a new DmrClusterTopologyIssueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *DmrClusterTopologyIssueResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *DmrClusterTopologyIssueResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *DmrClusterTopologyIssueResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *DmrClusterTopologyIssueResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *DmrClusterTopologyIssueResponse) GetData() DmrClusterTopologyIssue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DmrClusterTopologyIssueResponse) GetDataOk() (*DmrClusterTopologyIssue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DmrClusterTopologyIssueResponse) SetData(v DmrClusterTopologyIssue)`

SetData sets Data field to given value.

### HasData

`func (o *DmrClusterTopologyIssueResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *DmrClusterTopologyIssueResponse) GetLinks() DmrClusterTopologyIssueLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DmrClusterTopologyIssueResponse) GetLinksOk() (*DmrClusterTopologyIssueLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DmrClusterTopologyIssueResponse) SetLinks(v DmrClusterTopologyIssueLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DmrClusterTopologyIssueResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *DmrClusterTopologyIssueResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DmrClusterTopologyIssueResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DmrClusterTopologyIssueResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


