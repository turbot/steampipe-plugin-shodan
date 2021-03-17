---
organization: Turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/shodan.svg"
brand_color: "#B94039"
display_name: Shodan
name: shodan
description: Steampipe plugin to query host, DNS and exploit information using Shodan.
---

# Shodan

Shodan provides host, DNS and exploit information from scanning the Internet.


## Installation

To download and install the latest shodan plugin:

```bash
steampipe plugin install shodan
```

## Credentials

Shodan requires an API token for all requests, but offers a free tier. Sign up on the [shodan website](https://shodan.com) to get your free token. It looks like `ZGloRBAl2Tvur3tBTu84NkZIf3i5Cc5U`.


## Connection Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

Installing the latest shodan plugin will create a default connection named `shodan` in the `~/.steampipe/config/shodan.spc` file.  You must edit this connection to include your API token:

```hcl
connection "shodan" {
  plugin  = "shodan"
  api_key = "ZGloRBAl2Tvur3tBTu84NkZIf3i5Cc5U"
}
```
