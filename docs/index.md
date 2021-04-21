---
organization: Turbot
category: ["security"]
icon_url: "/images/plugins/turbot/shodan.svg"
brand_color: "#C83237"
display_name: Shodan
name: shodan
description: Steampipe plugin to query host, DNS and exploit information using Shodan.
social_about: Query host, DNS and exploit information using SQL. Open source CLI. No DB required. 
social_preview: "/images/plugins/turbot/shodan-social-graphic.png"
---

# Shodan + Steampipe

[Shodan](https://shodan.io) is a search engine for Internet-connected devices. Shodan gathers information about devices directly connected to the Internet. The types of devices that are indexed can vary tremendously: ranging from internet connected cameras to cloud hosted servers, and everything in-between.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For Example:

```sql
select
  i.title,
  h.ip,
  h.ports,
  h.isp
from
  aws_ec2_instance i,
  shodan_host h
where
  i.public_ip_address is not null
  and i.public_ip_address = h.ip;
```
```
+---------------------+-------------+---------+------------------+
| title               | ip          | ports   | isp              |
+---------------------+-------------+---------+------------------+
| Ubuntu 18 Test      | 3.94.22.81  | [80,22] | Amazon.com, Inc. |
| Redhat 8 Test       | 3.94.22.148 | [22]    | Amazon.com, Inc. |
| Ubuntu 20 Test      | 3.29.46.31  | [22]    | Amazon.com, Inc. |
| Amazon Linux 2 Test | 3.39.34.4   | [111]   | Amazon.com, Inc. |
+---------------------+-------------+---------+------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/shodan/tables)**

## Get started

### Install

Download and install the latest Shodan plugin:

```bash
steampipe plugin install shodan
```

### Credentials

| Item | Description |
| - | - |
| Credentials | Shodan requires an API token for all requests, but offers a free tier. Sign up on the [shodan website](https://shodan.com) to get your free token. A valid token looks like `ZGloRBAl2Tvur3tBTu84NkZIf3i5Cc5U`. |
| Permissions | N/A |
| Radius | Each connection represents a single Shodan account. |
| Resolution |  1. Credentials specified in environment variables e.g. `SHODAN_API_KEY`.<br />2. Credentials in the credential file (`~/.shodan/credentials`) |

### Configuration

Installing the latest aws plugin will create a config file (`~/.steampipe/config/shodan.spc`) with a single connection named `shodan`:

```hcl
connection "shodan" {
  plugin  = "shodan"
  api_key = "ZGloRBAl2Tvur3tBTu84NkZIf3i5Cc5U"
}
```

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-shodan
* Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)