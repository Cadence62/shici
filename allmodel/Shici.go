package allmodel

import "github.com/jinzhu/gorm"

type Shi struct {
	gorm.Model
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content" gorm:"type:text"`
}

type Ci struct {
	gorm.Model
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content" gorm:"type:text"`
}

type Yan struct {
	gorm.Model
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content" gorm:"type:text"`
}
