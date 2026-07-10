package main

import (
	"fmt"

	"github.com/TaowuZhang/metanode-go-homeworks/internal/homework02"
)

func main() {
	n := 5
	homework02.AddTen(&n)
	fmt.Println("add ten:", n)
	nums := []int{1, 2, 3}
	homework02.DoubleSlice(&nums)
	fmt.Println("double slice:", nums)
	odds, evens := homework02.OddEvenNumbers()
	fmt.Println("odds:", odds, "evens:", evens)
	results := homework02.RunTasks([]homework02.Task{func() error { return nil }})
	fmt.Println("task results:", results)
	rect, _ := homework02.NewRectangle(3, 4)
	circle, _ := homework02.NewCircle(2)
	fmt.Println("rectangle:", rect.Area(), rect.Perimeter(), "circle:", circle.Area(), circle.Perimeter())
	employee := homework02.Employee{Person: homework02.Person{Name: "Alice", Age: 30}, EmployeeID: "E001"}
	fmt.Println(employee.PrintInfo())
	for v := range homework02.GenerateNumbers(3) {
		fmt.Println("channel:", v)
	}
	fmt.Println("buffered:", homework02.BufferedProducerConsumer(5, 2))
	fmt.Println("mutex counter:", homework02.MutexCounter(10, 1000))
	fmt.Println("atomic counter:", homework02.AtomicCounter(10, 1000))
}
