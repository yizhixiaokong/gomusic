package service

import (
	"gomusic/common"
	"gomusic/dao"
	"gomusic/pkg/logger"
	"gomusic/serializer"
)

type GetMusicService struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

// GetMusic 获取单首歌曲服务
func (service *GetMusicService) GetMusic() (serializer.Music, common.WebError) {
	id := service.ID
	var exist bool
	music, err := dao.GetMusic(nil, id)
	exist, err = dao.ExistRow(err)
	if !exist {
		return serializer.Music{}, common.ErrNotExist().AddMsg(" :歌曲id不存在")
	}
	if err != nil {
		logger.Errorp(service, err)
		return serializer.Music{}, common.ErrServer()
	}

	return serializer.BuildMusic(music), nil
}
