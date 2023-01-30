package types

import (
	"github.com/d5/tengo/v2"
	"github.com/pocketbase/pocketbase/models"
)

type Collection struct {
	tengo.ObjectImpl
	collection *models.Collection
}

func (c *Collection) TypeName() string {
	return "collection(" + c.collection.Name + ")"
}

func (c *Collection) String() string {
	return c.TypeName()
}

func (c *Collection) IndexGet(index tengo.Object) (tengo.Object, error) {
	k, _ := index.(*tengo.String)
	switch k.Value {
	case "name":
		return &tengo.UserFunction{
			Name: "name",
			Value: func(_ ...tengo.Object) (tengo.Object, error) {
				return &tengo.String{Value: c.collection.Name}, nil
			},
		}, nil
	}

	return tengo.UndefinedValue, nil
}
