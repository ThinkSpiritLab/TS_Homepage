package interfaceDataStruct

type sortRule struct {
	Column string `json:"column"`
	Order string `json:"order"`
}
type filterRule struct {
	Identity []int `json:"identity"`
	Privilege []int `json:"privilege"`
	Grade []string `json:"grade"`
}
type recordRule struct {
	Size int `json:"size"`
	Offset int `json:"offset"`
	Keyword string `json:"keyword"`
}
type QueryRecordInfo struct {
	SortRule sortRule `json:"sortRule"`
	FilterRule filterRule `json:"filterRule"`
	RecordRule recordRule `json:"recordRule"`
}
type Auth struct {
	Stid string `valid:"Required; MaxSize(50)"`
	Psw string `valid:"Required; MaxSize(50)"`
}
type GradeInterface struct {
	Text string `json:"text"`
	Value string `json:"value"`
}