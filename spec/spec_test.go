package spec

import "testing"

func Test(t *testing.T) {
	if err := Expect(2+2, Eq(4)); err != nil {
		t.Error(err)
	}

	if err := Expect(2+2, Eq(5)); err == nil {
		t.Error(err)
	}
}
