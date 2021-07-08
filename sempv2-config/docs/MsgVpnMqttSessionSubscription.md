# MsgVpnMqttSessionSubscription

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MqttSessionClientId** | **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | [optional] [default to null]
**MqttSessionVirtualRouter** | **string** | The virtual router of the MQTT Session. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - The MQTT Session belongs to the primary virtual router. \&quot;backup\&quot; - The MQTT Session belongs to the backup virtual router. &lt;/pre&gt;  | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**SubscriptionQos** | **int64** | The quality of service (QoS) for the subscription as either 0 (deliver at most once) or 1 (deliver at least once). QoS 2 is not supported, but QoS 2 messages attracted by QoS 0 or QoS 1 subscriptions are accepted and delivered accordingly. The default value is &#x60;0&#x60;. | [optional] [default to null]
**SubscriptionTopic** | **string** | The MQTT subscription topic. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

