package types

type IFitzProvider interface {
	New(filePath string) error
	GetText() ([]string, error)
}
