package act

type Queue struct {
	ch chan interface{}
}

func NewQueue() *Queue {
	q := &Queue{
		ch: make(chan interface{}),
	}
	return q
}

func (q *Queue) Push() chan<- interface{} {
	return q.ch
}

func (q *Queue) Pop() <-chan interface{} {
	return q.ch
}
