package providers

import (
	"github.com/gen2brain/go-fitz"
)

type FitzProvider struct {
	doc *fitz.Document
}

func NewFitzProvider() *FitzProvider {
	return &FitzProvider{}
}

func (fp *FitzProvider) New(filePath string) error {
	doc, err := fitz.New(filePath)
	if err != nil {
		return err
	}

	fp.doc = doc
	return nil
}

func (fp *FitzProvider) GetText() ([]string, error) {
	defer fp.doc.Close()
	pages := make([]string,0)
	for n := 0; n < fp.doc.NumPage(); n++ {
		text, err := fp.doc.Text(n)
		if err != nil {
			return nil, err
		}

		pages = append(pages, text)
	}

	return pages, nil
}