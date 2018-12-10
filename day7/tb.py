

class CoolRule:
    def __init__(self,id,pre,post):
        self.id = id
        self.pre = pre
        self.post = post
    def __repr__(self):
        return self.pre + "-" + self.post

class Treenode:
    def __init__(self,val):
        self.val = val
        self.pointers = []
    def add(self,child_node):
        self.pointers.append(child_node)
    def __repr__(self):
        ts = str(self.val)
        for c in self.pointers:
            ts = ts + "-" + str(c)
        return ts
def findNode(val,head):
    if head.val == val:
        return val
    for c in head.pointers:
        res = findNode(val,c)
        if res:
            return c

def mainline():
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
        print(pre,post)
        rules.append(CoolRule(id,pre,post))
        id = id + 1


    for r in rules:
        if head is None:
            head = Treenode(r.pre)
            head.add(Treenode(r.post))
            print(head)
        else:
            print("need to fill out tree")
            ip = findNode(r.pre,head)
            print("found this node",ip)
            ip.pointers.append(Treenode(r.pre))

if __name__ == "__main__":
    mainline()
