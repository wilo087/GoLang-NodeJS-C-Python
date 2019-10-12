# Node vs C vs Python
Scripts that show which is faster between NodeJS, C, Python and Go

## Result on MacBook Pro 
1.  Go looped 1,000,000,000 time in `292.138231 ms`
2.  NodeJS looped 1,000,000,000 times in `0.587 seconds` (587 ms)
3.  C looped 1,000,000,000 times in `2.394147 seconds`
4.  Python3 looped 1,000,000,000 times in `22.139064056000002 seconds`

## Do you want to test by you? 

### Languages versions
* Golang version go1.13.1 darwin/amd64
* Node v10.16.3
* Clang version 11.0.0 (clang-1100.0.33.8)
* Python 3.7.3

### Setup
```sh
# Clone or download the repo
git clone git@github.com:wilo087/Node-vs-C-Python.git && cd Node-vs-C-Python

# Compile in Go
go build speed-loop.go

# Compile in C
gcc speed-loop.c -o speed-loop-c
```
### Running the test
```bash
# Run in Python
python speed-loop.py

# Run in C
./speed-loop-c

# Run in Node
node speed-loop.js

# Run in Go
./speed-loop
```

### This test was made with
![MacBook Pro]('https://github.com/wilo087/Node-vs-C-Python/blob/master/mac-specifications.png')
