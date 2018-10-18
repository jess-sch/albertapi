package main

import "github.com/lumi-sch/albertapi"

func main() {
	albertapi.Handle(
		"Albertapi",
		"1.0",
		"Lumi Schallenberg",
		[]string{"xclip"},
		"e",
		nil,
		nil,
		func(s string) (r []albertapi.Item) {
			r = append(r, albertapi.Item{
				ID:          "0",
				Name:        s,
				Description: s,
				Icon:        "download",
				Actions: []albertapi.Action{albertapi.Action{
					Name:      "s",
					Command:   "s",
					Arguments: []string{},
				}},
			})
			return
		},
	)
}
