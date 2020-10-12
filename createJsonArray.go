package main

import (
	"github.com/NovikovRoman/gmdb/templates"
	"path/filepath"
	"strings"
)

func createJsonArray() (err error) {
	if err = createDir(*jsonArrayPath); err != nil {
		return err
	}

	packageName := getPackageName(*jsonArrayPath)
	filename := filepath.Join(*jsonArrayPath, *jsonArrayName+".go")
	data := struct {
		Package    string
		Backtick   string
		Struct     string
		StructSymb string
	}{
		Package:    packageName,
		Backtick:   backtick,
		Struct:     strings.Title(*jsonArrayName),
		StructSymb: strings.ToLower(string([]rune(*jsonArrayName)[0])),
	}

	err = saveTemplate(filename, templates.StringArray, data)
	return
}
