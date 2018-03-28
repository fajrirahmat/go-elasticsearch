package client

import (
	"fmt"
	"testing"
)

func TestInfo(t *testing.T) {
	c := New("http://localhost:9200")
	idx, err := c.Info()
	if err != nil {
		t.Fatal()
	}
	fmt.Println(idx)

}

type Gotest struct {
	Message string `json:"message"`
}

func TestCreateIndex(t *testing.T) {
	c := New("http://localhost:9200")
	data := &Gotest{Message: "Testing dulu lah"}
	resp, _ := c.CreateIndex("gotest", "gotest", *data)
	if resp.ID == "" {
		t.Fail()
	}
}

func TestCreateIndexWithID(t *testing.T) {
	c := New("http://localhost:9200")
	data := &Gotest{Message: "Testing dulu lah"}
	resp, _ := c.CreateIndexWithID("gotest", "gotest", "1324", *data)
	if resp.ID == "" {
		t.Fail()
	}
}

func TestCreateIndexWithIDAndVersion(t *testing.T) {
	c := New("http://localhost:9200")
	data := &Gotest{Message: "Testing dulu lah"}
	resp, err := c.CreateIndexWithIDAndVersion("gotest", "gotest", "1324", 4, *data)
	if err != nil {
		t.Fail()
	}

	fmt.Println(resp.Result)
}

func TestGetIndexWithTypeAndID(t *testing.T) {
	c := New("http://localhost:9200")
	resp, err := c.GetIndexWithTypeAndID("gotest", "gotest", "AWJmRO8NhMSe7XqCM5Vt")
	if err != nil {
		t.Fail()
	}
	if !resp.Found {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	c := New("http://localhost:9200")
	data := &Gotest{Message: "Testing dulu lah"}
	c.CreateIndexWithID("gotest", "gotest", "AWJmRO8NhMSe7XqCM5Vt", *data)
	if !c.Delete("gotest", "gotest", "AWJmRO8NhMSe7XqCM5Vt") {
		t.Fail()
	}
}

func TestDeleteByQuery(t *testing.T) {
	c := New("http://localhost:9200")
	data := &Gotest{Message: "Testing dulu lah"}
	c.CreateIndexWithID("gotest", "gotest", "AWJmRO8NhMSe7XqCM5Vt", *data)
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"_id": "AWJmRO8NhMSe7XqCM5Vt",
			},
		},
	}
	if !c.DeleteByQuery([]string{"gotest"}, query) {
		t.Fail()
	}
}

func TestDeleteByQuery_NotExist(t *testing.T) {
	c := New("http://localhost:9200")
	//data := &Gotest{Message: "Testing dulu lah"}
	//c.CreateIndexWithID("gotest", "gotest", "AWJmRO8NhMSe7XqCM5Vt", *data)
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"_id": "AWJmRO8NhMSe7XqCM5V",
			},
		},
	}
	if c.DeleteByQuery([]string{"gotest"}, query) {
		t.Fail()
	}
}

func TestUpdate(t *testing.T) {
	c := New("http://localhost:9200")
	script := map[string]interface{}{
		"script": map[string]interface{}{
			"inline": "ctx._source.message='Ini update yang baru loh ads'",
		},
	}
	u := &UpdateDocument{
		Body:         script,
		DocumentID:   []string{"AWJnFD6ThMSe7XqCM5Vz"},
		DocumentType: []string{"gotest"},
		IndexName:    []string{"gotest"},
	}

	if !c.Update(u) {
		t.Fail()
	}
}

func TestMultiGet(t *testing.T) {
	c := New("http://localhost:9200")
	docs := []DocumentHeader{
		DocumentHeader{
			ID:    "AWJnFD6ThMSe7XqCM5Vz",
			Index: "gotest",
		},
		DocumentHeader{
			ID:    "AWJnFHaxhMSe7XqCM5V0",
			Index: "gotest",
		},
	}
	c.MultiGet(docs)
}

func TestGenerateBulkRequestBody(t *testing.T) {
	pq := &BulkProcessRequest{
		Items: []BulkItem{
			BulkItem{
				Process: INDEX,
				Header:  HitsHeader{ID: "BULK1", IndexName: "gotest", Type: "gotest"},
				Source: map[string]interface{}{
					"field1": "value1",
					"field2": "value2",
				},
			},
			BulkItem{
				Process: DELETE,
				Header:  HitsHeader{ID: "AWJnFD6ThMSe7XqCM5Vz", IndexName: "gotest", Type: "gotest"},
			},
		},
	}
	str, _ := generateBulkRequestBody(*pq)
	if str == "" {
		t.Fail()
	}
}

func TestBulk(t *testing.T) {
	c := New("http://localhost:9200")
	pq := &BulkProcessRequest{
		Items: []BulkItem{
			BulkItem{
				Process: INDEX,
				Header:  HitsHeader{ID: "BULK1", IndexName: "gotest", Type: "gotest"},
				Source: map[string]interface{}{
					"field1": "value1",
					"field2": "value2",
				},
			},
			BulkItem{
				Process: DELETE,
				Header:  HitsHeader{ID: "BULK1", IndexName: "gotest", Type: "gotest"},
			},
		},
	}
	if c.Bulk(*pq) == nil {
		t.Fail()
	}
}

func TestBulk_UnknonwProcess(t *testing.T) {
	c := New("http://localhost:9200")
	pq := &BulkProcessRequest{
		Items: []BulkItem{
			BulkItem{
				Process: "contek",
				Header:  HitsHeader{ID: "BULK1", IndexName: "gotest", Type: "gotest"},
				Source: map[string]interface{}{
					"field1": "value1",
					"field2": "value2",
				},
			},
		},
	}
	if c.Bulk(*pq) != nil {
		t.Fail()
	}
}

func TestSearch(t *testing.T) {
	c := New("http://localhost:9200")
	q := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"_id": "AWJnFHaxhMSe7XqCM5V0",
			},
		},
	}
	s := &SearchRequest{
		IndexName:    "gotest",
		DocumentType: "gotest",
		Query:        q,
	}

	response, _ := c.Search(s)
	//assume the data with the index and _id has been in Elasticsearch
	if response.Hits.Total == 0 {
		t.Fail()
	}

}
