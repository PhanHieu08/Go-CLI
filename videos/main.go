package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getAll := getCmd.Bool("all", false, "Get all videos")
	getID := getCmd.String("id", "", "Youtube video ID")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addID := addCmd.String("id", "", "Youtube video ID")
	addTitle := addCmd.String("title", "", "Youtube video title")
	addUrl := addCmd.String("url", "", "Youtube video url")
	addImageUrl := addCmd.String("imageurl", "", "Youtube video image url")
	addDesc := addCmd.String("desc", "", "Youtube video description")

	if len(os.Args) < 2 {
		fmt.Println("expect 'get' or 'add' command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getAll, getID)

	case "add":
		HandleAdd(addCmd, addID, addTitle, addUrl, addImageUrl, addDesc)

	default:
	}

}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
		fmt.Println("id is required or '--all' for all video")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		videos := getVideos()
		fmt.Println("ID \t Title \t URL \t ImageURL \t Description \n")

		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.ImageUrl, video.Description)
		}

		return
	}

	if *id != "" {
		videos := getVideos()
		id := *id

		for _, video := range videos {
			if id == video.Id {
				fmt.Println("ID \t Title \t URL \t ImageURL \t Description \n")
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.ImageUrl, video.Description)

			}
		}
	}
}

func ValidateVideo(addCmd *flag.FlagSet, id *string, title *string, url *string, imageurl *string, desc *string) {
	if *id != "" || *title != "" || *url != "" || *imageurl != "" || *desc != "" {
		fmt.Println("all fields required")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
}

func HandleAdd(addCmd *flag.FlagSet, id *string, title *string, url *string, imageurl *string, desc *string) {
	ValidateVideo(addCmd, id, title, url, imageurl, desc)

	video := Video{
		Id:          *id,
		Title:       *title,
		Description: *desc,
		Url:         *url,
		ImageUrl:    *imageurl,
	}

	videos := getVideos()
	videos = append(videos, video)

	saveVideos(videos)
}
