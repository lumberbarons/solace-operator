# MsgVpnBridgeRemoteMsgVpn

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeName** | **string** | The name of the Bridge. | [optional] [default to null]
**BridgeVirtualRouter** | **string** | The virtual router of the Bridge. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - The Bridge is used for the primary virtual router. \&quot;backup\&quot; - The Bridge is used for the backup virtual router. \&quot;auto\&quot; - The Bridge is automatically assigned a virtual router at creation, depending on the broker&#x27;s active-standby role. &lt;/pre&gt;  | [optional] [default to null]
**ClientUsername** | **string** | The Client Username the Bridge uses to login to the remote Message VPN. This per remote Message VPN value overrides the value provided for the Bridge overall. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] [default to null]
**CompressedDataEnabled** | **bool** | Enable or disable data compression for the remote Message VPN connection. The default value is &#x60;false&#x60;. | [optional] [default to null]
**ConnectOrder** | **int32** | The preference given to incoming connections from remote Message VPN hosts, from 1 (highest priority) to 4 (lowest priority). The default value is &#x60;4&#x60;. | [optional] [default to null]
**EgressFlowWindowSize** | **int64** | The number of outstanding guaranteed messages that can be transmitted over the remote Message VPN connection before an acknowledgement is received. The default value is &#x60;255&#x60;. | [optional] [default to null]
**Enabled** | **bool** | Enable or disable the remote Message VPN. The default value is &#x60;false&#x60;. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**Password** | **string** | The password for the Client Username. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] [default to null]
**QueueBinding** | **string** | The queue binding of the Bridge in the remote Message VPN. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] [default to null]
**RemoteMsgVpnInterface** | **string** | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, &#x60;remoteMsgVpnLocation&#x60; must not be a virtual router name. | [optional] [default to null]
**RemoteMsgVpnLocation** | **string** | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \&quot;v:\&quot;). | [optional] [default to null]
**RemoteMsgVpnName** | **string** | The name of the remote Message VPN. | [optional] [default to null]
**TlsEnabled** | **bool** | Enable or disable encryption (TLS) for the remote Message VPN connection. The default value is &#x60;false&#x60;. | [optional] [default to null]
**UnidirectionalClientProfile** | **string** | The Client Profile for the unidirectional Bridge of the remote Message VPN. The Client Profile must exist in the local Message VPN, and it is used only for the TCP parameters. Note that the default client profile has a TCP maximum window size of 2MB. The default value is &#x60;\&quot;#client-profile\&quot;&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

