package arch

type UseCase[A any, R, any] interface {
	Execute(args A) (*R, error)
}
