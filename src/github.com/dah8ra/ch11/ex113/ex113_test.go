package ex113

import (
	"math/rand"
	"time"
)

import "testing"

func randomNoPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		/////////////////////
		// Create no palindrome
		/////////////////////
		t := rune(rng.Intn(0x1000))
		for r == t {
			t = rune(rng.Intn(0x1000))
		}
		runes[i] = r
		runes[n-1-i] = t // Input different string for no palindrome
	}
	return string(runes)
}

func TestRandomNoPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 10; i++ {
		p := randomNoPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}
