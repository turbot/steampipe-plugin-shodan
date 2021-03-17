package shodan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableShodanAPIInfo(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shodan_api_info",
		Description: "Information about the API plan belonging to the given API key.",
		List: &plugin.ListConfig{
			Hydrate: listAPIInfo,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "plan", Type: proto.ColumnType_STRING, Description: "Plan for the API key."},
			{Name: "query_credits", Type: proto.ColumnType_INT, Description: "Query credits."},
			{Name: "scan_credits", Type: proto.ColumnType_INT, Description: "Scan credits."},
			{Name: "https", Type: proto.ColumnType_BOOL, Description: "HTTPS support."},
			{Name: "monitored_ips", Type: proto.ColumnType_INT, Description: "Monitored IPs."},
			{Name: "unlocked", Type: proto.ColumnType_BOOL, Description: "Unlocked."},
			{Name: "unlocked_left", Type: proto.ColumnType_INT, Description: "Unlocked left."},
			{Name: "usage_limits", Type: proto.ColumnType_JSON, Description: "Usage limits."},
			{Name: "telnet", Type: proto.ColumnType_BOOL, Description: "Telnet support."},
		},
	}
}

func listAPIInfo(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shodan_api_info.listAPIInfo", "connection_error", err)
		return nil, err
	}
	result, err := conn.ApiInfo(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("shodan_api_info.listAPIInfo", "query_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, result)
	return nil, nil
}
