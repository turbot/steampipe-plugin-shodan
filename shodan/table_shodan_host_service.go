package shodan

import (
	"context"
	"regexp"

	"github.com/shadowscatcher/shodan/models"
	"github.com/shadowscatcher/shodan/search"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableShodanHostService(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:             "shodan_host_service",
		Description:      "All services that have been found on the given host.",
		DefaultTransform: transform.FromJSONTag(),
		List: &plugin.ListConfig{
			Hydrate:    listHostService,
			KeyColumns: plugin.SingleColumn("ip"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "ip", Type: proto.ColumnType_IPADDR, Hydrate: ipString, Transform: transform.FromValue(), Description: "The IP address of the host as a string."},
			{Name: "ipv6", Type: proto.ColumnType_STRING, Description: "The IPv6 address of the host as a string."},
			{Name: "port", Type: proto.ColumnType_INT, Description: "Port number that the service is operating on."},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp for when the banner was fetched from the device."},
			{Name: "hash", Type: proto.ColumnType_INT, Description: "Numeric hash of the data property."},
			{Name: "domains", Type: proto.ColumnType_JSON, Description: "An array of strings containing the top-level domains for the hostnames of the device."},
			{Name: "link", Type: proto.ColumnType_STRING, Description: "The network link type. Possible values are: \"Ethernet or modem\", \"generic tunnel or VPN\", \"DSL\", \"IPIP or SIT\", \"SLIP\", \"IPSec or GRE\", \"VLAN\", \"jumbo Ethernet\", \"Google\", \"GIF\", \"PPTP\", \"loopback\", \"AX.25 radio modem\"."},
			{Name: "uptime", Type: proto.ColumnType_INT, Description: "Uptime of the IP (in minutes)."},
			{Name: "transport", Type: proto.ColumnType_STRING, Description: "Uptime of the IP (in minutes)."},
			{Name: "product", Type: proto.ColumnType_STRING, Transform: transform.FromMethod("ProductString"), Description: "Name of the software running the service."},
			{Name: "version", Type: proto.ColumnType_STRING, Transform: transform.FromMethod("VersionString"), Description: "Version of the software running the service."},
			{Name: "cpe", Type: proto.ColumnType_JSON, Transform: transform.FromField("CPE"), Description: "Common Platform Enumeration."},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "Title of the website as extracted from the HTML source."},
			{Name: "devicetype", Type: proto.ColumnType_STRING, Description: "The type of device (webcam, router, etc.)."},
			{Name: "info", Type: proto.ColumnType_STRING, Description: "Miscellaneous information that was extracted about the product."},
			{Name: "shodan", Type: proto.ColumnType_JSON, Transform: transform.FromField("Shodan"), Description: "Information about how the banner was generated. It doesn’t store any data about the port/service itself."},
			{Name: "vulns", Type: proto.ColumnType_JSON, Description: "The vulns property contains information about vulnerabilities that may exist in the service represented by the banner. In general, the Shodan crawlers don’t perform vulnerability testing as a result the vulnerabilities stored in vulns are inferred from the banner and haven’t been verified. Availability: Banners where the software/version has been identified and there exist known CVEs for it."},
			{Name: "banners", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(banner), Description: "TODO"},

			{Name: "ssl", Type: proto.ColumnType_JSON, Transform: transform.FromField("SSL"), Description: "Services that require SSL (ex. HTTPS) or support upgrading a connection to SSL/TLS (ex. POP3 with STARTTLS)."},
			{Name: "cassandra", Type: proto.ColumnType_JSON, Transform: transform.FromField("Cassandra"), Description: "Cassandra database services that allow connections to the client Thrift port (default: 9160/ tcp)."},
			{Name: "db2", Type: proto.ColumnType_JSON, Transform: transform.FromField("DB2"), Description: "Services running the IBM DB2 DRDA protocol."},
			{Name: "dns", Type: proto.ColumnType_JSON, Transform: transform.FromField("DNS"), Description: "DNS servers that support either UDP or TCP (typically on port 53)."},
			{Name: "docker", Type: proto.ColumnType_JSON, Transform: transform.FromField("Docker"), Description: "Docker services that allow remote connections and don’t have authentication enabled."},
			{Name: "elastic", Type: proto.ColumnType_JSON, Transform: transform.FromField("Elastic"), Description: "The elastic property is available in banners that are identified as belonging to an Elastic service."},
			{Name: "etcd", Type: proto.ColumnType_JSON, Transform: transform.FromField("Etcd"), Description: "The etcd service provides a distributed key/value store used by projects such as Kubernetes. Ports that are running the etcd service."},
			{Name: "ethernet_ip", Type: proto.ColumnType_JSON, Transform: transform.FromField("EthernetIP"), Description: "Devices that complete a handshake in either TCP or UDP for the industrial Ethernet/IP protocol."},
			{Name: "ftp", Type: proto.ColumnType_JSON, Transform: transform.FromField("FTP"), Description: "FTP services running on the default port 21/TCP. If the FTP service supports STARTTLS then the starttls tag will be added to the list of tags on the banner and it will also have a top-level ssl property which contains the certificate, SSL testing results and more."},
			{Name: "hive", Type: proto.ColumnType_JSON, Transform: transform.FromField("Hive"), Description: "Devices running Apache Hive servers on any port that Shodan crawls."},
			{Name: "http", Type: proto.ColumnType_JSON, Transform: transform.FromField("HTTP"), Description: "The banner was generated by a HTTP module (http, https, http-simple-new, https-simple-new) and successfully completed a HTTP handshake."},
			{Name: "isakmp", Type: proto.ColumnType_JSON, Transform: transform.FromField("ISAKMP"), Description: "VPN services that use the ISAKMP protocol (such as IKE)."},
			{Name: "lantronix", Type: proto.ColumnType_JSON, Transform: transform.FromField("Lantronix"), Description: "Lantronix devices that are running the configuration service."},
			{Name: "monero", Type: proto.ColumnType_JSON, Transform: transform.FromField("Monero"), Description: "If the Monero RPC service is enabled and accepting remote connections. Most results are on port 18081, but it can also be available on other ports."},
			{Name: "mongodb", Type: proto.ColumnType_JSON, Transform: transform.FromField("MongoDB"), Description: "MongoDB services that support the binary protocol to interact with the database."},
			{Name: "mqtt", Type: proto.ColumnType_JSON, Transform: transform.FromField("MQTT"), Description: "MQTT services that allow remote connections."},
			{Name: "netbios", Type: proto.ColumnType_JSON, Transform: transform.FromField("NetBIOS"), Description: "Services that run on port 137 and complete a NetBIOS handshake."},
			{Name: "ntp", Type: proto.ColumnType_JSON, Transform: transform.FromField("NTP"), Description: "NTP daemons supporting at least version 1 or version 2."},
			{Name: "redis", Type: proto.ColumnType_JSON, Transform: transform.FromField("Redis"), Description: "Redis services running on the default port 6379/TCP."},
			{Name: "rip", Type: proto.ColumnType_JSON, Transform: transform.FromField("RIP"), Description: "Services on port 520 that successfully respond to a RIP request."},
			{Name: "rsync", Type: proto.ColumnType_JSON, Transform: transform.FromField("Rsync"), Description: "rsync service information."},
			{Name: "smb", Type: proto.ColumnType_JSON, Transform: transform.FromField("SMB"), Description: "Services that run on port 445 and support either SMBv1 or SMBv2."},
			{Name: "snmp", Type: proto.ColumnType_JSON, Transform: transform.FromField("SNMP"), Description: "Any banner generated by the snmp module (typically on 161/UDP)."},
			{Name: "ssh", Type: proto.ColumnType_JSON, Transform: transform.FromField("SSH"), Description: "Any service banner where the initial response starts with “SSH” and subsequently completes a SSH handshake."},
			{Name: "vertx", Type: proto.ColumnType_JSON, Transform: transform.FromField("VertX"), Description: "Devices running the VertX/Edge door controllers."},
			{Name: "minecraft", Type: proto.ColumnType_JSON, Transform: transform.FromField("Minecraft"), Description: "Devices running the Minecraft game server."},
			{Name: "influxdb", Type: proto.ColumnType_JSON, Transform: transform.FromField("InfluxDB"), Description: "Devices running InfluxDB time-series database."},
			{Name: "coap", Type: proto.ColumnType_JSON, Transform: transform.FromField("CoAP"), Description: "Devices running CoAP IoT protocol service."},

			// Other columns
			{Name: "opts", Type: proto.ColumnType_JSON, Description: "Contains experimental and supplemental data for the service. This can include the SSL certificate, robots.txt and other raw information that hasn't yet been formalized into the Banner Specification."},
		},
	}
}

func banner(_ context.Context, d *transform.TransformData) (interface{}, error) {
	re := regexp.MustCompile(`\\r?\\n`)
	s := d.Value.(*models.Service)
	banners := []string{}
	for _, b := range re.Split(s.Data, -1) {
		if b != "" {
			banners = append(banners, b)
		}
	}
	return banners, nil
}

func listHostService(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shodan_ip.listHostService", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	ip := quals["ip"].GetInetValue().GetAddr()
	result, err := conn.Host(ctx, search.HostParams{IP: ip})
	if err != nil {
		plugin.Logger(ctx).Error("shodan_ip.listHostService", "query_error", err)
		if isErrorWithMessage(err, []string{"No information available for that IP", "Invalid IP"}) {
			// Not found or invalid (so not found)
			return nil, nil
		}
		return nil, err
	}
	for _, i := range result.Services {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
