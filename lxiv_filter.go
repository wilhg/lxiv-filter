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
	if k > 32 || k <= 1 {
		panic("k should be: 1 < k <= 32")
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
	i := 0
	for ; i < lf.k; i += 2 {
		h1, h2 := h128(append(data, byte(i)))
		if lf.read(h1) == false {
			return false
		}
		if lf.read(h2) == false {
			return false
		}
	}
	if i > lf.k {
		h := h64(append(data, byte(i)))
		return lf.read(h)
	}
	return true
}

// Add will set records to filter
func (lf *lxivFilter) Add(data []byte) {
	i := 0
	for ; i < lf.k; i += 2 {
		h1, h2 := h128(append(data, byte(i)))
		lf.switchOn(h1)
		lf.switchOn(h2)
	}
	if i > lf.k {
		h := h64(append(data, byte(i)))
		lf.switchOn(h)
	}
}

func (lf lxivFilter) read(position uint64) bool {
	mapIdx, cellIdx := lf.calcPosition(position)
	return lf.cells[mapIdx].at(cellIdx)
}

func (lf *lxivFilter) switchOn(position uint64) {
	mapIdx, cellIdx := lf.calcPosition(position)
	lf.cells[mapIdx] = lf.cells[mapIdx].switchOn(cellIdx)
}

func (lf lxivFilter) calcPosition(hashCode uint64) (uint64, uint8) {
	mapIdx := uint64((hashCode >> 5) & uint64(lf.size-1)) // == (hashCode >> 5) % lf.size
	cellIdx := uint8(hashCode & (1<<6 - 1))               // ==  hashCode % 64
	return mapIdx, cellIdx
}

func h64(data []byte) uint64 {
	hash := murmur3.New64()
	hash.Write(data)
	return hash.Sum64()
}

func h128(data []byte) (uint64, uint64) {
	hash := murmur3.New128()
	hash.Write(data)
	return hash.Sum128()
}
