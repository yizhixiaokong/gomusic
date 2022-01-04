package service

import (
	"gomusic/common"
	"gomusic/dao"
	"gomusic/pkg/logger"
	"gomusic/serializer"
)

type GetRandomMusicService struct {
	Count int `form:"count" json:"count" binding:"gte=1,lte=30"`
}

// GetRandomMusic 获取随机若干首歌曲服务
func (service *GetRandomMusicService) GetRandomMusic() (serializer.GetMusicList, common.WebError) {
	count := service.Count
	musics, err := dao.GetRandomMusic(nil, count)
	if err != nil {
		logger.Errorp(service, err)
		return serializer.GetMusicList{}, common.ErrServer()
	}

	return serializer.BuildMusicList(musics, len(musics)), nil
}
