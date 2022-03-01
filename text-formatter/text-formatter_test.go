package text_formatter

import (
    "testing"
)

func TestTextFormatterCanReturnCorrectString(t *testing.T) {
    actual := Format("My friend name is %s, he is %d years old.", "Jonas", 28)
    expected := "My friend name is Jonas, he is 28 years old."

    if(actual != expected) {
        t.Fatalf("Strings: %s and %s are not equals!", actual, expected)
    }
}
