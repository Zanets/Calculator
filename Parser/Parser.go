package Parser

import (
	"Calculator/Collection"
	"fmt"
)

// Parse infix to postfix(Reverse Polish Notation)
// with Shunting Yard Algorithm
// https://en.wikipedia.org/wiki/Shunting-yard_algorithm

type Parser struct {
	rpn Collection.Queue
	ops Collection.Stack
}

func (p *Parser) Parse(s string) []interface{} {
	p.ops.Clear()
	p.rpn.Clear()
	return p.shunting_yard([]rune(s))
}

func (p *Parser) shunting_yard(s []rune) []interface{} {
	for i := 0; i < len(s); i++ {
		if t := s[i]; p.isNum(t) {
			ti := float64(t-'0')
			// if current s is number and also
			// next one, then its the same number
			for j := i + 1; j < len(s); j++ {
				if p.isNum(s[j]) {
					ti = ti * 10 + float64(s[j]-'0') 
					i = j
				} else {
					break
				}
 			}
			p.rpn.Push(ti)
		} else if t == '(' {
			p.ops.Push(t)
		} else if t == ')' {
			leftPare := false
			for !p.ops.IsEmpty() {
				op := p.ops.Peek()
				if p.isOp(op) {
					p.rpn.Push(p.ops.Pop())
				} else if op == '(' {
					leftPare = true
					p.ops.Pop()
					break
				} else {
					// TODO: ERROR
				}
			}
			// mismatched parentheses
			if !leftPare {
				// TODO: ERROR
			}
		} else if p.isOp(t) {
			// TODO: right association is ignored
			// such as ++,--,>,<,=

			// keep pop op in ops which priority
			// is equal to or larger than t
			for !p.ops.IsEmpty() {
				op := p.ops.Peek()
				if !p.isOp(op) {
					break
				}
				if p.gp(t) <= p.gp(op) {
					p.rpn.Push( p.ops.Pop() )
				} else {
					break
				}  
			}
			p.ops.Push(t)
		}
 	}

	// pop all op in ops to rpn,
	// if '(' exist, then there are mismatched parentheses
 	for !p.ops.IsEmpty() {
		switch op := p.ops.Pop(); op {
		case '+', '-', '*', '/':
			p.rpn.Push(op)
		default:
			// TODO: ERROR
		}
	}

	return p.rpn.Dup()
}

func (p *Parser) isNum(t rune) bool {
	ts := []rune{'0','1','2','3','4','5','6','7','8','9'}
	for i := 0; i < len(ts); i++ {
		if t == ts[i] {
			return true
		}
	}
	return false
}


func (p *Parser) isOp(t interface{}) bool {
	ts := []rune{'+','-','*','/'}
	for i := 0; i < len(ts); i++ {
		if t.(rune) == ts[i] {
			return true
		}
	}
	return false
}

// the greater, the higher priority
func (p *Parser) gp(t interface{}) int {
	switch t.(rune) {
	case '+', '-':
		return 0
	case '*', '/':
		return 1
	default:
		// TODO: ERROR
		return 0
	}
}

func (p *Parser) Status() {
	fmt.Println("Status")
	
	ops := p.ops.Dup()
	fmt.Println("	OPS Stack")
	for i := 0; i < len(ops); i++ {
		op := ops[i]
		switch op.(type) {
		case float64:
			fmt.Printf("		%f\n", op.(float64))
		case rune:
			fmt.Printf("		%s\n", string(op.(rune)))
		}
	}
	
	rpn := p.rpn.Dup()
	fmt.Println("	RPN Queue")
	for i := 0; i < len(rpn); i++ {
		op := rpn[i]
		switch op.(type) {
		case float64:
			fmt.Printf("		%f\n", op.(float64))
		case rune:
			fmt.Printf("		%s\n", string(op.(rune)))
		}
	}

	fmt.Println()
}
