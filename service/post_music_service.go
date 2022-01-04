package service

import (
	"gomusic/common"
	"gomusic/dao"
	"gomusic/model"
	"gomusic/pkg/logger"
	"gomusic/serializer"
)

type PostMusicService struct {
	MusicName        string `form:"musicName" json:"musicName" binding:"required"`           //歌名
	OriginalSinger   string `form:"originalSinger" json:"originalSinger" binding:"required"` //原唱
	Language         string `form:"language" json:"language" binding:"required"`             //语言
	MusicType        string `form:"musicType" json:"musicType" binding:"required"`           //分类
	ProficiencyLevel string `form:"proficiencyLevel" json:"proficiencyLevel"`                //熟练度
	DesignateBy      string `form:"designateBy" json:"designateBy"`                          //指定者
}

// PostMusic 添加歌曲服务
func (service *PostMusicService) PostMusic() (serializer.Music, common.WebError) {
	var exist bool
	_, err := dao.GetMusicByName(nil, service.MusicName)
	exist, err = dao.ExistRow(err)
	if exist {
		return serializer.Music{}, common.ErrIsExist().AddMsg(" :歌曲name已存在")
	}
	if err != nil {
		logger.Errorp(service, err)
		return serializer.Music{}, common.ErrServer()
	}

	music := model.Music{
		MusicName:        service.MusicName,
		OriginalSinger:   service.OriginalSinger,
		Language:         service.Language,
		MusicType:        service.MusicType,
		ProficiencyLevel: service.ProficiencyLevel,
		DesignateBy:      service.DesignateBy,
	}

	err = dao.TxCreate(nil, &music)
	if err != nil {
		logger.Errorp(service, err)
		return serializer.Music{}, common.ErrServer()
	}

	return serializer.BuildMusic(music), nil
}
