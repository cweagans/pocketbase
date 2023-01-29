package api

import (
	"github.com/d5/tengo/v2"
	"github.com/pocketbase/pocketbase"
)

var metaModule = map[string]tengo.Object{
	"version": &tengo.String{Value: pocketbase.Version},
}
