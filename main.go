package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	PhotoApi = "http://api.pexels.com/v1"
	VideoApi = "http://api.pexels.com/video"
)

type Client struct {
	Token          string
	hc             http.Client
	RemainingTimes int32
}

func NewClient(token string) *Client {
	c := http.Client{}
	return &Client{Token: token, hc: c}
}

type SerachResults struct {
	Page         int32   `json:"page"`
	Perpage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results`
	NextPage     string  `json: "next_page"`
	Photos       []Photo `json: "photos"`
}

type Photos struct {
	Id             int         `json:"id"`
	Width          int         `json:"width"`
	Height         int         `json:"height"`
	Url            string      `json:"url"`
	Photographer   string      `json:"photographer"`
	PhotographeUrl string      `json:"photographer_url"`
	Source         PhotoSource `json: "src"`
}

type PhotoSource struct {
}

func main() {

	os.Setenv("pexelsToken", "wDpFgtSdGHlK6PHKZXxfEmb1TVIPJIywEoTyPwlYxRysmlRAGc6gZd9b")
	var TOKEN = os.Getenv("pexelsToken")

	var c = NewClient(TOKEN)

	results, err := c.SearchPhotos("waves")

	if err != nil {
		fmt.Errorf("search error:%v", err)
	}
	if result.Page == 0 {
		fmt.Errorf("search result is wrong")
	}
	fmt.Println("results")
}
