from typing import List, Tuple

def parseMask(inpt: str) -> List[Tuple[int,int]]:
    return [(i, int(x)) for i, x in enumerate(inpt[::-1]) if (x == "1"  or x == "0")]

def applyMask(number: int, mask: List[Tuple[int,int]]) -> int:
    for i, j in mask:
        if j == 1:
            number |= (1<<i)
        else:
            number &=~ (1<<i)
    return number


with open('./data/input14.txt', 'r') as f:
    data = [x.strip() for x in f.readlines()]

mask = []
mem = {}
for i in data:
    split = i.split()
    if split[0] == "mask":
        mask = parseMask(split[2])
    else:
        loc = int(split[0][4:-1])
        mem[loc] = applyMask(int(split[2]), mask)


sumVal = sum(mem.values())
print(sumVal)
