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
//Reference: https://www.elastic.co/guide/en/elasticsearch/reference/5.6/docs-delete-by-query.html
func (c *Context) DeleteByQuery(indexName []string, query map[string]interface{}) bool {
	url := generateDeleteByQueryURLString(indexName)
	return deleteByQueryProcess(c, url, query)
}

//DeleteByQueryProceedConflict to delete elasticsearch document by specific query with proceed conflict
//Reference: https://www.elastic.co/guide/en/elasticsearch/reference/5.6/docs-delete-by-query.html
func (c *Context) DeleteByQueryProceedConflict(indexName []string, query map[string]interface{}) bool {
	url := generateURLStringWithProceedConflict(generateDeleteByQueryURLString(indexName))
	return deleteByQueryProcess(c, url, query)
}

func deleteByQueryProcess(c *Context, url string, query map[string]interface{}) bool {
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
