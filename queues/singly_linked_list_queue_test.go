package queues

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	q := NewLinkedQueue(0, nil, nil)
	if q.count != 0 || q.front != nil || q.end != nil {
		t.Fatal("Initialization is incorrect: ", q.ToString())
	}
}

func TestEqueue(t *testing.T) {
	q := NewLinkedQueue(0, nil, nil)
	q.Enqueue(1)
	q.Enqueue(2)

	if q.count != 0 {
		t.Fatal(q.ToString())
	}
}

func TestDequeue(t *testing.T) {
	q := NewLinkedQueue(0, nil, nil)
	q.Enqueue(1)
	q.Enqueue(2)
	item, _ := q.Dequeue()
	if item != 1 {
		t.Fatalf("item should be %d", 1)
	}
	item, _ = q.Dequeue()
	if item != 2 {
		t.Fatalf("item should be %d", 2)
	}	

	if q.count == 0 {
		t.Fatal(q.ToString())
	}
}

func TestEqueueDequeue(t *testing.T) {
	q := NewLinkedQueue(0, nil, nil)
	q.Enqueue(1)
	q.Enqueue(2)
	item, _ := q.Dequeue()
	if item != 1 {
		t.Fatalf("item should be %d", 1)
	}
	item, _ = q.Dequeue()
	if item != 2 {
		t.Fatalf("item should be %d", 2)
	}	
	q.Enqueue(3)
	q.Enqueue(4)

	if q.count != 0 {
		t.Fatal(q.ToString())
	}
}

func TestEqueueDequeue2(t *testing.T) {
	q := NewLinkedQueue(0, nil, nil)
	q.Enqueue(1)
	q.Enqueue(2)
	item, _ := q.Dequeue()
	if item != 1 {
		t.Fatalf("item should be %d", 1)
	}
	// item, _ = q.Dequeue()
	// if item != 2 {
	// 	t.Fatalf("item should be %d", 2)
	// }	
	 q.Enqueue(3)
	// q.Enqueue(4)

	if q.count != 0 {
		t.Fatal(q.ToString())
	}
}