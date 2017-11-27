# Go Testing 
[src](https://golang.org/pkg/testing/)
Go includes package for testing already as a std library.
This is intended to be used in concet with the `go test` command,
which automates execution of any function of the form:

```go
func TestXxx(*testing.T)
```

Where `Xxx` can be any alphanumeric string (But the first letter [a-z])

To write a new test suite, create a file whose name ends in `_test.go` 
that contains the TestXxx function as described here.  
