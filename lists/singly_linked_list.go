package lists

type SinglyLinkedList struct {
	First  *SinglyLinkedNode
	Length int
}

type SinglyLinkedNode struct {
	Next *SinglyLinkedNode
	Value interface{}
}

func (l *SinglyLinkedList) Append(value interface{}) {
	if l.First == nil {
		l.First = &SinglyLinkedNode{Value:value}
		l.Length++
		return
	}
	firstCopy := l.First
	for {
		if firstCopy.Next == nil {
			break
		}
		firstCopy = firstCopy.Next
	}
	firstCopy.Next = &SinglyLinkedNode{Value:value}
	l.Length++
}

func (l *SinglyLinkedList) Reverse() {
	// Check that we can manipulate pointers
	if l.Length <= 1 || l.First == nil || l.First.Next == nil {
		return
	}
	// Now we create a copy of our current top node
	curr := l.First
	var (
		// Let's create a nil pointer for our new top node
		newFirst *SinglyLinkedNode = nil
		// Let's create a nil pointer for storing every next pointer
		temp *SinglyLinkedNode = nil
	)

	for {
		// if current is nil then break
		if curr == nil {
			break
		}
		// Let's store next value in temporary pointer
		temp = curr.Next
		// Next is now newFirst (in first iteration it's nil)
		curr.Next = newFirst
		// newFirst is now the current node
		newFirst = curr
		// curr becomes 'the current next'
		curr = temp
	}

	l.First = newFirst
}