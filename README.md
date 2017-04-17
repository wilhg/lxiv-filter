# lxiv-filter
> LXIV is the number 64 in Roman number.

A Bloom filter is a representation of a set of n items, where the main requirement is to make membership queries; i.e., whether an item may be a member of a set.

This bloom filter implementation is backed by uint64 array.

And the hashing functions used is [murmur3](github.com/spaolacci/murmur3), a fast and good hashing function.

***WARNING***: Before you using any implementation of bloom filter, please have a view of this [article](http://pages.cs.wisc.edu/~cao/papers/summary-cache/node8.html), to know how to config your parameters.

## Installation

```bash
go get -u github.com/cuebyte/lxiv-filter
```

## Usage
```go
import "github.com/cuebyte/lxiv-filter"

lf := lxivFilter.NewDefault() // == lf.ManualNew(1<<32, 5)

// The size has to be a power of 2, and greater than 64.
lf.Size()                           // Return 1 << 32
lf.K()                              // Return 5

lf.MayExist([]byte("Hello World!")) // Return false

lf.Add([]byte("Hello World!"))

lf.MayExist([]byte("Hello World!")) // Return true

lf.Reset()                          // Clean the bit-map
```

However, before calling `Add()`, we don't need to check whether the data is existed.
