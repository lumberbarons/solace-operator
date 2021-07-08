# Go API client for swagger

SEMP (starting in `v2`, see note 1) is a RESTful API for configuring, monitoring, and administering a Solace PubSub+ broker.  SEMP uses URIs to address manageable **resources** of the Solace PubSub+ broker. Resources are individual **objects**, **collections** of objects, or (exclusively in the action API) **actions**. This document applies to the following API:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Monitoring|/SEMP/v2/monitor|Querying operational parameters|See note 2    The following APIs are also available:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Action|/SEMP/v2/action|Performing actions|See note 2 Configuration|/SEMP/v2/config|Reading and writing config state|See note 2    Resources are always nouns, with individual objects being singular and collections being plural.  Objects within a collection are identified by an `obj-id`, which follows the collection name with the form `collection-name/obj-id`.  Actions within an object are identified by an `action-id`, which follows the object name with the form `obj-id/action-id`.  Some examples:  ``` /SEMP/v2/config/msgVpns                        ; MsgVpn collection /SEMP/v2/config/msgVpns/a                      ; MsgVpn object named \"a\" /SEMP/v2/config/msgVpns/a/queues               ; Queue collection in MsgVpn \"a\" /SEMP/v2/config/msgVpns/a/queues/b             ; Queue object named \"b\" in MsgVpn \"a\" /SEMP/v2/action/msgVpns/a/queues/b/startReplay ; Action that starts a replay on Queue \"b\" in MsgVpn \"a\" /SEMP/v2/monitor/msgVpns/a/clients             ; Client collection in MsgVpn \"a\" /SEMP/v2/monitor/msgVpns/a/clients/c           ; Client object named \"c\" in MsgVpn \"a\" ```  ## Collection Resources  Collections are unordered lists of objects (unless described as otherwise), and are described by JSON arrays. Each item in the array represents an object in the same manner as the individual object would normally be represented. In the configuration API, the creation of a new object is done through its collection resource.  ## Object and Action Resources  Objects are composed of attributes, actions, collections, and other objects. They are described by JSON objects as name/value pairs. The collections and actions of an object are not contained directly in the object's JSON content; rather the content includes an attribute containing a URI which points to the collections and actions. These contained resources must be managed through this URI. At a minimum, every object has one or more identifying attributes, and its own `uri` attribute which contains the URI pointing to itself.  Actions are also composed of attributes, and are described by JSON objects as name/value pairs. Unlike objects, however, they are not members of a collection and cannot be retrieved, only performed. Actions only exist in the action API.  Attributes in an object or action may have any combination of the following properties:   Property|Meaning|Comments :---|:---|:--- Identifying|Attribute is involved in unique identification of the object, and appears in its URI| Required|Attribute must be provided in the request| Read-Only|Attribute can only be read, not written.|See note 3 Write-Only|Attribute can only be written, not read, unless the attribute is also opaque|See the documentation for the opaque property Requires-Disable|Attribute can only be changed when object is disabled| Deprecated|Attribute is deprecated, and will disappear in the next SEMP version| Opaque|Attribute can be set or retrieved in opaque form when the `opaquePassword` query parameter is present|See the `opaquePassword` query parameter documentation    In some requests, certain attributes may only be provided in certain combinations with other attributes:   Relationship|Meaning :---|:--- Requires|Attribute may only be changed by a request if a particular attribute or combination of attributes is also provided in the request Conflicts|Attribute may only be provided in a request if a particular attribute or combination of attributes is not also provided in the request    In the monitoring API, any non-identifying attribute may not be returned in a GET.  ## HTTP Methods  The following HTTP methods manipulate resources in accordance with these general principles. Note that some methods are only used in certain APIs:   Method|Resource|Meaning|Request Body|Response Body|Missing Request Attributes :---|:---|:---|:---|:---|:--- POST|Collection|Create object|Initial attribute values|Object attributes and metadata|Set to default PUT|Object|Create or replace object (see note 5)|New attribute values|Object attributes and metadata|Set to default, with certain exceptions (see note 4) PUT|Action|Performs action|Action arguments|Action metadata|N/A PATCH|Object|Update object|New attribute values|Object attributes and metadata|unchanged DELETE|Object|Delete object|Empty|Object metadata|N/A GET|Object|Get object|Empty|Object attributes and metadata|N/A GET|Collection|Get collection|Empty|Object attributes and collection metadata|N/A    ## Common Query Parameters  The following are some common query parameters that are supported by many method/URI combinations. Individual URIs may document additional parameters. Note that multiple query parameters can be used together in a single URI, separated by the ampersand character. For example:  ``` ; Request for the MsgVpns collection using two hypothetical query parameters ; \"q1\" and \"q2\" with values \"val1\" and \"val2\" respectively /SEMP/v2/monitor/msgVpns?q1=val1&q2=val2 ```  ### select  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. Use this query parameter to limit the size of the returned data for each returned object, return only those fields that are desired, or exclude fields that are not desired.  The value of `select` is a comma-separated list of attribute names. If the list contains attribute names that are not prefaced by `-`, only those attributes are included in the response. If the list contains attribute names that are prefaced by `-`, those attributes are excluded from the response. If the list contains both types, then the difference of the first set of attributes and the second set of attributes is returned. If the list is empty (i.e. `select=`), no attributes are returned.  All attributes that are prefaced by `-` must follow all attributes that are not prefaced by `-`. In addition, each attribute name in the list must match at least one attribute in the object.  Names may include the `*` wildcard (zero or more characters). Nested attribute names are supported using periods (e.g. `parentName.childName`).  Some examples:  ``` ; List of all MsgVpn names /SEMP/v2/monitor/msgVpns?select=msgVpnName ; List of all MsgVpn and their attributes except for their names /SEMP/v2/monitor/msgVpns?select=-msgVpnName ; Authentication attributes of MsgVpn \"finance\" /SEMP/v2/monitor/msgVpns/finance?select=authentication* ; All attributes of MsgVpn \"finance\" except for authentication attributes /SEMP/v2/monitor/msgVpns/finance?select=-authentication* ; Access related attributes of Queue \"orderQ\" of MsgVpn \"finance\" /SEMP/v2/monitor/msgVpns/finance/queues/orderQ?select=owner,permission ```  ### where  Include in the response only objects where certain conditions are true. Use this query parameter to limit which objects are returned to those whose attribute values meet the given conditions.  The value of `where` is a comma-separated list of expressions. All expressions must be true for the object to be included in the response. Each expression takes the form:  ``` expression  = attribute-name OP value OP          = '==' | '!=' | '&lt;' | '&gt;' | '&lt;=' | '&gt;=' ```  `value` may be a number, string, `true`, or `false`, as appropriate for the type of `attribute-name`. Greater-than and less-than comparisons only work for numbers. A `*` in a string `value` is interpreted as a wildcard (zero or more characters). Some examples:  ``` ; Only enabled MsgVpns /SEMP/v2/monitor/msgVpns?where=enabled==true ; Only MsgVpns using basic non-LDAP authentication /SEMP/v2/monitor/msgVpns?where=authenticationBasicEnabled==true,authenticationBasicType!=ldap ; Only MsgVpns that allow more than 100 client connections /SEMP/v2/monitor/msgVpns?where=maxConnectionCount>100 ; Only MsgVpns with msgVpnName starting with \"B\": /SEMP/v2/monitor/msgVpns?where=msgVpnName==B* ```  ### count  Limit the count of objects in the response. This can be useful to limit the size of the response for large collections. The minimum value for `count` is `1` and the default is `10`. There is also a per-collection maximum value to limit request handling time. For example:  ``` ; Up to 25 MsgVpns /SEMP/v2/monitor/msgVpns?count=25 ```  ### cursor  The cursor, or position, for the next page of objects. Cursors are opaque data that should not be created or interpreted by SEMP clients, and should only be used as described below.  When a request is made for a collection and there may be additional objects available for retrieval that are not included in the initial response, the response will include a `cursorQuery` field containing a cursor. The value of this field can be specified in the `cursor` query parameter of a subsequent request to retrieve the next page of objects. For convenience, an appropriate URI is constructed automatically by the broker and included in the `nextPageUri` field of the response. This URI can be used directly to retrieve the next page of objects.  ### opaquePassword  Attributes with the opaque property are also write-only and so cannot normally be retrieved in a GET. However, when a password is provided in the `opaquePassword` query parameter, attributes with the opaque property are retrieved in a GET in opaque form, encrypted with this password. The query parameter can also be used on a POST, PATCH, or PUT to set opaque attributes using opaque attribute values retrieved in a GET, so long as:  1. the same password that was used to retrieve the opaque attribute values is provided; and  2. the broker to which the request is being sent has the same major and minor SEMP version as the broker that produced the opaque attribute values.  The password provided in the query parameter must be a minimum of 8 characters and a maximum of 128 characters.  The query parameter can only be used in the configuration API, and only over HTTPS.  ## Authentication  When a client makes its first SEMPv2 request, it must supply a username and password using HTTP Basic authentication.  If authentication is successful, the broker returns a cookie containing a session key. The client can omit the username and password from subsequent requests, because the broker now uses the session cookie for authentication instead. When the session expires or is deleted, the client must provide the username and password again, and the broker creates a new session.  There are a limited number of session slots available on the broker. The broker returns 529 No SEMP Session Available if it is not able to allocate a session. For this reason, all clients that use SEMPv2 should support cookies.  If certain attributes—such as a user's password—are changed, the broker automatically deletes the affected sessions. These attributes are documented below. However, changes in external user configuration data stored on a RADIUS or LDAP server do not trigger the broker to delete the associated session(s), therefore you must do this manually, if required.  A client can retrieve its current session information using the /about/user endpoint, delete its own session using the /about/user/logout endpoint, and manage all sessions using the /sessions endpoint.  ## Help  Visit [our website](https://solace.com) to learn more about Solace.  You can also download the SEMP API specifications by clicking [here](https://solace.com/downloads/).  If you need additional support, please contact us at [support@solace.com](mailto:support@solace.com).  ## Notes  Note|Description :---:|:--- 1|This specification defines SEMP starting in \"v2\", and not the original SEMP \"v1\" interface. Request and response formats between \"v1\" and \"v2\" are entirely incompatible, although both protocols share a common port configuration on the Solace PubSub+ broker. They are differentiated by the initial portion of the URI path, one of either \"/SEMP/\" or \"/SEMP/v2/\" 2|This API is partially implemented. Only a subset of all objects are available. 3|Read-only attributes may appear in POST and PUT/PATCH requests. However, if a read-only attribute is not marked as identifying, it will be ignored during a PUT/PATCH. 4|On a PUT, if the SEMP user is not authorized to modify the attribute, its value is left unchanged rather than set to default. In addition, the values of write-only attributes are not set to their defaults on a PUT, except in the following two cases: there is a mutual requires relationship with another non-write-only attribute, both attributes are absent from the request, and the non-write-only attribute is not currently set to its default value; or the attribute is also opaque and the `opaquePassword` query parameter is provided in the request. 5|On a PUT, if the object does not exist, it is created first.  

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.21
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [http://www.solace.com](http://www.solace.com)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AboutApi* | [**GetAbout**](docs/AboutApi.md#getabout) | **Get** /about | Get an About object.
*AboutApi* | [**GetAboutApi**](docs/AboutApi.md#getaboutapi) | **Get** /about/api | Get an API Description object.
*AboutApi* | [**GetAboutUser**](docs/AboutApi.md#getaboutuser) | **Get** /about/user | Get a User object.
*AboutApi* | [**GetAboutUserMsgVpn**](docs/AboutApi.md#getaboutusermsgvpn) | **Get** /about/user/msgVpns/{msgVpnName} | Get a User Message VPN object.
*AboutApi* | [**GetAboutUserMsgVpns**](docs/AboutApi.md#getaboutusermsgvpns) | **Get** /about/user/msgVpns | Get a list of User Message VPN objects.
*AclProfileApi* | [**GetMsgVpnAclProfile**](docs/AclProfileApi.md#getmsgvpnaclprofile) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Get an ACL Profile object.
*AclProfileApi* | [**GetMsgVpnAclProfileClientConnectException**](docs/AclProfileApi.md#getmsgvpnaclprofileclientconnectexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress} | Get a Client Connect Exception object.
*AclProfileApi* | [**GetMsgVpnAclProfileClientConnectExceptions**](docs/AclProfileApi.md#getmsgvpnaclprofileclientconnectexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions | Get a list of Client Connect Exception objects.
*AclProfileApi* | [**GetMsgVpnAclProfilePublishException**](docs/AclProfileApi.md#getmsgvpnaclprofilepublishexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic} | Get a Publish Topic Exception object.
*AclProfileApi* | [**GetMsgVpnAclProfilePublishExceptions**](docs/AclProfileApi.md#getmsgvpnaclprofilepublishexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions | Get a list of Publish Topic Exception objects.
*AclProfileApi* | [**GetMsgVpnAclProfilePublishTopicException**](docs/AclProfileApi.md#getmsgvpnaclprofilepublishtopicexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions/{publishTopicExceptionSyntax},{publishTopicException} | Get a Publish Topic Exception object.
*AclProfileApi* | [**GetMsgVpnAclProfilePublishTopicExceptions**](docs/AclProfileApi.md#getmsgvpnaclprofilepublishtopicexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions | Get a list of Publish Topic Exception objects.
*AclProfileApi* | [**GetMsgVpnAclProfileSubscribeException**](docs/AclProfileApi.md#getmsgvpnaclprofilesubscribeexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic} | Get a Subscribe Topic Exception object.
*AclProfileApi* | [**GetMsgVpnAclProfileSubscribeExceptions**](docs/AclProfileApi.md#getmsgvpnaclprofilesubscribeexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions | Get a list of Subscribe Topic Exception objects.
*AclProfileApi* | [**GetMsgVpnAclProfileSubscribeShareNameException**](docs/AclProfileApi.md#getmsgvpnaclprofilesubscribesharenameexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions/{subscribeShareNameExceptionSyntax},{subscribeShareNameException} | Get a Subscribe Share Name Exception object.
*AclProfileApi* | [**GetMsgVpnAclProfileSubscribeShareNameExceptions**](docs/AclProfileApi.md#getmsgvpnaclprofilesubscribesharenameexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions | Get a list of Subscribe Share Name Exception objects.
*AclProfileApi* | [**GetMsgVpnAclProfileSubscribeTopicException**](docs/AclProfileApi.md#getmsgvpnaclprofilesubscribetopicexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions/{subscribeTopicExceptionSyntax},{subscribeTopicException} | Get a Subscribe Topic Exception object.
*AclProfileApi* | [**GetMsgVpnAclProfileSubscribeTopicExceptions**](docs/AclProfileApi.md#getmsgvpnaclprofilesubscribetopicexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions | Get a list of Subscribe Topic Exception objects.
*AclProfileApi* | [**GetMsgVpnAclProfiles**](docs/AclProfileApi.md#getmsgvpnaclprofiles) | **Get** /msgVpns/{msgVpnName}/aclProfiles | Get a list of ACL Profile objects.
*AllApi* | [**GetAbout**](docs/AllApi.md#getabout) | **Get** /about | Get an About object.
*AllApi* | [**GetAboutApi**](docs/AllApi.md#getaboutapi) | **Get** /about/api | Get an API Description object.
*AllApi* | [**GetAboutUser**](docs/AllApi.md#getaboutuser) | **Get** /about/user | Get a User object.
*AllApi* | [**GetAboutUserMsgVpn**](docs/AllApi.md#getaboutusermsgvpn) | **Get** /about/user/msgVpns/{msgVpnName} | Get a User Message VPN object.
*AllApi* | [**GetAboutUserMsgVpns**](docs/AllApi.md#getaboutusermsgvpns) | **Get** /about/user/msgVpns | Get a list of User Message VPN objects.
*AllApi* | [**GetBroker**](docs/AllApi.md#getbroker) | **Get** / | Get a Broker object.
*AllApi* | [**GetCertAuthorities**](docs/AllApi.md#getcertauthorities) | **Get** /certAuthorities | Get a list of Certificate Authority objects.
*AllApi* | [**GetCertAuthority**](docs/AllApi.md#getcertauthority) | **Get** /certAuthorities/{certAuthorityName} | Get a Certificate Authority object.
*AllApi* | [**GetCertAuthorityOcspTlsTrustedCommonName**](docs/AllApi.md#getcertauthorityocsptlstrustedcommonname) | **Get** /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName} | Get an OCSP Responder Trusted Common Name object.
*AllApi* | [**GetCertAuthorityOcspTlsTrustedCommonNames**](docs/AllApi.md#getcertauthorityocsptlstrustedcommonnames) | **Get** /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames | Get a list of OCSP Responder Trusted Common Name objects.
*AllApi* | [**GetClientCertAuthorities**](docs/AllApi.md#getclientcertauthorities) | **Get** /clientCertAuthorities | Get a list of Client Certificate Authority objects.
*AllApi* | [**GetClientCertAuthority**](docs/AllApi.md#getclientcertauthority) | **Get** /clientCertAuthorities/{certAuthorityName} | Get a Client Certificate Authority object.
*AllApi* | [**GetClientCertAuthorityOcspTlsTrustedCommonName**](docs/AllApi.md#getclientcertauthorityocsptlstrustedcommonname) | **Get** /clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName} | Get an OCSP Responder Trusted Common Name object.
*AllApi* | [**GetClientCertAuthorityOcspTlsTrustedCommonNames**](docs/AllApi.md#getclientcertauthorityocsptlstrustedcommonnames) | **Get** /clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames | Get a list of OCSP Responder Trusted Common Name objects.
*AllApi* | [**GetDmrCluster**](docs/AllApi.md#getdmrcluster) | **Get** /dmrClusters/{dmrClusterName} | Get a Cluster object.
*AllApi* | [**GetDmrClusterLink**](docs/AllApi.md#getdmrclusterlink) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName} | Get a Link object.
*AllApi* | [**GetDmrClusterLinkChannel**](docs/AllApi.md#getdmrclusterlinkchannel) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/channels/{msgVpnName} | Get a Cluster Link Channels object.
*AllApi* | [**GetDmrClusterLinkChannels**](docs/AllApi.md#getdmrclusterlinkchannels) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/channels | Get a list of Cluster Link Channels objects.
*AllApi* | [**GetDmrClusterLinkRemoteAddress**](docs/AllApi.md#getdmrclusterlinkremoteaddress) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses/{remoteAddress} | Get a Remote Address object.
*AllApi* | [**GetDmrClusterLinkRemoteAddresses**](docs/AllApi.md#getdmrclusterlinkremoteaddresses) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses | Get a list of Remote Address objects.
*AllApi* | [**GetDmrClusterLinkTlsTrustedCommonName**](docs/AllApi.md#getdmrclusterlinktlstrustedcommonname) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
*AllApi* | [**GetDmrClusterLinkTlsTrustedCommonNames**](docs/AllApi.md#getdmrclusterlinktlstrustedcommonnames) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
*AllApi* | [**GetDmrClusterLinks**](docs/AllApi.md#getdmrclusterlinks) | **Get** /dmrClusters/{dmrClusterName}/links | Get a list of Link objects.
*AllApi* | [**GetDmrClusterTopologyIssue**](docs/AllApi.md#getdmrclustertopologyissue) | **Get** /dmrClusters/{dmrClusterName}/topologyIssues/{topologyIssue} | Get a Cluster Topology Issue object.
*AllApi* | [**GetDmrClusterTopologyIssues**](docs/AllApi.md#getdmrclustertopologyissues) | **Get** /dmrClusters/{dmrClusterName}/topologyIssues | Get a list of Cluster Topology Issue objects.
*AllApi* | [**GetDmrClusters**](docs/AllApi.md#getdmrclusters) | **Get** /dmrClusters | Get a list of Cluster objects.
*AllApi* | [**GetDomainCertAuthorities**](docs/AllApi.md#getdomaincertauthorities) | **Get** /domainCertAuthorities | Get a list of Domain Certificate Authority objects.
*AllApi* | [**GetDomainCertAuthority**](docs/AllApi.md#getdomaincertauthority) | **Get** /domainCertAuthorities/{certAuthorityName} | Get a Domain Certificate Authority object.
*AllApi* | [**GetMsgVpn**](docs/AllApi.md#getmsgvpn) | **Get** /msgVpns/{msgVpnName} | Get a Message VPN object.
*AllApi* | [**GetMsgVpnAclProfile**](docs/AllApi.md#getmsgvpnaclprofile) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Get an ACL Profile object.
*AllApi* | [**GetMsgVpnAclProfileClientConnectException**](docs/AllApi.md#getmsgvpnaclprofileclientconnectexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress} | Get a Client Connect Exception object.
*AllApi* | [**GetMsgVpnAclProfileClientConnectExceptions**](docs/AllApi.md#getmsgvpnaclprofileclientconnectexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions | Get a list of Client Connect Exception objects.
*AllApi* | [**GetMsgVpnAclProfilePublishException**](docs/AllApi.md#getmsgvpnaclprofilepublishexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic} | Get a Publish Topic Exception object.
*AllApi* | [**GetMsgVpnAclProfilePublishExceptions**](docs/AllApi.md#getmsgvpnaclprofilepublishexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions | Get a list of Publish Topic Exception objects.
*AllApi* | [**GetMsgVpnAclProfilePublishTopicException**](docs/AllApi.md#getmsgvpnaclprofilepublishtopicexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions/{publishTopicExceptionSyntax},{publishTopicException} | Get a Publish Topic Exception object.
*AllApi* | [**GetMsgVpnAclProfilePublishTopicExceptions**](docs/AllApi.md#getmsgvpnaclprofilepublishtopicexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions | Get a list of Publish Topic Exception objects.
*AllApi* | [**GetMsgVpnAclProfileSubscribeException**](docs/AllApi.md#getmsgvpnaclprofilesubscribeexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic} | Get a Subscribe Topic Exception object.
*AllApi* | [**GetMsgVpnAclProfileSubscribeExceptions**](docs/AllApi.md#getmsgvpnaclprofilesubscribeexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions | Get a list of Subscribe Topic Exception objects.
*AllApi* | [**GetMsgVpnAclProfileSubscribeShareNameException**](docs/AllApi.md#getmsgvpnaclprofilesubscribesharenameexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions/{subscribeShareNameExceptionSyntax},{subscribeShareNameException} | Get a Subscribe Share Name Exception object.
*AllApi* | [**GetMsgVpnAclProfileSubscribeShareNameExceptions**](docs/AllApi.md#getmsgvpnaclprofilesubscribesharenameexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions | Get a list of Subscribe Share Name Exception objects.
*AllApi* | [**GetMsgVpnAclProfileSubscribeTopicException**](docs/AllApi.md#getmsgvpnaclprofilesubscribetopicexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions/{subscribeTopicExceptionSyntax},{subscribeTopicException} | Get a Subscribe Topic Exception object.
*AllApi* | [**GetMsgVpnAclProfileSubscribeTopicExceptions**](docs/AllApi.md#getmsgvpnaclprofilesubscribetopicexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions | Get a list of Subscribe Topic Exception objects.
*AllApi* | [**GetMsgVpnAclProfiles**](docs/AllApi.md#getmsgvpnaclprofiles) | **Get** /msgVpns/{msgVpnName}/aclProfiles | Get a list of ACL Profile objects.
*AllApi* | [**GetMsgVpnAuthenticationOauthProvider**](docs/AllApi.md#getmsgvpnauthenticationoauthprovider) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Get an OAuth Provider object.
*AllApi* | [**GetMsgVpnAuthenticationOauthProviders**](docs/AllApi.md#getmsgvpnauthenticationoauthproviders) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders | Get a list of OAuth Provider objects.
*AllApi* | [**GetMsgVpnAuthorizationGroup**](docs/AllApi.md#getmsgvpnauthorizationgroup) | **Get** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Get an LDAP Authorization Group object.
*AllApi* | [**GetMsgVpnAuthorizationGroups**](docs/AllApi.md#getmsgvpnauthorizationgroups) | **Get** /msgVpns/{msgVpnName}/authorizationGroups | Get a list of LDAP Authorization Group objects.
*AllApi* | [**GetMsgVpnBridge**](docs/AllApi.md#getmsgvpnbridge) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Get a Bridge object.
*AllApi* | [**GetMsgVpnBridgeLocalSubscription**](docs/AllApi.md#getmsgvpnbridgelocalsubscription) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/localSubscriptions/{localSubscriptionTopic} | Get a Bridge Local Subscriptions object.
*AllApi* | [**GetMsgVpnBridgeLocalSubscriptions**](docs/AllApi.md#getmsgvpnbridgelocalsubscriptions) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/localSubscriptions | Get a list of Bridge Local Subscriptions objects.
*AllApi* | [**GetMsgVpnBridgeRemoteMsgVpn**](docs/AllApi.md#getmsgvpnbridgeremotemsgvpn) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Get a Remote Message VPN object.
*AllApi* | [**GetMsgVpnBridgeRemoteMsgVpns**](docs/AllApi.md#getmsgvpnbridgeremotemsgvpns) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns | Get a list of Remote Message VPN objects.
*AllApi* | [**GetMsgVpnBridgeRemoteSubscription**](docs/AllApi.md#getmsgvpnbridgeremotesubscription) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic} | Get a Remote Subscription object.
*AllApi* | [**GetMsgVpnBridgeRemoteSubscriptions**](docs/AllApi.md#getmsgvpnbridgeremotesubscriptions) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions | Get a list of Remote Subscription objects.
*AllApi* | [**GetMsgVpnBridgeTlsTrustedCommonName**](docs/AllApi.md#getmsgvpnbridgetlstrustedcommonname) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
*AllApi* | [**GetMsgVpnBridgeTlsTrustedCommonNames**](docs/AllApi.md#getmsgvpnbridgetlstrustedcommonnames) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
*AllApi* | [**GetMsgVpnBridges**](docs/AllApi.md#getmsgvpnbridges) | **Get** /msgVpns/{msgVpnName}/bridges | Get a list of Bridge objects.
*AllApi* | [**GetMsgVpnClient**](docs/AllApi.md#getmsgvpnclient) | **Get** /msgVpns/{msgVpnName}/clients/{clientName} | Get a Client object.
*AllApi* | [**GetMsgVpnClientConnection**](docs/AllApi.md#getmsgvpnclientconnection) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/connections/{clientAddress} | Get a Client Connection object.
*AllApi* | [**GetMsgVpnClientConnections**](docs/AllApi.md#getmsgvpnclientconnections) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/connections | Get a list of Client Connection objects.
*AllApi* | [**GetMsgVpnClientProfile**](docs/AllApi.md#getmsgvpnclientprofile) | **Get** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Get a Client Profile object.
*AllApi* | [**GetMsgVpnClientProfiles**](docs/AllApi.md#getmsgvpnclientprofiles) | **Get** /msgVpns/{msgVpnName}/clientProfiles | Get a list of Client Profile objects.
*AllApi* | [**GetMsgVpnClientRxFlow**](docs/AllApi.md#getmsgvpnclientrxflow) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/rxFlows/{flowId} | Get a Client Receive Flow object.
*AllApi* | [**GetMsgVpnClientRxFlows**](docs/AllApi.md#getmsgvpnclientrxflows) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/rxFlows | Get a list of Client Receive Flow objects.
*AllApi* | [**GetMsgVpnClientSubscription**](docs/AllApi.md#getmsgvpnclientsubscription) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/subscriptions/{subscriptionTopic} | Get a Client Subscription object.
*AllApi* | [**GetMsgVpnClientSubscriptions**](docs/AllApi.md#getmsgvpnclientsubscriptions) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/subscriptions | Get a list of Client Subscription objects.
*AllApi* | [**GetMsgVpnClientTransactedSession**](docs/AllApi.md#getmsgvpnclienttransactedsession) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions/{sessionName} | Get a Client Transacted Session object.
*AllApi* | [**GetMsgVpnClientTransactedSessions**](docs/AllApi.md#getmsgvpnclienttransactedsessions) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions | Get a list of Client Transacted Session objects.
*AllApi* | [**GetMsgVpnClientTxFlow**](docs/AllApi.md#getmsgvpnclienttxflow) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/txFlows/{flowId} | Get a Client Transmit Flow object.
*AllApi* | [**GetMsgVpnClientTxFlows**](docs/AllApi.md#getmsgvpnclienttxflows) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/txFlows | Get a list of Client Transmit Flow objects.
*AllApi* | [**GetMsgVpnClientUsername**](docs/AllApi.md#getmsgvpnclientusername) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Get a Client Username object.
*AllApi* | [**GetMsgVpnClientUsernames**](docs/AllApi.md#getmsgvpnclientusernames) | **Get** /msgVpns/{msgVpnName}/clientUsernames | Get a list of Client Username objects.
*AllApi* | [**GetMsgVpnClients**](docs/AllApi.md#getmsgvpnclients) | **Get** /msgVpns/{msgVpnName}/clients | Get a list of Client objects.
*AllApi* | [**GetMsgVpnConfigSyncRemoteNode**](docs/AllApi.md#getmsgvpnconfigsyncremotenode) | **Get** /msgVpns/{msgVpnName}/configSyncRemoteNodes/{remoteNodeName} | Get a Config Sync Remote Node object.
*AllApi* | [**GetMsgVpnConfigSyncRemoteNodes**](docs/AllApi.md#getmsgvpnconfigsyncremotenodes) | **Get** /msgVpns/{msgVpnName}/configSyncRemoteNodes | Get a list of Config Sync Remote Node objects.
*AllApi* | [**GetMsgVpnDistributedCache**](docs/AllApi.md#getmsgvpndistributedcache) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Get a Distributed Cache object.
*AllApi* | [**GetMsgVpnDistributedCacheCluster**](docs/AllApi.md#getmsgvpndistributedcachecluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Get a Cache Cluster object.
*AllApi* | [**GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](docs/AllApi.md#getmsgvpndistributedcacheclusterglobalcachinghomecluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName} | Get a Home Cache Cluster object.
*AllApi* | [**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](docs/AllApi.md#getmsgvpndistributedcacheclusterglobalcachinghomeclustertopicprefix) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix} | Get a Topic Prefix object.
*AllApi* | [**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes**](docs/AllApi.md#getmsgvpndistributedcacheclusterglobalcachinghomeclustertopicprefixes) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes | Get a list of Topic Prefix objects.
*AllApi* | [**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters**](docs/AllApi.md#getmsgvpndistributedcacheclusterglobalcachinghomeclusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters | Get a list of Home Cache Cluster objects.
*AllApi* | [**GetMsgVpnDistributedCacheClusterInstance**](docs/AllApi.md#getmsgvpndistributedcacheclusterinstance) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Get a Cache Instance object.
*AllApi* | [**GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster**](docs/AllApi.md#getmsgvpndistributedcacheclusterinstanceremoteglobalcachinghomecluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteGlobalCachingHomeClusters/{homeClusterName} | Get a Remote Home Cache Cluster object.
*AllApi* | [**GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters**](docs/AllApi.md#getmsgvpndistributedcacheclusterinstanceremoteglobalcachinghomeclusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteGlobalCachingHomeClusters | Get a list of Remote Home Cache Cluster objects.
*AllApi* | [**GetMsgVpnDistributedCacheClusterInstanceRemoteTopic**](docs/AllApi.md#getmsgvpndistributedcacheclusterinstanceremotetopic) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteTopics/{topic} | Get a Remote Topic object.
*AllApi* | [**GetMsgVpnDistributedCacheClusterInstanceRemoteTopics**](docs/AllApi.md#getmsgvpndistributedcacheclusterinstanceremotetopics) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteTopics | Get a list of Remote Topic objects.
*AllApi* | [**GetMsgVpnDistributedCacheClusterInstances**](docs/AllApi.md#getmsgvpndistributedcacheclusterinstances) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances | Get a list of Cache Instance objects.
*AllApi* | [**GetMsgVpnDistributedCacheClusterTopic**](docs/AllApi.md#getmsgvpndistributedcacheclustertopic) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics/{topic} | Get a Topic object.
*AllApi* | [**GetMsgVpnDistributedCacheClusterTopics**](docs/AllApi.md#getmsgvpndistributedcacheclustertopics) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics | Get a list of Topic objects.
*AllApi* | [**GetMsgVpnDistributedCacheClusters**](docs/AllApi.md#getmsgvpndistributedcacheclusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters | Get a list of Cache Cluster objects.
*AllApi* | [**GetMsgVpnDistributedCaches**](docs/AllApi.md#getmsgvpndistributedcaches) | **Get** /msgVpns/{msgVpnName}/distributedCaches | Get a list of Distributed Cache objects.
*AllApi* | [**GetMsgVpnDmrBridge**](docs/AllApi.md#getmsgvpndmrbridge) | **Get** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Get a DMR Bridge object.
*AllApi* | [**GetMsgVpnDmrBridges**](docs/AllApi.md#getmsgvpndmrbridges) | **Get** /msgVpns/{msgVpnName}/dmrBridges | Get a list of DMR Bridge objects.
*AllApi* | [**GetMsgVpnJndiConnectionFactories**](docs/AllApi.md#getmsgvpnjndiconnectionfactories) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories | Get a list of JNDI Connection Factory objects.
*AllApi* | [**GetMsgVpnJndiConnectionFactory**](docs/AllApi.md#getmsgvpnjndiconnectionfactory) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Get a JNDI Connection Factory object.
*AllApi* | [**GetMsgVpnJndiQueue**](docs/AllApi.md#getmsgvpnjndiqueue) | **Get** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Get a JNDI Queue object.
*AllApi* | [**GetMsgVpnJndiQueues**](docs/AllApi.md#getmsgvpnjndiqueues) | **Get** /msgVpns/{msgVpnName}/jndiQueues | Get a list of JNDI Queue objects.
*AllApi* | [**GetMsgVpnJndiTopic**](docs/AllApi.md#getmsgvpnjnditopic) | **Get** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Get a JNDI Topic object.
*AllApi* | [**GetMsgVpnJndiTopics**](docs/AllApi.md#getmsgvpnjnditopics) | **Get** /msgVpns/{msgVpnName}/jndiTopics | Get a list of JNDI Topic objects.
*AllApi* | [**GetMsgVpnMqttRetainCache**](docs/AllApi.md#getmsgvpnmqttretaincache) | **Get** /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName} | Get an MQTT Retain Cache object.
*AllApi* | [**GetMsgVpnMqttRetainCaches**](docs/AllApi.md#getmsgvpnmqttretaincaches) | **Get** /msgVpns/{msgVpnName}/mqttRetainCaches | Get a list of MQTT Retain Cache objects.
*AllApi* | [**GetMsgVpnMqttSession**](docs/AllApi.md#getmsgvpnmqttsession) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Get an MQTT Session object.
*AllApi* | [**GetMsgVpnMqttSessionSubscription**](docs/AllApi.md#getmsgvpnmqttsessionsubscription) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Get a Subscription object.
*AllApi* | [**GetMsgVpnMqttSessionSubscriptions**](docs/AllApi.md#getmsgvpnmqttsessionsubscriptions) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions | Get a list of Subscription objects.
*AllApi* | [**GetMsgVpnMqttSessions**](docs/AllApi.md#getmsgvpnmqttsessions) | **Get** /msgVpns/{msgVpnName}/mqttSessions | Get a list of MQTT Session objects.
*AllApi* | [**GetMsgVpnQueue**](docs/AllApi.md#getmsgvpnqueue) | **Get** /msgVpns/{msgVpnName}/queues/{queueName} | Get a Queue object.
*AllApi* | [**GetMsgVpnQueueMsg**](docs/AllApi.md#getmsgvpnqueuemsg) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs/{msgId} | Get a Queue Message object.
*AllApi* | [**GetMsgVpnQueueMsgs**](docs/AllApi.md#getmsgvpnqueuemsgs) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs | Get a list of Queue Message objects.
*AllApi* | [**GetMsgVpnQueuePriorities**](docs/AllApi.md#getmsgvpnqueuepriorities) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/priorities | Get a list of Queue Priority objects.
*AllApi* | [**GetMsgVpnQueuePriority**](docs/AllApi.md#getmsgvpnqueuepriority) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/priorities/{priority} | Get a Queue Priority object.
*AllApi* | [**GetMsgVpnQueueSubscription**](docs/AllApi.md#getmsgvpnqueuesubscription) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic} | Get a Queue Subscription object.
*AllApi* | [**GetMsgVpnQueueSubscriptions**](docs/AllApi.md#getmsgvpnqueuesubscriptions) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions | Get a list of Queue Subscription objects.
*AllApi* | [**GetMsgVpnQueueTemplate**](docs/AllApi.md#getmsgvpnqueuetemplate) | **Get** /msgVpns/{msgVpnName}/queueTemplates/{queueTemplateName} | Get a Queue Template object.
*AllApi* | [**GetMsgVpnQueueTemplates**](docs/AllApi.md#getmsgvpnqueuetemplates) | **Get** /msgVpns/{msgVpnName}/queueTemplates | Get a list of Queue Template objects.
*AllApi* | [**GetMsgVpnQueueTxFlow**](docs/AllApi.md#getmsgvpnqueuetxflow) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/txFlows/{flowId} | Get a Queue Transmit Flow object.
*AllApi* | [**GetMsgVpnQueueTxFlows**](docs/AllApi.md#getmsgvpnqueuetxflows) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/txFlows | Get a list of Queue Transmit Flow objects.
*AllApi* | [**GetMsgVpnQueues**](docs/AllApi.md#getmsgvpnqueues) | **Get** /msgVpns/{msgVpnName}/queues | Get a list of Queue objects.
*AllApi* | [**GetMsgVpnReplayLog**](docs/AllApi.md#getmsgvpnreplaylog) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Get a Replay Log object.
*AllApi* | [**GetMsgVpnReplayLogMsg**](docs/AllApi.md#getmsgvpnreplaylogmsg) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/msgs/{msgId} | Get a Message object.
*AllApi* | [**GetMsgVpnReplayLogMsgs**](docs/AllApi.md#getmsgvpnreplaylogmsgs) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/msgs | Get a list of Message objects.
*AllApi* | [**GetMsgVpnReplayLogs**](docs/AllApi.md#getmsgvpnreplaylogs) | **Get** /msgVpns/{msgVpnName}/replayLogs | Get a list of Replay Log objects.
*AllApi* | [**GetMsgVpnReplicatedTopic**](docs/AllApi.md#getmsgvpnreplicatedtopic) | **Get** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Get a Replicated Topic object.
*AllApi* | [**GetMsgVpnReplicatedTopics**](docs/AllApi.md#getmsgvpnreplicatedtopics) | **Get** /msgVpns/{msgVpnName}/replicatedTopics | Get a list of Replicated Topic objects.
*AllApi* | [**GetMsgVpnRestDeliveryPoint**](docs/AllApi.md#getmsgvpnrestdeliverypoint) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Get a REST Delivery Point object.
*AllApi* | [**GetMsgVpnRestDeliveryPointQueueBinding**](docs/AllApi.md#getmsgvpnrestdeliverypointqueuebinding) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Get a Queue Binding object.
*AllApi* | [**GetMsgVpnRestDeliveryPointQueueBindings**](docs/AllApi.md#getmsgvpnrestdeliverypointqueuebindings) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings | Get a list of Queue Binding objects.
*AllApi* | [**GetMsgVpnRestDeliveryPointRestConsumer**](docs/AllApi.md#getmsgvpnrestdeliverypointrestconsumer) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Get a REST Consumer object.
*AllApi* | [**GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**](docs/AllApi.md#getmsgvpnrestdeliverypointrestconsumeroauthjwtclaim) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims/{oauthJwtClaimName} | Get a Claim object.
*AllApi* | [**GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims**](docs/AllApi.md#getmsgvpnrestdeliverypointrestconsumeroauthjwtclaims) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims | Get a list of Claim objects.
*AllApi* | [**GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**](docs/AllApi.md#getmsgvpnrestdeliverypointrestconsumertlstrustedcommonname) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
*AllApi* | [**GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames**](docs/AllApi.md#getmsgvpnrestdeliverypointrestconsumertlstrustedcommonnames) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
*AllApi* | [**GetMsgVpnRestDeliveryPointRestConsumers**](docs/AllApi.md#getmsgvpnrestdeliverypointrestconsumers) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Get a list of REST Consumer objects.
*AllApi* | [**GetMsgVpnRestDeliveryPoints**](docs/AllApi.md#getmsgvpnrestdeliverypoints) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints | Get a list of REST Delivery Point objects.
*AllApi* | [**GetMsgVpnTopicEndpoint**](docs/AllApi.md#getmsgvpntopicendpoint) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Get a Topic Endpoint object.
*AllApi* | [**GetMsgVpnTopicEndpointMsg**](docs/AllApi.md#getmsgvpntopicendpointmsg) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId} | Get a Topic Endpoint Message object.
*AllApi* | [**GetMsgVpnTopicEndpointMsgs**](docs/AllApi.md#getmsgvpntopicendpointmsgs) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs | Get a list of Topic Endpoint Message objects.
*AllApi* | [**GetMsgVpnTopicEndpointPriorities**](docs/AllApi.md#getmsgvpntopicendpointpriorities) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/priorities | Get a list of Topic Endpoint Priority objects.
*AllApi* | [**GetMsgVpnTopicEndpointPriority**](docs/AllApi.md#getmsgvpntopicendpointpriority) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/priorities/{priority} | Get a Topic Endpoint Priority object.
*AllApi* | [**GetMsgVpnTopicEndpointTemplate**](docs/AllApi.md#getmsgvpntopicendpointtemplate) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Get a Topic Endpoint Template object.
*AllApi* | [**GetMsgVpnTopicEndpointTemplates**](docs/AllApi.md#getmsgvpntopicendpointtemplates) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates | Get a list of Topic Endpoint Template objects.
*AllApi* | [**GetMsgVpnTopicEndpointTxFlow**](docs/AllApi.md#getmsgvpntopicendpointtxflow) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/txFlows/{flowId} | Get a Topic Endpoint Transmit Flow object.
*AllApi* | [**GetMsgVpnTopicEndpointTxFlows**](docs/AllApi.md#getmsgvpntopicendpointtxflows) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/txFlows | Get a list of Topic Endpoint Transmit Flow objects.
*AllApi* | [**GetMsgVpnTopicEndpoints**](docs/AllApi.md#getmsgvpntopicendpoints) | **Get** /msgVpns/{msgVpnName}/topicEndpoints | Get a list of Topic Endpoint objects.
*AllApi* | [**GetMsgVpnTransaction**](docs/AllApi.md#getmsgvpntransaction) | **Get** /msgVpns/{msgVpnName}/transactions/{xid} | Get a Replicated Local Transaction or XA Transaction object.
*AllApi* | [**GetMsgVpnTransactionConsumerMsg**](docs/AllApi.md#getmsgvpntransactionconsumermsg) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/consumerMsgs/{msgId} | Get a Transaction Consumer Message object.
*AllApi* | [**GetMsgVpnTransactionConsumerMsgs**](docs/AllApi.md#getmsgvpntransactionconsumermsgs) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/consumerMsgs | Get a list of Transaction Consumer Message objects.
*AllApi* | [**GetMsgVpnTransactionPublisherMsg**](docs/AllApi.md#getmsgvpntransactionpublishermsg) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/publisherMsgs/{msgId} | Get a Transaction Publisher Message object.
*AllApi* | [**GetMsgVpnTransactionPublisherMsgs**](docs/AllApi.md#getmsgvpntransactionpublishermsgs) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/publisherMsgs | Get a list of Transaction Publisher Message objects.
*AllApi* | [**GetMsgVpnTransactions**](docs/AllApi.md#getmsgvpntransactions) | **Get** /msgVpns/{msgVpnName}/transactions | Get a list of Replicated Local Transaction or XA Transaction objects.
*AllApi* | [**GetMsgVpns**](docs/AllApi.md#getmsgvpns) | **Get** /msgVpns | Get a list of Message VPN objects.
*AllApi* | [**GetSession**](docs/AllApi.md#getsession) | **Get** /sessions/{sessionUsername},{sessionId} | Get a Session object.
*AllApi* | [**GetSessions**](docs/AllApi.md#getsessions) | **Get** /sessions | Get a list of Session objects.
*AllApi* | [**GetStandardDomainCertAuthorities**](docs/AllApi.md#getstandarddomaincertauthorities) | **Get** /standardDomainCertAuthorities | Get a list of Standard Domain Certificate Authority objects.
*AllApi* | [**GetStandardDomainCertAuthority**](docs/AllApi.md#getstandarddomaincertauthority) | **Get** /standardDomainCertAuthorities/{certAuthorityName} | Get a Standard Domain Certificate Authority object.
*AllApi* | [**GetVirtualHostname**](docs/AllApi.md#getvirtualhostname) | **Get** /virtualHostnames/{virtualHostname} | Get a Virtual Hostname object.
*AllApi* | [**GetVirtualHostnames**](docs/AllApi.md#getvirtualhostnames) | **Get** /virtualHostnames | Get a list of Virtual Hostname objects.
*AuthenticationOauthProviderApi* | [**GetMsgVpnAuthenticationOauthProvider**](docs/AuthenticationOauthProviderApi.md#getmsgvpnauthenticationoauthprovider) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Get an OAuth Provider object.
*AuthenticationOauthProviderApi* | [**GetMsgVpnAuthenticationOauthProviders**](docs/AuthenticationOauthProviderApi.md#getmsgvpnauthenticationoauthproviders) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders | Get a list of OAuth Provider objects.
*AuthorizationGroupApi* | [**GetMsgVpnAuthorizationGroup**](docs/AuthorizationGroupApi.md#getmsgvpnauthorizationgroup) | **Get** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Get an LDAP Authorization Group object.
*AuthorizationGroupApi* | [**GetMsgVpnAuthorizationGroups**](docs/AuthorizationGroupApi.md#getmsgvpnauthorizationgroups) | **Get** /msgVpns/{msgVpnName}/authorizationGroups | Get a list of LDAP Authorization Group objects.
*BridgeApi* | [**GetMsgVpnBridge**](docs/BridgeApi.md#getmsgvpnbridge) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Get a Bridge object.
*BridgeApi* | [**GetMsgVpnBridgeLocalSubscription**](docs/BridgeApi.md#getmsgvpnbridgelocalsubscription) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/localSubscriptions/{localSubscriptionTopic} | Get a Bridge Local Subscriptions object.
*BridgeApi* | [**GetMsgVpnBridgeLocalSubscriptions**](docs/BridgeApi.md#getmsgvpnbridgelocalsubscriptions) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/localSubscriptions | Get a list of Bridge Local Subscriptions objects.
*BridgeApi* | [**GetMsgVpnBridgeRemoteMsgVpn**](docs/BridgeApi.md#getmsgvpnbridgeremotemsgvpn) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Get a Remote Message VPN object.
*BridgeApi* | [**GetMsgVpnBridgeRemoteMsgVpns**](docs/BridgeApi.md#getmsgvpnbridgeremotemsgvpns) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns | Get a list of Remote Message VPN objects.
*BridgeApi* | [**GetMsgVpnBridgeRemoteSubscription**](docs/BridgeApi.md#getmsgvpnbridgeremotesubscription) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic} | Get a Remote Subscription object.
*BridgeApi* | [**GetMsgVpnBridgeRemoteSubscriptions**](docs/BridgeApi.md#getmsgvpnbridgeremotesubscriptions) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions | Get a list of Remote Subscription objects.
*BridgeApi* | [**GetMsgVpnBridgeTlsTrustedCommonName**](docs/BridgeApi.md#getmsgvpnbridgetlstrustedcommonname) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
*BridgeApi* | [**GetMsgVpnBridgeTlsTrustedCommonNames**](docs/BridgeApi.md#getmsgvpnbridgetlstrustedcommonnames) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
*BridgeApi* | [**GetMsgVpnBridges**](docs/BridgeApi.md#getmsgvpnbridges) | **Get** /msgVpns/{msgVpnName}/bridges | Get a list of Bridge objects.
*CertAuthorityApi* | [**GetCertAuthorities**](docs/CertAuthorityApi.md#getcertauthorities) | **Get** /certAuthorities | Get a list of Certificate Authority objects.
*CertAuthorityApi* | [**GetCertAuthority**](docs/CertAuthorityApi.md#getcertauthority) | **Get** /certAuthorities/{certAuthorityName} | Get a Certificate Authority object.
*CertAuthorityApi* | [**GetCertAuthorityOcspTlsTrustedCommonName**](docs/CertAuthorityApi.md#getcertauthorityocsptlstrustedcommonname) | **Get** /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName} | Get an OCSP Responder Trusted Common Name object.
*CertAuthorityApi* | [**GetCertAuthorityOcspTlsTrustedCommonNames**](docs/CertAuthorityApi.md#getcertauthorityocsptlstrustedcommonnames) | **Get** /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames | Get a list of OCSP Responder Trusted Common Name objects.
*ClientApi* | [**GetMsgVpnClient**](docs/ClientApi.md#getmsgvpnclient) | **Get** /msgVpns/{msgVpnName}/clients/{clientName} | Get a Client object.
*ClientApi* | [**GetMsgVpnClientConnection**](docs/ClientApi.md#getmsgvpnclientconnection) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/connections/{clientAddress} | Get a Client Connection object.
*ClientApi* | [**GetMsgVpnClientConnections**](docs/ClientApi.md#getmsgvpnclientconnections) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/connections | Get a list of Client Connection objects.
*ClientApi* | [**GetMsgVpnClientRxFlow**](docs/ClientApi.md#getmsgvpnclientrxflow) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/rxFlows/{flowId} | Get a Client Receive Flow object.
*ClientApi* | [**GetMsgVpnClientRxFlows**](docs/ClientApi.md#getmsgvpnclientrxflows) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/rxFlows | Get a list of Client Receive Flow objects.
*ClientApi* | [**GetMsgVpnClientSubscription**](docs/ClientApi.md#getmsgvpnclientsubscription) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/subscriptions/{subscriptionTopic} | Get a Client Subscription object.
*ClientApi* | [**GetMsgVpnClientSubscriptions**](docs/ClientApi.md#getmsgvpnclientsubscriptions) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/subscriptions | Get a list of Client Subscription objects.
*ClientApi* | [**GetMsgVpnClientTransactedSession**](docs/ClientApi.md#getmsgvpnclienttransactedsession) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions/{sessionName} | Get a Client Transacted Session object.
*ClientApi* | [**GetMsgVpnClientTransactedSessions**](docs/ClientApi.md#getmsgvpnclienttransactedsessions) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions | Get a list of Client Transacted Session objects.
*ClientApi* | [**GetMsgVpnClientTxFlow**](docs/ClientApi.md#getmsgvpnclienttxflow) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/txFlows/{flowId} | Get a Client Transmit Flow object.
*ClientApi* | [**GetMsgVpnClientTxFlows**](docs/ClientApi.md#getmsgvpnclienttxflows) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/txFlows | Get a list of Client Transmit Flow objects.
*ClientApi* | [**GetMsgVpnClients**](docs/ClientApi.md#getmsgvpnclients) | **Get** /msgVpns/{msgVpnName}/clients | Get a list of Client objects.
*ClientCertAuthorityApi* | [**GetClientCertAuthorities**](docs/ClientCertAuthorityApi.md#getclientcertauthorities) | **Get** /clientCertAuthorities | Get a list of Client Certificate Authority objects.
*ClientCertAuthorityApi* | [**GetClientCertAuthority**](docs/ClientCertAuthorityApi.md#getclientcertauthority) | **Get** /clientCertAuthorities/{certAuthorityName} | Get a Client Certificate Authority object.
*ClientCertAuthorityApi* | [**GetClientCertAuthorityOcspTlsTrustedCommonName**](docs/ClientCertAuthorityApi.md#getclientcertauthorityocsptlstrustedcommonname) | **Get** /clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName} | Get an OCSP Responder Trusted Common Name object.
*ClientCertAuthorityApi* | [**GetClientCertAuthorityOcspTlsTrustedCommonNames**](docs/ClientCertAuthorityApi.md#getclientcertauthorityocsptlstrustedcommonnames) | **Get** /clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames | Get a list of OCSP Responder Trusted Common Name objects.
*ClientProfileApi* | [**GetMsgVpnClientProfile**](docs/ClientProfileApi.md#getmsgvpnclientprofile) | **Get** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Get a Client Profile object.
*ClientProfileApi* | [**GetMsgVpnClientProfiles**](docs/ClientProfileApi.md#getmsgvpnclientprofiles) | **Get** /msgVpns/{msgVpnName}/clientProfiles | Get a list of Client Profile objects.
*ClientUsernameApi* | [**GetMsgVpnClientUsername**](docs/ClientUsernameApi.md#getmsgvpnclientusername) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Get a Client Username object.
*ClientUsernameApi* | [**GetMsgVpnClientUsernames**](docs/ClientUsernameApi.md#getmsgvpnclientusernames) | **Get** /msgVpns/{msgVpnName}/clientUsernames | Get a list of Client Username objects.
*ConfigSyncRemoteNodeApi* | [**GetMsgVpnConfigSyncRemoteNode**](docs/ConfigSyncRemoteNodeApi.md#getmsgvpnconfigsyncremotenode) | **Get** /msgVpns/{msgVpnName}/configSyncRemoteNodes/{remoteNodeName} | Get a Config Sync Remote Node object.
*ConfigSyncRemoteNodeApi* | [**GetMsgVpnConfigSyncRemoteNodes**](docs/ConfigSyncRemoteNodeApi.md#getmsgvpnconfigsyncremotenodes) | **Get** /msgVpns/{msgVpnName}/configSyncRemoteNodes | Get a list of Config Sync Remote Node objects.
*DistributedCacheApi* | [**GetMsgVpnDistributedCache**](docs/DistributedCacheApi.md#getmsgvpndistributedcache) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Get a Distributed Cache object.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheCluster**](docs/DistributedCacheApi.md#getmsgvpndistributedcachecluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Get a Cache Cluster object.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](docs/DistributedCacheApi.md#getmsgvpndistributedcacheclusterglobalcachinghomecluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName} | Get a Home Cache Cluster object.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](docs/DistributedCacheApi.md#getmsgvpndistributedcacheclusterglobalcachinghomeclustertopicprefix) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix} | Get a Topic Prefix object.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes**](docs/DistributedCacheApi.md#getmsgvpndistributedcacheclusterglobalcachinghomeclustertopicprefixes) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes | Get a list of Topic Prefix objects.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters**](docs/DistributedCacheApi.md#getmsgvpndistributedcacheclusterglobalcachinghomeclusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters | Get a list of Home Cache Cluster objects.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheClusterInstance**](docs/DistributedCacheApi.md#getmsgvpndistributedcacheclusterinstance) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Get a Cache Instance object.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster**](docs/DistributedCacheApi.md#getmsgvpndistributedcacheclusterinstanceremoteglobalcachinghomecluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteGlobalCachingHomeClusters/{homeClusterName} | Get a Remote Home Cache Cluster object.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters**](docs/DistributedCacheApi.md#getmsgvpndistributedcacheclusterinstanceremoteglobalcachinghomeclusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteGlobalCachingHomeClusters | Get a list of Remote Home Cache Cluster objects.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheClusterInstanceRemoteTopic**](docs/DistributedCacheApi.md#getmsgvpndistributedcacheclusterinstanceremotetopic) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteTopics/{topic} | Get a Remote Topic object.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheClusterInstanceRemoteTopics**](docs/DistributedCacheApi.md#getmsgvpndistributedcacheclusterinstanceremotetopics) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteTopics | Get a list of Remote Topic objects.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheClusterInstances**](docs/DistributedCacheApi.md#getmsgvpndistributedcacheclusterinstances) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances | Get a list of Cache Instance objects.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheClusterTopic**](docs/DistributedCacheApi.md#getmsgvpndistributedcacheclustertopic) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics/{topic} | Get a Topic object.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheClusterTopics**](docs/DistributedCacheApi.md#getmsgvpndistributedcacheclustertopics) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics | Get a list of Topic objects.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheClusters**](docs/DistributedCacheApi.md#getmsgvpndistributedcacheclusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters | Get a list of Cache Cluster objects.
*DistributedCacheApi* | [**GetMsgVpnDistributedCaches**](docs/DistributedCacheApi.md#getmsgvpndistributedcaches) | **Get** /msgVpns/{msgVpnName}/distributedCaches | Get a list of Distributed Cache objects.
*DmrBridgeApi* | [**GetMsgVpnDmrBridge**](docs/DmrBridgeApi.md#getmsgvpndmrbridge) | **Get** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Get a DMR Bridge object.
*DmrBridgeApi* | [**GetMsgVpnDmrBridges**](docs/DmrBridgeApi.md#getmsgvpndmrbridges) | **Get** /msgVpns/{msgVpnName}/dmrBridges | Get a list of DMR Bridge objects.
*DmrClusterApi* | [**GetDmrCluster**](docs/DmrClusterApi.md#getdmrcluster) | **Get** /dmrClusters/{dmrClusterName} | Get a Cluster object.
*DmrClusterApi* | [**GetDmrClusterLink**](docs/DmrClusterApi.md#getdmrclusterlink) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName} | Get a Link object.
*DmrClusterApi* | [**GetDmrClusterLinkChannel**](docs/DmrClusterApi.md#getdmrclusterlinkchannel) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/channels/{msgVpnName} | Get a Cluster Link Channels object.
*DmrClusterApi* | [**GetDmrClusterLinkChannels**](docs/DmrClusterApi.md#getdmrclusterlinkchannels) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/channels | Get a list of Cluster Link Channels objects.
*DmrClusterApi* | [**GetDmrClusterLinkRemoteAddress**](docs/DmrClusterApi.md#getdmrclusterlinkremoteaddress) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses/{remoteAddress} | Get a Remote Address object.
*DmrClusterApi* | [**GetDmrClusterLinkRemoteAddresses**](docs/DmrClusterApi.md#getdmrclusterlinkremoteaddresses) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses | Get a list of Remote Address objects.
*DmrClusterApi* | [**GetDmrClusterLinkTlsTrustedCommonName**](docs/DmrClusterApi.md#getdmrclusterlinktlstrustedcommonname) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
*DmrClusterApi* | [**GetDmrClusterLinkTlsTrustedCommonNames**](docs/DmrClusterApi.md#getdmrclusterlinktlstrustedcommonnames) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
*DmrClusterApi* | [**GetDmrClusterLinks**](docs/DmrClusterApi.md#getdmrclusterlinks) | **Get** /dmrClusters/{dmrClusterName}/links | Get a list of Link objects.
*DmrClusterApi* | [**GetDmrClusterTopologyIssue**](docs/DmrClusterApi.md#getdmrclustertopologyissue) | **Get** /dmrClusters/{dmrClusterName}/topologyIssues/{topologyIssue} | Get a Cluster Topology Issue object.
*DmrClusterApi* | [**GetDmrClusterTopologyIssues**](docs/DmrClusterApi.md#getdmrclustertopologyissues) | **Get** /dmrClusters/{dmrClusterName}/topologyIssues | Get a list of Cluster Topology Issue objects.
*DmrClusterApi* | [**GetDmrClusters**](docs/DmrClusterApi.md#getdmrclusters) | **Get** /dmrClusters | Get a list of Cluster objects.
*DomainCertAuthorityApi* | [**GetDomainCertAuthorities**](docs/DomainCertAuthorityApi.md#getdomaincertauthorities) | **Get** /domainCertAuthorities | Get a list of Domain Certificate Authority objects.
*DomainCertAuthorityApi* | [**GetDomainCertAuthority**](docs/DomainCertAuthorityApi.md#getdomaincertauthority) | **Get** /domainCertAuthorities/{certAuthorityName} | Get a Domain Certificate Authority object.
*JndiApi* | [**GetMsgVpnJndiConnectionFactories**](docs/JndiApi.md#getmsgvpnjndiconnectionfactories) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories | Get a list of JNDI Connection Factory objects.
*JndiApi* | [**GetMsgVpnJndiConnectionFactory**](docs/JndiApi.md#getmsgvpnjndiconnectionfactory) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Get a JNDI Connection Factory object.
*JndiApi* | [**GetMsgVpnJndiQueue**](docs/JndiApi.md#getmsgvpnjndiqueue) | **Get** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Get a JNDI Queue object.
*JndiApi* | [**GetMsgVpnJndiQueues**](docs/JndiApi.md#getmsgvpnjndiqueues) | **Get** /msgVpns/{msgVpnName}/jndiQueues | Get a list of JNDI Queue objects.
*JndiApi* | [**GetMsgVpnJndiTopic**](docs/JndiApi.md#getmsgvpnjnditopic) | **Get** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Get a JNDI Topic object.
*JndiApi* | [**GetMsgVpnJndiTopics**](docs/JndiApi.md#getmsgvpnjnditopics) | **Get** /msgVpns/{msgVpnName}/jndiTopics | Get a list of JNDI Topic objects.
*MqttRetainCacheApi* | [**GetMsgVpnMqttRetainCache**](docs/MqttRetainCacheApi.md#getmsgvpnmqttretaincache) | **Get** /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName} | Get an MQTT Retain Cache object.
*MqttRetainCacheApi* | [**GetMsgVpnMqttRetainCaches**](docs/MqttRetainCacheApi.md#getmsgvpnmqttretaincaches) | **Get** /msgVpns/{msgVpnName}/mqttRetainCaches | Get a list of MQTT Retain Cache objects.
*MqttSessionApi* | [**GetMsgVpnMqttSession**](docs/MqttSessionApi.md#getmsgvpnmqttsession) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Get an MQTT Session object.
*MqttSessionApi* | [**GetMsgVpnMqttSessionSubscription**](docs/MqttSessionApi.md#getmsgvpnmqttsessionsubscription) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Get a Subscription object.
*MqttSessionApi* | [**GetMsgVpnMqttSessionSubscriptions**](docs/MqttSessionApi.md#getmsgvpnmqttsessionsubscriptions) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions | Get a list of Subscription objects.
*MqttSessionApi* | [**GetMsgVpnMqttSessions**](docs/MqttSessionApi.md#getmsgvpnmqttsessions) | **Get** /msgVpns/{msgVpnName}/mqttSessions | Get a list of MQTT Session objects.
*MsgVpnApi* | [**GetMsgVpn**](docs/MsgVpnApi.md#getmsgvpn) | **Get** /msgVpns/{msgVpnName} | Get a Message VPN object.
*MsgVpnApi* | [**GetMsgVpnAclProfile**](docs/MsgVpnApi.md#getmsgvpnaclprofile) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Get an ACL Profile object.
*MsgVpnApi* | [**GetMsgVpnAclProfileClientConnectException**](docs/MsgVpnApi.md#getmsgvpnaclprofileclientconnectexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress} | Get a Client Connect Exception object.
*MsgVpnApi* | [**GetMsgVpnAclProfileClientConnectExceptions**](docs/MsgVpnApi.md#getmsgvpnaclprofileclientconnectexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions | Get a list of Client Connect Exception objects.
*MsgVpnApi* | [**GetMsgVpnAclProfilePublishException**](docs/MsgVpnApi.md#getmsgvpnaclprofilepublishexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic} | Get a Publish Topic Exception object.
*MsgVpnApi* | [**GetMsgVpnAclProfilePublishExceptions**](docs/MsgVpnApi.md#getmsgvpnaclprofilepublishexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions | Get a list of Publish Topic Exception objects.
*MsgVpnApi* | [**GetMsgVpnAclProfilePublishTopicException**](docs/MsgVpnApi.md#getmsgvpnaclprofilepublishtopicexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions/{publishTopicExceptionSyntax},{publishTopicException} | Get a Publish Topic Exception object.
*MsgVpnApi* | [**GetMsgVpnAclProfilePublishTopicExceptions**](docs/MsgVpnApi.md#getmsgvpnaclprofilepublishtopicexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions | Get a list of Publish Topic Exception objects.
*MsgVpnApi* | [**GetMsgVpnAclProfileSubscribeException**](docs/MsgVpnApi.md#getmsgvpnaclprofilesubscribeexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic} | Get a Subscribe Topic Exception object.
*MsgVpnApi* | [**GetMsgVpnAclProfileSubscribeExceptions**](docs/MsgVpnApi.md#getmsgvpnaclprofilesubscribeexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions | Get a list of Subscribe Topic Exception objects.
*MsgVpnApi* | [**GetMsgVpnAclProfileSubscribeShareNameException**](docs/MsgVpnApi.md#getmsgvpnaclprofilesubscribesharenameexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions/{subscribeShareNameExceptionSyntax},{subscribeShareNameException} | Get a Subscribe Share Name Exception object.
*MsgVpnApi* | [**GetMsgVpnAclProfileSubscribeShareNameExceptions**](docs/MsgVpnApi.md#getmsgvpnaclprofilesubscribesharenameexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions | Get a list of Subscribe Share Name Exception objects.
*MsgVpnApi* | [**GetMsgVpnAclProfileSubscribeTopicException**](docs/MsgVpnApi.md#getmsgvpnaclprofilesubscribetopicexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions/{subscribeTopicExceptionSyntax},{subscribeTopicException} | Get a Subscribe Topic Exception object.
*MsgVpnApi* | [**GetMsgVpnAclProfileSubscribeTopicExceptions**](docs/MsgVpnApi.md#getmsgvpnaclprofilesubscribetopicexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions | Get a list of Subscribe Topic Exception objects.
*MsgVpnApi* | [**GetMsgVpnAclProfiles**](docs/MsgVpnApi.md#getmsgvpnaclprofiles) | **Get** /msgVpns/{msgVpnName}/aclProfiles | Get a list of ACL Profile objects.
*MsgVpnApi* | [**GetMsgVpnAuthenticationOauthProvider**](docs/MsgVpnApi.md#getmsgvpnauthenticationoauthprovider) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Get an OAuth Provider object.
*MsgVpnApi* | [**GetMsgVpnAuthenticationOauthProviders**](docs/MsgVpnApi.md#getmsgvpnauthenticationoauthproviders) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders | Get a list of OAuth Provider objects.
*MsgVpnApi* | [**GetMsgVpnAuthorizationGroup**](docs/MsgVpnApi.md#getmsgvpnauthorizationgroup) | **Get** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Get an LDAP Authorization Group object.
*MsgVpnApi* | [**GetMsgVpnAuthorizationGroups**](docs/MsgVpnApi.md#getmsgvpnauthorizationgroups) | **Get** /msgVpns/{msgVpnName}/authorizationGroups | Get a list of LDAP Authorization Group objects.
*MsgVpnApi* | [**GetMsgVpnBridge**](docs/MsgVpnApi.md#getmsgvpnbridge) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Get a Bridge object.
*MsgVpnApi* | [**GetMsgVpnBridgeLocalSubscription**](docs/MsgVpnApi.md#getmsgvpnbridgelocalsubscription) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/localSubscriptions/{localSubscriptionTopic} | Get a Bridge Local Subscriptions object.
*MsgVpnApi* | [**GetMsgVpnBridgeLocalSubscriptions**](docs/MsgVpnApi.md#getmsgvpnbridgelocalsubscriptions) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/localSubscriptions | Get a list of Bridge Local Subscriptions objects.
*MsgVpnApi* | [**GetMsgVpnBridgeRemoteMsgVpn**](docs/MsgVpnApi.md#getmsgvpnbridgeremotemsgvpn) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Get a Remote Message VPN object.
*MsgVpnApi* | [**GetMsgVpnBridgeRemoteMsgVpns**](docs/MsgVpnApi.md#getmsgvpnbridgeremotemsgvpns) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns | Get a list of Remote Message VPN objects.
*MsgVpnApi* | [**GetMsgVpnBridgeRemoteSubscription**](docs/MsgVpnApi.md#getmsgvpnbridgeremotesubscription) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic} | Get a Remote Subscription object.
*MsgVpnApi* | [**GetMsgVpnBridgeRemoteSubscriptions**](docs/MsgVpnApi.md#getmsgvpnbridgeremotesubscriptions) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions | Get a list of Remote Subscription objects.
*MsgVpnApi* | [**GetMsgVpnBridgeTlsTrustedCommonName**](docs/MsgVpnApi.md#getmsgvpnbridgetlstrustedcommonname) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
*MsgVpnApi* | [**GetMsgVpnBridgeTlsTrustedCommonNames**](docs/MsgVpnApi.md#getmsgvpnbridgetlstrustedcommonnames) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
*MsgVpnApi* | [**GetMsgVpnBridges**](docs/MsgVpnApi.md#getmsgvpnbridges) | **Get** /msgVpns/{msgVpnName}/bridges | Get a list of Bridge objects.
*MsgVpnApi* | [**GetMsgVpnClient**](docs/MsgVpnApi.md#getmsgvpnclient) | **Get** /msgVpns/{msgVpnName}/clients/{clientName} | Get a Client object.
*MsgVpnApi* | [**GetMsgVpnClientConnection**](docs/MsgVpnApi.md#getmsgvpnclientconnection) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/connections/{clientAddress} | Get a Client Connection object.
*MsgVpnApi* | [**GetMsgVpnClientConnections**](docs/MsgVpnApi.md#getmsgvpnclientconnections) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/connections | Get a list of Client Connection objects.
*MsgVpnApi* | [**GetMsgVpnClientProfile**](docs/MsgVpnApi.md#getmsgvpnclientprofile) | **Get** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Get a Client Profile object.
*MsgVpnApi* | [**GetMsgVpnClientProfiles**](docs/MsgVpnApi.md#getmsgvpnclientprofiles) | **Get** /msgVpns/{msgVpnName}/clientProfiles | Get a list of Client Profile objects.
*MsgVpnApi* | [**GetMsgVpnClientRxFlow**](docs/MsgVpnApi.md#getmsgvpnclientrxflow) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/rxFlows/{flowId} | Get a Client Receive Flow object.
*MsgVpnApi* | [**GetMsgVpnClientRxFlows**](docs/MsgVpnApi.md#getmsgvpnclientrxflows) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/rxFlows | Get a list of Client Receive Flow objects.
*MsgVpnApi* | [**GetMsgVpnClientSubscription**](docs/MsgVpnApi.md#getmsgvpnclientsubscription) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/subscriptions/{subscriptionTopic} | Get a Client Subscription object.
*MsgVpnApi* | [**GetMsgVpnClientSubscriptions**](docs/MsgVpnApi.md#getmsgvpnclientsubscriptions) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/subscriptions | Get a list of Client Subscription objects.
*MsgVpnApi* | [**GetMsgVpnClientTransactedSession**](docs/MsgVpnApi.md#getmsgvpnclienttransactedsession) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions/{sessionName} | Get a Client Transacted Session object.
*MsgVpnApi* | [**GetMsgVpnClientTransactedSessions**](docs/MsgVpnApi.md#getmsgvpnclienttransactedsessions) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions | Get a list of Client Transacted Session objects.
*MsgVpnApi* | [**GetMsgVpnClientTxFlow**](docs/MsgVpnApi.md#getmsgvpnclienttxflow) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/txFlows/{flowId} | Get a Client Transmit Flow object.
*MsgVpnApi* | [**GetMsgVpnClientTxFlows**](docs/MsgVpnApi.md#getmsgvpnclienttxflows) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/txFlows | Get a list of Client Transmit Flow objects.
*MsgVpnApi* | [**GetMsgVpnClientUsername**](docs/MsgVpnApi.md#getmsgvpnclientusername) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Get a Client Username object.
*MsgVpnApi* | [**GetMsgVpnClientUsernames**](docs/MsgVpnApi.md#getmsgvpnclientusernames) | **Get** /msgVpns/{msgVpnName}/clientUsernames | Get a list of Client Username objects.
*MsgVpnApi* | [**GetMsgVpnClients**](docs/MsgVpnApi.md#getmsgvpnclients) | **Get** /msgVpns/{msgVpnName}/clients | Get a list of Client objects.
*MsgVpnApi* | [**GetMsgVpnConfigSyncRemoteNode**](docs/MsgVpnApi.md#getmsgvpnconfigsyncremotenode) | **Get** /msgVpns/{msgVpnName}/configSyncRemoteNodes/{remoteNodeName} | Get a Config Sync Remote Node object.
*MsgVpnApi* | [**GetMsgVpnConfigSyncRemoteNodes**](docs/MsgVpnApi.md#getmsgvpnconfigsyncremotenodes) | **Get** /msgVpns/{msgVpnName}/configSyncRemoteNodes | Get a list of Config Sync Remote Node objects.
*MsgVpnApi* | [**GetMsgVpnDistributedCache**](docs/MsgVpnApi.md#getmsgvpndistributedcache) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Get a Distributed Cache object.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheCluster**](docs/MsgVpnApi.md#getmsgvpndistributedcachecluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Get a Cache Cluster object.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](docs/MsgVpnApi.md#getmsgvpndistributedcacheclusterglobalcachinghomecluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName} | Get a Home Cache Cluster object.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](docs/MsgVpnApi.md#getmsgvpndistributedcacheclusterglobalcachinghomeclustertopicprefix) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix} | Get a Topic Prefix object.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes**](docs/MsgVpnApi.md#getmsgvpndistributedcacheclusterglobalcachinghomeclustertopicprefixes) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes | Get a list of Topic Prefix objects.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters**](docs/MsgVpnApi.md#getmsgvpndistributedcacheclusterglobalcachinghomeclusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters | Get a list of Home Cache Cluster objects.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheClusterInstance**](docs/MsgVpnApi.md#getmsgvpndistributedcacheclusterinstance) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Get a Cache Instance object.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster**](docs/MsgVpnApi.md#getmsgvpndistributedcacheclusterinstanceremoteglobalcachinghomecluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteGlobalCachingHomeClusters/{homeClusterName} | Get a Remote Home Cache Cluster object.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters**](docs/MsgVpnApi.md#getmsgvpndistributedcacheclusterinstanceremoteglobalcachinghomeclusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteGlobalCachingHomeClusters | Get a list of Remote Home Cache Cluster objects.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheClusterInstanceRemoteTopic**](docs/MsgVpnApi.md#getmsgvpndistributedcacheclusterinstanceremotetopic) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteTopics/{topic} | Get a Remote Topic object.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheClusterInstanceRemoteTopics**](docs/MsgVpnApi.md#getmsgvpndistributedcacheclusterinstanceremotetopics) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteTopics | Get a list of Remote Topic objects.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheClusterInstances**](docs/MsgVpnApi.md#getmsgvpndistributedcacheclusterinstances) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances | Get a list of Cache Instance objects.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheClusterTopic**](docs/MsgVpnApi.md#getmsgvpndistributedcacheclustertopic) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics/{topic} | Get a Topic object.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheClusterTopics**](docs/MsgVpnApi.md#getmsgvpndistributedcacheclustertopics) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics | Get a list of Topic objects.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheClusters**](docs/MsgVpnApi.md#getmsgvpndistributedcacheclusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters | Get a list of Cache Cluster objects.
*MsgVpnApi* | [**GetMsgVpnDistributedCaches**](docs/MsgVpnApi.md#getmsgvpndistributedcaches) | **Get** /msgVpns/{msgVpnName}/distributedCaches | Get a list of Distributed Cache objects.
*MsgVpnApi* | [**GetMsgVpnDmrBridge**](docs/MsgVpnApi.md#getmsgvpndmrbridge) | **Get** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Get a DMR Bridge object.
*MsgVpnApi* | [**GetMsgVpnDmrBridges**](docs/MsgVpnApi.md#getmsgvpndmrbridges) | **Get** /msgVpns/{msgVpnName}/dmrBridges | Get a list of DMR Bridge objects.
*MsgVpnApi* | [**GetMsgVpnJndiConnectionFactories**](docs/MsgVpnApi.md#getmsgvpnjndiconnectionfactories) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories | Get a list of JNDI Connection Factory objects.
*MsgVpnApi* | [**GetMsgVpnJndiConnectionFactory**](docs/MsgVpnApi.md#getmsgvpnjndiconnectionfactory) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Get a JNDI Connection Factory object.
*MsgVpnApi* | [**GetMsgVpnJndiQueue**](docs/MsgVpnApi.md#getmsgvpnjndiqueue) | **Get** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Get a JNDI Queue object.
*MsgVpnApi* | [**GetMsgVpnJndiQueues**](docs/MsgVpnApi.md#getmsgvpnjndiqueues) | **Get** /msgVpns/{msgVpnName}/jndiQueues | Get a list of JNDI Queue objects.
*MsgVpnApi* | [**GetMsgVpnJndiTopic**](docs/MsgVpnApi.md#getmsgvpnjnditopic) | **Get** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Get a JNDI Topic object.
*MsgVpnApi* | [**GetMsgVpnJndiTopics**](docs/MsgVpnApi.md#getmsgvpnjnditopics) | **Get** /msgVpns/{msgVpnName}/jndiTopics | Get a list of JNDI Topic objects.
*MsgVpnApi* | [**GetMsgVpnMqttRetainCache**](docs/MsgVpnApi.md#getmsgvpnmqttretaincache) | **Get** /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName} | Get an MQTT Retain Cache object.
*MsgVpnApi* | [**GetMsgVpnMqttRetainCaches**](docs/MsgVpnApi.md#getmsgvpnmqttretaincaches) | **Get** /msgVpns/{msgVpnName}/mqttRetainCaches | Get a list of MQTT Retain Cache objects.
*MsgVpnApi* | [**GetMsgVpnMqttSession**](docs/MsgVpnApi.md#getmsgvpnmqttsession) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Get an MQTT Session object.
*MsgVpnApi* | [**GetMsgVpnMqttSessionSubscription**](docs/MsgVpnApi.md#getmsgvpnmqttsessionsubscription) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Get a Subscription object.
*MsgVpnApi* | [**GetMsgVpnMqttSessionSubscriptions**](docs/MsgVpnApi.md#getmsgvpnmqttsessionsubscriptions) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions | Get a list of Subscription objects.
*MsgVpnApi* | [**GetMsgVpnMqttSessions**](docs/MsgVpnApi.md#getmsgvpnmqttsessions) | **Get** /msgVpns/{msgVpnName}/mqttSessions | Get a list of MQTT Session objects.
*MsgVpnApi* | [**GetMsgVpnQueue**](docs/MsgVpnApi.md#getmsgvpnqueue) | **Get** /msgVpns/{msgVpnName}/queues/{queueName} | Get a Queue object.
*MsgVpnApi* | [**GetMsgVpnQueueMsg**](docs/MsgVpnApi.md#getmsgvpnqueuemsg) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs/{msgId} | Get a Queue Message object.
*MsgVpnApi* | [**GetMsgVpnQueueMsgs**](docs/MsgVpnApi.md#getmsgvpnqueuemsgs) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs | Get a list of Queue Message objects.
*MsgVpnApi* | [**GetMsgVpnQueuePriorities**](docs/MsgVpnApi.md#getmsgvpnqueuepriorities) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/priorities | Get a list of Queue Priority objects.
*MsgVpnApi* | [**GetMsgVpnQueuePriority**](docs/MsgVpnApi.md#getmsgvpnqueuepriority) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/priorities/{priority} | Get a Queue Priority object.
*MsgVpnApi* | [**GetMsgVpnQueueSubscription**](docs/MsgVpnApi.md#getmsgvpnqueuesubscription) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic} | Get a Queue Subscription object.
*MsgVpnApi* | [**GetMsgVpnQueueSubscriptions**](docs/MsgVpnApi.md#getmsgvpnqueuesubscriptions) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions | Get a list of Queue Subscription objects.
*MsgVpnApi* | [**GetMsgVpnQueueTemplate**](docs/MsgVpnApi.md#getmsgvpnqueuetemplate) | **Get** /msgVpns/{msgVpnName}/queueTemplates/{queueTemplateName} | Get a Queue Template object.
*MsgVpnApi* | [**GetMsgVpnQueueTemplates**](docs/MsgVpnApi.md#getmsgvpnqueuetemplates) | **Get** /msgVpns/{msgVpnName}/queueTemplates | Get a list of Queue Template objects.
*MsgVpnApi* | [**GetMsgVpnQueueTxFlow**](docs/MsgVpnApi.md#getmsgvpnqueuetxflow) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/txFlows/{flowId} | Get a Queue Transmit Flow object.
*MsgVpnApi* | [**GetMsgVpnQueueTxFlows**](docs/MsgVpnApi.md#getmsgvpnqueuetxflows) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/txFlows | Get a list of Queue Transmit Flow objects.
*MsgVpnApi* | [**GetMsgVpnQueues**](docs/MsgVpnApi.md#getmsgvpnqueues) | **Get** /msgVpns/{msgVpnName}/queues | Get a list of Queue objects.
*MsgVpnApi* | [**GetMsgVpnReplayLog**](docs/MsgVpnApi.md#getmsgvpnreplaylog) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Get a Replay Log object.
*MsgVpnApi* | [**GetMsgVpnReplayLogMsg**](docs/MsgVpnApi.md#getmsgvpnreplaylogmsg) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/msgs/{msgId} | Get a Message object.
*MsgVpnApi* | [**GetMsgVpnReplayLogMsgs**](docs/MsgVpnApi.md#getmsgvpnreplaylogmsgs) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/msgs | Get a list of Message objects.
*MsgVpnApi* | [**GetMsgVpnReplayLogs**](docs/MsgVpnApi.md#getmsgvpnreplaylogs) | **Get** /msgVpns/{msgVpnName}/replayLogs | Get a list of Replay Log objects.
*MsgVpnApi* | [**GetMsgVpnReplicatedTopic**](docs/MsgVpnApi.md#getmsgvpnreplicatedtopic) | **Get** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Get a Replicated Topic object.
*MsgVpnApi* | [**GetMsgVpnReplicatedTopics**](docs/MsgVpnApi.md#getmsgvpnreplicatedtopics) | **Get** /msgVpns/{msgVpnName}/replicatedTopics | Get a list of Replicated Topic objects.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPoint**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypoint) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Get a REST Delivery Point object.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPointQueueBinding**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypointqueuebinding) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Get a Queue Binding object.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPointQueueBindings**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypointqueuebindings) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings | Get a list of Queue Binding objects.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPointRestConsumer**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypointrestconsumer) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Get a REST Consumer object.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypointrestconsumeroauthjwtclaim) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims/{oauthJwtClaimName} | Get a Claim object.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypointrestconsumeroauthjwtclaims) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims | Get a list of Claim objects.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypointrestconsumertlstrustedcommonname) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypointrestconsumertlstrustedcommonnames) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPointRestConsumers**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypointrestconsumers) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Get a list of REST Consumer objects.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPoints**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypoints) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints | Get a list of REST Delivery Point objects.
*MsgVpnApi* | [**GetMsgVpnTopicEndpoint**](docs/MsgVpnApi.md#getmsgvpntopicendpoint) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Get a Topic Endpoint object.
*MsgVpnApi* | [**GetMsgVpnTopicEndpointMsg**](docs/MsgVpnApi.md#getmsgvpntopicendpointmsg) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId} | Get a Topic Endpoint Message object.
*MsgVpnApi* | [**GetMsgVpnTopicEndpointMsgs**](docs/MsgVpnApi.md#getmsgvpntopicendpointmsgs) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs | Get a list of Topic Endpoint Message objects.
*MsgVpnApi* | [**GetMsgVpnTopicEndpointPriorities**](docs/MsgVpnApi.md#getmsgvpntopicendpointpriorities) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/priorities | Get a list of Topic Endpoint Priority objects.
*MsgVpnApi* | [**GetMsgVpnTopicEndpointPriority**](docs/MsgVpnApi.md#getmsgvpntopicendpointpriority) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/priorities/{priority} | Get a Topic Endpoint Priority object.
*MsgVpnApi* | [**GetMsgVpnTopicEndpointTemplate**](docs/MsgVpnApi.md#getmsgvpntopicendpointtemplate) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Get a Topic Endpoint Template object.
*MsgVpnApi* | [**GetMsgVpnTopicEndpointTemplates**](docs/MsgVpnApi.md#getmsgvpntopicendpointtemplates) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates | Get a list of Topic Endpoint Template objects.
*MsgVpnApi* | [**GetMsgVpnTopicEndpointTxFlow**](docs/MsgVpnApi.md#getmsgvpntopicendpointtxflow) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/txFlows/{flowId} | Get a Topic Endpoint Transmit Flow object.
*MsgVpnApi* | [**GetMsgVpnTopicEndpointTxFlows**](docs/MsgVpnApi.md#getmsgvpntopicendpointtxflows) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/txFlows | Get a list of Topic Endpoint Transmit Flow objects.
*MsgVpnApi* | [**GetMsgVpnTopicEndpoints**](docs/MsgVpnApi.md#getmsgvpntopicendpoints) | **Get** /msgVpns/{msgVpnName}/topicEndpoints | Get a list of Topic Endpoint objects.
*MsgVpnApi* | [**GetMsgVpnTransaction**](docs/MsgVpnApi.md#getmsgvpntransaction) | **Get** /msgVpns/{msgVpnName}/transactions/{xid} | Get a Replicated Local Transaction or XA Transaction object.
*MsgVpnApi* | [**GetMsgVpnTransactionConsumerMsg**](docs/MsgVpnApi.md#getmsgvpntransactionconsumermsg) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/consumerMsgs/{msgId} | Get a Transaction Consumer Message object.
*MsgVpnApi* | [**GetMsgVpnTransactionConsumerMsgs**](docs/MsgVpnApi.md#getmsgvpntransactionconsumermsgs) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/consumerMsgs | Get a list of Transaction Consumer Message objects.
*MsgVpnApi* | [**GetMsgVpnTransactionPublisherMsg**](docs/MsgVpnApi.md#getmsgvpntransactionpublishermsg) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/publisherMsgs/{msgId} | Get a Transaction Publisher Message object.
*MsgVpnApi* | [**GetMsgVpnTransactionPublisherMsgs**](docs/MsgVpnApi.md#getmsgvpntransactionpublishermsgs) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/publisherMsgs | Get a list of Transaction Publisher Message objects.
*MsgVpnApi* | [**GetMsgVpnTransactions**](docs/MsgVpnApi.md#getmsgvpntransactions) | **Get** /msgVpns/{msgVpnName}/transactions | Get a list of Replicated Local Transaction or XA Transaction objects.
*MsgVpnApi* | [**GetMsgVpns**](docs/MsgVpnApi.md#getmsgvpns) | **Get** /msgVpns | Get a list of Message VPN objects.
*QueueApi* | [**GetMsgVpnQueue**](docs/QueueApi.md#getmsgvpnqueue) | **Get** /msgVpns/{msgVpnName}/queues/{queueName} | Get a Queue object.
*QueueApi* | [**GetMsgVpnQueueMsg**](docs/QueueApi.md#getmsgvpnqueuemsg) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs/{msgId} | Get a Queue Message object.
*QueueApi* | [**GetMsgVpnQueueMsgs**](docs/QueueApi.md#getmsgvpnqueuemsgs) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs | Get a list of Queue Message objects.
*QueueApi* | [**GetMsgVpnQueuePriorities**](docs/QueueApi.md#getmsgvpnqueuepriorities) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/priorities | Get a list of Queue Priority objects.
*QueueApi* | [**GetMsgVpnQueuePriority**](docs/QueueApi.md#getmsgvpnqueuepriority) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/priorities/{priority} | Get a Queue Priority object.
*QueueApi* | [**GetMsgVpnQueueSubscription**](docs/QueueApi.md#getmsgvpnqueuesubscription) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic} | Get a Queue Subscription object.
*QueueApi* | [**GetMsgVpnQueueSubscriptions**](docs/QueueApi.md#getmsgvpnqueuesubscriptions) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions | Get a list of Queue Subscription objects.
*QueueApi* | [**GetMsgVpnQueueTxFlow**](docs/QueueApi.md#getmsgvpnqueuetxflow) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/txFlows/{flowId} | Get a Queue Transmit Flow object.
*QueueApi* | [**GetMsgVpnQueueTxFlows**](docs/QueueApi.md#getmsgvpnqueuetxflows) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/txFlows | Get a list of Queue Transmit Flow objects.
*QueueApi* | [**GetMsgVpnQueues**](docs/QueueApi.md#getmsgvpnqueues) | **Get** /msgVpns/{msgVpnName}/queues | Get a list of Queue objects.
*QueueTemplateApi* | [**GetMsgVpnQueueTemplate**](docs/QueueTemplateApi.md#getmsgvpnqueuetemplate) | **Get** /msgVpns/{msgVpnName}/queueTemplates/{queueTemplateName} | Get a Queue Template object.
*QueueTemplateApi* | [**GetMsgVpnQueueTemplates**](docs/QueueTemplateApi.md#getmsgvpnqueuetemplates) | **Get** /msgVpns/{msgVpnName}/queueTemplates | Get a list of Queue Template objects.
*ReplayLogApi* | [**GetMsgVpnReplayLog**](docs/ReplayLogApi.md#getmsgvpnreplaylog) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Get a Replay Log object.
*ReplayLogApi* | [**GetMsgVpnReplayLogMsg**](docs/ReplayLogApi.md#getmsgvpnreplaylogmsg) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/msgs/{msgId} | Get a Message object.
*ReplayLogApi* | [**GetMsgVpnReplayLogMsgs**](docs/ReplayLogApi.md#getmsgvpnreplaylogmsgs) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/msgs | Get a list of Message objects.
*ReplayLogApi* | [**GetMsgVpnReplayLogs**](docs/ReplayLogApi.md#getmsgvpnreplaylogs) | **Get** /msgVpns/{msgVpnName}/replayLogs | Get a list of Replay Log objects.
*ReplicatedTopicApi* | [**GetMsgVpnReplicatedTopic**](docs/ReplicatedTopicApi.md#getmsgvpnreplicatedtopic) | **Get** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Get a Replicated Topic object.
*ReplicatedTopicApi* | [**GetMsgVpnReplicatedTopics**](docs/ReplicatedTopicApi.md#getmsgvpnreplicatedtopics) | **Get** /msgVpns/{msgVpnName}/replicatedTopics | Get a list of Replicated Topic objects.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPoint**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypoint) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Get a REST Delivery Point object.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPointQueueBinding**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypointqueuebinding) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Get a Queue Binding object.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPointQueueBindings**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypointqueuebindings) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings | Get a list of Queue Binding objects.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPointRestConsumer**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypointrestconsumer) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Get a REST Consumer object.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypointrestconsumeroauthjwtclaim) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims/{oauthJwtClaimName} | Get a Claim object.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypointrestconsumeroauthjwtclaims) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims | Get a list of Claim objects.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypointrestconsumertlstrustedcommonname) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypointrestconsumertlstrustedcommonnames) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPointRestConsumers**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypointrestconsumers) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Get a list of REST Consumer objects.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPoints**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypoints) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints | Get a list of REST Delivery Point objects.
*SessionApi* | [**GetSession**](docs/SessionApi.md#getsession) | **Get** /sessions/{sessionUsername},{sessionId} | Get a Session object.
*SessionApi* | [**GetSessions**](docs/SessionApi.md#getsessions) | **Get** /sessions | Get a list of Session objects.
*StandardDomainCertAuthorityApi* | [**GetStandardDomainCertAuthorities**](docs/StandardDomainCertAuthorityApi.md#getstandarddomaincertauthorities) | **Get** /standardDomainCertAuthorities | Get a list of Standard Domain Certificate Authority objects.
*StandardDomainCertAuthorityApi* | [**GetStandardDomainCertAuthority**](docs/StandardDomainCertAuthorityApi.md#getstandarddomaincertauthority) | **Get** /standardDomainCertAuthorities/{certAuthorityName} | Get a Standard Domain Certificate Authority object.
*TopicEndpointApi* | [**GetMsgVpnTopicEndpoint**](docs/TopicEndpointApi.md#getmsgvpntopicendpoint) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Get a Topic Endpoint object.
*TopicEndpointApi* | [**GetMsgVpnTopicEndpointMsg**](docs/TopicEndpointApi.md#getmsgvpntopicendpointmsg) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId} | Get a Topic Endpoint Message object.
*TopicEndpointApi* | [**GetMsgVpnTopicEndpointMsgs**](docs/TopicEndpointApi.md#getmsgvpntopicendpointmsgs) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs | Get a list of Topic Endpoint Message objects.
*TopicEndpointApi* | [**GetMsgVpnTopicEndpointPriorities**](docs/TopicEndpointApi.md#getmsgvpntopicendpointpriorities) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/priorities | Get a list of Topic Endpoint Priority objects.
*TopicEndpointApi* | [**GetMsgVpnTopicEndpointPriority**](docs/TopicEndpointApi.md#getmsgvpntopicendpointpriority) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/priorities/{priority} | Get a Topic Endpoint Priority object.
*TopicEndpointApi* | [**GetMsgVpnTopicEndpointTxFlow**](docs/TopicEndpointApi.md#getmsgvpntopicendpointtxflow) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/txFlows/{flowId} | Get a Topic Endpoint Transmit Flow object.
*TopicEndpointApi* | [**GetMsgVpnTopicEndpointTxFlows**](docs/TopicEndpointApi.md#getmsgvpntopicendpointtxflows) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/txFlows | Get a list of Topic Endpoint Transmit Flow objects.
*TopicEndpointApi* | [**GetMsgVpnTopicEndpoints**](docs/TopicEndpointApi.md#getmsgvpntopicendpoints) | **Get** /msgVpns/{msgVpnName}/topicEndpoints | Get a list of Topic Endpoint objects.
*TopicEndpointTemplateApi* | [**GetMsgVpnTopicEndpointTemplate**](docs/TopicEndpointTemplateApi.md#getmsgvpntopicendpointtemplate) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Get a Topic Endpoint Template object.
*TopicEndpointTemplateApi* | [**GetMsgVpnTopicEndpointTemplates**](docs/TopicEndpointTemplateApi.md#getmsgvpntopicendpointtemplates) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates | Get a list of Topic Endpoint Template objects.
*TransactionApi* | [**GetMsgVpnTransaction**](docs/TransactionApi.md#getmsgvpntransaction) | **Get** /msgVpns/{msgVpnName}/transactions/{xid} | Get a Replicated Local Transaction or XA Transaction object.
*TransactionApi* | [**GetMsgVpnTransactionConsumerMsg**](docs/TransactionApi.md#getmsgvpntransactionconsumermsg) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/consumerMsgs/{msgId} | Get a Transaction Consumer Message object.
*TransactionApi* | [**GetMsgVpnTransactionConsumerMsgs**](docs/TransactionApi.md#getmsgvpntransactionconsumermsgs) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/consumerMsgs | Get a list of Transaction Consumer Message objects.
*TransactionApi* | [**GetMsgVpnTransactionPublisherMsg**](docs/TransactionApi.md#getmsgvpntransactionpublishermsg) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/publisherMsgs/{msgId} | Get a Transaction Publisher Message object.
*TransactionApi* | [**GetMsgVpnTransactionPublisherMsgs**](docs/TransactionApi.md#getmsgvpntransactionpublishermsgs) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/publisherMsgs | Get a list of Transaction Publisher Message objects.
*TransactionApi* | [**GetMsgVpnTransactions**](docs/TransactionApi.md#getmsgvpntransactions) | **Get** /msgVpns/{msgVpnName}/transactions | Get a list of Replicated Local Transaction or XA Transaction objects.
*VirtualHostnameApi* | [**GetVirtualHostname**](docs/VirtualHostnameApi.md#getvirtualhostname) | **Get** /virtualHostnames/{virtualHostname} | Get a Virtual Hostname object.
*VirtualHostnameApi* | [**GetVirtualHostnames**](docs/VirtualHostnameApi.md#getvirtualhostnames) | **Get** /virtualHostnames | Get a list of Virtual Hostname objects.

## Documentation For Models

 - [About](docs/About.md)
 - [AboutApi](docs/AboutApi.md)
 - [AboutApiCollections](docs/AboutApiCollections.md)
 - [AboutApiLinks](docs/AboutApiLinks.md)
 - [AboutApiResponse](docs/AboutApiResponse.md)
 - [AboutCollections](docs/AboutCollections.md)
 - [AboutLinks](docs/AboutLinks.md)
 - [AboutResponse](docs/AboutResponse.md)
 - [AboutUser](docs/AboutUser.md)
 - [AboutUserCollections](docs/AboutUserCollections.md)
 - [AboutUserCollectionsMsgvpns](docs/AboutUserCollectionsMsgvpns.md)
 - [AboutUserLinks](docs/AboutUserLinks.md)
 - [AboutUserMsgVpn](docs/AboutUserMsgVpn.md)
 - [AboutUserMsgVpnCollections](docs/AboutUserMsgVpnCollections.md)
 - [AboutUserMsgVpnLinks](docs/AboutUserMsgVpnLinks.md)
 - [AboutUserMsgVpnResponse](docs/AboutUserMsgVpnResponse.md)
 - [AboutUserMsgVpnsResponse](docs/AboutUserMsgVpnsResponse.md)
 - [AboutUserResponse](docs/AboutUserResponse.md)
 - [Broker](docs/Broker.md)
 - [BrokerCollections](docs/BrokerCollections.md)
 - [BrokerCollectionsCertauthorities](docs/BrokerCollectionsCertauthorities.md)
 - [BrokerCollectionsClientcertauthorities](docs/BrokerCollectionsClientcertauthorities.md)
 - [BrokerCollectionsDmrclusters](docs/BrokerCollectionsDmrclusters.md)
 - [BrokerCollectionsDomaincertauthorities](docs/BrokerCollectionsDomaincertauthorities.md)
 - [BrokerCollectionsMsgvpns](docs/BrokerCollectionsMsgvpns.md)
 - [BrokerCollectionsSessions](docs/BrokerCollectionsSessions.md)
 - [BrokerCollectionsStandarddomaincertauthorities](docs/BrokerCollectionsStandarddomaincertauthorities.md)
 - [BrokerCollectionsVirtualhostnames](docs/BrokerCollectionsVirtualhostnames.md)
 - [BrokerLinks](docs/BrokerLinks.md)
 - [BrokerResponse](docs/BrokerResponse.md)
 - [CertAuthoritiesResponse](docs/CertAuthoritiesResponse.md)
 - [CertAuthority](docs/CertAuthority.md)
 - [CertAuthorityCollections](docs/CertAuthorityCollections.md)
 - [CertAuthorityCollectionsOcsptlstrustedcommonnames](docs/CertAuthorityCollectionsOcsptlstrustedcommonnames.md)
 - [CertAuthorityLinks](docs/CertAuthorityLinks.md)
 - [CertAuthorityOcspTlsTrustedCommonName](docs/CertAuthorityOcspTlsTrustedCommonName.md)
 - [CertAuthorityOcspTlsTrustedCommonNameCollections](docs/CertAuthorityOcspTlsTrustedCommonNameCollections.md)
 - [CertAuthorityOcspTlsTrustedCommonNameLinks](docs/CertAuthorityOcspTlsTrustedCommonNameLinks.md)
 - [CertAuthorityOcspTlsTrustedCommonNameResponse](docs/CertAuthorityOcspTlsTrustedCommonNameResponse.md)
 - [CertAuthorityOcspTlsTrustedCommonNamesResponse](docs/CertAuthorityOcspTlsTrustedCommonNamesResponse.md)
 - [CertAuthorityResponse](docs/CertAuthorityResponse.md)
 - [ClientCertAuthoritiesResponse](docs/ClientCertAuthoritiesResponse.md)
 - [ClientCertAuthority](docs/ClientCertAuthority.md)
 - [ClientCertAuthorityCollections](docs/ClientCertAuthorityCollections.md)
 - [ClientCertAuthorityCollectionsOcsptlstrustedcommonnames](docs/ClientCertAuthorityCollectionsOcsptlstrustedcommonnames.md)
 - [ClientCertAuthorityLinks](docs/ClientCertAuthorityLinks.md)
 - [ClientCertAuthorityOcspTlsTrustedCommonName](docs/ClientCertAuthorityOcspTlsTrustedCommonName.md)
 - [ClientCertAuthorityOcspTlsTrustedCommonNameCollections](docs/ClientCertAuthorityOcspTlsTrustedCommonNameCollections.md)
 - [ClientCertAuthorityOcspTlsTrustedCommonNameLinks](docs/ClientCertAuthorityOcspTlsTrustedCommonNameLinks.md)
 - [ClientCertAuthorityOcspTlsTrustedCommonNameResponse](docs/ClientCertAuthorityOcspTlsTrustedCommonNameResponse.md)
 - [ClientCertAuthorityOcspTlsTrustedCommonNamesResponse](docs/ClientCertAuthorityOcspTlsTrustedCommonNamesResponse.md)
 - [ClientCertAuthorityResponse](docs/ClientCertAuthorityResponse.md)
 - [DmrCluster](docs/DmrCluster.md)
 - [DmrClusterCollections](docs/DmrClusterCollections.md)
 - [DmrClusterCollectionsLinks](docs/DmrClusterCollectionsLinks.md)
 - [DmrClusterCollectionsTopologyissues](docs/DmrClusterCollectionsTopologyissues.md)
 - [DmrClusterLink](docs/DmrClusterLink.md)
 - [DmrClusterLinkChannel](docs/DmrClusterLinkChannel.md)
 - [DmrClusterLinkChannelCollections](docs/DmrClusterLinkChannelCollections.md)
 - [DmrClusterLinkChannelLinks](docs/DmrClusterLinkChannelLinks.md)
 - [DmrClusterLinkChannelResponse](docs/DmrClusterLinkChannelResponse.md)
 - [DmrClusterLinkChannelsResponse](docs/DmrClusterLinkChannelsResponse.md)
 - [DmrClusterLinkCollections](docs/DmrClusterLinkCollections.md)
 - [DmrClusterLinkCollectionsChannels](docs/DmrClusterLinkCollectionsChannels.md)
 - [DmrClusterLinkCollectionsRemoteaddresses](docs/DmrClusterLinkCollectionsRemoteaddresses.md)
 - [DmrClusterLinkCollectionsTlstrustedcommonnames](docs/DmrClusterLinkCollectionsTlstrustedcommonnames.md)
 - [DmrClusterLinkLinks](docs/DmrClusterLinkLinks.md)
 - [DmrClusterLinkRemoteAddress](docs/DmrClusterLinkRemoteAddress.md)
 - [DmrClusterLinkRemoteAddressCollections](docs/DmrClusterLinkRemoteAddressCollections.md)
 - [DmrClusterLinkRemoteAddressLinks](docs/DmrClusterLinkRemoteAddressLinks.md)
 - [DmrClusterLinkRemoteAddressResponse](docs/DmrClusterLinkRemoteAddressResponse.md)
 - [DmrClusterLinkRemoteAddressesResponse](docs/DmrClusterLinkRemoteAddressesResponse.md)
 - [DmrClusterLinkResponse](docs/DmrClusterLinkResponse.md)
 - [DmrClusterLinkTlsTrustedCommonName](docs/DmrClusterLinkTlsTrustedCommonName.md)
 - [DmrClusterLinkTlsTrustedCommonNameCollections](docs/DmrClusterLinkTlsTrustedCommonNameCollections.md)
 - [DmrClusterLinkTlsTrustedCommonNameLinks](docs/DmrClusterLinkTlsTrustedCommonNameLinks.md)
 - [DmrClusterLinkTlsTrustedCommonNameResponse](docs/DmrClusterLinkTlsTrustedCommonNameResponse.md)
 - [DmrClusterLinkTlsTrustedCommonNamesResponse](docs/DmrClusterLinkTlsTrustedCommonNamesResponse.md)
 - [DmrClusterLinks](docs/DmrClusterLinks.md)
 - [DmrClusterLinksResponse](docs/DmrClusterLinksResponse.md)
 - [DmrClusterResponse](docs/DmrClusterResponse.md)
 - [DmrClusterTopologyIssue](docs/DmrClusterTopologyIssue.md)
 - [DmrClusterTopologyIssueCollections](docs/DmrClusterTopologyIssueCollections.md)
 - [DmrClusterTopologyIssueLinks](docs/DmrClusterTopologyIssueLinks.md)
 - [DmrClusterTopologyIssueResponse](docs/DmrClusterTopologyIssueResponse.md)
 - [DmrClusterTopologyIssuesResponse](docs/DmrClusterTopologyIssuesResponse.md)
 - [DmrClustersResponse](docs/DmrClustersResponse.md)
 - [DomainCertAuthoritiesResponse](docs/DomainCertAuthoritiesResponse.md)
 - [DomainCertAuthority](docs/DomainCertAuthority.md)
 - [DomainCertAuthorityCollections](docs/DomainCertAuthorityCollections.md)
 - [DomainCertAuthorityLinks](docs/DomainCertAuthorityLinks.md)
 - [DomainCertAuthorityResponse](docs/DomainCertAuthorityResponse.md)
 - [EventThreshold](docs/EventThreshold.md)
 - [EventThresholdByPercent](docs/EventThresholdByPercent.md)
 - [EventThresholdByValue](docs/EventThresholdByValue.md)
 - [MsgVpn](docs/MsgVpn.md)
 - [MsgVpnAclProfile](docs/MsgVpnAclProfile.md)
 - [MsgVpnAclProfileClientConnectException](docs/MsgVpnAclProfileClientConnectException.md)
 - [MsgVpnAclProfileClientConnectExceptionCollections](docs/MsgVpnAclProfileClientConnectExceptionCollections.md)
 - [MsgVpnAclProfileClientConnectExceptionLinks](docs/MsgVpnAclProfileClientConnectExceptionLinks.md)
 - [MsgVpnAclProfileClientConnectExceptionResponse](docs/MsgVpnAclProfileClientConnectExceptionResponse.md)
 - [MsgVpnAclProfileClientConnectExceptionsResponse](docs/MsgVpnAclProfileClientConnectExceptionsResponse.md)
 - [MsgVpnAclProfileCollections](docs/MsgVpnAclProfileCollections.md)
 - [MsgVpnAclProfileCollectionsClientconnectexceptions](docs/MsgVpnAclProfileCollectionsClientconnectexceptions.md)
 - [MsgVpnAclProfileCollectionsPublishexceptions](docs/MsgVpnAclProfileCollectionsPublishexceptions.md)
 - [MsgVpnAclProfileCollectionsPublishtopicexceptions](docs/MsgVpnAclProfileCollectionsPublishtopicexceptions.md)
 - [MsgVpnAclProfileCollectionsSubscribeexceptions](docs/MsgVpnAclProfileCollectionsSubscribeexceptions.md)
 - [MsgVpnAclProfileCollectionsSubscribesharenameexceptions](docs/MsgVpnAclProfileCollectionsSubscribesharenameexceptions.md)
 - [MsgVpnAclProfileCollectionsSubscribetopicexceptions](docs/MsgVpnAclProfileCollectionsSubscribetopicexceptions.md)
 - [MsgVpnAclProfileLinks](docs/MsgVpnAclProfileLinks.md)
 - [MsgVpnAclProfilePublishException](docs/MsgVpnAclProfilePublishException.md)
 - [MsgVpnAclProfilePublishExceptionCollections](docs/MsgVpnAclProfilePublishExceptionCollections.md)
 - [MsgVpnAclProfilePublishExceptionLinks](docs/MsgVpnAclProfilePublishExceptionLinks.md)
 - [MsgVpnAclProfilePublishExceptionResponse](docs/MsgVpnAclProfilePublishExceptionResponse.md)
 - [MsgVpnAclProfilePublishExceptionsResponse](docs/MsgVpnAclProfilePublishExceptionsResponse.md)
 - [MsgVpnAclProfilePublishTopicException](docs/MsgVpnAclProfilePublishTopicException.md)
 - [MsgVpnAclProfilePublishTopicExceptionCollections](docs/MsgVpnAclProfilePublishTopicExceptionCollections.md)
 - [MsgVpnAclProfilePublishTopicExceptionLinks](docs/MsgVpnAclProfilePublishTopicExceptionLinks.md)
 - [MsgVpnAclProfilePublishTopicExceptionResponse](docs/MsgVpnAclProfilePublishTopicExceptionResponse.md)
 - [MsgVpnAclProfilePublishTopicExceptionsResponse](docs/MsgVpnAclProfilePublishTopicExceptionsResponse.md)
 - [MsgVpnAclProfileResponse](docs/MsgVpnAclProfileResponse.md)
 - [MsgVpnAclProfileSubscribeException](docs/MsgVpnAclProfileSubscribeException.md)
 - [MsgVpnAclProfileSubscribeExceptionCollections](docs/MsgVpnAclProfileSubscribeExceptionCollections.md)
 - [MsgVpnAclProfileSubscribeExceptionLinks](docs/MsgVpnAclProfileSubscribeExceptionLinks.md)
 - [MsgVpnAclProfileSubscribeExceptionResponse](docs/MsgVpnAclProfileSubscribeExceptionResponse.md)
 - [MsgVpnAclProfileSubscribeExceptionsResponse](docs/MsgVpnAclProfileSubscribeExceptionsResponse.md)
 - [MsgVpnAclProfileSubscribeShareNameException](docs/MsgVpnAclProfileSubscribeShareNameException.md)
 - [MsgVpnAclProfileSubscribeShareNameExceptionCollections](docs/MsgVpnAclProfileSubscribeShareNameExceptionCollections.md)
 - [MsgVpnAclProfileSubscribeShareNameExceptionLinks](docs/MsgVpnAclProfileSubscribeShareNameExceptionLinks.md)
 - [MsgVpnAclProfileSubscribeShareNameExceptionResponse](docs/MsgVpnAclProfileSubscribeShareNameExceptionResponse.md)
 - [MsgVpnAclProfileSubscribeShareNameExceptionsResponse](docs/MsgVpnAclProfileSubscribeShareNameExceptionsResponse.md)
 - [MsgVpnAclProfileSubscribeTopicException](docs/MsgVpnAclProfileSubscribeTopicException.md)
 - [MsgVpnAclProfileSubscribeTopicExceptionCollections](docs/MsgVpnAclProfileSubscribeTopicExceptionCollections.md)
 - [MsgVpnAclProfileSubscribeTopicExceptionLinks](docs/MsgVpnAclProfileSubscribeTopicExceptionLinks.md)
 - [MsgVpnAclProfileSubscribeTopicExceptionResponse](docs/MsgVpnAclProfileSubscribeTopicExceptionResponse.md)
 - [MsgVpnAclProfileSubscribeTopicExceptionsResponse](docs/MsgVpnAclProfileSubscribeTopicExceptionsResponse.md)
 - [MsgVpnAclProfilesResponse](docs/MsgVpnAclProfilesResponse.md)
 - [MsgVpnAuthenticationOauthProvider](docs/MsgVpnAuthenticationOauthProvider.md)
 - [MsgVpnAuthenticationOauthProviderCollections](docs/MsgVpnAuthenticationOauthProviderCollections.md)
 - [MsgVpnAuthenticationOauthProviderLinks](docs/MsgVpnAuthenticationOauthProviderLinks.md)
 - [MsgVpnAuthenticationOauthProviderResponse](docs/MsgVpnAuthenticationOauthProviderResponse.md)
 - [MsgVpnAuthenticationOauthProvidersResponse](docs/MsgVpnAuthenticationOauthProvidersResponse.md)
 - [MsgVpnAuthorizationGroup](docs/MsgVpnAuthorizationGroup.md)
 - [MsgVpnAuthorizationGroupCollections](docs/MsgVpnAuthorizationGroupCollections.md)
 - [MsgVpnAuthorizationGroupLinks](docs/MsgVpnAuthorizationGroupLinks.md)
 - [MsgVpnAuthorizationGroupResponse](docs/MsgVpnAuthorizationGroupResponse.md)
 - [MsgVpnAuthorizationGroupsResponse](docs/MsgVpnAuthorizationGroupsResponse.md)
 - [MsgVpnBridge](docs/MsgVpnBridge.md)
 - [MsgVpnBridgeCollections](docs/MsgVpnBridgeCollections.md)
 - [MsgVpnBridgeCollectionsLocalsubscriptions](docs/MsgVpnBridgeCollectionsLocalsubscriptions.md)
 - [MsgVpnBridgeCollectionsRemotemsgvpns](docs/MsgVpnBridgeCollectionsRemotemsgvpns.md)
 - [MsgVpnBridgeCollectionsRemotesubscriptions](docs/MsgVpnBridgeCollectionsRemotesubscriptions.md)
 - [MsgVpnBridgeCollectionsTlstrustedcommonnames](docs/MsgVpnBridgeCollectionsTlstrustedcommonnames.md)
 - [MsgVpnBridgeCounter](docs/MsgVpnBridgeCounter.md)
 - [MsgVpnBridgeLinks](docs/MsgVpnBridgeLinks.md)
 - [MsgVpnBridgeLocalSubscription](docs/MsgVpnBridgeLocalSubscription.md)
 - [MsgVpnBridgeLocalSubscriptionCollections](docs/MsgVpnBridgeLocalSubscriptionCollections.md)
 - [MsgVpnBridgeLocalSubscriptionLinks](docs/MsgVpnBridgeLocalSubscriptionLinks.md)
 - [MsgVpnBridgeLocalSubscriptionResponse](docs/MsgVpnBridgeLocalSubscriptionResponse.md)
 - [MsgVpnBridgeLocalSubscriptionsResponse](docs/MsgVpnBridgeLocalSubscriptionsResponse.md)
 - [MsgVpnBridgeRate](docs/MsgVpnBridgeRate.md)
 - [MsgVpnBridgeRemoteMsgVpn](docs/MsgVpnBridgeRemoteMsgVpn.md)
 - [MsgVpnBridgeRemoteMsgVpnCollections](docs/MsgVpnBridgeRemoteMsgVpnCollections.md)
 - [MsgVpnBridgeRemoteMsgVpnLinks](docs/MsgVpnBridgeRemoteMsgVpnLinks.md)
 - [MsgVpnBridgeRemoteMsgVpnResponse](docs/MsgVpnBridgeRemoteMsgVpnResponse.md)
 - [MsgVpnBridgeRemoteMsgVpnsResponse](docs/MsgVpnBridgeRemoteMsgVpnsResponse.md)
 - [MsgVpnBridgeRemoteSubscription](docs/MsgVpnBridgeRemoteSubscription.md)
 - [MsgVpnBridgeRemoteSubscriptionCollections](docs/MsgVpnBridgeRemoteSubscriptionCollections.md)
 - [MsgVpnBridgeRemoteSubscriptionLinks](docs/MsgVpnBridgeRemoteSubscriptionLinks.md)
 - [MsgVpnBridgeRemoteSubscriptionResponse](docs/MsgVpnBridgeRemoteSubscriptionResponse.md)
 - [MsgVpnBridgeRemoteSubscriptionsResponse](docs/MsgVpnBridgeRemoteSubscriptionsResponse.md)
 - [MsgVpnBridgeResponse](docs/MsgVpnBridgeResponse.md)
 - [MsgVpnBridgeTlsTrustedCommonName](docs/MsgVpnBridgeTlsTrustedCommonName.md)
 - [MsgVpnBridgeTlsTrustedCommonNameCollections](docs/MsgVpnBridgeTlsTrustedCommonNameCollections.md)
 - [MsgVpnBridgeTlsTrustedCommonNameLinks](docs/MsgVpnBridgeTlsTrustedCommonNameLinks.md)
 - [MsgVpnBridgeTlsTrustedCommonNameResponse](docs/MsgVpnBridgeTlsTrustedCommonNameResponse.md)
 - [MsgVpnBridgeTlsTrustedCommonNamesResponse](docs/MsgVpnBridgeTlsTrustedCommonNamesResponse.md)
 - [MsgVpnBridgesResponse](docs/MsgVpnBridgesResponse.md)
 - [MsgVpnClient](docs/MsgVpnClient.md)
 - [MsgVpnClientCollections](docs/MsgVpnClientCollections.md)
 - [MsgVpnClientCollectionsConnections](docs/MsgVpnClientCollectionsConnections.md)
 - [MsgVpnClientCollectionsRxflows](docs/MsgVpnClientCollectionsRxflows.md)
 - [MsgVpnClientCollectionsSubscriptions](docs/MsgVpnClientCollectionsSubscriptions.md)
 - [MsgVpnClientCollectionsTransactedsessions](docs/MsgVpnClientCollectionsTransactedsessions.md)
 - [MsgVpnClientCollectionsTxflows](docs/MsgVpnClientCollectionsTxflows.md)
 - [MsgVpnClientConnection](docs/MsgVpnClientConnection.md)
 - [MsgVpnClientConnectionCollections](docs/MsgVpnClientConnectionCollections.md)
 - [MsgVpnClientConnectionLinks](docs/MsgVpnClientConnectionLinks.md)
 - [MsgVpnClientConnectionResponse](docs/MsgVpnClientConnectionResponse.md)
 - [MsgVpnClientConnectionsResponse](docs/MsgVpnClientConnectionsResponse.md)
 - [MsgVpnClientLinks](docs/MsgVpnClientLinks.md)
 - [MsgVpnClientProfile](docs/MsgVpnClientProfile.md)
 - [MsgVpnClientProfileCollections](docs/MsgVpnClientProfileCollections.md)
 - [MsgVpnClientProfileLinks](docs/MsgVpnClientProfileLinks.md)
 - [MsgVpnClientProfileResponse](docs/MsgVpnClientProfileResponse.md)
 - [MsgVpnClientProfilesResponse](docs/MsgVpnClientProfilesResponse.md)
 - [MsgVpnClientResponse](docs/MsgVpnClientResponse.md)
 - [MsgVpnClientRxFlow](docs/MsgVpnClientRxFlow.md)
 - [MsgVpnClientRxFlowCollections](docs/MsgVpnClientRxFlowCollections.md)
 - [MsgVpnClientRxFlowLinks](docs/MsgVpnClientRxFlowLinks.md)
 - [MsgVpnClientRxFlowResponse](docs/MsgVpnClientRxFlowResponse.md)
 - [MsgVpnClientRxFlowsResponse](docs/MsgVpnClientRxFlowsResponse.md)
 - [MsgVpnClientSubscription](docs/MsgVpnClientSubscription.md)
 - [MsgVpnClientSubscriptionCollections](docs/MsgVpnClientSubscriptionCollections.md)
 - [MsgVpnClientSubscriptionLinks](docs/MsgVpnClientSubscriptionLinks.md)
 - [MsgVpnClientSubscriptionResponse](docs/MsgVpnClientSubscriptionResponse.md)
 - [MsgVpnClientSubscriptionsResponse](docs/MsgVpnClientSubscriptionsResponse.md)
 - [MsgVpnClientTransactedSession](docs/MsgVpnClientTransactedSession.md)
 - [MsgVpnClientTransactedSessionCollections](docs/MsgVpnClientTransactedSessionCollections.md)
 - [MsgVpnClientTransactedSessionLinks](docs/MsgVpnClientTransactedSessionLinks.md)
 - [MsgVpnClientTransactedSessionResponse](docs/MsgVpnClientTransactedSessionResponse.md)
 - [MsgVpnClientTransactedSessionsResponse](docs/MsgVpnClientTransactedSessionsResponse.md)
 - [MsgVpnClientTxFlow](docs/MsgVpnClientTxFlow.md)
 - [MsgVpnClientTxFlowCollections](docs/MsgVpnClientTxFlowCollections.md)
 - [MsgVpnClientTxFlowLinks](docs/MsgVpnClientTxFlowLinks.md)
 - [MsgVpnClientTxFlowResponse](docs/MsgVpnClientTxFlowResponse.md)
 - [MsgVpnClientTxFlowsResponse](docs/MsgVpnClientTxFlowsResponse.md)
 - [MsgVpnClientUsername](docs/MsgVpnClientUsername.md)
 - [MsgVpnClientUsernameCollections](docs/MsgVpnClientUsernameCollections.md)
 - [MsgVpnClientUsernameLinks](docs/MsgVpnClientUsernameLinks.md)
 - [MsgVpnClientUsernameResponse](docs/MsgVpnClientUsernameResponse.md)
 - [MsgVpnClientUsernamesResponse](docs/MsgVpnClientUsernamesResponse.md)
 - [MsgVpnClientsResponse](docs/MsgVpnClientsResponse.md)
 - [MsgVpnCollections](docs/MsgVpnCollections.md)
 - [MsgVpnCollectionsAclprofiles](docs/MsgVpnCollectionsAclprofiles.md)
 - [MsgVpnCollectionsAuthenticationoauthproviders](docs/MsgVpnCollectionsAuthenticationoauthproviders.md)
 - [MsgVpnCollectionsAuthorizationgroups](docs/MsgVpnCollectionsAuthorizationgroups.md)
 - [MsgVpnCollectionsBridges](docs/MsgVpnCollectionsBridges.md)
 - [MsgVpnCollectionsClientprofiles](docs/MsgVpnCollectionsClientprofiles.md)
 - [MsgVpnCollectionsClients](docs/MsgVpnCollectionsClients.md)
 - [MsgVpnCollectionsClientusernames](docs/MsgVpnCollectionsClientusernames.md)
 - [MsgVpnCollectionsConfigsyncremotenodes](docs/MsgVpnCollectionsConfigsyncremotenodes.md)
 - [MsgVpnCollectionsDistributedcaches](docs/MsgVpnCollectionsDistributedcaches.md)
 - [MsgVpnCollectionsDmrbridges](docs/MsgVpnCollectionsDmrbridges.md)
 - [MsgVpnCollectionsJndiconnectionfactories](docs/MsgVpnCollectionsJndiconnectionfactories.md)
 - [MsgVpnCollectionsJndiqueues](docs/MsgVpnCollectionsJndiqueues.md)
 - [MsgVpnCollectionsJnditopics](docs/MsgVpnCollectionsJnditopics.md)
 - [MsgVpnCollectionsMqttretaincaches](docs/MsgVpnCollectionsMqttretaincaches.md)
 - [MsgVpnCollectionsMqttsessions](docs/MsgVpnCollectionsMqttsessions.md)
 - [MsgVpnCollectionsQueues](docs/MsgVpnCollectionsQueues.md)
 - [MsgVpnCollectionsQueuetemplates](docs/MsgVpnCollectionsQueuetemplates.md)
 - [MsgVpnCollectionsReplaylogs](docs/MsgVpnCollectionsReplaylogs.md)
 - [MsgVpnCollectionsReplicatedtopics](docs/MsgVpnCollectionsReplicatedtopics.md)
 - [MsgVpnCollectionsRestdeliverypoints](docs/MsgVpnCollectionsRestdeliverypoints.md)
 - [MsgVpnCollectionsTopicendpoints](docs/MsgVpnCollectionsTopicendpoints.md)
 - [MsgVpnCollectionsTopicendpointtemplates](docs/MsgVpnCollectionsTopicendpointtemplates.md)
 - [MsgVpnCollectionsTransactions](docs/MsgVpnCollectionsTransactions.md)
 - [MsgVpnConfigSyncRemoteNode](docs/MsgVpnConfigSyncRemoteNode.md)
 - [MsgVpnConfigSyncRemoteNodeCollections](docs/MsgVpnConfigSyncRemoteNodeCollections.md)
 - [MsgVpnConfigSyncRemoteNodeLinks](docs/MsgVpnConfigSyncRemoteNodeLinks.md)
 - [MsgVpnConfigSyncRemoteNodeResponse](docs/MsgVpnConfigSyncRemoteNodeResponse.md)
 - [MsgVpnConfigSyncRemoteNodesResponse](docs/MsgVpnConfigSyncRemoteNodesResponse.md)
 - [MsgVpnCounter](docs/MsgVpnCounter.md)
 - [MsgVpnDistributedCache](docs/MsgVpnDistributedCache.md)
 - [MsgVpnDistributedCacheCluster](docs/MsgVpnDistributedCacheCluster.md)
 - [MsgVpnDistributedCacheClusterCollections](docs/MsgVpnDistributedCacheClusterCollections.md)
 - [MsgVpnDistributedCacheClusterCollectionsGlobalcachinghomeclusters](docs/MsgVpnDistributedCacheClusterCollectionsGlobalcachinghomeclusters.md)
 - [MsgVpnDistributedCacheClusterCollectionsInstances](docs/MsgVpnDistributedCacheClusterCollectionsInstances.md)
 - [MsgVpnDistributedCacheClusterCollectionsTopics](docs/MsgVpnDistributedCacheClusterCollectionsTopics.md)
 - [MsgVpnDistributedCacheClusterGlobalCachingHomeCluster](docs/MsgVpnDistributedCacheClusterGlobalCachingHomeCluster.md)
 - [MsgVpnDistributedCacheClusterGlobalCachingHomeClusterCollections](docs/MsgVpnDistributedCacheClusterGlobalCachingHomeClusterCollections.md)
 - [MsgVpnDistributedCacheClusterGlobalCachingHomeClusterCollectionsTopicprefixes](docs/MsgVpnDistributedCacheClusterGlobalCachingHomeClusterCollectionsTopicprefixes.md)
 - [MsgVpnDistributedCacheClusterGlobalCachingHomeClusterLinks](docs/MsgVpnDistributedCacheClusterGlobalCachingHomeClusterLinks.md)
 - [MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse](docs/MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse.md)
 - [MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix](docs/MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix.md)
 - [MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixCollections](docs/MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixCollections.md)
 - [MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixLinks](docs/MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixLinks.md)
 - [MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse](docs/MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse.md)
 - [MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse](docs/MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse.md)
 - [MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse](docs/MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse.md)
 - [MsgVpnDistributedCacheClusterInstance](docs/MsgVpnDistributedCacheClusterInstance.md)
 - [MsgVpnDistributedCacheClusterInstanceCollections](docs/MsgVpnDistributedCacheClusterInstanceCollections.md)
 - [MsgVpnDistributedCacheClusterInstanceCollectionsRemoteglobalcachinghomeclusters](docs/MsgVpnDistributedCacheClusterInstanceCollectionsRemoteglobalcachinghomeclusters.md)
 - [MsgVpnDistributedCacheClusterInstanceCollectionsRemotetopics](docs/MsgVpnDistributedCacheClusterInstanceCollectionsRemotetopics.md)
 - [MsgVpnDistributedCacheClusterInstanceCounter](docs/MsgVpnDistributedCacheClusterInstanceCounter.md)
 - [MsgVpnDistributedCacheClusterInstanceLinks](docs/MsgVpnDistributedCacheClusterInstanceLinks.md)
 - [MsgVpnDistributedCacheClusterInstanceRate](docs/MsgVpnDistributedCacheClusterInstanceRate.md)
 - [MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster](docs/MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster.md)
 - [MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterCollections](docs/MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterCollections.md)
 - [MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterLinks](docs/MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterLinks.md)
 - [MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse](docs/MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse.md)
 - [MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse](docs/MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse.md)
 - [MsgVpnDistributedCacheClusterInstanceRemoteTopic](docs/MsgVpnDistributedCacheClusterInstanceRemoteTopic.md)
 - [MsgVpnDistributedCacheClusterInstanceRemoteTopicCollections](docs/MsgVpnDistributedCacheClusterInstanceRemoteTopicCollections.md)
 - [MsgVpnDistributedCacheClusterInstanceRemoteTopicLinks](docs/MsgVpnDistributedCacheClusterInstanceRemoteTopicLinks.md)
 - [MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse](docs/MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse.md)
 - [MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse](docs/MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse.md)
 - [MsgVpnDistributedCacheClusterInstanceResponse](docs/MsgVpnDistributedCacheClusterInstanceResponse.md)
 - [MsgVpnDistributedCacheClusterInstancesResponse](docs/MsgVpnDistributedCacheClusterInstancesResponse.md)
 - [MsgVpnDistributedCacheClusterLinks](docs/MsgVpnDistributedCacheClusterLinks.md)
 - [MsgVpnDistributedCacheClusterResponse](docs/MsgVpnDistributedCacheClusterResponse.md)
 - [MsgVpnDistributedCacheClusterTopic](docs/MsgVpnDistributedCacheClusterTopic.md)
 - [MsgVpnDistributedCacheClusterTopicCollections](docs/MsgVpnDistributedCacheClusterTopicCollections.md)
 - [MsgVpnDistributedCacheClusterTopicLinks](docs/MsgVpnDistributedCacheClusterTopicLinks.md)
 - [MsgVpnDistributedCacheClusterTopicResponse](docs/MsgVpnDistributedCacheClusterTopicResponse.md)
 - [MsgVpnDistributedCacheClusterTopicsResponse](docs/MsgVpnDistributedCacheClusterTopicsResponse.md)
 - [MsgVpnDistributedCacheClustersResponse](docs/MsgVpnDistributedCacheClustersResponse.md)
 - [MsgVpnDistributedCacheCollections](docs/MsgVpnDistributedCacheCollections.md)
 - [MsgVpnDistributedCacheCollectionsClusters](docs/MsgVpnDistributedCacheCollectionsClusters.md)
 - [MsgVpnDistributedCacheLinks](docs/MsgVpnDistributedCacheLinks.md)
 - [MsgVpnDistributedCacheResponse](docs/MsgVpnDistributedCacheResponse.md)
 - [MsgVpnDistributedCachesResponse](docs/MsgVpnDistributedCachesResponse.md)
 - [MsgVpnDmrBridge](docs/MsgVpnDmrBridge.md)
 - [MsgVpnDmrBridgeCollections](docs/MsgVpnDmrBridgeCollections.md)
 - [MsgVpnDmrBridgeLinks](docs/MsgVpnDmrBridgeLinks.md)
 - [MsgVpnDmrBridgeResponse](docs/MsgVpnDmrBridgeResponse.md)
 - [MsgVpnDmrBridgesResponse](docs/MsgVpnDmrBridgesResponse.md)
 - [MsgVpnJndiConnectionFactoriesResponse](docs/MsgVpnJndiConnectionFactoriesResponse.md)
 - [MsgVpnJndiConnectionFactory](docs/MsgVpnJndiConnectionFactory.md)
 - [MsgVpnJndiConnectionFactoryCollections](docs/MsgVpnJndiConnectionFactoryCollections.md)
 - [MsgVpnJndiConnectionFactoryLinks](docs/MsgVpnJndiConnectionFactoryLinks.md)
 - [MsgVpnJndiConnectionFactoryResponse](docs/MsgVpnJndiConnectionFactoryResponse.md)
 - [MsgVpnJndiQueue](docs/MsgVpnJndiQueue.md)
 - [MsgVpnJndiQueueCollections](docs/MsgVpnJndiQueueCollections.md)
 - [MsgVpnJndiQueueLinks](docs/MsgVpnJndiQueueLinks.md)
 - [MsgVpnJndiQueueResponse](docs/MsgVpnJndiQueueResponse.md)
 - [MsgVpnJndiQueuesResponse](docs/MsgVpnJndiQueuesResponse.md)
 - [MsgVpnJndiTopic](docs/MsgVpnJndiTopic.md)
 - [MsgVpnJndiTopicCollections](docs/MsgVpnJndiTopicCollections.md)
 - [MsgVpnJndiTopicLinks](docs/MsgVpnJndiTopicLinks.md)
 - [MsgVpnJndiTopicResponse](docs/MsgVpnJndiTopicResponse.md)
 - [MsgVpnJndiTopicsResponse](docs/MsgVpnJndiTopicsResponse.md)
 - [MsgVpnLinks](docs/MsgVpnLinks.md)
 - [MsgVpnMqttRetainCache](docs/MsgVpnMqttRetainCache.md)
 - [MsgVpnMqttRetainCacheCollections](docs/MsgVpnMqttRetainCacheCollections.md)
 - [MsgVpnMqttRetainCacheLinks](docs/MsgVpnMqttRetainCacheLinks.md)
 - [MsgVpnMqttRetainCacheResponse](docs/MsgVpnMqttRetainCacheResponse.md)
 - [MsgVpnMqttRetainCachesResponse](docs/MsgVpnMqttRetainCachesResponse.md)
 - [MsgVpnMqttSession](docs/MsgVpnMqttSession.md)
 - [MsgVpnMqttSessionCollections](docs/MsgVpnMqttSessionCollections.md)
 - [MsgVpnMqttSessionCollectionsSubscriptions](docs/MsgVpnMqttSessionCollectionsSubscriptions.md)
 - [MsgVpnMqttSessionCounter](docs/MsgVpnMqttSessionCounter.md)
 - [MsgVpnMqttSessionLinks](docs/MsgVpnMqttSessionLinks.md)
 - [MsgVpnMqttSessionResponse](docs/MsgVpnMqttSessionResponse.md)
 - [MsgVpnMqttSessionSubscription](docs/MsgVpnMqttSessionSubscription.md)
 - [MsgVpnMqttSessionSubscriptionCollections](docs/MsgVpnMqttSessionSubscriptionCollections.md)
 - [MsgVpnMqttSessionSubscriptionLinks](docs/MsgVpnMqttSessionSubscriptionLinks.md)
 - [MsgVpnMqttSessionSubscriptionResponse](docs/MsgVpnMqttSessionSubscriptionResponse.md)
 - [MsgVpnMqttSessionSubscriptionsResponse](docs/MsgVpnMqttSessionSubscriptionsResponse.md)
 - [MsgVpnMqttSessionsResponse](docs/MsgVpnMqttSessionsResponse.md)
 - [MsgVpnQueue](docs/MsgVpnQueue.md)
 - [MsgVpnQueueCollections](docs/MsgVpnQueueCollections.md)
 - [MsgVpnQueueCollectionsMsgs](docs/MsgVpnQueueCollectionsMsgs.md)
 - [MsgVpnQueueCollectionsPriorities](docs/MsgVpnQueueCollectionsPriorities.md)
 - [MsgVpnQueueCollectionsSubscriptions](docs/MsgVpnQueueCollectionsSubscriptions.md)
 - [MsgVpnQueueCollectionsTxflows](docs/MsgVpnQueueCollectionsTxflows.md)
 - [MsgVpnQueueLinks](docs/MsgVpnQueueLinks.md)
 - [MsgVpnQueueMsg](docs/MsgVpnQueueMsg.md)
 - [MsgVpnQueueMsgCollections](docs/MsgVpnQueueMsgCollections.md)
 - [MsgVpnQueueMsgLinks](docs/MsgVpnQueueMsgLinks.md)
 - [MsgVpnQueueMsgResponse](docs/MsgVpnQueueMsgResponse.md)
 - [MsgVpnQueueMsgsResponse](docs/MsgVpnQueueMsgsResponse.md)
 - [MsgVpnQueuePrioritiesResponse](docs/MsgVpnQueuePrioritiesResponse.md)
 - [MsgVpnQueuePriority](docs/MsgVpnQueuePriority.md)
 - [MsgVpnQueuePriorityCollections](docs/MsgVpnQueuePriorityCollections.md)
 - [MsgVpnQueuePriorityLinks](docs/MsgVpnQueuePriorityLinks.md)
 - [MsgVpnQueuePriorityResponse](docs/MsgVpnQueuePriorityResponse.md)
 - [MsgVpnQueueResponse](docs/MsgVpnQueueResponse.md)
 - [MsgVpnQueueSubscription](docs/MsgVpnQueueSubscription.md)
 - [MsgVpnQueueSubscriptionCollections](docs/MsgVpnQueueSubscriptionCollections.md)
 - [MsgVpnQueueSubscriptionLinks](docs/MsgVpnQueueSubscriptionLinks.md)
 - [MsgVpnQueueSubscriptionResponse](docs/MsgVpnQueueSubscriptionResponse.md)
 - [MsgVpnQueueSubscriptionsResponse](docs/MsgVpnQueueSubscriptionsResponse.md)
 - [MsgVpnQueueTemplate](docs/MsgVpnQueueTemplate.md)
 - [MsgVpnQueueTemplateCollections](docs/MsgVpnQueueTemplateCollections.md)
 - [MsgVpnQueueTemplateLinks](docs/MsgVpnQueueTemplateLinks.md)
 - [MsgVpnQueueTemplateResponse](docs/MsgVpnQueueTemplateResponse.md)
 - [MsgVpnQueueTemplatesResponse](docs/MsgVpnQueueTemplatesResponse.md)
 - [MsgVpnQueueTxFlow](docs/MsgVpnQueueTxFlow.md)
 - [MsgVpnQueueTxFlowCollections](docs/MsgVpnQueueTxFlowCollections.md)
 - [MsgVpnQueueTxFlowLinks](docs/MsgVpnQueueTxFlowLinks.md)
 - [MsgVpnQueueTxFlowResponse](docs/MsgVpnQueueTxFlowResponse.md)
 - [MsgVpnQueueTxFlowsResponse](docs/MsgVpnQueueTxFlowsResponse.md)
 - [MsgVpnQueuesResponse](docs/MsgVpnQueuesResponse.md)
 - [MsgVpnRate](docs/MsgVpnRate.md)
 - [MsgVpnReplayLog](docs/MsgVpnReplayLog.md)
 - [MsgVpnReplayLogCollections](docs/MsgVpnReplayLogCollections.md)
 - [MsgVpnReplayLogCollectionsMsgs](docs/MsgVpnReplayLogCollectionsMsgs.md)
 - [MsgVpnReplayLogLinks](docs/MsgVpnReplayLogLinks.md)
 - [MsgVpnReplayLogMsg](docs/MsgVpnReplayLogMsg.md)
 - [MsgVpnReplayLogMsgCollections](docs/MsgVpnReplayLogMsgCollections.md)
 - [MsgVpnReplayLogMsgLinks](docs/MsgVpnReplayLogMsgLinks.md)
 - [MsgVpnReplayLogMsgResponse](docs/MsgVpnReplayLogMsgResponse.md)
 - [MsgVpnReplayLogMsgsResponse](docs/MsgVpnReplayLogMsgsResponse.md)
 - [MsgVpnReplayLogResponse](docs/MsgVpnReplayLogResponse.md)
 - [MsgVpnReplayLogsResponse](docs/MsgVpnReplayLogsResponse.md)
 - [MsgVpnReplicatedTopic](docs/MsgVpnReplicatedTopic.md)
 - [MsgVpnReplicatedTopicCollections](docs/MsgVpnReplicatedTopicCollections.md)
 - [MsgVpnReplicatedTopicLinks](docs/MsgVpnReplicatedTopicLinks.md)
 - [MsgVpnReplicatedTopicResponse](docs/MsgVpnReplicatedTopicResponse.md)
 - [MsgVpnReplicatedTopicsResponse](docs/MsgVpnReplicatedTopicsResponse.md)
 - [MsgVpnResponse](docs/MsgVpnResponse.md)
 - [MsgVpnRestDeliveryPoint](docs/MsgVpnRestDeliveryPoint.md)
 - [MsgVpnRestDeliveryPointCollections](docs/MsgVpnRestDeliveryPointCollections.md)
 - [MsgVpnRestDeliveryPointCollectionsQueuebindings](docs/MsgVpnRestDeliveryPointCollectionsQueuebindings.md)
 - [MsgVpnRestDeliveryPointCollectionsRestconsumers](docs/MsgVpnRestDeliveryPointCollectionsRestconsumers.md)
 - [MsgVpnRestDeliveryPointLinks](docs/MsgVpnRestDeliveryPointLinks.md)
 - [MsgVpnRestDeliveryPointQueueBinding](docs/MsgVpnRestDeliveryPointQueueBinding.md)
 - [MsgVpnRestDeliveryPointQueueBindingCollections](docs/MsgVpnRestDeliveryPointQueueBindingCollections.md)
 - [MsgVpnRestDeliveryPointQueueBindingLinks](docs/MsgVpnRestDeliveryPointQueueBindingLinks.md)
 - [MsgVpnRestDeliveryPointQueueBindingResponse](docs/MsgVpnRestDeliveryPointQueueBindingResponse.md)
 - [MsgVpnRestDeliveryPointQueueBindingsResponse](docs/MsgVpnRestDeliveryPointQueueBindingsResponse.md)
 - [MsgVpnRestDeliveryPointResponse](docs/MsgVpnRestDeliveryPointResponse.md)
 - [MsgVpnRestDeliveryPointRestConsumer](docs/MsgVpnRestDeliveryPointRestConsumer.md)
 - [MsgVpnRestDeliveryPointRestConsumerCollections](docs/MsgVpnRestDeliveryPointRestConsumerCollections.md)
 - [MsgVpnRestDeliveryPointRestConsumerCollectionsOauthjwtclaims](docs/MsgVpnRestDeliveryPointRestConsumerCollectionsOauthjwtclaims.md)
 - [MsgVpnRestDeliveryPointRestConsumerCollectionsTlstrustedcommonnames](docs/MsgVpnRestDeliveryPointRestConsumerCollectionsTlstrustedcommonnames.md)
 - [MsgVpnRestDeliveryPointRestConsumerCounter](docs/MsgVpnRestDeliveryPointRestConsumerCounter.md)
 - [MsgVpnRestDeliveryPointRestConsumerLinks](docs/MsgVpnRestDeliveryPointRestConsumerLinks.md)
 - [MsgVpnRestDeliveryPointRestConsumerOauthJwtClaim](docs/MsgVpnRestDeliveryPointRestConsumerOauthJwtClaim.md)
 - [MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimCollections](docs/MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimCollections.md)
 - [MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimLinks](docs/MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimLinks.md)
 - [MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse](docs/MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse.md)
 - [MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse](docs/MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse.md)
 - [MsgVpnRestDeliveryPointRestConsumerResponse](docs/MsgVpnRestDeliveryPointRestConsumerResponse.md)
 - [MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName](docs/MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName.md)
 - [MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameCollections](docs/MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameCollections.md)
 - [MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameLinks](docs/MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameLinks.md)
 - [MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse](docs/MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse.md)
 - [MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse](docs/MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse.md)
 - [MsgVpnRestDeliveryPointRestConsumersResponse](docs/MsgVpnRestDeliveryPointRestConsumersResponse.md)
 - [MsgVpnRestDeliveryPointsResponse](docs/MsgVpnRestDeliveryPointsResponse.md)
 - [MsgVpnTopicEndpoint](docs/MsgVpnTopicEndpoint.md)
 - [MsgVpnTopicEndpointCollections](docs/MsgVpnTopicEndpointCollections.md)
 - [MsgVpnTopicEndpointCollectionsMsgs](docs/MsgVpnTopicEndpointCollectionsMsgs.md)
 - [MsgVpnTopicEndpointCollectionsPriorities](docs/MsgVpnTopicEndpointCollectionsPriorities.md)
 - [MsgVpnTopicEndpointCollectionsTxflows](docs/MsgVpnTopicEndpointCollectionsTxflows.md)
 - [MsgVpnTopicEndpointLinks](docs/MsgVpnTopicEndpointLinks.md)
 - [MsgVpnTopicEndpointMsg](docs/MsgVpnTopicEndpointMsg.md)
 - [MsgVpnTopicEndpointMsgCollections](docs/MsgVpnTopicEndpointMsgCollections.md)
 - [MsgVpnTopicEndpointMsgLinks](docs/MsgVpnTopicEndpointMsgLinks.md)
 - [MsgVpnTopicEndpointMsgResponse](docs/MsgVpnTopicEndpointMsgResponse.md)
 - [MsgVpnTopicEndpointMsgsResponse](docs/MsgVpnTopicEndpointMsgsResponse.md)
 - [MsgVpnTopicEndpointPrioritiesResponse](docs/MsgVpnTopicEndpointPrioritiesResponse.md)
 - [MsgVpnTopicEndpointPriority](docs/MsgVpnTopicEndpointPriority.md)
 - [MsgVpnTopicEndpointPriorityCollections](docs/MsgVpnTopicEndpointPriorityCollections.md)
 - [MsgVpnTopicEndpointPriorityLinks](docs/MsgVpnTopicEndpointPriorityLinks.md)
 - [MsgVpnTopicEndpointPriorityResponse](docs/MsgVpnTopicEndpointPriorityResponse.md)
 - [MsgVpnTopicEndpointResponse](docs/MsgVpnTopicEndpointResponse.md)
 - [MsgVpnTopicEndpointTemplate](docs/MsgVpnTopicEndpointTemplate.md)
 - [MsgVpnTopicEndpointTemplateCollections](docs/MsgVpnTopicEndpointTemplateCollections.md)
 - [MsgVpnTopicEndpointTemplateLinks](docs/MsgVpnTopicEndpointTemplateLinks.md)
 - [MsgVpnTopicEndpointTemplateResponse](docs/MsgVpnTopicEndpointTemplateResponse.md)
 - [MsgVpnTopicEndpointTemplatesResponse](docs/MsgVpnTopicEndpointTemplatesResponse.md)
 - [MsgVpnTopicEndpointTxFlow](docs/MsgVpnTopicEndpointTxFlow.md)
 - [MsgVpnTopicEndpointTxFlowCollections](docs/MsgVpnTopicEndpointTxFlowCollections.md)
 - [MsgVpnTopicEndpointTxFlowLinks](docs/MsgVpnTopicEndpointTxFlowLinks.md)
 - [MsgVpnTopicEndpointTxFlowResponse](docs/MsgVpnTopicEndpointTxFlowResponse.md)
 - [MsgVpnTopicEndpointTxFlowsResponse](docs/MsgVpnTopicEndpointTxFlowsResponse.md)
 - [MsgVpnTopicEndpointsResponse](docs/MsgVpnTopicEndpointsResponse.md)
 - [MsgVpnTransaction](docs/MsgVpnTransaction.md)
 - [MsgVpnTransactionCollections](docs/MsgVpnTransactionCollections.md)
 - [MsgVpnTransactionCollectionsConsumermsgs](docs/MsgVpnTransactionCollectionsConsumermsgs.md)
 - [MsgVpnTransactionCollectionsPublishermsgs](docs/MsgVpnTransactionCollectionsPublishermsgs.md)
 - [MsgVpnTransactionConsumerMsg](docs/MsgVpnTransactionConsumerMsg.md)
 - [MsgVpnTransactionConsumerMsgCollections](docs/MsgVpnTransactionConsumerMsgCollections.md)
 - [MsgVpnTransactionConsumerMsgLinks](docs/MsgVpnTransactionConsumerMsgLinks.md)
 - [MsgVpnTransactionConsumerMsgResponse](docs/MsgVpnTransactionConsumerMsgResponse.md)
 - [MsgVpnTransactionConsumerMsgsResponse](docs/MsgVpnTransactionConsumerMsgsResponse.md)
 - [MsgVpnTransactionLinks](docs/MsgVpnTransactionLinks.md)
 - [MsgVpnTransactionPublisherMsg](docs/MsgVpnTransactionPublisherMsg.md)
 - [MsgVpnTransactionPublisherMsgCollections](docs/MsgVpnTransactionPublisherMsgCollections.md)
 - [MsgVpnTransactionPublisherMsgLinks](docs/MsgVpnTransactionPublisherMsgLinks.md)
 - [MsgVpnTransactionPublisherMsgResponse](docs/MsgVpnTransactionPublisherMsgResponse.md)
 - [MsgVpnTransactionPublisherMsgsResponse](docs/MsgVpnTransactionPublisherMsgsResponse.md)
 - [MsgVpnTransactionResponse](docs/MsgVpnTransactionResponse.md)
 - [MsgVpnTransactionsResponse](docs/MsgVpnTransactionsResponse.md)
 - [MsgVpnsResponse](docs/MsgVpnsResponse.md)
 - [SempError](docs/SempError.md)
 - [SempMeta](docs/SempMeta.md)
 - [SempMetaOnlyResponse](docs/SempMetaOnlyResponse.md)
 - [SempPaging](docs/SempPaging.md)
 - [SempRequest](docs/SempRequest.md)
 - [Session](docs/Session.md)
 - [SessionCollections](docs/SessionCollections.md)
 - [SessionLinks](docs/SessionLinks.md)
 - [SessionResponse](docs/SessionResponse.md)
 - [SessionsResponse](docs/SessionsResponse.md)
 - [StandardDomainCertAuthoritiesResponse](docs/StandardDomainCertAuthoritiesResponse.md)
 - [StandardDomainCertAuthority](docs/StandardDomainCertAuthority.md)
 - [StandardDomainCertAuthorityCollections](docs/StandardDomainCertAuthorityCollections.md)
 - [StandardDomainCertAuthorityLinks](docs/StandardDomainCertAuthorityLinks.md)
 - [StandardDomainCertAuthorityResponse](docs/StandardDomainCertAuthorityResponse.md)
 - [VirtualHostname](docs/VirtualHostname.md)
 - [VirtualHostnameCollections](docs/VirtualHostnameCollections.md)
 - [VirtualHostnameLinks](docs/VirtualHostnameLinks.md)
 - [VirtualHostnameResponse](docs/VirtualHostnameResponse.md)
 - [VirtualHostnamesResponse](docs/VirtualHostnamesResponse.md)

## Documentation For Authorization

## basicAuth
- **Type**: HTTP basic authentication

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

## Author

support@solace.com
