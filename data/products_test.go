package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "JJ",
		Price: 1,
		SKU:   "abcd-adfj-adf",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
