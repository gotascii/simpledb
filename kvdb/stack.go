package kvdb

type stack []*Transaction

func NewStack() stack {
	return make(stack, 0)
}

func (s stack) Push(t *Transaction) stack {
	return append(s, t)
}

func (s stack) Pop() (stack, *Transaction) {
	l := len(s)
	return s[:l-1], s[l-1]
}
