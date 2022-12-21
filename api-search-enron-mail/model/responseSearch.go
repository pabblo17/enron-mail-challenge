package model

/*
{
    "took": 0,
    "timed_out": false,
    "_shards": {
        "total": 1,
        "successful": 1,
        "skipped": 0,
        "failed": 0
    },
    "hits": {
        "total": {
            "value": 16
        },
        "max_score": 1.2701945663230216,
        "hits": [
            {
                "_index": "mail",
                "_type": "_doc",
                "_id": "1Uhe6OkZ5T4",
                "_score": 1.2701945663230216,
                "@timestamp": "2022-12-09T08:52:26.169928704Z",
                "_source": {
                    "content": "",
                    "from": "",
                    "path": "",
                    "subject": "",
                    "to": []
                },
				highlight : [
					{
						content: []
					}
				]
            },
		]
*/

// Response to zinc
type ResponseZincQuery struct {
	Took int        `json:"took"`
	Hits HitsValues `json:"hits"`
}

type HitsValues struct {
	Total       TotalHitsValues     `json:"total"`
	MaxScore    float32             `json:"max_score"`
	HitsResumen []HitsResumenValues `json:"hits"`
}

type TotalHitsValues struct {
	Value int `json:"value"`
}

type HitsResumenValues struct {
	Id         string          `json:"_id"`
	ItemSource Mail            `json:"_source"`
	Highlight  HighlightValues `json:"highlight"`
}

type HighlightValues struct {
	Content []string `json:"content"`
}
