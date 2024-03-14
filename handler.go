package main

/*
	Deployment

	Deployment to Azure for running Golang using Azure Functions is not ideal. The deployment process currently requires manual effort as 'Deployment Center' is not available.
	https://learn.microsoft.com/en-us/azure/azure-functions/create-first-function-vs-code-other?tabs=go%2Cwindows#compile-the-custom-handler-for-azure

	In order to make this an automated step, Buildkite or Github Actions would need to be setup with the following steps:
		- 	Run IaC (Terraform, Bicep, ARM etc.) to create Azure Functions resource.
			-	Example name: 'cash-links-functions'.
		-	Install Golang Runtime Dependencies
		-	Run
				set GOOS=linux
				set GOARCH=amd64
				go build handler.go
		-	Deploy to existing 'cash-links-functions' resource.
*/

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := "This HTTP triggered function executed successfully. Pass a name in the query string for a personalized response.\n"
	name := r.URL.Query().Get("name")

	if name != "" {
		message = fmt.Sprintf("Hello, %s. This HTTP triggered function executed successfully.\n", name)
	}

	fmt.Fprint(w, message)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	http.HandleFunc("/api/HttpExample", helloHandler)

	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
