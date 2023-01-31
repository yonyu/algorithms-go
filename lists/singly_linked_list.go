package lists

type Node struct {
	Item interface{}
	Next *Node
}

func NewNode(item interface{}, node *Node) *Node {
	return &Node {item, node}
}