package model

import (
	"encoding/json"
)

type ResponseMultipleMail struct {
	Total int    `json:"total"`
	Mail  []Mail `json:"mails"`
}

type Mail struct {
	Id        string   `json:"_id"`
	MessageId string   `json:"message-id,omitempty"`
	Date      string   `json:"date,omitempty"`
	From      string   `json:"from"`
	To        []string `json:"to"`
	Subject   string   `json:"subject"`
	Cc        []string `json:"cc,omitempty"`
	Bcc       []string `json:"bcc,omitempty"`
	Xfrom     string   `json:"x-from,omitempty"`
	Xto       []string `json:"x-to,omitempty"`
	Xcc       []string `json:"x-cc,omitempty"`
	Xbcc      []string `json:"x-bcc,omitempty"`
	Content   string   `json:"content,omitempty"`
	Highlight []string `json:"highlight,omitempty"`
	Path      string   `json:"path,omitempty"`
	Mode      string   `json:"mode,omitempty"`
}

func CreateResponseAPISearch(hitsValues HitsValues) ResponseMultipleMail {
	mails := []Mail{}
	for _, item := range hitsValues.HitsResumen {
		mails = append(mails, Mail{
			Id:        item.Id,
			Subject:   item.ItemSource.Subject,
			From:      item.ItemSource.From,
			To:        item.ItemSource.To,
			Path:        item.ItemSource.Path,
			Content:   item.ItemSource.Content,
			Highlight: item.Highlight.Content,
		})
	}
	return ResponseMultipleMail{
		Total: hitsValues.Total.Value,
		Mail:  mails,
	}
}

func CreateResponseAPIMail(hitsValues HitsValues) Mail {
	result := hitsValues.HitsResumen
	mail := Mail{}
	if len(result) > 0 {
		bodyJson, _ := json.Marshal(result[0].ItemSource)
		json.Unmarshal([]byte(string(bodyJson)), &mail)
		mail.Id = result[0].Id
	}
	return mail
}
