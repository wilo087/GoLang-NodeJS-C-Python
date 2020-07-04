# Speed test for Go, Node, C and Python
Scripts that compare looping speed between NodeJS, C, Python and Go

## Result on MacBook Pro 
1.  Go looped 1,000,000,000 times in `0.292 seconds` (292 ms) 
2.  NodeJS looped 1,000,000,000 times in `0.587 seconds` (587 ms)
3.  C looped 1,000,000,000 times in `2.394147 seconds`
4.  Python3 looped 1,000,000,000 times in `22.139064056000002 seconds`

## Do you want to test by you self? 

### Language versions
* Golang version go1.13.1 darwin/amd64
* Node v10.16.3
* Clang version 11.0.0 (clang-1100.0.33.8)
* Python 3.7.3

### Running the test
```bash
# Run in Python
python speed-loop.py

# Run in C
gcc speed-loop.c -o speed-loop-c && ./speed-loop-c

# Run in Node
node speed-loop.js

# Run go
go run speed-loop.go
```

### Specs for the results showed above
![Alt text](mac-specifications.png?raw=true "MacBook Pro specs")
