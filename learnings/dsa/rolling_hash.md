## Rolling Hash

* Rolling hash is a technique used to compute the incremental hash, instead of completely hashing from the scratch.
* Typically hashing is achieved by a modulo operation.
* Modulo operation is used for wrap within a range and say for eg strings represented as integers and exceeding the integers max limit (modulo fn can be used to bring it into range).
* Choosing a good modulo value/hashing function is important for getting a unique (non-clashing) hashed value.


In this case 10^9 might not be a good modulo value as it results in clashes.


Modulus operations fits data within a range. Eg.., modulo (m) of 1000 restricts the hash space from 0-999. hash values are wrapped around once they exceed m-1

Examples

```
A poor hash function: choosing a modulo value as 10^9

1. 1212 % 10^9 = 1212
2. 100001212 % 10^9 = 1212

Base 10, modulus = 1000

s = abc

hash = 10^2 * 1 + 10^1 * 2 + 10^0 * 3 = 123
hash with moduls = (10^2 * 1 + 10^1 * 2 + 10^0 * 3) % 1000 = 123


```

### Why multiple by a base?

To encode positional information
* Multiplying by a base is not just for randomness. Multiplying by a base gives each character in a string example a positional weight and makes ordering important. Two strings "abc" and "cab" produce the same hash if we do not multiply by a base

To create unique hashes
* Adding a base increases the randomess for producing a unique hash

To Enhance rolling hash
* Rolling hash (such as rabin-karp), the base simplifies updating the hash when new characters are added or removed.

### Why modulo?

* Modulo operation is used bring the hash within the desired range.
* Just multiplying by base can result in a possible overflow.
* It is also used to provide another dimension of randomness.


### Prefer Prime Number as Modulo Values/Hash Fn

* Suppose we have the strings: 12, 120, 1200, 12000. All of this results in the same hash because the base 10 matches with the structure of the data (powers of 10).
* Using a prime number = 31, m = 1000, the hashes of 12, 120, 1200, 12000 are unique and there are no collisions
* Primes ensure that when performing modulo operations with a large number 𝑚, the generated hash values are less likely to overlap, leading to fewer collisions.

**Example with b = 31 and modulo = 1000000007.**


