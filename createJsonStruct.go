package main

import (
	"path/filepath"
	"strings"
)

func createJsonStruct() (err error) {
	if err = createDir(*jsonStructPath); err != nil {
		return err
	}

	packageName := getPackageName(*jsonStructPath)
	filename := filepath.Join(*jsonStructPath, *jsonStructName+".go")
	data := struct {
		Package    string
		Backtick   string
		Struct     string
		StructSymb string
	}{
		Package:    packageName,
		Backtick:   backtick,
		Struct:     strings.Title(*jsonStructName),
		StructSymb: strings.ToLower(string([]rune(*jsonStructName)[0])),
	}

	err = saveTemplate(filename, tmplStringArray, data)
	return
}
