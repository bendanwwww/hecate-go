package tools

import "strconv"

func Float64String(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func String2Int64(s string) (i int64, err error) {
	return strconv.ParseInt(s, 10, 64)
}

func String2Int(s string) (i int, err error) {
	v, err := strconv.ParseInt(s, 10, 64)
	return int(v), err
}

func String2Int32(s string) (i int32, err error) {
	v, err := strconv.ParseInt(s, 10, 64)
	return int32(v), err
}

func Int64String(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Int2String(i int) string {
	return Int64String(int64(i))
}

func String2Float64(s string) (i float64, err error) {
	return strconv.ParseFloat(s, 64)
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func MinI64(x, y int64) int64 {
	if x < y {
		return x
	} else {
		return y
	}
}

// Accumulation n!
func Accumulation(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return n * (n - 1) / 2
}
