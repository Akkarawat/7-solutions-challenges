package binarytree

type Node struct {
    parent *Node
    left *Node
    right *Node

    value int
}

func NewNode(value int) Node {
    return Node{
        value: value,
    }
}

func (n *Node) SetParent(parent *Node) {
    n.parent = parent
}

func (n *Node) AddLeft(node *Node)  {
    node.SetParent(n)
    n.left = node
}

func (n *Node) AddRight(node *Node)  {
    node.SetParent(n)
    n.right = node
}

func (n *Node) GetLeft() *Node {
    return n.left
}

func (n *Node) GetRight() *Node {
    return n.right
}

func (n *Node) GetValue() int {
    return n.value
}

func NodesEqual(n1, n2 *Node) bool {
    if n1 == nil && n2 == nil {
        return true
    }
    if n1 == nil || n2 == nil {
        return false
    }
    if n1.value != n2.value {
        return false
    }
    return NodesEqual(n1.left, n2.left) && NodesEqual(n1.right, n2.right)
}

func (n *Node) GetMaxSum() int {
    if n == nil {
        return 0
    }
    leftSum := n.left.GetMaxSum()
    rightSum := n.right.GetMaxSum()
    if leftSum > rightSum {
        return n.value + leftSum
    }
    return n.value + rightSum
}