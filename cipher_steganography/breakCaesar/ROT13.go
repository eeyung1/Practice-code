package main

func ROT13(text string) string {
	return CaesarCipher(text, 13)
}