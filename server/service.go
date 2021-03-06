package server

import (
	"fmt"
	"log"
	"net/http"
)

const (
	STATUS_RUNNING  = "RUNNING"
	STAUTS_DEGRADED = "DEGRADED"
	STATUS_STOPPED  = "STOPPED"
	STATUS_ERROR    = "ERROR"
)

type Servicer interface {
	FullAddress() string
	String() string
	IsHealthy() bool
}

type Service struct {
	Name string `json:"Name"`
	Ip   string `json:"Ip"`
	Port int    `json:"Port"`
}

func (s Service) FullAddress() string {
	return fmt.Sprintf("%s:%d", s.Ip, s.Port)
}

func (s Service) String() string {
	return fmt.Sprintf("%s<%s>", s.Name, s.FullAddress())
}

func (s Service) IsHealthy() bool {
	_, err := http.Get(fmt.Sprintf("http://%s:%d/healthCheck", s.Ip, s.Port))
	if err != nil {
		log.Printf("ERR: %s Caused By: %s", s.String(), err.Error())
		return false
	}
	log.Printf("OK : %s", s.String())
	return true
}
