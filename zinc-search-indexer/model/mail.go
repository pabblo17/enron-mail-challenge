package model

type Mail struct {
	MessageId string   `json:"message-id"`
	Date      string   `json:"date"`
	From      string   `json:"from"`
	To        []string `json:"to"`
	Subject   string   `json:"subject"`
	Cc        []string `json:"cc,omitempty"`
	Bcc       []string `json:"bcc,omitempty"`
	Xfrom     string   `json:"x-from"`
	Xto       []string `json:"x-to"`
	Xcc       []string `json:"x-cc,omitempty"`
	Xbcc      []string `json:"x-bcc,omitempty"`
	Content   string   `json:"content"`
	Path      string   `json:"path"`
	Mode      string   `json:"mode"`
}
