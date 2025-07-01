package belajar_golang_context

import (
	"context"
	"fmt"
	"testing"
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
