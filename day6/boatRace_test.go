package day6

import (
	"testing"
)

func Test_CalcWaysToWin_Returns_Expected_Value_For_Example_Data(t *testing.T) {
	test1 := exampleData[0]
	actual := CalcWaysToWin(test1)
	expected := 4
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	test2 := exampleData[1]
	actual = CalcWaysToWin(test2)
	expected = 8
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	test3 := exampleData[2]
	actual = CalcWaysToWin(test3)
	expected = 9
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func Test_CalcErrorMargin_Returns_Expected_Result_For_Example(t *testing.T) {
	expected := 288
	actual := CalcErrorMargin(exampleData)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func Test_CalcErrorMargin_Returns_Expected_Result_For_Real_Data(t *testing.T) {
	expected := 138915
	actual := CalcErrorMargin(realData)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
