package sortalgorithm

// MergerSort 归并排序结构体
type MergerSort struct{}

// o 被处理的数组区间
// l 被处理的区间的开始索引
// r 被处理的区间的结束索引
// t 外部传入的临时数组，避免在merger函数内部频繁创建空间带来性能损耗，
// t 是有足够空间保存数据的，因为t的长度等于被排序的数组的长度
func merger(o *[]int, l, m, r int, t *[]int) {
	// 复制一份当前需要被处理的数据
	for i := l; i <= r; i++ {
		// 临时数组下标从0开始，每一次归并源数组下标从l开始
		(*t)[i-l] = (*o)[i]
	}

	// j左部有序数组的开始索引，k右部有序数组的开始索引
	j, k := l, m+1

	// 修改被排序的数组区间的元素，排序区间的起始和结束索引是l、r
	for n := l; n <= r; n++ {
		if j > m {
			// 左边的下标已经超出其最大值，说明左边归并完但右边还没有
			(*o)[n] = (*t)[k-l]
			k++
		} else if k > r {
			// 右边的下标已经超出其最大值，说明右边归并完但左边还没有
			(*o)[n] = (*t)[j-l]
			j++
		} else if (*t)[j-l] < (*t)[k-l] {
			// 比较，找出左部合适的数据填到合适的位置
			(*o)[n] = (*t)[j-l] // 注意升序还是降序
			j++
		} else {
			// 比较，找出右部合适的数据填到合适的位置
			(*o)[n] = (*t)[k-l]
			k++
		}
	}
}

// o 被处理的数组区间
// l 被处理的区间的开始索引
// r 被处理的区间的结束索引
// t 临时数组
func mergerSortU2D(o *[]int, l, r int, t *[]int) {
	if o == nil || l >= r {
		return
	}

	// 优化代码，递归到一定程度改用插入排序
	if r-l <= 60 {
		new(InsertionSort).InsertSort2(o, l, r)
		return
	}

	// 计算中间值
	m := l + (r-l)/2

	mergerSortU2D(o, l, m, t)   // 对数组左边部分递归排序处理
	mergerSortU2D(o, m+1, r, t) // 对数组右边部分递归排序处理

	// 优化代码，多加一个if判断，使得无序时才归并，要注意是递增还是递减排序来确定比较
	if (*o)[m] < (*o)[m+1] {
		merger(o, l, m, r, t) // 归并为一个数组
	}
}

// MergerSortUp2Down 自上而下归并排序
func (ms *MergerSort) MergerSortUp2Down(o *[]int, args ...interface{}) interface{} {
	t := make([]int, len(*o))
	l, r := 0, len(*o)-1
	mergerSortU2D(o, l, r, &t)
	return 0
}

/*+-------------------------------------------------------+*/
// 思路分析：
// 1、将n个元素的数组看成已经排好序的n个有1个元素组成的数组
// 2、n个1个元素组成的数组两个两个归并排序，得到排好序的有2个元素组成的数组，这是第1轮
// 3、2个元素组成的数组，再次两个两个归并排序，得到排好序的有4个元素组成的数组，这是第2轮
// 4、4个元素组成的数组，再次两个两个归并排序，得到排好序的有8个元素组成的数组，这是第3轮
// 5、直到n个元素排好序
func mergerSortD2U(o *[]int, t *[]int) {
	length := len(*o)
	for i := 1; i < length; i *= 2 {
		//对 arr[j...j+i-1] 和 arr[j+i...j+2*i-1] 进行归并
		for j := 0; j+i < length; j += 2 * i { // 控制对哪个区间进行归并，j+i<length保证右部数组的存在
			l, m, r := j, j+i-1, (j+2*i)-1 // 指定归并排序的范围
			if r > length-1 {              // 防止数组越界
				r = length - 1
			}
			if (*o)[j+i-1] > (*o)[j+i] { // 优化代码，多加if判断
				merger(o, l, m, r, t)
			}
		}
	}

	/*for l := 0; l < length; l += 16 {
		r := l + 15
		if r > length-1 {
			r = length - 1
		}
		ins.InsertionSort2(o, l, r)
	}

	for i := 16; i <= length; i += i {
		for j := 0; j < length-i; j += 2 * i {
			r := (j + 2*i) - 1
			if r > length-1 {
				r = length - 1
			}
			if (*o)[j+i-1] < (*o)[j+i] {
				merger(o, j, j+i-1, r, t)
			}
		}
	}*/
}

// MergerSortDown2Up 自下而上归并排序
func (ms *MergerSort) MergerSortDown2Up(o *[]int, args ...interface{}) interface{} {
	t := make([]int, len(*o))
	mergerSortD2U(o, &t)
	return 0
}
