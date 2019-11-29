/*
 * Copyright 2019 VMware, Inc.  All rights reserved.  Licensed under the Apache v2 License.
 */

package types

import "encoding/xml"

// FirewallConfigWithXml allows to enable/disable firewall on a specific edge gateway
// Reference: vCloud Director API for NSX Programming Guide
// https://code.vmware.com/docs/6900/vcloud-director-api-for-nsx-programming-guide
//
// Warning. It nests all firewall rules because Edge Gateway API is done so that if this data is not
// sent while enabling it would wipe all firewall rules. InnerXML type field is used with struct tag
//`innerxml` to prevent any manipulation of configuration and sending it verbatim
type FirewallConfigWithXml struct {
	XMLName       xml.Name              `xml:"firewall"`
	Enabled       bool                  `xml:"enabled"`
	DefaultPolicy FirewallDefaultPolicy `xml:"defaultPolicy"`

	// Each configuration change has a version number
	Version string `xml:"version,omitempty"`

	// The below field has `innerxml` tag so that it is not processed but instead
	// sent verbatim
	FirewallRules InnerXML `xml:"firewallRules,omitempty"`
	GlobalConfig  InnerXML `xml:"globalConfig,omitempty"`
}

// FirewallDefaultPolicy represent default rule
type FirewallDefaultPolicy struct {
	LoggingEnabled bool   `xml:"loggingEnabled"`
	Action         string `xml:"action"`
}

// LbGeneralParamsWithXml allows to enable/disable load balancing capabilities on specific edge gateway
// Reference: vCloud Director API for NSX Programming Guide
// https://code.vmware.com/docs/6900/vcloud-director-api-for-nsx-programming-guide
//
// Warning. It nests all components (LbMonitor, LbPool, LbAppProfile, LbAppRule, LbVirtualServer)
// because Edge Gateway API is done so that if this data is not sent while enabling it would wipe
// all load balancer configurations. InnerXML type fields are used with struct tag `innerxml` to
// prevent any manipulation of configuration and sending it verbatim
type LbGeneralParamsWithXml struct {
	XMLName             xml.Name   `xml:"loadBalancer"`
	Enabled             bool       `xml:"enabled"`
	AccelerationEnabled bool       `xml:"accelerationEnabled"`
	Logging             *LbLogging `xml:"logging"`

	// This field is not used anywhere but needs to be passed through
	EnableServiceInsertion bool `xml:"enableServiceInsertion"`
	// Each configuration change has a version number
	Version string `xml:"version,omitempty"`

	// The below fields have `innerxml` tag so that they are not processed but instead
	// sent verbatim
	VirtualServers []InnerXML `xml:"virtualServer,omitempty"`
	Pools          []InnerXML `xml:"pool,omitempty"`
	AppProfiles    []InnerXML `xml:"applicationProfile,omitempty"`
	Monitors       []InnerXML `xml:"monitor,omitempty"`
	AppRules       []InnerXML `xml:"applicationRule,omitempty"`
}

// LbLogging represents logging configuration for load balancer
type LbLogging struct {
	Enable   bool   `xml:"enable"`
	LogLevel string `xml:"logLevel"`
}

// InnerXML is meant to be used when unmarshaling a field into text rather than struct
// It helps to avoid missing out any fields which may not have been specified in the struct.
type InnerXML struct {
	Text string `xml:",innerxml"`
}

// LbMonitor defines health check parameters for a particular type of network traffic
// Reference: vCloud Director API for NSX Programming Guide
// https://code.vmware.com/docs/6900/vcloud-director-api-for-nsx-programming-guide
type LbMonitor struct {
	XMLName    xml.Name `xml:"monitor"`
	ID         string   `xml:"monitorId,omitempty"`
	Type       string   `xml:"type"`
	Interval   int      `xml:"interval,omitempty"`
	Timeout    int      `xml:"timeout,omitempty"`
	MaxRetries int      `xml:"maxRetries,omitempty"`
	Method     string   `xml:"method,omitempty"`
	URL        string   `xml:"url,omitempty"`
	Expected   string   `xml:"expected,omitempty"`
	Name       string   `xml:"name,omitempty"`
	Send       string   `xml:"send,omitempty"`
	Receive    string   `xml:"receive,omitempty"`
	Extension  string   `xml:"extension,omitempty"`
}

type LbMonitors []LbMonitor

// LbPool represents a load balancer server pool as per "vCloud Director API for NSX Programming Guide"
// Type: LBPoolHealthCheckType
// https://code.vmware.com/docs/6900/vcloud-director-api-for-nsx-programming-guide
type LbPool struct {
	XMLName             xml.Name      `xml:"pool"`
	ID                  string        `xml:"poolId,omitempty"`
	Name                string        `xml:"name"`
	Description         string        `xml:"description,omitempty"`
	Algorithm           string        `xml:"algorithm"`
	AlgorithmParameters string        `xml:"algorithmParameters,omitempty"`
	Transparent         bool          `xml:"transparent"`
	MonitorId           string        `xml:"monitorId,omitempty"`
	Members             LbPoolMembers `xml:"member,omitempty"`
}

type LbPools []LbPool

// LbPoolMember represents a single member inside LbPool
type LbPoolMember struct {
	ID          string `xml:"memberId,omitempty"`
	Name        string `xml:"name"`
	IpAddress   string `xml:"ipAddress"`
	Weight      int    `xml:"weight,omitempty"`
	MonitorPort int    `xml:"monitorPort,omitempty"`
	Port        int    `xml:"port"`
	MaxConn     int    `xml:"maxConn,omitempty"`
	MinConn     int    `xml:"minConn,omitempty"`
	Condition   string `xml:"condition,omitempty"`
}

type LbPoolMembers []LbPoolMember

// LbAppProfile represents a load balancer application profile as per "vCloud Director API for NSX
// Programming Guide"
// https://code.vmware.com/docs/6900/vcloud-director-api-for-nsx-programming-guide
type LbAppProfile struct {
	XMLName                       xml.Name                  `xml:"applicationProfile"`
	ID                            string                    `xml:"applicationProfileId,omitempty"`
	Name                          string                    `xml:"name,omitempty"`
	SslPassthrough                bool                      `xml:"sslPassthrough"`
	Template                      string                    `xml:"template,omitempty"`
	HttpRedirect                  *LbAppProfileHttpRedirect `xml:"httpRedirect,omitempty"`
	Persistence                   *LbAppProfilePersistence  `xml:"persistence,omitempty"`
	InsertXForwardedForHttpHeader bool                      `xml:"insertXForwardedFor"`
	ServerSslEnabled              bool                      `xml:"serverSslEnabled"`
}

type LbAppProfiles []LbAppProfile

// LbAppProfilePersistence defines persistence profile settings in LbAppProfile
type LbAppProfilePersistence struct {
	XMLName    xml.Name `xml:"persistence"`
	Method     string   `xml:"method,omitempty"`
	CookieName string   `xml:"cookieName,omitempty"`
	CookieMode string   `xml:"cookieMode,omitempty"`
	Expire     int      `xml:"expire,omitempty"`
}

// LbAppProfileHttpRedirect defines http redirect settings in LbAppProfile
type LbAppProfileHttpRedirect struct {
	XMLName xml.Name `xml:"httpRedirect"`
	To      string   `xml:"to,omitempty"`
}

// LbAppRule represents a load balancer application rule as per "vCloud Director API for NSX
// Programming Guide"
// https://code.vmware.com/docs/6900/vcloud-director-api-for-nsx-programming-guide
type LbAppRule struct {
	XMLName xml.Name `xml:"applicationRule"`
	ID      string   `xml:"applicationRuleId,omitempty"`
	Name    string   `xml:"name,omitempty"`
	Script  string   `xml:"script,omitempty"`
}

type LbAppRules []LbAppRule

// LbVirtualServer represents a load balancer virtual server as per "vCloud Director API for NSX
// Programming Guide"
// https://code.vmware.com/docs/6900/vcloud-director-api-for-nsx-programming-guide
type LbVirtualServer struct {
	XMLName              xml.Name `xml:"virtualServer"`
	ID                   string   `xml:"virtualServerId,omitempty"`
	Name                 string   `xml:"name,omitempty"`
	Description          string   `xml:"description,omitempty"`
	Enabled              bool     `xml:"enabled"`
	IpAddress            string   `xml:"ipAddress"`
	Protocol             string   `xml:"protocol"`
	Port                 int      `xml:"port"`
	AccelerationEnabled  bool     `xml:"accelerationEnabled"`
	ConnectionLimit      int      `xml:"connectionLimit,omitempty"`
	ConnectionRateLimit  int      `xml:"connectionRateLimit,omitempty"`
	ApplicationProfileId string   `xml:"applicationProfileId,omitempty"`
	DefaultPoolId        string   `xml:"defaultPoolId,omitempty"`
	ApplicationRuleIds   []string `xml:"applicationRuleId,omitempty"`
}

// EdgeNatRule contains shared structure for SNAT and DNAT rule configuration using
// NSX-V proxied edge gateway endpoint
// https://code.vmware.com/docs/6900/vcloud-director-api-for-nsx-programming-guide
type EdgeNatRule struct {
	XMLName           xml.Name `xml:"natRule"`
	ID                string   `xml:"ruleId,omitempty"`
	RuleType          string   `xml:"ruleType,omitempty"`
	RuleTag           string   `xml:"ruleTag,omitempty"`
	Action            string   `xml:"action"`
	Vnic              *int     `xml:"vnic,omitempty"`
	OriginalAddress   string   `xml:"originalAddress"`
	TranslatedAddress string   `xml:"translatedAddress"`
	LoggingEnabled    bool     `xml:"loggingEnabled"`
	Enabled           bool     `xml:"enabled"`
	Description       string   `xml:"description,omitempty"`
	Protocol          string   `xml:"protocol,omitempty"`
	OriginalPort      string   `xml:"originalPort,omitempty"`
	TranslatedPort    string   `xml:"translatedPort,omitempty"`
	IcmpType          string   `xml:"icmpType,omitempty"`
}

// EdgeFirewall holds data for creating firewall rule using proxied NSX-V API
// https://code.vmware.com/docs/6900/vcloud-director-api-for-nsx-programming-guide
type EdgeFirewallRule struct {
	XMLName         xml.Name                `xml:"firewallRule" `
	ID              string                  `xml:"id,omitempty"`
	Name            string                  `xml:"name,omitempty"`
	RuleType        string                  `xml:"ruleType,omitempty"`
	RuleTag         string                  `xml:"ruleTag,omitempty"`
	Source          EdgeFirewallEndpoint    `xml:"source" `
	Destination     EdgeFirewallEndpoint    `xml:"destination"`
	Application     EdgeFirewallApplication `xml:"application"`
	MatchTranslated *bool                   `xml:"matchTranslated,omitempty"`
	Direction       string                  `xml:"direction,omitempty"`
	Action          string                  `xml:"action,omitempty"`
	Enabled         bool                    `xml:"enabled"`
	LoggingEnabled  bool                    `xml:"loggingEnabled"`
}

// EdgeFirewallEndpoint can contains slices of objects for source or destination in EdgeFirewall
type EdgeFirewallEndpoint struct {
	Exclude           bool     `xml:"exclude"`
	VnicGroupIds      []string `xml:"vnicGroupId,omitempty"`
	GroupingObjectIds []string `xml:"groupingObjectId,omitempty"`
	IpAddresses       []string `xml:"ipAddress,omitempty"`
}

// EdgeFirewallApplication Wraps []EdgeFirewallApplicationService for multiple protocol/port specification
type EdgeFirewallApplication struct {
	ID       string                           `xml:"applicationId,omitempty"`
	Services []EdgeFirewallApplicationService `xml:"service,omitempty"`
}

// EdgeFirewallApplicationService defines port/protocol details for one service in EdgeFirewallRule
type EdgeFirewallApplicationService struct {
	Protocol   string `xml:"protocol,omitempty"`
	Port       string `xml:"port,omitempty"`
	SourcePort string `xml:"sourcePort,omitempty"`
}

// EdgeIpSet defines a group of IP addresses that you can add as the source or destination in a
// firewall rule or in DHCP relay configuration. The object itself has more fields in API response,
// however vCD UI only uses the below mentioned. It looks as if the other fields are used in NSX
// internally and are simply proxied back.
//
// Note. Only advanced edge gateways support IP sets
type EdgeIpSet struct {
	XMLName xml.Name `xml:"ipset"`
	// ID holds composite ID of IP set which is formatted as
	// 'f9daf2da-b4f9-4921-a2f4-d77a943a381c:ipset-4' where the first segment before colon is vDC id
	// and the second one is IP set ID
	ID string `xml:"objectId,omitempty"`
	// Name is mandatory and must be unique
	Name string `xml:"name"`
	// Description - optional
	Description string `xml:"description,omitempty"`
	// IPAddresses is a mandatory field with comma separated values. The API is known to re-order
	// data after submiting and may shuffle components even if re-submitted as it was return from
	// API itself
	// (eg: "192.168.200.1,192.168.200.1/24,192.168.200.1-192.168.200.24")
	IPAddresses string `xml:"value"`
	// InheritanceAllowed defines visibility at underlying scopes
	InheritanceAllowed *bool `xml:"inheritanceAllowed"`
	// Revision is a "version" of IP set configuration. During read current revision is being
	// returned and when update is performed this latest version must be sent as it validates if no
	// updates ocurred in between. When not the latest version is being sent during update one can
	// expect similar error response from API: "The object ipset-27 used in this operation has an
	// older version 0 than the current system version 1. Refresh UI or fetch the latest copy of the
	// object and retry operation."
	Revision *int `xml:"revision,omitempty"`
}

// EdgeIpSets is a slice of pointers to EdgeIpSet
type EdgeIpSets []*EdgeIpSet
