package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "hemant"
	ch <- "aakash"
	ch <- "shekhar"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
