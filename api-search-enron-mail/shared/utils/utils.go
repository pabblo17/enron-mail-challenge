package utils

import (
	"flag"
	"os"
)

// Revisar el parametro de entrada -p para path de la carpeta
func CheckParam() string {
	port := flag.String("port", "", "path of directory")
	flag.Parse()
	if *port == "" {
		panic("-p Path of directory Argument required")
	}
	return *port
}

func GetEnviroment(enviroment string) string {
	return os.Getenv(enviroment)
}
