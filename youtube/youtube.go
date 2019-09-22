package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Response define the JSON response structure via Youtube API
type Response struct {
	Kind  string  `json:"kind"`
	Items []Items `json:"items"` // slice of items
}

// Items define the structure of each element of an Items slice
type Items struct {
	Kind  string `json:"kind"`
	ID    string `json:"ID"`
	Stats Stats  `json:"statistics"`
}

// Stats define the Statistics attribute
type Stats struct {
	Views       string `json:"viewcount"`
	Subscribers string `json:"subscribercount"`
	Videos      string `json:"videocount"`
}

// GetSubscribers handles the actual interaction with the Youtube API
func GetSubscribers() (Items, error) {

	newRequest, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil) // instantiate a new http before adding query paramters e.g. API Key
	if err != nil {
		log.Println("Unable to instantiate a new httpRequest.", err)
		return Items{}, err // return an empty Items struct and the error
	}

	// start populating the newRequest parameter fields, upon successful instantiation
	queryParameters := newRequest.URL.Query() //instantiate the query object
	queryParameters.Add("key", os.Getenv("YOUTUBE_KEY"))
	queryParameters.Add("id", os.Getenv("CHANNEL_ID"))
	queryParameters.Add("part", "statistics")

	newRequest.URL.RawQuery = queryParameters.Encode() // encode into a URL encoded format seen in browsers

	// instantiate an http Client to connect to Youtube API
	newClient := &http.Client{}
	apiResponse, err := newClient.Do(newRequest) // pass the updated `newRequest` for the client to make the connection
	if err != nil {
		log.Println("Unable to connect to Youtube's API.", err)
		return Items{}, err
	}
	defer apiResponse.Body.Close() // keep the response open until we can extract data from it

	fmt.Println("Response Status.", apiResponse.Status)

	apiResponseBody, err := ioutil.ReadAll(apiResponse.Body) // extract the response body data and returns as bytes
	if err != nil {
		log.Println("Unable to read the apiResponseBody data.", err)
		return Items{}, err
	}

	var response Response
	err = json.Unmarshal(apiResponseBody, &response) // decode via UnMarshall since the data is already bytes JSON format in memory [i.e. not from a ioReader]
	if err != nil {
		log.Println("Unable to unmarshall the read apiResponseBody data.", err)
		return Items{}, err
	}
	return response.Items[0], nil // we are returning the first item in the body of data
}
