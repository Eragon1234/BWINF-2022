package main

import (
	"Aufgabe4/fahrradwerkstatt"
	"flag"
	"os"
)

var filename string

func init() {
	flag.StringVar(&filename, "f", "", "Name der Datei mit der Liste der Aufträge")
	flag.StringVar(&filename, "file", "", "Name der Datei mit der Liste der Aufträge")
	flag.StringVar(&filename, "filename", "", "Name der Datei mit der Liste der Aufträge")
	flag.Parse()
}

func main() {
	if filename == "" && len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		flag.PrintDefaults()
		return
	}

	auftraege, err := fahrradwerkstatt.AuftraegeAusDateiLesen(filename)
	if err != nil {
		panic(err)
	}

	fahrradwerkstatt.Fahrradwerkstatt(auftraege)
}
