package main

import (
	"Simple_Programs/Day4/Hands_On_Lab/loading"
	"Simple_Programs/Day4/Hands_On_Lab/parsing"
	"fmt"
)

func main() {

	logs := loading.LoadLogs{FileName: "Day4/Hands_On_Lab/large_log.txt"}

	fmt.Println("Execution is Started....")
	//
	//loading.WaitingGroup.Add(1)

	go logs.Load()
	//
	//loading.WaitingGroup.Wait()
	//
	//parsing.WaitingGroup.Add(1)

	parsing.Parse()

	//parsing.WaitingGroup.Wait()
}
