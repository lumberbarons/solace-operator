# MsgVpnClientUsername

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfileName** | **string** | The ACL Profile of the Client Username. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] [default to null]
**ClientProfileName** | **string** | The Client Profile of the Client Username. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] [default to null]
**ClientUsername** | **string** | The name of the Client Username. | [optional] [default to null]
**Enabled** | **bool** | Enable or disable the Client Username. When disabled, all clients currently connected as the Client Username are disconnected. The default value is &#x60;false&#x60;. | [optional] [default to null]
**GuaranteedEndpointPermissionOverrideEnabled** | **bool** | Enable or disable guaranteed endpoint permission override for the Client Username. When enabled all guaranteed endpoints may be accessed, modified or deleted with the same permission as the owner. The default value is &#x60;false&#x60;. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**Password** | **string** | The password for the Client Username. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] [default to null]
**SubscriptionManagerEnabled** | **bool** | Enable or disable the subscription management capability of the Client Username. This is the ability to manage subscriptions on behalf of other Client Usernames. The default value is &#x60;false&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

