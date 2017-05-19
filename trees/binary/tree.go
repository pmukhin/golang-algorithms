package binary

type TreeNode struct {
  Left  *TreeNode
  Right *TreeNode
  Value int
}

func New(v int) *TreeNode {
  tn := new(TreeNode)
  tn.Value = v

  return tn
}

func (t *TreeNode) Append(node *TreeNode) {
  if node.Value < t.Value {
    if t.Left == nil {
      t.Left = node
    } else {
      t.Left.Append(node)
    }
  } else {
    if t.Right == nil {
      t.Right = node
    } else {
      t.Right.Append(node)
    }
  }
}
