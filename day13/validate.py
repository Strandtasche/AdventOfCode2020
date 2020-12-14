import numpy as np

def gcdExtended(a, b):
    # Base Case
    if a == 0 :
        return b,0,1

    gcd,x1,y1 = gcdExtended(b%a, a)

    # Update x and y using results of recursive
    # call
    x = y1 - (b//a) * x1
    y = x1

    return gcd,x,y

# input = "13,x,x,41,x,x,x,x,x,x,x,x,x,467,x,x,x,x,x,x,x,x,x,x,x,19,x,x,x,x,17,x,x,x,x,x,x,x,x,x,x,x,29,x,353,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,23".split(',')
input = "1789,37,47,1889".split(',')
inputparsed = [(input[i], i) for i in range(len(input))]
inputparsed2 = list(filter(lambda x: x[0] != 'x', inputparsed))

mods = [int(i[0]) for i in inputparsed2]
offsets = [-1 * i[1] for i in inputparsed2]

print(mods)
print(offsets)

# mods = [int(x) for x in "13 41 467 19 17 29 353 37 23".split()]
# offsets = [-1 * int(x) for x in "0 3 13 25 30 42 44 50 67".split()]

m = np.prod(mods)

z = [m/x for x in mods]
y = [gcdExtended(z[i], mods[i])[1]  for i in range(len(mods))]

w = [y[i] * z[i] % m for i in range(len(mods))]

x = [offsets[i] * w[i] for i in range(len(offsets))]

print(sum(x) % m)

