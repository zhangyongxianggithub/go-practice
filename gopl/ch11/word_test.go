package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("abba") {
		t.Error(`IsPalindrome("abba") = false`)
	}
	fmt.Println("teststststs")
}

func ExampleIsPalindrome() {
	fmt.Println("example test" + strconv.FormatBool(IsPalindrome("abba")))
	// 输出
}
