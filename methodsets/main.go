package main

import (
	"fmt"
)

type Interface interface {
	method()
	method1()
}

type Implementation struct{}

func (i Implementation) method() {
	fmt.Println("method with value receiver called")
}

func (i *Implementation) method1() {
	fmt.Println("method with pointer receiver called")
}

func panicRecover() {
	if recover() != nil {
		fmt.Println("panic recovered")
	}
}

func main() {

	var v Implementation

	v.method()
	v.method1()

	var p *Implementation

	func() {
		defer panicRecover()
		p.method()
	}()
	p.method1()
	p2 := &Implementation{}
	p2.method()

	var i Interface

	// i = v // does not compile: cannot use v (type Implementation) as type Interface in assignment: Implementation does not implement Interface (method1 method has pointer receiver)

	i = p

	i.method1()
	func() {
		defer panicRecover()
		i.method()
	}()

	i = &Implementation{}
	i.method()
}

