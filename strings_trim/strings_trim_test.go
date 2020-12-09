package strings_trim

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringsTrim(t *testing.T) {
	testInput := "/ce{cecececececece111111}"
	fmt.Println(strings.Trim(testInput, `{/ce`))
}
