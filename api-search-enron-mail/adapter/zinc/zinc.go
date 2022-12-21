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

func getFieldsToSearch() []string {
	return []string{
		"from",
		"to",
		"x-to",
		"content",
		"subject",
		"path",
	}
}

func SearchByIdAPI(id string) string {
	url := getZincSearchHostAPI() + getZincSearchIndex() + "/_search"
	body := model.RequestAPISearchQuery{
		SearchType: "term",
		Query: model.QueryAPIOptions{
			Term:  id,
			Field: "_id",
		},
	}
	bodyJson, _ := json.Marshal(body)
	return _post(url, string(bodyJson))
}

func SearchByContentAPI(term string, from int, maxResult int) string {
	url := getZincSearchHostAPI() + getZincSearchIndex() + "/_search"

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
	return _post(url, string(bodyJson))
}

func SearchByContentEs(term string) string {
	url := getZincSearchHostEs() + getZincSearchIndex() + "/_search"

	body := model.RequestEsSearchQuery{
		Query: model.QueryOptions{
			Match: model.MatchOptions{
				Content: term,
			},
		},
		From:       0,
		MaxResults: 1,
		Source:     getFieldsToSearch(),
		Highlight: model.HighlightOptions{
			Fields: model.FieldsOptions{
				Content: map[string]string{},
			},
		},
	}
	bodyJson, _ := json.Marshal(body)
	return _post(url, string(bodyJson))
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

func _post(url string, _body string) string {
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
	return string(body)
}

func getZincSearchHostEs() string {
	return utils.GetEnviroment("ZINC_SEARCH_HOST_ES")
}

func getZincSearchHostAPI() string {
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
