package service

import (
	"gomusic/common"
	"gomusic/dao"
	"gomusic/pkg/logger"
	"gomusic/serializer"
)

type PutMusicService struct {
	ID               uint   `form:"id" json:"id" binding:"required"`
	MusicName        string `form:"musicName" json:"musicName" binding:"required"`           //歌名
	OriginalSinger   string `form:"originalSinger" json:"originalSinger" binding:"required"` //原唱
	Language         string `form:"language" json:"language" binding:"required"`             //语言
	MusicType        string `form:"musicType" json:"musicType" binding:"required"`           //分类
	ProficiencyLevel string `form:"proficiencyLevel" json:"proficiencyLevel"`                //熟练度
	DesignateBy      string `form:"designateBy" json:"designateBy"`                          //指定者
}

// PutMusic 修改歌曲服务
func (service *PutMusicService) PutMusic() (serializer.Music, common.WebError) {
	var exist bool
	music, err := dao.GetMusic(nil, service.ID)
	exist, err = dao.ExistRow(err)
	if !exist {
		return serializer.Music{}, common.ErrNotExist().AddMsg(" :歌曲id不存在")
	}
	if err != nil {
		logger.Errorp(service, err)
		return serializer.Music{}, common.ErrServer()
	}

	_, err = dao.GetMusicByName(nil, service.MusicName)
	exist, err = dao.ExistRow(err)
	if exist && service.MusicName != music.MusicName {
		return serializer.Music{}, common.ErrIsExist().AddMsg(" :歌曲name已存在")
	}
	if err != nil {
		logger.Errorp(service, err)
		return serializer.Music{}, common.ErrServer()
	}

	if service.MusicName != "" {
		music.MusicName = service.MusicName
	}
	if service.OriginalSinger != "" {
		music.OriginalSinger = service.OriginalSinger
	}
	if service.Language != "" {
		music.Language = service.Language
	}
	if service.MusicType != "" {
		music.MusicType = service.MusicType
	}
	music.ProficiencyLevel = service.ProficiencyLevel
	music.DesignateBy = service.DesignateBy

	err = dao.TxSave(nil, &music)
	if err != nil {
		logger.Errorp(service, err)
		return serializer.Music{}, common.ErrServer()
	}

	return serializer.BuildMusic(music), nil
}
