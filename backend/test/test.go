package main

import (
	"encoding/json"
	"fmt"
)

type sortRule struct {
	Column string `json:"column"`
	Order string `json:"order"`
}
type filterRule struct {
	Identity []string `json:"identity"`
	Privilege []string `json:"privilege"`
	Grade []string `json:"grade"`
}
type recordRule struct {
	Size string `json:"size"`
	Offset string `json:"offset"`
	Keyword string `json:"keyword"`
}
type queryRecordInfo struct {
	SortRule sortRule `json:"sortRule"`
	FilterRule filterRule `json:"filterRule"`
	RecordRule recordRule `json:"recordRule"`
}

func main() {
	js := `{"sortRule":{"column":"","order":""},"filterRule":{"identity":["1","2"],"privilege":[],"grade":[]},"recordRule":{"size":10,"offset":0,"keyword":""}}`
	var xm queryRecordInfo
	_ = json.Unmarshal([]byte(js), &xm)
	fmt.Println(xm)
}