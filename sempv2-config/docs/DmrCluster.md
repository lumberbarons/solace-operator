# DmrCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationBasicEnabled** | Pointer to **bool** | Enable or disable basic authentication for Cluster Links. The default value is &#x60;true&#x60;. | [optional] 
**AuthenticationBasicPassword** | Pointer to **string** | The password used to authenticate incoming Cluster Links when using basic internal authentication. The same password is also used by outgoing Cluster Links if a per-Link password is not configured. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**AuthenticationBasicType** | Pointer to **string** | The type of basic authentication to use for Cluster Links. The default value is &#x60;\&quot;internal\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;internal\&quot; - Use locally configured password. \&quot;none\&quot; - No authentication. &lt;/pre&gt;  | [optional] 
**AuthenticationClientCertContent** | Pointer to **string** | The PEM formatted content for the client certificate used to login to the remote node. It must consist of a private key and between one and three certificates comprising the certificate trust chain. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**AuthenticationClientCertEnabled** | Pointer to **bool** | Enable or disable client certificate authentication for Cluster Links. The default value is &#x60;true&#x60;. | [optional] 
**AuthenticationClientCertPassword** | Pointer to **string** | The password for the client certificate. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**DirectOnlyEnabled** | Pointer to **bool** | Enable or disable direct messaging only. Guaranteed messages will not be transmitted through the cluster. The default value is &#x60;false&#x60;. | [optional] 
**DmrClusterName** | Pointer to **string** | The name of the Cluster. | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable the Cluster. The default value is &#x60;false&#x60;. | [optional] 
**NodeName** | Pointer to **string** | The name of this node in the Cluster. This is the name that this broker (or redundant group of brokers) is know by to other nodes in the Cluster. The name is chosen automatically to be either this broker&#39;s Router Name or Mate Router Name, depending on which Active Standby Role (primary or backup) this broker plays in its redundancy group. | [optional] 
**TlsServerCertEnforceTrustedCommonNameEnabled** | Pointer to **bool** | Enable or disable the enforcing of the common name provided by the remote broker against the list of trusted common names configured for the Link. If enabled, the certificate&#39;s common name must match one of the trusted common names for the Link to be accepted. Common Name validation is not performed if Server Certificate Name Validation is enabled, even if Common Name validation is enabled. The default value is &#x60;true&#x60;. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation. | [optional] 
**TlsServerCertMaxChainDepth** | Pointer to **int64** | The maximum allowed depth of a certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate. The default value is &#x60;3&#x60;. | [optional] 
**TlsServerCertValidateDateEnabled** | Pointer to **bool** | Enable or disable the validation of the \&quot;Not Before\&quot; and \&quot;Not After\&quot; validity dates in the certificate. When disabled, the certificate is accepted even if the certificate is not valid based on these dates. The default value is &#x60;true&#x60;. | [optional] 
**TlsServerCertValidateNameEnabled** | Pointer to **bool** | Enable or disable the standard TLS authentication mechanism of verifying the name used to connect to the bridge. If enabled, the name used to connect to the bridge is checked against the names specified in the certificate returned by the remote router. Legacy Common Name validation is not performed if Server Certificate Name Validation is enabled, even if Common Name validation is also enabled. The default value is &#x60;true&#x60;. Available since 2.18. | [optional] 

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

### GetAuthenticationBasicPassword

`func (o *DmrCluster) GetAuthenticationBasicPassword() string`

GetAuthenticationBasicPassword returns the AuthenticationBasicPassword field if non-nil, zero value otherwise.

### GetAuthenticationBasicPasswordOk

`func (o *DmrCluster) GetAuthenticationBasicPasswordOk() (*string, bool)`

GetAuthenticationBasicPasswordOk returns a tuple with the AuthenticationBasicPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationBasicPassword

`func (o *DmrCluster) SetAuthenticationBasicPassword(v string)`

SetAuthenticationBasicPassword sets AuthenticationBasicPassword field to given value.

### HasAuthenticationBasicPassword

`func (o *DmrCluster) HasAuthenticationBasicPassword() bool`

HasAuthenticationBasicPassword returns a boolean if a field has been set.

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

### GetAuthenticationClientCertContent

`func (o *DmrCluster) GetAuthenticationClientCertContent() string`

GetAuthenticationClientCertContent returns the AuthenticationClientCertContent field if non-nil, zero value otherwise.

### GetAuthenticationClientCertContentOk

`func (o *DmrCluster) GetAuthenticationClientCertContentOk() (*string, bool)`

GetAuthenticationClientCertContentOk returns a tuple with the AuthenticationClientCertContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationClientCertContent

`func (o *DmrCluster) SetAuthenticationClientCertContent(v string)`

SetAuthenticationClientCertContent sets AuthenticationClientCertContent field to given value.

### HasAuthenticationClientCertContent

`func (o *DmrCluster) HasAuthenticationClientCertContent() bool`

HasAuthenticationClientCertContent returns a boolean if a field has been set.

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

### GetAuthenticationClientCertPassword

`func (o *DmrCluster) GetAuthenticationClientCertPassword() string`

GetAuthenticationClientCertPassword returns the AuthenticationClientCertPassword field if non-nil, zero value otherwise.

### GetAuthenticationClientCertPasswordOk

`func (o *DmrCluster) GetAuthenticationClientCertPasswordOk() (*string, bool)`

GetAuthenticationClientCertPasswordOk returns a tuple with the AuthenticationClientCertPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationClientCertPassword

`func (o *DmrCluster) SetAuthenticationClientCertPassword(v string)`

SetAuthenticationClientCertPassword sets AuthenticationClientCertPassword field to given value.

### HasAuthenticationClientCertPassword

`func (o *DmrCluster) HasAuthenticationClientCertPassword() bool`

HasAuthenticationClientCertPassword returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


