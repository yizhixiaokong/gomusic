package service

import (
	"gomusic/common"
	"gomusic/dao"
	"gomusic/pkg/logger"
	"gomusic/serializer"
)

// GetMusicListService 获取歌曲列表的服务
type GetMusicListService struct {
	common.Page
	MusicName        string `form:"musicName" json:"musicName"`               //歌名
	OriginalSinger   string `form:"originalSinger" json:"originalSinger"`     //原唱
	Language         string `form:"language" json:"language"`                 //语言
	MusicType        string `form:"musicType" json:"musicType"`               //分类
	ProficiencyLevel string `form:"proficiencyLevel" json:"proficiencyLevel"` //熟练度
	DesignateBy      string `form:"designateBy" json:"designateBy"`           //指定者
}

// GetMusicList 获取歌曲列表服务
func (service *GetMusicListService) GetMusicList() (serializer.GetMusicList, common.WebError) {
	total, musics, err := dao.GetMusicList(nil, service.Page,
		service.MusicName,
		service.OriginalSinger,
		service.Language,
		service.MusicType,
		service.ProficiencyLevel,
		service.DesignateBy)
	if err != nil {
		logger.Errorp(service, err)
		return serializer.GetMusicList{}, common.ErrServer()
	}

	return serializer.BuildMusicList(musics, total), nil
}
