package main

import (
	"projeto-interface/inter"
	"projeto-interface/str"
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
	inter := str.Str{}

	main := NewMain(inter)

	main.inter.Execute()
}
