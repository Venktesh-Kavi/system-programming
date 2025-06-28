# PostgreSQL Autovacuum Optimization in Temporal Database

## Problem Statement

Our Temporal production database running on PostgreSQL RDS faces performance challenges due to heavy write operations and dead tuple accumulation.

### Background
- Temporal writes large amounts of data, causing frequent table updates
- Updates create dead tuples (outdated row versions)
- PostgreSQL's autovacuum process removes these dead tuples
- The analyze phase requires reading existing data, increasing read IOPS

### Issues
- Peak activity causes read IOPS to exceed provisioned limits
- Dead tuple accumulation leads to:
  - Increased disk usage
  - Slower queries
  - Performance degradation
  - Higher database load during analyze phase

![high read io][high_read_io.png]
![perf insights][perf_insights_dashboard.png]

## Solution

We optimized the autovacuum process through parameter tuning and testing:

### Approach
1. Tested different autovacuum parameters under production-like workload
2. Increased resources for autovacuum operations
3. Fine-tuned key parameters for better efficiency

![Default vs Changed Parameter IO Performance][default_vs_changed_params.png]
![Parameteres Changes][param_changes.png]

### Parameters Modified
- `autovacuum_vacuum_cost_limit`
- `autovacuum_work_mem`
- `maintenance_work_mem`
- `autovacuum_max_workers`

### Results
- 40% reduction in read IOPS
- 40% reduction in total IOPS usage
- Improved database performance and stability
- Prevention of IOPS spikes

## Parameter Details

### autovacuum_vacuum_cost_limit
- **Purpose**: Controls CPU and I/O resource consumption by autovacuum workers
- **Default**: `GREATEST({log(DBInstanceClassMemory/21474836480)*600},200)` units
- **Our Configuration**: Set to 2000 (50% of total IOPS)
- **Result**: Significant performance improvement in AutoVacuum execution

### autovacuum_work_mem
- **Purpose**: Controls memory used by autovacuum for maintenance tasks
- **Default**: 64 MB in Amazon RDS for PostgreSQL

### maintenance_work_mem
- **Purpose**: Controls memory for maintenance operations (VACUUM, ANALYZE, index creation)
- **Default**: 64 MB
- **Recommendation**: At least 1 GB for tables with many dead tuples

### autovacuum_max_workers
- **Purpose**: Controls maximum number of autovacuum worker processes
- **Default**: `GREATEST(DBInstanceClassMemory/64371566592,3)` workers

## References

1. [AWS Case Study: Tuning Autovacuum in RDS PostgreSQL](https://aws.amazon.com/blogs/database/a-case-study-of-tuning-autovacuum-in-amazon-rds-for-postgresql/)
2. [Understanding Autovacuum in RDS PostgreSQL](https://aws.amazon.com/blogs/database/understanding-autovacuum-in-amazon-rds-for-postgresql-environments/)
3. [RDS PostgreSQL Common DBA Tasks](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.PostgreSQL.CommonDBATasks.Parameters.html)
4. [Autovacuum Tuning Basics - EnterpriseDB](https://www.enterprisedb.com/blog/autovacuum-tuning-basics)
