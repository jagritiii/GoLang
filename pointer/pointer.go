package pointer

import "fmt"

func Pointer() {
	num := 45
	var p *int
	p = &num

	fmt.Println(*p)
}
