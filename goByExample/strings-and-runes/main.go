package main

/*
A Go string is a read-only slice of bytes. The language and the standard library treat strings specially - as containers of text encoded in UTF-8. In other languages, strings are made of “characters”. In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point. This Go blog post is a good introduction to the topic.
*/

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี" // s is a string assigned a literal value representing the word “hello” in the Thai language. Go string literals are UTF-8 encoded text.

	fmt.Println("Len:", len(s)) // Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x \n", s[i]) // Indexing into a string produces the raw byte values at each index. This loop generates the hex values of all the bytes that constitute the code points in s.
	}

	fmt.Println("Rune Count:", utf8.RuneCountInString(s)) // To count how many runes are in a string, we can use the utf8 package. Note that the run-time of RuneCountInString depends on the size of the string, because it has to decode each UTF-8 rune sequentially. Some Thai characters are represented by multiple UTF-8 code points, so the result of this count may be surprising.

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx) // A range loop handles strings specially and decodes each rune along with its offset in the string.
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:]) // We can achieve the same iteration by using the utf8.DecodeRuneInString function explicitly.
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' { // Values enclosed in single quotes are rune literals. We can compare a rune value to a rune literal directly.
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
