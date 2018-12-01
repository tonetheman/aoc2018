

from bitarray import bitarray
import hashlib

def t(s):
    data = s.split()
    import string
    res=[]
    for d in data:
        res.append(int(d.strip(", \n")))
    
    ba = bitarray(1000000)
    ba.setall(0)

    def add(num):
        COUNT=3
        for i in range(COUNT):
            m = hashlib.sha256()
            m.update(i,num)
            
    add(0)

    freq = 0
    index=0
    while True:
        i = res[index]
        freq = freq + i
            
data=open("data_day1.txt","r").readlines()
ts=""
for d in data:
    ts=ts+d
t(ts)
