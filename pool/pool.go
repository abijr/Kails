// Package pool makes a pool of template sets
package pool

import (
	"errors"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/nicksnyder/go-i18n/i18n"
)

const (
	templateExtension = "*.tmpl.html"
)

// Pool is a
type Pool struct {
	pathToTemplates    string
	pathToTranslations string
	languages          []string
	templates          map[string]set
}

type set map[string]*template.Template

// Checks if directory is accesible returns
// error if there are missing permissions.
func isValidDir(file string) (bool, error) {
	fi, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false, err
	}
	if err == nil && fi.IsDir() {
		return true, nil
	}
	return false, nil
}

// addTemplateDir returns a slice with the list of directories (strings)
// inside the given directory
func (p *Pool) addTemplateDir(lang, dir string, fmap template.FuncMap) {
	p.templates[lang][dir], _ = template.New("").
		Funcs(fmap).                                                    // Load translation function
		ParseGlob(path.Join(p.pathToTemplates, dir, templateExtension)) // parse the templates in pathToTemplates
}

// NewPool loads the template sets in the given path
func NewPool(pathToTemplates, pathToTranslations string, languages []string) (*Pool, error) {
	//loads translation files
	// TODO: load the files programatically (not hardcoded)
	i18n.MustLoadTranslationFile("_translations/all/en-US.all.json")
	i18n.MustLoadTranslationFile("_translations/all/es-MX.all.json")

	if len(languages) == 0 {
		return nil, errors.New("no languages given for pool")
	}

	dirList, err := ioutil.ReadDir(pathToTemplates)
	if err != nil {
		return nil, err
	}

	p := new(Pool)
	p.pathToTemplates = pathToTemplates
	p.pathToTranslations = pathToTranslations
	p.templates = make(map[string]set)

	// Adds the template sets to the pool variable
	for _, lang := range languages {

		T, _ := i18n.Tfunc(lang) // Make the translation function
		p.templates[lang] = make(set)

		for _, dir := range dirList {
			p.addTemplateDir(lang, dir.Name(), template.FuncMap{"T": T})
			if err != nil {
				log.Println(err)
			}
		}

	}

	return p, nil
}

// Render executes the template in the poolset with the given data
func (p *Pool) Render(poolSet, template, lang string, data interface{}, writer io.Writer) error {
	return p.templates[lang][poolSet].ExecuteTemplate(writer, template, data)
}
