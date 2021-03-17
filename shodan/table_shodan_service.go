package shodan

import (
	"context"
	"strconv"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableShodanService(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shodan_service",
		Description: "List of the services Shodan can detect.",
		List: &plugin.ListConfig{
			Hydrate: listService,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "port", Type: proto.ColumnType_INT, Description: "Port of the service (e.g. 5432)."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the service (e.g. postgresql)."},
		},
	}
}

type service struct {
	Port int
	Name string
}

func listService(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shodan_service.listService", "connection_error", err)
		return nil, err
	}
	result, err := conn.Services(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("shodan_service.listService", "query_error", err)
		return nil, err
	}
	for k, v := range result {
		i, err := strconv.Atoi(k)
		if err != nil {
			// Should never happen, log and continue
			plugin.Logger(ctx).Error("shodan_service.listService", "query_error", err)
			continue
		}
		d.StreamListItem(ctx, service{Port: i, Name: v})
	}
	return nil, nil
}
