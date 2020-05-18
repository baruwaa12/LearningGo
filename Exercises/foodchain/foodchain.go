package foodchain

import (
	"fmt"
	"strings"
)

type Phrase struct {
	noun string
	next string
	last string
}

var Phrases = []Phrase{
	{
		noun: "fly",
		next: "I don't know why she swallowed the fly. Perhaps she'll die.",
	},
	{
		noun: "spider",
		next: "It wriggled and jiggled and tickled inside her.",
		last: ".\n",
	},
	{
		noun: "bird",
		next: "How absurd to swallow a bird!",
		last: " that wriggled and jiggled and tickled inside her.\n",
	},
	{
		noun: "cat",
		next: "Imagine that, to swallow a cat!",
		last: ".\n",
	},
	{
		noun: "dog",
		next: "What a hog, to swallow a dog!",
		last: ".\n",
	},
	{
		noun: "goat",
		next: "Just opened her throat and swallowed a goat!",
		last: ".\n",
	},
	{
		noun: "cow",
		next: "I don't know how she swallowed a cow!",
		last: ".\n",
	},
	{
		noun: "horse",
		next: "She's dead, of course!",
	},
}

func opening(phrase Phrase) string {
	return fmt.Sprintf("I know an old lady who swallowed a %s.\n%s", phrase.noun, phrase.next)
}

func next(curr, prev Phrase) string {
	return fmt.Sprintf("\nShe swallowed the %s to catch the %s%s", curr.noun, prev.noun, curr.last)
}

func closing(v int) string {
	previous := strings.Split(Verse(v-1), "\n")
	var ending []string
	if len(previous) < 3 {
		ending = previous[1:]
	} else {
		ending = previous[2:]
	}
	return strings.Join(ending, "\n")
}

func Verse(v int) string {
	current := Phrases[v-1]
	if v == 1 || v == 8 {
		return opening(current)
	}
	previous := Phrases[v-2]
	
	return opening(current) + next(current, previous) + closing(v)
}

func Verses(from, to int) string {
	var vs []string
	for v := from; v <= to; v++ {
		vs = append(vs, Verse(v))
	}
	return strings.Join(vs, "\n")
}

func Song() string {
	return Verses(1, 8)
}