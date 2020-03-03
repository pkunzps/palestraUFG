package main

import (
	"localhost/awesomeProject/providers"
	"testing"
)

var (
	svc      *Service
	mockFitz *providers.FitzMock
	mockFile *providers.FileMock
)

func TestService(t *testing.T) {
	setup()

	t.Run("Sucesso", func(t *testing.T) {
		err := svc.withInterfaces()
		if err != nil {
			t.Error()
		}
	})
}

func setup() {
	setupFitz()
	setupFile()
	svc = &Service{fitz: mockFitz, file: mockFile}
}

func setupFitz() {
	mockFitz = providers.NewFitzMock()
}

func setupFile() {
	mockFile = providers.NewFileMock()
}
