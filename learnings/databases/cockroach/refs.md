Cockroach DB underlying uses pebble (Rocks DB) as the KV store for writing and retrieving data.
An uniform SQL API is exposed, internally they are coveted to KV and eventually stored or retrieved from pebble.
Each node can have multiple stores.
In case of a disk stall, WAL failover can be configured to another store in a different node.

Questions
What is Rocks DB?
Motivation to Pebble?
If each node has multiple stores, does the CDC work over ranges? How do we track till what point we have read?
Which store to tap for CDC?


References

Strimzi - Strimzi is an open source tool for managing kafka clusters on kubernetes/openshift. It provides a set of operators to manage the kafka cluster. Ref: [https://strimzi.io/docs/operators/latest/overview]
Cockroach Architecture: https://github.com/cockroachdb/cockroach/blob/master/docs/design.md
Pebble DB: https://www.cockroachlabs.com/blog/pebble-rocksdb-kv-store/
Use case of Pebble over Rocks: https://www.cockroachlabs.com/blog/bulk-data-import/
RocksDB 101: https://zhangyuchi.gitbooks.io/rocksdbbook/content/RocksDB-Basics.html 
SMT Connectors