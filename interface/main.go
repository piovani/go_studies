package main

import (
	"fmt"
	"projeto-interface/inter"
	"projeto-interface/str"
	"reflect"
)

type Main struct {
	inter inter.Inter
}

func NewMain(
	inter inter.Inter,
) *Main {
	return &Main{
		inter: inter,
	}
}

func main() {
	inter := str.NewStr()

	fmt.Println(reflect.TypeOf(inter))

	main := NewMain(inter)

	main.inter.Execute()
}
