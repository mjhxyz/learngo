package queue

type E interface{}
type Queue []E

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	// 强制转换
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
