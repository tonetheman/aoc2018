

class CoolRule:
    def __init__(self,id,pre,post):
        self.id = id
        self.pre = pre
        self.post = post
    def __repr__(self):
        return self.pre + "-" + self.post

def loadrules():
    data = open("example-input","r").readlines()
    import re
    P = re.compile("Step (\w+) must be finished before step (\w+) can begin.")
    rules = []
    id = 0
    head = None
    for line in data:
        m = P.search(line)
        pre = m.group(1)
        post = m.group(2)
        # print(pre,post)
        rules.append(CoolRule(id,pre,post))
        id = id + 1
    return rules

class Node:
    def __init__(self,val):
        self.degree = 0
        self.val = val
        self.next = None

def mainline():
    rules = loadrules()
    ia = []

    for rule in rules:
        print(rule)
        mything = None
        for thing in ia:
            if thing.val == rule.pre:
                mything = thing
        print("mything is",mything)
        if mything is None:
            # in this case we are making a new one
            tmp = Node(rule.pre)
            tmp.next = Node(rule.post)
            ia.append(tmp)
        else:
            # we have one already
            tmp = tmp.next
            while tmp.next is not None:
                tmp = tmp.next
            tmp.next = None(rule.post)

if __name__ == "__main__":
    mainline()
