package shodan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-shodan",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"shodan_account_profile": tableShodanAccountProfile(ctx),
			"shodan_api_info":        tableShodanAPIInfo(ctx),
			"shodan_dns_reverse":     tableShodanDNSReverse(ctx),
			"shodan_domain":          tableShodanDomain(ctx),
			"shodan_exploit":         tableShodanExploit(ctx),
			"shodan_host":            tableShodanHost(ctx),
			"shodan_host_service":    tableShodanHostService(ctx),
			"shodan_port":            tableShodanPort(ctx),
			"shodan_protocol":        tableShodanProtocol(ctx),
			"shodan_scan":            tableShodanScan(ctx),
			"shodan_search":          tableShodanSearch(ctx),
			"shodan_service":         tableShodanService(ctx),
		},
	}
	return p
}
