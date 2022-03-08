package domain

type FizzBuzz struct {
	num uint
}

func NewFizzBuzz(num uint) *FizzBuzz {
	return &FizzBuzz{num: num}
}

func (f FizzBuzz) Fizzbuzz() string {
	return ""
}
