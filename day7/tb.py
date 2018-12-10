

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


class Tree:
    def __init__(self,val):
        self.val = val
        self.children = []
    def addchild(self,child):
        self.children.append(child)
    def __repr__(self):
        return str(self.val)

def searchTree(val,head):
    if head.val == val:
        return head
    for c in head.children:
        res = searchTree(val,c)
        if res is not None:
            return res
    return None

def pt(head,level):
    print(head,level)
    for c in head.children:
        pt(c,level+1)
    
def mainline():
    rules = loadrules()

    head = None
    for rule in rules:
        if head is None:
            head = Tree(rule.pre)
            head.addchild(Tree(rule.post))
        else:
            res = searchTree(rule.pre,head)
            res2 = searchTree(rule.post,head)
            if res2 is None:
                tmp2 = Tree(rule.post)
            res.children.append(tmp2)

    pt(head,1)

if __name__ == "__main__":
    mainline()
