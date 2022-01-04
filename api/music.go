package api

import (
	"gomusic/common"
	"gomusic/pkg/logger"
	"gomusic/service"

	"github.com/gin-gonic/gin"
)

// 获取歌曲列表
func GetMusicList(c *gin.Context) {
	var service service.GetMusicListService
	if err := c.ShouldBind(&service); err != nil {
		logger.Info(err.Error())
		common.ResJson(c, nil, common.ErrInvalidParams(err))
		return
	}
	if err := service.Page.Check(); err != nil {
		logger.Info(err.Error())
		common.ResJson(c, nil, err)
		return
	}
	res, err := service.GetMusicList()
	common.ResJson(c, res, err)
}

// 获取随机若干首歌曲
func GetRandomMusic(c *gin.Context) {
	var service service.GetRandomMusicService
	if err := c.ShouldBind(&service); err != nil {
		logger.Info(err.Error())
		common.ResJson(c, nil, common.ErrInvalidParams(err))
		return
	}
	res, err := service.GetRandomMusic()
	common.ResJson(c, res, err)
}

// 添加歌曲
func PostMusic(c *gin.Context) {
	var service service.PostMusicService
	if err := c.ShouldBind(&service); err != nil {
		logger.Info(err.Error())
		common.ResJson(c, nil, common.ErrInvalidParams(err))
		return
	}
	res, err := service.PostMusic()
	common.ResJson(c, res, err)
}

// 修改歌曲
func PutMusic(c *gin.Context) {
	var service service.PutMusicService
	if err := c.ShouldBind(&service); err != nil {
		logger.Info(err.Error())
		common.ResJson(c, nil, common.ErrInvalidParams(err))
		return
	}
	res, err := service.PutMusic()
	common.ResJson(c, res, err)
}

// 删除歌曲
func DeleteMusic(c *gin.Context) {
	var service service.DeleteMusicService
	if err := c.ShouldBind(&service); err != nil {
		logger.Info(err.Error())
		common.ResJson(c, nil, common.ErrInvalidParams(err))
		return
	}
	err := service.DeleteMusic()
	common.ResJson(c, nil, err)
}

// 批量添加歌曲
func PostBatchMusic(c *gin.Context) {
	var service service.PostBatchMusicService
	if err := c.ShouldBind(&service); err != nil {
		logger.Info(err.Error())
		common.ResJson(c, nil, common.ErrInvalidParams(err))
		return
	}
	res, err := service.PostBatchMusic()
	common.ResJson(c, res, err)
}

// 批量删除歌曲
func DeleteBatchMusic(c *gin.Context) {
	var service service.DeleteBatchMusicService
	if err := c.ShouldBind(&service); err != nil {
		logger.Info(err.Error())
		common.ResJson(c, nil, common.ErrInvalidParams(err))
		return
	}
	err := service.DeleteBatchMusic()
	common.ResJson(c, nil, err)
}
