package fahrradwerkstatt

import (
	"Aufgabe4/Quelltext/utils"
	"fmt"
	"sort"
	"time"
)

func Fahrradwerkstatt(auftraege []*Auftrag) {
	shortestResults := Simulation(utils.CopyPointerSlice(auftraege), shortestAuftrag)
	aeltestResults := Simulation(utils.CopyPointerSlice(auftraege), aeltesterAuftrag)
	highestRatioResults := Simulation(utils.CopyPointerSlice(auftraege), highestRatioAuftrag)

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
