package model

/*
POST  http://localhost:4080/api/mail/_search

{
    "search_type": "term",
    "query": {
        "term": "1UsHHP51tqp",
        "field": "_id"
    }
}

{
    "search_type": "matchphrase",
    "query": {
        "term": "you"
    },
    //"sort_fields": ["-@timestamp"],
    "from": 0,
    "max_results": 1,
    "_source": ["from", "to", "content", "subject", "path", "mode"],
    "highlight": {
        "pre_tags": ["<p color='red'>"],
        "post_tags": ["</p>"],
        "number_of_fragments": 1,
        "fragment_size": 0,
        "fields": {
            "content": {
            }
        }
    }
}*/

type RequestAPISearchQuery struct {
	SearchType string               `json:"search_type"`
	Query      QueryAPIOptions      `json:"query"`
	SortFields []string             `json:"sort_fields,omitempty"`
	From       int                  `json:"from,omitempty"`
	MaxResults int                  `json:"max_results,omitempty"`
	Source     []string             `json:"_source,omitempty"`
	Highlight  *HighlightAPIOptions `json:"highlight,omitempty"`
}

type QueryAPIOptions struct {
	Term  string `json:"term"`
	Field string `json:"field,omitempty"`
}

type HighlightAPIOptions struct {
	PreTags           []string         `json:"pre_tags,omitempty"`
	PostTags          []string         `json:"post_tags,omitempty"`
	NumberOfFragments int              `json:"number_of_fragments,omitempty"`
	FragmentSize      int              `json:"fragment_size,omitempty"`
	Fields            FieldsAPIOptions `json:"fields,omitempty"`
}

type FieldsAPIOptions struct {
	Content interface{} `json:"content,omitempty"`
}
