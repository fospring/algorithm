package binary_tree

import (
	"fmt"
)

type TreeNode struct {
	val   int32
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) String() string {
	return "value: " + string(t.val)
}

func (t *TreeNode) Visit() {
	fmt.Printf("node.val: %d \n", t.val)
}

// 前序遍历：根节点 -> 左孩子节点 -> 右孩子节点，递归版本
func PreOrderRecursive(t *TreeNode) {
	if t == nil {
		return
	}
	t.Visit()
	PreOrderRecursive(t.left)  // 访问左孩子
	PreOrderRecursive(t.right) // 访问右孩子
}

// 前序遍历：根节点 -> 左孩子节点 -> 右孩子节点，非递归版本，利用栈
func PreOrderTraversal(t *TreeNode) {
	treeStack := NewStack()
	if t == nil {
		return
	}
	treeStack.Push(t)
	for !treeStack.IsEmpty() {
		tempNode := treeStack.Pop()
		if tempNode != nil {
			tempNode.Visit()
			treeStack.Push(tempNode.right)
			treeStack.Push(tempNode.left)
		}
	}
}

// 中序遍历： 左孩子 -> 根节点 -> 右孩子； 递归版本
func MinOrderRecursion(t *TreeNode) {
	if t == nil {
		return
	}
	MinOrderRecursion(t.left)  // 访问左孩子
	t.Visit()                  // 访问根节点
	MinOrderRecursion(t.right) // 访问右孩子
}

// 后续遍历(LDR，先遍历子树，再遍历根的遍历方式)
func PostOrderTraversal(t *TreeNode) {
	if t == nil {
		return
	}
	PostOrderTraversal(t.left)
	PostOrderTraversal(t.right)
	t.Visit()
}

// 层次遍历，广度优先搜索，一层一层搜索，非递归代码如下
func LayerTraversal(t *TreeNode) [][]int32 {
	resultList := make([][]int32, 0)
	levelNum := 0 // 记录某层具有多少个节点
	treeQueue := NewLinkQueue()
	treeQueue.EnQueue(t)
	for !treeQueue.IsEmpty() {
		levelNum = treeQueue.Size()
		levelList := make([]int32, 0)
		for levelNum > 0 {
			tempNode := treeQueue.DeQueue()
			levelList = append(levelList, tempNode.val)
			treeQueue.EnQueue(tempNode.left)
			treeQueue.EnQueue(tempNode.right)
		}
		levelNum--
		if len(levelList) > 0 {
			resultList = append(resultList, levelList)
		}
	}
	return resultList
}

// 层次遍历，蛇形广度优先搜索，一层一层搜索，非递归代码如下
func SnackLayerTraversal(t *TreeNode) [][]int32 {
	isLeft := true
	resultList := make([][]int32, 0)
	levelNum := 0 // 记录某层具有多少个节点
	treeStack := NewStack()
	treeStack.Push(t)
	for !treeStack.IsEmpty() {
		levelNum = treeStack.Size()
		levelList := make([]int32, 0)
		for levelNum > 0 && isLeft {
			tempNode := treeStack.Pop()
			levelList = append(levelList, tempNode.val)
			treeStack.Push(tempNode.left)
			treeStack.Push(tempNode.right)
		}
		levelNum--
		if len(levelList) > 0 {
			resultList = append(resultList, levelList)
		}
		isLeft = !isLeft
	}
	return resultList
}

type TStack struct {
	stack []*TreeNode
}

func NewStack() TStack {
	return TStack{stack: make([]*TreeNode, 0)}
}

func (s *TStack) Pop() *TreeNode {
	if s.IsEmpty() {
		return nil
	}
	node := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return node
}

func (s *TStack) Push(node *TreeNode) {
	if node == nil {
		return
	}
	s.stack = append(s.stack, node)
}

func (s *TStack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *TStack) Size() int {
	return len(s.stack)
}

type LinkedQueue struct {
	queue []*TreeNode
}

func NewLinkQueue() LinkedQueue {
	return LinkedQueue{queue: make([]*TreeNode, 0)}
}

func (q *LinkedQueue) EnQueue(t *TreeNode) {
	if t == nil {
		return
	}
	q.queue = append(q.queue, t)
}

func (q *LinkedQueue) DeQueue() *TreeNode {
	if len(q.queue) == 0 {
		return nil
	}
	h := q.queue[0]
	q.queue = q.queue[1:]
	return h
}

func (q *LinkedQueue) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *LinkedQueue) Size() int {
	return len(q.queue)
}
