package internal

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"unique"
)

// Before running this program, run ./generate.sh to
// produce the file Le_Comte_de_Monte-Cristo_x100.txt

func InterningClone() {
	const bookPath = "./Le_Comte_de_Monte-Cristo_x100.txt"
	data, err := os.ReadFile(bookPath)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Read", len(data), "bytes from", bookPath)
	memClone()
	book := string(data)
	bWords := findBWordsClone(book)
	memClone()
	// Use bWords
	fmt.Printf("The last B-word is %q\n", bWords[len(bWords)-1].Value())
	memClone()
}

func init() {
	for _, c := range wordChars {
		isWordChar[c] = true
	}
}

func findBWordsClone(book string) []unique.Handle[string] {
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
					// to the original string being interned.
					// To prevent this unwanted behavior, we're explicitly cloning the
					// substring, and then intern the clone.
					// This workaround is no longer necessary, since
					// https://github.com/golang/go/issues/69370 was fixed.

					cloned := strings.Clone(word)
					handle := unique.Make(cloned)
					bWords = append(bWords, handle)
				}
			}
			a = -1
		}
	}
	fmt.Println("Found", len(bWords), "B-words out of", n, "words")
	return bWords
}

var memStatClone = runtime.MemStats{}

func memClone() {
	runtime.GC()
	runtime.ReadMemStats(&memStatClone)
	const MiB = 1024 * 1024
	fmt.Println("The program is now using", memStatClone.Alloc/MiB, "MiB")
}
