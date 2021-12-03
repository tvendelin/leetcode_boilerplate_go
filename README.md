# Boilerplate for Leetcode Tasks in Go

Boilerplate for solving [leetcode.com](http://leetcode.com) tasks locally in Go. Steps to take:

- add your test input and expected result into `tcases.txt` separated by space, one test case per line. 
- append the skeleton of the function from your task to `solution.go`
- edit `runIt(input, expected string, t *testing.T)` function in `runit.go`
- run `go test`
- copy-and-paste your solution on [leetcode.com](http://leetcode.com) when done
