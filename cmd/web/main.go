package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adamdboudreau/hello-world-web/pkg/config"
	"github.com/adamdboudreau/hello-world-web/pkg/handlers"
	"github.com/adamdboudreau/hello-world-web/pkg/render"
)

// mac run bunch of go files together "go run *.go"
// windows run bunch of go files together "go run ."
const portNumber = ":3030"

func main() {
	var app config.AppConfig
	app.UseCache = false
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot load templates")
	}
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "hello world !<div>ok</div>")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	})*/
	fmt.Println(fmt.Sprintf("starting app on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
