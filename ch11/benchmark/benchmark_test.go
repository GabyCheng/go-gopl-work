package benchmark

import (
	word "go-study/ch11/word1"
	"testing"
)

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		word.IsPalindrome("A man, a plan, a canal: Panama")
	}
}
