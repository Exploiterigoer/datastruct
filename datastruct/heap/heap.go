package heap

// Heap 堆结构体，本质是完全二叉树的数组表示形式
// 堆分为最大堆和最小堆
type Heap interface {
	Size() int
	IsEmpty() bool
	Insert(v int)
	Extract() int
	Elements(int) interface{}
	HeapSort() []int
}

// MaxHeap 最大堆，即每个父节点的值都大于其子节点值的完全二叉树
// 假设最大堆保存的是整数
// 关系通式：第i个节点的父节点下标 -> (i - 1) / 2
//			第i个节点的左节点下标 -> 2 * i + 1
// 			第i个节点的右节点下标 -> 2 * i + 2
// i为数组的下标
type MaxHeap struct {
	data  []int
	count int
}

// CreateMaxHeap 堆的初始化，通过存入单个元素
func CreateMaxHeap() *MaxHeap {
	xh := new(MaxHeap)
	xh.data = make([]int, 0)
	return xh
}

// CreateMaxHeap2 堆的初始化，通过存入一个数组
// O(n)
func CreateMaxHeap2(a []int) *MaxHeap {
	xh := new(MaxHeap)
	xh.data = make([]int, 0)
	length := len(a)
	if length == 0 {
		return xh
	}
	// 把数组元素复制到堆的数据域并设置到堆元素的个数
	xh.data = append(xh.data, a...)
	xh.count = length

	// (xh.count-1) / 2 表示第1个非叶节点
	// for循环是一个heapify过程：只要存在子节点值大于父节点值的就交换
	for i := (xh.count - 1) / 2; i >= 0; i-- {
		xh.shiftDown(i)
	}
	return xh
}

// Size 堆的大小
func (xh *MaxHeap) Size() int {
	return xh.count
}

// IsEmpty 堆的是否为空
func (xh *MaxHeap) IsEmpty() bool {
	return xh.Size() == 0
}

// Elements 最大堆元素值组成的数组
func (xh *MaxHeap) Elements(index int) interface{} {
	if index < 0 || index > xh.count-1 {
		return nil
	}
	if index > 0 {
		return xh.data[index]
	}

	return xh.data
}

// Insert 增加堆元素
// O(nlogn)
func (xh *MaxHeap) Insert(v int) {
	xh.data = append(xh.data, v) // 动态扩容
	xh.count++
	xh.shiftUp(xh.count - 1) // 维持最大堆的定义,xh.count - 1表示新增元素所在的下标
}

// 保持新加入节点后的完全二叉树仍然是最大堆的操作
// 根据最大堆的特征，增加的节点必须与其父节点作比较
func (xh *MaxHeap) shiftUp(c int) {
	// c表示当前节点所在的下标，可能是左节点也可能是右节点
	// (c-1)/2表示当前节点的父节点所在下标
	for (xh.data)[c] > (xh.data)[(c-1)/2] {
		(xh.data)[c], (xh.data)[(c-1)/2] = (xh.data)[(c-1)/2], (xh.data)[c]
		c = (c - 1) / 2
	}
}

// Extract 取出堆元素，只能取出堆的根元素，因为是最大堆，每次取出的值一定是堆中的最大值
func (xh *MaxHeap) Extract() int {
	if xh.Size() == 0 {
		return -1
	}
	root := xh.data[0]               // 注意，第二个元素才是堆的根元素
	xh.data[0] = xh.data[xh.count-1] // 交换最后一个元素与第1个元素
	xh.count--                       // 元素个数减1
	xh.data = xh.data[:xh.count]     // 舍弃一个存储单元
	xh.shiftDown(0)                  // 少了一个元素需要维持堆的定义
	return root
}

func (xh *MaxHeap) shiftDown(c int) {
	// 2*c+2 表示右节点所在下标,有右节点必有左节点
	for 2*c+1 < xh.count {
		// j表示当前节点的左节点的下标，j+1表示前节点的右节点的下标
		j := 2*c + 1

		// 有右节点且右节点的值比左节点的值大
		if j+1 < xh.count && xh.data[j+1] > xh.data[j] {
			j++ // 这里j自增后恰好是右节点的下标
		}

		// 父节点比左、右节点的值还要大，无需任何交换
		if xh.data[c] >= xh.data[j] {
			break
		}
		// 当前节点与当前节点的左右节点中值较大的节点交换
		xh.data[c], xh.data[j] = xh.data[j], xh.data[c]
		c = j
	}
}

// HeapSort 堆排序,借助额外空间O(n)把每次取出的元素值保存起来
func (xh *MaxHeap) HeapSort() []int {
	element := make([]int, 0)
	for !xh.IsEmpty() {
		element = append(element, xh.Extract())
	}
	return element
}

/*+++++++++++++++++++++++分割线+++++++++++++++++++++++*/

// MinHeap 最小堆，即每个父节点的值都小于其子节点值的完全二叉树
type MinHeap struct {
	data  []int
	count int
}
