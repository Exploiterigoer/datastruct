package search

// OrderSearch 顺序表查找结构体
type OrderSearch struct{}

// OrdSearch 顺序表查找
// 最好的情况1次就查找到
// 最坏的情况n次才能找到
// 时间复杂度o(n)
// 平均查找次数(n+1)/2
// 当数据规模很大时效率极低
func (os *OrderSearch) OrdSearch(o *[]int, args ...interface{}) interface{} {
	key := (args[0].([]interface{}))[0].(int) // 类型断言两次得到被查找的对象
	for k, v := range *o {
		if v == key {
			return k
		}
	}
	return -1
}

// OrderSearch2 优化版的顺序查找
// 时间复杂度o(n)
// 平均查找次数(n+1)/2
// 只是避免了for k, v := range中不断赋值操作
// 可变参数 args ... T 最终会被转化为 args []T
// 与OrderSearch2相比，如果数组中存在重复的数据，得到的下标是不一样的
// 因为OrderSearch2从末尾往回查找，OrderSearch从开头往末尾查找
// 经过实际测试OrderSearch2并不比OrderSearch好
func (os *OrderSearch) OrderSearch2(o *[]int, args ...interface{}) interface{} {
	length := len(*o)
	tmp := *o
	index := length - 1
	key := (args[0].([]interface{}))[0].(int) // 类型断言两次得到被查找的对象

	if (*o)[0] == key { // 表示第1次就查到
		return 0
	}

	tmp[0] = key // 这一行使得当index=0时结束循环

	for tmp[index] != key {
		index--
	}

	if index == 0 { // 此时返回0是表示找不到，为避免混淆，用负数表示
		return -1
	}
	return index
}
