package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type SearchRequest struct {
	IndexName    string
	DocumentType string
	Query        interface{}
}

type SearchResponse struct {
	Took     int64              `json:"took"`
	TimedOut bool               `json:"timed_out"`
	Shards   Shard              `json:"_shards"`
	Hits     HitsSearchResponse `json:"hits"`
}

type HitsSearchResponse struct {
	Total    int64   `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hits  `json:"hits"`
}

type Hits struct {
	HitsHeader
	Score  float64     `json:"_score"`
	Source interface{} `json:"_source"`
}

func (c *Context) Search(searchRequest *SearchRequest) (SearchResponse, error) {
	var b []byte
	if searchRequest.Query != nil {
		b, _ = json.Marshal(searchRequest.Query)
	}
	request, _ := http.NewRequest("GET", generateSearchURLString(*searchRequest), bytes.NewBuffer(b))
	resp, err := c.C.Do(request)
	if err != nil {
		return SearchResponse{}, err
	}
	defer resp.Body.Close()

	var f SearchResponse
	json.NewDecoder(resp.Body).Decode(&f)

	return f, nil

}

func generateSearchURLString(searchRequest SearchRequest) string {
	urlString := "/"
	if searchRequest.IndexName != "" {
		urlString += searchRequest.IndexName + "/"
	}

	if searchRequest.DocumentType != "" {
		urlString += searchRequest.DocumentType + "/"
	}
	return urlString + "_search"
}
