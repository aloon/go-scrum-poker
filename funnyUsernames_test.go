package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunnyUsernames(t *testing.T) {
	results := make([]string, 0)
	distinctMap := make(map[string]bool)

	for i := 0; i < 3; i++ {
		results = append(results, generateFunnyUsername())
	}

	assert.NotEmpty(t, results[0])

	for _, item := range results {
		if distinctMap[item] {
			t.Errorf("A duplicate result was found")
		}
		distinctMap[item] = true
	}

}
