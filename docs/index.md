---
organization: Turbot
category: ["security"]
icon_url: "/images/plugins/turbot/shodan.svg"
brand_color: "#C83237"
display_name: Shodan
name: shodan
description: Steampipe plugin to query host, DNS and exploit information using Shodan.
og_description: "Query Shodan with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/shodan-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Shodan + Steampipe

[Shodan](https://www.shodan.io/) provides host, DNS and exploit information from scanning the Internet.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select ip,port,timestamp,dns from shodan_host_service where ip = '8.8.8.8'
```
```
+---------+------+---------------------+--------------------------------------------------------------------------------+
| ip      | port | timestamp           | dns                                                                            |
+---------+------+---------------------+--------------------------------------------------------------------------------+
| 8.8.8.8 | 53   | 2021-08-06 07:12:16 | {"recursive":true,"resolver_hostname":null,"resolver_id":null,"software":null} |
+---------+------+---------------------+--------------------------------------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/shodan/tables)**

## Get started

### Install

Download and install the latest Shodan plugin:

```shell
steampipe plugin install shodan
```

### Credentials

Shodan requires an API token for all requests, but offers a free tier. Sign up on the [Shodan website](https://www.shodan.io) to get your free token. It looks like `ZGloRBAl2Tvur3tBTu84NkZIf3i5Cc5U`.

### Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

Installing the latest shodan plugin will create a default connection named `shodan` in the `~/.steampipe/config/shodan.spc` file.  You must edit this connection to include your API token:

```hcl
connection "shodan" {
  plugin  = "shodan"
  api_key = "ZGloRBAl2Tvur3tBTu84NkZIf3i5Cc5U"
}
```


