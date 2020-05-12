package main

import "fmt"

func help(err error) {
	if err != nil {
		fmt.Println(wrapError(err.Error()))
	}

	fmt.Println(wrapBold("---------------------------------- help ----------------------------------"))
	fmt.Println("mgdb command [name] [path]\n\ncommands:")

	fmt.Println(wrapSuccess("i, init"), "creates a model interface")
	fmt.Println("\t", wrapBold("name"), "file name", wrapGray("(default interface)"))
	fmt.Println("\t", wrapBold("path"), "directory path", wrapGray("(default models)"))

	fmt.Println(wrapSuccess("j, json-struct"), "creates a structure template for json columns")
	fmt.Println("\t", wrapBold("name"), "structure name")
	fmt.Println("\t", wrapBold("path"), "directory path", wrapGray("(default models)"))

	fmt.Println(wrapSuccess("m, model"), "creates a model with a repository")
	fmt.Println("\t", wrapBold("name"), "model name")
	fmt.Println("\t", wrapBold("path"), "directory path", wrapGray("(default models)"))
}

func wrapBold(text string) string {
	return "\u001B[1m" + text + "\u001B[0m"
}

func wrapGray(text string) string {
	return "\u001B[90m" + text + "\u001B[0m"
}

func wrapSuccess(text string) string {
	return "\u001B[32m" + text + "\u001B[0m"
}

func wrapError(text string) string {
	return "\u001B[97;41m" + text + "\u001B[0m"
}
