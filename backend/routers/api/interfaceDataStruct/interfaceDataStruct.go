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
type RecordRule struct {
	Size int `json:"size"`
	Offset int `json:"offset"`
	Keyword string `json:"keyword"`
}
type QueryRecordInfo struct {
	SortRule   sortRule   `json:"sortRule"`
	FilterRule filterRule `json:"filterRule"`
	RecordRule RecordRule `json:"recordRule"`
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

type ContestEditForm struct {
	Cid       int        `json:"cid"`
	Extras    string     `json:"extras"`
}

type ContestListBrief struct {
	Cid        int    `json:"c_id"`
	NameZh     string `json:"c_name_zh"`
	NameEn     string `json:"c_name_en"`
	LocationZH string `json:"c_location_zh"`
	LocationEN string `json:"c_location_en"`
	Ctime      string `json:"c_time"`
	AwardType  string `json:"r_type"`
	Num0       int    `json:"num_0"`
	Num1       int    `json:"num_1"`
	Num2       int    `json:"num_2"`
	Num3       int    `json:"num_3"`
	Num4       int    `json:"num_4"`
}

type BulletinAddForm struct {
	BulletinTitle  string `json:"bulletinTitle"`
	BulletinDetail string `json:"bulletinDetail"`
}

type BulletinListBrief struct {
	Bid        int    `json:"b_id"`
	Title      string `json:"title"`
	Date       string `json:"date"`
	Promulgator       string `json:"promulgator"`
}

type EditBulletinForm struct {
	Bid            int    `json:"bid"`
	BulletinTitle  string `json:"bulletinTitle"`
	BulletinDetail string `json:"bulletinDetail"`
}

type NewsAddForm struct {
	NewsTitle  string `json:"newsTitle"`
	NewsDetail string `json:"newsDetail"`
	NewsDate   string `json:"newsDate"`
}

type NewsListBrief struct {
	Nid          int    `json:"n_id"`
	Title        string `json:"title"`
	Date         string `json:"date"`
	Promulgator  string `json:"promulgator"`
}

type EditNewsForm struct {
	Nid         int    `json:"nid"`
	NewsTitle   string `json:"newsTitle"`
	NewsDetail  string `json:"newsDetail"`
	NewsDate    string `json:"newsDate"`
}

type UserAllInfo struct {
	Uid          int    `json:"uid"`
	Stid         string `json:"stid"`
	Name         string `json:"name"`
	Identity     int    `json:"identity"`
	Privilege    int    `json:"privilege"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	QQ           string `json:"QQ"`
	URL          string `json:"URL"`
	Introduction string `json:"introduction"`
	AvatarUrl    string `json:"avatarUrl"`
	Education    string `json:"education"`
	Career string `json:"career"`
}