package client

//Index API
//Reference: https://www.elastic.co/guide/en/elasticsearch/reference/5.5/docs-index_.html

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type CreateIndexResponse struct {
	Shards  Shard  `json:"_shards"`
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	ID      string `json:"_id"`
	Version string `json:"_version"`
	Created bool   `json:"created"`
	Result  string `json:"result"`
}

type Shard struct {
	Total      int `json:"total"`
	Failed     int `json:"failed"`
	SuccessFul int `json:"successful"`
}

type Info struct {
	Name        string  `json:"name"`
	ClusterName string  `json:"cluster_name"`
	ClusterUUID string  `json:"cluster_uuid"`
	Version     Version `json:"version"`
	TagLine     string  `json:"tagline"`
}

type Version struct {
	Number        string    `json:"number"`
	BuildHash     string    `json:"build_hash"`
	BuildDate     time.Time `json:"build_date"`
	BuildSnapshot bool      `json:"build_snapshot"`
	LuceneVersion string    `json:"lucene_version"`
}

//Info Get Elasticsearch Info
//By request using GET method to root index ("/") of Elasticsearch
func (i *Context) Info() (Info, error) {
	resp, err := i.C.Get("/")
	if err != nil {
		return Info{}, err
	}
	if resp.StatusCode != 200 {
		return Info{}, errors.New(resp.Status)
	}

	defer resp.Body.Close()
	var info Info
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return Info{}, err
	}

	return info, nil
}

//CreateIndex to create index with Auto-Generation document ID
func (i *Context) CreateIndex(indexName string, documentType string, data interface{}) (CreateIndexResponse, error) {
	response := CreateIndexResponse{}
	url := "/" + indexName + "/" + documentType + "/"
	b, err := json.Marshal(&data)
	if err != nil {
		return response, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return response, err
	}

	resp, err := i.C.Do(request)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)

	return response, err
}

//CreateIndexWithRouting to create index with Auto-Generation document ID and routing defined
func (i *Context) CreateIndexWithRouting(indexName string, documentType string, routing string, data interface{}) (CreateIndexResponse, error) {
	response := CreateIndexResponse{}
	url := "/" + indexName + "/" + documentType + "?routing=" + routing
	b, err := json.Marshal(&data)
	if err != nil {
		return response, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return response, err
	}

	resp, err := i.C.Do(request)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)

	return response, err
}

//CreateIndexWithID Create index with specific indexname, documenttype and defined document ID
func (i *Context) CreateIndexWithID(indexName string, documentType string, documentID string, data interface{}) (CreateIndexResponse, error) {
	response := CreateIndexResponse{}
	url := "/" + indexName + "/" + documentType + "/" + documentID
	b, err := json.Marshal(&data)
	if err != nil {
		return response, err
	}
	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(b))
	if err != nil {
		return response, err
	}

	resp, err := i.C.Do(request)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)

	return response, err
}

//CreateIndexWithIDAndVersion Create index with specific indexname, documenttype and defined document ID and alson set version
func (i *Context) CreateIndexWithIDAndVersion(indexName string, documentType string, documentID string, version int64, data interface{}) (CreateIndexResponse, error) {
	response := CreateIndexResponse{}
	url := "/" + indexName + "/" + documentType + "/" + documentID + "?version=" + string(version)
	b, err := json.Marshal(&data)
	if err != nil {
		return response, err
	}
	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(b))
	if err != nil {
		return response, err
	}

	resp, err := i.C.Do(request)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)

	return response, err
}
