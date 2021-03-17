# Table: shodan_host_service

Get full information about all services discovered running on the host at a given IP address.

Note: An `ip` must be provided in all queries to this table.

## Examples

### List all service information for an IP

```sql
select
  *
from
  shodan_host_service
where
  ip = '8.8.8.8'
```

### SSL certificate details for services

```sql
select
  ip,
  port,
  ssl->'cert' as ssl_cert
from
  shodan_host_service
where
  ip = '140.82.112.4'
  and ssl is not null
```

### Check Heartbleed status for each service

```sql
select
  ip,
  port,
  opts->>'heartbleed'
from
  shodan_host_service
where
  ip = '140.82.112.4'
```
