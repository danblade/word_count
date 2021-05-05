package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type wordCount struct {
	Word  string
	Count int
}

type byWordCount []wordCount

func (w byWordCount) Len() int           { return len(w) }
func (w byWordCount) Less(i, j int) bool { return w[i].Count < w[j].Count }
func (w byWordCount) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }

func countWords(s string) []wordCount {
	input := strings.Fields(s)

	wc := make(map[string]int)
	for _, word := range input {
		wc[word] += 1
	}

	var countedWords []wordCount
	for word, i := range wc {
		countedWords = append(countedWords, wordCount{Word: word, Count: i})
	}

	sort.Sort(sort.Reverse(byWordCount(countedWords)))
	return countedWords
}

func main() { //My WordCount
	var r io.Reader

	if len(os.Args) == 2 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		r = file
	} else if len(os.Args) == 1 {
		r = os.Stdin
	} else {
		fmt.Println("Usage: wordcount  <file to examine>")
		os.Exit(0)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	countedWords := countWords(buf.String())

	for _, line := range countedWords {
		fmt.Printf("%v %v\n", line.Word, line.Count)
	}

}
