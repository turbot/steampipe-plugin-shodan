---
title: "Steampipe Table: shodan_host_service - Query Shodan Host Services using SQL"
description: "Allows users to query Shodan Host Services, providing insights into the services running on a particular host, including the port, transport protocol, and product information."
---

# Table: shodan_host_service - Query Shodan Host Services using SQL

Shodan is a search engine for internet-connected devices. It provides a wealth of information about hosts, including open ports, protocols, and services. This information is crucial for security researchers, network administrators, and penetration testers to understand the digital footprint of a host.

## Table Usage Guide

The `shodan_host_service` table provides insights into the services running on a specific host within the Shodan search engine. As a security researcher or network administrator, explore service-specific details through this table, including ports, protocols, and product information. Utilize it to uncover information about a host's digital footprint, such as open ports and the services running on them.

**Important Notes**
- You must specify the `ip` in the `where` clause to query this table.

## Examples

### List all service information for an IP
Explore which services are associated with a specific IP address. This can be useful for understanding the functionality and potential vulnerabilities of the device or system associated with that IP.

```sql+postgres
select
  *
from
  shodan_host_service
where
  ip = '8.8.8.8';
```

```sql+sqlite
select
  *
from
  shodan_host_service
where
  ip = '8.8.8.8';
```

### SSL certificate details for services
Analyze the settings to understand the SSL certificate details for specific services on a given IP address. This is useful for ensuring secure connections by checking the validity and details of SSL certificates.

```sql+postgres
select
  ip,
  port,
  ssl->'cert' as ssl_cert
from
  shodan_host_service
where
  ip = '140.82.112.4'
  and ssl is not null;
```

```sql+sqlite
select
  ip,
  port,
  json_extract(ssl, '$.cert') as ssl_cert
from
  shodan_host_service
where
  ip = '140.82.112.4'
  and ssl is not null;
```

### Check Heartbleed status for each service
Determine the Heartbleed vulnerability status for each service on a specific IP address. This can be crucial in identifying potential security risks and taking appropriate measures to mitigate them.

```sql+postgres
select
  ip,
  port,
  opts->>'heartbleed' as heartbleed
from
  shodan_host_service
where
  ip = '140.82.112.4';
```

```sql+sqlite
select
  ip,
  port,
  json_extract(opts, '$.heartbleed') as heartbleed
from
  shodan_host_service
where
  ip = '140.82.112.4';
```