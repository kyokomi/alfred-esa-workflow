package main

import (
	alfred "github.com/pascalw/go-alfred"
)

func alfredPrintError(err error) {
	response := alfred.NewResponse()
	response.AddItem(&alfred.AlfredResponseItem{
		Valid: true,
		Uid:   "1",
		Title: err.Error(),
	})
	response.Print()
}

func alfredPrintItems(items []*alfred.AlfredResponseItem) {
	response := alfred.NewResponse()
	for _, item := range items {
		response.AddItem(item)
	}
	response.Print()
}

func alfredPrintMessage(message string) {
	response := alfred.NewResponse()
	response.AddItem(&alfred.AlfredResponseItem{
		Valid: true,
		Uid:   "1",
		Title: message,
	})
	response.Print()
}
