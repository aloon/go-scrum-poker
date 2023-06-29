package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"sync"
)

var translationsCache map[string]map[string]string
var cacheMutex sync.RWMutex

func stringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func getLanguageCodeFromHeader(acceptLanguageHeader string) string {
	defaultCode := "en"
	validLanguges := []string{"en", "es", "ca", "de", "fr", "it"}
	if acceptLanguageHeader == "" {
		return defaultCode
	} else {
		result := strings.Split(strings.Split(acceptLanguageHeader, ",")[0], "-")[0]
		if stringInSlice(result, validLanguges) {
			return result
		} else {
			return defaultCode
		}
	}
}

func getTranslations(languageCode string) map[string]string {
	cacheMutex.RLock()
	cachedData, found := translationsCache[languageCode]
	cacheMutex.RUnlock()

	if found {
		return cachedData
	}

	content, err := os.ReadFile("resources/i18n.json")
	if err != nil {
		log.Fatal(err)
	}
	jsonData := string(content)

	var data map[string]map[string]string

	err = json.Unmarshal([]byte(jsonData), &data)
	if err == nil {
		rData := make(map[string]string)
		for k, v := range data {
			value, exists := v[languageCode]
			if exists {
				rData[k] = value
			} else {
				rData[k] = v["en"]
			}
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
