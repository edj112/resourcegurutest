package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	text, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	morse := getMorse(string(text))
	obfuscation := getObfuscation(morse)
	fmt.Println(obfuscation)
}

func getMorse(text string) string {
	var morse string
	for i, t := range text {
		if code, found := morseCode[t]; found {
			morse = morse + code
		}else if t == ' ' {
			morse = morse + "/"
			continue
		} else {
			morse = morse + string(t)
			continue
		}

		if i != len(text)-1 && text[i+1] != ' ' && text[i+1] != '\n' {
			morse = morse + "|"
		}
	}
	return morse
}



func getObfuscation(morse string) string {
	morseRune := []rune(morse)
	var obfuscation []rune

	for i := 0; i < len(morseRune); i++ {
		if morseRune[i] == '.' || morseRune[i] == '-' {
			count := 0
			m2 := morseRune[i]
			for j := i;  j < len(morseRune) && morseRune[j] == m2; j++ {
				count++
			}
			i = count + i - 1
			if morseRune[i] == '.' {
				obfuscation = append(obfuscation, obfuscationInts[count])
			}
			if morseRune[i] == '-' {
				obfuscation = append(obfuscation, obfuscationChars[count])
			}

		} else {

			obfuscation = append(obfuscation, morseRune[i])
		}

	}
	return string(obfuscation)
}
