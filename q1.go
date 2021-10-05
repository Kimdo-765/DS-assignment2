package ds_hw_0

import (
	"fmt"
	"sort"
    "io/ioutil"
    "strings"
    "regexp"
)

// Find the top K most common words in a text document.
// 	path: location of the document
//	numWords: number of words to return (i.e. k)
//	charThreshold: character threshold for whether a token qualifies as a word,
//		e.g. charThreshold = 5 means "apple" is a word but "pear" is not.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the same word.
// A word comprises alphanumeric characters only. All punctuations and other characters
// are removed, e.g. "don't" becomes "dont".
// You should use `checkError` to handle potential errors.
func topWords(path string, numWords int, charThreshold int) []WordCount {
	// TODO: implement me
	// HINT: You may find the `strings.Fields` and `strings.ToLower` functions helpful
	// HINT: To keep only alphanumeric characters, use the regex "[^0-9a-zA-Z]+"
    doc, _ := ioutil.ReadFile(path)
    lowString := strings.ToLower(string(doc))
    r, _ := regexp.Compile("[\r\n]")
    removeReturn := r.ReplaceAllString(lowString, " ")
    r, _ = regexp.Compile("[^0-9a-zA-Z ]+")
    regexDoc := r.ReplaceAllString(removeReturn, "")
    dat := strings.Fields(regexDoc)

    var wordCount []WordCount

    for _, word  := range dat {
        flag := false
        if len(word) < charThreshold {
            continue
        }
        for i, v := range wordCount {
            if v.Word == word {
                wordCount[i].Count++
                flag = true
                break
            }
        }
        if flag == false {
            wordCount = append(wordCount, WordCount{word, 1})
        }
    }

    sortWordCounts(wordCount)

    return wordCount[:numWords]
}

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.
// DO NOT MODIFY THIS FUNCTION!
func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}
