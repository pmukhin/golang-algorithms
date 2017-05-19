package queue

import "testing"

func TestQueue_IsEmpty(t *testing.T) {
	queue := New()
	if !queue.IsEmpty() {
		t.Error("Queue is empty")
	}
	queue.Enqueue(5)
	if queue.IsEmpty() {
		t.Error("Queue is not really empty")
	}
}

func TestQueue_Len(t *testing.T) {
	queue := New()
	if queue.Len() != 0 {
		t.Error("Queue is empty")
	}

	for i := 0; i < 20; i++ {
		queue.Enqueue(i)
		if queue.Len() != (i + 1) {
			t.Errorf("Len() must return %d", (i + 1))
		}
	}
}

func TestQueue_Enqueue(t *testing.T) {
	queue := New()
	if _, err := queue.Pop(); err == nil {
		t.Fatal("Queue is empty, there must be an error")
	}

	for i := 0; i < 20; i++ {
		queue.Enqueue(i)
	}

	for i := 0; i < 20; i++ {
		v, err := queue.Pop()
		if err != nil {
			t.Error("Unexptected error")
		}
		if v != i {
			t.Errorf("%d must equal to %d", v, i)
		}
	}
}

func TestQueue_Pop(t *testing.T) {
	queue := New()
	for i := 0; i < 20; i++ {
		queue.Enqueue(i)
	}
	if queue.Len() != 20 {
		t.Error("Wrong length returned: %d", queue.Len())
	}
	currLen := 20
	for i := 0; i < 20; i++ {
		v, _ := queue.Pop()
		if v != i {
			t.Errorf("%d must equal to %d", v, i)
		}
		currLen--
		if queue.Len() != currLen {
			t.Errorf("Unexpected length: %d", queue.Len())
		}
	}
}
