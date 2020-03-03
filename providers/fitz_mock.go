package providers

type FitzMock struct {
	OnNew     func(filePath string) error
	OnGetText func() ([]string, error)
}

func NewFitzMock() *FitzMock {
	return &FitzMock{
		OnNew:     func(filePath string) error { return nil },
		OnGetText: func() (strings []string, err error) { return []string{"test 1", "test 2"}, nil },
	}
}

func (mockFitz *FitzMock) New(filePath string) error {
	return mockFitz.OnNew(filePath)
}

func (mockFitz *FitzMock) GetText() ([]string, error) {
	return mockFitz.OnGetText()
}
