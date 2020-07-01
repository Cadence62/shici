package allmodel

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"shici/config"
	"time"
)

var (
	DB    *gorm.DB
	DBERR error
)

func InitDB() {
	sqlserver := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", config.PostgreIP,
		config.PostgrePort, config.PostgreUser, config.PostgreDbname, config.PostgreSSL, config.PostgrePassword)
	DB, DBERR = gorm.Open("postgres", sqlserver)
	if DBERR != nil {
		print("db error")
	}
	// 是否打印sql语句
	DB.LogMode(true)
	// 是否开启表名复数
	DB.SingularTable(true)
	//同步表结构
	DB.AutoMigrate(&Shi{}, &Ci{}, &Yan{})

}

type JsonTime time.Time

//实现它的json序列化方法
func (s JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(s).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type BaseModel struct {
	ID          uint      `gorm:"column:id" json:"id"`
	Create_date time.Time `gorm:"default:'null'" json:"-"`
	Delete_date time.Time `gorm:"default:'null'" json:"-"`
}

func (s *BaseModel) Create(value interface{}) error {
	s.Create_date = time.Now().UTC()
	return DB.Create(value).Error
}

func (s *BaseModel) FindALL(value interface{}) error {
	return DB.Model(value).Where("delete_date isnull").Find(value).Error
}

func (s *BaseModel) Delete(value interface{}, id int) error {
	delete_date := time.Now()
	return DB.Model(value).Where("id = ?", id).Update("delete_date", delete_date).Error
}

func (s *BaseModel) Update(value interface{}, id int) error {
	return DB.Model(value).Where("id = ?", id).Update(value).Error
}
