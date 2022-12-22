The `go test` command executes test functions 
## * (whose names begin with Test) in test files 
## * (whose names end with `_test`.go).


You can add the `-v` flag to get verbose output that lists all of the tests and their results.

The tests should pass.
```bash 
$ go test -v
```

```go
package math

// Add is our function that sums two integers
func Add(x, y int) (res int) {
	return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) (res int) {
	return x - y
}
```


```go
package math

import "testing"

func TestAdd(t *testing.T){

    got := Add(4, 6)
    want := 10

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
```


### To summarize, the following are characteristics of a test in Go:

#### * The first and only parameter must be t *testing.T
#### * The testing function begins with the word Test followed by a word or phrase that starts with a capital letter (the convention is to use the name of the method under test, e.g., TestAdd)
#### * The test calls t.Error or t.Fail to indicate a failure (you are calling t.Error because it returns more detail than t.Fail)
#### * You can use t.Log to provide non-failing debug information
#### * Tests are saved in files using this naming convention: foo_test.go, such as math_test.go.





