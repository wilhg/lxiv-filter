package lxivFilter

import (
	"math"
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

// New will calculate the best m and k by n and p.
// n is the number of entries the data structure is expected to support.
// p is the false positive rate that is considered acceptable.
func NewWithEstimate(n uint64, p float64) *lxivFilter {
	m := uint64(math.Ceil(-1 * float64(n) * math.Log(p) / math.Pow(math.Log(2), 2)))
	m2 := uint64(1 << 6) // 64 is the beginning
	for m2 < m {
		m2 <<= 1
	}
	k := int(math.Ceil(math.Log(2) * float64(m) / float64(n)))
	return New(m2, k)
}

// NewDefault will new an LxivFilter who's size=1<<32, k = 4
// It will cost 1GB memory
func NewDefault() *lxivFilter { return New((1 << 32), 4) }

// Reset will clean the whole filter
func (lf *lxivFilter) Reset()      { lf.cells = make([]cell, lf.size) }
func (lf lxivFilter) Size() uint64 { return lf.size << 6 }
func (lf lxivFilter) K() int       { return lf.k }

// MayExist will check whether the data may exist (true), or definite not exist (false)
func (lf lxivFilter) MayExist(data []byte) bool {
	i := 0
	for ; i < lf.k; i += 2 {
		h1, h2 := h128(append(data, byte(i)))
		if !lf.isOn(h1) || !lf.isOn(h2) {
			return false
		}
	}
	if i > lf.k {
		h := h64(append(data, byte(i)))
		return lf.isOn(h)
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

func (lf lxivFilter) isOn(position uint64) bool {
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
