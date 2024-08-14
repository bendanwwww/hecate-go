package tools

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+{}:?><~`1234567890-=[];/.,"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandFloat64() float64 {
	return rand.Float64()
}

func RandFloat32() float32 {
	return rand.Float32()
}

func RandPerm(n int) []int {
	return rand.Perm(n)
}

func RandIntN(n int) int {
	return rand.Intn(n)
}

// RandRangeIntN from [start, end) random a number
func RandRangeIntN(start, end int) int {
	return start + rand.Intn(end-start)
}

// RandRangeIntNs from [start, end) random n numbers
func RandRangeIntNs(start, end, n int) []int {
	length := end - start
	randArray := rand.Perm(length)
	resArray := make([]int, n)
	for i := 0; i < n; i++ {
		resArray[i] = randArray[i] + start
	}
	return resArray
}

func RandInt63n(n int64) int64 {
	return rand.Int63n(n)
}

func ShuffleFloat64Slice(items []float64) []float64 {
	if len(items) <= 1 {
		return items
	}

	ii := RandPerm(len(items))

	randItems := make([]float64, len(items))

	for i, randIdx := range ii {
		randItems[i] = items[randIdx]
	}
	return randItems
}

func ShuffleStringSlice(items []string) []string {
	if len(items) <= 1 {
		return items
	}

	ii := RandPerm(len(items))

	randItems := make([]string, len(items))

	for i, randIdx := range ii {
		randItems[i] = items[randIdx]
	}
	return randItems
}

func ShuffleInt64Slice(items []int64) []int64 {
	if len(items) <= 1 {
		return items
	}

	ii := RandPerm(len(items))

	randItems := make([]int64, len(items))

	for i, randIdx := range ii {
		randItems[i] = items[randIdx]
	}
	return randItems
}

func RandString(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
