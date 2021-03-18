package route

import (
	"goframe/module"

	"github.com/imkouga/gocore/http/route"
)

func Init() error {
	route.RawHandleFunc("/hello", module.Hello)
	return nil
}
