package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// A simple renders templates

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing html template", err)
		return
	}
}

//func createTemplateCache() (map[string]*template.Template, error) {
//	myCache := map[string]*template.Template{}
//}

// Second Render Template with caching.

//
//var tc = make(map[string]*template.Template)
//
//func RenderTemplate(w http.ResponseWriter, t string) {
//	var tmpl *template.Template
//	var err error
//
//	_, inMap := tc[t]
//	if !inMap {
//		fmt.Println("creating new template", t)
//		err = createTemplateCache(t)
//		if err != nil {
//			log.Println(err)
//		}
//	} else {
//		fmt.Println("useing cached template")
//	}
//
//	tmpl = tc[t]
//	err = tmpl.Execute(w, nil)
//	if err != nil {
//		log.Println(err)
//	}
//
//}
//
//func createTemplateCache(t string) error {
//	templates := []string{
//		fmt.Sprintf("./templates/%s", t),
//		"./templates/base.layout.html",
//	}
//
//	tmpl, err := template.ParseFiles(templates...)
//	if err != nil {
//		return err
//	}
//	//add template to cache
//	tc[t] = tmpl
//
//	return nil
//}
