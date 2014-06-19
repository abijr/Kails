package localization

import (
	"bitbucket.com/abijr/kails/middleware"

	"github.com/go-martini/martini"
	"github.com/nicksnyder/go-i18n/i18n/locale"
)

const (
	DefaultLanguage = "en-US"
)

type localizer struct {
	language string
}

func (l *localizer) Get() string {
	return l.language
}

func (l *localizer) Set(language string) {
	lang, err := locale.New(language)
	if err != nil {
		l.language = DefaultLanguage
	} else {
		l.language = lang.ID
	}
}

type Options struct {
	DefaultLanguage string
}

type Localizer interface {
	// Get the language
	Get() string
	// Set the language
	Set(language string)
}

func NewLocalizer(options ...Options) martini.Handler {
	opt := prepareOptions(options)
	return func(ctx *middleware.Context) {
		// Get language from session
		sesLang := ctx.Session.Get("Language")

		var language *locale.Locale

		// check if language no set in session
		if _, ok := sesLang.(string); !ok {
			var err error

			// get language from http header
			reqLang := ctx.Req.Header.Get("Accept-Language")
			language, err = locale.New(reqLang)
			if err != nil {
				language, _ = locale.New(opt.DefaultLanguage)
			}
			ctx.Session.Set("Language", language.ID)
		} else {
			language, _ = locale.New(sesLang.(string))
		}
		ctx.Language = language.ID
	}
}

func prepareOptions(opts []Options) Options {
	var opt Options
	if len(opts) > 0 {
		opt = opts[0]
	}

	// Defaults
	if len(opt.DefaultLanguage) == 0 {
		opt.DefaultLanguage = DefaultLanguage
	}

	return opt
}
