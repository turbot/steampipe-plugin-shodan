# Table: shodan_protocol

List all protocols that can be used when performing on-demand Internet scans via Shodan.

## Examples

### List the protocols

```sql
select
  *
from
  shodan_protocol
order by
  name
```

### Get the Postgres protocol

```sql
select
  *
from
  shodan_protocol
where
  name = 'postgresql'
```
