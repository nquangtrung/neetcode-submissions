type Stack struct {
	items []int
}

func (s *Stack) Push(o int) {
	s.items = append(s.items, o)
}

func (s *Stack) Pop() int {
	top, rest := s.items[len(s.items) - 1], s.items[:len(s.items) - 1]
	s.items = rest
	return top
}

func (s Stack) Top() int {
	return s.items[len(s.items) - 1]
}

func (s Stack) Size() int {
	return len(s.items)
}

func NewStack() Stack {
	return Stack{
		items: []int{},
	}
}

func dailyTemperatures(temperatures []int) []int {
	stack := NewStack()
	result := make([]int, len(temperatures))

	for i, t := range temperatures {
		for ;stack.Size() > 0 && t > temperatures[stack.Top()]; {
			result[stack.Top()] = i - stack.Top()
			stack.Pop()
		}

		stack.Push(i)
	}

	for ;stack.Size() > 0; {
		top := stack.Pop()
		result[top] = 0
	}

	return result
}
