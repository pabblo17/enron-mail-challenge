package main

import (
	"indexer/domain"
	"indexer/shared/utils"
)

const (
	// extencion de los archivos a procesar
	subffixFiles     = "."
	customConversion = true
)

func main() {

	domain.Init()
	pathDirectory := utils.CheckOnlyParam()
	pathFiles := utils.GetFilesPath(pathDirectory, subffixFiles)
	//domain.ProcessFiles(pathFiles, customConversion)
	domain.ProcessFilesBatch(pathFiles, customConversion)
}
