//go:generate mockgen -source interface.go -destination ./mock/interface_gen.go

package usecase

type Service interface {
	Rand() int
}

type Usecase interface {
	Gacha() int
}
