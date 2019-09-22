package websocket

import (
	"encoding/json"
	"fmt"
	"goSocialMediaSubscriberMonitor/youtube"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// define an upgrader i.e. attributes of the upgrader used to convert the http.conn to websocket.conn
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Upgrade creates a wsConn from normal http connections
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) { // returns a websocket connection and error value

	// required to handle CORS i.e. check if request from a different domain is allowed to connect
	upgrader.CheckOrigin = func(r *http.Request) bool { // the anonymous function returns a true or false i.e. true being allowed
		return true //let's assume for testing purposes that the requesting domain is allowed, reqgardless of origin/source
	}

	wsConn, err := upgrader.Upgrade(w, r, nil) //convert the http to ws i.e. convert the w to wsConn
	if err != nil {
		log.Println("Unable to connect to websocket.", err)
		return wsConn, err // wsConn pointing to an address with no value
	}

	return wsConn, nil // wsConn pointing to an address with a value
}

// Writer uses the established wsConn to start calling youtube.GetSubscribers() every 5s
func Writer(wsConn *websocket.Conn) {
	for { // for loop is always true i.e. always active

		newTicker := time.NewTicker(5 * time.Second) // instantiate a new ticker timer

		// use the time `ticks` as a trigger to active the `youtube.GetSubscribers()`
		for tick := range newTicker.C { // i.e. whenever there is a `tick` in the channel `C` ... do...

			fmt.Printf("Updating Stats: %+v\n", tick) // notify user of stats update

			items, err := youtube.GetSubscribers() // retrieve the subscribers stats
			if err != nil {
				log.Println("Failed to retrieve Subscriber stats.", err)
			}

			// encode the retrieved data to a JSON string
			jsonString, err := json.Marshal(items)
			if err != nil {
				log.Println("Failed to encode the retrieved Subscriber stats.", err)
			}

			// we now send the encoded data back to the websocket connection i.e. as bytes of textmessage
			err = wsConn.WriteMessage(websocket.TextMessage, []byte(jsonString))
			if err != nil { // that means message as NOT successfully written back
				log.Println("Failed to send the retrieved Subscriber stats.", err)
				return
			}
		}
	}
}
