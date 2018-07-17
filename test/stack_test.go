package test

import (
	"Calculator/Collection"
	"fmt"
)

func main() {
	t := Collection.Stack{}
	t.Push(5)
	t.Push(6)
	t.Push(7)
	t.Push(8)
	t.Push(9)

	fmt.Println(t.Peek())
	fmt.Println(t.Pop())
	fmt.Println(t.Pop())
	fmt.Println(t.Pop())
	fmt.Println(t.Pop())
	fmt.Println(t.Pop())
	fmt.Println(t.Pop())
	fmt.Println(t.Size())
}
