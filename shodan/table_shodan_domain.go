package shodan

import (
	"context"

	"github.com/shadowscatcher/shodan/search"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableShodanDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shodan_domain",
		Description: "Get all the subdomains and other DNS entries for the given domain.",
		List: &plugin.ListConfig{
			Hydrate:    listDomain,
			KeyColumns: plugin.SingleColumn("domain"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "domain", Type: proto.ColumnType_STRING, Hydrate: domain, Transform: transform.FromValue(), Description: "Requested domain."},
			// Other columns
			// This API returns a set of tags, indicating some capabilities of the domain.
			// We could return the same result for every row but that's wasteful and not
			// accurate. So, for now, skip the tags field.
			//{Name: "tags", Type: proto.ColumnType_JSON, Description: "TODO"},
			{Name: "subdomain", Type: proto.ColumnType_STRING, Description: "The subdomain, which is prepended to the domain name."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the subdomain (e.g. MX, CNAME, etc)."},
			{Name: "value", Type: proto.ColumnType_STRING, Description: "Value set for this subdomain."},
			{Name: "last_seen", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the subdomain was last seen."},
		},
	}
}

// Better to use the input IP string instead of net.IP.String() to be more confident it
// will match when returning.
func domain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	i := quals["domain"].GetStringValue()
	return i, nil
}

func listDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shodan_domain.listDomain", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	domain := quals["domain"].GetStringValue()
	q := search.DomainQuery{
		Domain: domain,
		Page:   1,
	}
	for {
		result, err := conn.DnsDomain(ctx, q)
		if err != nil {
			plugin.Logger(ctx).Error("shodan_domain.listDomain", "query_error", err)
			if isErrorWithMessage(err, []string{"No information available for that domain", "requires enterprise access"}) {
				// Cases:
				// Not found
				// Paging is only allowed for enterprise. In other cases just return
				// the single page of results and stop.
				break
			}
			return nil, err
		}
		for _, i := range result.Data {
			d.StreamListItem(ctx, i)
		}
		if !result.More {
			break
		}
		q.Page++
	}
	return nil, nil
}
