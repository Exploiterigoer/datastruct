package search

// BinarySearch 二分法查找结构体
type BinarySearch struct{}

// BinSearch 二分查找，非递归版
// 如果找到了key并不是结束函数而是结束循环
// 那么right不断往下标0方向移动，left不断往length-1方向移动
func (bs *BinarySearch) BinSearch(o *[]int, args ...interface{}) interface{} {
	length := len(*o)
	left, right := 0, length-1
	mid := left + (right-left)/2 // 避免溢出
	key := (args[0].([]interface{}))[0].(int)
	for left <= right { // 等号不能少，不然恰好被查找的值在left = right处取不到
		if key < (*o)[mid] {
			right = mid - 1 // 被查找的值可定还在左半区，调整右半区索引
		} else if key > (*o)[mid] {
			left = mid + 1 // 被查找的值可定还在右半区，调整左半区索引
		} else {
			return mid // 找到了就结束函数
		}
		mid = (left + right) / 2 // 存在整型数据溢出的风险
	}
	return -1
}

// BinSearch2 递归版二分查找，这种递归只能做到检测是否存在被搜索的值而无法做到值具体的下标
// 要想找到具体所在的下标，必须改造整个递归函数，此时需要传入左，右两个边界下标值
func (bs *BinarySearch) BinSearch2(o *[]int, args ...interface{}) interface{} {
	if len(*o) > 0 {
		mid := (len(*o) - 1) / 2
		var key int
		switch args[0].(type) {
		case []interface{}:
			key = (args[0].([]interface{}))[0].(int)
		case int:
			key = args[0].(int)
		}

		if key == (*o)[mid] {
			return 0
		} else if key < (*o)[mid] {
			tmp := (*o)[:mid]
			return bs.BinSearch2(&tmp, key)
		} else {
			tmp := (*o)[mid+1:]
			return bs.BinSearch2(&tmp, key)
		}
	}
	return -1
}

// 二分查找有很多变种的问题：
// 找出第一个与key相等的元素
// 找出最后一个与key相等的元素
// 找出第一个等于或者大于Key的元素
// 找出最后一个等于或者小于key的元素
// 找出第一个大于key的元素
// 找出最后一个小于key的元素
// 关键词：
// 第一个等于或者大于Key	(return left,mid-1)	 --- 最后一个等于或者小于key	(return right)
// 第一个大于Key			(return left,mid+1)	--- 最后一个小于key			   (return right)
// 第一个与key相等			(return left,mid-1)	--- 最后一个与key相等		   (return right)

// GetFirstEqKey 找出第一个等于key的元素，尽量保持left少移动
func (bs *BinarySearch) GetFirstEqKey(o *[]int, args ...interface{}) interface{} {
	length := len(*o)
	left, right := 0, length-1
	mid := left + (right-left)/2
	key := (args[0].([]interface{}))[0].(int)

	for left <= right {
		if key <= (*o)[mid] { // 这里相等时加速了right的移动，left尽量不移动
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = left + (right-left)/2
	}

	if left < length && (*o)[left] == key { // 找到相等的元素的下标
		return left
	}
	return -1
}

// GetLastEqKey 找出最后一个与key相等的元素，尽量保持right少移动
func (bs *BinarySearch) GetLastEqKey(o *[]int, args ...interface{}) interface{} {
	length := len(*o)
	left, right := 0, length-1
	mid := left + (right-left)/2
	key := (args[0].([]interface{}))[0].(int)
	for left <= right {
		if key < (*o)[mid] {
			right = mid - 1
		} else { // key=(*o)[mid]的情况使得left加速右移
			left = mid + 1
		}
		mid = left + (right-left)/2
	}
	if right >= 0 && (*o)[right] == key { // 找到相等的元素的下标
		return right
	}
	return -1
}

// GetFirstGtEqKey 找出第一个等于或者大于Key的元素，尽量保持left少移动
// 因为越往右值越大，只能rigth加速向右移动才能更快找到第一个等于或者大于Key的元素
// 所以等号边界在 mid - 1分支
func (bs *BinarySearch) GetFirstGtEqKey(o *[]int, args ...interface{}) interface{} {
	length := len(*o)
	left, right := 0, length-1
	mid := left + (right-left)/2
	key := (args[0].([]interface{}))[0].(int)
	for left <= right {
		if key <= (*o)[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = left + (right-left)/2
	}
	return left
}

// GetLastLteqkey 表示查找最后一个等于或者小于key的元素，尽量保持right少移动
// 因为越往右值越大，只能left加速向右移动才能更快找到最后一个等于或者小于key的元素
// 所以等号边界在 mid + 1分支
func (bs *BinarySearch) GetLastLteqkey(o *[]int, args ...interface{}) interface{} {
	length := len(*o)
	left, right := 0, length-1
	mid := left + (right-left)/2
	key := (args[0].([]interface{}))[0].(int)
	for left <= right {
		if key < (*o)[mid] {
			right = mid - 1
		} else {
			left = mid + 1 // 这里包含了key >= (*o)[mid]两情况，相等时加速了left的移动，right尽量不移动
		}
		mid = left + (right-left)/2
	}

	// 如果数组中一定存在多个相同的值，for循环结束后，right一定是偏左半区的
	return right
}

// GetFirstGtKey 找出第一个大于key的元素，尽量保持right少移动
// 因为越往左值越小，只能left加速向右移动才能更快找到第一个大于key的元素
// 所以等号边界在 mid + 1分支
func GetFirstGtKey(o *[]int, args ...interface{}) interface{} {
	length := len(*o)
	left, right := 0, length-1
	mid := left + (right-left)/2
	key := (args[0].([]interface{}))[0].(int)

	for left <= right {
		if key < (*o)[mid] {
			right = mid - 1
		} else {
			left = mid + 1 // 这里包含了key >= (*o)[mid]两情况，相等时加速了left的移动，right尽量不移动
		}
		mid = left + (right-left)/2
	}
	// 如果数组中一定存在多个相同的值，for循环结束后，left一定是偏右半区的
	return left
}

// GetLastLtKey 找出最后一个小于key的元素，尽量保持left少移动
// 因为越往右值越大，只能right加速向左移动才能更快找到最后一个小于key的元素
// 所以等号边界在 mid - 1分支
func GetLastLtKey(o *[]int, args ...interface{}) interface{} {
	length := len(*o)
	left, right := 0, length-1
	mid := left + (right-left)/2
	key := (args[0].([]interface{}))[0].(int)
	for left <= right {
		if key <= (*o)[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = left + (right-left)/2
	}
	return right
}
