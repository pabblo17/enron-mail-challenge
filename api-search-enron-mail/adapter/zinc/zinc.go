package zinc

import (
	"backend/model"
	"backend/shared/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func CreateIndex(_body string) string {
	url := getZincSearchHostAPI() + "index/"
	return _executeHttpRequest(url, _body, "POST")
}

func DeleteIndex(index string) string {
	url := getZincSearchHostAPI() + "index/" + index
	return _executeHttpRequest(url, "", "DELETE")
}

func SaveDocumentBulk(index string, mails []model.Mail) string {
	url := getZincSearchHostAPI() + index + "/_search"
	bulk := Bulk{
		Index:   index,
		Records: mails,
	}
	bulkJson, err := json.Marshal(bulk)
	if err != nil {
		log.Fatal(err)
	}
	return _executeHttpRequest(url, string(bulkJson), "POST")
}

func SaveDocument(index string, _body string) string {
	url := getZincSearchHostAPI() + index + "/_doc"
	return _executeHttpRequest(url, _body, "POST")
}

func SearchByIdAPI(index string, id string) string {
	url := getZincSearchHostAPI() + index + "/_search"
	body := model.RequestAPISearchQuery{
		SearchType: "term",
		Query: model.QueryAPIOptions{
			Term:  id,
			Field: "_id",
		},
	}
	bodyJson, _ := json.Marshal(body)
	return _executeHttpRequest(url, string(bodyJson), "POST")
}

func SearchByContentAPI(index string, term string, from int, maxResult int) string {
	url := getZincSearchHostAPI() + index + "/_search"

	body := model.RequestAPISearchQuery{
		SearchType: "matchphrase",
		Query: model.QueryAPIOptions{
			Term: term,
		},
		From:       from * maxResult,
		MaxResults: maxResult,
		Source:     []string{},
		Highlight: &model.HighlightAPIOptions{
			PreTags:           []string{"<mark>"},
			PostTags:          []string{"</mark>"},
			NumberOfFragments: 1,
			FragmentSize:      10,
			Fields: model.FieldsAPIOptions{
				Content: map[string]string{},
			},
		},
	}
	bodyJson, _ := json.Marshal(body)
	return _executeHttpRequest(url, string(bodyJson), "POST")
}

func CreateResponseSearchService(responseQuery string) model.ResponseMultipleMail {
	data := model.ResponseZincQuery{}
	json.Unmarshal([]byte(responseQuery), &data)
	return model.CreateResponseAPISearch(data.Hits)
}

func CreateResponseSearchIdService(responseQuery string) model.Mail {
	data := model.ResponseZincQuery{}
	json.Unmarshal([]byte(responseQuery), &data)
	return model.CreateResponseAPIMail(data.Hits)
}

func _executeHttpRequest(url string, _body string, method string) string {
	req, err := http.NewRequest(method, url, strings.NewReader(_body))
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
		log.Printf("Error %v", err)
	}
	return string(body)
}

func getZincSearchHostAPI() string {
	return utils.GetEnviroment("ZINC_SEARCH_HOST")
}

func getZincSearchUser() string {
	return utils.GetEnviroment("ZINC_SEARCH_USER")
}

func getZincSearchPassword() string {
	return utils.GetEnviroment("ZINC_SEARCH_PASSWORD")
}
