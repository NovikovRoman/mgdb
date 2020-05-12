package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	permDir    = 0755
	backtick   = "`"
	defaultDir = "models"
)

func main() {
	var err error

	if len(os.Args) < 2 ||
		os.Args[1] == "h" || os.Args[1] == "-h" ||
		os.Args[1] == "help" || os.Args[1] == "-help" || os.Args[1] == "--help" {
		help(nil)
		os.Exit(0)
	}

	switch os.Args[1] {

	// создать директорию для моделей (по-умолчанию models) с интерфесом модели, методами.
	case "i", "init":
		err = createInterface()
		break

	// создать структуру типа []string для json-столбца
	case "j", "json-struct":
		if len(os.Args) < 3 {
			help(errors.New("No structure name specified. "))
			return
		}
		err = createJsonStruct()
		break

	// создать модель с репозиторием
	case "m", "model":
		if len(os.Args) < 3 {
			help(errors.New("No model name specified. "))
			return
		}
		err = createModel()
		break

	default:
		help(errors.New("Unknown command. "))
		os.Exit(1)
	}

	if err != nil {
		help(err)
		os.Exit(1)
	}

	os.Exit(0)
}

// getPackageName returns the package name from the directory name.
func getPackageName(dir string) string {
	return strings.ToLower(filepath.Base(dir))
}

// getDir returns a directory of models.
func getDir() (dir string, err error) {
	dir = defaultDir
	if len(os.Args) > 3 {
		dir = os.Args[3]
	}

	err = createDir(dir)
	return
}

// createDir creates a directory if it does not exist.
func createDir(dir string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, permDir)
		return
	}

	if os.IsExist(err) {
		err = nil
	}
	return
}

// fileDoesNotExist returns nil if the file does not exist.
func fileDoesNotExist(filename string) (err error) {
	if _, err = os.Stat(filename); os.IsNotExist(err) {
		err = nil

	} else {
		err = errors.New(filename + " file exists. ")
	}
	return
}

func saveTemplate(filename string, tmpl string, data interface{}) (err error) {
	var f *os.File
	if err = fileDoesNotExist(filename); err != nil {
		return
	}

	f, err = os.Create(filename)
	if err != nil {
		return
	}

	t := template.Must(template.New("").Parse(tmpl))
	err = t.Execute(f, data)

	ferr := f.Close()
	if err == nil {
		err = ferr
	}
	return
}
