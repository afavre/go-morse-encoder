package main

import (
	"strings"
)

type MorseChar struct {
	letter    string
	morseCode string
}

var delimiterWordSlash = " / "
var delimiterWord3Spaces = "   "
var delimiterWord7Spaces = "       "
var delimiterLetterSpace = " "

var alphabet = []MorseChar{
	MorseChar{"A", ".-"},
	MorseChar{"B", "-..."},
	MorseChar{"C", "-.-."},
	MorseChar{"D", "-.."},
	MorseChar{"E", "."},
	MorseChar{"F", "..-."},
	MorseChar{"G", "--."},
	MorseChar{"H", "...."},
	MorseChar{"I", ".."},
	MorseChar{"J", ".---"},
	MorseChar{"K", "-.-"},
	MorseChar{"L", ".-.."},
	MorseChar{"M", "--"},
	MorseChar{"N", "-."},
	MorseChar{"O", "---"},
	MorseChar{"P", ".--."},
	MorseChar{"Q", "--.-"},
	MorseChar{"R", ".-."},
	MorseChar{"S", "..."},
	MorseChar{"T", "-"},
	MorseChar{"U", "..-"},
	MorseChar{"V", "...-"},
	MorseChar{"W", ".--"},
	MorseChar{"X", "-..-"},
	MorseChar{"Y", "-.--"},
	MorseChar{"Z", "--.."},
	MorseChar{"1", ".----"},
	MorseChar{"2", "..---"},
	MorseChar{"3", "...--"},
	MorseChar{"4", "....-"},
	MorseChar{"5", "....."},
	MorseChar{"6", "-...."},
	MorseChar{"7", "--..."},
	MorseChar{"8", "---.."},
	MorseChar{"9", "----."},
	MorseChar{"0", "-----"}}

func dictionaryMorseToEnglish() map[string]string {
	var morseToEnglishMap = make(map[string]string)
	for index := 0; index < len(alphabet); index++ {
		morseToEnglishMap[alphabet[index].morseCode] = alphabet[index].letter
	}
	return morseToEnglishMap
}

func splitMorseWords(text string) []string {
	var words = strings.Split(text, delimiterWordSlash)
	if len(words) == 1 {
		return strings.Split(text, delimiterWord3Spaces)
	}
	if len(words) == 1 {
		return strings.Split(text, delimiterWord7Spaces)
	}
	return words
}

func splitMorseLetters(text string) []string {
	return strings.Split(text, delimiterLetterSpace)
}

func decode(text string) string {
	var dictionaryMorseToEnglish = dictionaryMorseToEnglish()
	var decodedText string
	var words = splitMorseWords(text)
	for i := 0; i < len(words); i++ {
		var sanitisedWord = strings.TrimSpace(words[i])
		var letters = splitMorseLetters(sanitisedWord)
		for letterIndex := 0; letterIndex < len(letters); letterIndex++ {
			var currentLetter = letters[letterIndex]
			var decodedLetter = dictionaryMorseToEnglish[currentLetter]
			decodedText += decodedLetter
		}
		if i < len(words)-1 {
			decodedText += delimiterLetterSpace
		}
	}
	return decodedText
}

func splitEnglishWords(text string) []string {
	return strings.FieldsFunc(text, Split)
}

func Split(r rune) bool {
	return r == '.' || r == ' ' || r == '?' || r == '!' || r == ',' || r == ';' || r == ':'
}

func dictionaryEnglishToMorse() map[string]string {
	var dict = make(map[string]string)
	for index := 0; index < len(alphabet); index++ {
		dict[alphabet[index].letter] = alphabet[index].morseCode
	}
	return dict
}

func encode(text string) string {
	var dictionary = dictionaryEnglishToMorse()
	var encodedText string
	var words = splitEnglishWords(text)
	for i := 0; i < len(words); i++ {
		var sanitisedWord = strings.ToUpper(strings.TrimSpace(words[i]))
		for letterIndex := 0; letterIndex < len(sanitisedWord); letterIndex++ {
			var currentLetter = sanitisedWord[letterIndex : letterIndex+1]
			var encodedLetter = dictionary[currentLetter]
			encodedText += encodedLetter
			if letterIndex < len(sanitisedWord)-1 {
				encodedText += delimiterLetterSpace
			}
		}
		if i < len(words)-1 {
			encodedText += delimiterWordSlash
		}
	}
	return encodedText
}
