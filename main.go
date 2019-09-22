package main

import (
	"fmt"
	"goSocialMediaSubscriberMonitor/websocket"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePageEndpointHandler(w http.ResponseWriter, r *http.Request) {
	dataToRender := "Home Page"
	io.WriteString(w, dataToRender)
}

func statsEndpointHandler(w http.ResponseWriter, r *http.Request) {

	wsConn, err := websocket.Upgrade(w, r) //convert the http to ws i.e. convert the w to wsConn
	if err != nil {
		fmt.Fprintf(w, "Failed to connect via WebSockets %+v\n.", err)
	}
	go websocket.Writer(wsConn) // start writing the stats
}

func serviceRequestHandlers() {
	newRouter := mux.NewRouter().StrictSlash(true)
	newRouter.HandleFunc("/", homePageEndpointHandler)
	newRouter.HandleFunc("/stats", statsEndpointHandler)
	log.Fatal(http.ListenAndServe(":8080", newRouter))
}

func main() {
	fmt.Println("Social Media Subscriber Monitor")
	serviceRequestHandlers()
}
