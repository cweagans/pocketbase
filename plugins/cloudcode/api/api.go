package api

import (
	_ "embed"
	"github.com/d5/tengo/v2"
	"github.com/pocketbase/pocketbase/core"
)

//go:embed pb.tengo
var topLevelModule []byte

// _app is a reference to the core.App that was used to initialize the cloud code system.
// This is used by some modules (via getApp()) to interact with the main pocketbase application.
var _app *core.App = nil

func getApp() core.App {
	return *_app
}

// GetModules returns a tengo.ModuleMap of cloud code API modules.
func GetModules(a *core.App) *tengo.ModuleMap {
	_app = a

	mm := tengo.NewModuleMap()
	mm.AddSourceModule("pb", topLevelModule)

	// Important: make sure that each of these modules are registered in pb.tengo.
	// Otherwise, users won't be able to do pb.modulename.functionname() as expected.
	// TODO: There is probably some way to programmatically register each of these under the pb module.
	mm.AddBuiltinModule("pb.dao", daoModule)
	mm.AddBuiltinModule("pb.events", eventsModule)
	mm.AddBuiltinModule("pb.meta", metaModule)

	return mm
}
