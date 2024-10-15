package internal

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"unique"
)

// Before running this program, run ./generate.sh to
// produce the file Le_Comte_de_Monte-Cristo_x100.txt

func Interning() {
	const bookPath = "./Le_Comte_de_Monte-Cristo_x100.txt"
	data, err := os.ReadFile(bookPath)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Read", len(data), "bytes from", bookPath)
	mem()
	book := string(data)
	bWords := findBWords(book)
	mem()
	// Use Bwords
	fmt.Printf("The last B-word is %q\n", bWords[len(bWords)-1].Value())
	mem()
}

const wordChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-àâéèëêëîôû"

var isWordChar = map[rune]bool{}

func init() {
	for _, c := range wordChars {
		isWordChar[c] = true
	}
}

func findBWords(book string) []unique.Handle[string] {
	n := 0
	var bWords []unique.Handle[string]

	a := -1
	for i, c := range book {
		if isWordChar[c] {
			// current char is in a word e.g. 'a', 'à', 'm'
			if a == -1 {
				// start of a word
				a = i
			}
		} else {
			// current char is not in word e.g. ' ', ','
			if a != -1 {
				// just finished a word
				n++
				word := book[a:i]
				if word[0] == 'b' || word[0] == 'B' {
					// In Go 1.23.0 and 1.23.1, unique.Make would retain a reference
					// to the original substring being interned, which is part of a
					// very large string, which would not be garbage collected.
					//
					// This unwanted behavior has been fixed in
					// https://github.com/golang/go/issues/69370
					//
					// Now (after the fix), only small string cloned by unique.Make
					// are retained in the interning pool.
					handle := unique.Make(word)
					bWords = append(bWords, handle)
				}
			}
			a = -1
		}
	}
	fmt.Println("Found", len(bWords), "B-words out of", n, "words")
	return bWords
}

var memStats runtime.MemStats

func mem() {
	runtime.GC()
	runtime.ReadMemStats(&memStats)
	const MiB = 1024 * 1024
	fmt.Println("The program is now using", memStats.Alloc/MiB, "MiB")
}
