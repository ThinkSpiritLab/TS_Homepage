package models

type IndexImg struct {
	Id             int       `gorm:"primary_key" json:"name"`
	Url            string    `json:"url"`
}

func GetIndexImgList() (imgs []IndexImg) {
	db.Order("url asc").Find(&imgs)
	return
}

func SetIndexImgList(imgs []IndexImg) {
	db.Delete(IndexImg{})
	for _, val := range imgs {
		db.Create(&val)
	}
}