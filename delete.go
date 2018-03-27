package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//Delete to delete elasticsearch document
func (c *Context) Delete(indexName string, documentType string, documentID string) bool {
	url := "/" + indexName + "/" + documentType + "/" + documentID
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return false
	}

	resp, err := c.C.Do(request)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	var f map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&f)

	if f["result"] != "deleted" {
		return false
	}
	return true
}

//DeleteByQuery to delete elasticsearch document by specific query
func (c *Context) DeleteByQuery(indexName string, query map[string]interface{}) bool {
	url := "/" + indexName + "/_delete_by_query"
	b, _ := json.Marshal(query)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return false
	}

	resp, err := c.C.Do(request)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	var f map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&f)
	return true
}
