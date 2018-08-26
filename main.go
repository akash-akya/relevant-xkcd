package main

import (
	"log"
	"fmt"
	"time"
	"net/http"
	"os"
)

var comics []XKCDComic

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

func main() {
    addr, err := determineListenAddress()
    if err != nil {
	log.Fatal(err)
    }

    loadComics()

    go func() {
	for {
	    crawl()
	    loadComics()
	    time.Sleep(crawlInterval)
	}
    }()
    fmt.Printf("Started background crawl\n");

    http.HandleFunc("/", IndexHandler)
    http.HandleFunc("/search", SearchHandler)

    // http.ListenAndServe(":8080", nil)
    if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
