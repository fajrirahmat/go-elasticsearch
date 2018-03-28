package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

//SearchResponse type response from elasticsearch. hits
type SearchResponse struct {
	Index   string      `json:"_index"`
	Type    string      `json:"_type"`
	ID      string      `json:"_id"`
	Version int64       `json:"_version"`
	Found   bool        `json:"found"`
	Source  interface{} `json:"_source"`
}

//DocumentHeader document header that will send as request body for Multi Get process
type DocumentHeader struct {
	Index       string      `json:"_index"`
	Type        string      `json:"_type,omitempty"`
	ID          string      `json:"_id"`
	Source      interface{} `json:"_source,omitempty"`
	StoreFields []string    `json:"stored_fields,omitempty"`
	Routing     string      `json:"_routing,omitempty"`
}

//MultiGetResponse response for Multi Get
type MultiGetResponse struct {
	Docs []SearchResponse `json:"docs"`
}

//GetIndexWithTypeAndID get document index with index name, document type and document ID defined
func (c *Context) GetIndexWithTypeAndID(indexName string, documentType string, documentID string) (SearchResponse, error) {
	response := SearchResponse{}
	url := "/" + indexName + "/" + documentType + "/" + documentID
	resp, err := c.C.Get(url)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

//GetIndexWithTypeAndIDAndRouting get document index with index name, document type and document ID defined
func (c *Context) GetIndexWithTypeAndIDAndRouting(indexName string, documentType string, documentID string, routing string) (SearchResponse, error) {
	response := SearchResponse{}
	url := "/" + indexName + "/" + documentType + "/" + documentID + "?routing=" + routing
	resp, err := c.C.Get(url)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

//MultiGet Multi Get API
//Reference: https://www.elastic.co/guide/en/elasticsearch/reference/5.5/docs-multi-get.html
func (c *Context) MultiGet(docs []DocumentHeader) (MultiGetResponse, error) {
	url := "/_mget"
	return processMultiGet(c, url, docs)
}

//MultiGetWithIndex Multi Get API with defined index name
//Reference: https://www.elastic.co/guide/en/elasticsearch/reference/5.5/docs-multi-get.html
func (c *Context) MultiGetWithIndex(indexName string, docs []DocumentHeader) (MultiGetResponse, error) {
	url := "/" + indexName + "/_mget"
	return processMultiGet(c, url, docs)
}

//MultiGetWithIndexAndType Multi Get API with additional parameter index name and docs
//Referene: https://www.elastic.co/guide/en/elasticsearch/reference/5.5/docs-multi-get.html
func (c *Context) MultiGetWithIndexAndType(indexName, docType string, docs []DocumentHeader) (MultiGetResponse, error) {
	url := "/" + indexName + "/" + docType + "/_mget"
	return processMultiGet(c, url, docs)
}

func processMultiGet(c *Context, url string, docs []DocumentHeader) (MultiGetResponse, error) {
	response := MultiGetResponse{}
	body := map[string]interface{}{
		"docs": docs,
	}
	b, err := json.Marshal(&body)
	if err != nil {
		log.Println(err)
		return response, err
	}
	request, _ := http.NewRequest("GET", url, bytes.NewBuffer(b))
	resp, err := c.C.Do(request)
	if err != nil {
		log.Println(err)
		return response, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
