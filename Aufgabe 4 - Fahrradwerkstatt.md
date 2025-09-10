# Aufgabe 4: Fahrradwerkstatt

Team-ID: 00829

Team: Supercooles neues Team

Bearbeiter dieser Aufgabe:

    - Alexander Gerstberger
    - Oleg Klimenko
    - Arved Lieker

19. November 2022



## Inhaltsverzeichnis

Lösungsidee………………………………………………………………………………..….1

Umsetzung…………………………………………………………………………………..…1

Beispiele………………………………………………………………………………………...2

Quellcode……………………………………………………………………………………….4



## Lösungsidee

Wir verfolgen die aktuelle Zeit und loopen solange bis alle Aufgaben erledigt sind. Dafür sortieren wir die in aktueller Zeit verfügbaren Aufträge nach nur aktuell verfügbaren Aufträgen. Falls es keine gibt, setzen wir die Zeit auf den Eingangszeitpunkt des Auftrags mit dem frühsten Eingangszeitpunkt. Der nächste Auftrag wird basierend auf der aktuellen Bestimmungsmethode geholt. Danach loopt man, bis die Restzeit der Aufgabe 0 beträgt. Für den Auftrag berechnen wir die Zeit bis zum Feierabend und addieren dazu die Auftragsszeit.  Falls es zu lange dauert und über den Feierabend hinauskommt, wird ab dem nächsten Arbeitstag ab 9:00 Uhr weitergerechnet

## Umsetzung

Wir erstellen eine Funktion Simulation, welche die Liste aus Aufträgen namens "auftraege" und eine Funktion "nextOrder", welche aus einer Liste aus Auftraegen den nächsten Auftrag bestimmt, als Parameter erhält.

Wir erstellen  die Variable "currentTime" um immer den aktuellen Zeitpunkt im Blick zu behalten.

Dann erstellen wir eine leere Liste für Aufträge ("done") um alle erledigten Aufträge abzuspeichern.

Danach loopen wir solange bis wir alle Aufträge abgeschlossen haben.

In jeder Schleife filtern wir die Liste der Aufträge in eine neue Liste, nach nur zur "currentTime" verfügbaren Aufträgen.

Darauf lassen wir uns von der "nextOrder" Funktion die als nächstes zu erledigende Funktion geben.

Dann loopen wir solange bis die Restzeit des Auftrags 0 beträgt.

In jedem Durchlauf berchnen wir auf Basis der "currentTime" die Zeit bis Feierabend.

Falls die Restzeit geringer ist als die Zeit bis Feierabend, dann setzen wir die aktuelle Zeit auf "currentTime" + "Restzeit des Auftrags" und die Restdauer auf 0.

Wenn nicht, dann ziehen wir die Zeit bis Feierabend von der Restzeit ab und setzen die aktuelle Zeit auf 9:00 des nächsten Tages, also Arbeitsbeginn.

Dann entfernen wir denn Auftrag aus unserer Liste der offenen Aufträge und fügen ihn zur Liste der erledigten Aufträge hinzu.

Um dann unsere Fahrradwerkstatt zu simulieren, rufen wir diese Funktion mit unseren Aufträgen und auch unsere Funktion zum Berechnen des als nächstes zu erledigenden Auftrags. Das gibt uns dann alle erledigten Aufträge zurück, auf dessen Basis wir dann die durchschnittliche Wartezeit, maximale Wartezeit und das von uns hinzugefügte Wartezeit zu Auftragsdauer Verhältnis berechnet werden kann.

Zur Berechnung des nächsten Auftrags haben wir drei Funktionen erstellt, eine die den ältesten Auftrag zurückgibt, eine die den kürzesten Auftrag zurückgibt und einen, der den mit dem bisher höchsten Wartezeit zu Auftragsdauer Verhältnis zurückgibt.

## Beispiele

Wir rufen das Programm mit den BWINF-Beispieleingaben auf. Die ausführbaren Dateien befinden sich in einem Ordner AusfuehrbaresProgramm. Die Beispieleingabedateien befinden sich in einen Ordner Beispieleingaben.

```go
$ ./AusfuehrbaresProgramm/darwin-arm64 Beispieleingaben/fahrradwerkstatt0.txt
Kürzeste Aufträge zuerst
Durchschnittliche Wartezeit: 16981.28
Maximale Wartezeit: 188734
Durchschnittliches Wartezeit zu Auftragsdauer Verhältnis: 10.95

Älteste Aufträge zuerst
Durchschnittliche Wartezeit: 32753.50
Maximale Wartezeit: 68771
Durchschnittliches Wartezeit zu Auftragsdauer Verhältnis: 21.12

Höchstes Wartezeit zu Auftragsdauer Verhältnis zuerst
Durchschnittliche Wartezeit: 21518.54
Maximale Wartezeit: 93882
Durchschnittliches Wartezeit zu Auftragsdauer Verhältnis: 13.88
```

```go
$ ./AusfuehrbaresProgramm/darwin-arm64 Beispieleingaben/fahrradwerkstatt1.txt
Kürzeste Aufträge zuerst
Durchschnittliche Wartezeit: 11883.92
Maximale Wartezeit: 433563
Durchschnittliches Wartezeit zu Auftragsdauer Verhältnis: 51.42

Älteste Aufträge zuerst
Durchschnittliche Wartezeit: 63535.65
Maximale Wartezeit: 128321
Durchschnittliches Wartezeit zu Auftragsdauer Verhältnis: 274.92

Höchstes Wartezeit zu Auftragsdauer Verhältnis zuerst
Durchschnittliche Wartezeit: 21370.28
Maximale Wartezeit: 248126
Durchschnittliches Wartezeit zu Auftragsdauer Verhältnis: 92.47
```

```go
$ ./AusfuehrbaresProgramm/darwin-arm64 Beispieleingaben/fahrradwerkstatt2.txt
Kürzeste Aufträge zuerst
Durchschnittliche Wartezeit: 14813.53
Maximale Wartezeit: 327087
Durchschnittliches Wartezeit zu Auftragsdauer Verhältnis: 15.69

Älteste Aufträge zuerst
Durchschnittliche Wartezeit: 51194.49
Maximale Wartezeit: 110973
Durchschnittliches Wartezeit zu Auftragsdauer Verhältnis: 54.22

Höchstes Wartezeit zu Auftragsdauer Verhältnis zuerst
Durchschnittliche Wartezeit: 19082.87
Maximale Wartezeit: 135612
Durchschnittliches Wartezeit zu Auftragsdauer Verhältnis: 20.21

```

```go
$ ./AusfuehrbaresProgramm/darwin-arm64 Beispieleingaben/fahrradwerkstatt3.txt
Kürzeste Aufträge zuerst
Durchschnittliche Wartezeit: 17242.83
Maximale Wartezeit: 382016
Durchschnittliches Wartezeit zu Auftragsdauer Verhältnis: 21.62

Älteste Aufträge zuerst
Durchschnittliche Wartezeit: 30028.93
Maximale Wartezeit: 60821
Durchschnittliches Wartezeit zu Auftragsdauer Verhältnis: 37.65

Höchstes Wartezeit zu Auftragsdauer Verhältnis zuerst
Durchschnittliche Wartezeit: 19876.06
Maximale Wartezeit: 66263
Durchschnittliches Wartezeit zu Auftragsdauer Verhältnis: 24.92
```

```go
$ ./AusfuehrbaresProgramm/darwin-arm64 Beispieleingaben/fahrradwerkstatt4.txt
Kürzeste Aufträge zuerst
Durchschnittliche Wartezeit: 42200.90
Maximale Wartezeit: 363155
Durchschnittliches Wartezeit zu Auftragsdauer Verhältnis: 21.16

Älteste Aufträge zuerst
Durchschnittliche Wartezeit: 74427.52
Maximale Wartezeit: 167059
Durchschnittliches Wartezeit zu Auftragsdauer Verhältnis: 37.32

Höchstes Wartezeit zu Auftragsdauer Verhältnis zuerst
Durchschnittliche Wartezeit: 49841.12
Maximale Wartezeit: 193086
Durchschnittliches Wartezeit zu Auftragsdauer Verhältnis: 24.99
```

Man kann an den ganzen Beispielen erkennen, dass die Älteste Auftrag zuerst Methode immer die geringste Maximale Wartezeit hat, aber dafür eine sehr hohe durchschnittliche Wartezeit. Das resultiert auch in einem hohen Wartezeit zu Auftragsdauer Verhältnis.

Die kürzestes Auftrag zuerst Methode hat die geringste durchschnittliche Wartezeit, aber eine sehr hohe maximale Wartezeit. Das Wartezeit zu Auftragsdauer Verhältnis ist bei dieser Methode am geringsten.

Die Wartezeit zu Auftragsdauer Verhältnis Methode hat ein relativ geringe Durchschnittliche Wartezeit in allen Fällen. Und eine relativ geringer Maximale Wartezeit. Das Wartezeit zu Auftragsdauer Verhältnis ist auch relativ gering. Alles in allem ist diesem Methode die Beste Mischung aus geringer maximaler Wartezeit und geringer durchschnittlicher Wartzeit. Damit wären die Kunden im Endeffekt am zufriedensten.

## Quellcode

#### fahrradwerkstatt.go

```go
package fahrradwerkstatt

import (
	"Aufgabe4/utils"
	"fmt"
	"sort"
	"time"
)

func Fahrradwerkstatt(auftraege []*Auftrag) {
	shortestResults := Simulation(utils.CopyPointerSlice(auftraege), shortestAuftrag)
	aeltestResults := Simulation(utils.CopyPointerSlice(auftraege), aeltesterAuftrag)
	highestRatioResults := Simulation(utils.CopyPointerSlice(auftraege), highestRatioAuftrag)

```

```go
	printResult(shortestResults, "Kürzeste Aufträge zuerst")
	printResult(aeltestResults, "Älteste Aufträge zuerst")
	printResult(highestRatioResults, "Höchstes Wartezeit zu Auftragsdauer Verhältnis zuerst")
}

// Simulation erledigt sämtliche Aufträge und wählt den nächsten Auftrag mit der nextOrder function aus
// Gibt die erledigten Aufträge zurück
func Simulation(auftraege []*Auftrag, nextOrder func(auftraege []*Auftrag, currentTime int) *Auftrag) []*Auftrag {
	currentTime := time.Time{}.Add(time.Hour * 9)
	auftraege = sortArrayByEingangszeitpunkt(auftraege)
	var done []*Auftrag
	for len(auftraege) > 0 {
		aktuelleAuftraege := filterCurrentAvailable(auftraege, currentTime)
		if len(aktuelleAuftraege) == 0 {
			currentTime = timeFromMinutes(auftraege[0].Eingangszeitpunkt)
			continue
		}

		aktuellerAuftrag := nextOrder(aktuelleAuftraege, int(timeSinceBegin(currentTime).Minutes()))
		for aktuellerAuftrag.Restdauer > 0 {
			zeitBisFeierabend := durationTillClosingTime(currentTime)
			if float64(aktuellerAuftrag.Restdauer) <= zeitBisFeierabend.Minutes() {
				currentTime = addMinutesToTime(currentTime, aktuellerAuftrag.Restdauer)
				aktuellerAuftrag.Restdauer = 0
				break
			}
			aktuellerAuftrag.Restdauer -= int(zeitBisFeierabend.Minutes())
			currentTime = currentTime.Add(time.Hour * 16).Add(zeitBisFeierabend)
		}

		aktuellerAuftrag.Fertigstellungszeitpunkt = int(timeSinceBegin(currentTime).Minutes())

		i := utils.IndexOf(auftraege, aktuellerAuftrag)

		auftraege, done = utils.RemoveAndAppend(auftraege, done, i)
	}

	return done
}

func printResult(results []*Auftrag, name string) {
	fmt.Println(name)
	fmt.Printf("Durchschnittliche Wartezeit: %.2f\n", durchschnittsDauer(results))
	fmt.Printf("Maximale Wartezeit: %d\n", maxDauer(results))
	fmt.Printf("Durchschnittliches Wartezeit zu Auftragsdauer Verhältnis: %.2f\n", averageWaitToDurationRatio(results))
	fmt.Println()
}

```

```go
func timeFromMinutes(minutes int) time.Time {
	return time.Time{}.Add(time.Minute * time.Duration(minutes))
}

func addMinutesToTime(t time.Time, minutes int) time.Time {
	return t.Add(time.Minute * time.Duration(minutes))
}

func shortestAuftrag(auftraege []*Auftrag, _ int) *Auftrag {
	auftraege = sortArrayByShortest(auftraege)
	return auftraege[0]
}

func aeltesterAuftrag(auftraege []*Auftrag, _ int) *Auftrag {
	auftraege = sortArrayByEingangszeitpunkt(auftraege)
	return auftraege[0]
}

func highestRatioAuftrag(auftraege []*Auftrag, currentTime int) *Auftrag {
	var highestRatioAuftrag *Auftrag
	var highestRatio float64 = -1
	for _, auftrag := range auftraege {
		aktuelleWartezeit := float64(currentTime - auftrag.Eingangszeitpunkt)
		currentRatio := aktuelleWartezeit / float64(auftrag.Bearbeitungsdauer)

		if currentRatio > highestRatio {
			highestRatioAuftrag = auftrag
			highestRatio = currentRatio
		}
	}

	return highestRatioAuftrag
}

func sortArrayByShortest(array []*Auftrag) []*Auftrag {
	sort.Slice(array, func(i, j int) bool {
		return array[i].Bearbeitungsdauer < array[j].Bearbeitungsdauer
	})
	return array
}

func durationTillClosingTime(currentTime time.Time) time.Duration {
	feierabend := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 17, 0, 0, 0, currentTime.Location())
	return feierabend.Sub(currentTime)
}

```

```go
func filterCurrentAvailable(auftraege []*Auftrag, currentTime time.Time) []*Auftrag {
	currentTimeInMinutes := int(timeSinceBegin(currentTime).Minutes())

	var aktuelleAuftraege []*Auftrag
	for _, auftrag := range auftraege {
		if auftrag.Eingangszeitpunkt <= currentTimeInMinutes {
			aktuelleAuftraege = append(aktuelleAuftraege, auftrag)
		}
	}
	return aktuelleAuftraege
}

func sortArrayByEingangszeitpunkt(array []*Auftrag) []*Auftrag {
	sort.Slice(array, func(i, j int) bool {
		return array[i].Eingangszeitpunkt < array[j].Eingangszeitpunkt
	})
	return array
}

func durchschnittsDauer(auftraege []*Auftrag) float64 {
	var dauerSumme float64
	for _, auftrag := range auftraege {
		dauerSumme += float64(auftrag.Fertigstellungszeitpunkt) - float64(auftrag.Eingangszeitpunkt)
	}
	return dauerSumme / float64(len(auftraege))
}

func maxDauer(auftraege []*Auftrag) int {
	var longest int
	for _, auftrag := range auftraege {
		dauer := auftrag.Fertigstellungszeitpunkt - auftrag.Eingangszeitpunkt
		if longest < dauer {
			longest = dauer
		}
	}
	return longest
}

func averageWaitToDurationRatio(auftraege []*Auftrag) float64 {
	var dauerSumme float64
	var durationSumme float64
	for _, auftrag := range auftraege {
		dauerSumme += float64(auftrag.Fertigstellungszeitpunkt) - float64(auftrag.Eingangszeitpunkt)
		durationSumme += float64(auftrag.Bearbeitungsdauer)
	}
	return dauerSumme / durationSumme
}

func timeSinceBegin(currentTime time.Time) time.Duration {
	return currentTime.Sub(time.Time{})
}
```