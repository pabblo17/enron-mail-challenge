package model

/*
POST http://localhost:4080/es/{index}/_search
{
    "query": {
        "match": {
            "content": "you"
        }
    },
    "from": 0,
    "max_results": 5,
    "_source":["from", "to", "content", "subject", "path"],
    "highlight": {
        "fields": {
            "content": {}
        }
    }
}
*/

type RequestEsSearchQuery struct {
	Query      QueryOptions     `json:"query"`
	From       int              `json:"from"`
	MaxResults int              `json:"max_results,omitempty"`
	Source     []string         `json:"_source,omitempty"`
	Highlight  HighlightOptions `json:"highlight,omitempty"`
}

type QueryOptions struct {
	Match MatchOptions `json:"match,omitempty"`
}

type MatchOptions struct {
	Content string `json:"content,omitempty"`
}

type HighlightOptions struct {
	Fields FieldsOptions `json:"fields,omitempty"`
}

type FieldsOptions struct {
	Content interface{} `json:"content,omitempty"`
}
