package player

import (
    "fmt"
    "os"
    "os/exec"
)

func PlaySound(file string) {
    cmd := exec.Command("ffplay", "-nodisp", "-autoexit", "-loglevel", "quiet", file)
    cmd.Stdin = os.Stdin
    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
    }
}