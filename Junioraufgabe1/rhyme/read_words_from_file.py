def read_words_from_file(filename: str) -> list[str]:
    with open(filename, "r") as file:
        content = file.read()

    words = content.split("\n")
    return words
