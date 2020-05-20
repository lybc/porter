package main

import (
    "log"
    "os"
    "porter/commands"
)

func main() {
    err := commands.RootCmd.Run(os.Args)

    if err != nil {
        log.Fatal(err)
    }
}
