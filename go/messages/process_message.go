package messages

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/csarnataro/miniproject/utils"
)

// ProcessMessage processes the message received from redis.
// It should invoke the corresponding remote service and it
// should log the result
func ProcessMessage(message Message) {
	actualEndpointURL := utils.BuildURL(message.Endpoint.URL, message.Data)
	fmt.Printf("Calling %s\n", actualEndpointURL)
	// let's start with GET
	response, err := http.Get(actualEndpointURL)
	if err != nil {
		fmt.Printf("Error reading from %s: %s\n", actualEndpointURL, err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		response.Body.Close()
	}

	fmt.Println(actualEndpointURL)
}
