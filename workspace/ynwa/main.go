package main

import "fmt"

type MyStruct struct{}

func getNilPointer() *MyStruct {
	return nil // Возвращаем nil-указатель
}

func main() {
	var i interface{} = getNilPointer() // i теперь содержит тип *MyStruct, а значение - nil
	fmt.Println(i == nil)               // false, потому что тип не nil!
	fmt.Println(i)                      // <nil>
}
