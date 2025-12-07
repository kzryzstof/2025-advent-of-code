package abstractions

type BanksChannel interface {
	Banks() <-chan Bank
}
