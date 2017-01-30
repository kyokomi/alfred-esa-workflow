package main

import (
	"fmt"
	"os"

	alfred "github.com/pascalw/go-alfred"
)

// Alfred alfred printer
type Alfred struct {
}

// PrintError filter script print error
func (a Alfred) PrintError(err error) {
	response := alfred.NewResponse()
	response.AddItem(&alfred.AlfredResponseItem{
		Valid: true,
		Uid:   "1",
		Title: err.Error(),
	})
	response.Print()
}

// PrintItems filter script print items
func (a Alfred) PrintItems(items []*alfred.AlfredResponseItem) {
	response := alfred.NewResponse()
	for _, item := range items {
		response.AddItem(item)
	}
	response.Print()
}

// PrintMessage filter script print one message
func (a Alfred) PrintMessage(message string) {
	response := alfred.NewResponse()
	response.AddItem(&alfred.AlfredResponseItem{
		Valid: true,
		Uid:   "1",
		Title: message,
	})
	response.Print()
}

// Message normal response message
func (a Alfred) Message(message string) {
	fmt.Fprint(os.Stdout, message)
}

// Error normal response error message
func (a Alfred) Error(err error) {
	fmt.Fprint(os.Stdout, err)
}
