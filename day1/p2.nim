
import streams, strutils,tables

proc readfile(filename : string) : seq[int] =
    var line = ""
    var data : seq[int]
    var fs = newFileStream(filename,fmRead)
    if not isNil(fs):
        while fs.readLine(line):
            data.add(strutils.parseInt(line))
    fs.close()
    return data

var line = ""
var freq = 0
var old = {0:true}.newTable
var data : seq[int] = readfile("data_day1.txt")
var index = 0
while true:
    var v = data[index]
    freq = freq + v
    if old.hasKey(freq):
        echo "got it",freq
        break
    else:
        old[freq] = true
    index = index + 1
    if index >= len(data):
        index = 0

