package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func getTranslations(languageCode string) map[string]string {
	content, errReadFile := ioutil.ReadFile("resources/i18n.json")
	if errReadFile != nil {
		log.Fatal(errReadFile)
	}
	jsonData := string(content)

	var data map[string]map[string]string

	err := json.Unmarshal([]byte(jsonData), &data)
	if err == nil {
		rData := make(map[string]string)
		for k, v := range data {
			rData[k] = v[languageCode]
		}
		return rData
	}
	return map[string]string{}
}
