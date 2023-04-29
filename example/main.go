package main

import (
	"fmt"
	"io"
	"runtime"

	gerrors "errors"

	"github.com/cockroachdb/errors"
)

func main() {
	// foo()
	buzz()
}

func foo() {
	hoge()
	fizz()
}

func hoge() {
	var pc [64]uintptr
	n := runtime.Callers(0, pc[:])
	for _, c := range pc[0 : n-0] {
		fn := runtime.FuncForPC(c)
		if fn == nil {
			fmt.Println("nil fn")
			continue
		}
		file, line := fn.FileLine(c)
		fmt.Println(file, line, fn.Name())
	}
}

var (
	customeErr  = fmt.Errorf("this is sample error")
	customeErr2 = fmt.Errorf("this is sample error2")
)

func fizz() {
	// err := errors.New("hogehoge")
	err := errors.WithHint(customeErr, "this is hint!!")
	err = errors.WithHint(err, "this is hint2!!")
	err = errors.WithStackDepth(err, 1)
	err = errors.CombineErrors(err, customeErr2)
	err = errors.CombineErrors(err, nil)

	fmt.Printf("%+v\n", err)
	fmt.Println(errors.Is(err, customeErr))
	fmt.Println(gerrors.Is(err, customeErr))

	fmt.Println(errors.Is(err, customeErr2))
}

type password struct {
	value []byte
}

func (p password) String() string {
	return "[password]"
}

func (p password) Format(f fmt.State, verb rune) {
	io.WriteString(f, "[password]")
}

func buzz() {
	pass := password{
		value: []byte("hogehogefugafuga"),
	}
	fmt.Printf("%#v", pass)
}
