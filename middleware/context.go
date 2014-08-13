// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package middleware

import (
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"bitbucket.com/abijr/kails/models"
	"github.com/abijr/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
)

// Context represents context of a request.
type Context struct {
	render.Render
	c   martini.Context
	p   martini.Params
	Req *http.Request
	Res http.ResponseWriter
	// Flash   *Flash
	Session sessions.Session
	// Cache    cache.Cache
	User     models.User // <--- this is needed
	IsLogged bool
	Language string
	Data     map[string]interface{}
}

// Query querys form parameter.
func (ctx *Context) Query(name string) string {
	ctx.Req.ParseForm()
	return ctx.Req.Form.Get(name)
}

// func (ctx *Context) Param(name string) string {
// 	return ctx.p[name]
// }

// HasAPIError returns true if error occurs in form validation.
func (ctx *Context) HasAPIError() bool {
	hasErr, ok := ctx.Data["HasError"]
	if !ok {
		return false
	}
	return hasErr.(bool)
}

// GetErrMsg returns error message string
func (ctx *Context) GetErrMsg() string {
	return ctx.Data["ErrorMsg"].(string)
}

// HasError returns true if error occurs in form validation.
func (ctx *Context) HasError() bool {
	hasErr, ok := ctx.Data["HasError"]
	if !ok {
		return false
	}
	// ctx.Flash.ErrorMsg = ctx.Data["ErrorMsg"].(string)
	// ctx.Data["Flash"] = ctx.Flash
	return hasErr.(bool)
}

// HTML calls render.HTML underlying but reduce one argument.
func (ctx *Context) HTML(status int, name string) {
	ctx.Render.HTML(status, name, ctx.Data, ctx.Language)
}

func (ctx *Context) Redirect(url string) {
	http.Redirect(ctx.Res, ctx.Req, url, 303)
}

// RenderWithErr used for page has form validation but need to prompt error to users.
// func (ctx *Context) RenderWithErr(msg, tpl string, form auth.Form) {
// 	if form != nil {
// 		auth.AssignForm(form, ctx.Data)
// 	}
// 	ctx.Flash.ErrorMsg = msg
// 	ctx.Data["Flash"] = ctx.Flash
// 	ctx.HTML(200, tpl)
// }

// Handle handles and logs error by given status.
// func (ctx *Context) Handle(status int, title string, err error) {
// 	if err != nil {
// 		log.Error("%s: %v", title, err)
// 		if martini.Dev != martini.Prod {
// 			ctx.Data["ErrorMsg"] = err
// 		}
// 	}
//
// 	switch status {
// 	case 404:
// 		ctx.Data["Title"] = "Page Not Found"
// 	case 500:
// 		ctx.Data["Title"] = "Internal Server Error"
// 	}
// 	ctx.HTML(status, fmt.Sprintf("status/%d", status))
// }

// Debug writes given message to log
// func (ctx *Context) Debug(msg string, args ...interface{}) {
// 	log.Debug(msg, args...)
// }

func (ctx *Context) ServeFile(file string, names ...string) {
	var name string
	if len(names) > 0 {
		name = names[0]
	} else {
		name = filepath.Base(file)
	}
	ctx.Res.Header().Set("Content-Description", "File Transfer")
	ctx.Res.Header().Set("Content-Type", "application/octet-stream")
	ctx.Res.Header().Set("Content-Disposition", "attachment; filename="+name)
	ctx.Res.Header().Set("Content-Transfer-Encoding", "binary")
	ctx.Res.Header().Set("Expires", "0")
	ctx.Res.Header().Set("Cache-Control", "must-revalidate")
	ctx.Res.Header().Set("Pragma", "public")
	http.ServeFile(ctx.Res, ctx.Req, file)
}

func (ctx *Context) ServeContent(name string, r io.ReadSeeker, params ...interface{}) {
	modtime := time.Now()
	for _, p := range params {
		switch v := p.(type) {
		case time.Time:
			modtime = v
		}
	}
	ctx.Res.Header().Set("Content-Description", "File Transfer")
	ctx.Res.Header().Set("Content-Type", "application/octet-stream")
	ctx.Res.Header().Set("Content-Disposition", "attachment; filename="+name)
	ctx.Res.Header().Set("Content-Transfer-Encoding", "binary")
	ctx.Res.Header().Set("Expires", "0")
	ctx.Res.Header().Set("Cache-Control", "must-revalidate")
	ctx.Res.Header().Set("Pragma", "public")
	http.ServeContent(ctx.Res, ctx.Req, name, modtime, r)
}

// InitContext initializes a context for a request.
func InitContext() martini.Handler {
	return func(res http.ResponseWriter, r *http.Request, s sessions.Session, c martini.Context, rd render.Render) {

		ctx := &Context{
			c: c,
			// p:      p,
			Req:     r,
			Res:     res,
			Render:  rd,
			Session: s,
			Data:    make(map[string]interface{}),
		}

		ctx.Data["PageStartTime"] = time.Now()

		if username := s.Get("name"); username != nil {
			log.Println(username.(string))
			user, err := models.UserByName(username.(string))
			if err != nil {
				log.Println("Probably cannot marshall...", err)
			}
			ctx.User = *user
			ctx.IsLogged = true
		}
		//TODO: Use martini-contrib sessions here
		// start session

		// rw.Before(func(martini.ResponseWriter) {
		// 	ctx.Session.SessionRelease(res)
		//
		// 	//TODO: martini sessions has flash, do we remove this?
		// 	// or do we adapt it?
		// 	if flash := ctx.Flash.Encode(); len(flash) > 0 {
		// 		ctx.SetCookie("gogs_flash", ctx.Flash.Encode(), 0)
		// 	}
		// })

		// Get user from session if logined.
		// TODO: finish user model, and fix this
		// ......................................................
		// user := auth.SignedInUser(ctx.Session)
		// ctx.User = user
		// ctx.IsSigned = user != nil
		//
		// ctx.Data["IsSigned"] = ctx.IsSigned
		//
		// if user != nil {
		// 	ctx.Data["SignedUser"] = user
		// 	ctx.Data["SignedUserId"] = user.Id
		// 	ctx.Data["SignedUserName"] = user.Name
		// 	ctx.Data["IsAdmin"] = ctx.User.IsAdmin
		// }
		// ..........................................................

		c.Map(ctx)

		c.Next()
	}
}
