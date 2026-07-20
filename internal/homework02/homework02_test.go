package homework02

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"sync"
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
	const taskCount = 4
	started := make(chan int, taskCount)
	completed := make(chan int, taskCount)
	release := make(chan struct{})
	var releaseOnce sync.Once
	releaseTasks := func() { releaseOnce.Do(func() { close(release) }) }
	defer releaseTasks()

	errOne := errors.New("task one failed")
	errThree := fmt.Errorf("task three failed: %w", errors.New("underlying failure"))
	wantErrors := []error{nil, errOne, nil, errThree}
	tasks := make([]Task, taskCount)
	for id := range tasks {
		id := id
		tasks[id] = func() error {
			started <- id
			<-release
			completed <- id
			return wantErrors[id]
		}
	}

	resultsCh := make(chan []TaskResult, 1)
	go func() { resultsCh <- RunTasks(tasks) }()

	seenStarted := make(map[int]bool, taskCount)
	startTimeout := time.NewTimer(time.Second)
	defer startTimeout.Stop()
	for len(seenStarted) < taskCount {
		select {
		case id := <-started:
			seenStarted[id] = true
		case <-startTimeout.C:
			t.Fatalf("only %d/%d tasks started before release; tasks are not concurrent", len(seenStarted), taskCount)
		}
	}
	releaseTasks()

	var results []TaskResult
	select {
	case results = <-resultsCh:
	case <-time.After(time.Second):
		t.Fatal("RunTasks did not return after all tasks were released")
	}

	seenCompleted := make(map[int]bool, taskCount)
	for range taskCount {
		select {
		case id := <-completed:
			seenCompleted[id] = true
		case <-time.After(time.Second):
			t.Fatalf("only %d/%d tasks completed", len(seenCompleted), taskCount)
		}
	}
	if len(seenCompleted) != taskCount {
		t.Fatalf("completed task IDs = %v, want all %d tasks", seenCompleted, taskCount)
	}
	if len(results) != taskCount {
		t.Fatalf("got %d results, want %d", len(results), taskCount)
	}
	for id, result := range results {
		if result.ID != id {
			t.Errorf("results[%d].ID = %d, want %d", id, result.ID, id)
		}
		if result.Duration < 0 {
			t.Errorf("results[%d].Duration = %v, want non-negative", id, result.Duration)
		}
		if wantErrors[id] == nil {
			if result.Err != nil {
				t.Errorf("results[%d].Err = %v, want nil", id, result.Err)
			}
		} else if result.Err != wantErrors[id] || !errors.Is(result.Err, wantErrors[id]) {
			t.Errorf("results[%d].Err = %v, want preserved error %v", id, result.Err, wantErrors[id])
		}
	}
}

func TestShapes(t *testing.T) {
	const tolerance = 1e-9
	almostEqual := func(got, want float64) bool {
		return math.Abs(got-want) <= tolerance
	}

	rect, err := NewRectangle(3, 4)
	if err != nil {
		t.Fatalf("NewRectangle() error = %v", err)
	}
	var rectShape Shape = rect
	if !almostEqual(rectShape.Area(), 12) || !almostEqual(rectShape.Perimeter(), 14) {
		t.Fatalf("rectangle through Shape: area = %v, perimeter = %v", rectShape.Area(), rectShape.Perimeter())
	}

	circle, err := NewCircle(2)
	if err != nil {
		t.Fatalf("NewCircle() error = %v", err)
	}
	var circleShape Shape = circle
	if !almostEqual(circleShape.Area(), 4*math.Pi) || !almostEqual(circleShape.Perimeter(), 4*math.Pi) {
		t.Fatalf("circle through Shape: area = %v, perimeter = %v", circleShape.Area(), circleShape.Perimeter())
	}

	if _, err := NewRectangle(0, 1); err == nil {
		t.Fatal("NewRectangle() error = nil for invalid dimensions")
	}
	for _, radius := range []float64{0, -1} {
		if _, err := NewCircle(radius); err == nil {
			t.Errorf("NewCircle(%v) error = nil, want invalid radius error", radius)
		}
	}
}

func TestEmployeePrintInfo(t *testing.T) {
	e := Employee{Person: Person{Name: "Alice", Age: 30}, EmployeeID: "E001"}
	if got := e.PrintInfo(); got != "EmployeeID=E001 Name=Alice Age=30" {
		t.Fatalf("got %q", got)
	}
}

func TestChannels(t *testing.T) {
	wantNumbers := make([]int, 10)
	for i := range wantNumbers {
		wantNumbers[i] = i + 1
	}

	generatedCh := make(chan []int, 1)
	go func() {
		var got []int
		for n := range GenerateNumbers(10) {
			got = append(got, n)
		}
		generatedCh <- got
	}()
	select {
	case got := <-generatedCh:
		if !reflect.DeepEqual(got, wantNumbers) {
			t.Fatalf("GenerateNumbers() = %v, want %v", got, wantNumbers)
		}
	case <-time.After(time.Second):
		t.Fatal("ranging over GenerateNumbers did not finish after production")
	}

	wantBuffered := make([]int, 100)
	for i := range wantBuffered {
		wantBuffered[i] = i + 1
	}
	bufferedCh := make(chan []int, 1)
	go func() { bufferedCh <- BufferedProducerConsumer(100, 10) }()
	select {
	case got := <-bufferedCh:
		if !reflect.DeepEqual(got, wantBuffered) {
			t.Fatalf("BufferedProducerConsumer() = %v, want %v", got, wantBuffered)
		}
	case <-time.After(time.Second):
		t.Fatal("BufferedProducerConsumer did not finish")
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
