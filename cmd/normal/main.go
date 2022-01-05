package main

import (
	"fmt"
	"reflect"
	"yaegi-issues/common"
)

type A struct{}

func (A) Cheat() {}

func main() {
	var i common.Cheater = (*A)(nil)

	fmt.Println("Using interface natively:")
	fmt.Printf("reflect.TypeOf(i): %s\n", reflect.TypeOf(i))
	fmt.Println("")
}
