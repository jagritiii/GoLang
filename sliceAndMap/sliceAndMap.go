package sliceAndMap

import "fmt"

func Slice() {
	var slice1 = []string{"zero", "one", "two", "three", "four"}

	slice1 = append(slice1[:1], slice1[2:]...)
	fmt.Println(slice1)
}

func Maps() {
	myMap := make(map[string]string)

	myMap["one"] = "ONE"
	myMap["two"] = "TWO"
	myMap["three"] = "THREE"
	myMap["four"] = "FOUR"

	delete(myMap, "four")
	fmt.Println(myMap)

	for key, value := range myMap {
		fmt.Printf("key : %v and value : %v\n", key, value)
	}

}
