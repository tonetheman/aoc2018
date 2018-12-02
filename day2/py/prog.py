

import string
data = open("day2data.txt","r").readlines()
data = map(string.strip,data)
def chk(data):
    def hasit(s,VV):
        cc = {}
        for c in s:
            if cc.has_key(c):
                cc[c]=cc[c]+1
            else:
                cc[c]=1
        for k in cc.keys():
            if cc[k]==VV:
                return True
        return False
    def has2(s):
        return hasit(s,2)
    def has3(s):
        return hasit(s,3)
    return len(filter(has2,data)) * len(filter(has3,data))
def chk2(data):
    def hasit(s,VV):
        cc = {}
        for c in s:
            if cc.has_key(c):
                cc[c]=cc[c]+1
            else:
                cc[c]=1
        for k in cc.keys():
            if cc[k]==VV:
                return 1
        return 0
    def has2(s):
        return hasit(s,2)
    def has3(s):
        return hasit(s,3)
    def addr(a,b):
        return a+b
    return reduce(addr,map(has3,data)) * reduce(addr,map(has2,data))


print("res", chk2(data))
