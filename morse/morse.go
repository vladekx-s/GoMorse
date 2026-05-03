package morse

import "strings"

var Morse = map[rune]string{
    'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".", 'F': "..-.",
    'G': "--.", 'H': "....", 'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..",
    'M': "--", 'N': "-.", 'O': "---", 'P': ".--.", 'Q': "--.-", 'R': ".-.",
    'S': "...", 'T': "-", 'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-",
    'Y': "-.--", 'Z': "--..",
    ' ': "/", '!': "--..--", '?': "..--..", ',': ".-.-.-",
}

func WordToMorse(phrase string) string {
    phrase = strings.ToUpper(phrase)
    var result []string
    for _, letter := range phrase {
        if code, ok := Morse[letter]; ok {
            result = append(result, code)
        }
    }
    return strings.Join(result, " ")
}