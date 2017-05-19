package binary

import "github.com/redneck1/golang-algorithms/queue"

func IterByLevel(tn *TreeNode, q *queue.Queue, list *[]int) {
	if tn != nil {
		q.Enqueue(tn)
		for {
			v, err := q.Pop()
			if err != nil {
				break
			}
			n := v.(*TreeNode)
			list = append(list, n.Value)
			if n.Left != nil {
				q.Enqueue(tn.Left)
			}
			if n.Right != nil {
				q.Enqueue(tn.Right)
			}
		}
	}
}

func Iter(tn *TreeNode, vs *[]int) {
	if tn == nil {
		return
	}
	*vs = append(*vs, tn.Value)
	Iter(tn.Left, vs)
	Iter(tn.Right, vs)
}

func IterLeft(tn *TreeNode, vs *[]int) {
	if tn == nil {
		return
	}

	IterLeft(tn.Left, vs)
	*vs = append(*vs, tn.Value)
	IterLeft(tn.Right, vs)
}

func IterRight(tn *TreeNode, vs *[]int) {
	if tn == nil {
		return
	}

	IterRight(tn.Right, vs)
	*vs = append(*vs, tn.Value)
	IterRight(tn.Left, vs)
}
