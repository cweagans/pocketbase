package cloudcode

import (
	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/parser"
	"github.com/d5/tengo/v2/stdlib"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/cloudcode/api"
	"log"
	"os"
	"path/filepath"
)

// MustRegister initializes a Tengo environment and panics if anything goes wrong.
// See Register for param information.
func MustRegister(app core.App, dir string, init string) {
	if err := Register(app, dir, init); err != nil {
		panic(err)
	}
}

// Register initializes a Tengo environment for the application.
// `init` is the name of the first cloud code file to load. `dir` is the path to the cloud
// code directory (defaults to pb_data/../pb_cloudcode if empty).
func Register(app core.App, dir string, init string) error {
	if dir == "" {
		dir = filepath.Join(app.DataDir(), "..", "pb_cloudcode")
	}
	init = filepath.Join(dir, init)

	// If the cloud code or init file doesn't exist, just exit early.
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// TODO: Is there a better way to handle log output?
		log.Println("could not find cloud code dir")
		return nil
	}
	if _, err := os.Stat(init); os.IsNotExist(err) {
		// TODO: Is there a better way to handle log output?
		log.Println("could not find cloud code init")
		return nil
	}

	src, err := os.ReadFile(init)
	if err != nil {
		return err
	}

	modules := getModuleMap(&app)
	symbolTable := getSymbolTable()
	constants := []tengo.Object{}

	bytecode, err := compile(modules, src, init, symbolTable, constants)
	if err != nil {
		return err
	}

	machine := tengo.NewVM(bytecode, nil, -1)
	err = machine.Run()
	if err != nil {
		return err
	}

	return nil
}

func getModuleMap(app *core.App) *tengo.ModuleMap {
	// Enable all stdlib tengo modules except os.
	// Cloud code shouldn't be able to start other programs until we have better sandboxing.
	// TODO: also consider how to do outgoing http requests. something like https://gitlab.com/Ma_124/httpbox maybe?
	modules := stdlib.GetModuleMap(stdlib.AllModuleNames()...)
	modules.Remove("os")

	// Add the modules that are specific to the pocketbase cloud code system.
	modules.AddMap(api.GetModules(app))

	return modules
}

func getSymbolTable() *tengo.SymbolTable {
	symbolTable := tengo.NewSymbolTable()

	// Register builtin functions.
	for idx, fn := range tengo.GetAllBuiltinFunctions() {
		symbolTable.DefineBuiltin(idx, fn.Name)
	}

	// TODO: Add any custom builtins.

	return symbolTable
}

func compile(modules *tengo.ModuleMap, src []byte, inputFile string, symbolTable *tengo.SymbolTable, constants []tengo.Object) (*tengo.Bytecode, error) {
	fileSet := parser.NewFileSet()
	srcFile := fileSet.AddFile(filepath.Base(inputFile), -1, len(src))

	p := parser.NewParser(srcFile, src, nil)
	file, err := p.ParseFile()
	if err != nil {
		return nil, err
	}

	c := tengo.NewCompiler(srcFile, symbolTable, constants, modules, nil)
	c.EnableFileImport(true)
	c.SetImportDir(filepath.Dir(inputFile))

	if err := c.Compile(file); err != nil {
		return nil, err
	}

	bytecode := c.Bytecode()
	bytecode.RemoveDuplicates()
	return bytecode, nil
}
