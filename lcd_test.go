package lcd

import (
	"testing"
)

func TestPrintThreeDigitsNumberWithNewLine(t *testing.T) {
	lcd := Lcd{}
	expected := "" +
		" _  _    \n" +
		"| | _|  |\n" +
		"|_||_   |\n"
	result := lcd.print("021")
	if result != expected {
		t.Errorf("Expected %V but was %V", expected, result)
	}
}
