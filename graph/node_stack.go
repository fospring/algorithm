package graph

type NodeStack struct {
	Items []*Node
}


func NewNodeStack() *NodeStack {
	return &NodeStack{Items: []*Node{}}
}


func (n *NodeStack) push(q *Node) {
	n.Items = append(n.Items, q)
}


func (n *NodeStack) pop() *Node {


	item := n.Items[len(n.Items)-1]


	n.Items = n.Items[0: len(n.Items)-1]
	return item
}


func (n *NodeStack) IsEmpty() bool {
	return len(n.Items) == 0
}


func (n *NodeStack) Size() int {
	return len(n.Items)
}
