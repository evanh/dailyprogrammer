from itertools import combinations

"""
D O S
R P G
1 2 3
O H F
"""

lines = """DR1O
DP2H
SG3F""".split()

for cards in combinations(lines, 3):
    print cards
    if any(len(set(card[i] for card in cards)) == 2 for i in range(4)):
        continue
    for card in cards:
        print card,