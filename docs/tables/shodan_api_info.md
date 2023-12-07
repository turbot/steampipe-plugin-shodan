---
title: "Steampipe Table: shodan_api_info - Query Shodan API Information using SQL"
description: "Allows users to query Shodan API Information, specifically the details regarding the API's plan, usage, and capabilities."
---

# Table: shodan_api_info - Query Shodan API Information using SQL

Shodan is a search engine that lets users find specific types of computers (webcams, routers, servers, etc.) connected to the internet using a variety of filters. The Shodan API provides a method for accessing Shodan's raw data in order to integrate it into other applications. It allows users to explore the internet in real-time and provides a way to determine the geographical location of hosts, discover given services, or identify certain devices.

## Table Usage Guide

The `shodan_api_info` table provides insights into the Shodan API's plan, usage, and capabilities. As a security analyst, explore API-specific details through this table, including the plan's tier, the number of query credits available, and the scan credits remaining. Utilize it to monitor your usage of the Shodan API and to ensure it aligns with the capabilities of your current plan.

## Examples

### Get API status
Analyze the status of your APIs to ensure they're functioning as expected and to identify any potential issues that may need addressing. This can be crucial in maintaining smooth and efficient operations within your digital infrastructure.

```sql+postgres
select
  *
from
  shodan_api_info;
```

```sql+sqlite
select
  *
from
  shodan_api_info;
```