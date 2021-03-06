package main

import (
	"github.com/NovikovRoman/gmdb/templates"
	"path/filepath"
)

func createInterface() (err error) {
	if err = createDir(*interfacePath); err != nil {
		return err
	}

	packageName := getPackageName(*interfacePath)
	filename := filepath.Join(*interfacePath, *interfaceName+".go")
	data := struct {
		Package  string
		Backtick string
	}{
		Package:  packageName,
		Backtick: backtick,
	}

	tmpl := templates.Interface
	if *interfaceCtx {
		tmpl = templates.InterfaceWithContext
	}
	err = saveTemplate(filename, tmpl, data)
	return
}
