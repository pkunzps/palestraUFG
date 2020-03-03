package main

import (
	"fmt"
	"github.com/gen2brain/go-fitz"
	"io/ioutil"
	"localhost/awesomeProject/providers"
	"localhost/awesomeProject/types"
	"os"
	"path/filepath"
)

type Service struct {
	fitz types.IFitzProvider
	file types.IFileProvider
}

func main() {
	for i, arg := range os.Args {
		if i > 0 {
			fmt.Printf("Argumento: %v\n", arg)
		}
	}

	service := Service{}

	//service.withoutInterfaces()

	fitzProvider := providers.NewFitzProvider()
	fileProvider := providers.NewFileProvider()
	service.fitz = fitzProvider
	service.file = fileProvider
	service.withInterfaces()
}

func (s *Service) withoutInterfaces() {
	doc, err := fitz.New("assets/lorem-ipsum.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer doc.Close()

	tmpDir, err := ioutil.TempDir(os.TempDir(), "fitz")
	if err != nil {
		fmt.Println(err)
		return
	}

	for n := 0; n < doc.NumPage(); n++ {
		text, err := doc.Text(n)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(text)

		f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.txt", n)))
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = f.WriteString(text)
		if err != nil {
			fmt.Println(err)
			return
		}

		f.Close()
	}
}

func (s *Service) withInterfaces() error {
	err := s.fitz.New("assets/lorem-ipsum.pdf")
	if err != nil {
		fmt.Println(err)
		return err
	}

	tmpDir, err := s.file.TempDir("fitz")
	if err != nil {
		fmt.Println(err)
		return err
	}

	pages, err := s.fitz.GetText()
	if err != nil {
		fmt.Println(err)
		return err
	}

	for pageNumber, page := range pages {
		fmt.Println(page)

		err := s.file.Create(tmpDir, pageNumber)
		if err != nil {
			fmt.Println(err)
			return err
		}

		err = s.file.WriteString(page)
		if err != nil {
			fmt.Println(err)
			return err
		}

		s.file.Close()
	}

	return nil
}
