from Junioraufgabe1.rhyme.read_words_from_file import read_words_from_file
from Junioraufgabe1.rhyme.rhyming_words import rhyming_words


def main():
    words = read_words_from_file("Beispieleingaben/reimerei3.txt")
    print(rhyming_words(words))


if __name__ == '__main__':
    main()
