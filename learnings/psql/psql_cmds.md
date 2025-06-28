
### Replication Slot

```sql
select slot_name, database, active, pg_size_pretty(pg_wal_lsn_diff(pg_current_wal_lsn(), restart_lsn)) as replication_lag from pg_replication_slots order by replication_lag desc;
```

### Pg Stat Activity

Provides information of currently connected processes, users, running queries
```select * from pg_start_activity;```
```select pid, usename, query_start, state, query from pg_stat_activity where state != 'idle';```


### Max Connections

`SHOW max_connections;`

### Get Available Connections and Used connections
```sql
SELECT
    max_conn,
    used,
    reserved_connections,
    max_conn - used - reserved_connections AS available
FROM
    (SELECT setting::int AS max_conn FROM pg_settings WHERE name = 'max_connections') AS max_connections,
    (SELECT count(*) AS used FROM pg_stat_activity) AS used_connections,
    (SELECT setting::int AS reserved_connections FROM pg_settings WHERE name = 'superuser_reserved_connections') AS reserved;
```
