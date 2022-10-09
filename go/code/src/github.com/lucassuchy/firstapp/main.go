package main

// declaraçãod de variaveis globais precisam definir o tipo
import 
	"fmt"

type carta struct {
	jogador string
	numero  int
}


// https://github.com/silenceshell/algorithms-in-golang/blob/master/data-structure/linked-list.go
type Node struct {
	next *Node
	prev *Node
	key  interface{}
	carta carta
}

type List struct {
	head *Node
	tail *Node
}

func (list *List) insert(key interface{} carta) {
	node := &Node{next: list.head, prev: nil, key: key, carta: carta}
	if list.head != nil {
		list.head.prev = node
	}
	list.head = node

	tmp := list.head
	// todo: is this needed check every time?
	for tmp.next != nil {
		tmp = tmp.next
	}
	list.tail = tmp
}

func (list *List) display() {
	node := list.head
	for node != nil {
		fmt.Printf(" %v ", node.key)
		node = node.next
	}
	fmt.Println()
}

func main() {
	list.insert("lucas", 10);
	
}
