package str

import "fmt"

type Str struct{}

func NewStr() *Str {
	return &Str{}
}

func (s Str) Execute() error {

	fmt.Println("FUI EESCUTAD")
	
	return nil
}
