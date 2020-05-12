package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
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

func getPackageName(dir string) string {
	return strings.ToLower(filepath.Base(dir))
}

func getDir() (dir string) {
	dir = defaultDir
	if len(os.Args) > 3 {
		dir = os.Args[3]
	}
	return
}

func checkFilename(filename string) (err error) {
	if _, err = os.Stat(filename); os.IsNotExist(err) {
		err = nil

	} else {
		err = errors.New(filename + " file exists. ")
	}
	return
}
