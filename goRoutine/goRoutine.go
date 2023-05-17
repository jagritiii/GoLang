package goRoutine

import "fmt"

func Msg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, "=", i)
	}
}
func Msg1(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, "=", i)
	}
}
