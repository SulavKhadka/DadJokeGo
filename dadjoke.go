package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type joke struct {
	id     string
	joke   string
	status int
}

func main() {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://icanhazdadjoke.com/", nil)

	if err != nil {
		fmt.Println("Error %s\n", err)
	}
	req.Header.Set("Accept", "application/json")

	res, getErr := client.Do(req)
	if getErr != nil {
		fmt.Println("Error 2 %s\n", err)
	}

	dadJoke := new(joke)

	newJoke := json.NewDecoder(res.Body).Decode(dadJoke)

	fmt.Println(newJoke)
}
