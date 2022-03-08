package domain

import "fmt"

type FizzBuzz struct {
	num uint
}

func NewFizzBuzz(num uint) *FizzBuzz {
	return &FizzBuzz{num: num}
}

func (f *FizzBuzz) Fizzbuzz() string {
	if f.isFizzBuzz() {
		return "fizzbuzz"
	}
	if f.isBuzz() {
		return "buzz"
	}
	if f.isFizz() {
		return "fizz"
	}
	return fmt.Sprint(f.num)
}

func (f *FizzBuzz) isFizzBuzz() bool {
	return f.num%15 == 0
}

func (f *FizzBuzz) isFizz() bool {
	return f.num%3 == 0
}

func (f *FizzBuzz) isBuzz() bool {
	return f.num%5 == 0
}
