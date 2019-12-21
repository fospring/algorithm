package graph

// 先进先出队列
type NodeQueue struct {
	Items []*Node
}

func NewNodeQueue() *NodeQueue  {
	s := NodeQueue{Items: []*Node{}}
	return &s
}

func (s *NodeQueue) Enqueue(t *Node)  {
	s.Items = append(s.Items, t)
}

func (s *NodeQueue) Dequeue() *Node {
	item := s.Items[0]
	s.Items = s.Items[1:len(s.Items)]
	return item
}

func (s *NodeQueue) IsEmpty() bool  {
	return len(s.Items) == 0
}