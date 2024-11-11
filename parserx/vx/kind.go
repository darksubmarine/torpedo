package vx

const (
	KindEntity  = "entity"
	KindUseCase = "useCase"
	KindApp     = "app"
)

type K int

const (
	_ K = iota
	KInvalid
	KEntity
	KUseCase
	KApp
)

func Kind(kind string) K {
	if kind == KindEntity {
		return KEntity
	}

	if kind == KindUseCase {
		return KUseCase
	}

	if kind == KindApp {
		return KApp
	}

	return KInvalid
}
