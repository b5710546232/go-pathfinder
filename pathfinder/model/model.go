package model

// Node is a node in the grid
type Node struct {
	isWalkable bool // Is this node walkable?
	X          int  // X coordinate in Grid
	Y          int  // Y coordinate in Grid
}

// NewNode creates a new node
func NewNode(x, y int) Node {
	return Node{X: x, Y: y, isWalkable: true}
}

// SetWalkable sets whether this node is walkable
func (n *Node) SetWalkable(f bool) {
	n.isWalkable = f
}

// IsWalkable returns whether this node is walkable
func (n Node) IsWalkable() bool {
	return n.isWalkable
}

// PathNode is a node in the path
type PathNode struct {
	X int // X coordinate in Grid
	Y int // Y coordinate in Grid
}
