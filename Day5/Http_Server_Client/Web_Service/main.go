package main

import (
	"Simple_Programs/Day5/Http_Server_Client/Web_Service/Handlers"
	"Simple_Programs/Day5/Http_Server_Client/Web_Service/Validators_Makers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", Validators_Makers.MakeHandler(Handlers.ViewHandler))
	http.HandleFunc("/edit/", Validators_Makers.MakeHandler(Handlers.EditHandler))
	http.HandleFunc("/save/", Validators_Makers.MakeHandler(Handlers.SaveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
