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

isograms = dict()

with open("all-cards.txt", "r") as list:
    for card in list:
        ascii_alphabetized = normalize_alphabetize(card)
        result = is_isogram(ascii_alphabetized)
        if result:
            length = len(ascii_alphabetized)
            if length not in isograms:
                isograms[length] = []
            isograms[length].append(card)
max_len = max(isograms)

print("Top % lengths of isograms:\n", top_x)

i = 0
while i < top_x:
    length = max_len-i
    print("Length %:", length)
    print(isograms[length])
    i = i+1

total = 0
for sublist in isograms.values():
    total += len(sublist)

print("Total isograms in Magic: %", total)
