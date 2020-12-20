import numpy as np
from collections import Counter



with open('./data/input20.txt') as f:
    tiles = [l.rstrip('\n') for l in f.read().split('\n\n')]

matched = {}
sides = {}
allsides = []

for tile in tiles:
    tmp = tile.split('\n')
    number = int(tmp[0].split()[1][:-1])
    rep = np.array([list(line) for line in tmp[1:]])
    side_r = ''.join(rep[:, 0])
    side_t = ''.join(rep[0, :])
    side_l = ''.join(rep[:, -1])
    side_b = ''.join(rep[-1, :])
    sides[number] = [side_r, side_t, side_l, side_b]
    allsides.append(side_r)
    allsides.append(side_t)
    allsides.append(side_l)
    allsides.append(side_b)


# print("test")
sides_count = Counter(allsides)
print(max(sides_count.values()))
corners = []

for k, v in sides.items():
    free_sides = []
    for counter, side in enumerate(v):
        occurances = sides_count[side]
        occurances_flipped = sides_count[side[::-1]]
        if occurances + occurances_flipped == 1:
            free_sides.append(side)
    if len(free_sides) >= 2:
        corners.append(k)



print(np.prod(corners))


