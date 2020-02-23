package LinkedList

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) InsertFirst(i int) {
	newNode := &Node{data: i}
	if list.head != nil {
		newNode.next = list.head
	}
	list.head = newNode
}

func (list *LinkedList) InsertLast(i int) {
	newNode := &Node{data: i}
	if list.head == nil {
		list.head = newNode
		return
	}
	currentNode := list.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode
}

func (list *LinkedList) RemoveByValue(i int) bool {
	if list.head == nil {
		return false
	}
	if list.head.data == i {
		list.head = list.head.next
		return true
	}
	currentNode := list.head
	for currentNode.next != nil {
		if currentNode.next.data == i {
			currentNode.next = currentNode.next.next
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (list *LinkedList) RemoveByIndex(i int) bool {
	if list.head == nil {
		return false
	}
	if i < 0 {
		return false
	}
	if i == 0 {
		list.head = list.head.next
		return true
	}
	currentNode := list.head
	for u:= 1; u < i; u++ {
		if currentNode.next.next == nil {
			return false
		}
		
		currentNode = currentNode.next
	}
	currentNode.next = currentNode.next.next
	return true
}

func (list *LinkedList) SearchValue(i int) bool {
	if list.head == nil {
		return false
	}
	currentNode := list.head
	for currentNode != nil {
		if currentNode.data == i {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (list *LinkedList) GetFirst() (int, bool) {
	if list.head == nil {
		return 0, false
	}
	return list.head.data, true
}

func (list *LinkedList) GetLast() (int, bool) {
	if list.head == nil {
		return 0, false
	}
	currentNode := list.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode.data, true

}

func (list *LinkedList) GetSize() int {
	sizeCount := 0
	currentNode := list.head
	for currentNode != nil {
		sizeCount += 1
		currentNode = currentNode.next
	}
	return sizeCount
}

func (list *LinkedList) GetItems() []int {
	var items []int
	currentNode := list.head
	for currentNode != nil {
		items = append(items, currentNode.data)
		currentNode = currentNode.next
	}
	return items
}