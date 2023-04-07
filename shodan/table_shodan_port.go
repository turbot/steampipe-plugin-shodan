package shodan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableShodanPort(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shodan_port",
		Description: "Ports returns a list of port numbers that the crawlers are looking for.",
		List: &plugin.ListConfig{
			Hydrate: listPort,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "port", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Port number."},
		},
	}
}

func listPort(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shodan_port.listPort", "connection_error", err)
		return nil, err
	}
	result, err := conn.Ports(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("shodan_port.listPort", "query_error", err)
		return nil, err
	}
	for _, i := range result {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
