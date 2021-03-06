package tddbc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSay(t *testing.T) {
	actual := Say("Wow!")
	expected := "Wow! TDD BootCamp!!"

	if actual != expected {
		t.Errorf("actual=%s, expect=%s", actual, expected)
	}
}

func TestSay_testify(t *testing.T) {
	actual := Say("Hello!")
	assert.Equal(t, "Hello! TDD BootCamp!!", actual, "they should be equal")
}

func TestFizzbuzz(t *testing.T) {
	assert.Equal(t, "1", Fizzbuzz(1), "1")
	assert.Equal(t, "prime", Fizzbuzz(2), "2")
	assert.Equal(t, "prime", Fizzbuzz(3), "3")
	assert.Equal(t, "prime", Fizzbuzz(5), "5")
	assert.Equal(t, "fizz", Fizzbuzz(6), "6")
	assert.Equal(t, "buzz", Fizzbuzz(10), "10")
	assert.Equal(t, "fizzbuzz", Fizzbuzz(15), "15")
}
