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
	if !c.DeleteByQuery("gotest", query) {
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
	if c.DeleteByQuery("gotest", query) {
		t.Fail()
	}
}
