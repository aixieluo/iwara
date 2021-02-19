package models

type Video struct {
	Model
	Url    string `json:"url"`
	Poster string `json:"poster"`
	Title  string `json:"title"`
	View int `json:"view"`
	Star int `json:"star"`
	HashId string `json:"hash_id"`
}

type Videos []Video
