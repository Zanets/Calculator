package main

import (
	"flag"
	"Calculator/Parser"
	"Calculator/Collection"
	"fmt"
)


func main() {
	math := flag.String("math", "", "")
	flag.Parse()
	
	p := Parser.Parser{}
	rpn := p.Parse(*math)
	//p.Status()
	fmt.Printf("%s = %f\n", *math, calculate(rpn))	
}

func calculate(t []interface{}) float64 {
	var s Collection.Stack
	for i := 0; i < len(t); i++ {
		// NOTE: -,/ need to do with reverse order
		// of stack pop
		switch o := t[i]; o {
		case '+':
			s.Push(s.Pop().(float64) + s.Pop().(float64))
		case '-':
			a := s.Pop().(float64)
			b := s.Pop().(float64)
			s.Push(b - a)
		case '*':
			s.Push(s.Pop().(float64) * s.Pop().(float64))
		case '/':
			a := s.Pop().(float64)
			b := s.Pop().(float64)
			s.Push(b / a)
		default:
			s.Push(o.(float64))
		}

	} 
	if s.Size() == 1 {
		return s.Pop().(float64)
	} else {
		fmt.Println("error happen")
		return 0
	}

}
