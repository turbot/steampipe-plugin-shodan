---
title: "Steampipe Table: shodan_search - Query Shodan Search Results using SQL"
description: "Allows users to query Shodan Search Results, specifically the details of internet-connected devices, providing insights into device security and exposure."
---

# Table: shodan_search - Query Shodan Search Results using SQL

Shodan is a search engine for internet-connected devices. It provides a way to discover devices, where they are located and who is using them. Shodan is used for various purposes including cybersecurity research, software development, and educational research.

## Table Usage Guide

The `shodan_search` table provides insights into internet-connected devices as indexed by Shodan. As a cybersecurity analyst, explore device-specific details through this table, including IP addresses, hostnames, and potential vulnerabilities. Utilize it to uncover information about devices, such as their geographic location, the software they are running, and their exposure to potential cyber threats.

**Important Notes**
- You must specify the `query` in the `where` clause to query this table.

## Examples

### List all services for a network range
Explore all the services associated with a specific network range. This is useful for gaining insights into the various operations within a particular network segment, helping to better manage and secure your network infrastructure.

```sql+postgres
select
  *
from
  shodan_search
where
  query = 'net:34.98.0.0/26';
```

```sql+sqlite
Error: SQLite does not support CIDR operations.
```

### Find all Windows XP hosts
Determine the areas in which Windows XP is still being used to understand potential security vulnerabilities and outdated systems in your network.

```sql+postgres
select
  *
from
  shodan_search
where
  query = 'os:"windows xp"';
```

```sql+sqlite
select
  *
from
  shodan_search
where
  query = 'os:"windows xp"';
```

### Find all services for the GitHub organization
Discover all the services associated with a particular organization, in this case, GitHub. This is useful for gaining insights into the various services that an organization utilizes or is associated with.

```sql+postgres
select
  *
from
  shodan_search
where
  query = 'org:"GitHub"';
```

```sql+sqlite
select
  *
from
  shodan_search
where
  query = 'org:"GitHub"';
```