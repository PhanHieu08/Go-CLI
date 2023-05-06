package main

import (
	"encoding/json"
	"io/ioutil"
)

type Video struct {
	Id          string
	Title       string
	Description string
	Url         string
	ImageUrl    string
}

func getVideos() []Video {
	fileBytes, err := ioutil.ReadFile("./videos.json")

	if err != nil {
		panic(err)
	}

	var videos []Video

	err = json.Unmarshal(fileBytes, &videos)

	if err != nil {
		panic(err)
	}

	return videos

}

func saveVideos(videos []Video) {
	videoBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./videos.json", videoBytes, 0644)

	if err != nil {
		panic(err)
	}
}
