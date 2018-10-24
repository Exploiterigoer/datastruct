package sortalgorithm

// Heap 堆的接口
type Heap interface {
	Create([]int)
	Size() int
	IsEmpty() bool
	Extract() int
	HeapSort() []int
	Element(int) interface{}
	heapify(int)
}

// MaxHeap 大根堆结构体
type MaxHeap struct {
	data  []int
	count int
}

// Size 大根堆元素个数
func (xh *MaxHeap) Size() int {
	return xh.count
}

// IsEmpty 大根堆是否为空
func (xh *MaxHeap) IsEmpty() bool {
	return xh.Size() == 0
}

// Create 从数组建立大根堆
func (xh *MaxHeap) Create(arr []int) {
	length := len(arr)
	if arr == nil || length == 0 {
		return
	}

	for i := 0; i < length; i++ {
		xh.insert(arr, i)
		xh.count++
	}
}

// insert 从单个元素建立大根堆
func (xh *MaxHeap) insert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
	xh.data = arr
}

// Element 返回指定下标位置的大根堆元素
func (xh *MaxHeap) Element(i int) interface{} {
	if i < 0 {
		return xh.data[:xh.count]
	}

	if i > xh.count {
		return xh.data[xh.count-1]
	}

	return xh.data[i]
}

// Extract 取出大根堆元素
func (xh *MaxHeap) Extract() int {
	if xh.Size() == 0 {
		return -1
	}
	root := xh.data[0]               // 被删除的根元素
	xh.data[0] = xh.data[xh.count-1] // 交换最后一个元素与第一个元素
	xh.count--                       // 元素个数减1
	xh.data = xh.data[:xh.count]     // 舍弃一个存储单元
	xh.heapify(0)                    // 要维持堆的结构
	return root
}

// Heapify 维持大根堆的过程：值变小就往下沉的过程
func (xh *MaxHeap) heapify(index int) {
	left := index*2 + 1
	// lrmax表示左右节点中值最大节点所在下标
	var lrmax int
	for left < xh.count {
		// 假定左节点值大，取得左节点所在的下标
		lrmax = left

		// 存在右节点且右节点值大于左节点值则找出右节点所在下标
		if left+1 < xh.count && xh.data[left+1] > xh.data[left] {
			lrmax = left + 1
		}

		// xh.data[index]是xh.data[lrmax]的父节点，xh.data[lrmax]可能是左或右节点
		// 已经找到左右节点中值大的节点，与父节点的值比较
		if xh.data[lrmax] < xh.data[index] {
			lrmax = index
		}

		// 左右节点中的最大值就是父节点的值，此时可以退出循环，无需再比较
		if lrmax == index {
			break
		}

		// 将父节点的值与交换左右节点中值大的节点值交换
		xh.data[lrmax], xh.data[index] = xh.data[index], xh.data[lrmax]
		index = lrmax // 把左右节点其中之一作为新的的父节点再次heapify过程
		left = index*2 + 1
	}
}

// HeapSort 大根堆排序，从小到大排序
// 原理：根节点与最后一个子节点交换，每次交换后缩减节点个数，并重新调整数组为
// 大根堆，直至节点个数为0，这样每次交换后的根节点组成的数组就是从小到大排好序的
// 并且得到是数组还是小根堆
// 时间复杂度O(n*logn)，空间复杂度O(1)
func (xh *MaxHeap) HeapSort() []int {
	xh.data[0], xh.data[xh.count-1] = xh.data[xh.count-1], xh.data[0]
	xh.count--
	for xh.count > 0 {
		xh.heapify(0)
		xh.data[0], xh.data[xh.count-1] = xh.data[xh.count-1], xh.data[0]
		xh.count--
	}
	return xh.data
}

// HeapSort2 堆排序算法
func HeapSort2(o *[]int) {
	// heapify 将数组变成大根堆
	length := len(*o)
	for i := (length - 1) / 2; i >= 0; i-- {
		shiftDown(o, length, i)
	}

	// sort 将数组按照升序排好序
	for i := length - 1; i > 0; i-- {
		(*o)[0], (*o)[i] = (*o)[i], (*o)[0]
		shiftDown(o, i, 0)
	}
}

func shiftDown(o *[]int, n, i int) {
	for 2*i+1 < n {
		j := 2*i + 1
		if j+1 < n && (*o)[j+1] > (*o)[j] {
			j++
		}
		if (*o)[i] > (*o)[j] {
			break
		}
		(*o)[i], (*o)[j] = (*o)[j], (*o)[i]
		i = j
	}
}
