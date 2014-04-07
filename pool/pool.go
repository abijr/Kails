// Makes a pool of templates, and updates the pool if there are changes in the templates.
package pool

import (
	"html/template"
	"io"
)

var pool *template.Template

func init() {
	tmpl, err := template.ParseGlob("/home/potemkin/Projects/go/src/bitbucket.com/abijr/kails/templates/*.tmpl.html")
	if err != nil {
		panic(err)
	}

	pool = tmpl
}

func Render(template string, data interface{}, writer io.Writer) error {
	return pool.ExecuteTemplate(writer, template, data)
}
