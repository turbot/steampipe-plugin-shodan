# Table: shodan_domain

Get all the subdomains and other DNS entries for the given domain.

Note: A `domain` must be provided in all queries to this table.

## Examples

### Reverse DNS lookup

```sql
select
  *
from
  shodan_domain
where
  domain = 'steampipe.io'
```

### Find all A records for the domain and its subdomains

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

```sql
select
  *
from
  shodan_domain
where
  domain = 'github.com'
  and subdomain is null
```
