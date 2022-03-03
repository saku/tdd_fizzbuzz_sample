package tddbc

import "fmt"

func Say(greeting string) string {
	return greeting + " TDD BootCamp!!"
}

func Fizzbuzz(v int) string {
	if isPrime(v) {
		return "prime"
	}
	if v%15 == 0 {
		return "fizzbuzz"
	}
	if v%5 == 0 {
		return "buzz"
	}
	if v%3 == 0 {
		return "fizz"
	}
	return fmt.Sprint(v)
}

func isPrime(v int) bool {
	if v < 2 {
		return false
	}
	for i := 2; i < v; i++ {
		if v%i == 0 {
			return false
		}
	}
	return true
}
