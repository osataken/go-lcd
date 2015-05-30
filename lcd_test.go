package lcd

import "testing"

func printFirstLineOfZeroCharacter(t *testing.T) {
	lcd := Lcd();
	if " _ " != lcd.print("0") {
		t.Errorf("Expected ` _ ` but was %v", lcd.print("0"))
	}
}
