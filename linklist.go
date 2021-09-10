package lru

type node struct {
	key   string
	value interface{}
	next  *node
	prev  *node
}

type linklist struct {
	head *node
	tail *node
}

func newLinkedList() *linklist {
	return &linklist{}
}

func (l *linklist) add(key string, value interface{}) *node {
	newNode := &node{
		key:   key,
		value: value,
	}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		tmp := l.head
		l.head = newNode
		newNode.next = tmp
		tmp.prev = l.head
	}

	return newNode
}

func (l *linklist) delete(n *node) {
	// linklist is empty
	if l.head == nil {
		return
	}

	// just contains 1 item
	if l.head == l.tail && l.head == n {
		l.head = nil
		l.tail = nil
		return
	}

	// we are going to delete head
	if l.head == n {
		l.head = n.next
		return
	}

	if l.tail == n {
		l.tail = n.prev
		return
	}

	n.prev.next = n.next
	n.next.prev = n.prev
}
