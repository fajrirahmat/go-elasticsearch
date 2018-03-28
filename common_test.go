package client

import (
	"testing"
)

//shouldn't need to test private method. but I need it.. :D

func TestGenerateDeleteByQueryURLString_withOneIndex(t *testing.T) {
	indexname := []string{"gotest"}
	url := generateDeleteByQueryURLString(indexname)
	if url != "/gotest/_delete_by_query" {
		t.Fail()
	}
}

func TestGenerateDeleteByQueryWithTypeURLString_withOneIndexAndType(t *testing.T) {
	indexname := []string{"gotest"}
	docType := []string{"gotype"}
	url := generateDeleteByQueryWithTypeURLString(indexname, docType)
	if url != "/gotest/gotype/_delete_by_query" {
		t.Fail()
	}
}

func TestGenerateDeleteByQueryURLString_withMultiIndex(t *testing.T) {
	indexname := []string{"gotest", "goper", "golang"}
	url := generateDeleteByQueryURLString(indexname)
	if url != "/gotest,goper,golang/_delete_by_query" {
		t.Fail()
	}
}

func TestGenerateDeleteWithTypeByQueryURLString_withMultiIndex(t *testing.T) {
	indexname := []string{"gotest", "goper", "golang"}
	docType := []string{"gotype", "godoc"}
	url := generateDeleteByQueryWithTypeURLString(indexname, docType)
	if url != "/gotest,goper,golang/gotype,godoc/_delete_by_query" {
		t.Fail()
	}
}
