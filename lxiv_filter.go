package lxivFilter

import (
	"math"

	"github.com/spaolacci/murmur3"
)

type lxivFilter struct {
	cells []cell

	size    uint64
	bitSize uint64
	k       uint8
}

// New will new a LxivFilter
// size here is to define the size bit-map, it has to be a power of 2 and larger than 64
// k means the number of times to call hash functions, here is using murmur3
func New(size uint64, k uint8) *lxivFilter {
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
	return &lxivFilter{cells, size >> 6, size, k}

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
	k := uint8(math.Ceil(math.Log(2) * float64(m) / float64(n)))
	return New(m2, k)
}

// NewDefault will new an LxivFilter who's size=1<<32, k = 4
// It will cost 1GB memory
func NewDefault() *lxivFilter { return New(1<<32, 4) }

// Reset will clean the whole filter
func (lf *lxivFilter) Reset()      { lf.cells = make([]cell, lf.size) }
func (lf lxivFilter) Size() uint64 { return lf.bitSize }
func (lf lxivFilter) K() uint8     { return lf.k }

// MayExist will check whether the data may exist (true), or definite not exist (false)
func (lf lxivFilter) MayExist(data []byte) bool {
	baseHash := hash256(data)
	for i := uint8(0); i < lf.k; i++ {
		if !lf.isOn(baseHash, i) {
			return false
		}
	}
	return true
}

// Add will set records to filter
func (lf *lxivFilter) Add(data []byte) {
	baseHash := hash256(data)
	for i := uint8(0); i < lf.k; i++ {
		lf.switchOn(baseHash, i)
	}
}

func (lf lxivFilter) isOn(base [4]uint64, i uint8) bool {
	mapIdx, cellIdx := lf.calcPosition(base, i)
	return lf.cells[mapIdx].at(cellIdx)
}

func (lf *lxivFilter) switchOn(base [4]uint64, i uint8) {
	mapIdx, cellIdx := lf.calcPosition(base, i)
	lf.cells[mapIdx] = lf.cells[mapIdx].switchOn(cellIdx)
}

func hash256(data []byte) [4]uint64 {
	h := murmur3.New128()
	h.Write(data)
	x0, x1 := h.Sum128()
	h.Write([]byte{42})
	x2, x3 := h.Sum128()
	return [4]uint64{x0, x1, x2, x3}
}

func (lf lxivFilter) calcPosition(h [4]uint64, x uint8) (uint64, uint8) {
	ux := uint64(x)
	xn1 := x & 1
	hashCode := (h[xn1] + ux*h[(((x+xn1)&3)>>1)+2]) & (lf.bitSize - 1)
	mapIdx := (hashCode >> 6) & (lf.size - 1) // == (hashCode >> 6) % lf.size
	cellIdx := uint8(hashCode & 0x3f)         // ==  hashCode % 64
	return mapIdx, cellIdx
}
