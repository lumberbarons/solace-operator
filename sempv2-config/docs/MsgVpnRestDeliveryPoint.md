# MsgVpnRestDeliveryPoint

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientProfileName** | **string** | The Client Profile of the REST Delivery Point. It must exist in the local Message VPN. Its TCP parameters are used for all REST Consumers in this RDP. Its queue properties are used by the RDP client. The Client Profile is used inside the auto-generated Client Username for this RDP. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] [default to null]
**Enabled** | **bool** | Enable or disable the REST Delivery Point. When disabled, no connections are initiated or messages delivered to any of the contained REST Consumers. The default value is &#x60;false&#x60;. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**RestDeliveryPointName** | **string** | The name of the REST Delivery Point. | [optional] [default to null]
**Service** | **string** | The name of the service that this REST Delivery Point connects to. Internally the broker does not use this value; it is informational only. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.19. | [optional] [default to null]
**Vendor** | **string** | The name of the vendor that this REST Delivery Point connects to. Internally the broker does not use this value; it is informational only. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.19. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

