import collections
import math
import re
import sys


with open('./data/input19.txt') as f:
    rules, strings = [l.rstrip('\n') for l in f.read().split('\n\n')]
rules = dict([rule.split(': ', 1) for rule in rules.split('\n')])

def genRegex(number: str) -> str:
    if number == '8':
        return genRegex('42') + '+'
    elif number == '11':
        firstpart = genRegex('42')
        trailingpart = genRegex('31')
        return '(' + '|'.join(f'{firstpart}{{{n}}}{trailingpart}{{{n}}}' for n in range(1, 50)) + ')'

    rule = rules[number]
    if re.fullmatch(r'"."', rule):
        return rule[1]
    else:
        split0 = [r.strip() for r in rule.split('|')]
        regex = []
        for elem in split0:
            subregex = []
            split1 = elem.split()
            for spli in split1:
                subregex.append(genRegex(spli))
            regex.append(''.join(subregex))
        return '(' + '|'.join(regex) + ')'


# print(genRegex('0'))

count = 0
for s in strings.split('\n'):
    regex = genRegex('0')
    count += 1 if re.fullmatch(regex, s) else 0

    # print(f"{s} match: {value}")
print(count)
