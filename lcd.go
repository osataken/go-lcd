package lcd

import "strconv"

type Lcd struct {
}

func (l *Lcd) print(number string) string {
	firstLineTemplate := [...]string{" _ ", "   ", " _ "}

	result := ""
	for _, runeValue := range number {
		digit,_ := strconv.Atoi(string(runeValue))
		result += firstLineTemplate[digit]
	}

        return result
}
