package act

type Stack struct {
	in  chan interface{}
	out chan interface{}
}

func NewStack(size int) *Stack {
	s := &Stack{
		in:  make(chan interface{}, size),
		out: make(chan interface{}, size),
	}
	go func() {
		var last interface{}
		var buf []interface{}

		last = <-s.in

		for {
			select {
			case value := <-s.in:
				buf = append(buf, last)
				last = value

			case s.out <- last:
				sz := len(buf)
				if sz == 0 {
					last = <-s.in
				} else {
					last, buf = buf[sz-1], buf[:sz-1]
				}
			}
		}
	}()
	return s
}

func (s *Stack) Push() chan<- interface{} {
	return s.in
}

func (s *Stack) Pop() <-chan interface{} {
	return s.out
}
