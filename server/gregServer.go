package server

import (
    "fmt"
    "net/http"
    "time"
  )

type Service struct {
  Name string `json:"Name"`
  Ip string `json:"Ip"`
  Port int `json:"Port"`
}
func (s *Service) FullAddress() string{
  return fmt.Sprintf("%s:%d", s.Ip, s.Port)
}
func (s *Service) String() string {
  return fmt.Sprintf("%s<%s>", s.Name, s.FullAddress())
}

var services []Service

func Start(port int){
  fmt.Printf("Starting greg on port %d\n", port)
  http.HandleFunc("/register", registerCallback)
  http.HandleFunc("/uptime", uptimeCallback)
  http.HandleFunc("/services", servicesCallback)
  http.HandleFunc("/service", serviceLookupCallback)

  serverChannel := make(chan bool)
  go func() {
    http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
    serverChannel <- true
  }()
  pollServices()
  <-serverChannel
}

func registerCallback(writer http.ResponseWriter, r *http.Request){ fmt.Fprintf(writer, "<html><body><h1>Not Yet Implemented</h1></body></html>") }
func uptimeCallback(writer http.ResponseWriter, r *http.Request){ fmt.Fprintf(writer, "<html><body><h1>Not Yet Implemented</h1></body></html>") }
func servicesCallback(writer http.ResponseWriter, r *http.Request){ fmt.Fprintf(writer, "<html><body><h1>Not Yet Implemented</h1></body></html>") }
func serviceLookupCallback(writer http.ResponseWriter, r *http.Request){ fmt.Fprintf(writer, "<html><body><h1>Not Yet Implemented</h1></body></html>") }
func pollServices(){
  for true {
    fmt.Printf("Polling %d registered services", len(services))
    time.Sleep(time.Minute)
  }
}
