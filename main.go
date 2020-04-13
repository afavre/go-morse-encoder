package main

func main() {
	var text = "Hello world"
	println("Text to encode: " + text)
	var morse = encode(text)
	println("Morse code: ", morse)
	var textDecoded = decode(morse)
	println("Morse decoded: ", textDecoded)
}
