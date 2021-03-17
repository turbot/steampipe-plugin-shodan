# Table: shodan_scan

List all the on-demand scans active in the account.

## Examples

### List the scans

```sql
select
  *
from
  shodan_scan
```

### List scans that are not complete

```sql
select
  *
from
  shodan_scan
where
  status != 'DONE'
```
