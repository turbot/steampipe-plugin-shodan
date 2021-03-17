# Table: shodan_service

List of the services Shodan can detect.

## Examples

### List the services

```sql
select * from shodan_service order by port
```

### Find the service for a port

```sql
select * from shodan_service where port = 5432
```

### Get information about the Puppet service

```sql
select * from shodan_service where name ilike '%puppet%'
```
