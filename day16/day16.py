
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
# height = len(res[2]) - 1
for elements in res[2][1:]:
    splitted = elements.split(',')
    for i in splitted:
        if not valid(int(i), rangetuples):
            invalids += int(i)


print(invalids)