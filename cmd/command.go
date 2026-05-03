package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var roots = &cobra.Command{
    Use:   "gomorse",
    Short: "Переводчик текста в азбуку Морзе",
}

func CLI() {
    start()
    if err := roots.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}