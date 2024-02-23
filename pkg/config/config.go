package config

import (
	"html/template"

	"github.com/gorilla/sessions"
)

// holds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InProduction  bool
	Store         *sessions.CookieStore
}
