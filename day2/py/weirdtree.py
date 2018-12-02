

"""
import string
data = open("day2data.txt","r").readlines()
fdata= []
for d in data:
   fdata.append(d.strip("\r\n"))
# print(fdata)
"""

fdata = [
"abcde",
"fghij",
"klmno",
"pqrst",
"fguij",
"axcye",
"wvxyz"
]

class Pointer:
    def __init__(self,index,value):
        self.index = index
        self.value = value

first = {}
for word in fdata:
    firstchar = word[0]
    if firstchar in first:
        first[firstchar] = first[firstchar] + 1
    else:
        first[firstchar] = 1

print(first)

for word in fdata:
    secondchar = word[1]
