# DuckDB Overview [WIP]

## Introduction
* DuckDB is the "SQLite for OLAP operations"
* Provides SQL interface for querying file data
* In-memory analytical database optimized for analytics

## Core Concepts

### Process Architecture
* **Process Bound**:
  * Each process creates its own connection
  * Gets a DuckDB object with in-memory file access
  * No built-in connection sharing between hosts

### Data Flow
```
Source File → DuckDB Transformation (table creations) → .duckdb (meta file)
```

### Storage Options
* **File Formats**:
  * CSV
  * Parquet (memory efficient, queryable)
  * Native .duckdb files
* **Cloud Storage**:
  * Can store transformed files in S3
  * One file per MCD/resource

## Best Practices

### Performance Optimization
1. **Data Loading**:
   * Route file-specific jobs to dedicated workers
   * Minimize redundant memory loading

2. **Bulk Operations**:
   * Use `pg_copy` for bulk database inserts
   * Avoid batch-based insertions
   * Optimal flow: `source file → DuckDB → validate → CSV → pg_copy`

3. **Partitioning Strategy**:
   * Use SQL-based computations when possible
   * Consider partitioning only for complex computations
   * No significant benefit in partition-by-partition processing for simple loads

### Workflow Considerations
* **Process Flow**:
  ```
  Load Source → DuckDB → Validations → Database Insert
  ```

* **Bulk Framework Limitations**:
  * Validations may extend beyond framework scope
  * Context-dependent validation challenges
  * Consider use cases like Colending

## Performance
* **CSV Processing Speed**:
  * Competitive performance for analytical workloads
  * [Detailed benchmarks](https://datapythonista.me/blog/how-fast-can-we-process-a-csv-file)

## Limitations
* No native multi-host connection sharing
* Process-bound architecture
* Memory constraints based on host capacity

## Best Use Cases
1. **Analytics Workflows**:
   * Data transformation
   * File-based analytics
   * ETL processes

2. **Validation Pipelines**:
   * Data quality checks
   * Schema validation
   * Business rule verification

## References
* [DuckDB Overview](https://duckdb.org/docs/)
* [DuckDB CSV Load Speed](https://datapythonista.me/blog/how-fast-can-we-process-a-csv-file)
