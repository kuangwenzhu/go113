package main

import (
	"log"
	"net/http"

	"github.com/ServiceComb/go-chassis"
	"github.com/ServiceComb/go-chassis/server/restful"
)

type Service struct{}

func (s *Service) HelloWorld(c *restful.Context) {
	c.Write([]byte("hello world"))
}

func (r *Service) URLPatterns() []restful.Route {
	return []restful.Route{
		{http.MethodGet, "/go113", "HelloWorld"},
	}
}

func main() {
	chassis.RegisterSchema("rest", &Service{})
	if err := chassis.Init(); err != nil {
		log.Print("Init failed", err)
		return
	}
	err := http.ListenAndServe(":8090", nil)
        if err != nil {
          log.Print("ListenAndServe: ", err.Error())
     	}

	chassis.Run()
}
