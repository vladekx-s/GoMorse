package player

import (
	"fmt"
	"os"
	"os/exec"
)

func PlaySound(file string) {
	/*
	  *! -nodisp - БЕЗ графичкеского окна
	  *! -autoexit - автоматический выход после проигрыша звука
	  *! -loglevel quiet - БЕЗ вывод логов
	*/
	cmd := exec.Command("ffplay", "-nodisp", "-autoexit", "-loglevel", "quiet", file)

	cmd.Stdin = os.Stdin

	//* Запуск команды
	err := cmd.Run()

	if err != nil{
		fmt.Println(err)
	}
}