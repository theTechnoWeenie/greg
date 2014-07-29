package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Start(port int) {
	fmt.Printf("Starting greg on port %d\n", port)
	setupRestEndpoints()

	serverChannel := make(chan bool)
	startRegistry(port, serverChannel)
	<-serverChannel
}

func startRegistry(port int, serverChannel chan bool) {
	go func() {
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		serverChannel <- true
	}()
	go func() {
		pollServices()
		serverChannel <- true
	}()
}

func setupRestEndpoints() {
	http.HandleFunc("/register", registerCallback)
	http.HandleFunc("/uptime", uptimeCallback)
	http.HandleFunc("/services", servicesCallback)
	http.HandleFunc("/service", serviceLookupCallback)
}

var services []Servicer

func registerCallback(writer http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var service Service
	malformedError := json.Unmarshal(body, &service)
	if malformedError != nil {
		log.Printf("ERR: Malformed JSON %s", string(body))
	}
	services = append(services, service)
}
func uptimeCallback(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "<html><body><h1>Not Yet Implemented</h1></body></html>")
}
func servicesCallback(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "<html><body><h1>Not Yet Implemented</h1></body></html>")
}
func serviceLookupCallback(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "<html><body><h1>Not Yet Implemented</h1></body></html>")
}

func pollServices() {
	for true {
		fmt.Printf("Polling %d registered services\n", len(services))
		for index, service := range services {
			if !service.IsHealthy() { //remove the service from the list.
				services = services[:index+copy(services[index:], services[index+1:])]
			}
		}
		time.Sleep(time.Second * 10)
	}
}
