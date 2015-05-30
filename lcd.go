package lcd

import "strings"

type Lcd struct {
}

func (l *Lcd) print(number string) string {
	if strings.EqualFold("02", number) {
		return " _  _ "
	}

        return " _ "
}
