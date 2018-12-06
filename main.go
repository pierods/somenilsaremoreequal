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

func StructParm(impl InterfImpl) {
	impl.method1()
}

func StructParmPtrReceiver(impl *InterfImplPtr){
	impl.method1()
}

func panicrecover() {
	if recover() != nil {
		fmt.Println("panic recovered")
	}
}

func main() {
	var implPtrReceiver *InterfImplPtr

	fmt.Printf("implPtrReceiver is nil: %v\n", implPtrReceiver)
	fmt.Print("Can call method on nil struct* (ptr receiver): ")
	implPtrReceiver.method1()

	fmt.Print("Can use nil ptr as parameter:")
	StructParmPtrReceiver(implPtrReceiver)

	var implValReceiver *InterfImpl

	fmt.Printf("implValReceiver is nil: %v\n", implValReceiver)
	fmt.Print("Cannot call method on nil struct* (value receiver): ")
	func (){
		defer panicrecover()
		implValReceiver.method1()
	}()

	fmt.Print("Cannot use nil ptr as parameter (value receiver)")
	func () {
		defer panicrecover()
		StructParm(*implValReceiver)
	} ()

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
	fmt.Printf("i is assigned to nil struct* (ptr receiver): %v\n", i)
	fmt.Print("Can call method on value-assigned nil interface:")
	i.method1()
	fmt.Print("Can use nil-assigned interface as parameter:")
	InterfParm(i)


	i = implValReceiver
	fmt.Printf("i is assigned to nil struct* (value receiver): %v\n", i)
	fmt.Print("Cannot call method on nil-assigned interface (value receiver):")
	func() {
		defer panicrecover()
		i.method1()
	} ()
	fmt.Print("Cannot use nil-assigned interface (val receiver) as parameter:")
	func() {
		defer panicrecover()
		InterfParm(i)
	} ()

}
