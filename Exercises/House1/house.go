package house1

import (
	"strings"
)

func lyricSpitter(noun string, action string, oldVerse []string) []string {
	var verse []string


	return verse
}

func newLine(noun string) string {
	startingLine := "this is the " + noun
	return startingLine
}


func replaceThisIsWithThatAndAction(line string, action string) string{
	thisIsString:= "this is"
	newActionString:= "that " + action

	replacedLine := strings.Replace(line, thisIsString, newActionString, 1)

	return replacedLine
}