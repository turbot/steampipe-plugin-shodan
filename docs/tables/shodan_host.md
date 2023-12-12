---
title: "Steampipe Table: shodan_host - Query Shodan Hosts using SQL"
description: "Allows users to query Shodan Hosts, specifically providing details about devices connected to the internet, their characteristics, and potential vulnerabilities."
---

# Table: shodan_host - Query Shodan Hosts using SQL

Shodan is a search engine that lets users find specific types of computers connected to the internet using a variety of filters. Some have also described it as a search engine of service banners, which are metadata that the server sends back to the client. This can be information about the server software, what options the service supports, a welcome message or anything else that the client can find out before interacting with the server.

## Table Usage Guide

The `shodan_host` table provides insights into devices connected to the internet and their characteristics. As a security analyst, you can explore device-specific details through this table, including their IP addresses, hostnames, operating systems, and potential vulnerabilities. Utilize it to uncover information about devices, such as their geographical locations, open ports, and the services running on them.

**Important Notes**
- You must specify the `ip` in the `where` clause to query this table.

## Examples

### Basic host information
Analyze the settings to understand the basic details of a specific host, such as its location, operating system, and open ports. This can be useful for network administrators to assess the security posture of their systems.

```sql+postgres
select
  *
from
  shodan_host
where
  ip = '8.8.8.8';
```

```sql+sqlite
select
  *
from
  shodan_host
where
  ip = '8.8.8.8';
```

### Basic host information
Explore the basic information associated with a specific IP address. This can be useful to understand the characteristics of a host, which can aid in network management and security assessments.

```sql+postgres
select
  *
from
  shodan_host
where
  ip = '8.8.8.8';
```

```sql+sqlite
select
  *
from
  shodan_host
where
  ip = '8.8.8.8';
```

### Services open on the host
Explore which services are currently open on a specific host. This is useful for understanding potential vulnerabilities and security risks associated with open services on a host.

```sql+postgres
select
  ip,
  s.*
from
  shodan_host as h,
  jsonb_array_elements_text(h.ports) as host_port,
  shodan_service as s
where
  ip = '8.8.8.8'
  and host_port::bigint = s.port;
```

```sql+sqlite
select
  ip,
  s.*
from
  shodan_host as h,
  json_each(h.ports) as host_port,
  shodan_service as s
where
  ip = '8.8.8.8'
  and host_port.value = s.port;
```

### Location of the host
Analyze the geographical details of a specific internet host. This is useful for understanding the physical location of a host, which can be essential in security analysis or network management scenarios.

```sql+postgres
select
  ip,
  city,
  country_code
from
  shodan_host
where
  ip = '8.8.8.8';
```

```sql+sqlite
select
  ip,
  city,
  country_code
from
  shodan_host
where
  ip = '8.8.8.8';
```