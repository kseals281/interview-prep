"""
Write a function to determine if a word (or phrase) is a palindrome.

Then write a second function to receive a word (or phrase) list and determine which word is the longest palindrome.

See the template code / docstrings below for further requirements as well as the tests

A palindrome is a word, phrase, number, or other sequence of characters
which reads the same backward as forward"""
import string

DICTIONARY = "/home/kseals/go/src/interview-prep/setup/palindrome_dictionary.txt"


def load_dictionary():
    """Load dictionary (sample) and return as generator (done)"""
    with open(DICTIONARY, 'r') as f:
        return (word.lower().strip() for word in f.readlines())


def is_palindrome(word):
    """
    Return if word is palindrome, 'madam' would be one.
    Case insensitive, so Madam is valid too.
    It should work for phrases too so strip all but alphanumeric chars.
    So "No 'x' in 'Nixon'" should pass (see tests for more)
    """
    first_half = ''
    second_half = ''
    clean_word = sanitize(word)
    if len(clean_word) % 2 == 0:
        first_half = clean_word[:int(len(clean_word) / 2)]
        second_half = clean_word[int(len(clean_word) / 2):]
    else:
        first_half = clean_word[:int(len(clean_word) / 2)]
        second_half = clean_word[int(len(clean_word) / 2) + 1:]
    reversed_half = ''
    for letter in second_half:
        reversed_half = letter + reversed_half
    return first_half == reversed_half


def sanitize(word):
    """Takes in a word and returns the sanitized (lowercase alphanumeric) version"""
    word = word.lower()
    clean_word = ''
    for letter in word:
        if letter in string.ascii_lowercase:
            clean_word += letter
    return clean_word


def get_longest_palindrome(words=None):
    """Given a list of words return the longest palindrome
       If called without argument use the load_dictionary helper
       to populate the words list"""
    if words == None:
        words = load_dictionary()
    longest = ''
    for word in words:
        if len(word) < len(longest):
            continue
        if is_palindrome(word):
            longest = word
    return longest
