package utils

import (
	"flag"
	"os"
)

// Revisar el parametro de entrada -p para path de la carpeta
func CheckParam() string {
	port := flag.String("port", "", "port ")
	flag.Parse()
	if *port == "" {
		panic("-port ")
	}
	return *port
}

func GetEnviroment(enviroment string) string {
	return os.Getenv(enviroment)
}
