package main

import "fmt"

var DigitLetters = map[byte][]byte{
	'0': {'0'},
	'1': {'1'},
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

// O(4^n*n) time | O(4^n*n) space | n: len of number
func PhoneNumberMnemonics(phoneNumber string) []string {
	currentMnemonic := make([]byte, len(phoneNumber))
	for i := range currentMnemonic {
		currentMnemonic[i] = '0'
	}
	mnemonicsFound := make([]string, 0)

	phoneNumberMnemonicsHelper(0, phoneNumber, currentMnemonic, &mnemonicsFound)
	return mnemonicsFound
}

func phoneNumberMnemonicsHelper(idx int, phoneNumber string, currentMnemonic []byte, mnemonicsFound *[]string) {
	if idx == len(phoneNumber) {
		mnemonic := string(currentMnemonic)
		*mnemonicsFound = append(*mnemonicsFound, mnemonic)
	} else {
		digit := phoneNumber[idx]
		letters := DigitLetters[digit]
		for _, letter := range letters {
			currentMnemonic[idx] = letter
			phoneNumberMnemonicsHelper(idx+1, phoneNumber, currentMnemonic, mnemonicsFound)
		}
	}
}

func main() {
	phoneNumber := "1905"
	fmt.Println(PhoneNumberMnemonics(phoneNumber))
}