package game

import (
	"encoding/json"
	"io/ioutil"
)

type Story map[string]ChapterJson

type OptionsJson struct {
	Text        string `json:"text"`
	NextChapter string `json:"arc"`
}

type ChapterJson struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []OptionsJson `json:"options"`
}

func readFile(path string) ([]byte, error) {
	jsonData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func ReadJsonFromFile(path string) (*Story, error) {
	var jsonChapters Story
	jsonData, err := readFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonData, &jsonChapters)
	if err != nil {
		return nil, err
	}
	return &jsonChapters, nil
}
