package errorHandling

import "errors"

type dividebyzero struct {
	texterror string
}

func (c dividebyzero) Error() string {
	return c.texterror
}

//	func dividedbyzero() error {
//		return dividedbyzero("your error is ")
//	}
func Divison(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return -1, errors.New("division by zero")
		//can add the custom type also
	}
	return num1 / num2, nil
}
