package service

import (
	"gomusic/common"
	"gomusic/dao"
	"gomusic/pkg/logger"
	"gomusic/serializer"
)

type PutMusicService struct {
	ID               uint   `form:"id" json:"id"`
	MusicName        string `form:"musicName" json:"musicName"`               //歌名
	OriginalSinger   string `form:"originalSinger" json:"originalSinger"`     //原唱
	Language         string `form:"language" json:"language"`                 //语言
	MusicType        string `form:"musicType" json:"musicType"`               //分类
	ProficiencyLevel string `form:"proficiencyLevel" json:"proficiencyLevel"` //熟练度
	DesignateBy      string `form:"designateBy" json:"designateBy"`           //指定者
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
	if exist {
		return serializer.Music{}, common.ErrIsExist().AddMsg(" :歌曲name已存在")
	}
	if err != nil {
		logger.Errorp(service, err)
		return serializer.Music{}, common.ErrServer()
	}

	music.MusicName = service.MusicName
	music.OriginalSinger = service.OriginalSinger
	music.Language = service.Language
	music.MusicType = service.MusicType
	music.ProficiencyLevel = service.ProficiencyLevel
	music.DesignateBy = service.DesignateBy

	err = dao.TxSave(nil, &music)
	if err != nil {
		logger.Errorp(service, err)
		return serializer.Music{}, common.ErrServer()
	}

	return serializer.BuildMusic(music), nil
}
