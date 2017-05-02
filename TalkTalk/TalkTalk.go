/* tick-tock.go */
package main

import (
	"fmt"
	"log"
	"net/http"
)

var chat = make(map[string]string)

func post_message(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var message string = request.Form.Get("Message")
	var username string = request.Form.Get("Username")
	var channel string = request.Form.Get("Channel")
	if message != "" {
		if _, exists := chat[channel]; exists {

			chat[channel] += username + " ~> " + message + "\n"
		} else {
			chat[channel] = username + " ~> " + message + "\n"
		}
		fmt.Println(channel)
	}
	fmt.Fprintf(writer, chat[channel])
}

func checkMessage(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var channel string = request.Form.Get("Channel")
	var username string = request.Form.Get("Username")
	if _, exists := chat[channel]; !exists {
		chat[channel] = username + " Created " + channel + "\n"
	}
	fmt.Fprintf(writer, chat[channel])
	fmt.Println(channel)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".static")))
	http.HandleFunc("/post_message", post_message)
	http.HandleFunc("/messages", checkMessage)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
