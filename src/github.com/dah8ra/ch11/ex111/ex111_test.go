package ex111

import (
	"bytes"
	"testing"
)

func TestCharcount(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"", 0},
		{"a", 1},
		{"aa", 2},
		{"ab", 2},
		{"kayak", 5},
		{"detartrated", 11},
		{"A man, a plan, a canal: Panama", 30},
		{"Evil I did dwell; lewd did I live.", 34},
		{"Able was I ere I saw Elba", 25},
		{"été", 3},
		{"Et se resservir, ivresse reste.", 31},
		{"palindrome", 10}, // non-palindrome
		{"desserts", 8},    // semi-palindrome
	}
	for _, test := range tests {
		r := bytes.NewReader([]byte(test.input))
		if got := Charcount(r); got != test.want {
			t.Errorf("Charcount(%q) = %v", test.input, got)
		}
	}
}

/*
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}

//!-example

/*
//!+random
import "math/rand"

//!-random
*/

//!+random
// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.

/*
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
*/

//!-random

/*
// Answer for Exercicse 11.1: Modify randomPalindrome to exercise
// IsPalindrome's handling of punctuation and spaces.

// WARNING: the conversion r -> upper -> lower doesn't preserve
// the value of r in some cases, e.g., µ Μ, ſ S, ı I

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomNoisyPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x200)) // random rune up to \u99
		runes[i] = r
		r1 := r
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			r = unicode.ToUpper(r)
			if unicode.ToLower(r) != r1 {
				fmt.Printf("cap? %c %c\n", r1, r)
			}
		}
		runes[n-1-i] = r
	}
	return "?" + string(runes) + "!"
}

func TestRandomNoisyPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	n := 0
	for i := 0; i < 1000; i++ {
		p := randomNoisyPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsNoisyPalindrome(%q) = false", p)
			n++
		}
	}
	fmt.Fprintf(os.Stderr, "fail = %d\n", n)
}
*/
