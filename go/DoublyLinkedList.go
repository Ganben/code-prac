package DoublyLinkedList

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (list *LinkedList) InsertFirst(i int) {
	data := &Node{data: i}
	if list.head != nil {
		list.head.prev = data
		data.next = list.head
	}
	list.head = data
}

func (list *LinkedList) InsertLast(i int) {
	data := &Node{data: i}
	if list.head == nil {
		list.head = data
		list.tail = data
		return
	}
	if list.tail != nil {
		list.tail.next = data
		data.prev = list.tail
	}
	list.tail = data
}

func (list *LinkedList) RemoveByValue(i int) bool {
	if list.head == nil {
		return false
	}
	if list.head.data == i {
		list.head = list.head.next
		list.head.prev = nil
		return true
	}
	if list.tail.data == i {
		list.tail = list.tail.prev
		list.tail.next = nil
		return true
	}
	current := list.head
	for current.next != nil {
		if current.next.data == i {
			if current.next.next != nil {
				current.next.next.prev = current
			}
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList) RemoveByIndex(i int) bool {

}

func (list *LinkedList) SearchValue(i int) bool {

}

func (list *LinkedList) GetFirst() (int, bool) {

}

func (list *LinkedList) GetLast() (int, bool) {

}

func (list *LinkedList) GetSize() int {

}

func (list *LinkedList) GetItemsFromStart() []int {

}

func (list *LinkedList) GetItemsFromEnd() []int {
	
}