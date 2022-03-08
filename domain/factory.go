package domain

type FizzBuzzFactory struct{}

func NewFizzBuzzFactory() *FizzBuzzFactory {
	return &FizzBuzzFactory{}
}

func (factory FizzBuzzFactory) Create(num uint) (*FizzBuzz, error) {
	return NewFizzBuzz(num), nil
}
