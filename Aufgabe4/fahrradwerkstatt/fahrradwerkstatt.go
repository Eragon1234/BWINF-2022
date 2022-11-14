package fahrradwerkstatt

import (
	"Aufgabe4/utils"
	"fmt"
	"sort"
	"time"
)

func Fahrradwerkstatt(auftraege []*Auftrag) {
	var auftraege2 []*Auftrag
	for _, auftrag := range auftraege {
		auftrag2 := *auftrag
		auftraege2 = append(auftraege2, &auftrag2)
	}
	fmt.Println("Shortest:")
	fmt.Println(Simulation(auftraege, shortestAuftrag))
	fmt.Println("Ältester:")
	fmt.Println(Simulation(auftraege2, aeltesterAuftrag))
}

// Simulation erledigt sämtliche Aufträge und wählt den nächsten Auftrag mit der nextOrder function aus
// Gibt die durchschnittliche und die maximale Wartezeit zurück.
func Simulation(auftraege []*Auftrag, nextOrder func(auftraege []*Auftrag) *Auftrag) (float64, int) {
	var currentTime time.Time // aktueller Zeitpunkt in Minuten
	currentTime = currentTime.Add(time.Hour * 9)
	auftraege = sortArrayByEingangszeitpunkt(auftraege)
	var done []*Auftrag
	for len(auftraege) > 0 {
		aktuelleAuftraege := filterCurrentAvailable(auftraege, int(timeSinceBegin(currentTime).Minutes()))
		if len(aktuelleAuftraege) == 0 {
			currentTime = time.Time{}.Add(time.Minute * time.Duration(auftraege[0].Eingangszeitpunkt))
			continue
		}

		aktuellerAuftrag := nextOrder(aktuelleAuftraege)
		for aktuellerAuftrag.Restdauer > 0 {
			zeitBisFeierabend := durationTillClosingTime(currentTime)
			if float64(aktuellerAuftrag.Restdauer) <= zeitBisFeierabend.Minutes() {
				currentTime = currentTime.Add(time.Minute * time.Duration(aktuellerAuftrag.Bearbeitungsdauer))
				aktuellerAuftrag.Restdauer = 0
				break
			}
			aktuellerAuftrag.Restdauer -= int(zeitBisFeierabend.Minutes())
			currentTime = currentTime.Add(time.Hour * 16).Add(zeitBisFeierabend)
		}

		aktuellerAuftrag.Fertigstellungszeitpunkt = int(timeSinceBegin(currentTime).Minutes())

		i := utils.IndexOf(auftraege, aktuellerAuftrag)

		auftraege = utils.Remove(auftraege, i)
		done = append(done, aktuellerAuftrag)
	}

	return durchschnittsDauer(done), maxDauer(done)
}

func shortestAuftrag(auftraege []*Auftrag) *Auftrag {
	auftraege = sortArrayByShortest(auftraege)
	return auftraege[0]
}

func aeltesterAuftrag(auftraege []*Auftrag) *Auftrag {
	auftraege = sortArrayByEingangszeitpunkt(auftraege)
	return auftraege[0]
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

func filterCurrentAvailable(auftraege []*Auftrag, currentTime int) []*Auftrag {
	var aktuelleAuftraege []*Auftrag
	for _, auftrag := range auftraege {
		if auftrag.Eingangszeitpunkt <= currentTime {
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

func timeSinceBegin(currentTime time.Time) time.Duration {
	return currentTime.Sub(time.Time{})
}
