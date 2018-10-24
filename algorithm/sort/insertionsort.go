package sortalgorithm

// InsertionSort 插入排序结构体
type InsertionSort struct{}

// InsertSort 时间复杂度:O(n^2)
// 算法思想:遍历数组，比较当前元素(数组的第一个元素相当于排好序的，不用管)与前一个元素的值，
// 大于或小于，则将当前元素前移或前一个元素后移
// 对于一个近乎有序的数组，采用插入排序算法，效率非常高，时间复杂度接近O(n)
func (is *InsertionSort) InsertSort(o *[]int, args ...interface{}) interface{} {
	length := len(*o)
	for i := 1; i < length; i++ { // 比较的范围
		// 原始版
		/*for j := i; j > 0 && (*o)[j] > (*o)[j-1]; j-- {
			(*o)[j], (*o)[j-1] = (*o)[j-1], (*o)[j]
		}*/

		// 优化版 1
		var j int
		e := (*o)[i]                             // 当前的元素
		for j = i; j > 0 && e < (*o)[j-1]; j-- { // 当前元素与前一个元素比较并交换，直到找出适合的位置
			(*o)[j] = (*o)[j-1] // 只交换1次
		}
		(*o)[j] = e // 把当前的元素放到合适的位置
	}
	return 0
}

// InsertSort2 版本2
func (is *InsertionSort) InsertSort2(o *[]int, args ...interface{}) interface{} {
	var left, right int
	switch args[0].(type) {
	case []interface{}:
		param := args[0].([]interface{})
		left, right = param[0].(int), param[1].(int)
	case int:
		left, right = args[0].(int), args[1].(int)
	}

	for i := left + 1; i <= right; i++ {
		var j int
		e := (*o)[i]
		for j = i; j > left && e < (*o)[j-1]; j-- {
			(*o)[j] = (*o)[j-1]
		}
		(*o)[j] = e
	}
	return 0
}
