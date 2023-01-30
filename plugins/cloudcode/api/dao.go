package api

import (
	"github.com/d5/tengo/v2"
	"github.com/pocketbase/pocketbase/plugins/cloudcode/types"
)

var daoModule = map[string]tengo.Object{
	"findCollectionByNameOrId": &tengo.UserFunction{
		Name:  "findCollectionByNameOrId",
		Value: findCollectionByNameOrId,
	},
}

func findCollectionByNameOrId(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		return nil, tengo.ErrWrongNumArguments
	}

	collectionName, _ := tengo.ToString(args[0])

	c, err := getApp().Dao().FindCollectionByNameOrId(collectionName)
	if err != nil {
		return nil, err
	}

	return types.WrapCollection(c), nil
}
