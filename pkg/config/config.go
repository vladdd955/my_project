package config

import "html/template"

// AppConfig holds app config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
