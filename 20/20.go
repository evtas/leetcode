package main

type Stack struct {
	Top  int32
	data []int32
}

func (s *Stack) Push(val int32) {
	s.Top += 1
	s.data[s.Top] = val
}

func (s *Stack) Pop() int32 {
	val := s.data[s.Top]
	s.Top -= 1
	return val
}

func isValid(s string) bool {

	stack := Stack{
		Top:  -1,
		data: make([]int32, len(s)),
	}
	for _, char := range s {
		if char == '{' || char == '(' || char == '[' {
			stack.Push(char)
		} else {
			if stack.Top < 0 {
				return false
			}
			val := stack.Pop()
			switch {
			case val == '(' && char != ')':
				fallthrough
			case val == '{' && char != '}':
				fallthrough
			case val == '[' && char != ']':
				return false
			}
		}
	}
	if stack.Top > -1 {
		return false
	}
	return true
}

func main() {
	res := isValid("(")
	println(res)
}
