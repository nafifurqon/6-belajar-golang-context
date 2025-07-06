package belajar_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println("contextA", contextA)
	fmt.Println("contextB", contextB)
	fmt.Println("contextC", contextC)
	fmt.Println("contextD", contextD)
	fmt.Println("contextE", contextE)
	fmt.Println("contextF", contextF)
	fmt.Println("contextG", contextG)

	fmt.Println("=========================================")

	fmt.Println(contextF.Value("f")) // dapat dari context dirinya sendiri
	fmt.Println(contextF.Value("c")) // dapat dari context parent-nya
	fmt.Println(contextF.Value("b")) // tidak dapat karena beda parent
	fmt.Println(contextA.Value("b")) // tidak dapat karena punya context child-nya. hanya bisa ambil dari context dirinya sendiri dan parent-nya
	fmt.Println(contextG.Value("c")) // dapat dari context parent dari parent-nya
}

func CreateCounter() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			counter++
		}
	}()

	return destination
}

func CreateCounterWithContext(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithoutCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	destination := CreateCounter()
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

func TestContextWithContextCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounterWithContext(ctx)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cancel() // mengirim sinyal cancel ke context

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
