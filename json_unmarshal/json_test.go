package json_unmarshal

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	testSlice := []string{"result", "test", "hello"}
	resultStrSlice, _ := json.Marshal(testSlice)
	fmt.Println("json string slice", string(resultStrSlice)) // ["result","test","hello"]

	resultStringjson, _ :=json.Marshal(string(resultStrSlice))
	fmt.Println("json string:", string(resultStringjson))  // "[\"result\",\"test\",\"hello\"]"
}
