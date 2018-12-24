package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Dragon-taro/learn-status-code/db"
	"github.com/Dragon-taro/learn-status-code/status"
)

func getStatusCode() int {
	rand.Seed(time.Now().UnixNano())
	codeArray := []int{200, 200, 200, 200, 200, 200, 201, 205, 301, 400, 401, 402, 403, 404, 405, 413, 414, 429, 502, 503}

	return codeArray[rand.Intn(len(codeArray))]
}

func createBody(code int) string {
	return strconv.Itoa(code) + ": " + status.StatusText(code)
}

func handleStaus() (int, string) {
	code := getStatusCode()
	body := createBody(code)
	return code, body
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		code := http.StatusMethodNotAllowed
		body := createBody(code)
		w.WriteHeader(code)
		fmt.Fprint(w, body)
		return
	}

	status, body := handleStaus()
	w.WriteHeader(status)
	fmt.Fprint(w, body)
	return
}

func handleRequestUser(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		_, err := db.Connect()
		if err != nil {
			code := http.StatusInternalServerError
			body := createBody(code)
			w.WriteHeader(code)
			fmt.Fprint(w, body)
			return
		}
	}

	code := http.StatusPaymentRequired
	body := createBody(code)
	w.WriteHeader(code)
	fmt.Fprint(w, body)
	return
}

func handleRequestChiro(w http.ResponseWriter, req *http.Request) {
	code := http.StatusGatewayTimeout
	body := createBody(code)
	w.WriteHeader(code)
	fmt.Fprint(w, body)
	time.Sleep(3 * time.Second)
	return
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/user", handleRequestUser)
	http.HandleFunc("/chiro", handleRequestChiro)
	http.ListenAndServe(":8080", nil)
}
