package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/adamdboudreau/hello-world-web/pkg/config"
	"github.com/adamdboudreau/hello-world-web/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets config for template
func NewTemplates(a *config.AppConfig) {
	app = a
}
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// get template cache from app config

	// create template cache
	/*tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	} */
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from cache")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	// render template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("error writing template to browser", err)
	}

	// read from disc each request
	/* parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	} */
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files with page.tmpl
	for _, page := range pages {
		name := filepath.Base(page) // file.page.tmpl
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if len(matches) > 0 {
			// add any templates to the parsed page
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

/* initial template cache setup
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
*/
