package main

import (
	"encoding/json"
	"flag"
	"fmt"
	cyoa "gopher/adventure"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "Json file")
	flag.Parse()
	fmt.Printf("Using the stroy in %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
