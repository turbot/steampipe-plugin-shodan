package shodan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableShodanScan(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:             "shodan_scan",
		Description:      "List the scans requested for this account.",
		DefaultTransform: transform.FromJSONTag(),
		List: &plugin.ListConfig{
			Hydrate: listScan,
		},
		// NOTE: GetScan returns less data fields than ListScans, and I believe
		// this API is normally going to have a small number of results. So, the
		// get has not been implemented at this time.
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the scan."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status of the scan."},
			// Other columns
			// Weirdly, the shodan API returns the API key. We don't to keep it more secret.
			// {Name: "api_key", Type: proto.ColumnType_STRING, Description: "API key for the scan."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the scan was created."},
			{Name: "credits_left", Type: proto.ColumnType_INT, Description: "Credits left after this scan."},
			{Name: "size", Type: proto.ColumnType_INT, Description: "Size of the scan."},
			{Name: "status_check", Type: proto.ColumnType_TIMESTAMP, Description: "Status check for the scan."},
		},
	}
}

func listScan(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shodan_scan.listScan", "connection_error", err)
		return nil, err
	}
	page := uint(1)
	count := 0
	for {
		result, err := conn.ListScans(ctx, page)
		if err != nil {
			plugin.Logger(ctx).Error("shodan_scan.listScan", "query_error", err)
			return nil, err
		}
		for _, i := range result.Matches {
			d.StreamListItem(ctx, i)
			count++
		}
		if count >= result.Total {
			break
		}
		page++
	}
	return nil, nil
}
