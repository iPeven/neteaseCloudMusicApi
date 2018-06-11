package ncm

import (
	"encoding/json"
	"log"
)

type Search struct{
	S string `json:"s"`
	Type int `json:"type"`
	Offset int `json:"offset"`
	Total bool `json:"total"`
	Limit int `json:"limit"`
	Csrf_token string `json:"csrf_token"`
}

type Comments struct{
	Rid string `json:"rid"`
	Offset int `json:"offset"`
	Total bool `json:"total"`
	Limit int `json:"limit"`
	Csrf_token string `json:"csrf_token"`
}

type Detail struct{
	Id int `json:"id"`
	C []C `json:"c"`
	Csrf_token string `json:"csrf_token"`
}

type C struct{
	Id int `json:"id"`
}

type SongMP3 struct {
	Ids []int `json:"ids"`
	Br int `json:"br"`
	Csrf_token string `json:"csrf_token"`
}

type Lyric struct {
	Id int `json:"id"`
	Os string `json:"os"`
	Lv int `json:"lv"`
	Kv int `json:"kv"`
	Tv int `json:"tv"`
}

func (search Search)EncStr() string{
	jsonData, err := json.Marshal(search)
	if err != nil{
		log.Fatal(err)
	}
	return string(jsonData)
}

func (comments Comments)EncStr() string{
	jsonData, err := json.Marshal(comments)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonData)
}

func (detail Detail)EncStr() string{
	jsonData, err := json.Marshal(detail)
	if err != nil{
		log.Fatal(err)
	}
	return string(jsonData)
	}

func (songMp3 SongMP3)EncStr() string{
	jsonData, err := json.Marshal(songMp3)
	if err != nil{
		log.Fatal(err)
	}
	return string(jsonData)
}

func (lyric Lyric)EncStr() string{
	jsonData, err := json.Marshal(lyric)
	if err != nil{
		log.Fatal(err)
	}
	return string(jsonData)
}

