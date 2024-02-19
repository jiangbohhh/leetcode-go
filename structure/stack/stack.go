package stack

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{data: []int{}}
}

func (s *Stack) Push(n int) {
	s.data = append(s.data, n)
}

func (s *Stack) Pop() int {
	res := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return res
}

func (s *Stack) Peek() int {
	return s.data[len(s.data)-1]
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
