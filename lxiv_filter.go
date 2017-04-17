package lxivFilter

import (
	"github.com/spaolacci/murmur3"
)

type lxivFilter struct {
	cells []cell

	size uint64
	k    int
}

// New will new a LxivFilter
// size here is to define the size bit-map, it has to be a power of 2 and larger than 64
// k means the number of times to call hash functions, here is using murmur3
func New(size uint64, k int) *lxivFilter {
	if size&(size-1) != 0 {
		panic("Please set the size as a power of 2")
	}
	if size < 64 {
		panic("size should greater than 64")
	}
	if k > 32 {
		panic("k shouldn't greater than 32")
	}
	cells := make([]cell, size>>6)
	return &lxivFilter{cells, size >> 6, k}
}

// NewDefault will new an LxivFilter who's size=1<<32, k = 5
// It will cost 1GB memory
func NewDefault() *lxivFilter { return New((1 << 32), 5) }

// Reset will clean the whole filter
func (lf *lxivFilter) Reset()      { lf.cells = make([]cell, lf.size) }
func (lf lxivFilter) Size() uint64 { return lf.size << 6 }
func (lf lxivFilter) K() int       { return lf.k }

// MayExist will check whether the data may exist (true), or definite not exist (false)
func (lf lxivFilter) MayExist(data []byte) bool {
	for i := 0; i < lf.k; i++ {
		mapIdx, cellIdx := lf.calcPosition(append(data, byte(i)))
		if !lf.cells[mapIdx].at(cellIdx) {
			return false
		}
	}
	return true
}

// Add will set records to filter
func (lf *lxivFilter) Add(data []byte) {
	for i := 0; i < lf.k; i++ {
		mapIdx, cellIdx := lf.calcPosition(append(data, byte(i)))
		lf.cells[mapIdx] = lf.cells[mapIdx].turnOn(cellIdx)
	}
}

func (lf lxivFilter) calcPosition(data []byte) (uint64, uint8) {
	hash := murmur3.New64()
	hash.Write(data)
	hashCode := hash.Sum64()

	mapIdx := uint64((hashCode >> 5) & uint64(lf.size-1)) // == (hashCode >> 5) % lf.size
	cellIdx := uint8(hashCode & (1<<6 - 1))               // ==  hashCode % 64
	return mapIdx, cellIdx
}
