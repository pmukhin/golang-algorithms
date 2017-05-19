package lists

type Node struct {
	Next  *Node
	Value interface{}
}

func New(value interface{}) *Node {
	n := new(Node)
	n.Value = value

	return n
}

func (n *Node) Append(node *Node) {
	if n.Next == nil {
		n.Next = node
		return
	}
	n.Next.Append(node)
}

func Reversed(list *Node) *Node {
	if list == nil || list.Next == nil {
		return list
	}
	var (
		current *Node = list
		tmp     *Node = nil
		newF    *Node = nil
	)
	for current != nil {
		tmp = current.Next
		current.Next = newF
		newF = current
		current = tmp
	}

	return newF
}

func GetCycled(list *Node) bool {
	var (
		cp   *Node = list
		fast *Node = list.Next
	)

	for cp.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		if fast == cp {
			return true
		}

		cp = cp.Next
		fast = fast.Next.Next
	}

	return false
}
