package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "hemant"
	ch <- "vadhaiya"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
