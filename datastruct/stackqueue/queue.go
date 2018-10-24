package stackqueue

// Queue 队列结构体，这里利用slice来实现，并且是单端队列
type Queue struct {
	Container []interface{}
	size      int // 队列的长度
}

// CreateQueue 创建队列对象
func CreateQueue(size int) *Queue {
	q := new(Queue)
	q.size = size
	return q
}

// Enqueue 入队操作，这里从尾部入队，时间复杂度为O(1)
func (q *Queue) Enqueue(data interface{}) bool {
	if q == nil || len(q.Container) == q.size {
		return false
	}
	q.Container = append(q.Container, data)
	return true
}

// Dequeue 出队操作，这里从头部出队，时间复杂度为O(1)
// 得益于golang的slice的特性，如果利用链表也可以做到O(1)，但是如果入队就是O(n)
func (q *Queue) Dequeue() interface{} {
	if q == nil || q.Size() == 0 {
		return false
	}
	tmp := q.Container[0]
	q.Container = q.Container[1:]
	return tmp
}

// IsEmpty 判断是否为空队列
func (q *Queue) IsEmpty() bool {
	return q == nil || len(q.Container) == 0
}

// Size 获取队列元素个数
func (q *Queue) Size() int {
	return len(q.Container)
}
