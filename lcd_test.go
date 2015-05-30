package lcd

import (
	"testing"
	"strings"
)

func TestPrintFirstLineOfOneDigitNumber(t *testing.T) {
	lcd := Lcd{}
	expected := " _ "
	result := lcd.print("0")
	
	if !strings.EqualFold(expected, result) {
		t.Errorf("Expected %v but was %v", expected, result)
	}
}

func TestPrintFirstLineOfTwoDigitsNumber(t *testing.T) {
	lcd := Lcd{}
	expected := " _  _ "
	result := lcd.print("02")

	if !strings.EqualFold(expected, result) {
		t.Errorf("Expected %v but was %v", expected, result)
	}
}

func TestPrintFirstLineOfThreeDigitsNumber(t *testing.T) {
	lcd := Lcd{}
	expected := " _  _    "
	result := lcd.print("021")

	if !strings.EqualFold(expected, result) {
		t.Errorf("Expected %v but was %v", expected, result)
	}
}

func TestPrintTwoLineOfThreeDigitsNumber(t *testing.T) {
	lcd := Lcd{}
	expected := " _  _    " +
                    "| |  |  |"
	result := lcd.print("021")

	if !strings.EqualFold(expected, result) {
		t.Errorf("Expected %v but was %v", expected, result)	
	}
}
