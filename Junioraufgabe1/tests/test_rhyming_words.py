from unittest import TestCase
from ..main import rhyming_words, read_words_from_file


class Test(TestCase):
    def test_rhyming_words_reimerei0(self):
        words = read_words_from_file("../Beispieleingaben/reimerei0.txt")
        rhymes = rhyming_words(words)

        self.assertEqual(rhymes, {
            "ie": {
                "ne": [
                    "biene",
                    "schiene",
                    "hygiene"
                ]
            },
            "ue": {
                "hen": [
                    "bemuehen",
                    "gluehen"
                ]
            },
            "a": {
                "nk": [
                    "schlank",
                    "schwank"
                ],
                "gen": [
                    "hersagen"
                ]
            },
            "e": {
                "cht": [
                    "knecht",
                    "recht"
                ]
            }
        })

    def test_rhyming_words_reimerei1(self):
        words = read_words_from_file("../Beispieleingaben/reimerei1.txt")
        rhymes = rhyming_words(words)

        self.assertEqual(rhymes, {
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
        })

    def test_rhyming_words_reimerei2(self):
        words = read_words_from_file("../Beispieleingaben/reimerei2.txt")
        rhymes = rhyming_words(words)

        self.assertEqual(rhymes, {
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
        })
