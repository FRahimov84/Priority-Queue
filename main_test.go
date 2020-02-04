package main

import (
	"fmt"
	"testing"
)

func Test_queue_len(t *testing.T) {
	q := pQueue{}
	if q.len() != 0 {
		want := 0
		got := q.len()
		t.Errorf("method len() is not correct, for empty pQueue got: %v want: %v", got, want)
	}
}

func Test_queue_first(t *testing.T) {
	q := pQueue{}
	if got := q.first(); got != nil {
		t.Errorf("method first() is not correct, for empty pQueue got: %v want: %v", got, nil)
	}
	want := 1
	q.equeue(want, 0)
	if got := q.first().value; got != want {
		t.Errorf("method first() is not correct, for pQueue with one pElement got: %v want: %v", got, want)
	}
}

func Test_queue_last(t *testing.T) {
	q := pQueue{}
	if got := q.last(); got != nil {
		t.Errorf("method last() is not correct, for empty pQueue got: %v want: %v", got, nil)
	}
	q.equeue(1, 0)
	q.equeue(2, 0)
	q.equeue(3, 0)
	if got, want := q.last().value, 3; got != want {
		t.Errorf("method last() is not correct, for pQueue with three pElement got: %v want: %v", got, want)
	}
	if got, want := q.len(), 3; got != want {
		t.Errorf("method len() is not correct, for pQueue with three pElement got: %v want: %v", got, want)
	}

	if got, want := q.first().value, 1; got != want {
		t.Errorf("method first() is not correct, for pQueue with three pElement got: %v want: %v", got, want)
	}
}

func Test_queue_equeue(t *testing.T) {
	q := pQueue{}
	q.equeue("first", 0)
	if got, want := q.last().value, "first"; got != want {
		t.Errorf("method last() is not correct, for pQueue with one pElement got: %v want: %v", got, want)
	}
	if got, want := q.len(), 1; got != want {
		t.Errorf("method len() is not correct, for pQueue with one pElement got: %v want: %v", got, want)
	}
	if got, want := q.first().value, "first"; got != want {
		t.Errorf("method first() is not correct, for pQueue with one pElement got: %v want: %v", got, want)
	}
	q.equeue("second", 0)
	q.equeue("third", 0)
	q.equeue("fourth", 0)

	if got, want := q.last().value, "fourth"; got != want {
		t.Errorf("method last() is not correct, for pQueue with elements got: %v want: %v", got, want)
	}
	if got, want := q.len(), 4; got != want {
		t.Errorf("method len() is not correct, for pQueue with elements got: %v want: %v", got, want)
	}

	if got, want := q.first().value, "first"; got != want {
		t.Errorf("method first() is not correct, for pQueue with elements got: %v want: %v", got, want)
	}
}

func Test_queue_dequeue(t *testing.T) {
	q := pQueue{}
	if got := q.dequeue(); got != nil {
		t.Errorf("method dequeue() is not correct, for empty pQueue got: %v want: %v", got, nil)
	}
	q.equeue("first", 0)
	q.equeue("second", 0)
	q.equeue("third", 0)
	q.equeue("fourth", 0)
	if got, want := q.len(), 4; got != want {
		t.Errorf("method len() is not correct, for pQueue with elements got: %v want: %v", got, want)
	}
	if got, want := q.dequeue(), "first"; got != want {
		t.Errorf("method dequeue() is not correct, for pQueue with elements got: %v want: %v", got, want)
	}
	if got, want := q.len(), 3; got != want {
		t.Errorf("method len() is not correct, for pQueue with elements got: %v want: %v", got, want)
	}
	if got, want := q.first().value, "second"; got != want {
		t.Errorf("method first() is not correct, for pQueue with elements got: %v want: %v", got, want)
	}
	if got, want := q.last().value, "fourth"; got != want {
		t.Errorf("method last() is not correct, for pQueue with elements got: %v want: %v", got, want)
	}
	q.dequeue()
	q.dequeue()
	q.dequeue()
	if got, want := q.len(), 0; got != want {
		t.Errorf("method len() is not correct, for empty pQueue got: %v want: %v", got, want)
	}
	if got := q.first(); got != nil {
		t.Errorf("method first() is not correct, for pQueue with elements got: %v want: %v", got, nil)
	}
	if got := q.last(); got != nil {
		t.Errorf("method last() is not correct, for pQueue with elements got: %v want: %v", got, nil)
	}
}

func Test_pQueue_makeQueueOfPriotiry(t *testing.T) {
	q := pQueue{}
	if got := q.dequeue(); got != nil {
		t.Errorf("method dequeue() is not correct, for empty pQueue got: %v want: %v", got, nil)
	}
	q.equeue("first", 50)
	q.equeue("second", 20)
	q.equeue("third", 60)
	q.equeue("fourth", 0)
	if got, want := q.len(), 4; got != want {
		t.Errorf("method len() is not correct, for pQueue with elements got: %v want: %v", got, want)
	}
	if got, want := q.dequeue(), "fourth"; got != want {
		t.Errorf("method dequeue() is not correct, for pQueue with elements got: %v want: %v", got, want)
	}
	if got, want := q.len(), 3; got != want {
		t.Errorf("method len() is not correct, for pQueue with elements got: %v want: %v", got, want)
	}
	if got, want := q.first().value, "second"; got != want {
		t.Errorf("method first() is not correct, for pQueue with elements got: %v want: %v", got, want)
	}
	if got, want := q.last().value, "third"; got != want {
		t.Errorf("method last() is not correct, for pQueue with elements got: %v want: %v", got, want)
	}
	q.dequeue()
	q.dequeue()
	q.dequeue()
	if got, want := q.len(), 0; got != want {
		t.Errorf("method len() is not correct, for empty pQueue got: %v want: %v", got, want)
	}
	if got := q.first(); got != nil {
		t.Errorf("method first() is not correct, for pQueue with elements got: %v want: %v", got, nil)
	}
	if got := q.last(); got != nil {
		t.Errorf("method last() is not correct, for pQueue with elements got: %v want: %v", got, nil)
	}
}

func Example_for_PrioritiezeQueue() {
	q := pQueue{}
	q.equeue("first", 50)
	q.equeue("second", 20)
	q.equeue("third", 60)
	q.equeue("fourth", 0)
	i := q.first()
	for i != nil {
		fmt.Println(i.value, i.priority)
		i = i.next
	}
	// Output:fourth 0
	//second 20
	//first 50
	//third 60
}
