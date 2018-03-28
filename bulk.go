package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

//BulkHeader header of bulk item data
type BulkHeader struct {
	IndexName string `json:"_index"`
	Type      string `json:"_type"`
	ID        string `json:"_id"`
}

//BulkItem bulk item
type BulkItem struct {
	Process string
	Header  BulkHeader
	Source  map[string]interface{}
}

//constants of bulk process type
const (
	INDEX  = "index"
	DELETE = "delete"
	CREATE = "create"
	UPDATE = "update"
)

//BulkProcessRequest data type for bulk process
type BulkProcessRequest struct {
	Items []BulkItem
}

//Bulk bulk processing
//Reference: https://www.elastic.co/guide/en/elasticsearch/reference/5.5/docs-bulk.html
func (c *Context) Bulk(items BulkProcessRequest) interface{} {
	reqBodyString, err := generateBulkRequestBody(items)
	if err != nil {
		return nil
	}
	request, _ := http.NewRequest("POST", "/_bulk", bytes.NewBufferString(reqBodyString))
	resp, err := c.C.Do(request)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		log.Println(err)
		return nil
	}

	return result
}

func generateBulkRequestBody(items BulkProcessRequest) (string, error) {
	reqBody := ""
	for _, v := range items.Items {
		switch v.Process {
		case INDEX:
			header := map[string]interface{}{
				"index": v.Header,
			}
			h, _ := json.Marshal(header)
			reqBody += fmt.Sprintf("%s\n", string(h))
			s, _ := json.Marshal(v.Source)
			reqBody += fmt.Sprintf("%s\n", string(s))
		case DELETE:
			header := map[string]interface{}{
				"delete": v.Header,
			}
			b, _ := json.Marshal(header)
			reqBody += fmt.Sprintf("%s\n", string(b))
		case CREATE:
			header := map[string]interface{}{
				"create": v.Header,
			}
			h, _ := json.Marshal(header)
			reqBody += fmt.Sprintf("%s\n", string(h))
			s, _ := json.Marshal(v.Source)
			reqBody += fmt.Sprintf("%s\n", string(s))
		case UPDATE:
			header := map[string]interface{}{
				"update": v.Header,
			}
			h, _ := json.Marshal(header)
			reqBody += fmt.Sprintf("%s\n", string(h))
			s, _ := json.Marshal(v.Source)
			reqBody += fmt.Sprintf("%s\n", string(s))
		default:
			return reqBody, errors.New("Unsupport Bulk Process Type")
		}
	}
	return reqBody, nil
}
