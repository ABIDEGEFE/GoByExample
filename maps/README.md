# How Go Maps Work

A **Map** is a built-in data structure in Go used to store data associatively using a key-value relationship. It is built on top of an array to provide high-speed data access using a hashing algorithm.

This document explains the internal mechanics of Go maps, from initialization to data retrieval and resizing.

---

## 1. Map Declaration & Initialization

When you declare a map using the `make` keyword:

```go
mp := make(map[string]int)

```

The Go runtime immediately creates an internal structure called `hmap`.

### The `hmap` Structure

The base header of a map looks like this:

```text
hmap = {
    count:   0
    B:       0
    buckets: ----> Bucket[0]
}

```

* **`count`**: The total number of elements currently stored in the map.
* **`B`**: A base-2 logarithm integer that determines the number of buckets ($2^B$). For example, if $B = 0$, $2^0 = 1$, meaning there is exactly 1 bucket.
* **`buckets`**: A pointer to a fixed-size array holding the actual bucket structures.

### The Bucket Layout

Each bucket can hold up to 8 key-value pairs and has the following internal memory format:

```text
{
    "tophash": ["hash1", "hash2", "........", "hash8"],
    "key":     ["key1",  "key2",  "............", "key8"],
    "value":   ["val1",  "val2",  "............", "val8"]
}

```

* **`tophash`**: An array containing the high-order bits of the hash for each key in the bucket. It makes searching incredibly fast using hardware-accelerated, CPU-native number comparisons.

---

## 2. Adding an Item to the Map

When you assign a value to a map key:

```go
mp["one"] = 1

```

The Go runtime performs the following steps:

1. **Hashing**: The key string is passed to a cryptographic hash function: `hash("one")` $\rightarrow$ `0xF12A....3A`.
2. **Splitting the Hash**: The resulting hash output is split into two parts:
* **Low-bits**: Determines which specific bucket the key belongs to (e.g., `3A` resolves to `Bucket[0]`).
* **High-bits (`tophash`)**: Used for fast internal matching inside the bucket (e.g., `0xF1`).



### Bucket State After Insertion

```text
Bucket[0]:
  tophash: [0xF1,  hash2, hash3, ......, hash8]
  key:     ["one", key2,  key3,  ......, key8]
  value:   [1,     val2,  val3,  ......, val8]

```

---

## 3. How Key Lookup Works

When you read a value from the map:

```go
value := mp["one"]

```

The runtime executes a highly optimized search algorithm:

1. **Hash Generation**: It calculates `hash("one")` $\rightarrow$ `0xF12A....3A`.
2. **Bucket Selection**: The low-bits (`3A`) identify that the key is located in `Bucket[0]`.
3. **Tophash Scanning**: The CPU scans the `tophash` array for the high-bits (`0xF1`).
4. **CPU Optimization**: It finds `0xF1` at index 0. Because this is a simple 1-byte comparison, the CPU can execute it in a single CPU cycle.
5. **Offset Resolution**: Using index 0 as the memory offset, the runtime verifies that `key[0]` matches `"one"` and directly returns `value[0]` (which is `1`).

---

## 4. Evacuation and Resizing (Growth)

Buckets have a fixed maximum capacity of 8 items.

* **The Trigger**: When a bucket fills up and crosses a specific load factor threshold, the Go runtime triggers a map growth phase.
* **The Process**: The runtime doubles the bucket size ($B = B + 1$) and establishes a temporary linkage between the old buckets and the new ones.
* **Incremental Evacuation**: To prevent application pauses, data is not moved all at once. The runtime smoothly evacuates items to the new buckets incrementally as write or delete operations occur.
