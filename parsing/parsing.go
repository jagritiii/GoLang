package parsing

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parseing() {

	fmt.Println("enter a number")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	//var input2 int
	//fmt.Scanf("%d", &input2)

	number, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(number + 1)
	}
}
