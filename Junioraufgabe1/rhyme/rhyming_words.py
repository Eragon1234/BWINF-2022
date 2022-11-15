import re


def prepare_words(words: list[str]) -> list[str]:
    return [word.replace("ö", "oe").replace("ä", "ae").replace("ü", "ue") for word in words]


def rhyming_words(words: list[str]) -> dict[str, dict[str, list]]:
    words = prepare_words(words)
    massgebende_vokalgruppe(words[0])


def massgebende_vokalgruppe(word: str):
    vokalgruppen = re.findall("[aeiou]+", word)
    print(vokalgruppen[-2])
