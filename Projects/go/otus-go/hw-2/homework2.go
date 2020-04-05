package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	str := `a\4\\5a5`
	symbol := []rune(str)

	//countRune := utf8.RuneCountInString(str)
	if unicode.IsDigit(symbol[0]) || string(symbol[0]) == `\` {
		fmt.Printf("Не корректная строка: %s", str)
	} else {
		if strings.Count(str, `\`) > 1 {
			for i := 0; i < len(str); i++ {
				if !unicode.IsDigit(symbol[i]) {
					if i+2 > len(str) {
						break
					}
					if symbol[i] == 0x5c {
						fmt.Printf("%s", string(symbol[i+1]))
						i++
					} else {
						fmt.Printf("%s", string(symbol[i]))
					}

				} else {
					value, _ := strconv.Atoi(string(symbol[i]))
					y := strings.Repeat(string(symbol[i-1]), value)
					fmt.Print(y)
				}
			}

		}
	}

}
