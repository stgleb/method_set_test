# method_set_test
Testing method set performance in golang

Run tests with

 `go test -bench=.`
 
 Calling method where argument is interface type need
 to check that actual argument method set has all methods
 from interface implemented, so calling function in that
 way can be pretty longer.