package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getDadJokePlainText() string {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://icanhazdadjoke.com/", nil)

	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	req.Header.Set("Accept", "text/plain")

	res, getErr := client.Do(req)
	if getErr != nil {
		fmt.Printf("Error %s\n", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	return string(body)
}

func getDadJokeJSON() string {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://icanhazdadjoke.com/", nil)

	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	req.Header.Set("Accept", "application/json")

	res, getErr := client.Do(req)
	if getErr != nil {
		fmt.Printf("Error %s\n", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	return string(body)
}

func getMultipleDadJokes(numberOfJokes int, format string) []string {
	jokes := make([]string, numberOfJokes)
	for i := 0; i < numberOfJokes; i++ {
		if format == "json" {
			jokes[i] = getDadJokeJSON()
		} else if format == "plaintext" {
			jokes[i] = getDadJokePlainText()
		}
	}

	return jokes
}

func main() {

	formatPtr := flag.String("format", "json", "The format to return the joke in.")
	numberOfJokesPtr := flag.Int("jokes", 1, "The number of jokes to return.")

	flag.Parse()

	numberOfJokes := *numberOfJokesPtr
	format := *formatPtr

	jokes := getMultipleDadJokes(numberOfJokes, format)

	for i := 0; i < numberOfJokes; i++ {
		fmt.Printf("Joke #%d: %s\n", i+1, jokes[i])
	}
}
