package queues

import(
	"testing"
)

func TestNewArrayQueue(t *testing.T) {
	q := NewArrayQueue();

	if q.count != 0 || q.front != 0 || len(q.items) != 0 {
		t.Fatalf("count: %v front: %v len: %v Q:%#v \n", q.count, q.front, len(q.items), q.items)
	}
}

func TestEnqueue(t *testing.T) {
	q := NewArrayQueue();
	q.Enqueue(5)
	q.Enqueue(4)
	q.Dequeue()
	q.Dequeue()//
	q.Enqueue(3)
	q.Enqueue(2)
	q.Enqueue(1)
	q.Enqueue(0)
	q.Enqueue(-1)
	q.Enqueue(-2)
	if q.count != 0 || q.front != 0 || len(q.items) != 0 {
		t.Fatalf("%s \n", q.ToString())
	}
}

func TestEnqueue2(t *testing.T) {
	q := NewArrayQueue();
	q.Enqueue(5)
	q.Enqueue(4)
	q.Dequeue()
	q.Dequeue()//
	q.Enqueue(3)
	q.Enqueue(2)
	q.Enqueue(1)
	q.Enqueue(0)
	q.Enqueue(-1)
	q.Enqueue(-2)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()	
	if q.count != 0 || q.front != 0 || len(q.items) != 0 {
		t.Fatalf("%s \n", q.ToString())
	}
}

func TestEnqueue3(t *testing.T) {
	q := NewArrayQueue();
	q.Enqueue(5)
	q.Enqueue(4)
	q.Dequeue()
	q.Dequeue()//
	q.Enqueue(3)
	q.Enqueue(2)
	q.Enqueue(1)
	q.Enqueue(0)
	q.Enqueue(-1)
	q.Enqueue(-2)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Enqueue(-3)
	q.Enqueue(-4)
	q.Enqueue(-5)
	q.Enqueue(-6)	
	if q.count != 0 || q.front != 0 || len(q.items) != 0 {
		t.Fatalf("%s \n", q.ToString())
	}
}

func TestEnqueue4(t *testing.T) {
	q := NewArrayQueue();
	q.Enqueue(5)
	q.Enqueue(4)
	q.Enqueue(3)
	q.Dequeue()
	q.Dequeue()//
	q.Dequeue()

	q.Enqueue(2)
	q.Enqueue(1)
	q.Enqueue(0)
	q.Enqueue(-1)
	q.Enqueue(-2)
	q.Enqueue(-3)

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	q.Enqueue(-4)
	q.Enqueue(-5)
	q.Enqueue(-6)	
	if q.count != 0 || q.front != 0 || len(q.items) != 0 {
		t.Fatalf("%s \n", q.ToString())
	}
}