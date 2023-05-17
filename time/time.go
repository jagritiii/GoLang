package time

import (
	"fmt"
	"time"
)

func Time() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	format1 := presentTime.Format("01-02-2006")
	fmt.Println(format1)

	format2 := presentTime.Format("01-02-2006 15:05:04 Monday")
	fmt.Println(format2)

	format3 := presentTime.Format("01-02-2006 Monday")
	fmt.Println(format3)

}

func Pointers() {
	var ptr *int
	var a int = 5
	ptr = &a
	fmt.Println(ptr, *ptr)
}

func Arrays() {
	//var arr [2]int
	//arr[0] = 7

	var arr = []string{"ff", "ss"}
	fmt.Println(arr)
}
