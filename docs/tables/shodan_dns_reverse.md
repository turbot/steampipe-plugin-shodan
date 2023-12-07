---
title: "Steampipe Table: shodan_dns_reverse - Query Shodan DNS Reverse using SQL"
description: "Allows users to query DNS Reverse in Shodan, specifically returning information about the hostnames that have been defined for the given IP, providing insights into domain mapping and potential anomalies."
---

# Table: shodan_dns_reverse - Query Shodan DNS Reverse using SQL

Shodan DNS Reverse is a feature within Shodan that allows you to retrieve hostnames that have been defined for the given IP. It provides a comprehensive way to map and monitor domain records for various IP addresses. Shodan DNS Reverse helps you stay informed about the associated hostnames for your IP resources and take appropriate actions when predefined conditions are met.

**Important Notes**
- You must specify the `ip` in the `where` clause to query this table.

## Table Usage Guide

The `shodan_dns_reverse` table provides insights into the DNS Reverse within Shodan. As a Network Administrator, explore IP-specific details through this table, including associated hostnames, and related metadata. Utilize it to uncover information about the domain mapping for your IP addresses, such as those with multiple hostnames, the relationships between IPs and hostnames, and the verification of domain records.

## Examples

### Reverse DNS lookup
Analyze the settings to understand the details associated with a specific IP address. This can be useful for identifying the ownership and configuration of a network resource.

```sql+postgres
select
  *
from
  shodan_dns_reverse
where
  ip = '8.8.8.8';
```

```sql+sqlite
select
  *
from
  shodan_dns_reverse
where
  ip = '8.8.8.8';
```