package prose

import (
	"prose"
	"testing"
)

type testData struct {
	sending []string
	wanted  string
}

func TestElements(t *testing.T) {
	tests := []testData{
		{sending: []string{"apple"}, wanted: "apple"},
		{sending: []string{"apple", "pear"}, wanted: "apple and pear"},
		{sending: []string{"apple", "pear", "banana"}, wanted: "apple, pear, and banana"},
	}
	for _, value := range tests {
		got := prose.JoinWithCommas(value.sending)
		if got != value.wanted {
			t.Errorf("%s != %s", value.sending, value.wanted)
		}
	}
}

// it's time to study Go ahead

// func TestOneElement(t *testing.T) {
// 	interest := []string{"apples"}
// 	returned := prose.JoinWithCommas(interest)
// 	wanted := "apples"
// 	if returned != wanted {
// 		t.Errorf("%s != %s", returned, wanted)
// 	}
// }

// func TestTwoElements(t *testing.T) {
// 	interest := []string{"apples", "pears"}
// 	returned := prose.JoinWithCommas(interest)
// 	wanted := "apples and pears"
// 	if returned != wanted {
// 		t.Errorf("%s != %s", returned, wanted)
// 	}
// }

// func TestThreeElements(t *testing.T) {
// 	interest := []string{"apples", "pears", "bananas"}
// 	returned := prose.JoinWithCommas(interest)
// 	wanted := "apples, pears, and bananas"
// 	if returned != wanted {
// 		t.Errorf("%s != %s", returned, wanted)
// 	}
// }
