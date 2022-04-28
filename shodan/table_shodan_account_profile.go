package shodan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableShodanAccountProfile(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shodan_account_profile",
		Description: "Information about the Shodan account linked to the caller.",
		List: &plugin.ListConfig{
			Hydrate: listAccountProfile,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Account name."},
			{Name: "member", Type: proto.ColumnType_BOOL, Description: "True if the account is a member of Shodan."},
			{Name: "credits", Type: proto.ColumnType_INT, Description: "Export credits for this account."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "When the account was created."},
		},
	}
}

func listAccountProfile(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shodan_account_profile.listAccountProfile", "connection_error", err)
		return nil, err
	}
	result, err := conn.AccountProfile(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("shodan_account_profile.listAccountProfile", "query_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, result)
	return nil, nil
}
