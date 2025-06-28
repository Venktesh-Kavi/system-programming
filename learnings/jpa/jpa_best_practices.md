# JPA Best Practices

## Introduction
Hibernate shifts the developer mindset from SQL statements to entity state transitions. Once an entity is actively managed by Hibernate, all changes are automatically propagated to the database.

## Flush Ordering
* JPA has a specific order for executing database operations
* Example: Even if delete is performed before insert in code, JPA executes insert first
* This is due to the internal flush ordering mechanism

> [Read more about flush ordering](https://vladmihalcea.com/hibernate-facts-knowing-flush-operations-order-matters/)

## Relationship Mapping Best Practices

### One-to-Many Relationships
* **Best Practices**:
  * Let the many-to-one side handle state transitions
  * Use `@JoinColumn` on the owning (child) side
  * Use `mappedBy` in parent for bi-directional relationships
  * Avoid `@OneToMany` for large collections
    * Use queries instead for pagination/restriction
    * Hibernate cannot limit collection loading

* **Operation Flow**:
  1. Post entity insertion
  2. Post comments insertion
  3. Post comment update with post ID
  * Follows flush ordering: persist action first, then collection elements

* **Deletion Process**:
  1. Update `post_id` to null in `post_comments`
  2. Execute delete operation
  3. Process orphan removal

> [Detailed guide on One-to-Many mapping](https://vladmihalcea.com/the-best-way-to-map-a-onetomany-association-with-jpa-and-hibernate/)

### MappedBy Attribute
* Indicates that a column is mapped by an attribute in another table
* Used in bi-directional relationships
* Specifies the owning side of the relationship

### Child Entity Operations
* **Best Practices for Inserts**:
  * Always get parent reference for child insertion
  * Use `getByReferenceId()` to avoid SELECT on parent
  * Reduces unnecessary database operations

### Many-to-Many Relationships
* **Key Points**:
  * Uses a joining table
  * Prefer `Set` over `List` to avoid unnecessary inserts
  * Careful with cascade types:
    * Use `CascadeType.PERSIST` and `CascadeType.MERGE`
    * Avoid `CascadeType.ALL`
    * Prevents unintended deletions (e.g., shared books when deleting an author)

> [Many-to-Many best practices](https://vladmihalcea.com/the-best-way-to-use-the-manytomany-annotation-with-jpa-and-hibernate/)

## Cascading
* **Best Practices**:
  * Use cascading primarily on parent side
  * Avoid cascading on child side (considered a mapping smell)
  * Choose cascade types carefully based on relationship semantics

> [Comprehensive guide to cascade types](https://vladmihalcea.com/a-beginners-guide-to-jpa-and-hibernate-cascade-types/)

## Performance Considerations
* Use pagination for large collections
* Prefer lazy loading for relationships
* Consider using DTOs for specific use cases
* Use batch processing for bulk operations
