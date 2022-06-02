package main

import (
	sub "Rabbit-GOPkg/Consumer"
	pub "Rabbit-GOPkg/Publisher"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PostRequestBody struct {
	QueueName string `json:"queueName"`
	Letter    string `json:letter`
}

func faceError(w http.ResponseWriter, err error) {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	// return
}

func PublishMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		errorMessage := map[string]string{"Message": "POST method expected !!"}
		jsonValue, err := json.Marshal(errorMessage)
		faceError(w, err)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(jsonValue)
	}

	var reqBody PostRequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	faceError(w, err)

	resp, err := pub.Publish(reqBody.QueueName, []byte(reqBody.Letter))
	faceError(w, err)

	val, err := json.Marshal(resp)
	faceError(w, err)

	// w.WriteHeader(http.StatusAccepted)
	w.Write(val)
}

func ConsumerMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		errorMessage := map[string]string{"Message": "POST method expected !!"}
		jsonValue, err := json.Marshal(errorMessage)
		faceError(w, err)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(jsonValue)
	}

	queueName := r.URL.Query().Get("queueName")
	resp, err := sub.Consume(queueName)
	faceError(w, err)

	// w.WriteHeader(http.StatusAccepted)
	w.Write(resp)
}

func main() {
	fmt.Printf("Starting server...\n")
	http.HandleFunc("/publish", PublishMessage)
	http.HandleFunc("/consume", ConsumerMessage)

	log.Fatal(http.ListenAndServe(":8085", nil))
}
