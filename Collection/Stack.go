package Collection

type Stack struct {
	nodes [] interface{}
	count int
}

func (s *Stack) Push(n interface{}) {
	s.nodes = append( s.nodes[:s.count], n )
	s.count++
}

func (s *Stack) Pop() interface{} {
	if s.count <= 0 || s.count > len(s.nodes) {
		return nil
	}

	target := s.nodes[ s.count - 1 ]
	s.count--
	s.nodes = s.nodes[ : s.count ]
	return target
}

func (s *Stack) Peek() interface{} {
	if s.count <= 0 || s.count > len(s.nodes) {
		return nil
	}

	return s.nodes[ s.count -1 ]
}

func (s *Stack) Dup() []interface{} {
	dup := make([]interface{}, len(s.nodes))
	copy(dup, s.nodes)
	return dup
}

func (s *Stack) Build(t []interface{}) {
	s.nodes = append([]interface{}{nil}, t...)
	s.count = len(s.nodes)
}

func (s *Stack) Size() int {
	return s.count
}

func (s *Stack) IsEmpty() bool {
	if s.count <= 0 {
		return true
	} else {
		return false
	}
}

func (s *Stack) Clear() {
	s.nodes = nil
	s.count = 0
}
