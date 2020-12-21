import re
from collections import Counter

with open('/home/tobias/Downloads/input.txt') as f:
    inpt = f.read()
# Vinpt = "mxmxvkd kfcds sqjhc nhms (contains dairy, fish)\ntrh fvjkl sbzzf mxmxvkd (contains dairy)\n" \
#        "sqjhc fvjkl (contains soy)\nsqjhc mxmxvkd sbzzf (contains fish)"

lines = inpt.strip().split('\n')
dictionary = {}
pantry = set([])

for i in lines:
    result = re.search(r'\(contains\s([a-z,\s]+)\)', i)
    ingredients = set(i.split('(')[0].strip().split())
    pantry.update(ingredients)
    # print(result.group(1).split(', '))
    allergens = result.group(1).split(', ')
    for a in allergens:
        if a not in dictionary.keys():
            dictionary[a] = ingredients
        else:
            dictionary[a] = dictionary[a].intersection(ingredients)

changed = True
while changed:
    changed = False
    for k, v in dictionary.items():
        if len(v) == 1:
            pantry = pantry - v
            for k1, v1 in dictionary.items():
                if k1 != k and next(iter(v)) in v1:
                    dictionary[k1] = v1 - v
                    changed = True

print(dictionary)
print(pantry)
countertotal = Counter(inpt.split())

sum_val = 0
for p in pantry:
    sum_val += countertotal[p]

print(f'sum: {sum_val}')

keys = sorted(list(dictionary.keys()))
print(keys)
print(','.join([list(dictionary[k])[0] for k in keys]))






