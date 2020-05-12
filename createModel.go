package main

import (
	"os"
	"path/filepath"
	"strings"
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
	var dir string
	if dir, err = getDir(); err != nil {
		return err
	}

	modelName := os.Args[2]
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
	filename := filepath.Join(dir, data.Filename+".go")
	err = saveTemplate(filename, tmplModel, data)
	return
}

func saveRepository(dir string, data *dataModel) (err error) {
	filename := filepath.Join(dir, data.Filename+"Repository.go")
	err = saveTemplate(filename, tmplRepository, data)
	return
}
