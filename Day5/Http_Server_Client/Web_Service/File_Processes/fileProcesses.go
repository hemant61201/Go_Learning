package File_Processes

import (
	"html/template"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) Save() error {
	filename := page.Title + ".txt"
	return os.WriteFile(filename, page.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

var templates = template.Must(template.ParseFiles("Day5/Http_Server_Client/Web_Service/File_Processes/edit.html", "Day5/Http_Server_Client/Web_Service/File_Processes/view.html"))

func RenderTemplate(writer http.ResponseWriter, template string, page *Page) {
	err := templates.ExecuteTemplate(writer, template+".html", page)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
