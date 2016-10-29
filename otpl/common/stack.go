package common

type Stack []interface{}

func (stack *Stack) Push(v interface{}) interface{} {
	*stack = append(*stack, v)
	return v
}

func (stack *Stack) Pop() (interface{}, bool) {
	size := len(*stack)
	if size == 0 {
		return nil, false
	}
	v := (*stack)[size-1]
	*stack = (*stack)[:size-1]
	return v, true
}

func (stack *Stack) Len() int {
	return len(*stack)
}
