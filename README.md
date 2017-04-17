# lxiv-filter
> LXIV is the number 64 in Roman number.

A Bloom filter is a representation of a set of n items, where the main requirement is to make membership queries; i.e., whether an item may be a member of a set.

This bloom filter implementation is backed by uint64 array.

And the hashing functions used is [murmurhash](github.com/spaolacci/murmur3), a non-cryptographic hashing function.

## Installation

```bash
go get -u github.com/cuebyte/lxiv-filter
```

## Usage
```go
import "github.com/cuebyte/lxiv-filter"

lf := lxivFilter.NewDefault()

lf.Size() == 1 << 24 // True
lf.K() == 5 // True

lf.MayExist([]byte("Hello World!")) == false // True

lf.Add([]byte("Hello World!"))

lf.MayExist([]byte("Hello World!")) == true // True

lf.Reset() // Clean the bit-map
```

However, before calling `Add()` method, we don't need to check whether the data is existed.
