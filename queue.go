package act

type Queue struct {
	in  chan interface{}
	out chan interface{}
}

func NewQueue() *Queue {
	q := &Queue{
		in:  make(chan interface{}),
		out: make(chan interface{}),
	}
	go func() {
		var last interface{}
		var buf []interface{}

		last = <-q.in

		for {
			select {
			case value := <-q.in:
				buf = append(buf, value)

			case q.out <- last:
				sz := len(buf)
				if sz == 0 {
					last = <-q.in
				} else {
					last, buf = buf[0], buf[1:]
				}
			}
		}
	}()
	return q
}

func (q *Queue) Push() chan<- interface{} {
	return q.in
}

func (q *Queue) Pop() <-chan interface{} {
	return q.out
}
