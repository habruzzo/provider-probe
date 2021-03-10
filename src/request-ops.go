package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//make requests to nylas endpoint
const (
	detectionUrl string = "https://api.nylas.com/connect/detect-provider"
	clientId     string = ""
	clientSecret string = ""
)

type ProviderDetectionRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	EmailAddress string `json:"email_address"`
}

type ProviderDetectionResponse struct {
	AuthName     string
	Detected     bool
	EmailAddress string
	IsImap       bool
	ProviderName string
}

func SendRequest(emailAddress string) *http.Response {
	requestBody := ProviderDetectionRequest{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		EmailAddress: emailAddress,
	}
	requestBodyBytes, err := json.MarshalIndent(requestBody, "","\t")
	if err != nil {
		return nil
	}
	name := emailAddress //string(HashEmailForTimingLogs(emailAddress))
	RegisterTimer(name)
	StartTimer(name)
	response, err := http.Post(detectionUrl, "application/json", bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil
	}
	StopTimer(name)
	return response
}

func ValidateResponse(response *http.Response) bool {
	var providerResponse ProviderDetectionResponse
	if response == nil {
		return false
	}
	if response.StatusCode == http.StatusOK {
		responseBodyBytes, _ := ioutil.ReadAll(response.Body)

		err := json.Unmarshal(responseBodyBytes, &providerResponse)
		if err != nil {
			fmt.Println(err)
		}
		return true
	}
	return false
}
