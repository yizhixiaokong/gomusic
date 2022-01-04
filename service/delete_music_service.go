package service

import (
	"gomusic/common"
	"gomusic/dao"
	"gomusic/pkg/logger"
)

type DeleteMusicService struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

// DeleteMusic 删除歌曲服务
func (service *DeleteMusicService) DeleteMusic() common.WebError {

	var exist bool
	music, err := dao.GetMusic(nil, service.ID)
	exist, err = dao.ExistRow(err)
	if !exist {
		return common.ErrNotExist().AddMsg(" :歌曲id不存在")
	}
	if err != nil {
		logger.Errorp(service, err)
		return common.ErrServer()
	}

	_, err = dao.DeleteByID(&music, int64(music.ID))
	if err != nil {
		logger.Errorp(service, err)
		return common.ErrServer()
	}
	return nil
}
