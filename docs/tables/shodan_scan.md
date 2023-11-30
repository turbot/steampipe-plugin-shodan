---
title: "Steampipe Table: shodan_scan - Query Shodan Scan using SQL"
description: "Allows users to query Shodan Scans, specifically providing details about each individual scan performed on the Shodan platform."
---

# Table: shodan_scan - Query Shodan Scan using SQL

Shodan is a search engine that lets the user find specific types of computers (webcams, routers, servers, etc.) connected to the internet using a variety of filters. Shodan Scan is a feature of the Shodan platform that allows users to scan an IP/ netblock. It provides detailed information about the services running on the devices and their vulnerabilities.

## Table Usage Guide

The `shodan_scan` table provides insights into individual scans performed on the Shodan platform. As a security analyst, you can explore scan-specific details through this table, including IP addresses, hostnames, and associated metadata. Utilize it to uncover information about the services running on the devices, their vulnerabilities, and to monitor the security of your network infrastructure.

## Examples

### List the scans
Uncover the details of all the completed network scans to analyze potential vulnerabilities or security threats. This is useful in maintaining the security posture of your network and proactively addressing potential risks.

```sql
select
  *
from
  shodan_scan
```

### List scans that are not complete
Determine the areas in which scans are still in progress or have not been completed, allowing for improved tracking and management of scanning operations.

```sql
select
  *
from
  shodan_scan
where
  status != 'DONE'
```