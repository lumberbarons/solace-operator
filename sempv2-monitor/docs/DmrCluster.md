# DmrCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationBasicEnabled** | Pointer to **bool** | Indicates whether basic authentication is enabled for Cluster Links. | [optional] 
**AuthenticationBasicType** | Pointer to **string** | The type of basic authentication to use for Cluster Links. The allowed values and their meaning are:  &lt;pre&gt; \&quot;internal\&quot; - Use locally configured password. \&quot;none\&quot; - No authentication. &lt;/pre&gt;  | [optional] 
**AuthenticationClientCertEnabled** | Pointer to **bool** | Indicates whether client certificate authentication is enabled for Cluster Links. | [optional] 
**DirectOnlyEnabled** | Pointer to **bool** | Indicates whether this cluster only supports direct messaging. If true, guaranteed messages will not be transmitted through the cluster. | [optional] 
**DmrClusterName** | Pointer to **string** | The name of the Cluster. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Cluster is enabled. | [optional] 
**FailureReason** | Pointer to **string** | The failure reason for the Cluster being down. | [optional] 
**NodeName** | Pointer to **string** | The name of this node in the Cluster. This is the name that this broker (or redundant group of brokers) is know by to other nodes in the Cluster. The name is chosen automatically to be either this broker&#39;s Router Name or Mate Router Name, depending on which Active Standby Role (primary or backup) this broker plays in its redundancy group. | [optional] 
**SubscriptionDbBuildPercentage** | Pointer to **int64** | Cluster Subscription Database build completion percentage. Available since 2.20. | [optional] 
**TlsServerCertEnforceTrustedCommonNameEnabled** | Pointer to **bool** | Indicates whether the common name provided by the remote broker is enforced against the list of trusted common names configured for the Link. If enabled, the certificate&#39;s common name must match one of the trusted common names for the Link to be accepted. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation. | [optional] 
**TlsServerCertMaxChainDepth** | Pointer to **int64** | The maximum allowed depth of a certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate. | [optional] 
**TlsServerCertValidateDateEnabled** | Pointer to **bool** | Indicates whether validation of the \&quot;Not Before\&quot; and \&quot;Not After\&quot; validity dates in the certificate is enabled. When disabled, the certificate is accepted even if the certificate is not valid based on these dates. | [optional] 
**TlsServerCertValidateNameEnabled** | Pointer to **bool** | Enable or disable the standard TLS authentication mechanism of verifying the name used to connect to the bridge. If enabled, the name used to connect to the bridge is checked against the names specified in the certificate returned by the remote router. Legacy Common Name validation is not performed if Server Certificate Name Validation is enabled, even if Common Name validation is also enabled. Available since 2.18. | [optional] 
**Up** | Pointer to **bool** | Indicates whether the Cluster is operationally up. | [optional] 
**Uptime** | Pointer to **int64** | The amount of time in seconds since the Cluster was up. | [optional] 

## Methods

### NewDmrCluster

`func NewDmrCluster() *DmrCluster`

NewDmrCluster instantiates a new DmrCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDmrClusterWithDefaults

`func NewDmrClusterWithDefaults() *DmrCluster`

NewDmrClusterWithDefaults instantiates a new DmrCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationBasicEnabled

`func (o *DmrCluster) GetAuthenticationBasicEnabled() bool`

GetAuthenticationBasicEnabled returns the AuthenticationBasicEnabled field if non-nil, zero value otherwise.

### GetAuthenticationBasicEnabledOk

`func (o *DmrCluster) GetAuthenticationBasicEnabledOk() (*bool, bool)`

GetAuthenticationBasicEnabledOk returns a tuple with the AuthenticationBasicEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationBasicEnabled

`func (o *DmrCluster) SetAuthenticationBasicEnabled(v bool)`

SetAuthenticationBasicEnabled sets AuthenticationBasicEnabled field to given value.

### HasAuthenticationBasicEnabled

`func (o *DmrCluster) HasAuthenticationBasicEnabled() bool`

HasAuthenticationBasicEnabled returns a boolean if a field has been set.

### GetAuthenticationBasicType

`func (o *DmrCluster) GetAuthenticationBasicType() string`

GetAuthenticationBasicType returns the AuthenticationBasicType field if non-nil, zero value otherwise.

### GetAuthenticationBasicTypeOk

`func (o *DmrCluster) GetAuthenticationBasicTypeOk() (*string, bool)`

GetAuthenticationBasicTypeOk returns a tuple with the AuthenticationBasicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationBasicType

`func (o *DmrCluster) SetAuthenticationBasicType(v string)`

SetAuthenticationBasicType sets AuthenticationBasicType field to given value.

### HasAuthenticationBasicType

`func (o *DmrCluster) HasAuthenticationBasicType() bool`

HasAuthenticationBasicType returns a boolean if a field has been set.

### GetAuthenticationClientCertEnabled

`func (o *DmrCluster) GetAuthenticationClientCertEnabled() bool`

GetAuthenticationClientCertEnabled returns the AuthenticationClientCertEnabled field if non-nil, zero value otherwise.

### GetAuthenticationClientCertEnabledOk

`func (o *DmrCluster) GetAuthenticationClientCertEnabledOk() (*bool, bool)`

GetAuthenticationClientCertEnabledOk returns a tuple with the AuthenticationClientCertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationClientCertEnabled

`func (o *DmrCluster) SetAuthenticationClientCertEnabled(v bool)`

SetAuthenticationClientCertEnabled sets AuthenticationClientCertEnabled field to given value.

### HasAuthenticationClientCertEnabled

`func (o *DmrCluster) HasAuthenticationClientCertEnabled() bool`

HasAuthenticationClientCertEnabled returns a boolean if a field has been set.

### GetDirectOnlyEnabled

`func (o *DmrCluster) GetDirectOnlyEnabled() bool`

GetDirectOnlyEnabled returns the DirectOnlyEnabled field if non-nil, zero value otherwise.

### GetDirectOnlyEnabledOk

`func (o *DmrCluster) GetDirectOnlyEnabledOk() (*bool, bool)`

GetDirectOnlyEnabledOk returns a tuple with the DirectOnlyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectOnlyEnabled

`func (o *DmrCluster) SetDirectOnlyEnabled(v bool)`

SetDirectOnlyEnabled sets DirectOnlyEnabled field to given value.

### HasDirectOnlyEnabled

`func (o *DmrCluster) HasDirectOnlyEnabled() bool`

HasDirectOnlyEnabled returns a boolean if a field has been set.

### GetDmrClusterName

`func (o *DmrCluster) GetDmrClusterName() string`

GetDmrClusterName returns the DmrClusterName field if non-nil, zero value otherwise.

### GetDmrClusterNameOk

`func (o *DmrCluster) GetDmrClusterNameOk() (*string, bool)`

GetDmrClusterNameOk returns a tuple with the DmrClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmrClusterName

`func (o *DmrCluster) SetDmrClusterName(v string)`

SetDmrClusterName sets DmrClusterName field to given value.

### HasDmrClusterName

`func (o *DmrCluster) HasDmrClusterName() bool`

HasDmrClusterName returns a boolean if a field has been set.

### GetEnabled

`func (o *DmrCluster) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DmrCluster) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DmrCluster) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DmrCluster) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFailureReason

`func (o *DmrCluster) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *DmrCluster) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *DmrCluster) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *DmrCluster) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetNodeName

`func (o *DmrCluster) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *DmrCluster) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *DmrCluster) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *DmrCluster) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetSubscriptionDbBuildPercentage

`func (o *DmrCluster) GetSubscriptionDbBuildPercentage() int64`

GetSubscriptionDbBuildPercentage returns the SubscriptionDbBuildPercentage field if non-nil, zero value otherwise.

### GetSubscriptionDbBuildPercentageOk

`func (o *DmrCluster) GetSubscriptionDbBuildPercentageOk() (*int64, bool)`

GetSubscriptionDbBuildPercentageOk returns a tuple with the SubscriptionDbBuildPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionDbBuildPercentage

`func (o *DmrCluster) SetSubscriptionDbBuildPercentage(v int64)`

SetSubscriptionDbBuildPercentage sets SubscriptionDbBuildPercentage field to given value.

### HasSubscriptionDbBuildPercentage

`func (o *DmrCluster) HasSubscriptionDbBuildPercentage() bool`

HasSubscriptionDbBuildPercentage returns a boolean if a field has been set.

### GetTlsServerCertEnforceTrustedCommonNameEnabled

`func (o *DmrCluster) GetTlsServerCertEnforceTrustedCommonNameEnabled() bool`

GetTlsServerCertEnforceTrustedCommonNameEnabled returns the TlsServerCertEnforceTrustedCommonNameEnabled field if non-nil, zero value otherwise.

### GetTlsServerCertEnforceTrustedCommonNameEnabledOk

`func (o *DmrCluster) GetTlsServerCertEnforceTrustedCommonNameEnabledOk() (*bool, bool)`

GetTlsServerCertEnforceTrustedCommonNameEnabledOk returns a tuple with the TlsServerCertEnforceTrustedCommonNameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerCertEnforceTrustedCommonNameEnabled

`func (o *DmrCluster) SetTlsServerCertEnforceTrustedCommonNameEnabled(v bool)`

SetTlsServerCertEnforceTrustedCommonNameEnabled sets TlsServerCertEnforceTrustedCommonNameEnabled field to given value.

### HasTlsServerCertEnforceTrustedCommonNameEnabled

`func (o *DmrCluster) HasTlsServerCertEnforceTrustedCommonNameEnabled() bool`

HasTlsServerCertEnforceTrustedCommonNameEnabled returns a boolean if a field has been set.

### GetTlsServerCertMaxChainDepth

`func (o *DmrCluster) GetTlsServerCertMaxChainDepth() int64`

GetTlsServerCertMaxChainDepth returns the TlsServerCertMaxChainDepth field if non-nil, zero value otherwise.

### GetTlsServerCertMaxChainDepthOk

`func (o *DmrCluster) GetTlsServerCertMaxChainDepthOk() (*int64, bool)`

GetTlsServerCertMaxChainDepthOk returns a tuple with the TlsServerCertMaxChainDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerCertMaxChainDepth

`func (o *DmrCluster) SetTlsServerCertMaxChainDepth(v int64)`

SetTlsServerCertMaxChainDepth sets TlsServerCertMaxChainDepth field to given value.

### HasTlsServerCertMaxChainDepth

`func (o *DmrCluster) HasTlsServerCertMaxChainDepth() bool`

HasTlsServerCertMaxChainDepth returns a boolean if a field has been set.

### GetTlsServerCertValidateDateEnabled

`func (o *DmrCluster) GetTlsServerCertValidateDateEnabled() bool`

GetTlsServerCertValidateDateEnabled returns the TlsServerCertValidateDateEnabled field if non-nil, zero value otherwise.

### GetTlsServerCertValidateDateEnabledOk

`func (o *DmrCluster) GetTlsServerCertValidateDateEnabledOk() (*bool, bool)`

GetTlsServerCertValidateDateEnabledOk returns a tuple with the TlsServerCertValidateDateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerCertValidateDateEnabled

`func (o *DmrCluster) SetTlsServerCertValidateDateEnabled(v bool)`

SetTlsServerCertValidateDateEnabled sets TlsServerCertValidateDateEnabled field to given value.

### HasTlsServerCertValidateDateEnabled

`func (o *DmrCluster) HasTlsServerCertValidateDateEnabled() bool`

HasTlsServerCertValidateDateEnabled returns a boolean if a field has been set.

### GetTlsServerCertValidateNameEnabled

`func (o *DmrCluster) GetTlsServerCertValidateNameEnabled() bool`

GetTlsServerCertValidateNameEnabled returns the TlsServerCertValidateNameEnabled field if non-nil, zero value otherwise.

### GetTlsServerCertValidateNameEnabledOk

`func (o *DmrCluster) GetTlsServerCertValidateNameEnabledOk() (*bool, bool)`

GetTlsServerCertValidateNameEnabledOk returns a tuple with the TlsServerCertValidateNameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerCertValidateNameEnabled

`func (o *DmrCluster) SetTlsServerCertValidateNameEnabled(v bool)`

SetTlsServerCertValidateNameEnabled sets TlsServerCertValidateNameEnabled field to given value.

### HasTlsServerCertValidateNameEnabled

`func (o *DmrCluster) HasTlsServerCertValidateNameEnabled() bool`

HasTlsServerCertValidateNameEnabled returns a boolean if a field has been set.

### GetUp

`func (o *DmrCluster) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *DmrCluster) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *DmrCluster) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *DmrCluster) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUptime

`func (o *DmrCluster) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *DmrCluster) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *DmrCluster) SetUptime(v int64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *DmrCluster) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


