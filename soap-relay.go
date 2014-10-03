package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Response map[string]interface{}

func main() {

	http.HandleFunc("/", homeHandler)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}

	res := postData(body)

	fmt.Println("i can see you", string(res))
}

func (r Response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // Must be called before writing response
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		body := processRequest(r)
		fmt.Fprint(w, Response{"status": true, "body": body})
	}
}

func processRequest(r *http.Request) string {
	var body []string
	body = r.Form["body"]
	fmt.Println(body)

	result := putMas(body[0])
	//	n := bytes.Index(result, []byte{0})
	//return encode(body[0])
	return string(result[:])

}

func postData(data string) []byte {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", "https://url", strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
	req.Header.Set("SOAPAction", "/operation")
	req.SetBasicAuth("user", "pass")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body

}
