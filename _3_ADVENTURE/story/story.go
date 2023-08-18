package cyoa

import (
	"encoding/json"
	"fmt"
	"io"
)

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	fmt.Println(d)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
