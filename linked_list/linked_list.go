package linkedlist

import "fmt"

//Node type in linked list
type Node struct {
	data interface{}
	next *Node
}

//LinkedList container for Nodes
type LinkedList struct {
	head *Node
}

func (l *LinkedList) append(data interface{}) {

	if l.head == nil {
		l.head = &Node{data: data}
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}

	current.next = &Node{data: data}

}

func main() {
	lList := LinkedList{}
	lList.append("foo")
	lList.append("bar")
	for i := lList.head; i != nil; i = i.next {
		fmt.Println(i)
	}
}
