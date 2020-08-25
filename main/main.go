package main

import (
	"fmt"
	root "github.com/hardy8059/A_game_of_choices_in_golang"
	utilities "github.com/hardy8059/A_game_of_choices_in_golang/game"
	"log"
	"net/http"
)

func main() {
	paths := root.InitialisePath()

	//1. Read Json from file
	storyAsJson, err := utilities.ReadJsonFromFile(paths.PathToJsonFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", storyAsJson)

	//2. Read HTML template from file
	htmlTemplate, err := utilities.ReadHtmlFromFile(paths.PathToHtmlFile)
	if err != nil {
		panic(err)
	}
	h, err := utilities.RouterHandler(*storyAsJson, htmlTemplate, paths.PathToStaticFiles)
	fmt.Println(paths.PathToStaticFiles)

	fmt.Println("Started server...")
	log.Fatal(http.ListenAndServe(":9000", h))
}
