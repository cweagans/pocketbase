package api

import (
	"github.com/d5/tengo/v2"
	"github.com/pocketbase/pocketbase"
)

var databaseModule = map[string]tengo.Object{
	"version": &tengo.String{Value: pocketbase.Version},
}
