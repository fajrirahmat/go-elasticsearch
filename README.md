**Go Elasticsearch Client**
=======================
Elasticsearch client for Go Language.
The documention and library is in progress. This is not complete library or documentation

**Usage**
-----
First, get the library and put to GOPATH. Here how to do it:

    go get github.com/fajrirahmat/go-elasticsearch

Example: 

   

     package main
    
    import (
    	client "github.com/fajrirahmat/go-elasticsearch"
    )
    
    func main() {
    	c := client.New("http://localhost:9200")
    	query := map[string]interface{}{
    		"query": map[string]interface{}{
    			"term": map[string]interface{}{
    				"_id": "AWJnFHaxhMSe7XqCM5V0",
    			},
    		},
    	}
    	s := &client.SearchRequest{
    		IndexName:    "gotest",
    		DocumentType: "gotest",
    		Query:        query,
    	}
    	resp, _ := c.Search(s)
    }