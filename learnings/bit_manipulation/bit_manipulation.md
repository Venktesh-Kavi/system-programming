## Notes on Bit Manipulaion

Left Shift and Right Shift Operator

Basics 
- 8 bits = 1 byte
- 1024 byte = 1kb
- 1000kb = 1MB

* Till KB computers can represent only in binary (base 2 0,1's), the closest to 1000 is 2^10 = 1024
* From MB and GB we follow the manufaturer representation. (1

<<
* Moves bits to the left
* Each shift multiples the number by 2
* Tail/Right side is filled with zeros

```
// examples
1 << 1 = 2 => (0001 to 0010)
1 << 2 = 4 => (0010 to 0100)
1 << 3 = 2^3 (8) => (0100 to 1000)

(1 << 30) - 1
- 1 << 30 = 1073741824
- subtracting 1 gives = 1073741823
- Binary = 01111111111111111111111111111111 (represented in 32 bits)
```

Notes: Why do we start with 4 numbers
* 4 numbers represent 4 bits
* 8 bits represent 8 zeros/ones
* 16 bits represents 16 zeros/ones
* For a 32 or 64 bit architecture CPU, it 32/64 bits of zeros/ones

We start with 4 numbers to represent small numbers and grow as needed.

decimal 4 => 0100
decimal 8 => 1000

In go binary literal is defined as:
bl := 0b0100 // represent 4, 0b represents its binary, 0x for hexadecimal (used for compacting), 0o for octal

### Bit Mask Operators

* AND(&) (isolate bits)
* OR(|) (set bits)
* XOR(^) (toggle bits)
* NOT(~) (invert bits)


#### & (Isolate Bits)

value := uint8(0b00001101) // 13
// isolate the last 2 bits
mask := uint8(0b00000011) // 3  

fmt.Println(value&mask) // output 3, go automatically converts the binary rep to uint8

// Another example: Checking a specific bit (bit 3)
const bitToCheck uint8 = 1 << 3 // isolate bit 3 (4th bit from the right)
isBitSet := num & bitToCheck


