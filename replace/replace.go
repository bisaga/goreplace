package replace

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func FindAndReplace(replArgs *ReplaceArguments) error {
	files, err := filepath.Glob(filepath.Join(replArgs.SourceDir, replArgs.SourceFile))
	if err == nil {
		for _, file := range files {
			err = readAndWriteFile(replArgs, file)
			if err != nil {
				return err
			}
		}
	} else {
		return err
	}
	return nil
}

func readAndWriteFile(replArgs *ReplaceArguments, file string) error {
	fileStat, err := os.Stat(file)
	if err != nil {
		log.Fatalln(err)
	}
	data, err := ioutil.ReadFile(file)
	log.Println("Read from: " + file)
	if err == nil {
		markdown := string(data)
		markdown = strings.ReplaceAll(markdown, replArgs.SourceExp, replArgs.TargetExp)
		targetFile := path.Join(replArgs.TargetDir, fileStat.Name())
		err = ioutil.WriteFile(targetFile, []byte(markdown), 0777)
		log.Println("Write to:  " + strings.ReplaceAll(targetFile, "/", "\\"))
		if err != nil {
			log.Println(targetFile)
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}
	return nil
}
