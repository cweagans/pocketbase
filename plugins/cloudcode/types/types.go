package types

import (
	"github.com/pocketbase/pocketbase/models"
)

func WrapCollection(c *models.Collection) *Collection {
	return &Collection{collection: c}
}
