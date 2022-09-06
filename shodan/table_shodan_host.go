package shodan

import (
	"context"

	"github.com/shadowscatcher/shodan/search"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableShodanHost(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shodan_host",
		Description: "All services that have been found on the given host at IP.",
		List: &plugin.ListConfig{
			Hydrate:    listHost,
			KeyColumns: plugin.SingleColumn("ip"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "ip", Type: proto.ColumnType_IPADDR, Hydrate: ipString, Transform: transform.FromValue(), Description: "The IP address of the host as a string."},
			// Other columns
			{Name: "asn", Type: proto.ColumnType_STRING, Transform: transform.FromField("ASN"), Description: "The autonomous system number (ex. AS4837)."},
			//{Name: "host_location", Type: proto.ColumnType_JSON, Description: "TODO."},
			{Name: "hostnames", Type: proto.ColumnType_JSON, Description: "An array of strings containing all of the hostnames that have been assigned to the IP address for this device."},
			{Name: "isp", Type: proto.ColumnType_STRING, Transform: transform.FromField("ISP"), Description: "The ISP that is providing the organization with the IP space for this device. Consider this the \"parent\" of the organization in terms of IP ownership."},
			{Name: "last_update", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the IP was last updated."},
			{Name: "org", Type: proto.ColumnType_STRING, Description: "The name of the organization that is assigned the IP space for this device."},
			{Name: "os", Type: proto.ColumnType_STRING, Transform: transform.FromField("OS"), Description: "The operating system that powers the device."},
			{Name: "ports", Type: proto.ColumnType_JSON, Description: "Open ports for the IP."},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "List of tags that describe the characteristics of the device."},
			{Name: "vulns", Type: proto.ColumnType_JSON, Description: "A list of vulnerabilities for the IP."},
			// Location columns
			{Name: "area_code", Type: proto.ColumnType_INT, Description: "Area code for the device's location. Only available for the US."},
			{Name: "city", Type: proto.ColumnType_STRING, Description: "Name of the city where the device is located."},
			{Name: "country_code", Type: proto.ColumnType_STRING, Description: "2-letter country code where the device is located."},
			{Name: "country_code_3", Type: proto.ColumnType_STRING, Description: "3-letter country code where the device is located."},
			{Name: "country_name", Type: proto.ColumnType_STRING, Description: "Name of the country where the device is located."},
			{Name: "dma_code", Type: proto.ColumnType_INT, Description: "The designated market area code for the area where the device is located. Only available for the US."},
			{Name: "latitude", Type: proto.ColumnType_DOUBLE, Description: "Latitude for the geolocation of the device."},
			{Name: "longitude", Type: proto.ColumnType_DOUBLE, Description: "Longitude for the geolocation of the device."},
			{Name: "postal_code", Type: proto.ColumnType_STRING, Description: "The postal code for the device's location."},
			{Name: "region_code", Type: proto.ColumnType_STRING, Description: "Name of the region where the device is located."},
			// Large data of limited value so excluded for now
			//{Name: "services", Type: proto.ColumnType_JSON, Description: "Services found at the IP."},
			//{Name: "html", Type: proto.ColumnType_STRING, Description: "Raw HTML of the response."},
		},
	}
}

// Better to use the input IP string instead of net.IP.String() to be more confident it
// will match when returning.
func ipString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	ip := quals["ip"].GetInetValue().GetAddr()
	return ip, nil
}

func listHost(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shodan_host.listHost", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	ip := quals["ip"].GetInetValue().GetAddr()
	result, err := conn.Host(ctx, search.HostParams{IP: ip})
	if err != nil {
		plugin.Logger(ctx).Error("shodan_host.listHost", "query_error", err)
		if isErrorWithMessage(err, []string{"No information available for that IP", "Invalid IP"}) {
			// Not found or invalid (so not found)
			return nil, nil
		}
		return nil, err
	}
	d.StreamListItem(ctx, result)
	return nil, nil
}
