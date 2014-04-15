// Package that makes a pool of template sets
package pool

import (
	"errors"
	"github.com/nicksnyder/go-i18n/i18n"
	"html/template"
	"io"
	"log"
	"os"
	"path"
)

const (
	templateExtension = "*.tmpl.html"
)

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

// Returns a slice with the list of directories (strings)
// inside the given directory
func getSubDirs(dir string) ([]string, error) {

	dir = path.Clean(dir)

	if iv, err := isValidDir(dir); !iv {
		return nil, err
	}

	// gets the file descriptor
	file, err := os.Open(dir)
	if err != nil {
		return nil, err
	}

	fileList, err := file.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	dirList := make([]string, 0, len(fileList))
	for _, f := range fileList {
		iv, _ := isValidDir(path.Join(dir, f))
		if iv {
			dirList = append(dirList, f)
		}
	}

	return dirList, nil
}

func (p *Pool) addTemplateDir(lang, dir string, fmap template.FuncMap) {
	p.templates[lang][dir], _ = template.New("").
		Funcs(fmap). // Load translation function
		ParseGlob(path.Join(p.pathToTemplates, dir, templateExtension))
}

// Loads the template sets in the given path
func NewPool(pathToTemplates, pathToTranslations string, languages []string) (*Pool, error) {
	i18n.MustLoadTranslationFile("translations/all/en-US.all.json")

	if len(languages) == 0 {
		return nil, errors.New("no languages given for pool")
	}

	dirList, err := getSubDirs(pathToTemplates)
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
			p.addTemplateDir(lang, dir, template.FuncMap{"T": T})
			if err != nil {
				log.Println(err)
			}
		}

	}

	return p, nil

}

// Executes the template in the poolset with the given data
func (p *Pool) Render(poolSet, template, lang string, data interface{}, writer io.Writer) error {
	return p.templates[lang][poolSet].ExecuteTemplate(writer, template, data)
}
