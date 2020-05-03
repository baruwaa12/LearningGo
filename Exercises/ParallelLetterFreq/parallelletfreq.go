package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency function
func ConcurrentFrequency(inputs []string) FreqMap {
	resChannel := make(chan FreqMap)
	for _, s := range inputs {
		go func(input string) {
			resChannel <- Frequency(input)
		}(s)
	}

	res := make(FreqMap)
	for range inputs {
		for letter, freq := range <-resChannel {
			res[letter] += freq
		}
	}
	return res
}