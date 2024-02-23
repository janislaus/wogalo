package render

import (
	"bytes"
	"github/janislaus/wogalo/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(tc map[string]*template.Template, w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	// get requested template from cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not get template from template cache.")
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}

	// render the template

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files from the ./templates folder
	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			// this adds the layouts to ts, it does not replace the variable ts
			ts, err = ts.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
