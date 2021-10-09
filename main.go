package main

import (
	"fmt"
	"goreplace/replace"
	"os"
)

func main() {
	replArgs, err := replace.NewReplaceArguments(os.Args)
	if err != nil {
		fmt.Println(err)
		showHelp()
	} else {
		replace.FindAndReplace(replArgs)
	}
}

func showHelp() {
	fmt.Println()
	fmt.Println("Replace command parameters: ")
	fmt.Println("-s/-source {PATH/FILE/*.*}			: source files path or specific file name, wildcards are allowed")
	fmt.Println("-f/-find {expression}				: find specific text")
	fmt.Println("-r/-replace {expression}			: replace with a specific text")
	fmt.Println()
}
