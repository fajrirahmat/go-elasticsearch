package client

import (
	"encoding/json"
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

//GetIndexWithTypeAndID get document index with index name, document type and document ID defined
func (g *Context) GetIndexWithTypeAndID(indexName string, documentType string, documentID string) (SearchResponse, error) {
	response := SearchResponse{}
	url := "/" + indexName + "/" + documentType + "/" + documentID
	resp, err := g.C.Get(url)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

//GetIndexWithTypeAndIDAndRouting get document index with index name, document type and document ID defined
func (g *Context) GetIndexWithTypeAndIDAndRouting(indexName string, documentType string, documentID string, routing string) (SearchResponse, error) {
	response := SearchResponse{}
	url := "/" + indexName + "/" + documentType + "/" + documentID + "?routing=" + routing
	resp, err := g.C.Get(url)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
