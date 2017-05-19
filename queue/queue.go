package queue

import "errors"

type Queue struct {
	a []interface{}
}

func New() *Queue {
	q := new(Queue)
	q.a = make([]interface{}, 0)

	return q
}

func (q *Queue) Pop() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	value := q.a[0]
	q.a = q.a[1:]

	return value, nil
}

func (q *Queue) Enqueue(value interface{}) {
	q.a = append(q.a, value)
}

func (q *Queue) Len() int {
	return len(q.a)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
