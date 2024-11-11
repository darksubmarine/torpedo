package engine

type IStackEngine interface {
	Init() error
	Fire() []error
}
