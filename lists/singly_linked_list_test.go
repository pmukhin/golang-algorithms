package lists

import (
	"testing"
)

func TestSinglyLinkedList_Append(t *testing.T) {
	list := SinglyLinkedList{}
	list.Append(5)
	if list.First.Value != 5 {
		t.Fatal("Must equal")
	}
	if list.Length != 1 {
		t.Fatal("Length must be incremented")
	}
	list.Append(10)
	if list.First.Next.Value != 10 {
		t.Fatal("Must equal")
	}
	if list.Length != 2 {
		t.Fatal("Length must be incremented")
	}
}

func TestSinglyLinkedList_Reverse(t *testing.T) {
	list := SinglyLinkedList{}
	for i := 1; i <= 5; i ++ {
		list.Append(i)
	}
	list.Reverse()
	if list.First.Value != 5 {
		t.Fatalf("%d must equal %d", 5, list.First.Value)
	}
}