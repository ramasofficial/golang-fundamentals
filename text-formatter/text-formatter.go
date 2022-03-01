package text_formatter

import "fmt"

func Format(text string, arguments ...interface{}) string {
    message := fmt.Sprintf(text, arguments...)

    return message
}
