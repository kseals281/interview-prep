"""
Calculate the dictionary word that would have the most value in Scrabble.

There are 3 tasks to complete for this Bite:

First write a function to read in the scrabble_dictionary.txt file (= DICTIONARY constant), returning a list of words
(note that the words are separated by new lines).

Second write a function that receives a word and calculates its value. Use the scores stored in LETTER_SCORES.
Letters that are not in LETTER_SCORES should be omitted (= get a 0 score).

With these two pieces in place, write a third function that takes a list of words and returns the word with the highest
value.

Look at the TESTS tab to see what your code needs to pass. Enjoy!
"""
import os

# PREWORK
DICTIONARY = '/Users/khari/PycharmProjects/interview-prep/python/scrabble_dictionary.txt'
scrabble_scores = [(1, "E A O I N R T L S U"), (2, "D G"), (3, "B C M P"),
                   (4, "F H V W Y"), (5, "K"), (8, "J X"), (10, "Q Z")]
LETTER_SCORES = {letter: score for score, letters in scrabble_scores
                 for letter in letters.split()}


# start coding

def load_words():
    """Load the words dictionary (DICTIONARY constant) into a list and return it"""
    f = open(DICTIONARY, 'r')
    words = []
    for line in f:
        words.append(line.strip('\n'))
    return words


def calc_word_value(word):
    """Given a word calculate its value using the LETTER_SCORES dict"""
    score = 0
    for letter in word:
        letter = letter.upper()
        if letter in LETTER_SCORES:
            score += LETTER_SCORES[letter]
    return score


def max_word_value(words):
    """Given a list of words calculate the word with the maximum value and return it"""
    mx = (0, '')
    for word in words:
        score = calc_word_value(word)
        if score > mx[0]:
            mx = (score, word)
    return mx[1]

calc_word_value('bob')