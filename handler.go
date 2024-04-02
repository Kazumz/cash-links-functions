package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	// NOTE: In order for this to work in the cloud (AWS or Azure), you need to authenticate with OAuth and use a bearer token.
	req, err := http.NewRequest("GET", "https://oauth.reddit.com/r/beermoneyuk/hot.json?limit=25", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
	}(resp.Body)

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Set Writer to expect JSON, therefore consumers.
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(body))
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	http.HandleFunc("/api/HttpExample", getAllHandler)

	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
