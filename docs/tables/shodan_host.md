# Table: shodan_host

Get ISP, geolocation and other information about the host at a given IP address.

Note: An `ip` must be provided in all queries to this table.

## Examples

### Basic host information

```sql
select
  *
from
  shodan_host
where
  ip = '8.8.8.8'
```

### Basic host information

```sql
select
  *
from
  shodan_host
where
  ip = '8.8.8.8'
```

### Services open on the host

```sql
select
  ip,
  s.*
from
  shodan_host as h,
  jsonb_array_elements_text(h.ports) as host_port,
  shodan_service as s
where
  ip = '8.8.8.8'
  and host_port::bigint = s.port
```

### Location of the host

```sql
select
  ip,
  city,
  country_code
from
  shodan_host
where
  ip = '8.8.8.8'
```
