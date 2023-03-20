package shodan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableShodanProtocol(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shodan_protocol",
		Description: "List of the protocols that can be used when launching an Internet scan.",
		List: &plugin.ListConfig{
			Hydrate: listProtocol,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the protocol (e.g. postgresql)."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the protocol."},
		},
	}
}

type protocol struct {
	Name        string
	Description string
}

func listProtocol(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shodan_protocol.listProtocol", "connection_error", err)
		return nil, err
	}
	result, err := conn.Protocols(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("shodan_protocol.listProtocol", "query_error", err)
		return nil, err
	}
	for k, v := range result {
		d.StreamListItem(ctx, protocol{Name: k, Description: v})
	}
	return nil, nil
}
