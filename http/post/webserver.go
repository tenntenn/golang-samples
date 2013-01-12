package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/xml"
)

type Person struct {
	Name string
	Age int
	ServerResponse bool
}

func index(w http.ResponseWriter, r *http.Request) {

	var person Person
	body, _ := ioutil.ReadAll(r.Body)
	xml.Unmarshal(body, &person)
	person.ServerResponse = true

	responseXML, _ := xml.Marshal(person)
	fmt.Fprintf(w, string(responseXML))
}

func main() {
	http.HandleFunc("/", index);
	http.ListenAndServe(":8080", nil)
}
