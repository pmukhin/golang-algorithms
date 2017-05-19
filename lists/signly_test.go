package lists

import "testing"

func TestNode_Append(t *testing.T) {
	list := New(0)
	for i := 1; i < 20; i++ {
		list.Append(New(i))
	}

	cp := list
	for i := 0; i < 20; i++ {
		if cp.Value != i {
			t.Error("wrong order of values")
		}
		cp = cp.Next
	}
}

func TestReverse_Nonempty(t *testing.T) {
	list := New(0)
	for i := 1; i < 20; i++ {
		list.Append(New(i))
	}
	reversed := Reversed(list)
	cp := reversed

	for i := 19; i >= 0; i-- {
		if cp.Value != i {
			t.Error("wrong order of values")
		}
		cp = cp.Next
	}
}

func TestReverse_Empty(t *testing.T) {
	var list *Node = nil
	reversed := Reversed(list)

	if reversed != nil {
		t.Error("nil must be returned")
	}

	nonEmptyList := New(5)
	nonEmptyListReversed := Reversed(nonEmptyList)

	if nonEmptyList != nonEmptyListReversed {
		t.Error("must be same pointers")
	}
}

func TestGetCycled_Positive(t *testing.T) {
	var list *Node = New(0)

	cyclNode := New(1)
	list.Next = cyclNode
	list.Next.Next = New(2)
	list.Next.Next.Next = cyclNode

	found := GetCycled(list)
	if false == found {
		t.Error("cycle is not found but exists")
	}
}

func TestGetCycled_Negative(t *testing.T) {
	list := New(0)
	for i := 1; i < 20; i++ {
		list.Append(New(i))
	}

	found := GetCycled(list)
	if true == found {
		t.Error("cycle is found but not exists")
	}
}
