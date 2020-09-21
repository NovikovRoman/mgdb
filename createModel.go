package main

import (
	"path/filepath"
	"strings"
)

type dataModel struct {
	Filename       string
	Package        string
	Backtick       string
	Model          string
	SliceModelName string
	ModelSymb      string
	TableName      string
}

func createModel() (err error) {
	if err = createDir(*modelPath); err != nil {
		return err
	}

	table := toSnake(*modelName)
	sliceModelName := *modelName
	if []rune(table)[len(table)-1] == 's' {
		table += "es"
		sliceModelName += "es"

	} else {
		table += "s"
		sliceModelName += "s"
	}

	modelRunes := []rune(*modelName)

	data := &dataModel{
		Filename:       strings.ToLower(string(modelRunes[0])) + string(modelRunes[1:]),
		Package:        getPackageName(*modelPath),
		Backtick:       backtick,
		Model:          strings.Title(*modelName),
		SliceModelName: sliceModelName,
		ModelSymb:      strings.ToLower(string(modelRunes[0])),
		TableName:      table,
	}

	if err = saveModel(*modelPath, data); err != nil {
		return
	}

	err = saveRepository(*modelPath, data)
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
