package main

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

/*
*
Fetch RSS Feed data from URL using Http Client
*/
func urlToFeed(url string) (*RSSFeed, error) {
	httpClient := http.Client{Timeout: time.Second * 10}

	resp, err := httpClient.Get(url)
	if err != nil {
		log.Println("Error fetching RSS from this URL: ", url, err)
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading body RSS from this URL: ", url, err)
		return nil, err
	}

	var rssFeed RSSFeed

	err = xml.Unmarshal(respBody, &rssFeed)
	if err != nil {
		log.Println("Error parsing XML body RSS from this URL: ", url, err)
		return nil, err
	}

	return &rssFeed, nil
}
