package stackqueue

// Stack 栈结构体，这里利用slice来实现
type Stack struct {
	container []interface{}
	size      int // 栈的容量
}

// CreateStack 创建栈对象
func CreateStack(size int) *Stack {
	s := new(Stack)
	s.size = size
	return s
}

// Push 压栈操作
// 利用slice的append增加元素，从尾部添加元素，O(1)的时间复杂度
func (s *Stack) Push(data interface{}) bool {
	// 如果栈空间已用完，压栈失败，否则压栈成功
	if len(s.container) == s.size {
		return false
	}
	s.container = append(s.container, data)
	return true
}

// Pop 弹栈操作
func (s *Stack) Pop() interface{} {
	// 如果栈对象不存在或栈为空，弹栈失败
	length := s.Size()
	if s == nil || length == 0 {
		return false
	}
	// 记录最后的一个元素，再删除栈中对应的位置
	tmp := s.container[length-1]
	s.container = s.container[:length-1]
	return tmp
}

// Peek 获取栈顶元素
func (s *Stack) Peek() interface{} {
	// 如果栈对象不存在或栈为空，弹栈失败
	length := s.Size()
	if s == nil || length == 0 {
		return false
	}
	return s.container[length-1]
}

// IsEmpty 判断是否为空栈
func (s *Stack) IsEmpty() bool {
	return s == nil || s.Size() == 0
}

// Size 获取栈元素个数
func (s *Stack) Size() int {
	return len(s.container)
}
