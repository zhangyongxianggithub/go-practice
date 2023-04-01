package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("abba") {
		t.Error(`IsPalindrome("abba") = false`)
	}
}
