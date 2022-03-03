package tddbc

import "fmt"

func Say(greeting string) string {
	return greeting + " TDD BootCamp!!"
}

func Fizzbuzz(v int) string {
	if v%5 == 0 {
		return "buzz"
	}
	if v%3 == 0 {
		return "fizz"
	}
	return fmt.Sprint(v)
}
