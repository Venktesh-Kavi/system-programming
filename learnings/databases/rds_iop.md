# Amazon RDS Storage Types and IOPS [WIP]

## Storage Types Overview

### 1. General Purpose SSD (gp2/gp3)
* **Best for**: Development and testing environments
* **Features**:
  * gp2:
    * Baseline of 3 IOPS/GB
    * Bursts up to 3000 IOPS
    * Volume size: 20GB to 64TB
  * gp3:
    * Baseline of 3000 IOPS
    * Can provision up to 16000 IOPS
    * Independent IOPS scaling

### 2. Provisioned IOPS SSD (io1)
* **Best for**: Production environments
* **Features**:
  * Consistent performance
  * Up to 64,000 IOPS per volume
  * User-specified IOPS rate
  * Best for OLTP workloads

### 3. Magnetic Storage (Standard)
* **Status**: Legacy option, not recommended
* **Limited to**: Average of 100 IOPS
* **Best for**: Infrequently accessed workloads

## Understanding IOPS

### What are IOPS?
* Input/Output Operations Per Second
* Measures the number of read/write operations to storage
* Each operation can be up to 256KB in size

### IOPS Calculation Rules
1. **Operation Size**:
   * Operations ≤ 256KB = 1 IOPS
   * Operations > 256KB = Math.ceil(size/256KB) IOPS
   * Example: 512KB operation = 2 IOPS

2. **Performance Impact**:
   * Larger I/O operations consume more IOPS
   * Small, random I/O typically more expensive
   * Sequential operations more efficient

### IOPS Limits and Considerations
* **Volume Caps**:
  * gp2: 3 IOPS/GB, max 16,000
  * gp3: Base 3,000, max 16,000
  * io1: Up to 64,000

* **Monitoring Tips**:
  * Watch for IOPS throttling
  * Monitor I/O operation sizes
  * Consider upgrading storage type if consistently hitting limits

## RDS Performance Optimization

### Storage Striping
* RDS automatically stripes data across multiple volumes
* Benefits:
  * Improved I/O performance
  * Better throughput
  * Reduced latency

### Best Practices
1. **Storage Selection**:
   * Use gp3 for most workloads
   * Choose io1 for IOPS-intensive applications
   * Monitor and adjust based on workload

2. **Performance Monitoring**:
   * Track IOPS utilization
   * Monitor storage throughput
   * Watch for storage-related latency



## References
* [Understanding IOPS in Detail](https://cloudcasts.io/article/what-you-need-to-know-about-iops)
* [Amazon RDS Storage Types](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html)
