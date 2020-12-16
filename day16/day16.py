import numpy as np
from collections import defaultdict

def valid(number: int, ranges) -> bool:
    for lower, upper in ranges:
        if number >= lower and number <= upper:
            return True
    return False


with open('./data/input16.txt', 'r') as f:
    data = [x.strip() for x in f.readlines()]

rangetuples = []

indices = [i+1 for i, x in enumerate(data) if x == ""]
res = [data[i: j] for i, j in
        zip([0] + indices, indices +
        ([len(data)] if indices[-1] != len(data) else []))]

for validRanges in res[0][:-1]:
    splitted = validRanges.split()
    range1 = splitted[-1].split('-')
    rangetuples.append((int(range1[0]), int(range1[1])))
    range2 = splitted[-3].split('-')
    rangetuples.append((int(range2[0]), int(range2[1])))

print(rangetuples)

invalids = 0
validTickets = []
for element in res[2][1:]:
    splitted = element.split(',')
    validTicket = True
    for i in splitted:
        if not valid(int(i), rangetuples):
            validTicket = False
            break
    if validTicket:
        validTickets.append(element)

elements = np.fromstring(",".join(validTickets), sep=',', dtype=int).reshape(len(validTickets), -1)

row_to_col = defaultdict(list)

for col in range(elements.shape[1]):
    added = False
    for pos in range(0, len(rangetuples), 2):
        tmp = [valid(i, rangetuples[pos:pos+2]) for i in elements[:, col]]
        if all(tmp):
            row_to_col[pos/2].append(col)
            added = True
    if not added:
        raise ValueError("wrong assumption")

print(row_to_col)
final_match = {}
while len(final_match) != len(res[0]) - 1:
    for k in row_to_col.keys():
        if len(row_to_col[k]) == 1:
            final_match[k] = row_to_col[k][0]
            for k2 in row_to_col.keys():
                if final_match[k] in row_to_col[k2]:
                    row_to_col[k2].remove(final_match[k])

print(final_match)

splitted = res[1][1].split(',')
prod = 1
for i in range(6):
    prod *= int(splitted[final_match[i]])

print(prod)