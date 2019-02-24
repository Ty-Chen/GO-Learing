package main
import "fmt"

type error interface {
	Error() string
  }

type SomeError struct {
    Reason string
}

func (s SomeError) Error() string {
    return s.Reason
}

func testError() {
    var err error = SomeError{"something happened"}
    fmt.Println(err)
}