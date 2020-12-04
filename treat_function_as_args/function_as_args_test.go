package treat_function_as_args

import (
	"fmt"
	"testing"
)

type Options func(option *FunctionAsOption)

type FunctionAsOption struct {
	a string
	b string
	c string
}

func (f *FunctionAsOption) String() {
	fmt.Printf("%+v", *f)
}

func getOptions() []Options {
	return []Options{getAOption("print A"), getBOption("print B"), getCOption("print C")}
}

func getAOption(print string) Options {
	return func(option *FunctionAsOption) {
		option.a = print
	}
}

func getBOption(print string) Options {
	return func(option *FunctionAsOption) {
		option.b = print
	}
}

func getCOption(print string) Options {
	return func(option *FunctionAsOption) {
		option.c = print
	}
}

func TestFunctionAsOptions(t *testing.T) {
	opt := new(FunctionAsOption)
	opt.String()

	options := getOptions()

	for i := 0; i < len(options); i++ {
		options[i](opt)
	}
	opt.String()
}
