# Integration Testing in Go

Problem:  I want to write integration tests that require "components
of integration" to be running.  I want to write these tests in as
similar a manner as my unit tests as possible.  I don't want the
integration tests to run during normal build/test cycles, (i.e. `go
test`), but I would like to use Go's testing framework.

## Running External Components

There are no actual external components to this example.  If there
were, then they would not need to be running during normal (unit)
tests.  And, the external components would need to be started,
*manually* prior to running the integration tests

## Solultion

The first line of each _test.go file contains a build tag, (`// +build
...`).  That tag either sets, or inverts, (`!integration`) the
'integration' tag.  When `go test` is run, any test case ('\_test.go'
file) with the flag set will NOT be run.  Conversely, if `go test` is
run with the `--tags=integration` flag, it will cause 'integration' tagged
test cases to be run.  

Any test cases not setting the `integration` flag will also be run
when `--tags=integration` is added.  In order to keep non-integraton
tests from being run during integration tests, invert the flag as is
done in 'main_test.go'.

## Usage

Unit Tests:
```
go test -v ./...
```

Integration Tests:
```
go test -v --tags=integration ./...
```


## References & Citations

* https://mickey.dev/posts/go-build-tags-testing/
* https://blog.gojekengineering.com/golang-integration-testing-made-easy-a834e754fa4c
* https://medium.com/@victorsteven/understanding-unit-and-integrationtesting-in-golang-ba60becb778d
* https://www.ardanlabs.com/blog/2019/10/integration-testing-in-go-set-up-and-writing-tests.html
* https://pkg.go.dev/testing
