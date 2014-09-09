package middleware

import (
	"log"

	"github.com/go-martini/martini"
	locale "github.com/nicksnyder/go-i18n/i18n/language"
)

const (
	DefaultLanguage = "en-us"
)

// Initialized localizer for kails
var Localizer = NewLocalizer(LocalizerOptions{
	DefaultLanguage: "en-US",
})

// type localizer struct {
// 	language string
// }
//
// func (l *localizer) Get() string {
// 	return l.language
// }
//
// func (l *localizer) Set(language string) {
// 	lang := locale.Parse(language)
// 	if len(lang) == 0 {
// 		l.language = DefaultLanguage
// 	} else {
// 		l.language = lang[0].Tag
// 	}
// }

type LocalizerOptions struct {
	DefaultLanguage string
}

// type Localizer interface {
// 	// Get the language
// 	Get() string
// 	// Set the language
// 	Set(language string)
// }

func NewLocalizer(options ...LocalizerOptions) martini.Handler {
	opt := prepareLocalizerOptions(options)
	return func(ctx *Context) {
		// Get language from session
		sesLang := ctx.Session.Get("Language")

		var tmpLangs []*locale.Language

		// check if language no set in session
		if _, ok := sesLang.(string); !ok {
			// get language from http header
			reqLang := ctx.Req.Header.Get("Accept-Language")
			tmpLangs = locale.Parse(reqLang)
			log.Println(tmpLangs)
			if len(tmpLangs) == 0 {
				tmpLangs = locale.Parse(opt.DefaultLanguage)
			}
			ctx.Session.Set("Language", tmpLangs[0].Tag)
		} else {
			tmpLangs = locale.Parse(sesLang.(string))
		}
		ctx.Language = tmpLangs[0].Tag
		log.Println("ctx.Language: ", ctx.Language)
	}
}

func prepareLocalizerOptions(opts []LocalizerOptions) LocalizerOptions {
	var opt LocalizerOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	// Defaults
	if len(opt.DefaultLanguage) == 0 {
		opt.DefaultLanguage = DefaultLanguage
	}

	return opt
}
