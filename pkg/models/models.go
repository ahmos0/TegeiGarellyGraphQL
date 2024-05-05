package models

type Item struct {
	UUID     string `json:"uuid"`
	WorkName string `json:"works"`
	Author   string `json:"author"`
	ImageURL string `json:"url"`
	Other    string `json:"other"`
}
