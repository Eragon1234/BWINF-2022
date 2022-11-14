package fahrradwerkstatt

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Auftrag struct {
	Eingangszeitpunkt        int // der Eingangszeitpunkt des Auftrags in Minuten
	Bearbeitungsdauer        int // die Bearbeitungsdauer des Auftrags in Minuten
	Restdauer                int // die Restdauer des Auftrags am aktuellen Zeitpunkt
	Fertigstellungszeitpunkt int // der Fertigstellungszeitpunkt des Auftrags in Minuten
}

func AuftraegeAusDateiLesen(filename string) []*Auftrag {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var auftraege []*Auftrag
	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), " ")
		if len(line) != 2 {
			continue
		}
		eingangszeitpunkt, _ := strconv.Atoi(line[0])
		bearbeitungsdauer, _ := strconv.Atoi(line[1])

		auftraege = append(auftraege, &Auftrag{
			Eingangszeitpunkt: eingangszeitpunkt,
			Bearbeitungsdauer: bearbeitungsdauer,
		})
	}

	return auftraege
}
