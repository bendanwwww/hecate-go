package tools

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
)

// 长整型数的封装 线程安全
// 使用一个数组表示一个长整型数

const dataByte = 64

type BigBinaryTool interface {
	// BinarySize 二进制长度
	BinarySize() int
	// ChangeBitToTrue 将长整型数的某一位 置 1 (leftIndex: 从左到右计数 从 0 开始)
	ChangeBitToTrue(leftIndex int)
	// ChangeBitToFalse 将长整型数的某一位 置 0 (leftIndex: 从左到右计数 从 0 开始)
	ChangeBitToFalse(leftIndex int)
	// IsZero 判断长整型数是否为 0
	IsZero() bool
	// Equals 比较两个长整型数是否相等
	Equals(another *BigBinary) bool
	// EqualsWithLong 比较长整型数 和 一个 int64 的值是否相等
	EqualsWithLong(longData int64) bool
	// Or 长整型数 or 运算
	Or(another *BigBinary)
	// And 长整型数 and 运算
	And(another *BigBinary)
	// OrAndReturn 长整型数 or 运算 不改变原值
	OrAndReturn(another *BigBinary) *BigBinary
	// AndAndReturn 长整型数 and 运算 不改变原值
	AndAndReturn(another *BigBinary) *BigBinary
	// ShiftLeft 左移
	ShiftLeft(shiftNumber int)
	// ShiftRight 右移
	ShiftRight(shiftNumber int)
	// ToString 长整型数转字符串
	ToString() string
	// Copy 深拷贝
	Copy() *BigBinary
}

type BigBinary struct {
	/** 二进制长度 */
	binarySize int
	/** 数组长度 */
	dataSize int
	/** 二进制数组 */
	binaryData []int64
	/** 分段锁 数组的每一个元素持有一把锁 */
	binaryLock []sync.Mutex
	/** 全局锁 对数组多个元素进行操作的方法 使用全局锁 */
	globalLock sync.Mutex
}

func BuildNewBigBinary(binarySize int) *BigBinary {
	var dataSize, _ = getDataIndex(binarySize)
	return &BigBinary{
		binarySize: binarySize,
		dataSize:   dataSize + 1,
		binaryData: make([]int64, dataSize+1),
		binaryLock: make([]sync.Mutex, dataSize+1),
	}
}

func BuildNewBigBinaryFromLong(data int64) *BigBinary {
	return &BigBinary{
		binarySize: len(fmt.Sprintf("%b", data)),
		dataSize:   1,
		binaryData: []int64{data},
	}
}

func (bigBinary *BigBinary) BinarySize() int {
	return bigBinary.binarySize
}

func (bigBinary *BigBinary) ChangeBitToTrue(leftIndex int) {
	if leftIndex < 0 || leftIndex >= bigBinary.binarySize {
		return
	}
	index := bigBinary.binarySize - leftIndex - 1
	var arrayIndex, inArrayIndex = getDataIndex(index)
	// cas 一次
	oldData := bigBinary.binaryData[arrayIndex]
	newData := oldData | (1 << inArrayIndex)
	ok := atomic.CompareAndSwapInt64(&bigBinary.binaryData[arrayIndex], oldData, newData)
	if !ok {
		// 升级为悲观锁
		bigBinary.binaryLock[arrayIndex].Lock()
		bigBinary.binaryData[arrayIndex] |= 1 << inArrayIndex
		bigBinary.binaryLock[arrayIndex].Unlock()
	}
}

func (bigBinary *BigBinary) ChangeBitToFalse(leftIndex int) {
	if leftIndex < 0 || leftIndex >= bigBinary.binarySize {
		return
	}
	index := bigBinary.binarySize - leftIndex - 1
	var arrayIndex, inArrayIndex = getDataIndex(index)
	// cas 一次
	oldData := bigBinary.binaryData[arrayIndex]
	newData := oldData & (math.MaxInt64 ^ (1 << inArrayIndex))
	ok := atomic.CompareAndSwapInt64(&bigBinary.binaryData[arrayIndex], oldData, newData)
	if !ok {
		// 升级为悲观锁
		bigBinary.binaryLock[arrayIndex].Lock()
		bigBinary.binaryData[arrayIndex] &= math.MaxInt64 ^ (1 << inArrayIndex)
		bigBinary.binaryLock[arrayIndex].Unlock()
	}
}

func (bigBinary *BigBinary) IsZero() bool {
	for _, item := range bigBinary.binaryData {
		if item != 0 {
			return false
		}
	}
	return true
}

func (bigBinary *BigBinary) Equals(another *BigBinary) bool {
	if bigBinary.dataSize != another.dataSize {
		return false
	}
	for index := range bigBinary.binaryData {
		if bigBinary.binaryData[index] != another.binaryData[index] {
			return false
		}
	}
	return true
}

func (bigBinary *BigBinary) EqualsWithLong(longData int64) bool {
	if bigBinary.dataSize != 1 {
		return false
	}
	if bigBinary.binaryData[0] != longData {
		return false
	}
	return true
}

func (bigBinary *BigBinary) Or(another *BigBinary) {
	bigBinary.globalLock.Lock()
	var longBinary, shortBinary = getLongAndShortArray(bigBinary, another)
	var res = make([]int64, longBinary.dataSize)
	var onlyAppendSize = longBinary.dataSize - shortBinary.dataSize
	for i := len(longBinary.binaryData) - 1; i >= 0; i-- {
		if onlyAppendSize > 0 {
			res[i] = longBinary.binaryData[i]
			onlyAppendSize--
		} else {
			res[i] = longBinary.binaryData[i] | shortBinary.binaryData[i]
		}
	}
	bigBinary.binarySize = longBinary.binarySize
	bigBinary.dataSize = longBinary.dataSize
	bigBinary.binaryData = res
	bigBinary.globalLock.Unlock()
}

func (bigBinary *BigBinary) And(another *BigBinary) {
	bigBinary.globalLock.Lock()
	var longBinary, shortBinary = getLongAndShortArray(bigBinary, another)
	var res = make([]int64, longBinary.dataSize)
	var onlyAppendSize = longBinary.dataSize - shortBinary.dataSize
	for i := len(longBinary.binaryData) - 1; i >= 0; i-- {
		if onlyAppendSize > 0 {
			res[i] = 0
			onlyAppendSize--
		} else {
			res[i] = longBinary.binaryData[i] & shortBinary.binaryData[i]
		}
	}
	bigBinary.binarySize = longBinary.binarySize
	bigBinary.dataSize = longBinary.dataSize
	bigBinary.binaryData = res
	bigBinary.globalLock.Unlock()
}

func (bigBinary *BigBinary) OrAndReturn(another *BigBinary) *BigBinary {
	var longBinary, shortBinary = getLongAndShortArray(bigBinary, another)
	var res = make([]int64, longBinary.dataSize)
	var onlyAppendSize = longBinary.dataSize - shortBinary.dataSize
	for i := len(longBinary.binaryData) - 1; i >= 0; i-- {
		if onlyAppendSize > 0 {
			res[i] = longBinary.binaryData[i]
			onlyAppendSize--
		} else {
			res[i] = longBinary.binaryData[i] | shortBinary.binaryData[i]
		}
	}
	return &BigBinary{
		binarySize: longBinary.binarySize,
		dataSize:   longBinary.dataSize,
		binaryData: res,
		binaryLock: make([]sync.Mutex, longBinary.dataSize),
	}
}

func (bigBinary *BigBinary) AndAndReturn(another *BigBinary) *BigBinary {
	var longBinary, shortBinary = getLongAndShortArray(bigBinary, another)
	var res = make([]int64, longBinary.dataSize)
	var onlyAppendSize = longBinary.dataSize - shortBinary.dataSize
	for i := len(longBinary.binaryData) - 1; i >= 0; i-- {
		if onlyAppendSize > 0 {
			res[i] = 0
			onlyAppendSize--
		} else {
			res[i] = longBinary.binaryData[i] & shortBinary.binaryData[i]
		}
	}
	return &BigBinary{
		binarySize: longBinary.binarySize,
		dataSize:   longBinary.dataSize,
		binaryData: res,
		binaryLock: make([]sync.Mutex, longBinary.dataSize),
	}
}

func (bigBinary *BigBinary) ShiftLeft(shiftNumber int) {
	panic("not support")
}

func (bigBinary *BigBinary) ShiftRight(shiftNumber int) {
	bigBinary.globalLock.Lock()
	if shiftNumber >= bigBinary.binarySize {
		bigBinary.binarySize = 0
		bigBinary.binaryData = make([]int64, bigBinary.dataSize)
		return
	}

	moveIndex := shiftNumber / dataByte
	firstArrayMoveIndex := shiftNumber % dataByte
	bigBinary.binaryData[moveIndex] >>= firstArrayMoveIndex

	index := 0
	for i := moveIndex; i < bigBinary.dataSize; i++ {
		bigBinary.binaryData[index] = bigBinary.binaryData[moveIndex]
		index++
	}
	bigBinary.dataSize = index
	bigBinary.binarySize = bigBinary.binarySize - shiftNumber
	bigBinary.binaryData = bigBinary.binaryData[0:index]
	bigBinary.binaryLock = bigBinary.binaryLock[0:index]

	bigBinary.globalLock.Unlock()
}

func (bigBinary *BigBinary) ToString() string {
	var res = ""
	for i := bigBinary.dataSize - 1; i >= 0; i-- {
		var data = strconv.FormatInt(bigBinary.binaryData[i], 2)
		// 前面补 0
		if i < bigBinary.dataSize-1 && len(data) < dataByte {
			var zeroStr = ""
			for i := 0; i < dataByte-len(data); i++ {
				zeroStr += "0"
			}
			data = zeroStr + data
		}
		res += data
	}
	// 开头补 0
	var zeroStr = ""
	if len(res) < bigBinary.binarySize {
		for i := 0; i < bigBinary.binarySize-len(res); i++ {
			zeroStr += "0"
		}
	}
	return zeroStr + res
}

func (bigBinary *BigBinary) Copy() *BigBinary {
	bigBinary.globalLock.Lock()
	var copyDataArray = make([]int64, bigBinary.dataSize)
	copy(copyDataArray, bigBinary.binaryData)
	bigBinary.globalLock.Unlock()
	return &BigBinary{
		binarySize: bigBinary.binarySize,
		dataSize:   bigBinary.dataSize,
		binaryData: copyDataArray,
		binaryLock: make([]sync.Mutex, bigBinary.dataSize),
	}
}

func getLongAndShortArray(arrayOne *BigBinary, arrayTwo *BigBinary) (*BigBinary, *BigBinary) {
	if arrayOne.binarySize > arrayTwo.binarySize {
		return arrayOne, arrayTwo
	}
	return arrayTwo, arrayOne
}

func getDataIndex(binaryIndex int) (int, int) {
	if binaryIndex < dataByte {
		return 0, binaryIndex
	}
	if binaryIndex%dataByte == 0 {
		return binaryIndex / dataByte, 0
	}
	return binaryIndex / dataByte, binaryIndex % dataByte
}

var _ BigBinaryTool = (*BigBinary)(nil)
