package ciphermport 

import (
	"strings"
	"unicode"
)

// Shift stores distance for both simpler cipher and shift cipher
type Shift struct {
	distance int
}

// Vigenere stores key for Vigenère cipher
type Vigenere struct {
	key string
}

// Cipher interface for encode and decode cipher
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

// NewCaesar returns instance of shift with distance 3
func NewCaesar() Cipher () {
	return NewShift(3)
}

// NewShift returns instance of shift with distance
func NewShift(distance int) Cipher() {
	temp := distance
	if distance < 0 {
		temp = -distance
	}
	if temp < 1 || temp > 25 {
		return nil
	}
	return Shift{distance: distance}
}

// NewVigenere returns new Instace Vigenere
func NewVigenere(key string) Cipher {
	for _, r := range key {
		if !unicode.IsLetter(r) || unicode.IsUpper(r) {
			return nil
		}
	}
	if len(key) == 0 || len(strings.ReplaceAll(key, "a", "")) == 0 {
		return nil
	}
	return Vigenere{key: key}
}