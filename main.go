package main

import (
	"fmt"
	"time"
	
	//* Подключение файлов из сторонних пакетов
	"morsecode/morse"
	"morsecode/player"
)

func main() {
	for true {
		text := morse.CheckPhrase()

		//* Стоп фраза
		if text == "stop it"{
			fmt.Println("Спасибо за использование!")
			break
		}
		morseCode := morse.WordToMorse(text)

		fmt.Println(text, ":", morseCode)

		for _, i := range morseCode {
			switch i {
				case '.':
					player.PlaySound("sounds/dot.wav")
				case '-':
					player.PlaySound("sounds/dash.wav")
				case '/':
					//* Для наглядности, пауза между словами 0.5 секунд
					time.Sleep(500 * time.Millisecond)
				case ' ':
					//* Пауза между символами
					time.Sleep(100 * time.Millisecond)
			}
		}
	}
}
