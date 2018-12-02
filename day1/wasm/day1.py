
inf = open("data.txt","r")
data=inf.readlines()
inf.close()

import string
data=map(string.strip,data)

sum = 0
for d in data:
    sum = sum + int(d)
print(sum)
