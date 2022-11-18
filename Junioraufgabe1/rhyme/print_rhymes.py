def print_rhymes(rhymes: dict[str, dict[str, list]]):
    for mass_vokal_group in rhymes:
        for chars_after_vocal_group in rhymes[mass_vokal_group]:
            words = rhymes[mass_vokal_group][chars_after_vocal_group]

            if len(words) > 1:
                print(*words, sep=", ")
