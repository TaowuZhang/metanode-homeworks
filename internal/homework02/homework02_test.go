package homework02

import (
	"errors"
	"math"
	"reflect"
	"testing"
	"time"
)

func TestPointerFunctions(t *testing.T) {
	n := 5
	AddTen(&n)
	if n != 15 {
		t.Fatalf("got %d", n)
	}
	nums := []int{1, 2, 3}
	DoubleSlice(&nums)
	if !reflect.DeepEqual(nums, []int{2, 4, 6}) {
		t.Fatalf("got %v", nums)
	}
}

func TestOddEvenNumbers(t *testing.T) {
	odds, evens := OddEvenNumbers()
	if !reflect.DeepEqual(odds, []int{1, 3, 5, 7, 9}) || !reflect.DeepEqual(evens, []int{2, 4, 6, 8, 10}) {
		t.Fatalf("got odds=%v evens=%v", odds, evens)
	}
}

func TestRunTasks(t *testing.T) {
	results := RunTasks([]Task{
		func() error { time.Sleep(time.Millisecond); return nil },
		func() error { return errors.New("boom") },
	})
	if len(results) != 2 || results[0].Err != nil || results[1].Err == nil || results[0].Duration <= 0 {
		t.Fatalf("unexpected results: %+v", results)
	}
}

func TestShapes(t *testing.T) {
	rect, err := NewRectangle(3, 4)
	if err != nil || rect.Area() != 12 || rect.Perimeter() != 14 {
		t.Fatalf("bad rectangle")
	}
	circle, err := NewCircle(2)
	if err != nil || math.Abs(circle.Area()-4*math.Pi) > 0.0001 {
		t.Fatalf("bad circle")
	}
	if _, err := NewRectangle(0, 1); err == nil {
		t.Fatalf("expected invalid rectangle error")
	}
}

func TestEmployeePrintInfo(t *testing.T) {
	e := Employee{Person: Person{Name: "Alice", Age: 30}, EmployeeID: "E001"}
	if got := e.PrintInfo(); got != "EmployeeID=E001 Name=Alice Age=30" {
		t.Fatalf("got %q", got)
	}
}

func TestChannels(t *testing.T) {
	got := []int{}
	for n := range GenerateNumbers(3) {
		got = append(got, n)
	}
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("got %v", got)
	}
	if got := BufferedProducerConsumer(5, 2); !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5}) {
		t.Fatalf("got %v", got)
	}
}

func TestCounters(t *testing.T) {
	if got := MutexCounter(10, 1000); got != 10000 {
		t.Fatalf("mutex got %d", got)
	}
	if got := AtomicCounter(10, 1000); got != 10000 {
		t.Fatalf("atomic got %d", got)
	}
}
