package lxivFilter

import (
	"math/rand"
	"time"

	"github.com/spaolacci/murmur3"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func genRandByteArray(strlen int) []byte {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return result
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
