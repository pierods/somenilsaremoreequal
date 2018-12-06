package main

import (
	"fmt"
)

type Interf interface {
	method1()
}

type InterfImpl struct{}

func (i InterfImpl) method1() {
	fmt.Println("I am method1")
}

type InterfImplPtr struct{}

func (i *InterfImplPtr) method1() {
	fmt.Println("I am method1")
}

func InterfParm(i Interf) {
	i.method1()
}

func panicrecover() {
	if recover() != nil {
		fmt.Println("panic recovered")
	}
}

func main() {
	defer panicrecover()

	var implPtrReceiver *InterfImplPtr

	fmt.Printf("implPtrReceiver is nil: %v\n", implPtrReceiver)

	fmt.Print("Can call method on nil struct* (ptr receiver): ")
	implPtrReceiver.method1()

	var implValReceiver *InterfImpl

	fmt.Print("Cannot call method on nil struct* (value receiver): ")
	func (){
		defer panicrecover()
		implValReceiver.method1()
	}()

	var i Interf

	fmt.Printf("i is nil: %v\n", i)
	fmt.Print("Cannot call method on nil interface:")
	func(){
		defer panicrecover()
		i.method1()
	}()
	fmt.Print("Cannot use nil interface as parameter:")
	func() {
		defer panicrecover()
		InterfParm(i)
	}()


	i = implPtrReceiver
	fmt.Printf("i is nil but assigned to null struct* (ptr receiver): %v\n", i)
	fmt.Print("Can call method on value-assigned nil interface:")
	i.method1()
	fmt.Print("Can use value-assigned nil interface as parameter:")
	InterfParm(i)


	i = implValReceiver
	fmt.Printf("i is nil but assigned to null struct* (val receiver): %v\n", i)
	fmt.Print("Cannot call method on value-assigned nil interface:")
	func() {
		defer panicrecover()
		i.method1()
	} ()
	fmt.Print("Cannot use value-assigned nil interface (val receiver):")
	func() {
		defer panicrecover()
		InterfParm(i)
	} ()

}
