
import streams, strutils

var line = ""
var freq = 0
var fs = newFileStream("data_day1.txt", fmRead)
if not isNil(fs):
  while fs.readLine(line):
    freq = freq + strutils.parseInt(line)
  fs.close()
echo "freq ",freq