package Handlers

import (
	"Simple_Programs/Day5/Http_Server_Client/Web_Service/File_Processes"
	"net/http"
)

func ViewHandler(writer http.ResponseWriter, reader *http.Request, title string) {
	page, err := File_Processes.LoadPage(title)
	if err != nil {
		http.Redirect(writer, reader, "/edit/"+title, http.StatusFound)
		return
	}
	File_Processes.RenderTemplate(writer, "view", page)
}

func EditHandler(writer http.ResponseWriter, reader *http.Request, title string) {
	page, err := File_Processes.LoadPage(title)
	if err != nil {
		page = &File_Processes.Page{Title: title}
	}
	File_Processes.RenderTemplate(writer, "edit", page)
}

func SaveHandler(writer http.ResponseWriter, reader *http.Request, title string) {
	body := reader.FormValue("body")
	page := &File_Processes.Page{Title: title, Body: []byte(body)}
	err := page.Save()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(writer, reader, "/view/"+title, http.StatusFound)
}
