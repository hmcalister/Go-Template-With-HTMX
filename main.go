package main

import (
	"embed"
	"hmcalister/htmxTest/api"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

var (
	//go:embed static/css/output.css
	embedCSSFile []byte

	//go:embed static/htmx/htmx.js
	embedHTMXFile []byte

	//go:embed static/templates/*.html
	templatesFS embed.FS
)

func main() {
	var err error
	applicationState := api.NewApplicationState()

	templatesFS, err := fs.Sub(templatesFS, "static/templates")
	if err != nil {
		log.Fatalf("error during embedded file system: %v", err)
	}
	indexTemplate, err := template.ParseFS(templatesFS, "index.html")
	if err != nil {
		log.Fatalf("error parsing template: %v", err)
	}

	http.HandleFunc("/css/output.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		w.Write(embedCSSFile)
	})

	http.HandleFunc("/htmx/htmx.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(embedHTMXFile)
	})

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		err = indexTemplate.Execute(w, PageData{
			DataStructs: exampleDataArr,
		})
		if err != nil {
			log.Fatalf("error during template execute: %v", err)
		}
	})

	log.Printf("Serving template at http://localhost:8080/index")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("error during http serving: %v", err)
	}
}
