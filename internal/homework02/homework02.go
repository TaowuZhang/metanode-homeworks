package homework02

import (
	"errors"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

func AddTen(n *int) {
	if n != nil {
		*n += 10
	}
}

func DoubleSlice(nums *[]int) {
	if nums == nil {
		return
	}
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

func OddEvenNumbers() (odds []int, evens []int) {
	var wg sync.WaitGroup
	oddCh := make(chan []int, 1)
	evenCh := make(chan []int, 1)
	wg.Add(2)
	go func() {
		defer wg.Done()
		out := []int{}
		for i := 1; i <= 10; i += 2 {
			out = append(out, i)
		}
		oddCh <- out
	}()
	go func() {
		defer wg.Done()
		out := []int{}
		for i := 2; i <= 10; i += 2 {
			out = append(out, i)
		}
		evenCh <- out
	}()
	wg.Wait()
	return <-oddCh, <-evenCh
}

type Task func() error

type TaskResult struct {
	ID       int
	Duration time.Duration
	Err      error
}

func RunTasks(tasks []Task) []TaskResult {
	var wg sync.WaitGroup
	results := make([]TaskResult, len(tasks))
	for i, task := range tasks {
		wg.Add(1)
		go func(id int, fn Task) {
			defer wg.Done()
			start := time.Now()
			err := fn()
			results[id] = TaskResult{ID: id, Duration: time.Since(start), Err: err}
		}(i, task)
	}
	wg.Wait()
	return results
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func NewRectangle(width, height float64) (Rectangle, error) {
	if width <= 0 || height <= 0 {
		return Rectangle{}, errors.New("width and height must be positive")
	}
	return Rectangle{Width: width, Height: height}, nil
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Circle struct {
	Radius float64
}

func NewCircle(radius float64) (Circle, error) {
	if radius <= 0 {
		return Circle{}, errors.New("radius must be positive")
	}
	return Circle{Radius: radius}, nil
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() string {
	return fmt.Sprintf("EmployeeID=%s Name=%s Age=%d", e.EmployeeID, e.Name, e.Age)
}

func GenerateNumbers(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= n; i++ {
			ch <- i
		}
	}()
	return ch
}

func BufferedProducerConsumer(total int, bufferSize int) []int {
	if bufferSize <= 0 {
		bufferSize = 1
	}
	ch := make(chan int, bufferSize)
	var wg sync.WaitGroup
	collected := make([]int, 0, total)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			collected = append(collected, v)
		}
	}()
	for i := 1; i <= total; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
	return collected
}

func MutexCounter(workers int, increments int) int {
	var mu sync.Mutex
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	return counter
}

func AtomicCounter(workers int, increments int) int64 {
	var counter int64
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	return counter
}
