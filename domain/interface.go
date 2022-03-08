//go:generate mockgen -source interface.go -destination ./mock/interface_gen.go

package domain

type FizzBuzzRepository interface {
	List() []*FizzBuzz
	Append(f *FizzBuzz)
}
