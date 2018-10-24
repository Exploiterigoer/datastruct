package sortalgorithm

import (
	"data_struct_algorithm/assistant"
)

// QuickSort 快速排序结构体
type QuickSort struct{}

// QuicklySort 快速排序算法
func (q *QuickSort) QuicklySort(o *[]int, args ...interface{}) interface{} {
	l, r := 0, len(*o)-1
	quicklySort(o, l, r)
	return 0
}

// quicklySort 内部分区处理
func quicklySort(o *[]int, l, r int) {
	/*if l >= r {
		return
	}*/

	// 优化处理
	if r-l <= 60 {
		new(InsertionSort).InsertSort2(o, l, r)
		return
	}

	p := partition(o, l, r) // 数组分区
	quicklySort(o, l, p-1)  // 左半数组递归分区
	quicklySort(o, p+1, r)  // 右半数组递归分区
}

// QuicklySort2 快速排序算法
func (q *QuickSort) QuicklySort2(o *[]int, args ...interface{}) interface{} {
	l, r := 0, len(*o)-1
	quicklySort2(o, l, r)
	return 0
}

func quicklySort2(o *[]int, l, r int) {
	/*if l >= r {
		return
	}*/

	// 优化处理
	if r-l <= 60 {
		new(InsertionSort).InsertSort2(o, l, r)
		return
	}

	p := partition2(o, l, r) // 数组分区
	quicklySort2(o, l, p-1)  // 左半数组递归分区
	quicklySort2(o, p+1, r)  // 右半数组递归分区
}

// QuicklySort3 快速排序算法
func (q *QuickSort) QuicklySort3(o *[]int, args ...interface{}) interface{} {
	l, r := 0, len(*o)-1
	quicklySort3(o, l, r)
	return 0
}

func quicklySort3(o *[]int, l, r int) {
	// 优化处理
	if r-l <= 60 {
		new(InsertionSort).InsertSort2(o, l, r)
		return
	}
	ranIndex := assistant.RandIndex(l, r)
	(*o)[ranIndex], (*o)[r] = (*o)[r], (*o)[ranIndex]

	lt, gt := partition3(o, l, r) // 数组分区
	quicklySort3(o, l, lt)        // 左半数组递归分区
	quicklySort3(o, gt, r)        // 右半数组递归分区
}

// 返回下标j，这个j使得arr[l......j-1]<arr[j]<arr[j+1......r]
// 随机标定 + 单路扫描
func partition(o *[]int, l, r int) int {
	//e := (*o)[l] // 将数组分割成3部分的元素值，假定为第一个元素

	// 优化处理，避免了数组元素近乎有序时时间复杂度退化到O(n^2)
	e := (*o)[assistant.RandIndex(l, r)]
	j := l

	// arr[l+1......j]< e <arr[j+1......i)
	// i位置是当前考察的元素的索引
	for i := l + 1; i <= r; i++ {
		// 对小于设置的标准元素的其他元素进行交换
		if (*o)[i] < e {
			(*o)[j+1], (*o)[i] = (*o)[i], (*o)[j+1]
			j++
		}
	}

	// 把原数组第一个元素与第j个元素进行交换
	(*o)[l], (*o)[j] = (*o)[j], (*o)[l]

	return j
}

// 避免了大量重复元素导致时间复杂度退化为O(n^2)
// 随机标定 + 双路扫描(考察元素等于标定值的也要交换，性能损失)
func partition2(o *[]int, l, r int) int {
	idx := assistant.RandIndex(l, r)
	(*o)[l], (*o)[idx] = (*o)[idx], (*o)[l]
	e := (*o)[l] // 此时的e是随机选择的

	i, j := l+1, r
	for {
		for i <= r && (*o)[i] < e {
			i++
		}
		for j >= l+1 && (*o)[j] > e {
			j--
		}

		if i > j {
			break
		}

		(*o)[i], (*o)[j] = (*o)[j], (*o)[i]
		i++
		j--
	}

	(*o)[l], (*o)[j] = (*o)[j], (*o)[l]

	return j
}

// 随机标定 + 三路扫描
func partition3(o *[]int, l, r int) (int, int) {

	// 写法1
	/*lt, gt := l-1, r
	for l < gt {
		if (*o)[l] < (*o)[r] {
			lt++
			(*o)[lt], (*o)[l] = (*o)[l], (*o)[lt]
			l++
			continue
		}

		if (*o)[l] > (*o)[r] {
			gt--
			(*o)[gt], (*o)[l] = (*o)[l], (*o)[gt]
			continue
		}

		l++
	}

	(*o)[gt], (*o)[r] = (*o)[r], (*o)[gt]
	return lt, gt + 1*/

	// 写法2
	idx := assistant.RandIndex(l, r)
	(*o)[l], (*o)[idx] = (*o)[idx], (*o)[l]
	e := (*o)[l] // 此时的e是随机选择的
	i, j, k := l, r+1, l+1
	// arr[l+1...i] < e     arr[i+1...k) = e     arr[j...r]> e
	for k < j {
		if (*o)[k] < e {
			(*o)[k], (*o)[i+1] = (*o)[i+1], (*o)[k]
			i++
			k++
		} else if (*o)[k] > e {
			(*o)[k], (*o)[j-1] = (*o)[j-1], (*o)[k]
			j--
		} else {
			k++
		}
	}
	(*o)[l], (*o)[i] = (*o)[i], (*o)[l]
	i-- // 此时(*o)[i]是排好序的，i位置的元素无需处理，[l,i-1]区间是下一次被处理的，因而要自减操作
	return i, j
}
