package middleware

import (
	"log"

	"github.com/go-martini/martini"
)

const (
	DefaultLanguage = "en-us"
)

// Initialized localizer for kails
var Localizer = NewLocalizer(LocalizerOptions{
	DefaultLanguage: DefaultLanguage,
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
		if ctx.IsLogged {
			ctx.InterfaceLanguage = ctx.User.InterfaceLanguage
			return
		}
		sesLang := ctx.Session.Get("Language")

		// check if language no set in session
		if lang, ok := sesLang.(string); !ok {
			log.Printf("### Lang: `%v`, default: `%v`", lang, opt.DefaultLanguage)
			ctx.InterfaceLanguage = opt.DefaultLanguage
		} else {
			ctx.InterfaceLanguage = lang
		}
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
