# MongoDB Best Practices and Patterns [WIP]

## Document Relationships

### Embedding vs Referencing
* Choose based on querying patterns:
  * **Embedding**: When data is always accessed together
  * **Referencing**: When data needs to be accessed independently

### Document References in Java
* Two main annotations:
  * `@DocRef`:
    * References via `_id` field only
    * Legacy approach
  * `@DocumentReference`:
    * References via any field
    * Modern recommended approach
    * More flexible referencing options

### Example: Publisher-Books Relationship

```java
// Anti-pattern example with embedding
class Publisher {
    private String publisherName;
    private String publisherId;
    private List<Book> books;  // Unbounded array - potential issue
}

class Book {
    private String isbn;
    private String bookName;
}
```

## Schema Design Patterns

### Anti-Patterns to Avoid

1. **Massive Arrays**
   * Problem: Unbounded array growth
   * Impact:
     * Degraded index performance
     * Slower document updates
     * Memory usage issues

2. **Inappropriate Embedding**
   * Example: Books embedded in Publisher
   * Issues:
     * Growing array size
     * Performance degradation
     * Memory consumption

### Solutions and Trade-offs

1. **Flip Embedding Direction**
   * Example: Embed building in employees instead of employees in building
   * Pros:
     * Better control over array size
     * Improved query performance
   * Cons:
     * Data duplication
     * Update overhead

2. **Separate Collections with References**
   * Use `$ref` for linking documents
   * Pros:
     * No array size issues
     * Independent document management
   * Cons:
     * Requires `$lookup` operations
     * Join operations can be expensive

3. **Extended Reference Pattern**
   * Hybrid approach
   * Features:
     * Partial data duplication
     * Balance between performance and data consistency
     * Reduced need for joins

## Best Practices Summary

1. **Data Modeling**:
   * Consider access patterns first
   * Avoid unbounded arrays
   * Use appropriate referencing strategy

2. **Performance Optimization**:
   * Monitor array sizes
   * Index strategically
   * Consider read/write ratios

3. **Relationship Handling**:
   * Use embedding for 1:1 or small 1:N
   * Use references for large 1:N or N:N
   * Consider extended references for frequent reads

## References
* [MongoDB Schema Design Anti-Patterns](https://www.mongodb.com/developer/products/mongodb/schema-design-anti-pattern-summary/)
* [MongoDB Data Modeling](https://docs.mongodb.com/manual/core/data-modeling-introduction/)

