package main

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type dataModel struct {
	Filename  string
	Package   string
	Backtick  string
	Model     string
	ModelSymb string
	TableName string
}

func createModel() (err error) {
	modelName := os.Args[2]
	dir := getDir()
	if err = createDir(dir); err != nil {
		return err
	}

	table := modelName
	if []rune(table)[len(table)-1] == 's' {
		table += "es"

	} else {
		table += "s"
	}

	modelRunes := []rune(modelName)

	data := &dataModel{
		Filename:  strings.ToLower(string(modelRunes[0])) + string(modelRunes[1:]),
		Package:   getPackageName(dir),
		Backtick:  backtick,
		Model:     strings.Title(modelName),
		ModelSymb: strings.ToLower(string(modelRunes[0])),
		TableName: table,
	}

	if err = saveModel(dir, data); err != nil {
		return
	}

	err = saveRepository(dir, data)
	return
}

func saveModel(dir string, data *dataModel) (err error) {
	var fModel *os.File

	filename := filepath.Join(dir, data.Filename+".go")
	if err = checkFilename(filename); err != nil {
		return
	}

	fModel, err = os.Create(filename)
	if err != nil {
		return
	}

	defer func() {
		if derr := fModel.Close(); derr != nil {
			err = derr
		}
	}()

	t := template.Must(template.New("").Parse(tmplModel))
	err = t.Execute(fModel, data)
	return
}

func saveRepository(dir string, data *dataModel) (err error) {
	var fRepository *os.File

	filename := filepath.Join(dir, data.Filename+"Repository.go")
	if err = checkFilename(filename); err != nil {
		return
	}

	fRepository, err = os.Create(filename)
	if err != nil {
		return
	}

	defer func() {
		if derr := fRepository.Close(); derr != nil {
			err = derr
		}
	}()

	t := template.Must(template.New("").Parse(tmplRepository))
	err = t.Execute(fRepository, data)
	return
}
