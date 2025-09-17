package DataStructures

type Stack struct {
	elements []string
	len      int
}

func (s *Stack) Push(element string) {
	s.elements = append(s.elements, element)
	s.len++
}

func (s *Stack) Pop() string {
	if s.len == 0 {
		return ""
	}
	s.len--
	value := s.elements[s.len]
	s.elements = s.elements[:s.len]
	return value
}

func (s *Stack) Peek() string {
	if s.len == 0 {
		return ""
	}
	return s.elements[s.len-1]
}

func (s *Stack) Contains(item string) bool {
	for _, element := range s.elements {
		if element == item {
			return true
		}
	}

	return false
}

func (s *Stack) Len() int {
	return s.len
}

func CreateStack() *Stack {
	return &Stack{
		elements: []string{},
		len:      0,
	}
}
