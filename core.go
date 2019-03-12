package main

import (
    "encoding/base64"
    "regexp"
    "strings"
)

func junkEncode(input string) string {
    b64Str := base64.StdEncoding.EncodeToString([]byte(input))

    var junk strings.Builder
    iLast := len(b64Str) - 1
    for i, c := range b64Str {
        isLastOne := i == iLast
        switch i % 3 {
        case 0:
            if val, matched := nouns[c]; matched {
                junk.WriteString(val)
            } else {
                junk.WriteRune('$')
            }
        case 1:
            if val, matched := verbs[c]; matched {
                junk.WriteString(val)
            } else {
                junk.WriteRune('$')
            }
        default:
            if val, matched := nouns[c]; matched {
                junk.WriteString(val)
            } else {
                junk.WriteRune('$')
            }
            if !isLastOne {
                junk.WriteString("，")
            }
        }
        if isLastOne {
            junk.WriteString("。")
        }
    }
    return junk.String()
}

var sepPattern *regexp.Regexp

func init() {
    sepPattern = regexp.MustCompile(`[，。]`)
}

func junkDecode(input string) string {
    filtered := sepPattern.ReplaceAllString(input, "")
    runes := []rune(filtered)
    var head []rune

    var result strings.Builder
    for i := 0; len(runes) >= 2; i++ {
        head, runes = runes[:2], runes[2:]
        dict := revNouns
        if i%3 == 1 {
            dict = revVerbs
        }
        if val, matched := dict[string(head)]; matched {
            result.WriteRune(val)
        }
    }

    resultStr, _ := base64.StdEncoding.DecodeString(result.String())
    return string(resultStr)
}
