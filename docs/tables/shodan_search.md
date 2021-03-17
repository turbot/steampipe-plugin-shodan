# Table: shodan_search

Search the Internet for information about hosts and their services. Please see the [examples](https://beta.shodan.io/search/examples), a [cheat sheet](https://thedarksource.com/shodan-cheat-sheet/) and the [filter reference](https://beta.shodan.io/search/filters) to understand the query options.

Note: A `query` must be provided in all queries to this table.

## Examples

### List all services for a network range

```sql
select
  *
from
  shodan_search
where
  query = 'net:34.98.0.0/26'
```

### Find all Windows XP hosts

```sql
select
  *
from
  shodan_search
where
  query = 'os:"windows xp"'
```

### Find all services for the GitHub organization

```sql
select
  *
from
  shodan_search
where
  query = 'org:"GitHub"'
```
