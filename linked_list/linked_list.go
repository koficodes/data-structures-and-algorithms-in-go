package linkedlist

//Node type in linked list
type Node struct {
	data interface{}
	next *Node
}

//LinkedList container for Nodes
type LinkedList struct {
	head *Node
}
