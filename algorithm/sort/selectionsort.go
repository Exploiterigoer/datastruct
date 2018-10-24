package sortalgorithm

// SelectionSort 选择排序结构体
type SelectionSort struct{}

// SelectSort 选择排序算法
// 时间复杂度：O(n^2)
func (ss *SelectionSort) SelectSort(o *[]int, args ...interface{}) interface{} {
	if o == nil {
		return nil
	}
	length := len(*o)
	for i := 0; i < length; i++ { // 找到要比较的元素
		minIndex := i                     // 当前最小元素的下标
		for j := i + 1; j < length; j++ { // 一路比较到数组最后，得出最小的元素
			if (*o)[j] < (*o)[minIndex] {
				minIndex = j // 重新设置前最小元素的下标
			}
		}
		(*o)[i], (*o)[minIndex] = (*o)[minIndex], (*o)[i] // 交换最小元素，即小的值往前移动
	}
	return 0
}
