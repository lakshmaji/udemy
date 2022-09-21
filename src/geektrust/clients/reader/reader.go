package reader

type BaseReader interface {
	FileInput() ([]string, error)
}
