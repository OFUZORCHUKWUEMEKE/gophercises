package main

import (
	"flag"
	"fmt"
	cyoa "gophercises/_3_ADVENTURE/story"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "the json file with the cyao")
	flag.Parse()

	fmt.Println("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v+v\n", story)

}
