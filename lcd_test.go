package lcd

import (
	"testing"
	"strings"
)

func TestPrintFirstLineOfZeroCharacter(t *testing.T) {
	lcd := Lcd{}
	if !strings.EqualFold(" _ ", lcd.print("0")) {
		t.Errorf("Expected ` _ ` but was %v", lcd.print("0"))
	}
}
