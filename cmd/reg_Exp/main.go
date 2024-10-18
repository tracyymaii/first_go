package main

import (
	"fmt"
	"regexp" // package details: https://pkg.go.dev/regexp#pkg-index
)

// regexp cheat sheet for go: https://gist.github.com/harrietty/d737a350827e100712c5b62168358c88

func main() {
	//func MatchString(pattern string, s string) (matched bool) // look for pattern in the string, searching for string
	matched, err := regexp.MatchString(`foo.*`, "seafood")
	fmt.Println("matched:", matched, " error:", err)

	// func Match(pattern string, b []byte) (matched bool, err) // searchng for slice of bytes
	matched, err = regexp.Match(`\w+fo\w`, []byte(`seafood`))
	fmt.Println("matched:", matched, " errir:", err)

	// Compile the expressions once, usually at init time.
	// Use raw string to avoid having to quote the backslashes
	// func Compile(expr string) (*Regexp, error)

	var re *regexp.Regexp
	re, err = regexp.Compile(`\wat`)
	fmt.Println("re:", re, " error:", err)
	str1 := `bat mat pot sat cat rat pat vat hat`

	//func (re *Regexp) MAtch,String(s string) bool
	fmt.Println("MatchString:", re.MatchString(str1))

	//func (re *Regexp) Find (b []byte) [] byte
	fmt.Println("Find:", string(re.Find([]byte(str1))))

}
