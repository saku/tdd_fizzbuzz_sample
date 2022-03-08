package domain

import "fmt"

type FizzBuzzFactory struct{}

func NewFizzBuzzFactory() *FizzBuzzFactory {
	return &FizzBuzzFactory{}
}

func (factory FizzBuzzFactory) Create(num uint) (*FizzBuzz, error) {
	if num < 1 || 100 < num {
		return nil, fmt.Errorf("invalid number")
	}
	return NewFizzBuzz(num), nil
}
