package lcd

import (
	"testing"
	"strings"
)

func TestPrintFirstLineOfOneDigitNumber(t *testing.T) {
	lcd := Lcd{}
	if !strings.EqualFold(" _ ", lcd.print("0")) {
		t.Errorf("Expected ` _ ` but was %v", lcd.print("0"))
	}
}

func TestPrintFirstLineOfTwoDigitsNumber(t *testing.T) {
	lcd := Lcd{}
	if !strings.EqualFold(" _  _ ", lcd.print("02")) {
		t.Errorf("Expected ` _  _ ` but was %v", lcd.print("02"))
	}
}
