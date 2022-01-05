package main

import (
	"github.com/traefik/yaegi/interp"
	"yaegi-issues/cmd/yaegirun/symbols"
)

const FromSymbolCode = `
import (
	"fmt"
	"reflect"
	"yaegi-issues/common"
)

type A struct {}

func (A) Cheat() {}

func main() {
	var i common.Cheater = (*A)(nil)

	fmt.Println("Using interface defined outside:")
	fmt.Printf("reflect.TypeOf(i): %v\n", reflect.TypeOf(i))
	fmt.Println("")
}
`

const NotFromSymbolCode = `
import (
	"fmt"
	"reflect"
)

type Cheater interface {
	Cheat()
}

type A struct {}

func (A) Cheat() {
    fmt.Println("A cheat!")
}

func main() {
	var i Cheater = (*A)(nil)

	fmt.Println("Using interface defined inside:")
	fmt.Printf("reflect.TypeOf(i): %s\n", reflect.TypeOf(i))

    a, ok := i.(*A)
	fmt.Printf("Type assertion: %v, %v\n", a, ok)

	fmt.Println("")
}
`

func main() {
	// First code
	i := interp.New(interp.Options{})
	err := i.Use(symbols.Symbols)
	if err != nil {
		panic(err)
	}

	_, err = i.Eval(FromSymbolCode)
	if err != nil {
		panic(err)
	}

	// Second code
	i2 := interp.New(interp.Options{})
	err = i2.Use(symbols.Symbols)
	if err != nil {
		panic(err)
	}

	_, err = i2.Eval(NotFromSymbolCode)
	if err != nil {
		panic(err)
	}
}
