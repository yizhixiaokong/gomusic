package service

import (
	"gomusic/common"
	"gomusic/dao"
	"gomusic/model"
	"gomusic/pkg/logger"
	"gomusic/serializer"
)

type PostBatchMusicService struct {
	List []PostMusicService `form:"list" json:"list" binding:"required,gte=1"`
}

// PostBatchMusic 批量添加歌曲服务
func (service *PostBatchMusicService) PostBatchMusic() (failList serializer.FailMusicList, e common.WebError) {
	tx, err := dao.TxBegin()
	if err != nil {
		logger.Errorf("tx: %+v, err: %v", tx, err)
		e = common.ErrServer()
		return serializer.FailMusicList{}, e
	}
	defer func() {
		if e != nil {
			tx.Rollback()
		} else {
			if err = tx.Commit().Error; err != nil {
				e = common.ErrServer()
			}
		}
	}()
	success := 0
	fails := []serializer.Music{}
	for _, one := range service.List {
		var exist bool
		var failFlag = false
		_, err := dao.GetMusicByName(tx, one.MusicName)
		exist, err = dao.ExistRow(err)
		if exist {
			failFlag = true
		}
		if err != nil && !failFlag {
			failFlag = true
		}

		if !failFlag {
			music := model.Music{
				MusicName:        one.MusicName,
				OriginalSinger:   one.OriginalSinger,
				Language:         one.Language,
				MusicType:        one.MusicType,
				ProficiencyLevel: one.ProficiencyLevel,
				DesignateBy:      one.DesignateBy,
			}

			err = dao.TxCreate(tx, &music)
			if err != nil {
				failFlag = true
			}

			if !failFlag {
				success++
			}
		}

		if failFlag {
			fail := serializer.Music{
				ID:               0,
				MusicName:        one.MusicName,
				OriginalSinger:   one.OriginalSinger,
				Language:         one.Language,
				MusicType:        one.MusicType,
				ProficiencyLevel: one.ProficiencyLevel,
				DesignateBy:      one.DesignateBy,
			}
			fails = append(fails, fail)
		}
	}

	failList = serializer.FailMusicList{
		Success:  success,
		FailList: fails,
	}

	return failList, nil
}
