package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%v\n", time.Duration(75) * time.Millisecond)
}