package brackets

type RuneStack struct {
	Data []rune
	Pos  int
}

func NewRuneStack() *RuneStack {
	return &RuneStack{[]rune{}, -1}
}

func (stack *RuneStack) Push(r rune) {
	stack.Pos++
	if stack.Pos == len(stack.Data) {
		stack.Data = append(stack.Data, r)
	} else {
		stack.Data[stack.Pos] = r
	}
}

func (stack *RuneStack) Pop() (r rune, ok bool) {
	if stack.Pos == -1 {
		return
	}
	r, ok = stack.Data[stack.Pos], true
	stack.Pos--
	return
}

func (stack *RuneStack) Empty() bool {
	return stack.Pos == -1
}

var pair = map[rune]rune{
	']': '[',
	'}': '{',
	')': '(',
}

func Bracket(input string) bool {
	stack := NewRuneStack()
	for _, r := range input {
		switch r {
		case '[', '{', '(':
			stack.Push(r)
		case ']', '}', ')':
			got, ok := stack.Pop()
			if !ok || got != pair[r] {
				return false
			}
		}
	}
	return stack.Empty()
}
