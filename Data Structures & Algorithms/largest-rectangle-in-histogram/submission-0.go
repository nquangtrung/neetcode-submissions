
type HeightAndIndex struct {
	height int
	index int
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(o T) {
	s.items = append(s.items, o)
}

func (s *Stack[T]) Pop() T {
	top, rest := s.items[len(s.items) - 1], s.items[:len(s.items) - 1]
	s.items = rest
	return top
}

func (s Stack[T]) Top() T {
	return s.items[len(s.items) - 1]
}

func (s Stack[T]) Bottom() T {
	return s.items[0]
}

func (s Stack[T]) Size() int {
	return len(s.items)
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		items: []T{},
	}
}

func largestRectangleArea(heights []int) int {
	stack := NewStack[HeightAndIndex]()
	maxArea := 0
	for i, h := range heights {
		left := i
		for ;stack.Size() > 0 && stack.Top().height >= h; {
			current := stack.Pop()
			area := current.height * (i - current.index)
			maxArea = max(area, maxArea)
			left = current.index
		}

		if stack.Size() > 0 && stack.Top().height == h {
			left = stack.Top().index
		} 
		stack.Push(HeightAndIndex{
			height: h,
			index: left,
		})

		// fmt.Printf("stack: %v\n", stack)
	}

	for ;stack.Size() > 0; {
		current := stack.Pop()
		area := current.height * (len(heights) - current.index)
		maxArea = max(area, maxArea)
		// fmt.Printf("stack: %v\n", stack)
	}

	return maxArea
}
