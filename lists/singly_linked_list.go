package lists

type SinglyLinkedList struct {
	First  *SinglyLinkedNode
	Length int
}

type SinglyLinkedNode struct {
	Next  *SinglyLinkedNode
	Value interface{}
}

func (l *SinglyLinkedList) Append(value interface{}) {
	newNode := &SinglyLinkedNode{Value: value}

	if l.Length == 0 {
		l.First = newNode
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
	firstCopy.Next = newNode
	l.Length++
}

func (l *SinglyLinkedList) Reverse() {
	// First let's check length. If it's
	// less or equals to 1, there's nothing
	// to reverse, actually...
	if l.Length <= 1 {
		return
	}
	// Now will need three pointers
	// One for a copy of l.First
	// One for a temporary copy of every next node
	// And one for storing every previos copy (which
	// is going to be nil in first iteration)
	var (
		current *SinglyLinkedNode = l.First
		temp    *SinglyLinkedNode = nil
		prev    *SinglyLinkedNode = nil
	)
	for {
		// If current element is nil,
		// The list is reversed
		if current == nil {
			break
		}
		// Copy next pointer to temporary
		// Because we're going to lose it otherwise
		temp = current.Next
		// Now current.next is previos (which is nil
		// In 1st iteration
		current.Next = prev
		// Now prev is current element, in next iteration
		// It's not nil anymore
		prev = current
		// Current is not temp, which is a copy of current.Next
		current = temp
	}

	// Now the top of tree is the previous one (which is not nil)
	l.First = prev
}
