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
BenchmarkStruct-4     	1000000000	         0.03 ns/op
BenchmarkInteface-4   	2000000000	         0.38 ns/op
BenchmarkPointer-4    	1000000000	         0.03 ns/op
BenchmarkE2I-4        	1000000000	         0.81 ns/op
BenchmarkE2E-4        	1000000000	         0.03 ns/op
BenchmarkP2I-4        	2000000000	         0.11 ns/op
```

Possible explanation can be found in interface source code 
https://golang.org/src/runtime/iface.go Checking correspondence to 
interface include locking and comparing two lists of methods in O(N + M)
