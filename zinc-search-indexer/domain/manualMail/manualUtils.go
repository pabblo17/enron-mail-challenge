package manualMail

import (
	"bufio"
	"errors"
	"indexer/model"
	"log"
	"os"
	"regexp"
	"strings"

	"indexer/shared/utils"
)

const (
	// Constante para separar key: value
	splitMailValue = ":"

	// constantes para ajustar datos separados con comas y espacio por coma
	toListSplitValue = ","
	toListReplaceOld = ", "
	toListReplaceNew = ","
)

func ParseContentToMail(path string) model.Mail {
	mail := PaseseMailContent(path)

	return model.Mail{
		MessageId: mail["Message-ID"],
		Date:      mail["Date"],
		From:      mail["From"],
		To:        parseListItems(mail["To"]),
		Subject:   utils.CleanTabs(mail["Subject"]),
		Cc:        parseListItems(mail["Cc"]),
		Bcc:       parseListItems(mail["Bcc"]),
		Xfrom:     mail["X-From"],
		Xto:       parseListItems(mail["X-To"]),
		Xcc:       parseListItems(mail["X-Cc"]),
		Xbcc:      parseListItems(mail["X-Bcc"]),
		Content:   mail["content"],
		Path:      path,
		Mode:      "custom",
	}
}

// Separar el header y body cuando se tenga el primer salto de linea
func PaseseMailContent(path string) map[string]string {
	var currentKey string
	var currentValue string
	var bodyContent strings.Builder
	var err error

	flagHeader := true
	mail := map[string]string{}

	file, err := os.Open(path)
	if err != nil {
		log.Printf("Could not open the file due to this %s error \n", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var currentLine = scanner.Text()
		if currentLine == "" {
			flagHeader = false
		}
		if flagHeader {
			currentKey, currentValue, err = getKeyValuesToString(currentKey, currentValue, currentLine)
			if err != nil {
				delete(mail, currentKey)
			} else {
				mail[currentKey] = currentValue
			}
		} else {
			bodyContent.WriteString(currentLine + "\n")
		}
	}
	mail["content"] = bodyContent.String()

	return mail
}

func getKeyValuesToString(
	currentKey string,
	currentValue string,
	currentLine string,
) (string, string, error) {
	isLowerCase := regexp.MustCompile("[A-Z]")
	if isLowerCase.Match([]byte(string(currentLine[0]))) {
		return separeKeyValue(currentLine)
	}
	return currentKey, appendValues(currentValue, currentLine), nil
}

func separeKeyValue(content string) (string, string, error) {
	params := strings.Split(content, splitMailValue)
	var key = params[0]
	var value = strings.TrimSpace(strings.Join(params[1:], ""))
	if value == "" {
		return key, value, errors.New("Empty Value")
	}
	return key, value, nil
}

func appendValues(value string, appendValue string) string {
	return value + appendValue
}

func parseListItems(stringItems string) []string {
	if len(stringItems) <= 0 {
		return nil
	}
	replacer := strings.NewReplacer(", ", ",", "\t", "")
	stringValue := replacer.Replace(stringItems)
	//stringValue := strings.ReplaceAll(stringItems, ", ", ","),
	return strings.Split(
		stringValue,
		",",
	)
}
