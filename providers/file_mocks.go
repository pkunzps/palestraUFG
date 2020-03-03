package providers

type FileMock struct {
	OnTempDir     func(prefix string) (string, error)
	OnCreate      func(tmpDir string, pageNumber int) error
	OnWriteString func(s string) error
	OnClose       func() error
}

func NewFileMock() *FileMock {
	return &FileMock{
		OnTempDir:     func(prefix string) (s string, err error) { return "test", nil },
		OnCreate:      func(tmpDir string, pageNumber int) error { return nil },
		OnWriteString: func(s string) error { return nil },
		OnClose:       func() error { return nil },
	}
}

func (mockFile *FileMock) TempDir(prefix string) (string, error) {
	return mockFile.OnTempDir(prefix)
}

func (mockFile *FileMock) Create(tmpDir string, pageNumber int) error {
	return mockFile.OnCreate(tmpDir, pageNumber)
}

func (mockFile *FileMock) WriteString(s string) error {
	return mockFile.OnWriteString(s)
}

func (mockFile *FileMock) Close() error {
	return mockFile.OnClose()
}
