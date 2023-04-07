package shodan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableShodanDNSReverse(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shodan_dns_reverse",
		Description: "Hostnames defined for the given IP.",
		List: &plugin.ListConfig{
			Hydrate:    listDNSReverse,
			KeyColumns: plugin.SingleColumn("ip"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "ip", Type: proto.ColumnType_IPADDR, Description: "Requested IP address."},
			{Name: "host", Type: proto.ColumnType_STRING, Description: "Hostname associated with the IP."},
		},
	}
}

type reverseDNS struct {
	IP   string
	Host string
}

func listDNSReverse(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shodan_dns_reverse.listDNSReverse", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	ip := quals["ip"].GetInetValue().GetAddr()
	result, err := conn.DnsReverse(ctx, []string{ip})
	if err != nil {
		plugin.Logger(ctx).Error("shodan_dns_reverse.listDNSReverse", "query_error", err)
		return nil, err
	}
	if result[ip] == nil {
		// Not found
		return nil, nil
	}
	for resultIP, hosts := range result {
		for _, h := range hosts {
			d.StreamListItem(ctx, reverseDNS{IP: resultIP, Host: h})
		}
	}
	return nil, nil
}
