package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*.htm"))
}

func index(res http.ResponseWriter, req *http.Request) {
	// tpl, err := template.ParseFiles("index.htm")
	// if err != nil {
	// 	log.Fatalln("error in parsing template", err)
	// }

	err := tpl.ExecuteTemplate(res, "index.htm", nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("template"))))
	// http.Handle("/", http.FileServer(http.Dir("./template")))

	http.ListenAndServe(":8080", nil)
}
