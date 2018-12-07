

class CoolRule:
    def __init__(self,id,pre,post):
        self.id = id
        self.pre = pre
        self.post = post
    def __repr__(self):
        return self.pre + "-" + self.post

def searcharray(value,a):
    for tmp in a:
        if tmp == value:
            return True
    return False

def mainline():
    data = open("example-input","r").readlines()
    import re
    P = re.compile("Step (\w+) must be finished before step (\w+) can begin.")
    rules = []
    id = 0
    for line in data:
        m = P.search(line)
        pre = m.group(1)
        post = m.group(2)
        print(pre,post)
        rules.append(CoolRule(id,pre,post))
        id = id + 1

    res = []
    choices = ["C"]
    count = 0

    while True:
        print("---------------------------------")
        next = choices[0]
        choices = choices[1:]
        res.append(next)
        remove = []
        for cr in rules:
            if cr.pre == next:
                if not searcharray(cr.post,choices):
                    choices.append(cr.post)
                remove.append(cr.pre)
        for r in remove:
            for cr in rules:
                if r == cr.pre:
                    rules.remove(cr)
        choices.sort()
        
        print("res so far",res)
        print("choices so far",choices)
        print("rules",rules)

        count = count + 1
        if count >= 6:
            print("stopping from count")
            break







if __name__ == "__main__":
    mainline()
