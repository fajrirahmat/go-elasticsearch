package client

import (
	"log"
	"net/http"

	"github.com/olivere/balancers"
	"github.com/olivere/balancers/roundrobin"
)

//Context struct to store http.client
type Context struct {
	C *http.Client
}

//New create balancer if Elasticsearch has more than one client node
func New(urls ...string) *Context {
	balancer, err := roundrobin.NewBalancerFromURL(urls...)
	if err != nil {
		log.Fatal(err)
	}
	c := balancers.NewClient(balancer)
	return &Context{C: c}
}
