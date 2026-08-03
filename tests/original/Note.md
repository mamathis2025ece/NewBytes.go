# Port Tests

Go requires test files to reside in the same package as the source code.

Therefore, the executable test suite is located in:

src/index_test.go

This preserves idiomatic Go project structure while allowing `go test` to discover and execute all tests.
