package utils

import (
	"crypto/rand"
	"fmt"
	"hash/fnv"
	"math/big"
	"time"
)

// HashWithDate luck算法是在此input int + 114514, max是100
func HashWithDate(input int, max int) (int64, error) {
	h := fnv.New32a()
	_, err := h.Write(helper.StringToBytes(fmt.Sprintf("%d%s", input, time.Now().Format("20060102"))))
	if err != nil {
		return 0, err
	}
	return int64(h.Sum32() % uint32(max+1)), nil
}
