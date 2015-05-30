package lcd

import "strconv"
import "fmt"

type Lcd struct {
}

func (l *Lcd) print(number string) string {
	lineTemplate := [][]string{
		[]string{" _ ", "   ", " _ "},
		[]string{"| |", "  |", " _|"},
		[]string{"|_|", "  |", "|_ "},
	}

	line := []string{"", "", ""}
	for _, runeValue := range number {
		digit, _ := strconv.Atoi(string(runeValue))
		line[0] += lineTemplate[0][digit]
		line[1] += lineTemplate[1][digit]
		line[2] += lineTemplate[2][digit]
	}
	return fmt.Sprintf("%s\n%s\n%s\n", line[0], line[1], line[2])
}
