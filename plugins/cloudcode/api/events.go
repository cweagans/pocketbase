package api

import (
	"github.com/d5/tengo/v2"
)

// eventsModule makes it possible for cloud code to subscribe to Pocketbase events.
var eventsModule = map[string]tengo.Object{
	"something": &tengo.String{Value: "events module is registered properly"},
}
