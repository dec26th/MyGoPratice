package json_unmarshal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type B interface {
	Foo()
}

type b struct {
	A int `json:"a"`
}

func (b b) Foo() {

}

func TestJsonUnmarshal(t *testing.T) {
	{
		var o B = b{}
		err := json.Unmarshal([]byte(`{"a": 1}`), &o)
		fmt.Printf("%v %+v\n", err, o)
	}

	{
		var o B = &b{}
		err := json.Unmarshal([]byte(`{"a": 1}`), &o)
		fmt.Printf("%v %+v", err, o)
	}
}

type a struct {
	A int `json:"a"`
}

func TestJsonZeroMethodInterface(t *testing.T) {
	{
		test1 := []interface{}{a{}}[0]
		err := json.Unmarshal([]byte(`{"a": 1}`), &test1) // => map[string][interface]
		fmt.Printf("%v %+v\n", err, test1)
	}
	{
		test1 := []interface{}{&a{}}[0]
		err := json.Unmarshal([]byte(`{"a": 1}`), &test1) // => struct
		fmt.Printf("%v %+v, %v\n", err, test1, reflect.TypeOf(test1))
	}
}
