"""
You're given a dictionary of people and the number of games they've won.

Use a for loop to iterate over the dictionary and print out the users name and how many games they've won in the
following format: sara has won n games

To make it human readable, pluralise the word game to suit the number of games won.
"""

games_won = dict(sara=0, bob=1, tim=5, julian=3, jim=1)


def print_game_stats(games_won=games_won):
    """Loop through games_won's dict (key, value) pairs (dict.items)
       printing (print, not return) how many games each person has won,
       pluralize 'game' based on number.

       Expected output (ignore the docstring's indentation):

       sara has won 0 games
       bob has won 1 game
       tim has won 5 games
       julian has won 3 games
       jim has won 1 game

       (Note that as of Python 3.7 - which we're using atm - dict insert order is retained
        so no sorting is required for this Bite.)
    """
    for key in games_won:
        plural = "game"
        if games_won[key] != 1:
            plural = "games"
        print("{0} has won {1} {2}".format(key, games_won[key], plural))

print_game_stats(games_won)