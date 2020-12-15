from typing import List
import time
data = [1,0,18,10,19,6]

# data = [3,1,2]

def getSay(data : List[int], target: int):
    start = time.time()
    mem = {}
    for i in range(len(data) - 1):
        mem[data[i]] = i + 1

    last = data[-1]
    for i in range(len(data)+1, target+1):
        current = 0 if last not in mem.keys() else ((i-1) - mem[last])
        mem[last] = i - 1
        last = current
    print(f"finished after {time.time() - start}s")
    return last


print(getSay(data, 30000000))