package house1

import (
	"strings"
)

// LyricSpitter takes noun and action and and oldverse and replaces the lyrics
func lyricSpitter(noun string, action string, oldVerse []string) []string {
	oldVerse[0] = replaceThisIsWithThatAndAction(oldVerse[0], action)
	
	newLineSlice := make([]string, 1)
	newLineSlice[0] = newLine(noun)

	result := append(newLineSlice, oldVerse...)
	
	return result
}

// newLine creates a new line with the new noun
func newLine(noun string) string {
	startingLine := "this is the " + noun
	return startingLine
}

// This function replaces the "this is" to "that" and the action.
func replaceThisIsWithThatAndAction(line string, action string) string{
	thisIsString:= "this is"
	newActionString:= "that " + action

	replacedLine := strings.Replace(line, thisIsString, newActionString, 1)

	return replacedLine
}