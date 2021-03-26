---
organization: Turbot
category: ["security"]
icon_url: "/images/plugins/turbot/shodan.svg"
brand_color: "#C83237"
display_name: Shodan
name: shodan
description: Steampipe plugin to query host, DNS and exploit information using Shodan.
---

# Steampipe

The Steampipe CLI is open source software that allows you to perform real-time queries against cloud APIs using SQL; all without having to extract, transform and load data into a local DB. If you are just getting started checkout [steampipe.io](https://steampipe.io).

# Shodan

[Shodan](https://shodan.io) is a search engine for Internet-connected devices. Shodan gathers information about devices directly connected to the Internet. The types of devices that are indexed can vary tremendously: ranging from internet connected cameras to cloud hosted servers, and everything in-between.

# Steampipe + Shodan

Steampipe allows you to query Shodan's APIs with SQL; retrieving host metadata, open ports, DNS info and even potential exploits. This metadata can be made even more powerful by joining against other cloud service APIs:

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

Browse all [available tables and their schemas](shodan/tables).

## Installation

[Install or update Steampipe](https://steampipe.io/downloads) if not already installed.

To download and install the latest shodan plugin:

```bash
steampipe plugin install shodan
```

## Credentials

Shodan requires an API token for all requests, but offers a free tier. Sign up on the [shodan website](https://shodan.com) to get your free token. A valid token looks like `ZGloRBAl2Tvur3tBTu84NkZIf3i5Cc5U`.


## Connection Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

Installing the latest shodan plugin will create a default connection named `shodan` in the `~/.steampipe/config/shodan.spc` file.  You must edit this connection to include your API token:

```hcl
connection "shodan" {
  plugin  = "shodan"
  api_key = "ZGloRBAl2Tvur3tBTu84NkZIf3i5Cc5U"
}
```
