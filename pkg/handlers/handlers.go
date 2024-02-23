package handlers

import (
	"fmt"
	"github/janislaus/wogalo/pkg/config"
	"github/janislaus/wogalo/pkg/models"
	"github/janislaus/wogalo/pkg/render"
	"html/template"
	"log"
	"net/http"
)

type Repository struct {
	Config *config.AppConfig
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	session, _ := m.Config.Store.Get(r, "session-name")
	remoteIP := r.RemoteAddr

	session.Values["remote_ip"] = remoteIP
	fmt.Println(remoteIP)
	session.Save(r, w)
	render.RenderTemplate(getTemplateCache(m), w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	session, _ := m.Config.Store.Get(r, "session-name")
	remoteIP, _ := session.Values["remote_ip"].(string)
	fmt.Println(session.Values)

	stringMap := make(map[string]string)
	stringMap["remote_ip"] = remoteIP
	stringMap["test"] = "hello again!"

	render.RenderTemplate(getTemplateCache(m), w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// get template cache from cache or build up cache anew
func getTemplateCache(repo *Repository) map[string]*template.Template {
	if repo.Config.UseCache {
		return repo.Config.TemplateCache
	} else {
		tc, err := render.CreateTemplateCache()

		if err != nil {
			log.Fatal("Cannot create template cache.")
		}
		return tc
	}
}
