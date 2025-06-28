# PostgreSQL Data Types [WIP]

## Numeric Types

### Floating Point Types
* **Characteristics**:
  * Fixed precision and size
  * Inexact decimal representation
  * Types:
    * `real` - 4 bytes, 6 decimal digits precision
    * `double precision` - 8 bytes, 15 decimal digits precision
  * Use when exact precision isn't critical (e.g., scientific calculations)

### Arbitrary Precision Types
* **numeric/decimal**:
  * User-specified precision
  * Variable storage size
  * Exact decimal representation
  * Best for:
    * Financial calculations
    * Calculations requiring exact precision
    * Values where rounding errors are unacceptable

## JSON Types
* **jsonb**:
  * Maximum size: 255 MB
  * Binary storage format
  * Supports indexing
  * More efficient than `json` type
  * Best for:
    * Document storage
    * Complex data structures
    * When querying JSON content

## Common Data Type Sizes
* `smallint`: 2 bytes
* `integer`: 4 bytes
* `bigint`: 8 bytes
* `char(n)`: fixed-length, blank padded
* `varchar(n)`: variable-length with limit
* `text`: unlimited length
* `timestamp`: 8 bytes
* `uuid`: 16 bytes

## References
* [PostgreSQL Numeric Types](https://www.postgresql.org/docs/current/datatype-numeric.html)
* [Understanding Floating Point](https://floating-point-gui.de/basic/)
* [PostgreSQL Data Types Overview](https://www.postgresql.org/docs/current/datatype.html)

## Best Practices
* Use `numeric/decimal` for financial calculations
* Use `double precision` for scientific calculations
* Use `integer` for whole numbers unless `bigint` is specifically needed
* Use `text` instead of `varchar(n)` unless you need to enforce a length limit
* Use `jsonb` over `json` for better performance and functionality
