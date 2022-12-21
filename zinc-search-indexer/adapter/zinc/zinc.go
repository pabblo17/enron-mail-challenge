package zinc

import (
	"encoding/json"
	"indexer/model"
	"indexer/shared/utils"
	"io"
	"log"
	"net/http"
	"strings"
)

func DeleteIndex() {
	url := getZincSearchHost() + "index/" + getZincSearchIndex()
	_delete(url, "")
}

func CreateIndex() {
	url := getZincSearchHost() + "index"
	_post(url, GetIndex())
}

func SaveItemsBulk(mails []model.Mail) {
	url := getZincSearchHost() + "/_bulkv2"
	bulk := Bulk{
		Index:   getZincSearchIndex(),
		Records: mails,
	}
	bulkJson, err := json.Marshal(bulk)
	if err != nil {
		log.Fatal(err)
	}

	_post(url, string(bulkJson))
}

func SaveItem(jsonStr string) {
	url := getZincSearchHost() + getZincSearchIndex() + "/_doc"
	_post(url, jsonStr)
}

func _post(url string, _body string) {
	req, err := http.NewRequest("POST", url, strings.NewReader(_body))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(getZincSearchUser(), getZincSearchPassword())
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

func _delete(url string, _body string) {
	req, err := http.NewRequest("DELETE", url, strings.NewReader(_body))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(getZincSearchUser(), getZincSearchPassword())
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

func getZincSearchHost() string {
	return utils.GetEnviroment("ZINC_SEARCH_HOST")
}

func getZincSearchIndex() string {
	return utils.GetEnviroment("ZINC_SEARCH_INDEX")
}

func getZincSearchUser() string {
	return utils.GetEnviroment("ZINC_SEARCH_USER")
}

func getZincSearchPassword() string {
	return utils.GetEnviroment("ZINC_SEARCH_PASSWORD")
}
