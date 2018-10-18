package albertapi

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"
)

// metadata of the extension
type metadata struct {
	IID          string   `json:"iid"`
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	Author       string   `json:"author"`
	Dependencies []string `json:"dependencies"`
	Trigger      string   `json:"trigger"`
}

// Item is an item in the response
type Item struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Icon        string   `json:"icon"`
	Actions     []Action `json:"actions"`
}

// Action is an action for an Item
type Action struct {
	Name      string   `json:"name"`
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`
}

// queryResponse correctly formats the response of a QUERY request
type queryResponse struct {
	Items []Item `json:"items"`
}

// Handle should be the only function that is being called from main() directly
func Handle(name, version, author string, dependencies []string, triggerword string, loadStateFunc func() error, saveStateFunc func(), queryFunc func(query string) []Item) {
	switch os.Getenv("ALBERT_OP") {
	case "METADATA":
		b, _ := json.Marshal(metadata{
			IID:          "org.albert.extension.external/v3.0",
			Name:         name,
			Version:      version,
			Author:       author,
			Dependencies: dependencies,
			Trigger:      triggerword + " ",
		})
		io.Copy(os.Stdout, bytes.NewReader(b))
	case "INITIALIZE":
		if loadStateFunc != nil {
			if err := loadStateFunc(); err != nil {
				os.Exit(1)
			}
		}
	case "FINALIZE":
		if saveStateFunc != nil {
			saveStateFunc()
		}
	case "QUERY":
		b, _ := json.Marshal(queryResponse{Items: queryFunc(strings.TrimPrefix(os.Getenv("ALBERT_QUERY"), triggerword+" "))})
		io.Copy(os.Stdout, bytes.NewReader(b))
	default:
		io.Copy(os.Stdout, bytes.NewReader([]byte("This is an external Albert extension.\n")))
	}
}
