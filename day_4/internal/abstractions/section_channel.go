package abstractions

type SectionChannel interface {
	Sections() <-chan Section
}
