package _interface

import "fmt"

type calculator struct {
	num1, num2 int
}

type Calculation interface {
	sum() int
	mul() int
}

func (c calculator) sum() int {
	return c.num1 + c.num2
}

func (c calculator) mul() int {
	return c.num1 * c.num2
}

func useCalculation(cc Calculation) {
	fmt.Println(cc.sum(), cc.mul())
}

func Interface() {
	c := calculator{1, 2}
	useCalculation(c)

}
