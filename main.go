package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/Dragon-taro/learn-status-code/status"
)

func getStatusCode() int {
	rand.Seed(time.Now().UnixNano())
	codeArray := []int{200, 200, 200, 200, 200, 200, 201, 202, 301, 400, 401, 403, 404, 500, 502, 504}

	return codeArray[rand.Intn(len(codeArray))]
}

func handleStaus() (int, string) {
	code := getStatusCode()
	return code, status.StatusText(code)
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
