package lcd

import (
	"testing"
	"strings"
)

type Lcd struct {
}

func (l *Lcd) print(number string) string {
	return ""
}

func TestPrintFirstLineOfZeroCharacter(t *testing.T) {
	lcd := Lcd{}
	if !strings.EqualFold(" _ ", lcd.print("0")) {
		t.Errorf("Expected ` _ ` but was %v", lcd.print("0"))
	}
}
