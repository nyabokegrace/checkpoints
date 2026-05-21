package main

import (
    "fmt"
)

func LastWord(s string) string {
    word := ""
    last := ""

    for _, v := range s {
        if v != ' ' {
            word += string(v)
        } else if word != "" {
            last = word
            word = ""
        }
    }

    if word != "" {
        last = word
    }

    return last
}
