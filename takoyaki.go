package takoyaki

import (
	"github.com/spaolacci/murmur3"
)

type tako struct {
	podMap []octopod

	size uint32
	keys int
}

// NewDefault will create an array of [MaxInt32]uint64, with 3 keys
// It will cost 1GB memory
func NewDefault() *tako { return New((1<<32 - 1), 4) }

// New i
func New(size uint32, keys int) *tako {
	if size&(size-1) != 0 {
		panic("Please set the size as a power of 2")
	}
	podMap := make([]octopod, size)
	return &tako{podMap, size, keys}
}

func (t *tako) Reset()      { t.podMap = make([]octopod, t.size) }
func (t tako) Size() uint32 { return t.size }
func (t tako) Keys() int    { return t.keys }

func (t tako) MayExist(data []byte) bool {
	for i := 0; i < t.keys; i++ {
		mapIdx, podIdx := t.calcPosition(append(data, byte(i)))
		if t.podMap[mapIdx].at(podIdx) == false {
			return false
		}
	}
	return true
}

func (t *tako) Add(data []byte) {
	for i := 0; i < t.keys; i++ {
		mapIdx, podIdx := t.calcPosition(append(data, byte(i)))
		t.podMap[mapIdx] = t.podMap[mapIdx].turnOn(podIdx)
	}
}

func (t tako) calcPosition(data []byte) (mapIdx uint32, podIdx uint8) {
	hash := murmur3.New64()
	hash.Write(data)
	hashCode := hash.Sum64()

	mapIdx = uint32((hashCode >> 6) & uint64(t.size-1)) // == (hashCode >> 6) % t.size
	podIdx = uint8(hashCode & (1<<7 - 1))               // == hashCode % 64
	return
}
