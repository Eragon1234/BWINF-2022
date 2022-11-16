package search

import (
	"log"
	"os"
)

func SearchSentencesWithSearchFile(txt string, filename string) []string {
	searchBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return SearchSentences(txt, string(searchBytes))
}

func SearchSentences(txt string, search string) []string {
	return []string{}
}
