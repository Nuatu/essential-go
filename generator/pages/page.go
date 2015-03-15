package pages

import (
	"io/ioutil"

	"github.com/russross/blackfriday"
)

type Page struct {
	Title  string
	Author string
	HTML   string
	File   string
}

func NewPage(filename string) (*Page, error) {
	p := &Page{File: filename}

	// This return statement calls p.load first
	// then returns p
	return p, p.load()
}

func (p *Page) load() error {
	c, err := LoadConfig("config.json")
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(p.File)
	if err != nil {
		return err
	}

	p.HTML = string(blackfriday.MarkdownCommon(data))
	p.Title = c.Name()
	p.Author = c.Author()

	return nil
}
