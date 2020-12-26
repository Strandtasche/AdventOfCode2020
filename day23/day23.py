from typing import List
import logging

# inpt = "871369452"
inpt = "389125467"
inpt_int = [int(i) for i in inpt] + [i for i in range(10, int(1e6) + 1)]

def iterate(circle: List[int], current: int):
    # print(f"cups: {circle}")
    # print(f"current: {current}")
    current_loc = circle.index(current)
    followers_loc = [(current_loc + i) % len(circle) for i in range(1,4)]
    followers = [circle[i] for i in followers_loc]
    for f in sorted(followers_loc, reverse=True):
        del circle[f]
    target = current - 1
    while target not in circle:
        target = (target - 1) % (len(circle)+4)
    target_loc = circle.index(target)
    circle[target_loc+1:target_loc+1] = followers
    # print(f"followers: {followers}")
    # print(f"target: {target}")
    return circle, circle[(circle.index(current) + 1) % len(circle)]

pivot = inpt_int[0]
for i in range(int(1e7)):
    # print("-------------------")
    # print(f"move {i + 1}")
    inpt_int, pivot = iterate(inpt_int, pivot)
    if i % 1000 == 0:
        print("current state: " + str(i))

index_1 = inpt_int.index(1)
print(inpt_int[index_1 + 1] * inpt_int[index_1 + 2])

