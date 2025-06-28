Problem Statement:
Temporal Prod Database - Postgresql RDS DataBase

Temporal writes a large amount of data to the database, resulting in frequent updates to various tables. These updates create dead tuples, which are outdated row versions that remain in the database until they are cleaned up. PostgreSQL’s autovacuum process is responsible for removing these dead tuples to free up space and maintain performance. However, before vacuuming, the process must first analyze the table, which involves reading all the existing data. This analysis phase can cause a significant increase in read IOPS (Input/Output Operations Per Second), especially when large tables are involved.
During peak activity, the read IOPS usage can exceed the provisioned limit, leading to performance degradation. The autovacuum process is triggered based on the number of dead tuples and system-defined thresholds, ensuring that cleanup happens automatically. However, if too many dead tuples accumulate before autovacuum runs, the system may experience increased disk usage and slower queries. Additionally, the frequent reads during the analyze phase can put extra load on the database, further impacting its performance. Optimizing autovacuum settings and monitoring dead tuples is crucial to maintaining system stability.


Solution:
To optimize the autovacuum process, we tested different autovacuum parameters under SPOCTO's workload using a PostgreSQL database of the same size. we ensured that autovacuum runs more efficiently, reducing the impact of dead tuples on database performance. We specifically increased resources allocated to autovacuum, allowing it to clean up dead tuples faster and more effectively.
To optimize the autovacuum process, we tested and fine-tuned key autovacuum parameters to improve efficiency. We specifically increased the values of autovacuum_vacuum_cost_limit, autovacuum_work_mem, maintenance_work_mem, and autovacuum_max_workers to allow the vacuum process to run more effectively.
As a result of these optimizations, we observed a 40% reduction in read IOPS, meaning fewer disk reads were required during the analyze and vacuuming phases. Additionally, the total IOPS usage also dropped by 40%, improving overall database performance and stability. This optimization helped in preventing IOPS spikes, ensuring that the system operates within the provisioned limits without performance degradation.

Parameters:
All parameter Source: Working with parameters on your RDS for PostgreSQL DB instance - Amazon Relational Database Service
Autovacuum_vacuum_cost_limit:
Source : autovacuum_vacuum_cost_limit - AWS Prescriptive Guidance

The autovacuum_vacuum_cost_limit parameter controls the amount of CPU and I/O resources that an autovacuum worker can consume.
Default value: GREATEST({log(DBInstanceClassMemory/21474836480)*600},200) units of work
Note: Provisioned is 4000 IOPS, Alerts are setup at 3000 IOPS, We set the value at 2000 that's 50% of the total IOPS , we got performance improvement on AutoVacuum Execution.



Autovacuum_work_mem:
Source : autovacuum_work_mem - AWS Prescriptive Guidance

autovacuum_work_mem is a PostgreSQL configuration parameter that controls the amount of memory used by the autovacuum process when it performs table maintenance tasks such as vacuuming or analysis.

Default value: 64 MB in Amazon RDS for PostgreSQL.




Maintenance_work_mem:
Source : maintenance_work_mem - AWS Prescriptive Guidance

The maintenance_work_mem parameter controls the amount of memory used by maintenance operations such as VACUUM, ANALYZE, and index creation. 

Default value: 64 MB

Note: (This is common for both Autovacuum_work_mem and Maintenance_work_mem)
https://aws.amazon.com/blogs/database/understanding-autovacuum-in-amazon-rds-for-postgresql-environments/
As per the source above setting maintenance_work_mem to at least 1 GB will significantly improve the performance of vacuuming tables with a large number of dead tuples.


Autovacuum_max_workers:
Source :  autovacuum_max_workers - AWS Prescriptive Guidance

The autovacuum_max_workers parameter controls the maximum number of worker processes that the autovacuum process can create. Each worker process is responsible for vacuuming or analyzing a single table.

Default value: GREATEST(DBInstanceClassMemory/64371566592,3) workers
Note:
Reference for tuning Autovacuum_vacuum_cost_limit and Autovacuum_max_workers:
https://aws.amazon.com/blogs/database/a-case-study-of-tuning-autovacuum-in-amazon-rds-for-postgresql/









References:
https://aws.amazon.com/blogs/database/a-case-study-of-tuning-autovacuum-in-amazon-rds-for-postgresql/
https://aws.amazon.com/blogs/database/understanding-autovacuum-in-amazon-rds-for-postgresql-environments/

https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.PostgreSQL.CommonDBATasks.Parameters.html


https://www.enterprisedb.com/blog/autovacuum-tuning-basics
