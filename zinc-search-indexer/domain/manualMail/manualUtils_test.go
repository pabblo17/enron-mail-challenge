package manualMail

import (
	"encoding/json"
	"indexer/model"
	"indexer/shared/utils"
	"reflect"
	"testing"
)

func TestAllProces(t *testing.T) {
	pathDirectory := "./test"
	var pathFiles = utils.GetFilesPath(pathDirectory, ".")
	for _, pathFile := range pathFiles {
		var objMail = ParseContentToMail(pathFile)
		if _, err := json.Marshal(objMail); err != nil {
			//			zinc.SaveItem(string(jsonStr))
			t.Error("Error to process :", err)
		}

	}
}

func TestParseContentToMail(t *testing.T) {
	path := "./test/1."
	mail := ParseContentToMail(path)
	mailObj := model.Mail{}
	subject := mail.Subject
	if !reflect.DeepEqual(mail, mailObj) && subject != "UT Super Saturday Candidates" {
		t.Error("Error en parse de valores", mail.Subject)
	}
}

func TestGetKeyValuesToString1(t *testing.T) {

	currentLine := "X-From: Lexi Elliott"
	if key, value, err := getKeyValuesToString("", "", currentLine); err != nil {
		if !(key == "X-From" && value == "Lexi Elliott") {
			t.Error("Error to get key and values to ", key, value)
		}
	}
}

func TestGetKeyValuesToString2(t *testing.T) {

	currentLine := "X-From: 	Lexi Elliott"
	if key, value, err := getKeyValuesToString("", "", currentLine); err != nil {
		if !(key == "X-From" && value == "Lexi Elliott") {
			t.Error("Error to get key and values to ", key, value)
		}
	}
}

func TestGetKeyValuesToString3(t *testing.T) {
	currentLine := `To: richard.causey@enron.com, mike.deville@enron.com, mark.lindsey@enron.com, 
	brent.price@enron.com, stan.dowell@enron.com
`
	correctValue := "richard.causey@enron.com, mike.deville@enron.com, mark.lindsey@enron.com, brent.price@enron.com, stan.dowell@enron.com"
	if key, value, err := getKeyValuesToString("", "", currentLine); err != nil {
		if !(key == "To" && value == correctValue) {
			t.Error("Error to get key and values to ", key, value)
		}
	}
}

func TestParseList(t *testing.T) {
	testValue := "richard.causey@enron.com, mike.deville@enron.com, mark.lindsey@enron.com, brent.price@enron.com, stan.dowell@enron.com"
	correctValue := []string{
		"richard.causey@enron.com",
		"mike.deville@enron.com",
		"mark.lindsey@enron.com",
		"brent.price@enron.com",
		"stan.dowell@enron.com",
	}
	var testParseValue []string = parseListItems(testValue)
	if !(len(testParseValue) == len(correctValue)) {
		t.Error("Error to parse list")
	}
}
