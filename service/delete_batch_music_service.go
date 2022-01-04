package service

import (
	"gomusic/common"
	"gomusic/dao"
	"gomusic/model"
	"gomusic/pkg/logger"
)

type DeleteBatchMusicService struct {
	IDs []uint `json:"ids"`
}

// DeleteBatchMusic 批量删除歌曲服务
func (service *DeleteBatchMusicService) DeleteBatchMusic() common.WebError {
	var res []int64

	for _, one := range service.IDs {
		res = append(res, int64(one))
	}

	_, err := dao.DeleteByIDS(&model.Music{}, res)
	if err != nil {
		logger.Errorp(service, err)
		return common.ErrServer()
	}

	return nil
}
