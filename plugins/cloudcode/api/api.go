package api

import (
	_ "embed"
	"github.com/d5/tengo/v2"
)

//go:embed pb.tengo
var topLevelModule []byte

// GetModules returns a tengo.ModuleMap of cloud code API modules.
func GetModules() *tengo.ModuleMap {
	mm := tengo.NewModuleMap()
	mm.AddSourceModule("pb", topLevelModule)

	// Important: make sure that each of these modules are registered in pb.tengo.
	// Otherwise, users won't be able to do pb.modulename.functionname() as expected.
	// TODO: There is probably some way to programmatically register each of these under the pb module.
	mm.AddBuiltinModule("pb.meta", metaModule)
	mm.AddBuiltinModule("pb.events", eventsModule)

	return mm
}
