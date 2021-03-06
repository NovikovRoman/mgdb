package main

import (
	"errors"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"
	"text/template"
)

const (
	permDir    = 0755
	backtick   = "`"
	defaultDir = "models"
)

var (
	initCommand   = kingpin.Command("init", "Creates a model interface.").Alias("i")
	interfaceName = initCommand.Arg("name", "Interface name (default: interface).").
		Default("interface").String()
	interfacePath = initCommand.Arg("path", "Directory path (default: "+defaultDir+").").
		Default(defaultDir).String()
	interfaceCtx = initCommand.Flag("context", "Interface with context.").Short('c').
		Default("false").Bool()

	jsonArrayCommand = kingpin.Command("json-array", "Creates a array structure template for json columns.").
		Alias("a")
	jsonArrayName = jsonArrayCommand.Arg("name", "Structure name.").Required().String()
	jsonArrayPath = jsonArrayCommand.Arg("path", "Directory path (default: "+defaultDir+").").
		Default(defaultDir).String()

	jsonStructCommand = kingpin.Command("json-struct", "Creates a structure template for json columns.").
		Alias("s")
	jsonStructName = jsonStructCommand.Arg("name", "Structure name.").Required().String()
	jsonStructPath = jsonStructCommand.Arg("path", "Directory path (default: "+defaultDir+").").
		Default(defaultDir).String()

	modelCommand = kingpin.Command("model", "Creates a model with a repository.").Alias("m")
	modelName    = modelCommand.Arg("name", "Model name.").Required().String()
	modelPath    = modelCommand.Arg("path", "Directory path (default: "+defaultDir+").").
		Default(defaultDir).String()
	modelCtx = modelCommand.Flag("context", "Model with context.").Short('c').
		Default("false").Bool()

	migrateCommand = kingpin.Command("migrate", "").Alias("t")
	migratePath    = migrateCommand.Arg("path", "Directory path (default: migrate).").
		Default("migrate").String()
)

func main() {
	var err error
	kingpin.HelpFlag.Short('h')

	switch kingpin.Parse() {
	case initCommand.FullCommand():
		err = createInterface()

	case jsonArrayCommand.FullCommand():
		err = createJsonArray()

	case jsonStructCommand.FullCommand():
		err = createJsonStruct()

	case modelCommand.FullCommand():
		err = createModel()

	case migrateCommand.FullCommand():
		err = createMigrate()
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// getPackageName returns the package name from the directory name.
func getPackageName(dir string) string {
	return toSnake(filepath.Base(dir))
}

// createDir creates a directory if it does not exist.
func createDir(dir string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, permDir)
		return
	}

	if os.IsExist(err) {
		err = nil
	}
	return
}

// fileDoesNotExist returns nil if the file does not exist.
func fileDoesNotExist(filename string) (err error) {
	if _, err = os.Stat(filename); os.IsNotExist(err) {
		err = nil

	} else {
		err = errors.New(filename + " file exists. ")
	}
	return
}

func saveTemplate(filename string, tmpl string, data interface{}) (err error) {
	var f *os.File
	if err = fileDoesNotExist(filename); err != nil {
		return
	}

	if f, err = os.Create(filename); err != nil {
		return
	}

	defer func() {
		if derr := f.Close(); derr != nil {
			err = derr
		}
	}()

	t := template.Must(template.New("").Parse(tmpl))
	err = t.Execute(f, data)
	return
}
