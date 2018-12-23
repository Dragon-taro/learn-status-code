package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Dragon-taro/learn-status-code/status"
)

func getStatusCode() int {
	rand.Seed(time.Now().UnixNano())
	codeArray := []int{200, 200, 200, 200, 200, 200, 201, 205, 301, 400, 401, 402, 403, 404, 405, 413, 414, 429, 500, 502, 503, 504}

	return codeArray[rand.Intn(len(codeArray))]
}

func handleStaus() (int, string) {
	code := getStatusCode()
	body := strconv.Itoa(code) + ": " + status.StatusText(code)
	return code, body
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	status, body := handleStaus()
	w.WriteHeader(status)
	fmt.Fprint(w, body)
	return
}

func handleRequestUser(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		status, _ := handleStaus()
		w.WriteHeader(status)
		return
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/user", handleRequestUser)
	http.ListenAndServe(":8080", nil)
}
