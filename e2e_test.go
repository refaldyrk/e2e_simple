package e2e_simple

import "testing"

type test struct {
	input  string
	expect string
}

func TestE2E(t *testing.T) {
	testTable := []test{
		{
			input:  "Hello World!",
			expect: "Hello World!",
		},
		{
			input:  "Test Short Text",
			expect: "Test Short Text",
		},
		{
			input:  "Test Long Text, Umm Yes Long Text",
			expect: "Test Long Text, Umm Yes Long Text",
		},
		{
			input:  "Lorem ipsum sit dolor amet, casablanca",
			expect: "Lorem ipsum sit dolor amet, casablanca",
		},
	}

	for _, v := range testTable {
		enc, r, err := E2e_enc(v.input)
		if err != nil {
			t.Error(err)
			t.Fail()
		}

		dec, err := E2e_dec(r, enc)
		if err != nil {
			t.Error(err)
			t.Fail()
		}

		t.Log(string(dec))
	}
}
