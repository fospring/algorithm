package graph

import (
	"fmt"
)

type Item interface {
}

// node
type Node struct {
	Value Item
}
// 图的结构，顶点与边组成 V E
type ItemGraph struct {
	Nodes []*Node
	Edges map[*Node][]*Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

func (n *Node) Visit(s string) {
	fmt.Printf("%s, %v \n", s, n)
}

// add node
func (g *ItemGraph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

// add edge
func (g *ItemGraph) AddEdge(n1, n2 *Node) {
	if g.Edges == nil {
		g.Edges = make(map[*Node][]*Node)
	}

	// 无向图
	g.Edges[n1] = append(g.Edges[n1], n2) // n1 -> n2
	g.Edges[n2] = append(g.Edges[n2], n1) // n2 -> n1
}

func (g *ItemGraph) String() {
	s := ""

	for i := 0; i < len(g.Nodes); i++ {
		s += g.Nodes[i].String() + "->"
		near := g.Edges[g.Nodes[i]]

		for j := 0; j < len(near); j ++ {
			s += near[j].String()
		}
		s += "\n"
	}

	fmt.Println(s)
}

// 广度优先搜索 Breadth-first traversal
func (g *ItemGraph) Bfs () {
	q := NewNodeQueue()

	n := g.Nodes[0]
	q.Enqueue(n)

	visited := make(map[*Node]bool)
	visited[n] = true

	for {
		if q.IsEmpty() {
			break
		}
		node := q.Dequeue()
		near := g.Edges[node]

		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j] {
				q.Enqueue(j)
				visited[j] = true
			}
		}

		if node != nil {
			node.Visit("BFS visiting ")
		}
	}
}

// 深度优先搜索 Depth-First traversal
func (g *ItemGraph) Dfs () {
	stack := NewNodeStack()

	n := g.Nodes[0]
	stack.push(n)


	visited := make(map[*Node] bool)


	visited[n] = true


	for {
		if stack.IsEmpty(){
			break
		}


		node := stack.pop()
		if !visited[node]{
			visited[node] = true
		}
		near := g.Edges[node]


		for i:= 0; i< len(near); i++{
			j := near[i]
			if !visited[j]{
				visited[j] = true
				// 先进后出后访问
				stack.push(j)
			}
		}


		if node != nil{
			node.Visit("DFS visiting ")
		}
	}
}