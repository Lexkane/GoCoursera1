package main

import (
	"regexp"
	"strings"
	"testing"
)

var (
	browser = `Mozilla/ 5.0 (Windows NT 6.3;WOW64)AppleWebKit/537.36 (KHTML,like Gecko) Chrome/58.0.3029.110
	Safari/537.36`
	re = regexp.MustCompile("Android")
)

//Regexp Match compiles everytime
func BenchmarkRegExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = regextp.MatchString("Android", browser)
	}
}

//Precompiled Regulary Expression
func BenchmarkRegCompiled(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = re.MatchString(browser)
	}
}
func BenchmarkStrContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = strings.Contains(browser, "Android")
	}
}
func main() {

}

//go test -bench . -benchmem string_test.go
