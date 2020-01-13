package main

import "testing"

func TestGetMorse(t *testing.T) {
	text := "I AM IN TROUBLE"
	morse := getMorse(text)
	if morse != "../.-|--/..|-./-|.-.|---|..-|-...|.-..|." {
		t.Error(morse)
	}

	text = "HELLO\nI AM IN TROUBLE"
	morse = getMorse(text)
	if morse != "....|.|.-..|.-..|---\n../.-|--/..|-./-|.-.|---|..-|-...|.-..|." {
		t.Error(morse)
	}
}

func TestGetObfuscation(t *testing.T) {
	text := "I AM IN TROUBLE"
	morse := getMorse(text)
	obfuscation := getObfuscation(morse)
	if obfuscation != "2/1A|B/2|A1/A|1A1|C|2A|A3|1A2|1" {
		t.Error(obfuscation)
	}

	text = "HELLO\nI AM IN TROUBLE"
	morse = getMorse(text)
	obfuscation = getObfuscation(morse)
	if obfuscation != "4|1|1A2|1A2|C\n2/1A|B/2|A1/A|1A1|C|2A|A3|1A2|1" {
		t.Error(obfuscation)
	}
}