package main

import (
	"awesomeProject2/goRoutine"
	"time"
)

func main() {
	//parsing.Parseing()
	//time.Time()
	//time.Pointers()
	//time.Arrays()
	//sliceAndMap.Slice()
	//sliceAndMap.Maps()
	//_interface.Interface()
	//pointer.Pointer()

	//error handling--------
	//p, ok := errorHandling.Divison(20, 0)
	//if ok != nil {
	//	fmt.Println(ok.Error())
	//} else {
	//	fmt.Println(p, ok)
	//}

	//go routine------------
	go goRoutine.Msg("hi")
	go goRoutine.Msg1("byi")
	time.Sleep(time.Second)

	//for encapsulation--------------------
	//var firstname string
	//var lastname string
	//var age int
	//fmt.Scanln(&firstname)
	//fmt.Scanln(&lastname)
	//fmt.Scanln(&age)
	//p := oops.PersonCreate(firstname, lastname, age)
	//fmt.Println(p.Getdetails())

	//for polymorphism-----------------
	// use the common interface and in the right hand side use any struct which is implementing all its
	//functions
	//var c polymorphism.Shape = polymorphism.Circle{}
	//var s polymorphism.Shape = polymorphism.Square{}
	//c.Render()
	//s.Render()

	//for composition-----------
	//c := composition.NewCar("a", 600, 400)
	//fmt.Println(c)
}
