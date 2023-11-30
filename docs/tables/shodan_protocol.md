---
title: "Steampipe Table: shodan_protocol - Query Shodan Protocols using SQL"
description: "Allows users to query Shodan Protocols, providing insights into protocol details such as name, description, and port."
---

# Table: shodan_protocol - Query Shodan Protocols using SQL

Shodan is a search engine for internet-connected devices. It provides information about devices, including their location, type, operating system, and security status. Shodan's protocol is a crucial resource that provides details about various protocols used by these devices, including their names, descriptions, and ports.

## Table Usage Guide

The `shodan_protocol` table provides insights into the protocols used by internet-connected devices indexed by Shodan. As a cybersecurity analyst, you can explore protocol-specific details through this table, including names, descriptions, and associated ports. Utilize it to uncover information about protocols, such as those commonly used by specific types of devices, the description of each protocol, and the ports they typically use.

## Examples

### List the protocols
Explore the various protocols to understand their order and organization within the system. This can help in managing and optimizing system communications.

```sql
select
  *
from
  shodan_protocol
order by
  name
```

### Get the Postgres protocol
Explore which systems are using the PostgreSQL protocol to understand its distribution and usage across your network. This can be useful for maintaining protocol consistency and troubleshooting network issues.

```sql
select
  *
from
  shodan_protocol
where
  name = 'postgresql'
```