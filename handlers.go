package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}

func HandleHome(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func PostRequest(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var metadata MetaData

	err := decoder.Decode(&metadata)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "payload %v\n", metadata)
}

func UserPostRequest(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var user User

	err := decoder.Decode(&user)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := user.ToJson()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
