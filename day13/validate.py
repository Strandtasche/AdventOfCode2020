import numpy as np

def chinese_remainder(n, a):
    sum = 0
    prod = np.prod(n)
    for i in range(len(n)):
        p = prod // n[i]
        sum += a[i] * mul_inv(p, n[i]) * p
    return sum % prod

def mul_inv(a, b):
    b0 = b
    x0, x1 = 0, 1
    if b == 1: return 1
    while a > 1:
        q = a // b
        a, b = b, a%b
        x0, x1 = x1 - q * x0, x0
    if x1 < 0: x1 += b0
    return x1

input = "13,x,x,41,x,x,x,x,x,x,x,x,x,467,x,x,x,x,x,x,x,x,x,x,x,19,x,x,x,x,17,x,x,x,x,x,x,x,x,x,x,x,29,x,353,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,23".split(',')
# input = "1789,37,47,1889".split(',')
inputparsed = [(input[i], i) for i in range(len(input))]
inputparsed2 = list(filter(lambda x: x[0] != 'x', inputparsed))

mods = [int(i[0]) for i in inputparsed2]
offsets = [-1 * i[1] for i in inputparsed2]

print(mods)
print(offsets)


res = chinese_remainder(mods, offsets)
print(res)
# m = np.prod(mods)
# z = [m/x for x in mods]
# y = [gcdExtended(z[i], mods[i])[1]  for i in range(len(mods))]
# w = [(y[i] * z[i]) % m for i in range(len(mods))]
# x = [offsets[i] * w[i] for i in range(len(offsets))]


