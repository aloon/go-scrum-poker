package main

import "testing"

func TestFunnyUsernames(t *testing.T) {
	results := make([]string, 0)
	distinctMap := make(map[string]bool)

	for i := 0; i < 3; i++ {
		results = append(results, generateFunnyUsername())
	}

	if results[0] == "" {
		t.Errorf("A username was not generated")
	}

	for _, item := range results {
		if distinctMap[item] {
			t.Errorf("A duplicate result was found")
		}
		distinctMap[item] = true
	}

}
