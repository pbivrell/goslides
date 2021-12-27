package main

// START_TYPE OMIT
type Queue[T any] struct {
	items chan []T // non-empty slices only
	empty chan bool   // holds true if the queue is empty
}

func NewQueue[T any]() *Queue[T] {
	items := make(chan []T , 1)
	empty := make(chan bool, 1)
	empty <- true
	return &Queue{items, empty}
}

// END_TYPE OMIT

// START_FUNC OMIT
func (q *Queue[T]) Get() T {
	items := <-q.items
	item := items[0]
	items = items[1:]
	if len(items) == 0 {
		q.empty <- true
	} else {
		q.items <- items
	}
	return item
}

func (q *Queue[T]) Put(item T) {
	var items []T
	select {
	case items = <-q.items:
	case <-q.empty:
	}
	items = append(items, item)
	q.items <- items
}
// END_FUNC OMIT

// START_MAIN OMIT
func main() {
	q1 := NewQueue[string]()
	q1.Put("apple")
	q1.Put("banana")
	q1.Put("kiwi")

	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(q1.Get())
		}()
	}
}
// END_MAIN OMIT
