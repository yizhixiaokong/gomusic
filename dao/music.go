package dao

import (
	"gomusic/common"
	"gomusic/model"

	"github.com/jinzhu/gorm"
)

func GetMusicList(tx *gorm.DB, page common.Page, musicName, originalSinger, language, musicType, proficiencyLevel, designateBy string) (total int, items []model.Music, err error) {
	if tx == nil {
		tx = model.GetDB()
	}

	name := model.Music{}.TableName()
	tx = tx.Table(name)

	where := map[string]interface{}{}

	if musicName != "" {
		where["music_name"] = musicName
	}

	if originalSinger != "" {
		where["original_singer"] = originalSinger
	}

	if language != "" {
		where["language"] = language
	}

	if musicType != "" {
		where["music_type"] = musicType
	}

	if proficiencyLevel != "" {
		where["proficiency_level"] = proficiencyLevel
	}

	if designateBy != "" {
		where["designate_by"] = designateBy
	}

	err = tx.Where(where).Where("`deleted_at` IS NULL").Count(&total).Error //total
	if total == 0 {
		return total, items, err
	}

	//分页
	tx = tx.
		Order(page.Order("id")).
		Offset(page.Offset()).
		Limit(page.Limit())

	err = tx.Where(where).Find(&items).Error

	return
}

//通过id查找
func GetMusic(tx *gorm.DB, id uint) (item model.Music, err error) {
	if tx == nil {
		tx = model.GetDB()
	}

	tableName := model.Music{}.TableName()
	tx = tx.Table(tableName)

	where := map[string]interface{}{"id": id}

	err = tx.Where(where).First(&item).Error
	return
}

//通过name查找
func GetMusicByName(tx *gorm.DB, name string) (item model.Music, err error) {
	if tx == nil {
		tx = model.GetDB()
	}

	tableName := model.Music{}.TableName()
	tx = tx.Table(tableName)

	where := map[string]interface{}{"music_name": name}

	err = tx.Where(where).First(&item).Error
	return
}

func GetRandomMusic(tx *gorm.DB, count int) (items []model.Music, err error) {
	if tx == nil {
		tx = model.GetDB()
	}

	name := model.Music{}.TableName()
	tx = tx.Table(name)

	err = tx.Order("rand()").Limit(count).Find(&items).Error
	return
}
