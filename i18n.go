package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

var translationsCache map[string]map[string]string
var cacheMutex sync.RWMutex

func getTranslations(languageCode string) map[string]string {
	cacheMutex.RLock()
	cachedData, found := translationsCache[languageCode]
	cacheMutex.RUnlock()

	if found {
		return cachedData
	}

	content, err := ioutil.ReadFile("resources/i18n.json")
	if err != nil {
		log.Fatal(err)
	}
	jsonData := string(content)

	var data map[string]map[string]string

	err = json.Unmarshal([]byte(jsonData), &data)
	if err == nil {
		rData := make(map[string]string)
		for k, v := range data {
			rData[k] = v[languageCode]
		}

		cacheMutex.Lock()
		if translationsCache == nil {
			translationsCache = make(map[string]map[string]string)
		}
		translationsCache[languageCode] = rData
		cacheMutex.Unlock()

		return rData
	}

	return map[string]string{}
}
