import random

n = 50
pieces = list('OISZLJT')
bag = list(pieces)  # bag starts out with all the pieces
out = ''
len_out = 0 # decided to increment length_out rather than actually calculate it using len(out)

while len_out <= n:
    if len_out % 7 == 0:
        bag = list(pieces)  # refill the bag
    p = random.choice(bag)
    out = out + p
    bag.remove(p)
    len_out += 1

print(out)

