---
title: "Steampipe Table: shodan_port - Query Shodan Ports using SQL"
description: "Allows users to query Shodan Ports, specifically providing details about open ports detected by Shodan, the internet's search engine for IoT devices."
---

# Table: shodan_port - Query Shodan Ports using SQL

Shodan is a search engine that lets the user find specific types of computers (webcams, routers, servers, etc.) connected to the internet using a variety of filters. Some have also described it as a search engine of service banners, which are metadata that the server sends back to the client. This can be information about the server software, what options the service supports, a welcome message or anything else that the client can find out before interacting with the server.

## Table Usage Guide

The `shodan_port` table provides insights into open ports detected by Shodan, the internet's search engine for IoT devices. As a cybersecurity analyst, explore port-specific details through this table, including the services running on these ports, the devices they are connected to, and their geographical locations. Utilize it to uncover information about potential vulnerabilities, such as unsecured ports, outdated server software, and the presence of IoT devices on the network.

## Examples

### List the ports
Explore which ports are being used in your network. This can help in identifying potential vulnerabilities or unusual activity for security purposes.

```sql+postgres
select
  *
from
  shodan_port
order by
  port;
```

```sql+sqlite
select
  *
from
  shodan_port
order by
  port;
```