package markdown

import (
	"io/ioutil"

	"github.com/russross/blackfriday/v2"
)

type Markdown struct {
}

func New() *Markdown {
	return &Markdown{}
}

func (md *Markdown) Parse(filename string) (string, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(blackfriday.Run(f)), nil
}
