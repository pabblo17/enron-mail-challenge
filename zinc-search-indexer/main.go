package main

import (
	"indexer/domain"
	"indexer/shared/utils"

	"github.com/pkg/profile"
)

const (
	// extencion de los archivos a procesar
	subffixFiles     = "."
	customConversion = true
)

func main() {

	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

	domain.Init()
	pathDirectory := utils.CheckOnlyParam()
	pathFiles := utils.GetFilesPath(pathDirectory, subffixFiles)
	//domain.ProcessFiles(pathFiles, customConversion)
	domain.ProcessFilesBatch(pathFiles, customConversion)
}
