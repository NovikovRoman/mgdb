package main

import (
	"os"
	"path/filepath"
)

func createInterface() (err error) {
	var dir string
	if dir, err = getDir(); err != nil {
		return err
	}

	packageName := getPackageName(dir)
	name := "interface"
	if len(os.Args) > 2 {
		name = os.Args[2]
	}

	filename := filepath.Join(dir, name+".go")
	data := struct {
		Package  string
		Backtick string
	}{
		Package:  packageName,
		Backtick: backtick,
	}

	err = saveTemplate(filename, tmplModelInterface, data)
	return
}
