package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"indexer/adapter/zinc"
	"indexer/domain/manualMail"
	"indexer/domain/netMail"
	"indexer/model"
	"indexer/shared/utils"
	"log"
	"strconv"
)

func Init() {
	zinc.DeleteIndex()
	zinc.CreateIndex()
}

func ProcessFiles(pathFiles []string, customConversion bool) {
	for _, item := range pathFiles {
		//contentFile := utils.GetFileContent(item)
		jsonMail, err := getMailJsonStringFormat(item, customConversion)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		zinc.SaveItem(jsonMail)
	}
}

func ProcessFilesBatch(pathFiles []string, customConversion bool) {
	log.Println("Process Files Batch starting......")
	total := len(pathFiles)
	var batch []model.Mail
	batchSiZe, _ := strconv.Atoi(utils.GetEnviroment("BATCH_SIZE"))

	log.Printf("Batch size: %v \nTotal records: %v", batchSiZe, total)
	for i, item := range pathFiles {
		//contentFile := utils.GetFileContent(item)
		batch = append(batch, getMail(item, customConversion))
		if (i+1)%batchSiZe == 0 {
			log.Printf("Save documents for batch %v / %v", i+1, total)
			zinc.SaveItemsBulk(batch)
			batch = nil
		} else if batch != nil && (i+1) == total {
			log.Printf("Save documents for batch %v / %v", i+1, total)
			zinc.SaveItemsBulk(batch)
		}
	}
}

func getMail(path string, customConversion bool) model.Mail {
	if customConversion {
		return manualMail.ParseContentToMail(path)
	}
	return netMail.GetObjectMailMessage(path)
}

func getMailJsonStringFormat(path string, customConversion bool) (string, error) {
	var objMail model.Mail
	if customConversion {
		objMail = manualMail.ParseContentToMail(path)
	} else {
		objMail = netMail.GetObjectMailMessage(path)
	}
	jsonStr, err := json.Marshal(objMail)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return "", errors.New("Problems with parse content for file :" + path)
	}
	return string(jsonStr), nil
}
