
inf = open("data.txt","r")
data=inf.readlines()
inf.close()

import string
data=map(string.strip,data)

sum = 0
res = {}
res[sum] = 1

for d in data:
    sum = sum + int(d)
    if res.has_key(sum):
        print("gOT IT")
        break
print(sum)
