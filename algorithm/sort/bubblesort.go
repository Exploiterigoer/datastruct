package sortalgorithm

// Bubble 冒泡排序结构体
type Bubble struct{}

// BubbleSort 冒泡排序
// 时间复杂度O(n^2)
func (b *Bubble) BubbleSort(o *[]int, args ...interface{}) interface{} {
	if o == nil {
		return nil
	}
	length := len(*o)
	for i := length - 1; i > 0; i-- { // 确定比较的范围
		for j := 0; j < i; j++ { // 在范围内依次比较前后两个数
			if (*o)[i] < (*o)[i-1] { // 找出较大的元素并向后移动
				(*o)[i], (*o)[i-1] = (*o)[i-1], (*o)[i]
			}
		}
	}
	return 0
}
