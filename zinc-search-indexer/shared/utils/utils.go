package utils

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Revisar el parametro de entrada -p para path de la carpeta
func CheckParam() string {
	path := flag.String("p", "", "path of directory")
	flag.Parse()
	if *path == "" {
		panic("-p Path of directory Argument required")
	}
	CheckDirectory(*path)
	return *path
}

func CheckOnlyParam() string {
	if len(os.Args[1:]) > 0 {
		path := os.Args[1]
		CheckDirectory(path)
		return path
	}
	panic("Path of directory Argument required")
}

// Validari si el directorio existe deacuerdo a la ruta especificada
func CheckDirectory(path string) {
	fileInfo, _ := os.Stat(path)
	if !fileInfo.IsDir() {
		panic("Error in path : " + path + " not is a directory")
	}
}

// obtener todos los paths de los archivos dentro de la carpeta especificada
func GetFilesPath(path string, subffix string) []string {
	path_files := []string{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			//var extension = filepath.Ext(info.Name())
			if filepath.Ext(info.Name()) == subffix {
				path_files = append(path_files, path)
			}
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	return path_files
}

func GetFileContent(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func GetEnviroment(enviroment string) string {
	return os.Getenv(enviroment)
}

func CleanTabs(content string) string {
	return strings.Replace(content, "\t", "", -1)
}
