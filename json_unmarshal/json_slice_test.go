package json_unmarshal

import (
	"encoding/json"
	"fmt"
	"testing"

)

func TestJsonUnmarshalSlice(t *testing.T) {
	test1 := []string {"1","2", "3"}
	test2:= `["1", "2", "3"]`
	test3 := `["1","2","3"]`

	temp, err := json.Marshal(test1)

	fmt.Println(string(temp) == test2)
	fmt.Println(string(temp) == test3)
	fmt.Println(err)
}