package main

import (
	"assesment3/webserver"
	"html/template"
	"net/http"
)

type Status struct {
	water int
	wind  int
}

func main() {
	webserver.Start()
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	stat := Status{
		water: 4,
		wind:  3,
	}
	tpl, _ := template.ParseFiles("./webserver/index.html")
	tpl.Execute(w, map[string]interface{}{
		"Payload": stat,
		"Title":   "Assesment 3",
	})
}
