import json
import string

def normalize_alphabetize(card):
    return sorted([letter for letter in card.strip().lower() if letter in string.ascii_lowercase])

def is_isogram(card):
    if len(card) < 1:
        return True
    i = 0
    while i < len(card)-1:
        curr = card[i]
        next = card[i+1]
        if curr == next:
            return False
        i = i+1
    return True


top_x = 3
isograms = []

with open("all-cards.txt", "r") as list:
    cards = json.loads(list.read())
    for card in cards:
        ascii_alphabetized = normalize_alphabetize(card)
        result = is_isogram(ascii_alphabetized)
        if result:
            isograms.append(card)

with open("isograms.txt", "w") as output:
    for card in isograms:
        output.write(card)
        output.write("\n")
