package cmd

import (
    "fmt"
    "time"
    "github.com/spf13/cobra"
    "morsecode/morse"
    "morsecode/player"
)

var (
    text string
    play bool
)

var encode = &cobra.Command{
    Use:   "encode",
    Short: "Перевести текст в код Морзе (и проиграть)",
    Run: func(cmd *cobra.Command, args []string) {
        var input string

        if text != "" {
            input = text
        } else if len(args) > 0 {
            input = args[0]
        } else {
            fmt.Println("error: Укажите текст: gomorse encode --text 'TEXT' или gomorse encode 'TEXT'")
            return
        }

        result := morse.WordToMorse(input)
        fmt.Println(result)

        if play {
            for _, symbol := range result {
                switch symbol {
                case '.':
                    player.PlaySound("sounds/dot.wav")
                case '-':
                    player.PlaySound("sounds/dash.wav")
                case '/', ' ':
                    time.Sleep(300 * time.Millisecond)
                }
            }
        }
    },
}

func start() {
    roots.AddCommand(encode)
    encode.Flags().StringVarP(&text, "text", "t", "", "Текст для перевода")
    encode.Flags().BoolVarP(&play, "play", "p", false, "Проиграть результат")
}