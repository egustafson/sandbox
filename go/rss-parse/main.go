package main

import (
	"fmt"
	rss "github.com/jteeuwen/go-pkg-rss"
	// "github.com/jteeuwen/go-pkg-xmlx"
	"log"
)

func main() {
	const timeout = 5
	const uri = "http://feeds.twit.tv/tnt_video_large.xml"

	feed := rss.New(timeout, true, chanHandler, itemHandler)
	if err := feed.Fetch(uri, nil); err != nil {
		log.Printf("[error] %s: %s\n", uri, err)
	}
	fmt.Println("done.")
}

func chanHandler(feed *rss.Feed, newchans []*rss.Channel) {
	for _, ch := range newchans {
		fmt.Printf("Channel: %s\n", ch.Title)
	}
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	for _, item := range newitems {
		fmt.Printf("Item: (%s) %s\n", item.PubDate, item.Title)
		for _, encl := range item.Enclosures {
			fmt.Printf("  Encl: (%s) %s\n", encl.Type, encl.Url)
		}
	}
}
