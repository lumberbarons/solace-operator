# MsgVpnAuthorizationGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfileName** | **string** | The ACL Profile of the LDAP Authorization Group. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] [default to null]
**AuthorizationGroupName** | **string** | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as &#x27;#&#x27;, &#x27;+&#x27;, &#x27;;&#x27;, &#x27;&#x3D;&#x27; as the value of the group name returned from the LDAP server might prepend those characters with &#x27;\\&#x27;. For example a group name called &#x27;test#,lab,com&#x27; will be returned from the LDAP server as &#x27;test\\#,lab,com&#x27;. | [optional] [default to null]
**ClientProfileName** | **string** | The Client Profile of the LDAP Authorization Group. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] [default to null]
**Enabled** | **bool** | Enable or disable the LDAP Authorization Group in the Message VPN. The default value is &#x60;false&#x60;. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**OrderAfterAuthorizationGroupName** | **string** | Lower the priority to be less than this group. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default is not applicable. | [optional] [default to null]
**OrderBeforeAuthorizationGroupName** | **string** | Raise the priority to be greater than this group. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default is not applicable. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

