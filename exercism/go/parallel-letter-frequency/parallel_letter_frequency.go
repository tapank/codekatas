package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	var wg sync.WaitGroup
	fmchan := make(chan FreqMap, 10)
	anschan := make(chan FreqMap, 1)

	// create consumer
	go func(fmchan chan FreqMap, anschan chan FreqMap) {
		frequencies := FreqMap{}
		for fm := range fmchan {
			for k, v := range fm {
				frequencies[k] += v
			}
		}
		anschan <- frequencies
	}(fmchan, anschan)

	// create producers
	for _, s := range texts {
		wg.Add(1)
		go func(text string, fmchan chan FreqMap) {
			fmchan <- Frequency(text)
			wg.Done()
		}(s, fmchan)
	}
	wg.Wait()
	close(fmchan)

	return <-anschan
}
