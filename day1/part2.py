
def t(s):
    data = s.split()
    import string
    res=[]
    for d in data:
        res.append(int(d.strip(", \n")))

    mapper = []
    for i in range(160000):
        mapper.append(1)
    
    mapper[0] = 0
    freq = 0
    index=0
    max_freq=-10000000
    while True:
        i = res[index]
        freq = freq + i
        if freq>max_freq:
            max_freq=freq
        if mapper[freq]==0:
            print("GOT IT",freq)
            break
        mapper[freq]=0
        index=index+1
        if index>=len(res):
            index=0
    print("max freq",max_freq)
            
data=open("data_day1.txt","r").readlines()
ts=""
for d in data:
    ts=ts+d
t(ts)
