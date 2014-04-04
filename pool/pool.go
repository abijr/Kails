//kails is a simple spaced repetion web application.
package pool

import (
	"github.com/stathat/spitz"
	"io"
)

var pool *spitz.Pool

func init() {
	pool = spitz.New("/home/potemkin/Projects/go/src/bitbucket.com/abijr/kails/templates", true)
	err := pool.RegisterLayout("main", "main/header", "main/footer", "", "")
	if err != nil {
		panic(err)
	}

	err = pool.Register("main/index", "", "")
	if err != nil {
		panic(err)
	}
}

func Render(layout string, template string, data interface{}, writer io.Writer) error {
	return pool.Render(layout, template, data, writer)
}
