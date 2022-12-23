package main

import (
	"flag"
	"indexer/domain"
	"indexer/shared/utils"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
)

const (
	// extencion de los archivos a procesar
	subffixFiles     = "."
	customConversion = false
)

var cpuprofile = flag.String("cpuprofile", "cpu_library.pprof", "write cpu profile to file")

func main() {
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	domain.Init()
	pathDirectory := utils.CheckOnlyParam()
	pathFiles := utils.GetFilesPath(pathDirectory, subffixFiles)
	//domain.ProcessFiles(pathFiles, customConversion)
	domain.ProcessFilesBatch(pathFiles, customConversion)
}
