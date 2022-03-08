package domain

import "fmt"

type FizzBuzz struct {
	num uint
}

func NewFizzBuzz(num uint) *FizzBuzz {
	return &FizzBuzz{num: num}
}

func (f *FizzBuzz) Fizzbuzz() string {
	if f.num%15 == 0 {
		return "fizzbuzz"
	}
	if f.num%5 == 0 {
		return "buzz"
	}
	if f.num%3 == 0 {
		return "fizz"
	}
	return fmt.Sprint(f.num)
}
