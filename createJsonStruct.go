package main

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func createJsonStruct() (err error) {
	var f *os.File

	dir := getDir()
	if err = createDir(dir); err != nil {
		return err
	}

	structName := os.Args[2]
	packageName := getPackageName(dir)

	filename := filepath.Join(dir, structName+".go")
	if err = checkFilename(filename); err != nil {
		return
	}

	f, err = os.Create(filename)
	if err != nil {
		return
	}

	defer func() {
		if derr := f.Close(); derr != nil {
			err = derr
		}
	}()

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

	t := template.Must(template.New("").Parse(tmplStringArray))
	err = t.Execute(f, data)
	return
}
