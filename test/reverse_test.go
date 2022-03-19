package reverse

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []string{"Tentando Temporin", "", "!1234"}

	for _, c := range testcases {
		rev := Reverse(c)
		rev2 := Reverse(rev)

		if rev2 != c {
			t.Errorf("Esperado: %q, Recebido: %q", c, rev2)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Tentando Temporin", "", "!1234"}

	for _, c := range testcases {
		f.Add(c)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		rev2 := Reverse(rev)

		if orig != rev2 {
			t.Errorf("Esperado: %q, Obtido: %q", orig, rev2)
		}

		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Esperado: %q, Obtido: %q", orig, rev2)
		}
	})
}
