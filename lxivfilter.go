package lxivFilter

import (
	"github.com/spaolacci/murmur3"
)

type lxivFilter struct {
	cells []cell

	size uint32
	k    int
}

// New will new a LxivFilter
// size here is to define the size of []uint64 array
// k means the number of times to call hash functions, here is using murmur3
// they means you will get a 64*size bit-map, with k times hash call
func New(size uint32, k int) *lxivFilter {
	if size&(size-1) != 0 {
		panic("Please set the size as a power of 2")
	}
	cells := make([]cell, size)
	return &lxivFilter{cells, size, k}
}

// NewDefault will new an LxivFilter who's size=1<<32 - 1, k = 5
// It will cost 1GB memory
func NewDefault() *lxivFilter { return New((1<<32 - 1), 5) }

// Reset will clean the whole filter
func (lf *lxivFilter) Reset()      { lf.cells = make([]cell, lf.size) }
func (lf lxivFilter) Size() uint32 { return lf.size }
func (lf lxivFilter) K() int       { return lf.k }

// MayExist will check whether the data may exist (true), or definite not exist (false)
func (lf lxivFilter) MayExist(data []byte) bool {
	for i := 0; i < lf.k; i++ {
		mapIdx, cellIdx := lf.calcPosition(append(data, byte(i)))
		if lf.cells[mapIdx].at(cellIdx) == false {
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

func (lf lxivFilter) calcPosition(data []byte) (mapIdx uint32, cellIdx uint8) {
	hash := murmur3.New64()
	hash.Write(data)
	hashCode := hash.Sum64()

	mapIdx = uint32((hashCode >> 6) & uint64(lf.size-1)) // == (hashCode >> 6) % lf.size
	cellIdx = uint8(hashCode & (1<<7 - 1))               // ==  hashCode % 64
	return
}
