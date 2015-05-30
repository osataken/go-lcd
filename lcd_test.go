package lcd

import (
	"testing"
	"strings"
)

func TestPrintTwoLineOfThreeDigitsWithNewLine(t *testing.T) {
	lcd := Lcd{}
	expected := " _  _    \n" +
	     	    "| |  |  |\n"
	result := lcd.print("021")

	if !strings.EqualFold(expected, result) {
		t.Errorf("Expected %v but was %v", expected, result)
	} 
}
