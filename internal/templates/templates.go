package templates

import (
	"embed"
	"html/template"
	"io"
)

//go:embed templates/*
var templateFS embed.FS

func ExecuteTemplate(w io.Writer, file string, data interface{}) error {
	var err error
	layoutContent, err := templateFS.ReadFile("templates/_layout.gohtml")
	if err != nil {
		return err
	}

	layout := template.Must(template.New("_layout.gohtml").Parse(string(layoutContent)))//.Funcs(funcMap)

	content, err := templateFS.ReadFile("templates/content/" + file)
	if err != nil {
		return err
	}

	tmpl := template.Must(layout.Clone())
	_, err = tmpl.Parse(string(content))
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
