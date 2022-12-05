package main

import (
	"Aufgabe1/search"
	_ "embed"
	"flag"
	"fmt"
	"os"
)

//go:embed Alice_im_Wunderland.txt
var txt string

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

	sentences := search.SearchSentencesWithSearchFile(txt, filename)
	printSentences(sentences)
}

func printSentences(sentences []string) {
	for _, sentence := range sentences {
		fmt.Println(sentence)
	}
}
