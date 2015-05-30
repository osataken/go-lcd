package lcd

import (
	"testing"
	"strings"
)

func TestPrintThreeDigitsNumberWithNewLine(t * testing.T) {
	lcd := Lcd{}
	expected := " _  _    \n" +
	     	    "| | _|  |\n" +
		    "|_||_   |\n"
	result := lcd.print("021")
	if !strings.EqualFold(expected, result) {
		t.Errorf("Expected %v but was %v", expected, result) 
	}
}
