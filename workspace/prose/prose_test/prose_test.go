package prose

import (
	"prose"
	"testing"
)

func TestOneElement(t *testing.T) {
	interest := []string{"apples"}
	returned := prose.JoinWithCommas(interest)
	wanted := "apples"
	if returned != wanted {
		t.Errorf("%s != %s", returned, wanted)
	}
}

func TestTwoElements(t *testing.T) {
	interest := []string{"apples", "pears"}
	returned := prose.JoinWithCommas(interest)
	wanted := "apples and pears"
	if returned != wanted {
		t.Errorf("%s != %s", returned, wanted)
	}
}

func TestThreeElements(t *testing.T) {
	interest := []string{"apples", "pears", "bananas"}
	returned := prose.JoinWithCommas(interest)
	wanted := "apples, pears, and bananas"
	if returned != wanted {
		t.Errorf("%s != %s", returned, wanted)
	}
}
