# Table: shodan_dns_reverse

Reverse DNS lookup for a given IP address.

Note: An `ip` must be provided in all queries to this table.

## Examples

### Reverse DNS lookup

```sql
select
  *
from
  shodan_dns_reverse
where
  ip = '8.8.8.8'
```
