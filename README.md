# method_set_test
Testing method set performance in golang

Run tests with

 `go test -bench=.`
 
 Calling method where argument is interface type need
 to check that actual argument method set has all methods
 from interface implemented, so calling function in that
 way can be pretty longer.
 
 Results shows substantial difference when using interface
 as an argument.
 
```
BenchmarkStruct-4     	2000000000	         0.02 ns/op
BenchmarkInteface-4   	2000000000	         0.44 ns/op
PASS
ok  	_/home/stgleb/workspace/benchmethodset	25.838s
```

Possible explanation can be found in interface source code 
https://golang.org/src/runtime/iface.go Checking correspondence to 
interface include locking and comparing two lists of methods in O(N + M)