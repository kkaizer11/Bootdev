def vowel_counter(func_to_wrap):
    vowel_count = 0

    def wrapper(doc):
        nonlocal vowel_count
        vowels = "aeiou"
        for char in doc:
            if char in vowels:
                vowel_count += 1
        print(f"Vowel count: {vowel_count}")
        result = func_to_wrap(doc)
        return result

    return wrapper


@vowel_counter
def process_doc(doc):
    print(f"Document: {doc}")


process_doc("Hello")
# Prints:
# Vowel count: 2
# Document: Hello

process_doc("world")
# Prints:
# Vowel count: 3
# Document: world
