package types

type IFileProvider interface {
	TempDir(prefix string) (string, error)
	Create(tmpDir string, pageNumber int) error
	WriteString(s string) error
	Close() error
}
