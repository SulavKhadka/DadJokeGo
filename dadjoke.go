package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type joke struct {
	ID     string
	Joke   string
	Status int
}

func getDadJoke() joke {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://icanhazdadjoke.com/", nil)

	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	req.Header.Set("Accept", "application/json")

	res, getErr := client.Do(req)
	if getErr != nil {
		fmt.Printf("Error 2 %s\n", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	var dadJoke joke
	json.Unmarshal([]byte(string(body)), &dadJoke)

	return dadJoke
}

func main() {
	fmt.Println(getDadJoke().Joke)
}
