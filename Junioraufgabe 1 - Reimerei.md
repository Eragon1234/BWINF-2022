# Junioraufgabe 1: Reimerei

Team-ID: 00829

Team: Supercooles neues Team

Bearbeiter dieser Aufgabe:

    - Alexander Gerstberger
    - Arved Lieker
    - Oleg Klimenko

19.November 2022

## Inhaltsverzeichnis

Lösungsidee………………………………………………………………………………..….1

Umsetzung…………………………………………………………………………………..…1

Beispiele………………………………………………………………………………………...2

Quellcode……………………………………………………………………………………….3

## Lösungsidee

Wir berechnen für jedes Wort die maßgebliche Vokalgruppe und die Buchstaben nach der Vokalgruppe. Wir überprüfen dann ob die maßgebliche Vokalgruppe mit den Buchstaben danach mindestens die Hälfte der Buchstaben ausmacht. Wenn sie die Hälfte der Buchstaben ausmachen fügen wir das Wort in eine Datenstruktur unter dem Schlüssel maßgebliche Vokalgruppe und dann unter dem Schlüssel Buchstaben danach in eine Liste ein. Damit sortieren sich die Wörter selber in reimende Wörter ein. Zum Schluss müssen wir alle Wörter,  die auf andere Wörter in ihrer Reimgruppe enden rausfiltern.

## Umsetzung

Die Lösungsidee implementieren wir in Python.

Die ganzen Wörter lesen wir in eine Liste ein.

Zum Speichern der Reime erstellen wir das Dictionary "rhymes".

Dann loopen wir durch alle Wörter.

Wir berechenen die maßgebliche Vokalgruppe, indem wir uns alle Vokalgruppen mit einer Regular Expression die alle Vokale sucht. Dass gibt eine Liste aller Vokalgruppen zurück und wir holen uns die Vorletze davon.

Wir berechnen die Position der letzen Vokalgruppe.

Wir überprüfen ob "Länge der maßgeblichen Vokalgruppe" + "Länge der Buchstaben danach" größer oder gleich der Länge des Wortes geteilt durch 2 ist.

Wenn es das Wort nicht ist fahren wir mit dem nächsten fort.

Wir speichern das Wort in einer Datenstruktur im Format Schlüssel und noch ein Schlüssel und dann eine Liste. Wir hängen das Wort in die Liste am ersten Schlüssel von „maßgebliche Vokalgruppe“ und am zweiten Schlüssel „Buchstaben danach“ an.

Wenn wir mit allen Wörtern fertig sind filtern wir alle Wörter pro Reimgruppe, dass sie nicht auf ein anderes Wort in der Vokalgruppe enden. Das tun wir indem wir in jeder Reimgruppe für jedes Wort mit der "str.endswith" Funktion von Python überprüfen ob es mit einem anderen Wort endet und wenn es das tut entfernen wir es.

## Beispiele 

Wir rufen das Programm mit den BWINF-Beispieleingaben auf. Die Dateien befinden sich in dem Ordner Beispieleingaben

```shell
$ python main.py Beispieleingaben/reimerei0.txt
bemuehen, gluehen
Biene, Hygiene, Schiene
Knecht, Recht
```

```shell
$ python main.py Beispieleingaben/reimerei1.txt
Brote, Note
Bildnis, Wildnis
```

```shell
$ python main.py Beispieleingaben/reimerei2.txt
Epsilon, Ypsilon
```

```shell
$ python main.py Beispieleingaben/reimerei3.txt
Absender, Kalender
Dezember, November, September
Fest, Test
Kassette, Kette, Toilette
Keller, Teller
Ansage, Frage
Bahn, Zahn
Bank, Dank
Dame, Name
Flasche, Tasche
Gas, Glas
Hand, Land, Strand
Kanne, Panne
Kasse, Klasse, Tasse
Magen, Wagen
Platz, Satz
Sache, Sprache
Butter, Mutter
Drucker, Zucker
Durst, Wurst
Fuß, Gruß
```

```go
Gruppe, Suppe
Hund, Mund
Arbeit, Zeit
Bein, Wein
Gleis, Reis
Bild, Schild
Bitte, Mitte
Fisch, Tisch
Kind, Rind, Wind
Export, Import, Wort
Hose, Rose
Kopf, Topf
Lohn, Sohn
Rock, Stock
Glueck, Stueck
Baum, Raum
Bier, Tier
Feuer, Steuer
Idee, See, Tee
Gruppe, Suppe
Hund, Mund
Kunde, Stunde
Arbeit, Zeit
Bein, Wein
Gleis, Reis
Bild, Schild
Bitte, Mitte
Fisch, Tisch
Kind, Rind, Wind
Export, Import, Wort
Hose, Rose
Kopf, Topf
Lohn, Sohn
Rock, Stock
Glueck, Stueck
Baum, Raum
Bier, Tier
Feuer, Steuer
Idee, See, Tee
Kunde, Stunde
```

## 

## 

## 

## 

## Quellcode

#### main.py

```python
import sys

if not sys.version_info >= (3, 9):
    sys.exit("Benötigt python 3.9")

from rhyme.print_rhymes import print_rhymes
from rhyme.read_words_from_file import read_words_from_file
from rhyme.rhyming_words import rhyming_words


def main():
    filename = sys.argv[1]
    if not filename:
        sys.exit("missing filename")
    words = read_words_from_file(filename)
    rhyme_words = rhyming_words(words)
    print_rhymes(rhyme_words)


if __name__ == '__main__':
    main()

```

#### rhyming_words.py

```python
import re
from copy import copy


def rhyming_words(words: list[str]) -> dict[str, dict[str, list]]:
    words = prepare_words(words)
    rhymes = {}
    for word in words:
        mass_vokal = massgebende_vokalgruppe(word)
        if mass_vokal == "":
            continue

        chars_after_vokal = chars_after_vocal(word, mass_vokal)

        if len(word) / 2 > len(mass_vokal) + len(chars_after_vokal):
            continue

        if mass_vokal not in rhymes.keys():
            rhymes[mass_vokal] = {}


```

```python
        if chars_after_vokal not in rhymes[mass_vokal].keys():
            rhymes[mass_vokal][chars_after_vokal] = []

        rhymes[mass_vokal][chars_after_vokal].append(word)

    rhymes = filter_rhymes(rhymes)
    return rhymes

def filter_rhymes(rhymes: dict[str, dict[str, list]]) -> dict[str, dict[str, list]]:
    for mass_vokal_group in rhymes:
        for chars_after_vocal_group in rhymes[mass_vokal_group]:
            words = rhymes[mass_vokal_group][chars_after_vocal_group]

            for word in copy(words):
                current_words = list(filter(lambda w: w != word, words))
                for current_word in current_words:
                    if word.lower().endswith(current_word.lower()):
                        words.remove(word)
    return rhymes


def massgebende_vokalgruppe(word: str):
    vokalgruppen = re.findall("[aeiou]+", word)
    if len(vokalgruppen) == 0:
        return ""
    if len(vokalgruppen) == 1:
        return vokalgruppen[0]
    return vokalgruppen[-2]

def chars_after_vocal(word: str, mass_vokal: str) -> str:
    if word.count(mass_vokal) == 1:
        return word.split(mass_vokal)[-1]
    i_last = word.rindex(mass_vokal)
    i_sec_last = word.rindex(mass_vokal, None, i_last)
    return word[i_sec_last + 1:]


def prepare_words(words: list[str]) -> list[str]:
    return [word.replace("ö", "oe").replace("ä", "ae").replace("ü", "ue") for word in words]


def sort_rhymes(rhymes: dict[str, dict[str, list]]) -> dict[str, dict[str, list]]:
    for mass_vokal_group in rhymes:
        for chars_after_vocal_group in rhymes[mass_vokal_group]:
            rhymes[mass_vokal_group][chars_after_vocal_group].sort()
    return rhymes
```