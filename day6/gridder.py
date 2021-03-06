
class Player:
    def __init__(self,id,row,col):
        self.id = id
        self.row = row
        self.col = col
    def __repr__(self):
        return str(self.id)

class Grid:
    def __init__(self,row,col):
        self.row = row
        self.col = col
        self.grid = []
        for i in range(self.row*self.col):
            self.grid.append(0)
    def get(self,row,col):
        return self.grid[row*self.row+col]
    def set(self,row,col,val):
        self.grid[row*self.row+col]=val
    def __repr__(self):
        ts = ""
        for i in range(self.row):
            for j in range(self.col):
                ts = ts + str(self.get(i,j)) + " "
            ts = ts + "\n"
        return ts
    def read_inputfile(self,filename):
        data = open(filename,"r").readlines()
        import re
        P = re.compile("^(\d+), (\d+)")
        index = 1
        self.players = []
        for line in data:
            m = P.match(line)
            n1 = int(m.group(1)) # x - col
            n2 = int(m.group(2)) # y - row
            self.set(n2,n1,index)
            self.players.append(Player(index,n2,n1))
            index=index+1
    def dist(self,p1row,p1col,p2row,p2col):
        return abs(p1row-p2row)+abs(p1col-p2col)
    def fill_in_one(self,row,col):
        m = {}
        for pl in self.players:
            dd = self.dist(pl.row,pl.col,row,col)
            m[pl.id]=dd
        # print(m)
        smallest = 999999999
        smallest_id = -1
        for k in m.keys():
            v = m[k]
            if v<smallest:
                smallest_id = k
                smallest = v
        count = 0
        for k in m.keys():
            if m[k]==smallest:
                count = count + 1

        # print("smallest value and id",smallest,smallest_id)

        if count == 1:
            return smallest_id
        else:
            return -1
        
    def fill_in_empty(self):
        for i in range(self.row):
            for j in range(self.col):
                p = self.get(i,j)
                if p==0:
                    # empty
                    winner = self.fill_in_one(i,j)
                    if winner == -1:
                        pass
                    else:
                        self.set(i,j,winner)
    def countid(self,id):
        count = 0
        for i in range(self.row):
            for j in range(self.col):
                if self.get(i,j) == id:
                    count = count + 1
        return count               

    def allplayers_toprow(self):
        m = {}
        for i in range(self.col):
            p = self.get(0,i)
            m[p] = True
        return m
    def allplayers_bottomrow(self):
        m = {}
        for i in range(self.col):
            p = self.get(self.row-1,i)
            m[p] = True
        return m
    def allplayers_leftcol(self):
        m = {}
        for i in range(self.row):
            p = self.get(i,0)
            m[p]=True
        return m
    def allplayers_rightcol(self):
        m = {}
        for i in range(self.row):
            p = self.get(i,self.col-1)
            m[p]=True
        return m

def example():
    grid = Grid(10,10)
    grid.read_inputfile("example-input")
    print(grid)
    grid.fill_in_empty()
    print(grid)
    print(grid.countid(5))

def part1():
    print("part1 starting...")
    grid = Grid(500,500)
    print("reading input...")
    grid.read_inputfile("input")
    print("filling out empty spots...")
    grid.fill_in_empty()
    print("all the players",grid.players)

    pp = {}
    for p in grid.players:
        pp[p.id] = True
    print("pp in part1",pp)
    # maybe if they are on the top or bottom
    # they must be infinite
    # since my grid is way larger than the input
    toprow = grid.allplayers_toprow()
    bottomrow = grid.allplayers_bottomrow()
    # remove the top ones
    topkeys = list(toprow)
    for _tmp in topkeys:
        print("toprow player", _tmp, type(_tmp))
        if _tmp in pp:
            print("\t in the top ro!!!")
            del pp[_tmp]
    print("pp",pp)
    bottomkeys = list(bottomrow)
    for _tmp in bottomkeys:
        if _tmp in pp:
            del pp[_tmp]
    print("pp",pp)
    lcol = grid.allplayers_leftcol()
    lkeys = list(lcol)
    for _tmp in lkeys:
        if _tmp in pp:
            del pp[_tmp]
    print("pp",pp)
    rcol = grid.allplayers_rightcol()
    rkeys = list(rcol)
    for _tmp in rkeys:
        if _tmp in pp:
            del pp[_tmp]
    print("pp",pp)
    for _tmp in pp:
        print("pp",grid.countid(_tmp))    

def part2():
    # run once with this
    # grid = Grid(1000,1000)
    # run again and compare
    grid = Grid(1200,1200)
    grid.read_inputfile("input")
    lessthan10k = 0
    for i in range(grid.row):
        for j in range(grid.col):
            row = i
            col = j
            total = 0
            for curr_player in grid.players:
                cdist = grid.dist(row,col,curr_player.row,curr_player.col)
                total = total + cdist
            if total<10000:
                lessthan10k = lessthan10k + 1
            # print("cdist is",total)
    print(lessthan10k)
def mainline():
    part2()

if __name__ == "__main__":
    mainline()