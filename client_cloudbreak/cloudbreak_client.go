// Code generated by go-swagger; DO NOT EDIT.

package client_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/client_cloudbreak/v1accountpreferences"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1blueprints"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1clustertemplates"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1connectors"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1constraints"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1credentials"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1events"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1flexsubscriptions"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1imagecatalogs"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1ldap"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1networks"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1rdsconfigs"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1recipes"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1repositoryconfigs"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1securitygroups"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1securityrules"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1settings"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1smartsensesubscriptions"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1stacks"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1subscriptions"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1templates"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1topologies"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1usages"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1users"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1util"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v2connectors"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v2stacks"
)

// Default cloudbreak HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new cloudbreak HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Cloudbreak {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new cloudbreak HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Cloudbreak {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new cloudbreak client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Cloudbreak {
	cli := new(Cloudbreak)
	cli.Transport = transport

	cli.V1accountpreferences = v1accountpreferences.New(transport, formats)

	cli.V1blueprints = v1blueprints.New(transport, formats)

	cli.V1clustertemplates = v1clustertemplates.New(transport, formats)

	cli.V1connectors = v1connectors.New(transport, formats)

	cli.V1constraints = v1constraints.New(transport, formats)

	cli.V1credentials = v1credentials.New(transport, formats)

	cli.V1events = v1events.New(transport, formats)

	cli.V1flexsubscriptions = v1flexsubscriptions.New(transport, formats)

	cli.V1imagecatalogs = v1imagecatalogs.New(transport, formats)

	cli.V1ldap = v1ldap.New(transport, formats)

	cli.V1networks = v1networks.New(transport, formats)

	cli.V1rdsconfigs = v1rdsconfigs.New(transport, formats)

	cli.V1recipes = v1recipes.New(transport, formats)

	cli.V1repositoryconfigs = v1repositoryconfigs.New(transport, formats)

	cli.V1securitygroups = v1securitygroups.New(transport, formats)

	cli.V1securityrules = v1securityrules.New(transport, formats)

	cli.V1settings = v1settings.New(transport, formats)

	cli.V1smartsensesubscriptions = v1smartsensesubscriptions.New(transport, formats)

	cli.V1stacks = v1stacks.New(transport, formats)

	cli.V1subscriptions = v1subscriptions.New(transport, formats)

	cli.V1templates = v1templates.New(transport, formats)

	cli.V1topologies = v1topologies.New(transport, formats)

	cli.V1usages = v1usages.New(transport, formats)

	cli.V1users = v1users.New(transport, formats)

	cli.V1util = v1util.New(transport, formats)

	cli.V2connectors = v2connectors.New(transport, formats)

	cli.V2stacks = v2stacks.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Cloudbreak is a client for cloudbreak
type Cloudbreak struct {
	V1accountpreferences *v1accountpreferences.Client

	V1blueprints *v1blueprints.Client

	V1clustertemplates *v1clustertemplates.Client

	V1connectors *v1connectors.Client

	V1constraints *v1constraints.Client

	V1credentials *v1credentials.Client

	V1events *v1events.Client

	V1flexsubscriptions *v1flexsubscriptions.Client

	V1imagecatalogs *v1imagecatalogs.Client

	V1ldap *v1ldap.Client

	V1networks *v1networks.Client

	V1rdsconfigs *v1rdsconfigs.Client

	V1recipes *v1recipes.Client

	V1repositoryconfigs *v1repositoryconfigs.Client

	V1securitygroups *v1securitygroups.Client

	V1securityrules *v1securityrules.Client

	V1settings *v1settings.Client

	V1smartsensesubscriptions *v1smartsensesubscriptions.Client

	V1stacks *v1stacks.Client

	V1subscriptions *v1subscriptions.Client

	V1templates *v1templates.Client

	V1topologies *v1topologies.Client

	V1usages *v1usages.Client

	V1users *v1users.Client

	V1util *v1util.Client

	V2connectors *v2connectors.Client

	V2stacks *v2stacks.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Cloudbreak) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.V1accountpreferences.SetTransport(transport)

	c.V1blueprints.SetTransport(transport)

	c.V1clustertemplates.SetTransport(transport)

	c.V1connectors.SetTransport(transport)

	c.V1constraints.SetTransport(transport)

	c.V1credentials.SetTransport(transport)

	c.V1events.SetTransport(transport)

	c.V1flexsubscriptions.SetTransport(transport)

	c.V1imagecatalogs.SetTransport(transport)

	c.V1ldap.SetTransport(transport)

	c.V1networks.SetTransport(transport)

	c.V1rdsconfigs.SetTransport(transport)

	c.V1recipes.SetTransport(transport)

	c.V1repositoryconfigs.SetTransport(transport)

	c.V1securitygroups.SetTransport(transport)

	c.V1securityrules.SetTransport(transport)

	c.V1settings.SetTransport(transport)

	c.V1smartsensesubscriptions.SetTransport(transport)

	c.V1stacks.SetTransport(transport)

	c.V1subscriptions.SetTransport(transport)

	c.V1templates.SetTransport(transport)

	c.V1topologies.SetTransport(transport)

	c.V1usages.SetTransport(transport)

	c.V1users.SetTransport(transport)

	c.V1util.SetTransport(transport)

	c.V2connectors.SetTransport(transport)

	c.V2stacks.SetTransport(transport)

}
