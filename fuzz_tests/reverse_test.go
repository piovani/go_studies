package reverse

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Text simples", " ",  "!33241"}

	for _, c := range testcases {
		f.Add(c)
	}

	f.Fuzz(func(t *testing.T, orin string) {
		rev, err1 := Reverve(orin)
		if err1 != nil {
			return 
		}

		rev2, err2 := Reverve(rev)
		if err2 != nil {
			return
		}

		if orin != rev2 {
			t.Errorf("esperado: %q, obtido: %q", orin, rev2)
		}

		if utf8.ValidString(orin) && !utf8.ValidString(rev) {
			t.Errorf("a string reversa nao é utf8")
		}
	})
}
