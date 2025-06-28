# PostgreSQL Replication Slots [WIP]

Replication slots are pointers to the WAL (Write-Ahead Log). Each connector gets its own slot to track WAL consumption. PostgreSQL won't remove WAL files until all slots have processed them.

## Production Issue in OCI PostgreSQL

Hit severe performance issues:
- 100% CPU and disk usage
- No disk space available
- Connection failures

Root cause: 12 inactive replication slots causing WAL retention and DB growth.

Fixes:
- Cleared inactive slots -> recovered disk space
- 2 vCPU -> 4 vCPU for connection issues

Note: OCI missing ReplicationSlotDiskUsage metric for monitoring.

## Replica Identity Problems

Seen during gRPC updates:
- Replica identity not set to FULL
- Depends on how slot was created (implicit/explicit)
- Updates need both old/new transaction records

## Things to Watch

Monitoring:
- Active vs inactive slots
- Disk usage from WAL retention
- Replication lag

Maintenance:
- Clean up inactive slots regularly
- Check replica identity settings
- Monitor WAL growth

## References
- [Table Replica Identity](https://www.artie.com/blogs/postgres-table-replica-identity)
- [Replication Slot Management](https://www.morling.dev/blog/insatiable-postgres-replication-slot/)
