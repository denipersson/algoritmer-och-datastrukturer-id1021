package main

type Linked_List struct {
	head *Node
}

func (l Linked_List) get_head() *Node {
	return l.head
}
func (l *Linked_List) append(new_list Linked_List) {
	cur := l.head

	for cur.next != nil {
		cur = cur.next
	}
	cur.next = new_list.get_head()
}

func (l *Linked_List) add_to_front(n *Node) {
	n.next = l.head
	l.head = n
}
func (l *Linked_List) remove(node *Node) {
	cur := l.head
	last_item := cur
	for cur.next != nil { //search through until the node to remove is found, and while keeping track of the previous node
		if cur == node {
			last_item.next = cur.next
			break
		}
		last_item = cur
		cur = cur.next
	}
}

type Double_Linked_List struct {
	head *Double_Node
}

func (l Double_Linked_List) get_head() *Double_Node {
	return l.head
}
func (l *Double_Linked_List) append(new_list Double_Linked_List) {
	cur := l.head

	for cur.next != nil {
		cur = cur.next
	}
	cur.next = new_list.get_head()
}
func (l *Double_Linked_List) add_to_front(n *Double_Node) {
	k := l.head
	l.head = n
	n.next = k
	k.previous = n
}
func (l *Double_Linked_List) remove(node *Double_Node) {

	old := node
	if node.previous != nil { // if it is not the first node in the list
		node = node.previous //make node point to the removed nodes previous node
		node.next = old.next //and make the next reference to point to the old nodes next node
	} else { // if it is the first object in the list
		if node.next != nil { //and also is not the last object in the list
			node = node.next // set node to point to the next node in the list
			l.head = node    //and make this the head of the list
		}
	}
}
