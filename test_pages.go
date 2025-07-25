package main

import (
	"blog/model"
	"blog/conf"
	"fmt"
)

func main() {
	conf.Init()
	model.Init()
	pages, err := model.PostPagesAll()
	if err \!= nil {
		fmt.Printf("Error: %v
", err)
		return
	}
	fmt.Printf("Pages: %v
", pages)
	for _, page := range pages {
		fmt.Printf("Page - ID: %d, Title: %s, Path: %s
", page.Id, page.Title, page.Path)
	}
}
