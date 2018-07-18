package Collection

type Queue struct {
	nodes [] interface{}
	count int
}

func (q *Queue) Push(n interface{}) {
	q.nodes = append( q.nodes[:q.count], n )
	q.count++
}

func (q *Queue) Pop() interface{} {
	if q.count <= 0 || q.count > len(q.nodes) {
		return nil
	}

	target := q.nodes[0]
	q.nodes = q.nodes[ 1 : ]
	q.count--
	return target
}

func (q *Queue) Peek() interface{} {
	if q.count <= 0 || q.count > len(q.nodes) {
		return nil
	}

	return q.nodes[ 0 ]
}

func (q *Queue) Dup() []interface{} {
	dup := make([]interface{}, len(q.nodes))
	copy(dup, q.nodes)
	return dup
}

func (q *Queue) Build(t []interface{}) {
	q.nodes = append([]interface{}{nil}, t...)
	q.count = len(q.nodes)
}

func (q *Queue) Size() int {
	return q.count
}

func (q *Queue) IsEmpty() bool {
	if q.count <= 0 {
		return true
	} else {
		return false
	}
}

func (q *Queue) Clear() {
	q.nodes = nil
	q.count = 0
}
