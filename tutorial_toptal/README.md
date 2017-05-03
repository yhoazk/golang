# Tutorial: go programming step-by-step notes

Go is deliberately not an object-oriented language: There are no classes,
objects, or inheritence.

## Go's testing package

**Benchmarks** are like test unit test, but include a loop which runs the same code
many times. This allows the framework to time how long each iteration takes, averaging
out transient differences from disk seeks, cache misses, process scheduling and other 
predictable factors.

### Goroutines 

Goroutines are basically green thread which are threads managed by the VM or a runtime library instead of the 
underlying OS.

Green threads emulate multithread environments without relying on the native OS capabilities, they run on
user space rather than in kernel space.


[https://www.toptal.com/go/go-programming-a-step-by-step-introductory-tutorial](https://www.toptal.com/go/go-programming-a-step-by-step-introductory-tutorial)
