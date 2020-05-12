package main

import (
	"os"
	"path/filepath"
	"strings"
)

func createJsonStruct() (err error) {
	var dir string
	if dir, err = getDir(); err != nil {
		return err
	}

	structName := os.Args[2]
	packageName := getPackageName(dir)

	filename := filepath.Join(dir, structName+".go")
	data := struct {
		Package    string
		Backtick   string
		Struct     string
		StructSymb string
	}{
		Package:    packageName,
		Backtick:   backtick,
		Struct:     strings.Title(structName),
		StructSymb: strings.ToLower(string([]rune(structName)[0])),
	}

	err = saveTemplate(filename, tmplStringArray, data)
	return
}
