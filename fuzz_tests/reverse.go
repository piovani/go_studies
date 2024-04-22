package reverse

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverve(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("a restrig passada nãp é UTF-8 valida")
	}

	b := []byte(s)

	fmt.Printf("bytes %q\n", b)

	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b), nil
}
