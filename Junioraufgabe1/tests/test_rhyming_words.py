from unittest import TestCase
from Junioraufgabe1.rhyme.rhyming_words import rhyming_words, chars_after_vocal, massgebende_vokalgruppe
from Junioraufgabe1.rhyme.read_words_from_file import read_words_from_file


class Test(TestCase):
    def test_rhyming_words_reimerei0(self):
        words = read_words_from_file("../Beispieleingaben/reimerei0.txt")
        rhymes = rhyming_words(words)

        self.assertCountEqual({
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
        }, rhymes)

    def test_rhyming_words_reimerei1(self):
        words = read_words_from_file("../Beispieleingaben/reimerei1.txt")
        rhymes = rhyming_words(words)

        self.assertCountEqual({
            "o": {
                "rption": [
                    "absorption"
                ],
                "te": [
                    "note",
                    "brote"
                ]
            },
            "u": {
                "nktion": {
                    "xor-funktion"
                },
                "mption": {
                    "konsumption"
                }
            },
            "i": {
                "ldnis": [
                    "bildnis",
                    "wildnis"
                ],
            },
            "ae": {
                "ndnis": [
                    "gestaendnis",
                ]
            },
        }, rhymes)

    def test_rhyming_words_reimerei2(self):
        words = read_words_from_file("../Beispieleingaben/reimerei2.txt")
        rhymes = rhyming_words(words)

        self.assertCountEqual({
            "i": {
                "lon": [
                    "epsilon",
                    "ypsilon"

                ]
            },
            "o": {
                "to": [
                    "foto"
                ]
            },
            "e": {
                "mpel": [
                    "tempel"
                ]
            }
        }, rhymes)


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
            self.assertEqual(chars_after_vocal(**test["args"]), test["want"], test["name"])


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
            self.assertEqual(massgebende_vokalgruppe(**test["args"]), test["want"], test["name"])