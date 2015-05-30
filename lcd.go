package lcd

import "strconv"

type Lcd struct {
}

func (l *Lcd) print(number string) string {
	firstLineTemplate :=  [...]string{" _ ", "   ", " _ "}
	secondLineTemplate := [...]string{"| |", "  |", "  |"}

	firstLineResult := ""
	secondLineResult := ""
	for _, runeValue := range number {
		digit,_ := strconv.Atoi(string(runeValue))
		firstLineResult += firstLineTemplate[digit]
		secondLineResult += secondLineTemplate[digit]
	}

	result := firstLineResult + "\n"
	result += secondLineResult + "\n"
	
        return result
}
