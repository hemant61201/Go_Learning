package main

import (
	"Simple_Programs/Day4/Hands_On_Lab/loading"
	"Simple_Programs/Day4/Hands_On_Lab/parsing"
	"fmt"
)

func main() {

	logs := loading.LoadLogs{FileName: "Day4/Hands_On_Lab/large_log.txt"}

	fmt.Println("Execution is Started....")

	go logs.Load()

	parsing.Parse()
}
