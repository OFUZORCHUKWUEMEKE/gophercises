package main

import (
	"flag"
	"fmt"
	cyoa "gophercises/_3_ADVENTURE/story"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
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

	h := cyoa.NewHandler(story)

	fmt.Printf("Starting the server at: &d\n", *port)

	fmt.Printf("%v+v\n", story)

	log.fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
