package model

import (
	"github.com/jinzhu/gorm"
)

// Music 歌曲模型
type Music struct {
	gorm.Model
	MusicName        string //歌名
	OriginalSinger   string //原唱
	Language         string //语言
	MusicType        string //分类
	ProficiencyLevel string //熟练度
	DesignateBy      string //指定者
}

// TableName 表名
func (Music) TableName() string {
	return "music"
}
