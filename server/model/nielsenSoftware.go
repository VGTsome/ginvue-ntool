// 自动生成模板NielsenSoftware
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type NielsenSoftware struct {
      gorm.Model
      SoftName  string `json:"softName" form:"softName" gorm:"column:soft_name;comment:'';type:varchar(255)"`
      SoftDescription  string `json:"softDescription" form:"softDescription" gorm:"column:soft_description;comment:'';type:varchar(2000)"`
      Download  string `json:"download" form:"download" gorm:"column:download;comment:'';type:varchar(2000)"`
      SoftImg  string `json:"softImg" form:"softImg" gorm:"column:soft_img;comment:'';type:varchar(255)"`
      Version  string `json:"version" form:"version" gorm:"column:version;comment:'';type:varchar(255)"` 
}


func (NielsenSoftware) TableName() string {
  return "nielsen_software"
}
