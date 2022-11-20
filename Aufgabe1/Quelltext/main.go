package main

import (
	"Aufgabe1/Quelltext/search"
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
