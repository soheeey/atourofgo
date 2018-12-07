package main

import "fmt"

type MyError struct {
	Message string
	ErrCode int
}

/* errorインターフェースのメソッドを実装 */
func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "エラーが発生しました", ErrCode: 1234}
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}
}
