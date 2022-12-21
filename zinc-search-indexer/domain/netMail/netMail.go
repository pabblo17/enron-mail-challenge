package netMail

import (
	"indexer/model"
	"indexer/shared/utils"
	"io"
	"log"
	"net/mail"
	"strings"
)

const (
	toListSplitValue = ","
	toListReplaceOld = ", "
	toListReplaceNew = ","
)

func GetObjectMailMessage(path string) model.Mail {

	r := strings.NewReader(utils.GetFileContent(path))
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}
	header := m.Header
	body, err := io.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
	mail := model.Mail{
		MessageId: header.Get("Message-ID"),
		Date:      header.Get("Date"),
		From:      header.Get("From"),
		To:        parseListItems(header.Get("To")),
		Subject:   header.Get("Subject"),
		Cc:        parseListItems(header.Get("Cc")),
		Bcc:       parseListItems(header.Get("Bcc")),
		Xfrom:     header.Get("X-From"),
		Xto:       parseListItems(header.Get("X-To")),
		Xcc:       parseListItems(header.Get("X-Cc")),
		Xbcc:      parseListItems(header.Get("X-Bcc")),
		Content:   string(body),
		Path:      path,
		Mode:      "net/mail",
	}

	return mail
}

func parseListItems(stringItems string) []string {
	if len(stringItems) <= 0 {
		return nil
	}
	//replacer := strings.NewReplacer(", ", ",", "\t", "")
	//stringValue := replacer.Replace(stringItems)
	stringValue := strings.ReplaceAll(stringItems, ", ", ",")
	return strings.Split(
		stringValue,
		",",
	)
}
