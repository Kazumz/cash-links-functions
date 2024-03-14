# cash-links-functions
A Web API that returns cash-rewarding referral links.

https://cash-links-functions.azurewebsites.net/api/HttpExample

## Run
Execute in Git Bash
./run-windows.bat

For windows, ensure that the linux version of 'handler' does not co-exist locally else this will cause runtime lockups.

## Deployment
Execute in Git Bash
./deploy-linux.bat

Use VSCode Azure Function Tools to deploy directly to the Azure Function Resource.

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