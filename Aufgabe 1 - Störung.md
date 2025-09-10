# Aufgabe 1: Störung

Team-ID: 00829

Team: Supercooles neues Team

Bearbeiter dieser Aufgabe:

    - Alexander Gerstberger
    - Arved Lieker
    - Oleg Klimenko

19. November 2022

## Inhaltsverzeichnis

Lösungsidee………………………………………………………………………………..….1

Umsetzung…………………………………………………………………………………..…1

Beispiele………………………………………………………………………………………...2

Quellcode……………………………………………………………………………………….2

## Lösungsidee

Zuerst erstellen wir eine Liste mit allen Wörtern und allen Positionen an denen die Wörter vorkommen. Daraufhin fügen wir alle möglichen Positionen für das erste Wort in eine Liste hinzu. Danach zählen wir die Anzahl der Wörter seit dem erstem Wort und filtern alle Positionen für das erste Wort und schauen, ob das gesuchte nächste Wort an der Stelle Position erstes Wort + Anzahl der Wörter beträgt. Für alle verbleibenden Startpositionen holen wir uns Startposition bis Startposition + Länge des Wortes.

## Umsetzung

Wir lesen die gesamte Alice_im_Wunderland.txt Datei in einen String. Dann entfernen wir alle nicht benötigten Zeichen.

Danach erstellen wir eine map namens "wordPositions" . Jedes Wort soll als Schlüssel in diese map kopiert werden und die jewaligen Stellen an denen das Wort vorkommt sollen als Wert für den Schlüssel genommen werden.

Wir erstellen eine Variable "searchWords" mit dem Wert von unserem Such String, der an allen Leerzeichen geteilt wird. Damit erhalten wir eine Liste aller Wörter.

Daraus erstellen wir eine Liste von Integern, die wir "startIndexes" genannt haben.

Dann loopen wir über alle Wörter in "searchWords".

Beim ersten Durchgang fügen wir alle Positionen des Wortes aus "wordPositions" zu "startIndexes" hinzu.

Nach jedem Durchgang filtern wir ob an Stelle "Startindex" + "Anzahl der Wörter seit Startindex" das aktuelle Wort vorkommt. Wenn das nicht der Fall ist, trifft unsere Suche nicht auf den Startindex zu.

Danach suchen wir alle Sätze heraus indem wir uns die Wörter für jede Startposition von "Startposition bis Startposition" + "Länge des Satzes" holen.

## Beispiele

Wir rufen das Programm mit den BWINF-Beispieleingaben auf. Die ausführbaren Dateien befinden sich in einem Ordner AusfuehrbaresProgramm. Die Beispieleingabedateien befinden sich in einen Ordner Beispielaufgaben.

```shell
$ ./AusfuehrbaresProgramm/darwin-arm64 Beispielaufgaben/stoerung0.txt
Das kommt mir gar nicht richtig vor
```

```shell
$ ./AusfuehrbaresProgramm/darwin-arm64 Beispielaufgaben/stoerung1.txt
Ich muß in Clara verwandelt
Ich muß doch Clara sein
```

```shell
$ ./AusfuehrbaresProgramm/darwin-arm64 Beispielaufgaben/stoerung2.txt
Fressen Katzen gern Spatzen
Fressen Katzen gern Spatzen
Fressen Spatzen gern Katzen
```

Bei dem vorigen Beispiel kommt der eine Satz doppelt heraus, weil er im Buch zweimal vorkommt. Ob dies gewünscht ist, konnte man an der Aufgabenstellung nicht erkennen. 

```shell
$ ./AusfuehrbaresProgramm/darwin-arm64 Beispielaufgaben/stoerung3.txt
das Spiel fing an
Das Publikum fing an
```

```shell
$ ./AusfuehrbaresProgramm/darwin-arm64 Beispielaufgaben/stoerung4.txt
ein sehr schöner Tag
```

```shell
$ ./AusfuehrbaresProgramm/darwin-arm64 Beispielaufgaben/stoerung5.txt
Wollen Sie so gut sein
```

## 

## 

## 

## 

## Quellcode

#### main.go

```go
package main

import (
	"Aufgabe1/search"
	"flag"
	"fmt"
	"log"
	"os"
)

var filename string

func init() {
	flag.StringVar(&filename, "f", "", "Name der Datei mit dem Lückensatz")
	flag.StringVar(&filename, "file", "", "Name der Datei mit dem Lückensatz")
	flag.StringVar(&filename, "filename", "", "Name der Datei mit dem Lückensatz")
	flag.Parse()
}

func main() {
	if filename == "" && len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		flag.PrintDefaults()
		return
	}

	txt, err := os.ReadFile("Alice_im_Wunderland.txt")
	if err != nil {
		log.Fatal(err)
	}

	sentences := search.SearchSentencesWithSearchFile(string(txt), filename)
	printSentences(sentences)
}

func printSentences(sentences []string) {
	for _, sentence := range sentences {
		fmt.Println(sentence)
	}
}

```

#### search.go

```go
package search

import (
	"Aufgabe1/utils"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func SearchSentencesWithSearchFile(txt, filename string) []string {
	searchBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return SearchSentences(txt, string(searchBytes))
}

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

func RegexSearchSentences(txt, search string) []string {
	expression := fmt.Sprintf(`(?mi)%s`, strings.ReplaceAll(search, "_", "\\w+"))
	re := regexp.MustCompile(expression)
	return re.FindAllString(txt, -1)
}
```