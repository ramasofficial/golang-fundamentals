package main

import (
    "fmt"

    text_formatter "example.com/text-formatter"
)

func main() {
    message := text_formatter.Format("My friend name is %s, he is %d years old.", "Ramas", 26)
    fmt.Println(message)

    message_2 := text_formatter.Format("My friend name is %s, he is %d years old.", "Jonas", 28)
    fmt.Println(message_2)
}
