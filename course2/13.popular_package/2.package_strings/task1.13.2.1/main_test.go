package main

import "testing"

func TestCountWordsInText(t *testing.T) {

	type tests struct {
		test string
		exp  []string
		res  map[string]int
	}

	txttests := []tests{
		{test: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. \n        Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor. \n        Praesent et diam eget libero egestas mattis sit amet vitae augue.", exp: []string{"sit", "amet", "lorem"}, res: map[string]int{"amet": 3, "lorem": 1, "sit": 3}},
	}
	for _, txt1 := range txttests {
		res1 := CountWordsInText(txt1.test, txt1.exp)
		for k, v := range res1 {
			if txt1.res[k] != v {
				t.Errorf("CountWordsInText(txt1.test=text, exp=%v, res=%v)", txt1.res, res1)
			}

		}
	}
}
