package generator

import (
	"crypto/rand"
	"math/big"
)

type Chars struct {
	Digits             string
	LowerCaseAlphabets string
	UpperCaseAlphabets string
	SpecialChars       string
}

func OTP() *Chars {
	return &Chars{
		Digits:             "0123456789",
		LowerCaseAlphabets: "abcdefghijklmnopqrstuvwxyz",
		UpperCaseAlphabets: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		SpecialChars:       "#!&@",
	}
}

func (c *Chars) New(length int, digits bool, lowerCase bool, upperCase bool, specialChars bool) string {
	allowedChars := ""
	password := ""

	if digits {
		allowedChars += c.Digits
	}

	if lowerCase {
		allowedChars += c.LowerCaseAlphabets
	}

	if upperCase {
		allowedChars += c.UpperCaseAlphabets
	}

	if specialChars {
		allowedChars += c.SpecialChars
	}

	usedChars := make(map[rune]bool)

	for len(password) < length {
		charIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(allowedChars))))
		randomChar := rune(allowedChars[charIndex.Int64()])

		if !usedChars[randomChar] {
			password += string(randomChar)
			usedChars[randomChar] = true
		}
	}

	return password
}