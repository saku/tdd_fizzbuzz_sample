package controller

type FizzBuzzViewModel struct {
	num  uint
	text string
}

func NewFizzBuzzViewModel(num uint, text string) FizzBuzzViewModel {
	return FizzBuzzViewModel{
		num:  num,
		text: text,
	}
}

type FizzBuzzCountViewModel struct {
	num   uint
	text  string
	count uint
}

func NewFizzBuzzCountViewModel(num uint, text string, count uint) FizzBuzzCountViewModel {
	return FizzBuzzCountViewModel{
		num:   num,
		text:  text,
		count: count,
	}
}
