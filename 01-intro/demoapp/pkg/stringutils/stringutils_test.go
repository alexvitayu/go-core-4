package stringutils_test

import (
	"go-core-4/01-intro/demoapp/pkg/stringutils"
	"testing"
)

func TestRev(t *testing.T) {
	//Arrange
	inData := "Joint"
	expect := "tnioJ"
	//Act
	result := stringutils.Rev(inData)
	//Assert
	if result != expect {
		t.Errorf("Ожидалось %v, получили %v", expect, result)
	}
}

//Several cases test implementation

var cases = []struct {
	name   string
	inData string
	expect string
}{
	{name: "test1",
		inData: "ABC",
		expect: "CBA",
	},
	{name: "test2",
		inData: "HONEY",
		expect: "YENOH",
	},
	{name: "test3",
		inData: "",
		expect: "",
	},
}

func TestReverse(t *testing.T) {
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := stringutils.Reverse(tc.inData)
			if result != tc.expect {
				t.Errorf("Ожидалось %v, получили %v", tc.expect, result)
			}
		})
	}
}
