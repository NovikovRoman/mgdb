package main

import (
	"os"
	"path/filepath"
	"text/template"
)

func createInterface() (err error) {
	var f *os.File

	dir := getDir()
	if err = createDir(dir); err != nil {
		return err
	}

	packageName := getPackageName(dir)
	name := "interface"
	if len(os.Args) > 2 {
		name = os.Args[2]
	}

	filename := filepath.Join(dir, name+".go")
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
		Package  string
		Backtick string
	}{
		Package:  packageName,
		Backtick: backtick,
	}

	t := template.Must(template.New("").Parse(tmplModelInterface))
	err = t.Execute(f, data)
	return
}

func createDir(dir string) (err error) {
	_, err = os.Stat(dir)

	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, permDir)
		return
	}

	if os.IsExist(err) {
		err = nil
	}
	return
}
