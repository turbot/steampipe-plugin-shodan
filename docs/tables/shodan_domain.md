---
title: "Steampipe Table: shodan_domain - Query Shodan Domains using SQL"
description: "Allows users to query Shodan Domains, specifically providing information about domain names and the servers they point to."
---

# Table: shodan_domain - Query Shodan Domains using SQL

Shodan is a search engine that lets the user find specific types of computers (webcams, routers, servers, etc.) connected to the internet using a variety of filters. Some have also described it as a search engine of service banners, which are metadata that the server sends back to the client. The Shodan Domain is a resource in Shodan that provides information about domain names and the servers they point to.

## Table Usage Guide

The `shodan_domain` table provides insights into domain names and the servers they point to within Shodan. As a network administrator or security analyst, explore domain-specific details through this table, including server metadata, IP addresses, and associated hostnames. Utilize it to uncover information about servers, such as their geographical location, operating system, and open ports, to help in identifying potential security vulnerabilities or infrastructure improvements.

## Examples

### Reverse DNS lookup
Analyze the settings to understand the domain configuration for 'steampipe.io'. This can be useful in identifying potential security vulnerabilities or misconfigurations.

```sql
select
  *
from
  shodan_domain
where
  domain = 'steampipe.io'
```

### Find all A records for the domain and its subdomains
Explore all A records associated with a specific domain and its subdomains. This is useful for understanding the IP addresses linked to a domain, aiding in network mapping and security assessments.

```sql
select
  *
from
  shodan_domain
where
  domain = 'github.com'
  and type = 'A'
```

### Get all records for the domain, without subdomain enumeration
Explore the main records associated with a specific domain, without including any subdomains. This can be useful when you need an overview of the main domain's details, without the potential clutter of subdomain information.

```sql
select
  *
from
  shodan_domain
where
  domain = 'github.com'
  and subdomain is null
```