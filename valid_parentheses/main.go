package main

func isClose(s rune) (bool, int) {
	switch s {
	case '(':
		return true, 1
	case '[':
		return true, 2
	case '{':
		return true, 3
	case ')':
		return false, 1
	case ']':
		return false, 2
	case '}':
		return false, 3
	}

	return false, 0
}

func isValid(s string) bool {
	stack := []int{}
	for _, v := range s {
		isOpen, pType := isClose(v)

		if pType == 0 {
			continue
		}

		if isOpen {
			stack = append(stack, pType)
		} else {
			if len(stack) == 0 {
				return false
			}
			x, newStack := stack[len(stack)-1], stack[:len(stack)-1]
			stack = newStack

			if x != pType {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {

}
