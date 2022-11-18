from unittest import TestCase

from Junioraufgabe1.rhyme.read_words_from_file import read_words_from_file
from Junioraufgabe1.rhyme.rhyming_words import rhyming_words, chars_after_vocal, massgebende_vokalgruppe, sort_rhymes, \
    filter_rhymes


class TestRhymingWords(TestCase):
    def test_rhyming_words_reimerei0(self):
        words = read_words_from_file("../Beispieleingaben/reimerei0.txt")
        rhymes = rhyming_words(words)

        self.assertEquals(sort_rhymes({
            "ie": {
                "ne": [
                    "Biene",
                    "Schiene",
                    "Hygiene"
                ]
            },
            "ue": {
                "hen": [
                    "bemuehen",
                    "gluehen"
                ]
            },
            "a": {
                "gen": [
                    "hersagen"
                ]
            },
            "e": {
                "cht": [
                    "Knecht",
                    "Recht"
                ]
            }
        }), sort_rhymes(rhymes))

    def test_rhyming_words_reimerei1(self):
        words = read_words_from_file("../Beispieleingaben/reimerei1.txt")
        rhymes = rhyming_words(words)

        self.assertEquals(sort_rhymes({
            "o": {
                "rption": [
                    "Absorption"
                ],
                "te": [
                    "Note",
                    "Brote"
                ]
            },
            "u": {
                "nktion": [
                    "XOR-Funktion"
                ],
                "mption": [
                    "Konsumption"
                ]
            },
            "i": {
                "ldnis": [
                    "Bildnis",
                    "Wildnis"
                ],
            },
            "ae": {
                "ndnis": [
                    "Gestaendnis",
                ]
            },
        }), sort_rhymes(rhymes))

    def test_rhyming_words_reimerei2(self):
        words = read_words_from_file("../Beispieleingaben/reimerei2.txt")
        rhymes = rhyming_words(words)

        self.assertEquals(sort_rhymes({
            "i": {
                "lon": [
                    "Epsilon",
                    "Ypsilon"

                ]
            },
            "o": {
                "to": [
                    "Foto"
                ]
            },
            "e": {
                "mpel": [
                    "Tempel"
                ]
            }
        }), sort_rhymes(rhymes))


class TestCharsAfterVocal(TestCase):
    def test_chars_after_vocal(self):
        tests = [
            {
                "name": "testing basic",
                "args": {
                    "word": "Biene",
                    "mass_vokal": "ie"
                },
                "want": "ne"
            },
            {
                "name": "testing word with the mass_vokal two times",
                "args": {
                    "word": "Foto",
                    "mass_vokal": "o"
                },
                "want": "to"
            }
        ]

        for test in tests:
            self.assertEquals(test["want"], chars_after_vocal(**test["args"]), test["name"])


class TestMassgebendeVokalgruppe(TestCase):
    def test_massgebende_vokalgruppe(self):
        tests = [
            {
                "name": "testing basic",
                "args": {
                    "word": "Biene"
                },
                "want": "ie"
            },
            {
                "name": "testing word with only one vocal group",
                "args": {
                    "word": "Knecht",
                },
                "want": "e"
            }
        ]

        for test in tests:
            self.assertEquals(test["want"], massgebende_vokalgruppe(**test["args"]), test["name"])


class TestFilterRhymes(TestCase):
    def test_filter_rhymes(self):
        test_words = {
            "x": {
                "x": [
                    "Autobahn",
                    "Bahn"

                ]
            },
            "y": {
                "x": [
                    "Parkhaus",
                    "Kaufhaus",
                    "Haus"
                ]
            }
        }

        result = filter_rhymes(test_words)

        self.assertEquals({
            "x": {
                "x": [
                    "Bahn"
                ]
            },
            "y": {
                "x": [
                    "Haus"
                ]
            }
        }, result)
