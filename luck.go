package utils

import (
	"crypto/rand"
	"fmt"
	"hash/fnv"
	"math/big"
	"time"
)

// luck算法是在此input int + 114514, max是100
func HashWithDate(input int, max int) int {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%d-%s", input, time.Now().Format("2006-01-02"))))
	hashValue := h.Sum32()
	return int(hashValue % uint32(max+1)) // 将哈希值映射到 0 到 max 范围内
}
