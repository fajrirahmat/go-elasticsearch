package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type UpdateDocument struct {
	IndexName    []string
	DocumentType []string
	DocumentID   []string
	Body         interface{}
	Query        interface{}
}

func (c *Context) Update(updateDoc *UpdateDocument) bool {
	if updateDoc.DocumentID == nil && updateDoc.Query == nil {
		log.Println("Note allowed")
		return false
	}
	url := generateUpdateURLString(*updateDoc)

	b, err := json.Marshal(updateDoc.Body)
	if err != nil {
		log.Println(err)
		return false
	}

	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))

	resp, err := c.C.Do(request)

	if err != nil {
		log.Println(err)
		return false
	}

	log.Println(resp)

	return true
}
