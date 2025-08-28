package Validators_Makers

import (
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func MakeHandler(function func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(writer http.ResponseWriter, reader *http.Request) {
		match := validPath.FindStringSubmatch(reader.URL.Path)
		if match == nil {
			http.NotFound(writer, reader)
			return
		}
		function(writer, reader, match[2])
	}
}
