package assistant

import (
	"math/rand"
	"time"
)

// CalcTime 计算算法所用的时间
func CalcTime(f func(*[]int, ...interface{}) interface{}, array *[]int, args ...interface{}) (interface{}, int64) {
	startTime := time.Now().UnixNano()
	result := f(array, args)
	endTime := time.Now().UnixNano()
	return result, endTime - startTime
}

// ArrayGenerator 随机生成number在[0,flag]范围内的数据
func ArrayGenerator(number int, flag int) []int {
	array := make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < number; i++ {
		array = append(array, r.Intn(flag+1))
	}
	return array
}

// IsSorted 检查数组元素是否被排序
func IsSorted(array *[]int) bool {
	length := len(*array)
	for i := 0; i < length-1; i++ {
		if (*array)[i] < (*array)[i+1] {
			return false
		}
	}
	return true
}

// RandIndex 返回指定范围内的数组下标
func RandIndex(borderA, borderB int) int {
	if borderB < 0 || borderA < 0 || borderB < borderA {
		return 0
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(borderB-borderA+1) + borderA
	return index
}
