package lcd

import "strconv"

type Lcd struct {
}

func (l *Lcd) print(number string) string {
	firstLineTemplate :=  [...]string{" _ ", "   ", " _ "}
	secondLineTemplate := [...]string{"| |", "  |", " _|"}
	thirdLineTemplate :=  [...]string{"|_|", "  |", "|_ "}

	firstLineResult := ""
	secondLineResult := ""
	thirdLineResult := ""
	for _, runeValue := range number {
		digit,_ := strconv.Atoi(string(runeValue))
		firstLineResult += firstLineTemplate[digit]
		secondLineResult += secondLineTemplate[digit]
		thirdLineResult += thirdLineTemplate[digit]
	}

	result := firstLineResult + "\n"
	result += secondLineResult + "\n"
	result += thirdLineResult + "\n"
	
        return result
}
