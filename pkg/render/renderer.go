package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders templates
func RenderTemplateFirstTest(w http.ResponseWriter, tmpl string) {
	// read from disc each request
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	// cache template to be returned after
	var tmpl *template.Template
	var err error
	// check if we already have template in cache
	_, inMap := tc[t]
	log.Println("check template: ", t)
	if !inMap {
		log.Println("creating template and adding to cached")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
			//return err
		}
	} else {
		// have template
		log.Println("using cached template")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = tmpl
	return nil
}
