package main

import (
	"fmt"
	"time"
)

func init() {
	var err error
	time.Local, err = time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic(err)
	}
}
func main() {
	fmt.Println("Hello World")
}
