---
title: "Steampipe Table: shodan_service - Query Shodan Service using SQL"
description: "Allows users to query Shodan Services, specifically the service banners, providing insights into service details and potential vulnerabilities."
---

# Table: shodan_service - Query Shodan Service using SQL

Shodan is a search engine for internet-connected devices. It provides information on devices connected to the internet, including their location, operating system, software, and vulnerabilities. Shodan's service banners provide a snapshot of a device's state at the time it was indexed, including details about the server software and available services.

## Table Usage Guide

The `shodan_service` table provides insights into services indexed by Shodan. As a security analyst, explore service-specific details through this table, including service banners, server software, and available services. Utilize it to uncover information about potential vulnerabilities, the state of the device at the time it was indexed, and the location of internet-connected devices.

## Examples

### List the services
Analyze the settings to understand the arrangement of services based on their port numbers. This allows you to pinpoint the specific locations where services are operating, aiding in network management and security.

```sql
select
  *
from
  shodan_service
order by
  port
```

### Find the service for a port
Identify the specific service operating on a given port number. This is useful for understanding what software is running on your network and can help with network management and security.

```sql
select
  *
from
  shodan_service
where
  port = 5432
```

### Get information about the Puppet service
Explore which services are associated with Puppet, a configuration management tool. This query is useful in identifying instances where Puppet is being used, which can aid in network management and security audits.

```sql
select
  *
from
  shodan_service
where
  name ilike '%puppet%'
```