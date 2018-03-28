package client

import (
	"fmt"
	"strings"
)

//TODO
//find the best way to generate URL

//generateDeleteByQueryURLString to generate url string to delete document by query
//Indexname pass as parameter of function
//Return URL string
func generateDeleteByQueryURLString(indexName []string) string {
	return "/" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(indexName)), ","), "[]") + "/_delete_by_query"
}

//generateDeleteByQueryWithTypeURLString to generate url string to delete document by query
//Indexname and Document type pass as parameters of function
//Return URL string
func generateDeleteByQueryWithTypeURLString(indexName []string, documentType []string) string {
	return "/" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(indexName)), ","), "[]") + "/" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(documentType)), ","), "[]") + "/_delete_by_query"
}

func generateURLStringWithProceedConflict(urlstring string) string {
	return urlstring + "?conflicts=proceed"
}

func generateURLStringWithRouting(urlstring string, routing string) string {
	return urlstring + "?routing=" + routing
}

func generateURLStringWithScrollSize(urlstring string, size int) string {
	return urlstring + "?scroll_size=" + string(size)
}

func generateUpdateURLString(u UpdateDocument) string {
	url := "/"
	if index := writeOnelineArray(u.IndexName); index != "" {
		url += index
	}

	if docType := writeOnelineArray(u.DocumentType); docType != "" {
		url += ("/" + docType)
	}

	if docID := writeOnelineArray(u.DocumentID); docID != "" {
		url += ("/" + docID)
	}

	if u.Query == nil {
		url += "/_update"
	} else {
		url += "/_update_by_query"
	}
	return url
}

func writeOnelineArray(arr []string) string {
	if len(arr) > 0 {
		return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), ","), "[]")
	}
	return ""
}
