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
                    if word.endswith(current_word.lower()):
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
