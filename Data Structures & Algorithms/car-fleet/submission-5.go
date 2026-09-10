import (
	"slices"
	"cmp"
)

type Car struct {
	position int
	speed int
}

func (c Car) getTimeLeft(target int) float64 {
	return (float64(target) - float64(c.position)) / float64(c.speed)
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

func (s Stack[T]) Size() int {
	return len(s.items)
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		items: []T{},
	}
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))
	for i, _ := range position {
		cars[i] = Car{
			position: position[i],
			speed: speed[i],
		}
	}

	slices.SortFunc(cars, func(a, b Car) int {
		return cmp.Compare(b.position, a.position)
	})

	// fmt.Printf("sorted: %v \n", cars)
	stack := NewStack[Car]()
	fleet := 0
	topTimeLeft := 0.0
	for _, car := range cars {
		if stack.Size() == 0 {
			stack.Push(car)
			topTimeLeft = car.getTimeLeft(target)
			continue
		}

		// fmt.Printf("top car time %f, current car time %f\n", stack.Top().getTimeLeft(target), car.getTimeLeft(target))
		if topTimeLeft >= car.getTimeLeft(target) {
			// fmt.Printf("car @%d is catching up\n", car.position)
			// The car at the top stack (and the rest of the stack) will arrive 
			// after current car if they keep the same speed
			// => Current car will catch up to the top car, they're the same fleet
			stack.Push(car)
			continue
		}

		// fmt.Printf("car @%d is NOT catching up\n", car.position)
		
		// This car will never catch up to the previous fleet, pop all the car and
		// increase fleet count 
		fleet += 1
		for ;stack.Size() > 0; {
			stack.Pop()
		}

		stack.Push(car)
		topTimeLeft = car.getTimeLeft(target)
	}

	if stack.Size() > 0 {
		fleet += 1
	}

	return fleet
}
