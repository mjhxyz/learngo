package queue

// Queue 一个先进先出的队列
type Queue []int

// Push 将元素添加到队列中
//
//	e.g. q.Push(123)
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pop 将元素从队列中移除
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// IsEmpty 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
