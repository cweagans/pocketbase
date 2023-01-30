package api

import (
	"fmt"
	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
	"github.com/pocketbase/pocketbase/core"
)

// eventsModule makes it possible for cloud code to subscribe to Pocketbase events.
var eventsModule = map[string]tengo.Object{
	"onBeforeBootstrap": &tengo.UserFunction{
		Name:  "onBeforeBootstrap",
		Value: onBeforeBootstrap,
	},
	"onAfterBootstrap": &tengo.UserFunction{
		Name:  "onAfterBootstrap",
		Value: onAfterBootstrap,
	},
}

func onBeforeBootstrap(args ...tengo.Object) (ret tengo.Object, err error) {
	stdlib.FuncAR()

	if len(args) != 1 {
		return nil, tengo.ErrWrongNumArguments
	}
	if !args[0].CanCall() {
		return nil, tengo.ErrInvalidArgumentType{Name: "callback", Expected: "user-function", Found: args[0].TypeName()}
	}

	callback := args[0].(*tengo.CompiledFunction)

	getApp().OnBeforeBootstrap().Add(func(_ *core.BootstrapEvent) error {
		fmt.Println("called on before bootstrap hook defined in cloud code")
		_, err := callback.Call()
		return err
	})

	return nil, nil
}
func onAfterBootstrap(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		return nil, tengo.ErrWrongNumArguments
	}
	if !args[0].CanCall() {
		return nil, tengo.ErrInvalidArgumentType{Name: "callback", Expected: "user-function", Found: args[0].TypeName()}
	}

	callback := args[0].(*tengo.CompiledFunction)

	getApp().OnAfterBootstrap().Add(func(_ *core.BootstrapEvent) error {
		fmt.Println("called on after bootstrap hook defined in cloud code")
		_, err := callback.Call()
		return err
	})

	return nil, nil
}
