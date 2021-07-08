# CertAuthority

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertAuthorityName** | **string** | The name of the Certificate Authority. Deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]
**CertContent** | **string** | The PEM formatted content for the trusted root certificate of a Certificate Authority. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]
**CrlDayList** | **string** | The scheduled CRL refresh day(s), specified as \&quot;daily\&quot; or a comma-separated list of days. Days must be specified as \&quot;Sun\&quot;, \&quot;Mon\&quot;, \&quot;Tue\&quot;, \&quot;Wed\&quot;, \&quot;Thu\&quot;, \&quot;Fri\&quot;, or \&quot;Sat\&quot;, with no spaces, and in sorted order from Sunday to Saturday. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]
**CrlLastDownloadTime** | **int32** | The timestamp of the last successful CRL download. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]
**CrlLastFailureReason** | **string** | The reason for the last CRL failure. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]
**CrlLastFailureTime** | **int32** | The timestamp of the last CRL failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]
**CrlNextDownloadTime** | **int32** | The scheduled time of the next CRL download. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]
**CrlTimeList** | **string** | The scheduled CRL refresh time(s), specified as \&quot;hourly\&quot; or a comma-separated list of 24-hour times in the form hh:mm, or h:mm. There must be no spaces, and times must be in sorted order from 0:00 to 23:59. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]
**CrlUp** | **bool** | Indicates whether CRL revocation checking is operationally up. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]
**CrlUrl** | **string** | The URL for the CRL source. This is a required attribute for CRL to be operational and the URL must be complete with http:// included. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]
**OcspLastFailureReason** | **string** | The reason for the last OCSP failure. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]
**OcspLastFailureTime** | **int32** | The timestamp of the last OCSP failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]
**OcspLastFailureUrl** | **string** | The URL involved in the last OCSP failure. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]
**OcspNonResponderCertEnabled** | **bool** | Indicates whether a non-responder certificate is allowed to sign an OCSP response. Typically used with an OCSP override URL in cases where a single certificate is used to sign client certificates and OCSP responses. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]
**OcspOverrideUrl** | **string** | The OCSP responder URL to use for overriding the one supplied in the client certificate. The URL must be complete with http:// included. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]
**OcspTimeout** | **int64** | The timeout in seconds to receive a response from the OCSP responder after sending a request or making the initial connection attempt. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]
**RevocationCheckEnabled** | **bool** | Indicates whether Certificate Authority revocation checking is enabled. Deprecated since 2.19. certAuthorities replaced by clientCertAuthorities and domainCertAuthorities. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
