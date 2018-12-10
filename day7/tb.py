

class CoolRule:
    def __init__(self,id,pre,post):
        self.id = id
        self.pre = pre
        self.post = post
    def __repr__(self):
        return self.pre + "-" + self.post

class Tree:
    def __init__(self,val):
        self.val = val
        self.points_to = None
    def __repr__(self):
        ts = "val:" + str(self.val) + " -> " + str(self.points_to)
        return ts

def findInTree(head,val):
    if head.val == val:
        return head
    return findInTree(head.points_to)

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
        # print(pre,post)
        rules.append(CoolRule(id,pre,post))
        id = id + 1

    head = None
    count = 0
    for r in rules:
        pre = r.pre
        post = r.post

        if head is None:
            head = Tree(pre)
            child = Tree(post)
            head.points_to = child
        else:
            insert_point = findInTree(head,pre)
            print("found insert point")
            
        print("head is",head)
        count = count + 1
        if count==2:
            break

if __name__ == "__main__":
    mainline()
