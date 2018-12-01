
def t(s):
    data = s.split()
    import string
    res=[]
    for d in data:
        res.append(int(d.strip(", \n")))
    print(res)
    freq = 0
    old = {}
    old[freq] = True
    index = 0
    count= 0
    while True:
        i = res[index]
        freq = freq + i
        #print("freq is now",freq)
        if freq in old:
            print("found dup",freq)
            break
        old[freq]=True
        index=index+1
        count=count+1
        if index>=len(res):
            index=0
        if count>6000000:
            break
    print("len of old",len(old))
    return res

# t("+1 -1")
# t("+3, +3, +4, -2, -4")
# t("+1, -2, +3")
# t("-6, +3, +8, +5, -6")
data=open("data_day1.txt","r").readlines()
ts=""
for d in data:
    ts=ts+d
t(ts)