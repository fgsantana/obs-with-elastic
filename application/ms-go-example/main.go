package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Person struct {
	Name string            `json:"name"`
	Age  int               `json:"age"`
	Map  map[string]string `json:"map"`
}

func main() {

	http.HandleFunc("/go", test)

	http.ListenAndServe(":8880", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)

	body := Person{}
	err := json.Unmarshal(reqBody, &body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(fmt.Errorf("Error when parsing request body to json: %s ", err))
		return
	}

	log.Printf("JSON received:\n%s", reqBody)

	methodFromServer := doReq(r.Method)

	var resBody = make(map[string]any)
	resBody["person"] = body
	resBody["method"] = methodFromServer

	res, _ := json.Marshal(resBody)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func doReq(method string) string {
	log.Printf("Sending [%s] request to ms-kt server", method)
	req, _ := http.NewRequest(method, "http://ms-kt:8085/kt", nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)

	log.Printf("Received status code [ %s] from ms-kt server \nResponse body: %s", res.Status, resBody)

	return checkResBody(resBody)

}

func checkResBody(resBody []byte) string {
	if strings.Contains(string(resBody), "GET") {
		return "GET"
	}
	if strings.Contains(string(resBody), "POST") {
		return "POST"
	}

	return "Return from ms-kt server not supported"
}
