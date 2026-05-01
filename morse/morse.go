package morse

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Morse = map[rune]string{
	'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".", 'F': "..-.",
	'G': "--.", 'H': "....", 'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..",
	'M': "--", 'N': "-.", 'O': "---", 'P': ".--.", 'Q': "--.-", 'R': ".-.",
	'S': "...", 'T': "-", 'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-",
	'Y': "-.--", 'Z': "--..",
	' ': "/", '!': "--..-", '?': "..--..", ',': ".-.-.-",
}

var text string

func WordToMorse(phrase string) string {
	//* Перевод в верхний регистр
	phrase = strings.ToUpper(phrase)

	//* Добавление символов в массив
	var result []string

	//* Перевод символов в морзе
	for _, letter := range phrase {
		if i, ok := Morse[letter]; ok {
			result = append(result, i)
		}
	}

	//* Результат - фраза в морзе
	return strings.Join(result, " ")
}

func CheckPhrase() string {
	//* Ввод
	fmt.Print("Введите фразу: ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	//* Возвращается вся фраза
	return scan.Text()
}
