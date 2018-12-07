

class CoolRule:
    def __init__(self,pre,post):
        self.pre = pre
        self.post = post
    
def mainline():
    data = open("example-input","r").readlines()
    import re
    P = re.compile("Step (\w+) must be finished before step (\w+) can begin.")
    rules = []
    for line in data:
        m = P.search(line)
        pre = m.group(1)
        post = m.group(2)
        print(pre,post)
        rules.append(CoolRule(pre,post))
    
    res = []
    choices = []

    # first
    res.append("C")

    for cr in rules:
        if cr.pre == "C":
            choices.append(cr.post)
    
    print("res so far",res)
    print("choices so far",choices)

    choices.sort()
    

if __name__ == "__main__":
    mainline()
