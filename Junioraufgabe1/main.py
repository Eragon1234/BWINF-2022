import sys

if not sys.version_info >= (3, 9):
    sys.exit("Benötigt python 3.9")

from rhyme.print_rhymes import print_rhymes
from rhyme.read_words_from_file import read_words_from_file
from rhyme.rhyming_words import rhyming_words


def main():
    if len(sys.argv) < 2:
        sys.exit("missing filename")
    filename = sys.argv[1]
    words = read_words_from_file(filename)
    rhyme_words = rhyming_words(words)
    print_rhymes(rhyme_words)


if __name__ == '__main__':
    main()
