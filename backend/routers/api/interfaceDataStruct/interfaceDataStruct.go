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

type BaseInfo struct {
	NameZh     string `json:"c_name_zh"`
	NameEn     string `json:"c_name_en"`
	LocationZH string `json:"c_location_zh"`
	LocationEN string `json:"c_location_en"`
	Ctime      string `json:"c_time"`
}

type Result struct {
	NameZh   string `json:"t_name_zh"`
	NameEn   string `json:"t_name_en"`
	Mem1Stid string `json:"t_mem1_stid"`
	Mem2Stid string `json:"t_mem2_stid"`
	Mem3Stid string `json:"t_mem3_stid"`
	Star     string `json:"t_star"`
	Rank     int    `json:"t_rank"`
	Awards   string `json:"t_awards"`
}

type ResultInfo struct {
	AwardType  string `json:"r_type"`
	TotalTeams int    `json:"total_team_num"`
	Results []Result  `json:"results"`
}

type ContestAddForm struct {
	BaseInfo   BaseInfo   `json:"baseInfo"`
	ResultInfo ResultInfo `json:"resultInfo"`
	Extras    string     `json:"extras"`
}