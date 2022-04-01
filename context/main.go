package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second * 5))
	defer cancel()
	
	for {
		fmt.Println("Executando %v", ctx)
		time.Sleep(time.Second * 1)
	}
}