package search

import (
	"Aufgabe1/utils"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// SearchSentencesWithSearchFile loads search from filename and calls SearchSentences
func SearchSentencesWithSearchFile(txt, filename string) []string {
	searchBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return SearchSentences(txt, string(searchBytes))
}

// SearchSentences searches for a sentence in txt that matches the search string (gaps are underscores).
// returns all found applicable sentences
func SearchSentences(txt, search string) []string {
	txt = removeSpecialCharacters(txt)

	words := strings.Split(txt, " ")
	wordPositions := make(map[string][]int)
	for i, word := range words {
		word = strings.ToLower(word)
		wordPositions[word] = append(wordPositions[word], i)
	}

	searchWords := strings.Split(search, " ")
	var startIndexes []int
	for i, searchWord := range searchWords {
		if searchWord == "_" {
			continue
		}

		wordPosition := wordPositions[searchWord]

		if i == 0 {
			startIndexes = append(startIndexes, wordPosition...)
			continue
		}

		startIndexes = filterPossibleStartIndexes(startIndexes, wordPosition, i)
	}

	searchLength := len(searchWords)
	var found []string
	for _, startIndex := range startIndexes {
		endIndex := startIndex + searchLength

		found = append(found, strings.Join(words[startIndex:endIndex], " "))
	}

	return found
}

func filterPossibleStartIndexes(possibleStartIndexes, nextIndexes []int, gaps int) []int {
	var newPossibleStartIndexes []int
	for _, possibleStartIndex := range possibleStartIndexes {
		if utils.IndexOf(nextIndexes, possibleStartIndex+gaps) != -1 {
			newPossibleStartIndexes = append(newPossibleStartIndexes, possibleStartIndex)
		}
	}
	return newPossibleStartIndexes
}

func removeSpecialCharacters(txt string) string {
	txt = strings.ReplaceAll(txt, "\n", " ")
	reg := regexp.MustCompile("[^\\w äöüß]")
	return reg.ReplaceAllString(txt, "")
}

// RegexSearchSentences searches for a sentence in txt that matches the search string (gaps are underscores) by using a regular expression.
// returns all found applicable sentences
func RegexSearchSentences(txt, search string) []string {
	expression := fmt.Sprintf(`(?mi)%s`, strings.ReplaceAll(search, "_", "\\w+"))
	re := regexp.MustCompile(expression)
	return re.FindAllString(txt, -1)
}
