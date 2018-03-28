package client

import (
	"fmt"
	"strings"
)

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
